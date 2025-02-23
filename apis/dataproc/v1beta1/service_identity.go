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

// ServiceIdentity is the identifier for a DataprocService
type ServiceIdentity struct {
	ProjectID string
	Location  string
	Service   string
}

// ParseServiceIdentity parses a string into a ServiceIdentity.
// It should recognize the forms that can be provided in the "external" value of a ServiceRef
func ParseServiceIdentity(external string) (*ServiceIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "services" {
		return nil, fmt.Errorf("format of DataproceService external=%q was not known (use projects/{{projectID}}/locations/{{location}}/services/{{service}})", external)
	}
	id := &ServiceIdentity{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Service:   tokens[5],
	}
	return id, nil
}
