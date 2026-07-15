# Integration Contracts

These contracts define how Integration acts autonomously.

They are written for two readers:

- **Humans** deciding what the system is allowed to do.
- **Robots** selecting, scoring, executing, verifying, and auditing actions.

## Operating Rule

The LLM may plan and select actions, but it must express them as an `intent`.

```text
intent -> action exists -> environment allowed -> evidence present ->
risk score calculated -> pre-approval matched -> execute -> verify -> audit
```

If an action is not in `action-catalog.yaml`, it cannot be executed.

If an environment is not in `environment-catalog.yaml`, it cannot be targeted.

If no pre-approval policy matches, the action is held, escalated, or denied according to policy.

## Files

| File | Human purpose | Robot purpose |
|---|---|---|
| `action-catalog.yaml` | Shows what Integration may do | Validates `intent.action_id` and executor requirements |
| `environment-catalog.yaml` | Shows where actions may happen | Prevents accidental production or wrong-cluster actions |
| `pre-approval-policies.yaml` | Shows what runs without a human | Resolves `approval_resolution` |
| `risk-scoring.yaml` | Shows how risk is calculated | Computes deterministic risk score |
| `intent.schema.json` | Shows what the LLM must emit | Validates LLM-planned action requests |
| `execution-record.schema.json` | Shows what must be audited | Validates immutable execution records |

## Approval Resolutions

`pre_approved`: action may run automatically because policy matched.

`human_required`: action is valid, but requires human approval.

`denied`: action must not run.

`escalated`: action needs a higher authority, incident commander, or break-glass path.

## Integration Defaults

Most Integration actions should be pre-approved in CI or integration environments:

- Fetch source snapshot.
- Inspect manifests.
- Build containers.
- Run tests.
- Render Kubernetes manifests.
- Create ephemeral namespaces.
- Run Eye dependency/deprecation plugins.
- Emit readiness score.
- Clean up ephemeral resources.

Production-impacting or destructive actions are not pre-approved by default.

## Non-Goals

Integration does not own:

- Production release policy.
- Incident diagnosis.
- Long-term SLO math.
- Security scanner implementation.
- General code scanning as a product category.
- ESXi provisioning.

## Robot Contract

Robots must:

1. Read all contracts before selecting actions.
2. Produce an intent that validates against `intent.schema.json`.
3. Never invent an action id.
4. Never target an unknown environment.
5. Provide required evidence references.
6. Use deterministic scoring from `risk-scoring.yaml`.
7. Execute only after policy resolves to `pre_approved` or a valid human approval token exists.
8. Write an execution record that validates against `execution-record.schema.json`.
