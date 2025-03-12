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

// +tool:krm-identity
// proto.service: google.cloud.discoveryengine.v1.DataStoreService
// proto.message: google.cloud.discoveryengine.v1.DataStore
// crd.type: DiscoveryEngineDataStore
// crd.version: v1alpha1

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DataStoreIdentity defines the full identity for a DiscoveryEngineDataStore
//
// +k8s:deepcopy-gen=false
type DataStoreIdentity struct {
	*CollectionIdentity
	DataStore string
}

var _ identity.Identity = &DataStoreIdentity{}

func (i *DataStoreIdentity) FromExternal(external string) error {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" {
		collection := &CollectionIdentity{
			ProjectID:  tokens[1],
			Location:   tokens[3],
			Collection: tokens[5],
		}

		*i = DataStoreIdentity{
			CollectionIdentity: collection,
			DataStore:          tokens[7],
		}
		return nil
	}
	return fmt.Errorf("format of DiscoveryEngineDataStore external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}})", external)
}

func (i *DataStoreIdentity) String() string {
	return i.CollectionIdentity.String() + "/dataStores/" + i.DataStore
}

var _ identity.Resource = &DiscoveryEngineDataStore{}

func (obj *DiscoveryEngineDataStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" && obj.Status.ExternalRef != nil {
		savedID := &DataStoreIdentity{}
		if err := savedID.FromExternal(common.ValueOf(obj.Status.ExternalRef)); err != nil {
			return nil, err
		}
		resourceID = savedID.DataStore
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve DiscoveryEngineDataStore resource ID (obj.Status in %+v)", obj.Status)
	}

	id := &DataStoreIdentity{
		CollectionIdentity: parentID.(*CollectionIdentity),
		DataStore:          resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &DataStoreIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update DiscoveryEngineDataStore identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *DiscoveryEngineDataStore) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	id := obj.Spec.ProjectRef.External + "/locations/" + obj.Spec.Location + "/collections/" + obj.Spec.Collection
	parentID := &CollectionIdentity{}
	if err := parentID.FromExternal(id); err != nil {
		return nil, err
	}
	return parentID, nil
}

// func ParseDataStoreExternal(external string) (parent *DataStoreParent, resourceID string, err error) {
// 	tokens := strings.Split(external, "/")
// 	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "datastores" {
// 		return nil, "", fmt.Errorf("format of DiscoveryEngineDataStore external=%q was not known (use projects/{{projectID}}/locations/{{location}}/datastores/{{datastoreID}})", external)
// 	}
// 	parent = &DataStoreParent{
// 		ProjectID: tokens[1],
// 		Location:  tokens[3],
// 	}
// 	resourceID = tokens[5]
// 	return parent, resourceID, nil
// }

// func (i *DataStoreIdentity) ID() string {
// 	return i.id
// }

// func (i *DataStoreIdentity) Parent() *DataStoreParent {
// 	return i.parent
// }

// type DataStoreParent struct {
// 	ProjectID string
// 	Location  string
// }

// func (p *DataStoreParent) String() string {
// 	return "projects/" + p.ProjectID + "/locations/" + p.Location
// }

// // New builds a DataStoreIdentity from the Config Connector DataStore object.
// func NewDataStoreIdentity(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStore) (*DataStoreIdentity, error) {

// 	// Get Parent
// 	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
// 	if err != nil {
// 		return nil, err
// 	}
// 	projectID := projectRef.ProjectID
// 	if projectID == "" {
// 		return nil, fmt.Errorf("cannot resolve project")
// 	}
// 	location := obj.Spec.Location

// 	// Get desired ID
// 	resourceID := common.ValueOf(obj.Spec.ResourceID)
// 	if resourceID == "" {
// 		resourceID = obj.GetName()
// 	}
// 	if resourceID == "" {
// 		return nil, fmt.Errorf("cannot resolve resource ID")
// 	}

// 	// Use approved External
// 	externalRef := common.ValueOf(obj.Status.ExternalRef)
// 	if externalRef != "" {
// 		// Validate desired with actual
// 		actualParent, actualResourceID, err := ParseDataStoreExternal(externalRef)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if actualParent.ProjectID != projectID {
// 			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
// 		}
// 		if actualParent.Location != location {
// 			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
// 		}
// 		if actualResourceID != resourceID {
// 			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
// 				resourceID, actualResourceID)
// 		}
// 	}
// 	return &DataStoreIdentity{
// 		parent: &DataStoreParent{
// 			ProjectID: projectID,
// 			Location:  location,
// 		},
// 		id: resourceID,
// 	}, nil
// }
