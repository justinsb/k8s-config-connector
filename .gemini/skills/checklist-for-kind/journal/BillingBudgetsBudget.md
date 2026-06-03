# Migration Journal: BillingBudgetsBudget

* **Kind**: `BillingBudgetsBudget`
* **Current Step**: Step 1: Direct API Types (Issue opened)

## Progress Tracker

| Step | Description | Issue | PR | Status | Started | Completed |
|---|---|---|---|---|---|---|
| 1 | Direct API Types & generate.sh | [#9013](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9013) | - | `Open` | 2026-06-02 | - |
| 2 | Identity and Reference Types | - | - | `Not Started` | - | - |
| 3 | Create a Round-Trip KRM Fuzzer | - | - | `Not Started` | - | - |

## Notes / Logs
* **2026-06-02**: Initialized migration journal for `BillingBudgetsBudget`. Checked `apis/` and verified that `apis/billingbudgets` is completely missing (neither direct types nor identity/reference definitions exist).
* **2026-06-02**: Opened issue #9013 for `codebot-robot` to implement direct KRM types and configure `generate.sh` under `apis/billingbudgets/v1beta1/` (Step 1).
