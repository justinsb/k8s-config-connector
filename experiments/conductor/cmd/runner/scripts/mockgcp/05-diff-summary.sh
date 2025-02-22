I am creating mock implementations for some GCP APIs.  I have captured the HTTP logs from requests and responses to the real APIs.  I have also captured logs from requests and responses to my mock.  I would like my mock to match the real results, but currently I am getting some diffs.  I have attached the diff.  Please give me a summarized list of the fields that I need to change so that there is no diff.


"""
Install an additional SDK for JSON schema support Google AI Python SDK

$ pip install google.ai.generativelanguage
"""

import os
import time
import google.generativeai as genai
from google.ai.generativelanguage_v1beta.types import content

genai.configure(api_key=os.environ["GEMINI_API_KEY"])

def upload_to_gemini(path, mime_type=None):
  """Uploads the given file to Gemini.

  See https://ai.google.dev/gemini-api/docs/prompting_with_media
  """
  file = genai.upload_file(path, mime_type=mime_type)
  print(f"Uploaded file '{file.display_name}' as: {file.uri}")
  return file

def wait_for_files_active(files):
  """Waits for the given files to be active.

  Some files uploaded to the Gemini API need to be processed before they can be
  used as prompt inputs. The status can be seen by querying the file's "state"
  field.

  This implementation uses a simple blocking polling loop. Production code
  should probably employ a more sophisticated approach.
  """
  print("Waiting for file processing...")
  for name in (file.name for file in files):
    file = genai.get_file(name)
    while file.state.name == "PROCESSING":
      print(".", end="", flush=True)
      time.sleep(10)
      file = genai.get_file(name)
    if file.state.name != "ACTIVE":
      raise Exception(f"File {file.name} failed to process")
  print("...all files ready")
  print()

# Create the model
generation_config = {
  "temperature": 1,
  "top_p": 0.95,
  "top_k": 40,
  "max_output_tokens": 8192,
  "response_schema": content.Schema(
    type = content.Type.OBJECT,
    enum = [],
    required = ["operations"],
    properties = {
      "operations": content.Schema(
        type = content.Type.ARRAY,
        items = content.Schema(
          type = content.Type.OBJECT,
          enum = [],
          required = ["operation", "fields"],
          properties = {
            "operation": content.Schema(
              type = content.Type.STRING,
            ),
            "fields": content.Schema(
              type = content.Type.ARRAY,
              items = content.Schema(
                type = content.Type.OBJECT,
                enum = [],
                required = ["field"],
                properties = {
                  "field": content.Schema(
                    type = content.Type.STRING,
                  ),
                  "current_value": content.Schema(
                    type = content.Type.STRING,
                  ),
                  "new_value": content.Schema(
                    type = content.Type.STRING,
                  ),
                },
              ),
            ),
          },
        ),
      ),
    },
  ),
  "response_mime_type": "application/json",
}

model = genai.GenerativeModel(
  model_name="gemini-2.0-flash",
  generation_config=generation_config,
)

# TODO Make these files available on the local file system
# You may need to update the file paths
files = [
  upload_to_gemini("diff.txt", mime_type="text/plain"),
]

# Some files have a processing delay. Wait for them to be ready.
wait_for_files_active(files)

