package composerenvironment

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/composer/v1alpha1"
	composer "google.golang.org/api/composer/v1"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

var mapping = NewMapping(&v1alpha1.ComposerEnvironment{}, &composer.Environment{},
	Spec("config"),
	ResourceID("name"),
	Status("state"),
	Ignore("uuid"),
	// Ignore("createTime"),
	// Ignore("status.state"),
).
	MapType(&v1alpha1.EnvironmentConfig{}, &composer.EnvironmentConfig{},
		"nodeConfig",
		OutputOnly("airflowByoidUri"),
		OutputOnly("airflowUri"),
		OutputOnly("dagGcsPrefix"),
		TODO("databaseConfig"),
		TODO("encryptionConfig"), // TODO: also immutable
		"environmentSize",
		OutputOnly("gkeCluster"),
		TODO("maintenanceWindow"),
		TODO("masterAuthorizedNetworksConfig"),
		TODO("nodeConfig"),
		TODO("nodeCount"),
		TODO("privateEnvironmentConfig"),
		TODO("recoveryConfig"),
		TODO("resilienceMode"),
		TODO("softwareConfig"),
		TODO("webServerConfig"),
		TODO("webServerNetworkAccessControl"),
		TODO("workloadsConfig"),
	).
	MapType(&v1alpha1.EnvironmentNodeConfig{}, &composer.NodeConfig{},
		"diskSizeGb",
		"enableIpMasqAgent",
		"ipAllocationPolicy",
		"machineType",
		"maxPodsPerNode",
		"network", // todo: ref
		"oauthScopes",
		Ref("serviceAccount"),
		"subnetwork", // todo: ref
		"tags",
		"zone", // TODO: ref/normalize?
	).
	MapType(&v1alpha1.EnvironmentWebServerConfig{}, &composer.WebServerConfig{},
		"machineType", // TODO: Only supported in airflow v1
	).
	MustBuild()
