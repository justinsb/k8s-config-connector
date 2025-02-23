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

// ImageIdentity is the identifier for a ComputeImage.
type ImageIdentity struct {
	ProjectID string
	Image     string
}

// TODO: Context
//  Image examples:
//
//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/[image-id]`
//  * `projects/[project_id]/global/images/[image-id]`
//  * `image-id`
//
//  Image family examples. Dataproc will use the most recent
//  image from the family:
//
//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/family/[custom-image-family-name]`
//  * `projects/[project_id]/global/images/family/[custom-image-family-name]`
//

// ParseImageIdentity parses a string specified for an image into a ImageIdentity.
// It should recognize the forms that can be provided in the "external" value of a ComputeImageRef
func ParseImageIdentity(external string) (*ImageIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 5 || tokens[0] != "projects" || tokens[2] != "global" || tokens[3] != "images" {
		return nil, fmt.Errorf("format of ComputeImage external=%q was not known (use projects/{{projectID}}/global/images/{{image}})", external)
	}
	id := &ImageIdentity{
		ProjectID: tokens[1],
		Image:     tokens[4],
	}
	return id, nil
}
