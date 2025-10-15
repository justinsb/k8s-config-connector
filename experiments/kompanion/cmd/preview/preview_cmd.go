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

package preview

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/preview"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	kubeconfigFlag = "kubeconfig"
	timeoutFlag    = "timeout"

	reportNamePrefixFlag = "report-prefix"
	defaultScope         = "https://www.googleapis.com/auth/cloud-platform"
)

const (
	examples = `
	# preview Config Connector resources
	kompanion preview
	`
)

type PreviewOptions struct {
	kubeconfig       string
	timeout          int
	reportNamePrefix string
	fullReport       bool
}

func BuildPreviewCmd() *cobra.Command {
	var opts PreviewOptions
	cmd := &cobra.Command{
		Use:     "preview",
		Short:   "preview Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunPreview(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}
	// TODO: Add scope
	cmd.Flags().StringVarP(&opts.kubeconfig, kubeconfigFlag, "", opts.kubeconfig, "path to the kubeconfig file.")
	cmd.Flags().IntVarP(&opts.timeout, timeoutFlag, "", 15, "timeout in minutes. Default to 15 minutes.")
	cmd.Flags().StringVarP(&opts.reportNamePrefix, reportNamePrefixFlag, "", "preview-report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
	cmd.Flags().BoolVarP(&opts.fullReport, "full-report", "f", false, "Enable verbose logging.")

	return cmd
}

func getRESTConfig(ctx context.Context, opts *PreviewOptions) (*rest.Config, error) {
	// TODO: Add rate limiting.
	var loadingRules clientcmd.ClientConfigLoader
	if opts.kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func getGCPAuthorization(ctx context.Context, opts *PreviewOptions) (oauth2.TokenSource, error) {
	// TODO: Add scope
	scopes := []string{defaultScope}
	ts, err := google.DefaultTokenSource(ctx, scopes...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func RunPreview(ctx context.Context, opts *PreviewOptions) error {
	log.SetLogger(klogr.New())
	klog.Info("Starting preview tool.")
	upstreamRESTConfig, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}
	recorder := preview.NewRecorder()
	klog.Info("Preloading the list of resources to reconcile")
	if err := recorder.PreloadGKNN(ctx, upstreamRESTConfig); err != nil {
		return fmt.Errorf("error preload the list of resources to reconcile: %w", err)
	}
	klog.Info("Successfully preload the list of resources to reconcile.")
	authorization, err := getGCPAuthorization(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building GCP authorization: %w", err)
	}
	var upstreamGCPHTTPClient *http.Client
	if true /* HACK */ {
		upstreamGCPHTTPClient = buildSnapshotHTTPClient()
	}
	preview, err := preview.NewPreviewInstance(recorder, preview.PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    upstreamGCPHTTPClient,
	})
	if err != nil {
		return fmt.Errorf("building preview instance: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(opts.timeout)*time.Minute)
	defer printCapturedObjects(recorder, opts.reportNamePrefix, opts.fullReport)
	defer cancel()

	errChan := make(chan error, 1)
	go func() {
		klog.Info("Starting preview")
		errChan <- preview.Start(ctx)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("error starting preview: %w", err)
		}
	case <-ctx.Done():
		return fmt.Errorf("timeout reached: %w", ctx.Err())
	}
	klog.Info("Preview finished successfully")
	return nil
}

func printCapturedObjects(recorder *preview.Recorder, prefix string, full bool) error {
	now := time.Now()
	timestamp := now.Format("20060102-150405.000")
	summaryFile := fmt.Sprintf("%s-%s", prefix, timestamp)
	if err := recorder.SummaryReport(summaryFile); err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}

	if full {
		outputFile := fmt.Sprintf("%s-full-%s", prefix, timestamp)
		if err := recorder.ExportDetailObjectsEvent(outputFile); err != nil {
			return fmt.Errorf("error writing events: %w", err)
		}

	}
	return nil
}

func buildSnapshotHTTPClient() *http.Client {
	roundTripper := &snapshotRoundTripper{}

	return &http.Client{
		Transport: roundTripper,
		Timeout:   30 * time.Second,
	}
}

type snapshotRoundTripper struct {
}

var _ http.RoundTripper = &snapshotRoundTripper{}

func (s *snapshotRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	log := klog.FromContext(ctx)
	log.Info("snapshotRoundTripper", "req.method", req.Method, "req.url", req.URL.String())

	baseDir := "/usr/local/google/home/justinsb/b451536623/k8s-config-connector/gcpsnapshot"

	reqPath := req.URL.Path
	p := filepath.Join(baseDir, req.Host, reqPath+"."+strings.ToLower(req.Method))
	log.Info("looking for snapshot file", "path", p, "request.path", reqPath)

	// TODO: Sanitize to be a valid filename
	b, err := os.ReadFile(p)
	if err != nil {
		// TODO
		log.Error(err, "error reading snapshot file", "path", p)
	}

	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Header:     make(http.Header),
	}

	if b != nil {
		log.Info("snapshot found", "path", p)

		response.Body = io.NopCloser(strings.NewReader(string(b)))
		response.Header.Set("Content-Type", "application/json; charset=UTF-8")
		return response, nil
	}

	log.Info("snapshot NOT found, returning 401", "path", p)
	response.Status = "401 Unauthorized (FAKE)"
	response.StatusCode = 401

	return response, nil
}
