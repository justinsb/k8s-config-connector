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
// proto.service: google.cloud.discoveryengine.v1.EngineService
// proto.message: google.cloud.discoveryengine.v1.Engine
// crd.type: DiscoveryEngineEngine
// crd.version: v1alpha1

package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

// EngineIdentity defines the full identity for a DiscoveryEngineEngine
//
// +k8s:deepcopy-gen=false
type EngineIdentity struct {
	*CollectionIdentity
	Engine string
}

var _ identity.Identity = &EngineIdentity{}

func (i *EngineIdentity) FromExternal(external string) error {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "engines" {
		collection := &CollectionIdentity{
			ProjectID:  tokens[1],
			Location:   tokens[3],
			Collection: tokens[5],
		}

		*i = EngineIdentity{
			CollectionIdentity: collection,
			Engine:             tokens[7],
		}
		return nil
	}
	return fmt.Errorf("format of DiscoveryEngineEngine external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/engines/{{dataStoreID}}", external)
}

func (i *EngineIdentity) String() string {
	return i.CollectionIdentity.String() + "/engines/" + i.Engine
}

// // EngineIdentity defines the resource reference to DiscoveryEngineEngine, which "External" field
// // holds the GCP identifier for the KRM object.
// type EngineIdentity struct {
// 	parent *EngineParent
// 	id     string
// }

// func (i *EngineIdentity) String() string {
// 	return i.parent.String() + "/engines/" + i.id
// }

// func (i *EngineIdentity) ID() string {
// 	return i.id
// }

// func (i *EngineIdentity) Parent() *EngineParent {
// 	return i.parent
// }

// type EngineParent struct {
// 	ProjectID string
// 	Location  string
// }

// func (p *EngineParent) String() string {
// 	return "projects/" + p.ProjectID + "/locations/" + p.Location
// }

// // New builds a EngineIdentity from the Config Connector Engine object.
// func NewEngineIdentity(ctx context.Context, reader client.Reader, obj *DiscoveryEngineEngine) (*EngineIdentity, error) {

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
// 		actualParent, actualResourceID, err := ParseEngineExternal(externalRef)
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
// 	return &EngineIdentity{
// 		parent: &EngineParent{
// 			ProjectID: projectID,
// 			Location:  location,
// 		},
// 		id: resourceID,
// 	}, nil
// }

// func ParseEngineExternal(external string) (parent *EngineParent, resourceID string, err error) {
// 	tokens := strings.Split(external, "/")
// 	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "engines" {
// 		return nil, "", fmt.Errorf("format of DiscoveryEngineEngine external=%q was not known (use projects/{{projectID}}/locations/{{location}}/engines/{{engineID}})", external)
// 	}
// 	parent = &EngineParent{
// 		ProjectID: tokens[1],
// 		Location:  tokens[3],
// 	}
// 	resourceID = tokens[5]
// 	return parent, resourceID, nil
// }
