# Architecture

## Request Path

1. Customer opens the accessible frontend.
2. Customer uploads a transaction-policy bundle.
3. Middleware stores bundle metadata, creates ledger evidence, and publishes a `bundle.submitted` event.
4. Vetting workers inspect dependencies, schema compatibility, resource estimates, and security findings.
5. Delivery decides whether the bundle may be promoted.
6. Integration monitors live traffic and eval streams after promotion.
7. Debt diagnoses failures and scores operational interventions.
8. SecFlow can add optional security findings from third-party scanners.

## Runtime Topics

| Topic | Producer | Consumer |
|---|---|---|
| `forgeboard.bundle.submitted` | API | vetting workers, Urlog/Integration |
| `forgeboard.bundle.vetted` | vetting workers | API, Delivery |
| `forgeboard.run.events` | middleware/runtime | Integration, Debt |
| `forgeboard.ops.actions` | autonomous controller | Debt audit, operators |
| `forgeboard.secflow.findings` | SecFlow adapters | Delivery, Debt |
| `forgeboard.transactions.submitted` | middleware | vetting workers, Integration |
| `forgeboard.transactions.settled` | middleware | audit exports, Debt |

## Lifecycle

```text
upload -> static vetting -> dependency check -> schema check -> eval sample ->
deploy gate -> canary -> live monitoring -> incident diagnosis -> action scoring ->
safe automation or human approval
```

## Regulated Operating Model

ForgeBoard is aimed at EU fintech, medical, and green-tech buyers. The middleware behaves like a transaction system: it validates submitted operations, computes a deterministic risk score, records ledger-style evidence, and emits operational events. The sample uses a Bitcoin-like append-only evidence hash, but the settlement rail is abstract so a future implementation could use ordinary database ledgers, permissioned chains, or Bitcoin/Lightning-style proof anchors where appropriate.

The platform must be deployable with no public internet access. Air-gapped operators provide their own registry mirror, trusted base images, vulnerability feeds, model endpoints, and signing keys.
