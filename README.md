<p align="center">
  <img src="logo.png" alt="Urlog" width="360">
</p>

<h1 align="center">Urlog</h1>

<p align="center"><em>The primordial law. The original log.</em></p>

<p align="center">
  The SLO engine and incident layer for production AI systems —<br>
  quality error budgets, burn-rate alerting, and audit-grade records for agents in production.
</p>

---

## Why

Every AI observability tool on the market is a **dev-loop** tool: trace, debug, eval, iterate. Urlog is the **ops loop**: quality SLOs, error budgets, multi-window burn-rate alerting, incident lifecycle, and AI Act Article 12 audit records. Tracing tells you what your agent did. Urlog tells you whether someone should be paged about it — and proves to a regulator what happened when they ask.

Urlog ingests standard OTLP (OTel GenAI semantic conventions). No proprietary SDK. If your system is instrumented for anything, it is instrumented for Urlog.

## The modules

In Norse myth, three sisters sit at a well beneath the world tree and lay down *ørlǫg* — the primordial law: what has become, what is becoming, and what is owed. Odin gave an eye to that well in exchange for sight.

| Module | Greek | Role | What it owns |
|---|---|---|---|
| **Rheo** | ῥέω, *"I flow"* | Integration | OTLP ingestion, Redpanda consumers, tiered eval workers (classifiers on 100% of traffic, LLM-judge on stratified samples) |
| **Chreos** | χρέος, *"debt"* | Deployment | Quality SLOs, error budgets, multi-window multi-burn-rate alerting, eval-gated releases |
| **Aitia** | αἰτία, *"cause"* | Troubleshooting | Session forensics, incident lifecycle, hash-chained immutable audit log, AI Act Article 12 retention |
| **Auga** | Old Norse, *"eye"* | Security | Prompt-injection detection on the live stream, security deploy gates in Chreos, threat forensics in Aitia |

The modules are separate systems bound by one contract: the versioned protobuf schema in [`schema/`](schema/). Rheo writes it, Chreos reads it to gate, Aitia queries it forever, Auga annotates it. A change that touches two modules without going through the schema is a bug in the change.

## Architecture in one paragraph

Spans arrive over OTLP gRPC, buffer through Redpanda, and land in ClickHouse as narrow rows — full prompt/completion payloads live in object storage behind `payload_ref`. The unit of analysis is the **session with a goal-level outcome**, not the request: agent failures are multi-step causal chains and the data model has to admit that. Eval scores are a separate stream joined at query time, each stamped with `evaluator_version`, because judges drift exactly like the systems they judge. Chreos rolls scores into SLIs, tracks the error budget, and alerts on burn rate — fast-burn pages, slow-burn tickets.

## Quickstart (dev)

```bash
git clone <repo-url> && cd urlog
docker compose -f deploy/docker-compose.dev.yml up   # ClickHouse + Redpanda + otel-collector
```

Point any OTel GenAI-instrumented app (OpenLLMetry, OpenInference) at `localhost:4317` and spans start flowing.

## Repository layout

```
schema/    the contract — protobuf, versioned, breaking changes need their own MR
rheo/      integration
chreos/    deployment
aitia/     troubleshooting
auga/      security
deploy/    dev compose stack + Helm chart (EU-first, self-hostable)
assets/    brand — SVG sources of record
```

## Status

Pre-alpha. Design partner #0 is a production LangGraph retrieval system. Follow `ROADMAP.md` for the phase plan and `LOG.md` for the daily build record.

## The name

*Ørlǫg* — ór ("primordial") + lǫg ("laws, that which is laid down") — is the fate-decree the Norns carve at the well. Rendered in ASCII it becomes **Urlog**: the ur-log, the record that precedes all others. For once, the pun is etymologically load-bearing.
