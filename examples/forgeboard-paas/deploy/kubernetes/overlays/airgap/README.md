# Air-Gapped Overlay

This overlay models a high-security customer environment.

Assumptions:

- Images are mirrored into `registry.airgap.local`.
- LLM access is provided by a customer internal gateway or local model endpoint.
- Vulnerability feeds, SBOMs, and policy bundles are imported through an offline update process.
- No public model provider is allowed by default.

Render:

```bash
kubectl kustomize deploy/kubernetes/overlays/airgap
```
