// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitoring

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/monitoring/metricsscope/apiv1"

	monitoringpb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringMonitoredProjectGVK, NewMonitoringMonitoredProjectModel)
}

func NewMonitoringMonitoredProjectModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &monitoringMonitoredProjectModel{config: *config}, nil
}

var _ directbase.Model = &monitoringMonitoredProjectModel{}

type monitoringMonitoredProjectModel struct {
	config config.ControllerConfig
}

func (m *monitoringMonitoredProjectModel) client(ctx context.Context) (*gcp.MetricsScopesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	metricsScopeClient, err := gcp.NewMetricsScopesClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MetricsScopes client: %w", err)
	}
	return metricsScopeClient, err
}

func (m *monitoringMonitoredProjectModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MonitoringMonitoredProject{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	metricsScope := direct.ValueOf(obj.Spec.MetricsScope)

	// The `monitored_project.name` must be in the format:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	monitoredProject := metricsScope + "/projects/" + obj.GetName()

	id, err := krm.MonitoringMonitoredProjectRefFromExternal(monitoredProject)
	if err != nil {
		return nil, err
	}

	// Get monitoring GCP client
	metricsScopeClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &monitoringMonitoredProjectAdapter{
		id:                 id,
		metricsScopeClient: metricsScopeClient,
		desired:            obj,
	}, nil
}

func (m *monitoringMonitoredProjectModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type monitoringMonitoredProjectAdapter struct {
	id *krm.MonitoringMonitoredProjectRef

	metricsScopeClient *gcp.MetricsScopesClient
	desired            *krm.MonitoringMonitoredProject
	actual             *monitoringpb.MonitoredProject
}

var _ directbase.Adapter = &monitoringMonitoredProjectAdapter{}

func (a *monitoringMonitoredProjectAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MonitoringMonitoredProject", "name", a.id.External)

	metricsScopeID, err := a.id.Parent()
	if err != nil {
		return false, fmt.Errorf("getting metricsScope: %w", err)
	}

	req := &monitoringpb.GetMetricsScopeRequest{Name: metricsScopeID.String()}
	metricsScope, err := a.metricsScopeClient.GetMetricsScope(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MonitoringMonitoredProject %q: %w", a.id.External, err)
	}

	for _, project := range metricsScope.MonitoredProjects {
		// TODO: parse project.Name as a MonitoringMonitoredProjectRef?
		if project.Name == a.id.External {
			a.actual = project
			return true, nil
		}
	}
	return false, nil
}

func (a *monitoringMonitoredProjectAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating MonitoredProject", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MonitoringMonitoredProjectSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	metricsScope, err := a.id.Parent()
	if err != nil {
		return err
	}

	resource.Name = a.id.External

	req := &monitoringpb.CreateMonitoredProjectRequest{
		Parent:           metricsScope.String(),
		MonitoredProject: resource,
	}
	op, err := a.metricsScopeClient.CreateMonitoredProject(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MonitoredProject %s: %w", a.id.External, err)
	}

	// TODO: Update status here
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of MonitoredProject %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created MonitoredProject", "name", a.id.External)

	status := MonitoringMonitoredProjectStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// TODO: Populate ExternalRef
	// status.ExternalRef = &a.id.External
	return setMonitoringMonitoredProjectStatus(u, status)
}

func (a *monitoringMonitoredProjectAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("update not supported on MonitoredProject", "name", a.id.External)

	// TODO: update status (e.g. status.externalRef)?

	return nil
}

func (a *monitoringMonitoredProjectAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringMonitoredProject{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringMonitoredProjectSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// TODO(user): Update other resource reference
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}

	monitoredProjectID, err := a.id.MonitoredProjectID()
	if err != nil {
		return nil, err
	}

	// ref, err := krm.MonitoringMonitoredProjectRefFromExternal(obj.Name)
	// if err != nil {
	// 	return nil, err
	// }

	obj.Spec.MetricsScope = direct.PtrTo(parent.String())
	obj.Spec.ResourceID = direct.PtrTo(monitoredProjectID)
	obj.SetName(monitoredProjectID)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *monitoringMonitoredProjectAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MonitoredProject", "name", a.id.External)

	req := &monitoringpb.DeleteMonitoredProjectRequest{Name: a.id.External}
	op, err := a.metricsScopeClient.DeleteMonitoredProject(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting MonitoredProject %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted MonitoredProject", "name", a.id.External)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MonitoredProject %s: %w", a.id.External, err)
	}
	return true, nil
}

func setMonitoringMonitoredProjectStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}
