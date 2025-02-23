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

// +kcc:proto=google.cloud.dataproc.v1.AuxiliaryNodeGroup
type AuxiliaryNodeGroup struct {
	// Required. Node group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryNodeGroup.node_group
	NodeGroup *NodeGroup `json:"nodeGroup,omitempty"`

	// Optional. A node group ID. Generated if not specified.
	//
	//  The ID must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), and hyphens (-). Cannot begin or end with underscore
	//  or hyphen. Must consist of from 3 to 33 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryNodeGroup.node_group_id
	NodeGroupID *string `json:"nodeGroupID,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AuxiliaryServicesConfig
type AuxiliaryServicesConfig struct {
	// Optional. The Hive Metastore configuration for this workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryServicesConfig.metastore_config
	MetastoreConfig *MetastoreConfig `json:"metastoreConfig,omitempty"`

	// Optional. The Spark History Server configuration for the workload.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryServicesConfig.spark_history_server_config
	SparkHistoryServerConfig *SparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterMetrics
type ClusterMetrics struct {
	// The HDFS metrics.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterMetrics.hdfs_metrics
	HdfsMetrics map[string]int64 `json:"hdfsMetrics,omitempty"`

	// YARN metrics.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterMetrics.yarn_metrics
	YarnMetrics map[string]int64 `json:"yarnMetrics,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterStatus
type ClusterStatus struct {
}

// +kcc:proto=google.cloud.dataproc.v1.ConfidentialInstanceConfig
type ConfidentialInstanceConfig struct {
	// Optional. Defines whether the instance should have confidential compute
	//  enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ConfidentialInstanceConfig.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig
type DataprocMetricConfig struct {
	// Required. Metrics sources to enable.
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.metrics
	Metrics []DataprocMetricConfig_Metric `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig.Metric
type DataprocMetricConfig_Metric struct {
	// Required. A standard set of metrics is collected unless `metricOverrides`
	//  are specified for the metric source (see [Custom metrics]
	//  (https://cloud.google.com/dataproc/docs/guides/dataproc-metrics#custom_metrics)
	//  for more information).
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_source
	MetricSource *string `json:"metricSource,omitempty"`

	// Optional. Specify one or more [Custom metrics]
	//  (https://cloud.google.com/dataproc/docs/guides/dataproc-metrics#custom_metrics)
	//  to collect for the metric course (for the `SPARK` metric source (any
	//  [Spark metric]
	//  (https://spark.apache.org/docs/latest/monitoring.html#metrics) can be
	//  specified).
	//
	//  Provide metrics in the following format:
	//  <code><var>METRIC_SOURCE</var>:<var>INSTANCE</var>:<var>GROUP</var>:<var>METRIC</var></code>
	//  Use camelcase as appropriate.
	//
	//  Examples:
	//
	//  ```
	//  yarn:ResourceManager:QueueMetrics:AppsCompleted
	//  spark:driver:DAGScheduler:job.allJobs
	//  sparkHistoryServer:JVM:Memory:NonHeapMemoryUsage.committed
	//  hiveserver2:JVM:Memory:NonHeapMemoryUsage.used
	//  ```
	//
	//  Notes:
	//
	//  * Only the specified overridden metrics are collected for the
	//    metric source. For example, if one or more `spark:executive` metrics
	//    are listed as metric overrides, other `SPARK` metrics are not
	//    collected. The collection of the metrics for other enabled custom
	//    metric sources is unaffected. For example, if both `SPARK` andd `YARN`
	//    metric sources are enabled, and overrides are provided for Spark
	//    metrics only, all YARN metrics are collected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_overrides
	MetricOverrides []string `json:"metricOverrides,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig
type GkeNodePoolConfig struct {
	// Optional. The node pool configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.config
	Config *GkeNodePoolConfig_GkeNodeConfig `json:"config,omitempty"`

	// Optional. The list of Compute Engine
	//  [zones](https://cloud.google.com/compute/docs/zones#available) where
	//  node pool nodes associated with a Dataproc on GKE virtual cluster
	//  will be located.
	//
	//  **Note:** All node pools associated with a virtual cluster
	//  must be located in the same region as the virtual cluster, and they must
	//  be located in the same zone within that region.
	//
	//  If a location is not specified during node pool creation, Dataproc on GKE
	//  will choose the zone.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.locations
	Locations []string `json:"locations,omitempty"`

	// Optional. The autoscaler configuration for this node pool. The autoscaler
	//  is enabled only when a valid configuration is present.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.autoscaling
	Autoscaling *GkeNodePoolConfig_GkeNodePoolAutoscalingConfig `json:"autoscaling,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig
type GkeNodePoolConfig_GkeNodeConfig struct {
	// Optional. The name of a Compute Engine [machine
	//  type](https://cloud.google.com/compute/docs/machine-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. The number of local SSD disks to attach to the node, which is
	//  limited by the maximum number of disks allowable per zone (see [Adding
	//  Local SSDs](https://cloud.google.com/compute/docs/disks/local-ssd)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.local_ssd_count
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`

	// Optional. Whether the nodes are created as legacy [preemptible VM
	//  instances] (https://cloud.google.com/compute/docs/instances/preemptible).
	//  Also see
	//  [Spot][google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.spot]
	//  VMs, preemptible VM instances without a maximum lifetime. Legacy and Spot
	//  preemptible nodes cannot be used in a node pool with the `CONTROLLER`
	//  [role]
	//  (/dataproc/docs/reference/rest/v1/projects.regions.clusters#role)
	//  or in the DEFAULT node pool if the CONTROLLER role is not assigned (the
	//  DEFAULT node pool will assume the CONTROLLER role).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Optional. A list of [hardware
	//  accelerators](https://cloud.google.com/compute/docs/gpus) to attach to
	//  each node.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.accelerators
	Accelerators []GkeNodePoolConfig_GkeNodePoolAcceleratorConfig `json:"accelerators,omitempty"`

	// Optional. [Minimum CPU
	//  platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
	//  to be used by this instance. The instance may be scheduled on the
	//  specified or a newer CPU platform. Specify the friendly names of CPU
	//  platforms, such as "Intel Haswell"` or Intel Sandy Bridge".
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`

	// Optional. The [Customer Managed Encryption Key (CMEK)]
	//  (https://cloud.google.com/kubernetes-engine/docs/how-to/using-cmek)
	//  used to encrypt the boot disk attached to each node in the node pool.
	//  Specify the key using the following format:
	//  <code>projects/<var>KEY_PROJECT_ID</var>/locations/<var>LOCATION</var>/keyRings/<var>RING_NAME</var>/cryptoKeys/<var>KEY_NAME</var></code>.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.boot_disk_kms_key
	BootDiskKMSKey *string `json:"bootDiskKMSKey,omitempty"`

	// Optional. Whether the nodes are created as [Spot VM instances]
	//  (https://cloud.google.com/compute/docs/instances/spot).
	//  Spot VMs are the latest update to legacy
	//  [preemptible
	//  VMs][google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.preemptible].
	//  Spot VMs do not have a maximum lifetime. Legacy and Spot preemptible
	//  nodes cannot be used in a node pool with the `CONTROLLER`
	//  [role](/dataproc/docs/reference/rest/v1/projects.regions.clusters#role)
	//  or in the DEFAULT node pool if the CONTROLLER role is not assigned (the
	//  DEFAULT node pool will assume the CONTROLLER role).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig
type GkeNodePoolConfig_GkeNodePoolAcceleratorConfig struct {
	// The number of accelerator cards exposed to an instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.accelerator_count
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	// The accelerator type resource namename (see GPUs on Compute Engine).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Size of partitions to create on the GPU. Valid values are described in
	//  the NVIDIA [mig user
	//  guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.gpu_partition_size
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig
type GkeNodePoolConfig_GkeNodePoolAutoscalingConfig struct {
	// The minimum number of nodes in the node pool. Must be >= 0 and <=
	//  max_node_count.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig.min_node_count
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// The maximum number of nodes in the node pool. Must be >= min_node_count,
	//  and must be > 0.
	//  **Note:** Quota must be sufficient to scale up the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolTarget
type GkeNodePoolTarget struct {
	// Required. The target GKE node pool.
	//  Format:
	//  'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}'
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.node_pool
	NodePool *string `json:"nodePool,omitempty"`

	// Required. The roles associated with the GKE node pool.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.roles
	Roles []string `json:"roles,omitempty"`

	// Input only. The configuration for the GKE node pool.
	//
	//  If specified, Dataproc attempts to create a node pool with the
	//  specified shape. If one with the same name already exists, it is
	//  verified against all specified fields. If a field differs, the
	//  virtual cluster creation will fail.
	//
	//  If omitted, any node pool with the specified name is used. If a
	//  node pool with the specified name does not exist, Dataproc create a
	//  node pool with default values.
	//
	//  This is an input only field. It will not be returned by the API.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.node_pool_config
	NodePoolConfig *GkeNodePoolConfig `json:"nodePoolConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.IdentityConfig
type IdentityConfig struct {
	// Required. Map of user to service account.
	// +kcc:proto:field=google.cloud.dataproc.v1.IdentityConfig.user_service_account_mapping
	UserServiceAccountMapping map[string]string `json:"userServiceAccountMapping,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicy struct {
	// Optional. Defines how the Group selects the provisioning model to ensure
	//  required reliability.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.provisioning_model_mix
	ProvisioningModelMix *InstanceFlexibilityPolicy_ProvisioningModelMix `json:"provisioningModelMix,omitempty"`

	// Optional. List of instance selection options that the group will use when
	//  creating new VMs.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_list
	InstanceSelectionList []InstanceFlexibilityPolicy_InstanceSelection `json:"instanceSelectionList,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection
type InstanceFlexibilityPolicy_InstanceSelection struct {
	// Optional. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.machine_types
	MachineTypes []string `json:"machineTypes,omitempty"`

	// Optional. Preference of this instance selection. Lower number means
	//  higher preference. Dataproc will first try to create a VM based on the
	//  machine-type with priority rank and fallback to next rank based on
	//  availability. Machine types and instance selections with the same
	//  priority have the same preference.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.rank
	Rank *int32 `json:"rank,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResult struct {
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix
type InstanceFlexibilityPolicy_ProvisioningModelMix struct {
	// Optional. The base capacity that will always use Standard VMs to avoid
	//  risk of more preemption than the minimum capacity you need. Dataproc will
	//  create only standard VMs until it reaches standard_capacity_base, then it
	//  will start using standard_capacity_percent_above_base to mix Spot with
	//  Standard VMs. eg. If 15 instances are requested and
	//  standard_capacity_base is 5, Dataproc will create 5 standard VMs and then
	//  start mixing spot and standard VMs for remaining 10 instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_base
	StandardCapacityBase *int32 `json:"standardCapacityBase,omitempty"`

	// Optional. The percentage of target capacity that should use Standard VM.
	//  The remaining percentage will use Spot VMs. The percentage applies only
	//  to the capacity above standard_capacity_base. eg. If 15 instances are
	//  requested and standard_capacity_base is 5 and
	//  standard_capacity_percent_above_base is 30, Dataproc will create 5
	//  standard VMs and then start mixing spot and standard VMs for remaining 10
	//  instances. The mix will be 30% standard and 70% spot.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_percent_above_base
	StandardCapacityPercentAboveBase *int32 `json:"standardCapacityPercentAboveBase,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceReference
type InstanceReference struct {
	// The user-friendly name of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_name
	InstanceName *string `json:"instanceName,omitempty"`

	// The unique identifier of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// The public RSA key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_key
	PublicKey *string `json:"publicKey,omitempty"`

	// The public ECIES key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_ecies_key
	PublicEciesKey *string `json:"publicEciesKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.KubernetesClusterConfig
type KubernetesClusterConfig struct {
	// Optional. A namespace within the Kubernetes cluster to deploy into. If this
	//  namespace does not exist, it is created. If it exists, Dataproc verifies
	//  that another Dataproc VirtualCluster is not installed into it. If not
	//  specified, the name of the Dataproc Cluster is used.
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.kubernetes_namespace
	KubernetesNamespace *string `json:"kubernetesNamespace,omitempty"`

	// Required. The configuration for running the Dataproc cluster on GKE.
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.gke_cluster_config
	GkeClusterConfig *GkeClusterConfig `json:"gkeClusterConfig,omitempty"`

	// Optional. The software configuration for this Dataproc cluster running on
	//  Kubernetes.
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.kubernetes_software_config
	KubernetesSoftwareConfig *KubernetesSoftwareConfig `json:"kubernetesSoftwareConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.KubernetesSoftwareConfig
type KubernetesSoftwareConfig struct {
	// The components that should be installed in this Dataproc cluster. The key
	//  must be a string from the KubernetesComponent enumeration. The value is
	//  the version of the software to be installed.
	//  At least one entry must be specified.
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesSoftwareConfig.component_version
	ComponentVersion map[string]string `json:"componentVersion,omitempty"`

	// The properties to set on daemon config files.
	//
	//  Property keys are specified in `prefix:property` format, for example
	//  `spark:spark.kubernetes.container.image`. The following are supported
	//  prefixes and their mappings:
	//
	//  * spark:  `spark-defaults.conf`
	//
	//  For more information, see [Cluster
	//  properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesSoftwareConfig.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfig struct {
	// Optional. The duration to keep the cluster alive while idling (when no jobs
	//  are running). Passing this threshold will cause the cluster to be
	//  deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON
	//  representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_delete_ttl
	IdleDeleteTtl *string `json:"idleDeleteTtl,omitempty"`

	// Optional. The time when cluster will be auto-deleted (see JSON
	//  representation of
	//  [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_time
	AutoDeleteTime *string `json:"autoDeleteTime,omitempty"`

	// Optional. The lifetime duration of cluster. The cluster will be
	//  auto-deleted at the end of this period. Minimum value is 10 minutes;
	//  maximum value is 14 days (see JSON representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_ttl
	AutoDeleteTtl *string `json:"autoDeleteTtl,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfig struct {
}

// +kcc:proto=google.cloud.dataproc.v1.NodeGroup
type NodeGroup struct {
	// The Node group [resource name](https://aip.dev/122).
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.name
	Name *string `json:"name,omitempty"`

	// Required. Node group roles.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.roles
	Roles []string `json:"roles,omitempty"`

	// Optional. The node group instance group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.node_group_config
	NodeGroupConfig *InstanceGroupConfig `json:"nodeGroupConfig,omitempty"`

	// Optional. Node group labels.
	//
	//  * Label **keys** must consist of from 1 to 63 characters and conform to
	//    [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  * Label **values** can be empty. If specified, they must consist of from
	//    1 to 63 characters and conform to [RFC 1035]
	//    (https://www.ietf.org/rfc/rfc1035.txt).
	//  * The node group must have no more than 32 labels.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeInitializationAction
type NodeInitializationAction struct {
	// Required. Cloud Storage URI of executable file.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.executable_file
	ExecutableFile *string `json:"executableFile,omitempty"`

	// Optional. Amount of time executable has to complete. Default is
	//  10 minutes (see JSON representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	//  Cluster creation fails with an explanatory error message (the
	//  name of the executable that caused the error and the exceeded timeout
	//  period) if the executable is not completed at end of the timeout period.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ReservationAffinity
type ReservationAffinity struct {
	// Optional. Type of reservation to consume
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Optional. Corresponds to the label key of reservation resource.
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of reservation resource.
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SecurityConfig
type SecurityConfig struct {
	// Optional. Kerberos related configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.kerberos_config
	KerberosConfig *KerberosConfig `json:"kerberosConfig,omitempty"`

	// Optional. Identity related configuration, including service account based
	//  secure multi-tenancy user mappings.
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.identity_config
	IdentityConfig *IdentityConfig `json:"identityConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SoftwareConfig
type SoftwareConfig struct {
	// Optional. The version of software inside the cluster. It must be one of the
	//  supported [Dataproc
	//  Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported-dataproc-image-versions),
	//  such as "1.2" (including a subminor version, such as "1.2.29"), or the
	//  ["preview"
	//  version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions).
	//  If unspecified, it defaults to the latest Debian version.
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	// Optional. The properties to set on daemon config files.
	//
	//  Property keys are specified in `prefix:property` format, for example
	//  `core:hadoop.tmp.dir`. The following are supported prefixes
	//  and their mappings:
	//
	//  * capacity-scheduler: `capacity-scheduler.xml`
	//  * core:   `core-site.xml`
	//  * distcp: `distcp-default.xml`
	//  * hdfs:   `hdfs-site.xml`
	//  * hive:   `hive-site.xml`
	//  * mapred: `mapred-site.xml`
	//  * pig:    `pig.properties`
	//  * spark:  `spark-defaults.conf`
	//  * yarn:   `yarn-site.xml`
	//
	//  For more information, see [Cluster
	//  properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The set of components to activate on the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.optional_components
	OptionalComponents []string `json:"optionalComponents,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.StartupConfig
type StartupConfig struct {
	// Optional. The config setting to enable cluster creation/ updation to be
	//  successful only after required_registration_fraction of instances are up
	//  and running. This configuration is applicable to only secondary workers for
	//  now. The cluster will fail if required_registration_fraction of instances
	//  are not available. This will include instance creation, agent registration,
	//  and service registration (if enabled).
	// +kcc:proto:field=google.cloud.dataproc.v1.StartupConfig.required_registration_fraction
	RequiredRegistrationFraction *float64 `json:"requiredRegistrationFraction,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.VirtualClusterConfig
type VirtualClusterConfig struct {
	// Optional. A Cloud Storage bucket used to stage job
	//  dependencies, config files, and job driver console output.
	//  If you do not specify a staging bucket, Cloud
	//  Dataproc will determine a Cloud Storage location (US,
	//  ASIA, or EU) for your cluster's staging bucket according to the
	//  Compute Engine zone where your cluster is deployed, and then create
	//  and manage this project-level, per-location bucket (see
	//  [Dataproc staging and temp
	//  buckets](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
	//  **This field requires a Cloud Storage bucket name, not a `gs://...` URI to
	//  a Cloud Storage bucket.**
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.staging_bucket
	StagingBucket *string `json:"stagingBucket,omitempty"`

	// Required. The configuration for running the Dataproc cluster on
	//  Kubernetes.
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.kubernetes_cluster_config
	KubernetesClusterConfig *KubernetesClusterConfig `json:"kubernetesClusterConfig,omitempty"`

	// Optional. Configuration of auxiliary services used by this cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.auxiliary_services_config
	AuxiliaryServicesConfig *AuxiliaryServicesConfig `json:"auxiliaryServicesConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterStatus
type ClusterStatusObservedState struct {
	// Output only. The cluster's state.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.state
	State *string `json:"state,omitempty"`

	// Optional. Output only. Details of cluster's state.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.detail
	Detail *string `json:"detail,omitempty"`

	// Output only. Time when this state was entered (see JSON representation of
	//  [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.state_start_time
	StateStartTime *string `json:"stateStartTime,omitempty"`

	// Output only. Additional state information that includes
	//  status reported by the agent.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.substate
	Substate *string `json:"substate,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicyObservedState struct {
	// Output only. A list of instance selection results in the group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_results
	InstanceSelectionResults []InstanceFlexibilityPolicy_InstanceSelectionResult `json:"instanceSelectionResults,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResultObservedState struct {
	// Output only. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Output only. Number of VM provisioned with the machine_type.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.vm_count
	VmCount *int32 `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfigObservedState struct {
	// Output only. The time when cluster became idle (most recent job finished)
	//  and became eligible for deletion due to idleness (see JSON representation
	//  of
	//  [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_start_time
	IdleStartTime *string `json:"idleStartTime,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfigObservedState struct {
	// Output only. The name of the Instance Template used for the Managed
	//  Instance Group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_template_name
	InstanceTemplateName *string `json:"instanceTemplateName,omitempty"`

	// Output only. The name of the Instance Group Manager for this group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_name
	InstanceGroupManagerName *string `json:"instanceGroupManagerName,omitempty"`

	// Output only. The partial URI to the instance group manager for this group.
	//  E.g. projects/my-project/regions/us-central1/instanceGroupManagers/my-igm.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_uri
	InstanceGroupManagerURI *string `json:"instanceGroupManagerURI,omitempty"`
}
