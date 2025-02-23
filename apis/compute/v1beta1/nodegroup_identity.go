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

// NodeGroupIdentity defines the resource reference to ComputeNodeGroup, which "External" field
// holds the GCP identifier for the KRM object.
type NodeGroupIdentity struct {
	ProjectID string
	Location  string
	NodeGroup string
}

// TODO: Context
// full URL, partial URI, or node group name are valid. Examples:
// +
// +                               * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
// +                               * `projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
// +                               * `node-group-1`

func ParseNodeGroupIdentity(external string) (*NodeGroupIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "zones" || tokens[4] != "nodeGroups" {
		return nil, fmt.Errorf("format of ComputeNodeGroup external=%q was not known (use projects/{{projectID}}/zones/{{zone}}/nodeGroups/{{nodeGroup}})", external)
	}
	id := &NodeGroupIdentity{
		ProjectID: tokens[1],
		Location:  tokens[3],
		NodeGroup: tokens[5],
	}
	return id, nil
}
