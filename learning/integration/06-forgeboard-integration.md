# 06 - ForgeBoard Integration

Goal: use `examples/forgeboard-paas` as the Integration target.

Integration connects to the ForgeBoard repository, builds the containers, runs tests, imports Eye plugin findings, and emits a readiness score for Delivery.

## Scope

Integration handles:

- Centralized version-control metadata.
- Automated builds.
- Automated testing.
- Fast feedback loops.
- Security and dependency findings from Eye plugins.
- Locking gate input.
- Readiness scoring for Delivery.

Integration does not handle:

- General source-code scanning as a product category.
- Release policy.
- Incident forensics.
- Destructive operational actions.

## Done

`examples/forgeboard-paas/ops/integration/readiness-score.yaml` can be produced from repo, build, test, and Eye evidence.
