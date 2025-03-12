// Copyright 2025 Google LLC
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

// +tool:controller
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.TargetSite
// crd.type: DiscoveryEngineDataStoreTargetSite
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/googleapis/gax-go/v2/apierror"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineDataStoreTargetSiteGVK, NewDataStoreTargetSiteModel)
}

func NewDataStoreTargetSiteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataStoreTargetSiteModel{config: *config}, nil
}

var _ directbase.Model = &dataStoreTargetSiteModel{}

type dataStoreTargetSiteModel struct {
	config config.ControllerConfig
}

func (m *dataStoreTargetSiteModel) client(ctx context.Context, projectID string) (*gcp.SiteSearchEngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSiteSearchEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine datastoretargetsite client: %w", err)
	}

	return gcpClient, err
}

func (m *dataStoreTargetSiteModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	a := &dataStoreTargetSiteAdapter{}

	var projectID string
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	// This resource has a server-generated id, so we might not have an identity
	if i != nil {
		a.id = i.(*krm.TargetSiteIdentity)
		projectID = a.id.ProjectID
	}

	if a.id == nil {
		i, err := obj.GetParentIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		a.parentID = i.(*krm.DataStoreIdentity)
		projectID = a.parentID.ProjectID
	}

	client, err := m.client(ctx, projectID)
	if err != nil {
		return nil, err
	}
	a.client = client

	mapCtx := &direct.MapContext{}
	a.desired = DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return a, nil
}

func (m *dataStoreTargetSiteModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		var id krm.TargetSiteIdentity
		if err := id.FromExternal(url); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineDataStoreTargetSite format", "url", url)
		} else {
			client, err := m.client(ctx, id.ProjectID)
			if err != nil {
				return nil, err
			}

			return &dataStoreTargetSiteAdapter{
				client: client,
				id:     &id,
			}, nil
		}
	}
	return nil, nil
}

type dataStoreTargetSiteAdapter struct {
	client   *gcp.SiteSearchEngineClient
	id       *krm.TargetSiteIdentity
	parentID *krm.DataStoreIdentity
	desired  *pb.TargetSite
	actual   *pb.TargetSite
}

var _ directbase.Adapter = &dataStoreTargetSiteAdapter{}

func (a *dataStoreTargetSiteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine datastoretargetsite", "name", a.id)

	// Server-generated id; we can't find it without an id
	if a.id == nil {
		return false, nil
	}

	req := &pb.GetTargetSiteRequest{Name: a.id.String()}
	actual, err := a.client.GetTargetSite(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}

		// Not found seems to be a 500 error
		if apiError, ok := err.(*apierror.APIError); ok {
			klog.Infof("apiError error is %T %+v %v", apiError, apiError, apiError)
			klog.Infof("apiError Details is %v", apiError.Details())
			klog.Infof("apiError HTTPCode is %v", apiError.HTTPCode())
			if apiError.HTTPCode() == 500 {
				return false, nil
			}
		}
		klog.Infof("full error is %T %+v %v", err, err, err)
		return false, fmt.Errorf("getting discoveryengine datastoretargetsite %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataStoreTargetSiteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine datastoretargetsite", "parentID", a.parentID)

	desired := direct.ProtoClone(a.desired)
	// desired.Name = a.id.String()

	if a.parentID == nil {
		return fmt.Errorf("parent not ready")
	}

	parent := a.parentID.String() + "/siteSearchEngine"
	req := &pb.CreateTargetSiteRequest{
		Parent:     parent,
		TargetSite: desired,
	}
	op, err := a.client.CreateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine datastoretargetsite %s: %w", req.Parent, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine datastoretargetsite %s waiting creation: %w", req.Parent, err)
	}
	log.V(2).Info("successfully created discoveryengine datastoretargetsite in gcp", "name", created.Name)

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// A bit of a hack here, to avoid changing projectId <-> projectNumber
	targetSiteID := lastComponent(created.Name)
	id := &krm.TargetSiteIdentity{
		DataStoreIdentity: a.parentID,
		TargetSite:        targetSiteID,
	}
	status.ExternalRef = direct.PtrTo(id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreTargetSiteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine datastoretargetsite", "name", a.id)

	fqn := a.id.String()
	fqn = strings.ReplaceAll(fqn, "/justinsb-root-20220725/", "/719301307883/")

	desired := direct.ProtoClone(a.desired)
	desired.Name = fqn

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Type, a.actual.Type) {
		updateMask.Paths = append(updateMask.Paths, "type")
	}
	if !reflect.DeepEqual(a.desired.ExactMatch, a.actual.ExactMatch) {
		updateMask.Paths = append(updateMask.Paths, "exact_match")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	req := &pb.UpdateTargetSiteRequest{
		// UpdateMask: updateMask,
		TargetSite: desired,
	}
	op, err := a.client.UpdateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine datastoretargetsite %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine datastoretargetsite %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated discoveryengine datastoretargetsite", "name", a.id)

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// WARNING - the ID changes when we update the resource (!!!)

	// A bit of a hack here, to avoid changing projectId <-> projectNumber
	targetSiteID := lastComponent(updated.Name)
	id := &krm.TargetSiteIdentity{
		DataStoreIdentity: a.id.DataStoreIdentity,
		TargetSite:        targetSiteID,
	}
	status.ExternalRef = direct.PtrTo(id.String())

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreTargetSiteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	// obj.Spec.Location = a.id.Location
	obj.Spec.DataStoreRef = &krm.DataStoreRef{External: a.id.DataStoreIdentity.String()}
	// obj.Spec.Collection = a.id.Collection
	obj.Spec.ResourceID = direct.PtrTo(a.id.TargetSite)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.TargetSite)
	u.SetGroupVersionKind(krm.DiscoveryEngineDataStoreTargetSiteGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *dataStoreTargetSiteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine datastoretargetsite", "name", a.id)

	// Server-generated id; nothing to delete if not acquired
	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	fqn = strings.ReplaceAll(fqn, "/justinsb-root-20220725/", "/719301307883/")

	req := &pb.DeleteTargetSiteRequest{Name: fqn}
	op, err := a.client.DeleteTargetSite(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting discoveryengine datastoretargetsite %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted discoveryengine datastoretargetsite", "name", fqn)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine datastoretargetsite %s: %w", fqn, err)
		}
	}
	return true, nil
}
