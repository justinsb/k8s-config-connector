# Migration Journal: LoggingLogExclusion

* **Kind**: `LoggingLogExclusion`
* **Current Step**: Step 2: Identity and Reference Types (Issue opened)

## Progress Tracker

| Step | Description | Issue | PR | Status | Started | Completed |
|---|---|---|---|---|---|---|
| 1 | Direct API Types & generate.sh | [#9012](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9012) | [#9014](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9014) | `Completed` | 2026-06-02 | 2026-06-03 |
| 2 | Identity and Reference Types | [#9063](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9063) | - | `Open` | 2026-06-03 | - |
| 3 | Create a Round-Trip KRM Fuzzer | - | - | `Not Started` | - | - |

## Notes / Logs
* **2026-06-02**: Initialized migration journal for `LoggingLogExclusion`. Checked `apis/logging/v1beta1/` and verified that `LoggingLogExclusion` is missing direct types, identity, and reference definitions.
* **2026-06-02**: Opened issue #9012 for `codebot-robot` to implement direct KRM types and configure `generate.sh` under `apis/logging/v1beta1/` (Step 1).
* **2026-06-03**: `codebot-robot` successfully created Pull Request [#9014](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9014) for Step 1. The PR has been approved and is pending merge.
* **2026-06-03**: Pull Request [#9014](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/9014) merged successfully. Marked Step 1 as `Completed`. Opened issue #9063 for `codebot-robot` to move the resource to the modern identity and reference types pattern (Step 2).
