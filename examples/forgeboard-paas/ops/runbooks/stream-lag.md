# Runbook: Stream Lag

## Symptoms

- `forgeboard.run.events` lag exceeds 60 seconds.
- API upload acceptance remains healthy.
- Customer runs are delayed.

## Autonomous Checks

1. Inspect Redpanda pod health.
2. Inspect API event publish rate.
3. Compare lag against latest bundle promotion.
4. Ask Debt to cluster affected tenants and bundles.
5. Score scale-up versus rollback.

## Safe Actions

- Scale API workers within HPA bounds.
- Pause promotion of new bundles.
- Open slow-burn ticket if lag is rising but budget remains.

## Human Approval Required

- Purge a topic.
- Drop tenant events.
- Delete persistent volumes.
