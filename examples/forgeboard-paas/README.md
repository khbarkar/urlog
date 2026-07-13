# ForgeBoard PaaS

ForgeBoard is a fake 180-person EU-first B2B SaaS company used as an Urlog design-partner sandbox. It runs a multi-tenant operational finance platform where customers upload transaction-policy bundles that process event streams. Those uploads can break the platform through bad code, dependency drift, data-volume spikes, prompt-injection content, schema-incompatible events, or unsafe settlement logic.

The target customers are fintech, medical startups, and green-technology operators up to roughly 500 staff. The sample system is intentionally ordinary: two application containers plus a streaming backend on Kubernetes.

## System Shape

| Layer | Container / service | Purpose |
|---|---|---|
| Frontend | `forgeboard-frontend` | Accessible customer portal for uploading transaction policies, viewing runs, and managing tenants |
| Middleware | `forgeboard-api` | Transaction API, upload vetting, run orchestration, ledger evidence, and operational metadata |
| Streaming backend | Redpanda-compatible Kafka API | Event intake, run events, vetting results, and operational signals |

Initial target runtime:

- `k3s` for local cluster realism.
- Docker Desktop Kubernetes as a secondary local target.
- Air-gapped Kubernetes as a high-security target.
- Later cloud targets can be added as overlays without changing the core app manifests.

## Why This Can Break

Customers upload `bundle.zip` packages that contain transaction validation rules, operational workflows, and a manifest. A bad upload can:

- Pin vulnerable libraries.
- Emit events with incompatible schema.
- Create a runaway stream fan-out.
- Include prompt-injection content in customer-provided instructions.
- Increase CPU or memory pressure during validation.
- Trigger poor eval scores for task completion or faithfulness.
- Create incorrect payment, carbon-credit, subsidy, insurance, or medical billing outcomes.

This gives Urlog a realistic operational target: validate before deploy, monitor after deploy, score risk, gate releases, troubleshoot incidents, and require human approval for destructive actions.

## EU-First Controls

This sample is designed around EU-market expectations:

- **DORA-shaped operational resilience** for fintech and ICT-provider buyers.
- **AI Act logging and evidence** for high-risk workflows where AI assists operational decisions.
- **GDPR/data-minimisation posture**: payloads are references, not inlined into analytics rows.
- **Accessibility by default**: UI and generated operator reports target EN 301 549 / WCAG AA practices.
- **Air-gapped operation**: customers provide container registries, model endpoints, keys, and update bundles.
- **Pluggable LLMs**: no hard dependency on a hosted model provider; LLM access is supplied by the customer environment.

## Urlog Mapping

| Urlog module | In this sample |
|---|---|
| Integration | Collects OTLP signals, consumes stream events, runs tiered eval workers, and feeds operational state |
| Delivery | Gates uploads/releases using CI, dependency checks, quality SLOs, burn-rate alerts, and validation results |
| Debt | Diagnoses incidents, links sessions/spans/eval scores, scores intervention risk, and keeps audit evidence |
| Eye | Optional DevSecOps plugin layer for prompt-injection, jailbreak, and supply-chain security signals |

## Folder Layout

```text
apps/                    frontend and middleware container sources
platform/streaming/      streaming backend manifests
deploy/kubernetes/       base manifests and local overlays
ops/autonomy/            action catalog, lifecycle model, scoring rules
ops/llm/                 pluggable and air-gap capable LLM provider contract
ops/policies/            gates and approval policies
ops/runbooks/            operational runbooks for the autonomous system
ci/                      realistic CI/vetting pipeline sketch
compliance/              EU controls, accessibility, and air-gap requirements
docs/                    architecture notes
```

## Local Build Sketch

```bash
docker build -t forgeboard/frontend:dev apps/frontend
docker build -t forgeboard/api:dev apps/middleware
kubectl apply -k deploy/kubernetes/overlays/k3s
```

The manifests are static scaffolding. They are meant to define the target lifecycle surface before the real implementation is added.
