# Migration Journal: LoggingLogBucket

* **Kind**: `LoggingLogBucket`
* **Current Step**: Step 1: Direct API Types (PR approved, waiting for merge)

## Progress Tracker

| Step | Description | Issue | PR | Status | Started | Completed |
|---|---|---|---|---|---|---|
| 1 | Direct API Types & generate.sh | [#8995](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8995) | [#8996](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8996) | `PR Created` | 2026-06-02 | - |
| 2 | Identity and Reference Types | - | - | `Completed` | - | 2026-06-02 |

## Notes / Logs
* **2026-06-02**: Initialized migration journal for `LoggingLogBucket`. Checked `apis/logging/v1beta1/` and confirmed that identity and reference types (`logbucket_identity.go` and `logbucket_reference.go`) already exist, so Step 2 is marked as `Completed`.
* **2026-06-02**: Checked `apis/logging/v1beta1/generate.sh` and verified `LoggingLogBucket` direct types are missing. Opened issue #8995 for `codebot-robot` to implement direct KRM types and configure the generation script using the existing resource workflow (Step 1).
* **2026-06-02**: `codebot-robot` successfully created Pull Request [#8996](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8996) for Step 1. The PR has been approved and is pending merge.
