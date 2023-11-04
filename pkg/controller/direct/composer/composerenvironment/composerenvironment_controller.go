// Copyright 2022 Google LLC
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

package composerenvironment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/composer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"golang.org/x/oauth2"
	composer "google.golang.org/api/composer/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Add creates a new controller and adds it to the Manager.
func Add(mgr manager.Manager, tfProvider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) error {
	gvk := schema.GroupVersionKind{
		Group:   "composer.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "ComposerEnvironment",
	}
	return directbase.Add(mgr, tfProvider, smLoader, converter, dclConfig, gvk, &model{})
}

type model struct {
}

type adapter struct {
	projectID  string
	location   string
	resourceID string
	namespace  string

	desired *v1alpha1.ComposerEnvironment
	actual  *v1alpha1.ComposerEnvironment

	client *composer.Service
	k8s    client.Client
}

func (*model) AdapterForObject(ctx context.Context, kubeClient client.Client, u *unstructured.Unstructured) (directbase.Adapter, error) {
	var clientOptions []option.ClientOption

	{
		// TODO: Should we abuse the oauth2.HTTPClient like this, vs creating our own?
		v := ctx.Value(oauth2.HTTPClient)
		if v != nil {
			clientOptions = append(clientOptions, option.WithHTTPClient(v.(*http.Client)))
		}
	}

	client, err := composer.NewService(ctx, clientOptions...)
	if err != nil {
		return nil, fmt.Errorf("building composer client: %w", err)
	}

	// TODO: Just fetch this object from kube-apiserver?
	obj := &v1alpha1.ComposerEnvironment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	location := "us-central1" // TODO

	return &adapter{
		projectID:  projectID,
		location:   location,
		resourceID: resourceID,
		desired:    obj,
		client:     client,
		k8s:        kubeClient,
		namespace:  u.GetNamespace(),
	}, nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

type kccOperation struct {
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	fqn := "projects/" + a.projectID + "/locations/" + a.location + "/environments/" + a.resourceID
	key, err := a.client.Projects.Locations.Environments.Get(fqn).Context(ctx).Do()
	if IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	u := &v1alpha1.ComposerEnvironment{}
	mapCtx := &mappings.KCCOperation{
		Client:    a.k8s,
		Ctx:       ctx,
		Namespace: a.namespace,
	}
	if err := mapping.MapCloudToKRM(mapCtx, key, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *adapter) Delete(ctx context.Context) error {
	// TODO: Compute fqn from status?
	fqn := "projects/" + a.projectID + "/locations/" + a.location + "/environments/" + a.resourceID
	op, err := a.client.Projects.Locations.Environments.Delete(fqn).Context(ctx).Do()
	if err != nil {
		return err
	}

	for {
		status, err := a.client.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("waiting for operation to finish: %w", err)
		}
		// TODO: Check properly for exit criteria
		if status.Done {
			break
		}
	}

	return nil
}

func (a *adapter) Create(ctx context.Context) (*unstructured.Unstructured, error) {
	desired := &composer.Environment{}
	mapCtx := &mappings.KCCOperation{
		Client:    a.k8s,
		Ctx:       ctx,
		Namespace: a.namespace,
	}
	if err := mapping.MapKRMToCloud(mapCtx, a.desired, desired); err != nil {
		return nil, err
	}
	// TODO: Don't special case?
	fqn := "projects/" + a.projectID + "/locations/" + a.location + "/environments/" + a.resourceID
	desired.Name = fqn

	parent := "projects/" + a.projectID + "/locations/" + a.location
	op, err := a.client.Projects.Locations.Environments.Create(parent, desired).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("creating environment: %w", err)
	}

	for {
		status, err := a.client.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return nil, fmt.Errorf("waiting for operation to finish: %w", err)
		}
		// TODO: Check properly for exit criteria
		if status.Done {
			break
		}
	}

	// TODO: Return created object
	return nil, nil
}

func (a *adapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("Update not implemented")
}

// IsNotFound reports whether err is the result of the
// server replying with http.StatusNotFound.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	ae, ok := err.(*googleapi.Error)
	return ok && ae.Code == http.StatusNotFound
}
