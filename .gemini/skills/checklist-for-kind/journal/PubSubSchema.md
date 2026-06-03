# Migration Journal: PubSubSchema

* **Kind**: `PubSubSchema`
* **Current Step**: Step 1: Direct API Types (PR approved, waiting for merge)

## Progress Tracker

| Step | Description | Issue | PR | Status | Started | Completed |
|---|---|---|---|---|---|---|
| 1 | Direct API Types & generate.sh | [#9059](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9059) | [#9060](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9060) | `PR Created` | 2026-06-03 | - |
| 2 | Identity and Reference Types | - | - | `Not Started` | - | - |
| 3 | Create a Round-Trip KRM Fuzzer | - | - | `Not Started` | - | - |

## Notes / Logs
* **2026-06-03**: Initialized migration journal for `PubSubSchema`. Checked `apis/pubsub/v1beta1/` and verified that `PubSubSchema` is missing direct types, identity, and reference definitions.
* **2026-06-03**: Opened issue #9059 for `codebot-robot` to implement direct KRM types and configure `generate.sh` under `apis/pubsub/v1beta1/` (Step 1).
* **2026-06-03**: `codebot-robot` successfully created Pull Request [#9060](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9060) for Step 1. The PR has been approved and is pending merge.
