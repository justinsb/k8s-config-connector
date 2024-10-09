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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &MonitoringMonitoredProjectRef{}

// MonitoringMonitoredProjectRef defines the resource reference to MonitoringMonitoredProject, which "External" field
// holds the GCP identifier for the KRM object.
type MonitoringMonitoredProjectRef struct {
	// A reference to an externally managed MonitoringMonitoredProject resource.
	// Should be in the format "projects/<projectID>/locations/<location>/monitoredprojects/<monitoredprojectID>".
	External string `json:"external,omitempty"`

	// The name of a MonitoringMonitoredProject resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MonitoringMonitoredProject resource.
	Namespace string `json:"namespace,omitempty"`

	// parent *MonitoringMonitoredProjectParent
}

// NormalizedExternal provision the "External" value for other resource that depends on MonitoringMonitoredProject.
// If the "External" is given in the other resource's spec.MonitoringMonitoredProjectRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual MonitoringMonitoredProject object from the cluster.
func (r *MonitoringMonitoredProjectRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", MonitoringMonitoredProjectGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := r.parseExternalURL(); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(MonitoringMonitoredProjectGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", MonitoringMonitoredProjectGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("MonitoringMonitoredProject is not ready yet")
	}
	r.External = actualExternalRef
	return r.External, nil
}

func MonitoringMonitoredProjectRefFromExternal(external string) (*MonitoringMonitoredProjectRef, error) {
	r := &MonitoringMonitoredProjectRef{External: external}
	if _, err := r.parseExternalURL(); err != nil {
		return nil, err
	}
	return r, nil
}

/* NOTYET
TODO: MonitoringMonitoredProject currently uses project-id annotation?
// New builds a MonitoringMonitoredProjectRef from the Config Connector MonitoringMonitoredProject object.
func NewMonitoringMonitoredProjectRef(ctx context.Context, reader client.Reader, obj *MonitoringMonitoredProject) (*MonitoringMonitoredProjectRef, error) {
	var id *MonitoringMonitoredProjectRef

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location
	id.parent = &MonitoringMonitoredProjectParent{ProjectID: projectID, Location: location}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asMonitoringMonitoredProjectExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseMonitoringMonitoredProjectExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &MonitoringMonitoredProjectParent{ProjectID: projectID, Location: location}
	return id, nil
}
*/

func (r *MonitoringMonitoredProjectRef) Parent() (*MetricsScopeID, error) {
	// if r.parent != nil {
	// 	return r.parent, nil
	// }
	if r.External != "" {
		gcpURL, err := r.parseExternalURL()
		if err != nil {
			return nil, err
		}

		scopeID := &MetricsScopeID{}
		scopeID.ScopingProjectID, err = gcpURL.Get("metricsScopes")
		if err != nil {
			return nil, err
		}
		return scopeID, nil
	}
	return nil, fmt.Errorf("MonitoringMonitoredProjectRef not initialized from `NewMonitoringMonitoredProjectRef` or `NormalizedExternal`")
}

func (r *MonitoringMonitoredProjectRef) MonitoredProjectID() (string, error) {
	if r.External == "" {
		return "", fmt.Errorf("MonitoringMonitoredProjectRef not initialized")
	}

	gcpURL, err := r.parseExternalURL()
	if err != nil {
		return "", err
	}

	projectID, err := gcpURL.Get("projects")
	if err != nil {
		return "", err
	}
	return projectID, nil
}

type MetricsScopeID struct {
	ScopingProjectID string
}

func (p *MetricsScopeID) String() string {
	return "locations/global/metricsScopes/" + p.ScopingProjectID
}

// func asMonitoringMonitoredProjectExternal(parent *MonitoringMonitoredProjectParent, resourceID string) (external string) {
// 	return parent.String() + "/monitoredprojects/" + resourceID
// }

func (r *MonitoringMonitoredProjectRef) parseExternalURL() (*gcpURL, error) {
	external := r.External

	gcpURL := MatchGCPURL(external, "locations/global/metricsScopes/{SCOPING_PROJECT_ID}/projects/{MONITORED_PROJECT_ID}")
	if gcpURL != nil {
		return gcpURL, nil
	}

	return nil, fmt.Errorf("format of MonitoringMonitoredProject external=%q was not known (use locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER})", external)
}

type gcpURL struct {
	m           map[string]string
	originalURL string
}

func (u *gcpURL) Get(k string) (string, error) {
	v, found := u.m[k]
	if !found {
		return "", fmt.Errorf("key %q not found in url %q", k, u.originalURL)
	}
	return v, nil
}

func MatchGCPURL(s string, template string) *gcpURL {
	m := make(map[string]string)

	tokens := strings.Split(s, "/")
	templateTokens := strings.Split(template, "/")
	if len(tokens) != len(templateTokens) {
		return nil
	}
	for i, templateToken := range templateTokens {
		if strings.HasPrefix(templateToken, "{") {
			// Wildcard
			if i == 0 {
				klog.Fatalf("invalid GCP URL template %q", template)
			}
			key := templateTokens[i-1]
			v := tokens[i]
			m[key] = v
			continue
		}
		if templateToken != tokens[i] {
			return nil
		}
	}

	return &gcpURL{m: m, originalURL: s}
}
