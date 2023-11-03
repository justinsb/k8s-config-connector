// Copyright 2022 Google LLC
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

package e2e

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestAllInSeries(t *testing.T) {
	ctx := test.WithContext(context.Background(), t)

	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	var project testgcp.GCPProject
	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		projectNumber := time.Now().Unix()
		project = testgcp.GCPProject{
			ProjectID:     "mock-project-" + strconv.FormatInt(projectNumber, 10),
			ProjectNumber: projectNumber,
		}
		// Some fixed-value fake org-ids for testing.
		// We used fixed values so that the output is predictable (for golden testing)
		testgcp.TestOrgID.Set("123450001")
		testgcp.TestBillingAccountID.Set("123456-777777-000001")
		testgcp.IAMIntegrationTestsOrganizationID.Set("123450002")
		testgcp.IAMIntegrationTestsBillingAccountID.Set("123456-777777-000002")
	} else {
		project = testgcp.GetDefaultProject(t)
	}

	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	eventSinkMux := test.NewEventSinkMux()
	ctx = test.AddSinkToContext(ctx, eventSinkMux)

	testHarness := create.NewHarness(t, ctx)

	t.Run("samples", func(t *testing.T) {
		samples := create.LoadSamples(t, project)

		for _, s := range samples {
			s := s
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(s.Name, func(t *testing.T) {
				create.MaybeSkip(t, s.Name, s.Resources)

				h := testHarness.ForSubtest(t)

				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s}, project)

				// Hack: set project-id because mockkubeapiserver does not support webhooks
				for _, u := range s.Resources {
					annotations := u.GetAnnotations()
					if annotations == nil {
						annotations = make(map[string]string)
					}
					annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
					u.SetAnnotations(annotations)
				}

				create.RunCreateDeleteTest(h, create.CreateDeleteTestOptions{Create: s.Resources, CleanupResources: true})
			})
		}
	})

	t.Run("fixtures", func(t *testing.T) {
		fixtures := resourcefixture.Load(t)
		for _, fixture := range fixtures {
			fixture := fixture
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			uniqueID := testvariable.NewUniqueId()

			s := create.Sample{
				Name: fixture.Name,
			}

			createResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)
			s.Resources = append(s.Resources, createResource)

			exportResources := []*unstructured.Unstructured{createResource}

			if fixture.Dependencies != nil {
				dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
				for _, dependBytes := range dependencyYamls {
					depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
					s.Resources = append(s.Resources, depUnstruct)
				}
			}

			opt := create.CreateDeleteTestOptions{Create: s.Resources, CleanupResources: true}
			if fixture.Update != nil {
				u := bytesToUnstructured(t, fixture.Update, uniqueID, project)
				opt.Updates = append(opt.Updates, u)
			}

			t.Run(s.Name, func(t *testing.T) {
				create.MaybeSkip(t, s.Name, s.Resources)

				h := testHarness.ForSubtest(t)

				eventSink := test.NewMemoryEventSink()
				eventSinkMux.SetEventSink(eventSink)

				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s}, project)

				opt.CleanupResources = false // We delete explicitly below
				create.RunCreateDeleteTest(h, opt)

				for _, exportResource := range exportResources {
					exportURI := buildExportURI(t, exportResource, project)

					if exportURI == "" {
						continue
					}

					exportParams := h.ExportParams()
					exportParams.IAMFormat = "partialpolicy"
					exportParams.ResourceFormat = "krm"
					outputDir := h.TempDir()
					outputPath := filepath.Join(outputDir, "export.yaml")
					exportParams.Output = outputPath
					exportParams.URI = exportURI
					if err := export.Execute(h.Ctx, &exportParams); err != nil {
						t.Errorf("error from export.Execute: %v", err)
						continue
					}

					expectedPath := filepath.Join(fixture.SourceDir, "export.yaml")
					output := h.MustReadFile(outputPath)
					h.CompareGoldenFile(expectedPath, string(output),
						h.IgnoreComments,
						h.ReplaceString(project.ProjectID, "example-project-id"),
						h.ReplaceString(uniqueID, "${uniqueId}"))
				}

				create.DeleteResources(h, s.Resources)

				// Verify events against golden file
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					// TODO: Fix how we poll / wait for objects being ready.
					eventSink.RemoveRequests(func(e *test.LogEntry) bool {
						if e.Response.StatusCode == 404 && e.Request.Method == "GET" {
							return true
						}
						return false
					})

					eventSink.RemoveRequests(func(e *test.LogEntry) bool {
						if strings.Contains(e.Request.URL, "/operations/") {
							return true
						}
						return false
					})

					eventSink.RemoveRequests(func(e *test.LogEntry) bool {
						if !strings.Contains(e.Request.URL, "https://composer.googleapis.com/") {
							return true
						}
						return false
					})

					jsonMutators := []test.JSONMutator{}

					jsonMutators = append(jsonMutators, func(obj map[string]any) {
						_, found, _ := unstructured.NestedString(obj, "uniqueId")
						if found {
							unstructured.SetNestedField(obj, "111111111111111111111", "uniqueId")
						}
					})
					jsonMutators = append(jsonMutators, func(obj map[string]any) {
						_, found, _ := unstructured.NestedString(obj, "oauth2ClientId")
						if found {
							unstructured.SetNestedField(obj, "888888888888888888888", "oauth2ClientId")
						}
					})
					jsonMutators = append(jsonMutators, func(obj map[string]any) {
						_, found, _ := unstructured.NestedString(obj, "etag")
						if found {
							unstructured.SetNestedField(obj, "abcdef0123A=", "etag")
						}
					})

					for _, timeField := range []string{"createTime", "updateTime", "metadata.createTime"} {
						path := strings.Split(timeField, ".")
						jsonMutators = append(jsonMutators, func(obj map[string]any) {
							s, found, _ := unstructured.NestedString(obj, path...)
							if found {
								s = regexp.MustCompile(`\d`).ReplaceAllString(s, "0")
								unstructured.SetNestedField(obj, s, path...)
							}
						})
					}

					uuids := make(map[string]bool)
					jsonMutators = append(jsonMutators, func(obj map[string]any) {
						v, _, _ := unstructured.NestedString(obj, "uuid")
						if v != "" {
							uuids[v] = true
						}
					})
					eventSink.PrettifyJSON(jsonMutators...)

					eventSink.RemoveHTTPResponseHeader("Date")
					eventSink.RemoveHTTPResponseHeader("Alt-Svc")
					got := eventSink.FormatHTTP()
					expectedPath := filepath.Join(fixture.SourceDir, "_http.log")
					normalizers := []func(string) string{}
					normalizers = append(normalizers, h.IgnoreComments)
					normalizers = append(normalizers, h.ReplaceString(uniqueID, "${uniqueId}"))
					normalizers = append(normalizers, h.ReplaceString(project.ProjectID, "${projectId}"))

					for uuid := range uuids {
						normalizers = append(normalizers, h.ReplaceString(uuid, "${uuid}"))
						normalizers = append(normalizers, h.ReplaceString(strings.ReplaceAll(uuid, "-", ""), "${uuid-no-hyphens}"))
					}

					// Magic composer.googleapis.com "short code"
					{
						for uuid := range uuids {
							envName := fmt.Sprintf("%s-%.15s-%.8s", "us-central1", "minimal"+uniqueID, uuid)
							normalizers = append(normalizers, h.ReplaceString(envName, "${composer-system-name}"))
						}
					}

					h.CompareGoldenFile(expectedPath, got, normalizers...)
				}
			})
		}
	})

	// Do a cleanup while we can still handle the error.
	t.Logf("shutting down manager")
	cancel()
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}

func buildExportURI(t *testing.T, u *unstructured.Unstructured, project testgcp.GCPProject) string {
	// Some hints here: https://cloud.google.com/asset-inventory/docs/resource-name-format

	projectID := project.ProjectID
	resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
	if resourceID == "" {
		resourceID = u.GetName()
	}

	gvk := u.GroupVersionKind()
	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
		return "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + resourceID

	case schema.GroupKind{Group: "composer.cnrm.cloud.google.com", Kind: "ComposerEnvironment"}:
		location, _, _ := unstructured.NestedString(u.Object, "spec", "region")
		if location == "" {
			t.Fatalf("cannot determine spec.region for %v", u)
		}
		return "//composer.googleapis.com/projects/" + projectID + "/locations/" + location + "/environments/" + resourceID

	default:
		return ""
	}
}
