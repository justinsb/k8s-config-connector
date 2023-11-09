package mocks

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
)

func GetMockClient(t *testing.T) client.Client {
	ctx := context.TODO()

	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		t.Fatalf("error building mock kube-apiserver: %v", err)
	}

	addr, err := k8s.StartServing()
	if err != nil {
		t.Errorf("error starting mock kube-apiserver: %v", err)
	}

	if addr == nil {
		t.Fatalf("address of the mock kube-apiserver is nil")
	}

	t.Cleanup(func() {
		if err := k8s.Stop(); err != nil {
			t.Errorf("error stopping mock kube-apiserver: %v", err)
		}
	})

	restConfig := &rest.Config{
		Host: addr.String(),
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
	}

	scheme := runtime.NewScheme()
	if err := v1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}
	// TODO: Create an envtest style CRD loader ... that's pretty nice!
	k8s.RegisterType(v1beta1.ConfigConnectorContextGroupVersionKind, "configconnectorcontexts", meta.RESTScopeNamespace)
	k8s.RegisterType(v1beta1.ConfigConnectorGroupVersionKind, "configconnectors", meta.RESTScopeRoot)

	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}

	mgr, err := manager.New(restConfig, manager.Options{
		Scheme:                 scheme,
		HealthProbeBindAddress: "0",
		MetricsBindAddress:     "0",
	})
	if err != nil {
		t.Fatalf("error building manager: %v", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		if err := mgr.Start(ctx); err != nil {
			t.Logf("failed to start manager: %v", err)
		}
	}()

	t.Cleanup(func() {
		cancel()
	})

	if !mgr.GetCache().WaitForCacheSync(ctx) {
		t.Fatalf("cache did not sync")
	}

	return mgr.GetClient()
}
