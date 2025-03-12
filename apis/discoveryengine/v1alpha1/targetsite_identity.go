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

// +tool:krm-identity
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.TargetSite
// crd.type: DiscoveryEngineDataStoreTargetSite
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

var _ identity.Identity = &TargetSiteIdentity{}

// TargetSiteIdentity defines the full identity for a DataStoreTargetSite
//
// +k8s:deepcopy-gen=false
type TargetSiteIdentity struct {
	*DataStoreIdentity
	TargetSite string
}

func (i *TargetSiteIdentity) String() string {
	return i.DataStoreIdentity.String() + "/siteSearchEngine/targetSites/" + i.TargetSite
}

func (i *TargetSiteIdentity) FromExternal(external string) error {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 11 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" && tokens[9] == "targetSites" {
		collection := &CollectionIdentity{
			ProjectID:  tokens[1],
			Location:   tokens[3],
			Collection: tokens[5],
		}

		*i = TargetSiteIdentity{
			DataStoreIdentity: &DataStoreIdentity{
				CollectionIdentity: collection,
				DataStore:          tokens[7],
			},
			TargetSite: tokens[10],
		}
		return nil
	}
	return fmt.Errorf("format of DiscoveryEngineDataStoreTargetSite external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}}/siteSearchEngine/targetSites/{{targetSiteID}})", external)

}

var _ identity.Resource = &DiscoveryEngineDataStoreTargetSite{}

func (obj *DiscoveryEngineDataStoreTargetSite) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// resourceID is server-generated
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" && obj.Status.ExternalRef != nil {
		savedID := &TargetSiteIdentity{}
		if err := savedID.FromExternal(common.ValueOf(obj.Status.ExternalRef)); err != nil {
			return nil, err
		}
		resourceID = savedID.TargetSite
	}
	if resourceID == "" {
		return nil, nil
	}

	id := &TargetSiteIdentity{
		DataStoreIdentity: parentID.(*DataStoreIdentity),
		TargetSite:        resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &TargetSiteIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update DiscoveryEngineDataStoreTargetSite identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *DiscoveryEngineDataStoreTargetSite) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.DataStoreRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &DataStoreIdentity{}
	if err := parentID.FromExternal(obj.Spec.DataStoreRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}

// // NewTargetSiteIdentityFromObject builds a TargetSiteIdentity from the Config Connector object.
// func NewTargetSiteIdentityFromObject(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStoreTargetSite) (*DiscoveryEngineDataStoreID, *TargetSiteIdentity, error) {
// 	if obj.Spec.DataStoreRef == nil {
// 		return nil, nil, fmt.Errorf("spec.dataStoreRef not set")
// 	}
// 	dataStoreRef := *obj.Spec.DataStoreRef
// 	if _, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
// 		return nil, nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
// 	}
// 	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(dataStoreRef.External)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("parsing dataStoreRef.external=%q: %w", dataStoreRef.External, err)
// 	}

// 	var link *TargetSiteIdentity

// 	// Validate the status.externalRef, if set
// 	externalRef := valueOf(obj.Status.ExternalRef)
// 	if externalRef != "" {
// 		// Validate desired with actual
// 		externalLink, err := ParseTargetSiteExternal(externalRef)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		if externalLink.DiscoveryEngineDataStoreID.String() != dataStoreLink.String() {
// 			return nil, nil, fmt.Errorf("cannot change object key after creation; status=%q, new=%q",
// 				externalLink.DiscoveryEngineDataStoreID.String(), dataStoreLink.String())
// 		}
// 		link = externalLink
// 	}
// 	return dataStoreLink, link, nil
// }