chat_session = model.start_chat(
  history=[
    {
      "role": "user",
      "parts": [
        files[0],
        "I am creating mock implementations for some GCP APIs.  I have captured the HTTP logs from requests and responses to the real APIs.  I have also captured logs from requests and responses to my mock.  I would like my mock to match the real results, but currently I am getting some diffs.  I have attached the diff.  Please give me a summarized list of the fields that I need to change so that there is no diff.\n",
      ],
    },
    {
      "role": "model",
      "parts": [
        "```json\n{\n  \"operations\": [\n    {\n      \"fields\": [\n        {\n          \"field\": \"POST request URL\",\n          \"current_value\": \"requestId=76d9568f7a6d4ab6910d187ab905a7d9\",\n          \"new_value\": \"requestId=1969f7a04c504f3d87a06d4c9ba0e637\"\n        },\n        {\n          \"field\": \"Content-Length\",\n          \"current_value\": \"335\",\n          \"new_value\": \"315\"\n        },\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Create cluster with 2 workers\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"CREATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.warnings\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Create Cluster\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Create cluster with 2 workers\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"CREATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.statusHistory\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.warnings\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"response.status.state\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"RUNNING\\\"\"\n        },\n        {\n          \"field\": \"response.status.stateStartTime\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"2025-02-21T16:37:17.993843483Z\\\"\"\n        },\n        {\n          \"field\": \"response.statusHistory\",\n          \"current_value\": \"None\",\n          \"new_value\": \"[...]\"\n        }\n      ],\n      \"operation\": \"Get Operation for Create Cluster\"\n    },\n    {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Get Cluster\"\n    },\n   {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.workerConfig.numInstances\",\n          \"current_value\": \"2\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"status.state\",\n          \"current_value\": \"None\",\n          \"new_value\": \"\\\"RUNNING\\\"\"\n        }\n      ],\n      \"operation\": \"Get Cluster After Create\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"PATCH request URL\",\n          \"current_value\": \"requestId=cdd55cd75ab24f0e8344cb0fc7fcdf21\",\n          \"new_value\": \"requestId=5df6ba079e75463c8bfcb48c38acc41b\"\n        },\n        {\n          \"field\": \"Content-Length\",\n          \"current_value\": \"128\",\n          \"new_value\": \"118\"\n        },\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Add 1 workers.\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"UPDATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Update Cluster\"\n    },\n    {\n      \"fields\": [\n        {\n          \"field\": \"metadata.@type\",\n          \"current_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\",\n          \"new_value\": \"type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata\"\n        },\n        {\n          \"field\": \"metadata.clusterUuid\",\n          \"current_value\": \"3eaf1340-e4f2-487e-88bd-f76e537256a7\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.description\",\n          \"current_value\": \"\\\"Add 1 workers.\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.operationType\",\n          \"current_value\": \"\\\"UPDATE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.status\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"metadata.statusHistory\",\n          \"current_value\": \"[...]\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.config.workerConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"response.labels\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        }\n      ],\n      \"operation\": \"Get Operation for Update Cluster\"\n    },\n        {\n      \"fields\": [\n         {\n          \"field\": \"clusterUuid\",\n          \"current_value\": \"\\\"3eaf1340-e4f2-487e-88bd-f76e537256a7\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.configBucket\",\n          \"current_value\": \"\\\"dataproc-staging-us-central1-${projectNumber}-ch70stme\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.endpointConfig\",\n          \"current_value\": \"{}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.internalIpOnly\",\n          \"current_value\": \"true\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.networkUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.shieldedInstanceConfig\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.gceClusterConfig.zoneUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new_value\": \"None\"\n        },\n         {\n          \"field\": \"config.masterConfig.numInstances\",\n          \"current_value\": \"1\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.masterConfig.preemptibility\",\n          \"current_value\": \"\\\"NON_PREEMPTIBLE\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.imageVersion\",\n          \"current_value\": \"\\\"2.2.47-debian12\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.softwareConfig.properties\",\n          \"current_value\": \"{...}\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.tempBucket\",\n          \"current_value\": \"\\\"dataproc-temp-us-central1-${projectNumber}-30wedmya\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskSizeGb\",\n          \"current_value\": \"1000\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.diskConfig.bootDiskType\",\n          \"current_value\": \"\\\"pd-standard\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.imageUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.machineTypeUri\",\n          \"current_value\": \"\\\"https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-c/machineTypes/n2-standard-4\\\"\",\n          \"new_value\": \"None\"\n        },\n        {\n          \"field\": \"config.workerConfig.minCpuPlatform\",\n          \"current_value\": \"\\\"AUTOMATIC\\\"\",\n          \"new\n```",
      ],
    },
  ]
)

response = chat_session.send_message("INSERT_INPUT_HERE")

print(response.text)



I am implementing a mock for GCP, and I need to make the following changes so that it matches the behaviour of real GCP.

For the "Create Cluster" operation:

metadata.operationType is currently empty, it should be "CREATE"

metadata.status is currently empty, it should be  {     "innerState": "PENDING",      "state": "PENDING",      "stateStartTime": <current_time> }

operations": [
    {
      "fields": [
        {
          "field": "POST request URL",
          "current_value": "requestId=76d9568f7a6d4ab6910d187ab905a7d9",
          "new_value": "requestId=1969f7a04c504f3d87a06d4c9ba0e637"
        },
        {
          "field": "Content-Length",
          "current_value": "335",
          "new_value": "315"
        },
        {
          "field": "metadata.@type",
          "current_value": "type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata",
          "new_value": "type.googleapis.com/google.cloud.dataproc.v1.ClusterOperationMetadata"
        },
        {
          "field": "metadata.clusterUuid",
          "current_value": "3eaf1340-e4f2-487e-88bd-f76e537256a7",
          "new_value": "None"
        },
        {
          "field": "metadata.description",
          "current_value": "\"Create cluster with 2 workers\"",
          "new_value": "None"
        },
        {
          "field": "metadata.operationType",
          "current_value": "\"CREATE\"",
          "new_value": "None"
        },
        {
          "field": "metadata.status",
          "current_value": "{...}",
          "new_value": "None"
        },
        {
          "field": "metadata.warnings",
          "current_value": "[...]",
          "new_value": "None"
        }
      ],
      "operation": "Create Cluster"
    },