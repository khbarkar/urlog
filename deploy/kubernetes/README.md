# Urlog Kubernetes Install

Urlog runs on Kubernetes by default through Kustomize.

The Urlog control plane is stateless:

- `urlog-api`: API and static UI.
- `urlog-worker`: Integration/Delivery/Debt/Eye worker execution.
- `urlog-operator`: Kubernetes controller for future CRDs and cluster-local reconciliation.

State and dependencies belong outside these pods:

- Redpanda for event streams.
- ClickHouse for compact analytical metadata/evidence records.
- Object storage for reports, bundles, PDFs, transcripts, sampled payloads, and large artifacts.
- Optional OpenSearch for report/evidence search.
- Optional Loki or customer log systems for raw log collection.
- A secret backend for all sensitive values.

Urlog services should be `Deployment` resources. Use `StatefulSet` only for self-hosted dependencies such as Redpanda, ClickHouse, OpenSearch, or object storage.

Urlog is not a raw log collector. It stores operational metadata and references; high-volume application logs should stay in a dedicated logging system.

## Install

```bash
kubectl apply -k deploy/kubernetes/overlays/k3s
```

For Docker Desktop:

```bash
kubectl apply -k deploy/kubernetes/overlays/docker-desktop
```

For air-gapped installs, mirror images first and then apply:

```bash
kubectl apply -k deploy/kubernetes/overlays/airgap
```

## Required Bootstrap Config

Before the first useful run, provide non-secret bootstrap config:

- LLM provider references.
- Repository access references.
- Secret backend adapter configuration.
- Infrastructure scope.
- Integration system mode: `observe_only`, `existing`, or `install`.

See [`../../bootstrap/`](../../bootstrap/).

## Secrets

These manifests do not create default passwords or static credentials.

Use one of:

- SOPS/age encrypted manifests.
- External Secrets Operator connected to a customer secret manager.
- Cloud secret manager access through workload identity or MCP.
- A manually created Kubernetes Secret for development only.

Config must contain references only. Secret values must not appear in ConfigMaps, logs, traces, reports, prompts, or readiness packets.

## Images

The default image names are placeholders until the Go binaries are implemented:

- `ghcr.io/khbarkar/urlog-api:dev`
- `ghcr.io/khbarkar/urlog-worker:dev`
- `ghcr.io/khbarkar/urlog-operator:dev`
