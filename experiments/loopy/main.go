package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type Result struct {
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Operation string  `json:"operation"`
	Fields    []Field `json:"fields"`
}

type Field struct {
	Field        string `json:"field"`
	RemovedValue string `json:"removed_value,omitempty"`
	AddedValue   string `json:"added_value,omitempty"`
}

func BuildSchemaFor(t reflect.Type) *llm.Schema {
	out := &llm.Schema{}

	switch t.Kind() {
	case reflect.String:
		out.Type = llm.TypeString
	case reflect.Struct:
		out.Type = llm.TypeObject
		out.Properties = make(map[string]*llm.Schema)
		numFields := t.NumField()
		required := []string{}
		for i := 0; i < numFields; i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				continue
			}
			if strings.HasSuffix(jsonTag, ",omitempty") {
				jsonTag = strings.TrimSuffix(jsonTag, ",omitempty")
			} else {
				required = append(required, jsonTag)
			}

			fieldType := field.Type

			fieldSchema := BuildSchemaFor(fieldType)
			out.Properties[jsonTag] = fieldSchema
		}

		if len(required) != 0 {
			out.Required = required
		}
	case reflect.Slice:
		out.Type = llm.TypeArray
		out.Items = BuildSchemaFor(t.Elem())
	default:
		klog.Fatalf("unhandled kind %v", t.Kind())
	}

	return out
}

