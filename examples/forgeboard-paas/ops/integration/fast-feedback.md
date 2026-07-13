# Fast Feedback Loop

Integration should give a developer/operator an answer quickly, before Delivery spends time on deeper gates.

## Feedback Order

1. Source snapshot valid.
2. Dependency manifests found.
3. Docker build contexts found.
4. Kubernetes manifests render.
5. Fast tests pass.
6. Eye dependency/security monitor returns no blocking findings.
7. Readiness score emitted to Delivery.

## Non-Goals

- No general code scanning.
- No source-code style review.
- No prompt playground.
- No production release decision. Delivery owns that.
