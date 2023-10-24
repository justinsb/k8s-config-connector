package configconnector

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/restmapper"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/httprecorder"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/testharness"

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

func runGoldenReconcileTest(h *testharness.Harness, testdir string) {
	ctx := context.Background()

	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		h.Fatalf("error building mock kube-apiserver: %v", err)
	}
	defer func() {
		if err := k8s.Stop(); err != nil {
			h.Fatalf("error closing mock kube-apiserver: %v", err)
		}
	}()

	// k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}, "namespaces", meta.RESTScopeRoot)
	// k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, "configmaps", meta.RESTScopeNamespace)
	// k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Event"}, "events", meta.RESTScopeNamespace)
	// k8s.RegisterType(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}, "deployments", meta.RESTScopeNamespace)
	// k8s.RegisterType(schema.GroupVersionKind{Group: "addons.example.org", Version: "v1alpha1", Kind: "SimpleTest"}, "simpletests", meta.RESTScopeNamespace)

	// TODO: Register CRDs?
	// k8s.RegisterType(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"}, "configconnectors", meta.RESTScopeRoot)
	// k8s.RegisterType(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"}, "configconnectorcontexts", meta.RESTScopeNamespace)

	crdDir := "../../../config/crd/bases"
	crdFiles, err := os.ReadDir(crdDir)
	if err != nil {
		h.Fatalf("error reading directory %q: %v", crdDir, err)
	}
	for _, crdFile := range crdFiles {
		p := filepath.Join(crdDir, crdFile.Name())
		h.Logf("loading objects from %v", p)
		before := string(h.MustReadFile(p))
		if err := k8s.AddObjectsFromManifest(before); err != nil {
			h.Fatalf("error creating crd: %v", err)
		}
	}

	addr, err := k8s.StartServing()
	if err != nil {
		h.Errorf("error starting mock kube-apiserver: %v", err)
	}

	klog.Infof("mock kubeapiserver will listen on %v", addr)

	var requestLog httprecorder.RequestLog
	wrapTransport := func(rt http.RoundTripper) http.RoundTripper {
		return httprecorder.NewRecorder(rt, &requestLog)
	}
	restConfig := &rest.Config{
		Host:          addr.String(),
		WrapTransport: wrapTransport,
		QPS:           10000,
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
	}

	scheme := runtime.NewScheme()
	if err := customizev1alpha1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := customizev1beta1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := corev1beta1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := corev1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := corev1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := appsv1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}
	if err := rbacv1.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}

	logger := klogr.New()
	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme: scheme,
		// Metrics: metricsserver.Options{
		// 	BindAddress: "0", // Disable the metrics server
		// },
		MetricsBindAddress: "0",
		LeaderElection:     false,

		// MapperProvider provides the rest mapper used to map go types to Kubernetes APIs
		MapperProvider: restmapper.NewControllerRESTMapper,

		Logger: logger,
	})
	if err != nil {
		h.Fatalf("error starting manager: %v", err)
	}

	// reconciler := &SimpleTestReconciler{
	// 	Client:  mgr.GetClient(),
	// 	Scheme:  mgr.GetScheme(),
	// 	applier: applier,
	// 	status:  status,
	// }
	repoPath := "../../../channels"
	reconciler, err := newReconciler(mgr, repoPath)
	if err != nil {
		h.Fatalf("error from newReconciler: %v", err)
	}
	// mc, err := loaders.NewManifestLoader("testdata/channels")
	// if err != nil {
	// 	h.Fatalf("error from loaders.NewManifestLoader: %v", err)
	// }
	// reconciler.manifestController = mc

	if err = reconciler.SetupWithManager(mgr); err != nil {
		h.Fatalf("error creating reconciler: %v", err)
	}

	if p := filepath.Join(testdir, "before.yaml"); h.FileExists(p) {
		h.Logf("loading objects from %v", p)
		before := string(h.MustReadFile(p))
		if err := k8s.AddObjectsFromManifest(before); err != nil {
			h.Fatalf("error precreating objects: %v", err)
		}
	}

	mgrContext, mgrStop := context.WithCancel(ctx)
	go func() {
		time.Sleep(10 * time.Second)
		// time.Sleep(200 * time.Second)
		mgrStop()
	}()
	if err := mgr.Start(mgrContext); err != nil {
		h.Fatalf("error starting manager: %v", err)
	}

	h.Logf("replacing old url prefix %q", "http://"+restConfig.Host)
	requestLog.ReplaceURLPrefix("http://"+restConfig.Host, "http://kube-apiserver")
	requestLog.RemoveUserAgent()
	// requestLog.SortGETs()
	// Workaround for non-determinism in https://github.com/kubernetes/kubernetes/blob/79a62d62350fb600f97d1f6309c3274515b3587a/staging/src/k8s.io/client-go/tools/cache/reflector.go#L301
	requestLog.RegexReplaceURL("&timeoutSeconds=.*&", "&timeoutSeconds=<replaced>&")
	h.Logf("replacing real timestamp in request and response to a fake value")
	// requestLog.ReplaceTimestamp()

	requests := requestLog.FormatHTTP()

	h.CompareGoldenFile(filepath.Join(testdir, "expected-http.yaml"), requests)
}

func TestGolden(t *testing.T) {
	testharness.RunGoldenTests(t, "testdata/reconcile/", func(h *testharness.Harness, testdir string) {
		runGoldenReconcileTest(h, testdir)
	})
}
