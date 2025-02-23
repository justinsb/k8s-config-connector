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

package v1beta1

import (
	"fmt"
	"strings"
)

// NodePoolIdentity is the identifier for a ContainerNodePool.
type NodePoolIdentity struct {
	ProjectID string
	Location  string
	Cluster   string
	NodePool  string
}

// ParseNodePoolIdentity parses a string specified for an nodepool into a NodePoolIdentity.
// It should recognize the forms that can be provided in the "external" value of a ComputeNodePoolRef
func ParseNodePoolIdentity(external string) (*NodePoolIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "nodePools" {
		id := &NodePoolIdentity{
			ProjectID: tokens[1],
			Location:  tokens[3],
			Cluster:   tokens[5],
			NodePool:  tokens[7],
		}
		return id, nil
	}

	return nil, fmt.Errorf("format of ContainerNodePool external=%q was not known (use projects/{{projectID}}/locations/{{location}}/nodePools/{{nodepool}})", external)

}