func run(ctx context.Context) error {
	schema := BuildSchemaFor(reflect.TypeOf(Result{}))

	schemaJSON, _ := json.Marshal(schema)
	fmt.Printf("schema is %v", string(schemaJSON))

	llmClient, err := llm.BuildVertexAIClient(ctx)
	if err != nil {
		return err
	}
	defer llmClient.Close()

	// llmClient.WithModel("gemini-2.0-flash")

	systemPrompt := ""

	chat := llmClient.StartChat(systemPrompt)

	if err := chat.SetResponseSchema(schema); err != nil {
		return err
	}

	diff, err := llm.AttachFile("/home/justinsb/diff.txt")
	if err != nil {
		return err
	}

	response, err := chat.SendMessage(ctx, diff, `
I am creating mock implementations for some GCP APIs.
I have captured the HTTP logs from requests and responses to the real APIs.
I have also captured logs from requests and responses to my mock.
I would like my mock to match the real results, but currently I am getting some diffs.
I have attached the diff.
Please give me a summarized list of the fields that I need to change so that there is no diff.  Only give about 10 differences, so you don't exceed the response size limit.
Put the value that is deleted in the diff into removed_value.  Put the value that is added in the diff into added_value.
You can ignore any HTTP Content-Length or Cache-Control header changes.

Only return leaf fields, not the parent fields, the corresponding values should never be a JSON object.

For example, given the following diff:

<diff>
 GET foo/bar/baz

 200 OK
 Content-Type: application/json

 {
   "metadata": {
     "@type": "type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata",
-    "clusterName": "test123",
-    "clusterUuid": "3eaf1340-e4f2-487e-88bd-f76e537256a7",
-    "description": "Add 1 workers.",
-    "operationType": "UPDATE",
-    "status": {
-      "innerState": "PENDING",
-      "state": "PENDING",
-      "stateStartTime": "2025-02-21T15:08:27.682465Z"
-    }
+    "clusterName": "test456"
   },
   "name": "projects/${projectId}/regions/us-central1/operations/${operationID}"
 }
</diff>

You might reply with the following:

<response>
{ "operations": { "operation": "GET foo/bar/bz", "fields": [
  { "field": "metadata.clusterName", "removed_value": "test123", "added_value": "test456" },
  { "field": "metadata.clusterUuid", "removed_value": "3eaf1340-e4f2-487e-88bd-f76e537256a7" },
  { "field": "metadata.description", "removed_value": "Add 1 workers." },
  { "field": "metadata.status.innerState", "removed_value": "PENDING" },
  { "field": "metadata.status.state", "removed_value": "PENDING" },
  { "field": "metadata.status.stateStartTime", "removed_value": "2025-02-21T15:08:27.682465Z" },
]}}
</response>

`)
	if err != nil {
		return fmt.Errorf("error from llm: %w", err)
	}

	var responses []*Result
	for _, candidate := range response.Candidates() {
		for _, part := range candidate.Parts() {
			s, ok := part.AsText()
			if ok {
				result := &Result{}
				// decoder := json.NewDecoder(strings.NewReader(s))
				// // if err := decoder.Decode(result); err != nil {
				// // 	klog.Warningf("error during decode: %v", err)
				// // }
				if err := json.Unmarshal([]byte(s), result); err != nil {
					klog.Infof("response is %v", s)
					return fmt.Errorf("parsing json: %w", err)
				}
				responses = append(responses, result)

			} else {
				return fmt.Errorf("unexpected non-text response")
			}
		}
	}
	if len(responses) == 0 {
		return fmt.Errorf("did not get any responses from LLM")
	}
	if len(responses) > 1 {
		return fmt.Errorf("got multiple responses from LLM")
	}

	// b, err := json.MarshalIndent(responses[0], "", "  ")
	// if err != nil {
	// 	return fmt.Errorf("rendering json: %w", err)
	// }
	// fmt.Fprintf(os.Stdout, "%s\n", string(b))
	w := os.Stdout
	fmt.Fprintf(w, "I am implementing a mock for GCP, and I need to make the following changes so that it matches the behaviour of real GCP.  Please fix the following problems: \n")
	fmt.Fprintf(w, "\n")
	for _, operation := range responses[0].Operations {

		fmt.Fprintf(w, "# For the operation %q\n", operation.Operation)

		for _, field := range operation.Fields {
			fmt.Fprintf(w, "  * Field %q should be changed from %q to %q\n", field.Field, field.AddedValue, field.RemovedValue)
		}

	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Use the EditFile function to update cluster.go\n")

	return nil
}

// }
// def upload_to_gemini(path, mime_type=None):
//   """Uploads the given file to Gemini.

//   See https://ai.google.dev/gemini-api/docs/prompting_with_media
//   """
//   file = genai.upload_file(path, mime_type=mime_type)
//   print(f"Uploaded file '{file.display_name}' as: {file.uri}")
//   return file

// def wait_for_files_active(files):
//   """Waits for the given files to be active.

//   Some files uploaded to the Gemini API need to be processed before they can be
//   used as prompt inputs. The status can be seen by querying the file's "state"
//   field.

//   This implementation uses a simple blocking polling loop. Production code
//   should probably employ a more sophisticated approach.
//   """
//   print("Waiting for file processing...")
//   for name in (file.name for file in files):
//     file = genai.get_file(name)
//     while file.state.name == "PROCESSING":
//       print(".", end="", flush=True)
//       time.sleep(10)
//       file = genai.get_file(name)
//     if file.state.name != "ACTIVE":
//       raise Exception(f"File {file.name} failed to process")
//   print("...all files ready")
//   print()

// # Create the model
// generation_config = {
//   "temperature": 1,
//   "top_p": 0.95,
//   "top_k": 40,
//   "max_output_tokens": 8192,
//   "response_schema": content.Schema(
//     type = content.Type.OBJECT,
//     enum = [],
//     required = ["operations"],
//     properties = {
//       "operations": content.Schema(
//         type = content.Type.ARRAY,
//         items = content.Schema(
//           type = content.Type.OBJECT,
//           enum = [],
//           required = ["operation", "fields"],
//           properties = {
//             "operation": content.Schema(
//               type = content.Type.STRING,
//             ),
//             "fields": content.Schema(
//               type = content.Type.ARRAY,
//               items = content.Schema(
//                 type = content.Type.OBJECT,
//                 enum = [],
//                 required = ["field"],
//                 properties = {
//                   "field": content.Schema(
//                     type = content.Type.STRING,
//                   ),
//                   "current_value": content.Schema(
//                     type = content.Type.STRING,
//                   ),
//                   "new_value": content.Schema(
//                     type = content.Type.STRING,
//                   ),
//                 },
//               ),
//             ),
//           },
//         ),
//       ),
//     },
//   ),
//   "response_mime_type": "application/json",
// }

// model = genai.GenerativeModel(
//   model_name="gemini-2.0-flash",
//   generation_config=generation_config,
// )

// # TODO Make these files available on the local file system
// # You may need to update the file paths
// files = [
//   upload_to_gemini("diff.txt", mime_type="text/plain"),
// ]

// # Some files have a processing delay. Wait for them to be ready.
// wait_for_files_active(files)

// chat_session = model.start_chat(
//   history=[
//     {
//       "role": "user",
//       "parts": [
//         files[0],
//         "I am creating mock implementations for some GCP APIs.  I have captured the HTTP logs from requests and responses to the real APIs.  I have also captured logs from requests and responses to my mock.  I would like my mock to match the real results, but currently I am getting some diffs.  I have attached the diff.  Please give me a summarized list of the fields that I need to change so that there is no diff.\n",
//       ],
//     },
//     {
//       "role": "model",
//       "parts": [
//         "```json\n{\n  \"operations\": [\n    {\n      \"fields\": [\n        {\n          \"field\": \"POST request URL\",\n          \"current_value\": \"requestId=76d9568f7a6d4ab6910d187ab905a7d9\",\n          \"new_value\": \"requestId=1969f7a04c504f3d87a06d4c9ba0e637\"\n        },\n        {\n          \"field\": \"Content-Length\",\n          \"current_value\": \"335\",\n          \"new_value\": \"315\"\n        },\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Create cluster with 2 workers\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"CREATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.warnings\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Create Cluster\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Create cluster with 2 workers\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"CREATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.statusHistory\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.warnings\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"response.status.state\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"RUNNING\\\"\"\n        },\n        {\n          \"field\": \"response.status.stateStartTime\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"2025-02-21T16:37:17.993843483Z\\\"\"\n        },\n        {\n          \"field\": \"response.statusHistory\",\n          \"current_value\": \"None\",\n          \"new_value\": \"[...]\"\n        }\n      ],\n      \"operation\": \"Get Operation for Create Cluster\"\n    },\n    {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Get Cluster\"\n    },\n   {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.workerConfig.numInstances\",\n          \"current_value\": \"2\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"status.state\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"RUNNING\\\"\"\n        }\n      ],\n      \"operation\": \"Get Cluster After Create\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"PATCH request URL\",\n          \"current_value\": \"requestId=cdd55cd75ab24f0e8344cb0fc7fcdf21\",\n          \"new_value\": \"requestId=5df6ba079e75463c8bfcb48c38acc41b\"\n        },\n        {\n          \"field\": \"Content-Length\",\n          \"current_value\": \"128\",\n          \"new_value\": \"118\"\n        },\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Add 1 workers.\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"UPDATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Update Cluster\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Add 1 workers.\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"UPDATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.statusHistory\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Get Operation for Update Cluster\"\n    },\n        {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new\n```",
//       ],
//     },
//   ]
// )

// response = chat_session.send_message("INSERT_INPUT_HERE")

// print(response.text)

// 	return nil

// }
