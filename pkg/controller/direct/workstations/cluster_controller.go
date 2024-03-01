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

package workstations

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/klog/v2"

	workstations "cloud.google.com/go/workstations/apiv1beta"
	pb "cloud.google.com/go/workstations/apiv1beta/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

// AddWorkstationClusterController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddWorkstationClusterController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.WorkstationsWorkstationClusterGVK

	return directbase.Add(mgr, gvk, &workstationClusterModel{config: *config})
}

type workstationClusterModel struct {
	config controller.Config
}

var workstationClusterMapping = NewMapping(&pb.WorkstationCluster{}, &krm.WorkstationsWorkstationCluster{},
	Spec("network"),
	Spec("subnetwork"),
	Spec("displayName"),
	Spec("privateClusterConfig"),

	TODO("createTime"), // TODO: Is this useful?
	TODO("updateTime"), // TODO: Is this useful?
	TODO("deleteTime"), // TODO: Is this useful?

	TODO("degraded"),

	Ignore("etag"),

	TODO("conditions"),         // Map to observedState.conditions?
	TODO("resourceConditions"), // Map to observedState.conditions

	TODO("name"), // TODO: Map to selfLink?
	TODO("uid"),  // TODO: Is this useful?

	Ignore("projectRef"), // Handled in code
	Ignore("resourceID"), // Handled in code
	Ignore("location"),   // Handled in code

	TODO("labels"),
	TODO("annotations"),
).
	MapNested(&pb.WorkstationCluster_PrivateClusterConfig{}, &krm.WorkstationclusterPrivateClusterConfig{},
		"allowedProjects",
		"enablePrivateEndpoint",
		"clusterHostname",
		"serviceAttachmentUri",
	).
	MustBuild()

type workstationClusterAdapter struct {
	projectID            string
	location             string
	workstationclusterID string

	desired *krm.WorkstationsWorkstationCluster
	actual  *krm.WorkstationsWorkstationCluster

	gcp *workstations.Client
}

func (m *workstationClusterModel) client(ctx context.Context) (*workstations.Client, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := workstations.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building workstations client: %w", err)
	}
	return gcpClient, err
}

func (m *workstationClusterModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.WorkstationsWorkstationCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("unable to determine location")
	}

	workstationclusterID := ValueOf(obj.Spec.ResourceID)
	if workstationclusterID == "" {
		workstationclusterID = obj.GetName()
	}
	if workstationclusterID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	return &workstationClusterAdapter{
		projectID:            projectID,
		location:             location,
		workstationclusterID: workstationclusterID,
		desired:              obj,
		gcp:                  gcp,
	}, nil
}

func (a *workstationClusterAdapter) Find(ctx context.Context) (bool, error) {
	if a.workstationclusterID == "" {
		return false, nil
	}

	req := &pb.GetWorkstationClusterRequest{
		Name: a.fullyQualifiedName(),
	}
	gcpObject, err := a.gcp.GetWorkstationCluster(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("workstationCluster was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.WorkstationsWorkstationCluster{}
	if err := workstationClusterMapping.Map(gcpObject, u, nil); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *workstationClusterAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteWorkstationClusterRequest{
		Name: a.fullyQualifiedName(),
	}
	op, err := a.gcp.DeleteWorkstationCluster(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting workstationCluster: %w", err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for workstationCluster deletion: %w", err)
	}

	return true, nil
}

func (a *workstationClusterAdapter) Create(ctx context.Context, obj *unstructured.Unstructured) error {
	desired := &pb.WorkstationCluster{}
	if err := workstationClusterMapping.MapSpec(a.desired, desired); err != nil {
		return err
	}

	req := &pb.CreateWorkstationClusterRequest{
		Parent:               fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location),
		WorkstationClusterId: a.workstationclusterID,
		WorkstationCluster:   desired,
	}

	op, err := a.gcp.CreateWorkstationCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating workstationCluster: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for workstationCluster creation: %w", err)
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("created workstationCluster", "workstationCluster", created)
	// TODO: Populate status in obj
	return nil
}

func (a *workstationClusterAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	desired := &pb.WorkstationCluster{}
	if err := workstationClusterMapping.Map(a.desired, desired, nil); err != nil {
		return nil, err
	}

	desired.Name = a.fullyQualifiedName()

	updateMask := &fieldmaskpb.FieldMask{}
	if a.actual.Spec.DisplayName != a.desired.Spec.DisplayName {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	// TODO: Support more paths
	// TODO: Use a helper function to construct the UpdateMask

	req := &pb.UpdateWorkstationClusterRequest{
		WorkstationCluster: desired,
		UpdateMask:         updateMask,
	}

	updated, err := a.gcp.UpdateWorkstationCluster(ctx, req)
	if err != nil {
		return nil, err
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("updated workstationCluster", "workstationCluster", updated)
	// TODO: Populate status in obj
	return nil, nil
}

func (a *workstationClusterAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", a.projectID, a.location, a.workstationclusterID)
}
