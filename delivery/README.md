# Delivery

Norn name: **Skuld**  
Small name: **Fate**

Delivery owns whether a change is fit to ship. It consumes Integration scores, eval scores, CI/build/test evidence, security requirements, and operational SLO state, then decides whether a release may proceed.

## Owns

- Quality SLOs and error budgets.
- Multi-window multi-burn-rate alerting.
- Eval-gated releases.
- Deploy readiness gates.
- Locking gates for security, compliance, accessibility, and missing evidence.
- Release verdicts consumed by autonomous operations.

## Does Not Own

- Build execution or source repository integration. That is Integration.
- Security scanner internals. Eye provides optional findings.
- Forensic diagnosis and intervention risk scoring. That is Debt.

## ForgeBoard Role

Delivery uses Integration's readiness score for `examples/forgeboard-paas` to decide whether transaction-policy bundles, frontend images, middleware images, and Kubernetes manifests can advance to canary or production.
