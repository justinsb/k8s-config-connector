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

// AutoscalingPolicyIdentity defines the resource reference to DataprocAutoscalingPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type AutoscalingPolicyIdentity struct {
	ProjectID         string
	Location          string
	AutoscalingPolicy string
}

func ParseAutoscalingPolicyExternal(external string) (*AutoscalingPolicyIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "autoscalingpolicies" {
		return nil, fmt.Errorf("format of DataprocAutoscalingPolicy external=%q was not known (use projects/{{projectID}}/locations/{{location}}/autoscalingpolicies/{{autoscalingpolicyID}})", external)
	}
	id := &AutoscalingPolicyIdentity{
		ProjectID:         tokens[1],
		Location:          tokens[3],
		AutoscalingPolicy: tokens[5],
	}
	return id, nil
}
