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


// +kcc:proto=google.cloud.dataproc.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// Full URL, partial URI, or short name of the accelerator type resource to
	//  expose to this instance. See
	//  [Compute Engine
	//  AcceleratorTypes](https://cloud.google.com/compute/docs/reference/v1/acceleratorTypes).
	//
	//  Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `nvidia-tesla-t4`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the accelerator type
	//  resource, for example, `nvidia-tesla-t4`.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_type_uri
	AcceleratorTypeURI *string `json:"acceleratorTypeURI,omitempty"`

	// The number of the accelerator cards of this type exposed to this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AutoscalingConfig
type AutoscalingConfig struct {
	// Optional. The autoscaling policy used by the cluster.
	//
	//  Only resource names including projectid and location (region) are valid.
	//  Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]`
	//  * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]`
	//
	//  Note that the policy must be in the same project and Dataproc region.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingConfig.policy_uri
	PolicyURI *string `json:"policyURI,omitempty"`
}

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

// +kcc:proto=google.cloud.dataproc.v1.Cluster
type Cluster struct {
	// Required. The Google Cloud Platform project ID that the cluster belongs to.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The cluster name, which must be unique within a project.
	//  The name must start with a lowercase letter, and can contain
	//  up to 51 lowercase letters, numbers, and hyphens. It cannot end
	//  with a hyphen. The name of a deleted cluster can be reused.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Optional. The cluster config for a cluster of Compute Engine Instances.
	//  Note that Dataproc may set default values, and values may change
	//  when clusters are updated.
	//
	//  Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.config
	Config *ClusterConfig `json:"config,omitempty"`

	// Optional. The virtual cluster config is used when creating a Dataproc
	//  cluster that does not directly control the underlying compute resources,
	//  for example, when creating a [Dataproc-on-GKE
	//  cluster](https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke-overview).
	//  Dataproc may set default values, and values may change when
	//  clusters are updated. Exactly one of
	//  [config][google.cloud.dataproc.v1.Cluster.config] or
	//  [virtual_cluster_config][google.cloud.dataproc.v1.Cluster.virtual_cluster_config]
	//  must be specified.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.virtual_cluster_config
	VirtualClusterConfig *VirtualClusterConfig `json:"virtualClusterConfig,omitempty"`

	// Optional. The labels to associate with this cluster.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfig struct {
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
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.config_bucket
	ConfigBucket *string `json:"configBucket,omitempty"`

	// Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs
	//  data, such as Spark and MapReduce history files. If you do not specify a
	//  temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or
	//  EU) for your cluster's temp bucket according to the Compute Engine zone
	//  where your cluster is deployed, and then create and manage this
	//  project-level, per-location bucket. The default bucket has a TTL of 90
	//  days, but you can use any TTL (or none) if you specify a bucket (see
	//  [Dataproc staging and temp
	//  buckets](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
	//  **This field requires a Cloud Storage bucket name, not a `gs://...` URI to
	//  a Cloud Storage bucket.**
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.temp_bucket
	TempBucket *string `json:"tempBucket,omitempty"`

	// Optional. The shared Compute Engine config settings for
	//  all instances in a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.gce_cluster_config
	GCEClusterConfig *GCEClusterConfig `json:"gceClusterConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  the cluster's master instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfig `json:"masterConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  the cluster's worker instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.worker_config
	WorkerConfig *InstanceGroupConfig `json:"workerConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  a cluster's secondary worker instances
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.secondary_worker_config
	SecondaryWorkerConfig *InstanceGroupConfig `json:"secondaryWorkerConfig,omitempty"`

	// Optional. The config settings for cluster software.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.software_config
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	// Optional. Commands to execute on each node after config is
	//  completed. By default, executables are run on master and all worker nodes.
	//  You can test a node's `role` metadata to run an executable on
	//  a master or worker node, as shown below using `curl` (you can also use
	//  `wget`):
	//
	//      ROLE=$(curl -H Metadata-Flavor:Google
	//      http://metadata/computeMetadata/v1/instance/attributes/dataproc-role)
	//      if [[ "${ROLE}" == 'Master' ]]; then
	//        ... master specific actions ...
	//      else
	//        ... worker specific actions ...
	//      fi
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.initialization_actions
	InitializationActions []NodeInitializationAction `json:"initializationActions,omitempty"`

	// Optional. Encryption settings for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Autoscaling config for the policy associated with the cluster.
	//  Cluster does not autoscale if this field is unset.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.autoscaling_config
	AutoscalingConfig *AutoscalingConfig `json:"autoscalingConfig,omitempty"`

	// Optional. Security settings for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.security_config
	SecurityConfig *SecurityConfig `json:"securityConfig,omitempty"`

	// Optional. Lifecycle setting for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfig `json:"lifecycleConfig,omitempty"`

	// Optional. Port/endpoint configuration for this cluster
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfig `json:"endpointConfig,omitempty"`

	// Optional. Metastore configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.metastore_config
	MetastoreConfig *MetastoreConfig `json:"metastoreConfig,omitempty"`

	// Optional. The config for Dataproc metrics.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.dataproc_metric_config
	DataprocMetricConfig *DataprocMetricConfig `json:"dataprocMetricConfig,omitempty"`

	// Optional. The node group settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.auxiliary_node_groups
	AuxiliaryNodeGroups []AuxiliaryNodeGroup `json:"auxiliaryNodeGroups,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.DiskConfig
type DiskConfig struct {
	// Optional. Type of the boot disk (default is "pd-standard").
	//  Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive),
	//  "pd-ssd" (Persistent Disk Solid State Drive),
	//  or "pd-standard" (Persistent Disk Hard Disk Drive).
	//  See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Optional. Size in GB of the boot disk (default is 500GB).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`

	// Optional. Number of attached SSDs, from 0 to 8 (default is 0).
	//  If SSDs are not attached, the boot disk is used to store runtime logs and
	//  [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data.
	//  If one or more SSDs are attached, this runtime bulk
	//  data is spread across them, and the boot disk contains only basic
	//  config and installed binaries.
	//
	//  Note: Local SSD options may vary by machine type and number of vCPUs
	//  selected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.num_local_ssds
	NumLocalSsds *int32 `json:"numLocalSsds,omitempty"`

	// Optional. Interface type of local SSDs (default is "scsi").
	//  Valid values: "scsi" (Small Computer System Interface),
	//  "nvme" (Non-Volatile Memory Express).
	//  See [local SSD
	//  performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.local_ssd_interface
	LocalSsdInterface *string `json:"localSsdInterface,omitempty"`

	// Optional. Indicates how many IOPS to provision for the disk. This sets the
	//  number of I/O operations per second that the disk can handle. Note: This
	//  field is only supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_iops
	BootDiskProvisionedIops *int64 `json:"bootDiskProvisionedIops,omitempty"`

	// Optional. Indicates how much throughput to provision for the disk. This
	//  sets the number of throughput mb per second that the disk can handle.
	//  Values must be greater than or equal to 1. Note: This field is only
	//  supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_throughput
	BootDiskProvisionedThroughput *int64 `json:"bootDiskProvisionedThroughput,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EncryptionConfig
type EncryptionConfig struct {
	// Optional. The Cloud KMS key resource name to use for persistent disk
	//  encryption for all instances in the cluster. See [Use CMEK with cluster
	//  data]
	//  (https://cloud.google.com//dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_cluster_data)
	//  for more information.
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.gce_pd_kms_key_name
	GCEPDKMSKeyName *string `json:"gcePDKMSKeyName,omitempty"`

	// Optional. The Cloud KMS key resource name to use for cluster persistent
	//  disk and job argument encryption. See [Use CMEK with cluster data]
	//  (https://cloud.google.com//dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_cluster_data)
	//  for more information.
	//
	//  When this key resource name is provided, the following job arguments of
	//  the following job types submitted to the cluster are encrypted using CMEK:
	//
	//  * [FlinkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/FlinkJob)
	//  * [HadoopJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/HadoopJob)
	//  * [SparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkJob)
	//  * [SparkRJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkRJob)
	//  * [PySparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/PySparkJob)
	//  * [SparkSqlJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkSqlJob)
	//    scriptVariables and queryList.queries
	//  * [HiveJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/HiveJob)
	//    scriptVariables and queryList.queries
	//  * [PigJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PigJob)
	//    scriptVariables and queryList.queries
	//  * [PrestoJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PrestoJob)
	//    scriptVariables and queryList.queries
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfig struct {

	// Optional. If true, enable http access to specific ports on the cluster
	//  from external sources. Defaults to false.
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.enable_http_port_access
	EnableHTTPPortAccess *bool `json:"enableHTTPPortAccess,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GceClusterConfig
type GCEClusterConfig struct {
	// Optional. The Compute Engine zone where the Dataproc cluster will be
	//  located. If omitted, the service will pick a zone in the cluster's Compute
	//  Engine region. On a get request, zone will always be present.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]`
	//  * `projects/[project_id]/zones/[zone]`
	//  * `[zone]`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.zone_uri
	ZoneURI *string `json:"zoneURI,omitempty"`

	// Optional. The Compute Engine network to be used for machine
	//  communications. Cannot be specified with subnetwork_uri. If neither
	//  `network_uri` nor `subnetwork_uri` is specified, the "default" network of
	//  the project is used, if it exists. Cannot be a "Custom Subnet Network" (see
	//  [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for
	//  more information).
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/networks/default`
	//  * `projects/[project_id]/global/networks/default`
	//  * `default`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Optional. The Compute Engine subnetwork to be used for machine
	//  communications. Cannot be specified with network_uri.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `sub0`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.subnetwork_uri
	SubnetworkURI *string `json:"subnetworkURI,omitempty"`

	// Optional. This setting applies to subnetwork-enabled networks. It is set to
	//  `true` by default in clusters created with image versions 2.2.x.
	//
	//  When set to `true`:
	//
	//  * All cluster VMs have internal IP addresses.
	//  * [Google Private Access]
	//  (https://cloud.google.com/vpc/docs/private-google-access)
	//  must be enabled to access Dataproc and other Google Cloud APIs.
	//  * Off-cluster dependencies must be configured to be accessible
	//  without external IP addresses.
	//
	//  When set to `false`:
	//
	//  * Cluster VMs are not restricted to internal IP addresses.
	//  * Ephemeral external IP addresses are assigned to each cluster VM.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.internal_ip_only
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	// Optional. The type of IPv6 access for a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.private_ipv6_google_access
	PrivateIPV6GoogleAccess *string `json:"privateIPV6GoogleAccess,omitempty"`

	// Optional. The [Dataproc service
	//  account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc)
	//  (also see [VM Data Plane
	//  identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity))
	//  used by Dataproc cluster VM instances to access Google Cloud Platform
	//  services.
	//
	//  If not specified, the
	//  [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. The URIs of service account scopes to be included in
	//  Compute Engine instances. The following base set of scopes is always
	//  included:
	//
	//  * https://www.googleapis.com/auth/cloud.useraccounts.readonly
	//  * https://www.googleapis.com/auth/devstorage.read_write
	//  * https://www.googleapis.com/auth/logging.write
	//
	//  If no scopes are specified, the following defaults are also provided:
	//
	//  * https://www.googleapis.com/auth/bigquery
	//  * https://www.googleapis.com/auth/bigtable.admin.table
	//  * https://www.googleapis.com/auth/bigtable.data
	//  * https://www.googleapis.com/auth/devstorage.full_control
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account_scopes
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// The Compute Engine network tags to add to all instances (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. The Compute Engine metadata entries to add to all instances (see
	//  [Project and instance
	//  metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. Reservation Affinity for consuming Zonal reservation.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Optional. Node Group Affinity for sole-tenant clusters.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.node_group_affinity
	NodeGroupAffinity *NodeGroupAffinity `json:"nodeGroupAffinity,omitempty"`

	// Optional. Shielded Instance Config for clusters using [Compute Engine
	//  Shielded
	//  VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Optional. Confidential Instance Config for clusters using [Confidential
	//  VMs](https://cloud.google.com/compute/confidential-vm/docs).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.confidential_instance_config
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeClusterConfig
type GkeClusterConfig struct {
	// Optional. A target GKE cluster to deploy to. It must be in the same project
	//  and region as the Dataproc cluster (the GKE cluster can be zonal or
	//  regional). Format:
	//  'projects/{project}/locations/{location}/clusters/{cluster_id}'
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.gke_cluster_target
	GkeClusterTarget *string `json:"gkeClusterTarget,omitempty"`

	// Optional. GKE node pools where workloads will be scheduled. At least one
	//  node pool must be assigned the `DEFAULT`
	//  [GkeNodePoolTarget.Role][google.cloud.dataproc.v1.GkeNodePoolTarget.Role].
	//  If a `GkeNodePoolTarget` is not specified, Dataproc constructs a `DEFAULT`
	//  `GkeNodePoolTarget`. Each role can be given to only one
	//  `GkeNodePoolTarget`. All node pools must have the same location settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.node_pool_target
	NodePoolTarget []GkeNodePoolTarget `json:"nodePoolTarget,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfig struct {
	// Optional. The number of VM instances in the instance group.
	//  For [HA
	//  cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
	//  [master_config](#FIELDS.master_config) groups, **must be set to 3**.
	//  For standard cluster [master_config](#FIELDS.master_config) groups,
	//  **must be set to 1**.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.num_instances
	NumInstances *int32 `json:"numInstances,omitempty"`

	// Optional. The Compute Engine image resource used for cluster instances.
	//
	//  The URI can represent an image or image family.
	//
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
	//  If the URI is unspecified, it will be inferred from
	//  `SoftwareConfig.image_version` or the system default.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Optional. The Compute Engine machine type used for cluster instances.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `n1-standard-2`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the machine type
	//  resource, for example, `n1-standard-2`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.machine_type_uri
	MachineTypeURI *string `json:"machineTypeURI,omitempty"`

	// Optional. Disk option config settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.disk_config
	DiskConfig *DiskConfig `json:"diskConfig,omitempty"`

	// Optional. Specifies the preemptibility of the instance group.
	//
	//  The default value for master and worker groups is
	//  `NON_PREEMPTIBLE`. This default cannot be changed.
	//
	//  The default value for secondary instances is
	//  `PREEMPTIBLE`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.preemptibility
	Preemptibility *string `json:"preemptibility,omitempty"`

	// Optional. The Compute Engine accelerator configuration for these
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.accelerators
	Accelerators []AcceleratorConfig `json:"accelerators,omitempty"`

	// Optional. Specifies the minimum cpu platform for the Instance Group.
	//  See [Dataproc -> Minimum CPU
	//  Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`

	// Optional. The minimum number of primary worker instances to create.
	//  If `min_num_instances` is set, cluster creation will succeed if
	//  the number of primary workers created is at least equal to the
	//  `min_num_instances` number.
	//
	//  Example: Cluster creation request with `num_instances` = `5` and
	//  `min_num_instances` = `3`:
	//
	//  *  If 4 VMs are created and 1 instance fails,
	//     the failed VM is deleted. The cluster is
	//     resized to 4 instances and placed in a `RUNNING` state.
	//  *  If 2 instances are created and 3 instances fail,
	//     the cluster in placed in an `ERROR` state. The failed VMs
	//     are not deleted.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_num_instances
	MinNumInstances *int32 `json:"minNumInstances,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicy `json:"instanceFlexibilityPolicy,omitempty"`

	// Optional. Configuration to handle the startup of instances during cluster
	//  create and update process.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.startup_config
	StartupConfig *StartupConfig `json:"startupConfig,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.KerberosConfig
type KerberosConfig struct {
	// Optional. Flag to indicate whether to Kerberize the cluster (default:
	//  false). Set this field to true to enable Kerberos on a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.enable_kerberos
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the root
	//  principal password.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.root_principal_password_uri
	RootPrincipalPasswordURI *string `json:"rootPrincipalPasswordURI,omitempty"`

	// Optional. The URI of the KMS key used to encrypt sensitive
	//  files.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kms_key_uri
	KMSKeyURI *string `json:"kmsKeyURI,omitempty"`

	// Optional. The Cloud Storage URI of the keystore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_uri
	KeystoreURI *string `json:"keystoreURI,omitempty"`

	// Optional. The Cloud Storage URI of the truststore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_uri
	TruststoreURI *string `json:"truststoreURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided keystore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_password_uri
	KeystorePasswordURI *string `json:"keystorePasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided key. For the self-signed certificate, this
	//  password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.key_password_uri
	KeyPasswordURI *string `json:"keyPasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided truststore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_password_uri
	TruststorePasswordURI *string `json:"truststorePasswordURI,omitempty"`

	// Optional. The remote realm the Dataproc on-cluster KDC will trust, should
	//  the user enable cross realm trust.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_realm
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty"`

	// Optional. The KDC (IP or hostname) for the remote trusted realm in a cross
	//  realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_kdc
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty"`

	// Optional. The admin server (IP or hostname) for the remote trusted realm in
	//  a cross realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_admin_server
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  shared password between the on-cluster Kerberos realm and the remote
	//  trusted realm, in a cross realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_shared_password_uri
	CrossRealmTrustSharedPasswordURI *string `json:"crossRealmTrustSharedPasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  master key of the KDC database.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kdc_db_key_uri
	KdcDbKeyURI *string `json:"kdcDbKeyURI,omitempty"`

	// Optional. The lifetime of the ticket granting ticket, in hours.
	//  If not specified, or user specifies 0, then default value 10
	//  will be used.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.tgt_lifetime_hours
	TgtLifetimeHours *int32 `json:"tgtLifetimeHours,omitempty"`

	// Optional. The name of the on-cluster Kerberos realm.
	//  If not specified, the uppercased domain of hostnames will be the realm.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.realm
	Realm *string `json:"realm,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.MetastoreConfig
type MetastoreConfig struct {
	// Required. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.MetastoreConfig.dataproc_metastore_service
	DataprocMetastoreService *string `json:"dataprocMetastoreService,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.NodeGroupAffinity
type NodeGroupAffinity struct {
	// Required. The URI of a
	//  sole-tenant [node group
	//  resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
	//  that the cluster will be created on.
	//
	//  A full URL, partial URI, or node group name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
	//  * `projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
	//  * `node-group-1`
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroupAffinity.node_group_uri
	NodeGroupURI *string `json:"nodeGroupURI,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Optional. Defines whether instances have Secure Boot enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Optional. Defines whether instances have the vTPM enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`

	// Optional. Defines whether instances have integrity monitoring enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.SparkHistoryServerConfig
type SparkHistoryServerConfig struct {
	// Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	//  History Server for the workload.
	//
	//  Example:
	//
	//  * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkHistoryServerConfig.dataproc_cluster
	DataprocCluster *string `json:"dataprocCluster,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.Cluster
type ClusterObservedState struct {
	// Optional. The cluster config for a cluster of Compute Engine Instances.
	//  Note that Dataproc may set default values, and values may change
	//  when clusters are updated.
	//
	//  Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.config
	Config *ClusterConfigObservedState `json:"config,omitempty"`

	// Output only. Cluster status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.status
	Status *ClusterStatus `json:"status,omitempty"`

	// Output only. The previous cluster status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.status_history
	StatusHistory []ClusterStatus `json:"statusHistory,omitempty"`

	// Output only. A cluster UUID (Unique Universal Identifier). Dataproc
	//  generates this value when it creates the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.cluster_uuid
	ClusterUuid *string `json:"clusterUuid,omitempty"`

	// Output only. Contains cluster daemon metrics such as HDFS and YARN stats.
	//
	//  **Beta Feature**: This report is available for testing purposes only. It
	//  may be changed before final release.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.metrics
	Metrics *ClusterMetrics `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfigObservedState struct {
	// Optional. The Compute Engine config settings for
	//  the cluster's master instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfigObservedState `json:"masterConfig,omitempty"`

	// Optional. Lifecycle setting for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfigObservedState `json:"lifecycleConfig,omitempty"`

	// Optional. Port/endpoint configuration for this cluster
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfigObservedState `json:"endpointConfig,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfigObservedState struct {
	// Output only. The map of port descriptions to URLs. Will only be populated
	//  if enable_http_port_access is true.
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.http_ports
	HTTPPorts map[string]string `json:"httpPorts,omitempty"`
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

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfigObservedState struct {
	// Output only. The list of instance names. Dataproc derives the names
	//  from `cluster_name`, `num_instances`, and the instance group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_names
	InstanceNames []string `json:"instanceNames,omitempty"`

	// Output only. List of references to Compute Engine instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_references
	InstanceReferences []InstanceReference `json:"instanceReferences,omitempty"`

	// Output only. Specifies that this instance group contains preemptible
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.is_preemptible
	IsPreemptible *bool `json:"isPreemptible,omitempty"`

	// Output only. The config for Compute Engine Instance Group
	//  Manager that manages this group.
	//  This is only used for preemptible instance groups.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.managed_group_config
	ManagedGroupConfig *ManagedGroupConfig `json:"managedGroupConfig,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicyObservedState `json:"instanceFlexibilityPolicy,omitempty"`
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
