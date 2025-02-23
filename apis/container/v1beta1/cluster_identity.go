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

// ClusterIdentity is the identifier for a ContainerCluster.
type ClusterIdentity struct {
	ProjectID string
	Location  string
	Cluster   string
}

// ParseClusterIdentity parses a string specified for an cluster into a ClusterIdentity.
// It should recognize the forms that can be provided in the "external" value of a ComputeClusterRef
func ParseClusterIdentity(external string) (*ClusterIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		id := &ClusterIdentity{
			ProjectID: tokens[1],
			Location:  tokens[3],
			Cluster:   tokens[5],
		}
		return id, nil
	}
	return nil, fmt.Errorf("format of ContainerCluster external=%q was not known (use projects/{{projectID}}/locations/{{location}}/clusters/{{cluster}})", external)
}
