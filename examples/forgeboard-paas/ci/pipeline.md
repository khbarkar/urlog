# CI and Vetting Pipeline

This is the initial Delivery-facing pipeline shape.

## Pull Request / Bundle Checks

1. Build `forgeboard-frontend`.
2. Build `forgeboard-api`.
3. Generate SBOM.
4. Scan dependencies.
5. Validate Kubernetes manifests with `kubectl apply --dry-run=server`.
6. Run schema compatibility checks.
7. Run bundle resource estimate.
8. Run sampled eval fixtures.
9. Optionally run SecFlow security checks.
10. Emit deploy-gate evidence for Delivery.

## Gate Output

```json
{
  "gate": "bundle_upload",
  "decision": "hold",
  "reasons": ["dependency_vulnerability", "stream_fanout_watch"],
  "evaluator_version": "delivery-gate-0.1.0"
}
```
