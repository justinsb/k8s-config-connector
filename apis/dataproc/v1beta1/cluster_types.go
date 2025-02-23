// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

var DataprocClusterGVK = GroupVersion.WithKind("DataprocCluster")

// DataprocClusterSpec defines the desired state of DataprocCluster
// +kcc:proto=google.cloud.dataproc.v1.Cluster
type DataprocClusterSpec struct {
	// The DataprocCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The cluster name, which must be unique within a project.
	// The name must start with a lowercase letter, and can contain
	// up to 51 lowercase letters, numbers, and hyphens. It cannot end
	// with a hyphen. The name of a deleted cluster can be reused.
	Name *string `json:"name,omitempty"`

	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The region of this cluster.
	// +required
	Region string `json:"region,omitempty"`

	// Optional. The cluster config for a cluster of Compute Engine Instances.
	//  Note that Dataproc may set default values, and values may change
	//  when clusters are updated.
	//
	//  Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
	Config *ClusterConfig `json:"config,omitempty"`

	/* NOTYET
	// Optional. The virtual cluster config is used when creating a Dataproc
	//  cluster that does not directly control the underlying compute resources,
	//  for example, when creating a [Dataproc-on-GKE
	//  cluster](https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke-overview).
	//  Dataproc may set default values, and values may change when
	//  clusters are updated. Exactly one of
	//  [config][google.cloud.dataproc.v1.Cluster.config] or
	//  [virtual_cluster_config][google.cloud.dataproc.v1.Cluster.virtual_cluster_config]
	//  must be specified.
	VirtualClusterConfig *VirtualClusterConfig `json:"virtualClusterConfig,omitempty"`
	*/

	/* NOTYET
	// Optional. The labels to associate with this cluster.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a cluster.
	Labels map[string]string `json:"labels,omitempty"`
	*/

}

// DataprocClusterStatus defines the config connector machine state of DataprocCluster
type DataprocClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocClusterObservedState `json:"observedState,omitempty"`
}

// DataprocClusterObservedState is the state of the DataprocCluster resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataproc.v1.Cluster
type DataprocClusterObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataproccluster;gcpdataprocclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocCluster is the Schema for the DataprocCluster API
// +k8s:openapi-gen=true
type DataprocCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocClusterSpec   `json:"spec,omitempty"`
	Status DataprocClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocClusterList contains a list of DataprocCluster
type DataprocClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocCluster{}, &DataprocClusterList{})
}
