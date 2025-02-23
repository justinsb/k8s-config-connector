package v1beta1

import (
	compute "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.dataproc.v1.AutoscalingConfig
type AutoscalingConfig struct {
	// Optional. The autoscaling policy used by the cluster.
	//
	//  Note that the policy must be in the same project and Dataproc region.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingConfig.policy_uri
	PolicyRef *AutoscalingPolicyRef `json:"policyRef,omitempty"`
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
	ConfigBucket *refs.StorageBucketRef `json:"stagingBucketRef,omitempty"`

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
	TempBucket *refs.StorageBucketRef `json:"tempBucketRef,omitempty"`

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

	/* NOT_IN_LEGACY
	// Optional. The node group settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.auxiliary_node_groups
	AuxiliaryNodeGroups []AuxiliaryNodeGroup `json:"auxiliaryNodeGroups,omitempty"`
	*/
}

// +kcc:proto=google.cloud.dataproc.v1.EncryptionConfig
type EncryptionConfig struct {
	// Optional. The Cloud KMS key resource name to use for persistent disk
	//  encryption for all instances in the cluster. See [Use CMEK with cluster
	//  data]
	//  (https://cloud.google.com//dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_cluster_data)
	//  for more information.
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.gce_pd_kms_key_name
	GCEPDKMSKeyRef *refs.KMSCryptoKeyRef `json:"gcePdKmsKeyRef,omitempty"`

	/* NOT_IN_LEGACY
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
	*/
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
	ZoneURI *string `json:"zone,omitempty"`

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
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The Compute Engine subnetwork to be used for machine
	//  communications. Cannot be specified with network_uri.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `sub0`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.subnetwork_uri
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

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
	PrivateIPV6GoogleAccess *string `json:"privateIPv6GoogleAccess,omitempty"`

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
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

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
	NodeGroupRef *compute.NodeGroupRef `json:"nodeGroupRef,omitempty"`
}

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
	AcceleratorTypeURI *string `json:"acceleratorType,omitempty"`

	// The number of the accelerator cards of this type exposed to this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
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
	BootDiskSizeGB *int32 `json:"bootDiskSizeGb,omitempty"`

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

	/* NOT_IN_LEGACY
	// Optional. Indicates how many IOPS to provision for the disk. This sets the
	//  number of I/O operations per second that the disk can handle. Note: This
	//  field is only supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_iops
	BootDiskProvisionedIops *int64 `json:"bootDiskProvisionedIops,omitempty"`
	*/

	/* NOT_IN_LEGACY
	// Optional. Indicates how much throughput to provision for the disk. This
	//  sets the number of throughput mb per second that the disk can handle.
	//  Values must be greater than or equal to 1. Note: This field is only
	//  supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_throughput
	BootDiskProvisionedThroughput *int64 `json:"bootDiskProvisionedThroughput,omitempty"`
	*/
}

// +kcc:proto=google.cloud.dataproc.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Optional. Defines whether instances have Secure Boot enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Optional. Defines whether instances have the vTPM enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVtpm,omitempty"`

	// Optional. Defines whether instances have integrity monitoring enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfig struct {

	// Optional. If true, enable http access to specific ports on the cluster
	//  from external sources. Defaults to false.
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.enable_http_port_access
	EnableHTTPPortAccess *bool `json:"enableHttpPortAccess,omitempty"`
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
	ImageURI *compute.ImageRef `json:"imageRef,omitempty"`

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
	MachineTypeURI *string `json:"machineType,omitempty"`

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
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	/* NOT_IN_LEGACY
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
	*/

	/* NOT_IN_LEGACY
	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicy `json:"instanceFlexibilityPolicy,omitempty"`
	*/

	/* NOT_IN_LEGACY
	// Optional. Configuration to handle the startup of instances during cluster
	//  create and update process.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.startup_config
	StartupConfig *StartupConfig `json:"startupConfig,omitempty"`
	*/
}

// +kcc:proto=google.cloud.dataproc.v1.MetastoreConfig
type MetastoreConfig struct {
	// Required. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.MetastoreConfig.dataproc_metastore_service
	DataprocMetastoreService *DataprocServiceRef `json:"dataprocMetastoreServiceRef,omitempty"`
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
	RootPrincipalPasswordURI *string `json:"rootPrincipalPassword,omitempty"`

	// Optional. The URI of the KMS key used to encrypt sensitive
	//  files.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kms_key_uri
	KMSKeyURI *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Optional. The Cloud Storage URI of the keystore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_uri
	KeystoreURI *string `json:"keystore,omitempty"`

	// Optional. The Cloud Storage URI of the truststore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_uri
	TruststoreURI *string `json:"truststore,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided keystore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_password_uri
	KeystorePasswordURI *string `json:"keystorePassword,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided key. For the self-signed certificate, this
	//  password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.key_password_uri
	KeyPasswordURI *string `json:"keyPassword,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided truststore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_password_uri
	TruststorePasswordURI *string `json:"truststorePassword,omitempty"`

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
	CrossRealmTrustSharedPasswordURI *string `json:"crossRealmTrustSharedPassword,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  master key of the KDC database.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kdc_db_key_uri
	KdcDbKeyURI *string `json:"kdcDbKey,omitempty"`

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

// +kcc:proto=google.cloud.dataproc.v1.SparkHistoryServerConfig
type SparkHistoryServerConfig struct {
	// Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	//  History Server for the workload.
	//
	//  Example:
	//
	//  * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkHistoryServerConfig.dataproc_cluster
	DataprocCluster *refs.DataprocClusterRef `json:"dataprocClusterRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeClusterConfig
type GkeClusterConfig struct {
	// Optional. A target GKE cluster to deploy to. It must be in the same project
	//  and region as the Dataproc cluster (the GKE cluster can be zonal or
	//  regional). Format:
	//  'projects/{project}/locations/{location}/clusters/{cluster_id}'
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.gke_cluster_target
	GkeClusterTarget *container.ClusterRef `json:"gkeClusterTargetRef,omitempty"`

	// Optional. GKE node pools where workloads will be scheduled. At least one
	//  node pool must be assigned the `DEFAULT`
	//  [GkeNodePoolTarget.Role][google.cloud.dataproc.v1.GkeNodePoolTarget.Role].
	//  If a `GkeNodePoolTarget` is not specified, Dataproc constructs a `DEFAULT`
	//  `GkeNodePoolTarget`. Each role can be given to only one
	//  `GkeNodePoolTarget`. All node pools must have the same location settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.node_pool_target
	NodePoolTarget []GkeNodePoolTarget `json:"nodePoolTarget,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolTarget
type GkeNodePoolTarget struct {
	// Required. The target GKE node pool.
	//  Format:
	//  'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}'
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.node_pool
	NodePoolRef *container.NodePoolRef `json:"nodePoolRef,omitempty"`

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

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig
type GkeNodePoolConfig_GkeNodeConfig struct {
	// Optional. The name of a Compute Engine [machine
	//  type](https://cloud.google.com/compute/docs/machine-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	/* NOT_IN_LEGACY
	// Optional. The number of local SSD disks to attach to the node, which is
	//  limited by the maximum number of disks allowable per zone (see [Adding
	//  Local SSDs](https://cloud.google.com/compute/docs/disks/local-ssd)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.local_ssd_count
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
	*/

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
	BootDiskKMSKey *string `json:"bootDiskKmsKey,omitempty"`

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
