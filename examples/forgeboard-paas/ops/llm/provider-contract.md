# Pluggable LLM Provider Contract

LLMs are providers behind an interface. No module should hard-code a provider.

## Required Capabilities

| Capability | Used by | Notes |
|---|---|---|
| `summarize_evidence` | Debt | Summarize logs, spans, and eval evidence |
| `propose_action` | Debt | Suggest bounded operational actions from catalog only |
| `classify_upload_risk` | Delivery | Assist with bundle validation and release readiness |
| `judge_task_completion` | Integration | Sampled eval judge, stamped with `evaluator_version` |
| `security_reasoning` | Eye | Optional, only if security plugin is enabled |

## Provider Rules

- Provider must be swappable through configuration.
- Provider output must be treated as evidence, not authority.
- Proposed actions must reference `ops/autonomy/action-catalog.yaml`.
- Destructive actions require human approval regardless of provider confidence.
- Every model call records provider, model, prompt bundle version, and evaluator version.
- Air-gapped deployments must use customer-provided model endpoints or deterministic local stubs.
- Compliance profiles may block public model providers even when credentials are present.
