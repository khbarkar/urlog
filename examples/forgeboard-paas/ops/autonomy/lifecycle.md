# Autonomous Lifecycle

The autonomous system handles the full lifecycle, but actions are policy-gated.

## Flow

```text
Observe -> Classify -> LoadComplianceProfile -> Diagnose -> ProposeAction ->
ScoreRisk -> CheckPolicy -> ExecuteOrRequestApproval -> Verify -> Audit
```

## Action Classes

| Class | Examples | Approval |
|---|---|---|
| Read-only | inspect pods, query metrics, fetch logs | automatic |
| Reversible safe | scale deployment within bounds, restart one failed pod | automatic if score passes |
| Reversible risky | rollback canary, downgrade one minor version | human if tenant impact is high |
| Destructive | delete PVC, purge topic, rotate tenant secrets | human required |

## Module Ownership

- Integration observes and streams operational state.
- Delivery gates deploy readiness and release safety.
- Debt diagnoses and scores operational interventions.
- Eye adds optional security findings.

## Compliance-Aware Operation

Before an action is executed, the controller loads the active compliance profile from `compliance/profiles/`. The strictest active control wins:

- Missing accessibility evidence can block customer-facing releases.
- Missing model-provider traceability can block AI-assisted actions.
- Missing incident or audit evidence can block production promotion.
- Air-gapped deployments reject public model endpoints and public registries.
- Destructive actions require human approval in every profile.
