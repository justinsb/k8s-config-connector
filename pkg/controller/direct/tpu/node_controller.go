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

package tpu

// +tool:controller
// proto.service: google.cloud.tpu.v1.Tpu
// proto.message: google.cloud.tpu.v1.Node
// crd.type: TPUNode
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "cloud.google.com/go/tpu/apiv1"
	pb "cloud.google.com/go/tpu/apiv1/tpupb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TPUNodeGVK, NewTPUNodeModel)
}

func NewTPUNodeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTPUNode{config: config}, nil
}

var _ directbase.Model = &modelTPUNode{}

type modelTPUNode struct {
	config *config.ControllerConfig
}

func (m *modelTPUNode) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.TPUNode{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.TPUNodeIdentity)
	project := &refs.Project{ProjectID: id.ParentID.ProjectID}

	if err := common.NormalizeReferences(ctx, kube, obj, project); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	tpuClient, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := TPUNodeSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &TPUNodeAdapter{
		id: id,
		// k8sClient: reader,
		tpuClient: tpuClient,
		desired:   desiredProto,
	}, nil
}

// AdapterForURL returns an adapter for export of an object by URL
// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
func (m *modelTPUNode) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //tpu.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/nodes/NODE_ID
	if external, ok := strings.CutPrefix(url, "//tpu.googleapis.com/"); ok {
		id := &krm.TPUNodeIdentity{}
		if err := id.FromExternal(external); err == nil {

			gcpClient, err := newGCPClient(ctx, m.config)
			if err != nil {
				return nil, err
			}
			tpuClient, err := gcpClient.newClient(ctx)
			if err != nil {
				return nil, err
			}

			return &TPUNodeAdapter{
				id:        id,
				tpuClient: tpuClient,
			}, nil
		}
	}

	return nil, nil
}

type TPUNodeAdapter struct {
	id *krm.TPUNodeIdentity
	// k8sClient client.Reader
	tpuClient *api.Client
	desired   *pb.Node
	actual    *pb.Node
}

var _ directbase.Adapter = &TPUNodeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TPUNodeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TPUNode", "name", a.id)

	obj, err := a.tpuClient.GetNode(ctx, &pb.GetNodeRequest{
		Name: a.id.String(),
	})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TPUNode %q: %w", a.id, err)
	}

	a.actual = obj
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TPUNodeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TPUNode", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := direct.ProtoClone(a.desired)

	req := &pb.CreateNodeRequest{
		Parent: a.id.ParentID.String(),
		Node:   desired,
		NodeId: a.id.Node,
	}
	op, err := a.tpuClient.CreateNode(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tpu node %s: %w", a.id, err)
	}
	if !op.Done() {
		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for creating of tpu node %q: %w", a.id, err)
		}
	}

	created, err := a.tpuClient.GetNode(ctx, &pb.GetNodeRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created tpu node %q: %w", a.id, err)
	}

	log.V(2).Info("successfully created tpu node", "name", a.id)

	status := &krm.TPUNodeStatus{}
	status.ObservedState = TPUNodeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TPUNodeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating tpu node", "name", a.id)
	mapCtx := &direct.MapContext{}
	updateMask := fieldmaskpb.FieldMask{}

	if a.desired.Description != a.actual.Description {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.TPUNodeStatus{}
		status.ObservedState = TPUNodeObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	// Update does not appear to be in v1 of the API!!!
	return fmt.Errorf("tpu node update is not supported")
	// op, err := a.tpuClient.Patch(a.id.String(), resource).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
	// if err != nil {
	// 	return fmt.Errorf("updating TPUNode %s: %w", a.id, err)
	// }
	// if !op.Done() {
	// 	if _, err := op.Wait(ctx); err != nil {
	// 		return fmt.Errorf("waiting for tpu node update on %q: %w", a.id, err)
	// 	}
	// }
	// updated, err := a.tpuClient.GetNode(ctx, &pb.GetNodeRequest{Name: a.id.String()})
	// if err != nil {
	// 	return fmt.Errorf("getting updated tpu node %q: %w", a.id, err)
	// }
	// log.V(2).Info("successfully updated tpu node", "name", a.id)

	// status := &krm.TPUNodeStatus{}
	// status.ObservedState = TPUNodeObservedState_FromProto(mapCtx, updated)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	// return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TPUNodeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TPUNode{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TPUNodeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ParentID.ProjectID}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Node)
	u.SetGroupVersionKind(krm.TPUNodeGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TPUNodeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting tpu node", "name", a.id)

	op, err := a.tpuClient.DeleteNode(ctx, &pb.DeleteNodeRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent tpu node, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting tpu node %q: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted tpu node", "name", a.id)

	if !op.Done() {
		if _, err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for deletion of tpu node %q: %w", a.id, err)
		}
	}
	return true, nil
}
