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

package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

// CollectionIdentity defines the full identity for a collection, the parent of a DataStore
//
// +k8s:deepcopy-gen=false
type CollectionIdentity struct {
	ProjectID  string
	Location   string
	Collection string
}

var _ identity.Identity = &CollectionIdentity{}

func (i *CollectionIdentity) FromExternal(external string) error {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" {
		*i = CollectionIdentity{
			ProjectID:  tokens[1],
			Location:   tokens[3],
			Collection: tokens[5],
		}
		return nil
	}
	return fmt.Errorf("format of DiscoveryEngineCollection external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}})", external)

}

func (i *CollectionIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s", i.ProjectID, i.Location, i.Collection)
}
