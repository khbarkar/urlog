<p align="center">
  <img src="docs/assets/logo.png" alt="Urlog" width="540">
</p>

<h1 align="center">Urlog</h1>

<p align="center">
  SLO engine and incident layer for production AI systems:<br>
  quality error budgets, burn-rate alerting, eval-gated releases, audit-grade records.
</p>

---

## Overview

Existing AI observability tools serve the dev loop (trace, debug, eval, iterate). Urlog serves the ops loop: it decides whether someone should be paged, whether a release should proceed, and which operational action is safe to take — and it can execute pre-approved actions autonomously. Every action is traceable, scored, gated, and auditable.

Design constraints:

- **OTLP-native, no proprietary SDK.** Ingests standard OTel GenAI semantic conventions (OpenLLMetry, OpenInference). Existing instrumentation works unmodified.
- **Session as unit of analysis.** Agent failures are multi-step causal chains; the data model tracks sessions with goal-level outcomes, not individual requests.
- **Metadata store, not a log collector.** ClickHouse holds narrow metadata rows; prompt/completion payloads and large artifacts stay in object storage behind `payload_ref`. See [`docs/storage-policy.md`](docs/storage-policy.md).
- **Versioned eval scores.** Eval results are a separate stream joined at query time, each stamped with `evaluator_version`, because judges drift like the systems they judge.
- **Tiered evaluation.** Classifiers on 100% of traffic, LLM-judge on stratified samples, human review on disagreement.
- **Multi-window multi-burn-rate alerting** (SRE Workbook ch. 5), not naive thresholds.
- **EU-first, self-hostable.** No hard dependency on US-only managed services.
- **Contracts before models.** The LLM is pluggable and may plan work, but the system acts only through machine-readable contracts (see below).

## Modules

Five modules, bound by one contract: the versioned protobuf schema in [`schema/`](schema/). Integration writes it, Delivery reads it to gate, Debt queries it forever, Eye reports on it, SecFlow annotates it. A change that couples two modules without going through the schema is a defect.

| Module | Role | Owns |
|---|---|---|
| **Integration** | Ingest | OTLP gRPC ingest, Redpanda consumers, tiered eval workers, live operational signal feed |
| **Delivery** | Deployment | Quality SLOs, error budgets, multi-window burn-rate alerting, eval-gated releases, CI readiness checks |
| **Debt** | Troubleshooting | Session forensics, incident lifecycle, hash-chained immutable audit log, AI Act Article 12 retention, action risk scoring |
| **Eye** | Reporting | Live reports, metadata/evidence indexes, PDF/doc generation, report upload connectors, optional search sinks |
| **SecFlow** | Security (optional) | Prompt-injection detection, garak evidence, dependency/image/SBOM findings, security gate signals |

SecFlow can be omitted without breaking the core loop.

## Contract-gated execution

Every autonomous action passes this pipeline:

```text
intent -> action exists -> environment allowed -> evidence present ->
risk score calculated -> pre-approval matched -> execute -> verify -> audit
```

Contracts for Integration live in [`modules/integration/contracts/`](modules/integration/contracts/):

| File | Defines |
|---|---|
| `action-catalog.yaml` | Actions the autonomous operator may choose |
| `environment-catalog.yaml` | Environments where actions may run |
| `pre-approval-policies.yaml` | Actions permitted without live human approval |
| `risk-scoring.yaml` | Final action risk calculation |
| `intent.schema.json` | Shape an LLM-planned action must emit |
| `execution-record.schema.json` | Audit record every action must leave behind |

Enforcement is deny-by-default: an action absent from the catalog cannot run, an unlisted environment cannot be targeted, and missing evidence holds or denies the action. Predefined destructive actions always require human approval.

First-time installation starts with the bootstrap contract in [`bootstrap/`](bootstrap/), which defines the initial trust handoff: LLM access, repository access, secret backend, and infrastructure access. Configuration carries secret *references* only — secret values never appear in config, logs, prompts, traces, reports, or readiness packets.

## Data flow

1. Spans arrive over OTLP gRPC and buffer through Redpanda.
2. Consumers land narrow metadata rows in ClickHouse; payloads go to object storage behind references.
3. Eval workers score sessions on the tiered schedule; scores stream separately, stamped with `evaluator_version`.
4. Delivery joins scores into SLIs, tracks error budgets, and gates releases.
5. Debt records incidents and the hash-chained audit trail; Eye generates reports; SecFlow attaches security findings.

## Quickstart (dev)

```bash
git clone <repo-url> && cd urlog
docker compose -f deploy/docker-compose.dev.yml up   # ClickHouse + Redpanda + otel-collector
```

Point any OTel GenAI-instrumented app at `localhost:4317`.

## Deployment

Default target is Kubernetes via Kustomize:

```bash
kubectl apply -k deploy/kubernetes/overlays/k3s
```

Urlog services are stateless `Deployment` resources: `urlog-api`, `urlog-worker`, `urlog-operator`. Redpanda, ClickHouse, OpenSearch, and object storage are external dependencies, not pod-local state.

## Repository layout

```
bootstrap/ first-time install contracts, secret backend catalog, target examples
docs/      design notes and repository conventions
internal/  small tested Go packages used by commands and services
modules/   product module docs, contracts, static concept pages
schema/    versioned protobuf contracts (buf)
deploy/    Kustomize install, observability sketches, packaged deployments
examples/  ForgeBoard PaaS sample system
learning/  tutorial tracks for Phase 0 learning
```

See [`docs/repo-layout.md`](docs/repo-layout.md) before adding top-level folders.

[index.html](index.html) links to the static sample pages for each module, ForgeBoard, and the Integration learning track.

## Status

Pre-alpha. Design partner #0 is a production LangGraph retrieval system. `ROADMAP.md` holds the phase plan; `LOG.md` holds the daily build record.

## The name

Norns, wells, and a traded eye — see [docs/about.md](docs/about.md).
