# CLAUDE.md — Urlog

Context for every Claude Code session in this repo. Read fully before doing anything.

## What this is

Urlog is the SLO engine and incident layer for production AI systems: quality error
budgets, multi-window burn-rate alerting, incident lifecycle, AI Act Article 12 audit
records. It is an **ops-loop** product. It is NOT a tracing UI, prompt playground,
general APM, or training observability tool — if a task drifts toward those, stop and say so.

Kristin is a senior SRE building this solo, daily, in small increments. Your job is to
keep increments small, decisions consistent, and scope tight — not to be maximally
agreeable.

## Module boundaries (hard)

| Dir | Module | Owns | Does NOT own |
|---|---|---|---|
| `schema/` | the contract | protobuf: Session, Span, EvalScore | any logic |
| `rheo/` | integration | OTLP ingest, Redpanda consumers, tiered eval workers, sampler | storage schemas, alerting, gates |
| `chreos/` | deployment | SLOs, error budgets, burn-rate alerting, release gates | eval computation, trace queries |
| `aitia/` | troubleshooting | ClickHouse tables, session forensics, incidents, audit log | ingestion, alert policy |
| `auga/` | security | injection detection, security gate logic, threat forensics | its own storage or transport — it rides the others |

**The rule:** modules communicate only through `schema/` types. If a change requires two
modules to know about each other's internals, the change is wrong — redesign it through
the contract. Schema changes get their own MR, never smuggled into a feature MR.
`buf breaking` must pass; a breaking change to v(n) means creating v(n+1).

## Standing decisions — do not relitigate, do not "improve"

- **No proprietary SDK.** Ingest standard OTLP (OTel GenAI semconv / OpenInference). If asked to write a client SDK, push back.
- **ClickHouse** is the store. Do not propose Postgres/Timescale/Elastic alternatives.
- **Session with goal-level outcome** is the unit of analysis, not the request.
- **Payloads live in object storage** behind `payload_ref`. Never inline prompts/completions into ClickHouse rows.
- **Eval scores are a separate stream**, joined at query time, always stamped with `evaluator_version`.
- **Tiered evals:** classifiers on 100%, LLM-judge on stratified samples, human on disagreement. Never propose judging 100% of traffic with a frontier model.
- **Alerting is multi-window multi-burn-rate** (SRE Workbook ch. 5 math). No naive thresholds.
- **EU-first, self-hostable.** No hard dependencies on US-only managed services.

If a task genuinely requires revisiting one of these, say so explicitly and stop —
that's a ROADMAP decision, not a coding-session decision.

## How to work in this repo

- **Small MRs.** One concern per MR. If the diff crosses a module boundary and `schema/`, split it. Target: reviewable in 15 minutes.
- **Plan before code** for anything non-trivial: state the approach in 3-5 lines, wait for a go, then implement. Do not scaffold speculative abstractions.
- **No placeholder/demo code paths.** Everything merged must run against the dev compose stack. If it can't be tested end-to-end yet, it doesn't merge.
- **Tests ride with the change.** For Chreos math (burn rates, budgets, change-point detection): property-based tests with known-answer fixtures, not just happy-path unit tests.
- **Don't invent metrics or eval names.** The metric namespace (`task_completion`, `faithfulness`, `injection_risk`, ...) is defined in `schema/` docs — extend it there first.
- **Migrations are forward-only.** ClickHouse schema changes ship as numbered migrations; never edit an applied migration.
- **Update `LOG.md` is Kristin's job, not yours** — but end every session by proposing the one-line entry: `date · B/L: what · Next: smallest step`.

## Keep-me-in-check rules (Kristin's own guardrails — enforce them)

1. **Scope creep challenge.** If a session's request expands beyond the current
   ROADMAP phase, name it: "this is Phase 2 work, current phase is X — park it or
   re-plan?" Do not silently build ahead.
2. **Anti-goal tripwire.** If asked for anything resembling a prompt playground,
   custom SDK, trace-viewer UI polish, or fine-tuning observability: refuse the frame,
   cite the anti-goals, offer the nearest in-scope alternative.
3. **Demo-quality tripwire.** If asked to "just hack it in" or "mock it for now" on a
   core path (ingest, schema, SLO math, audit log): push back once, clearly. Mocks are
   fine in tests, never on the write path of the audit log — that path is the product.
4. **Daily increment bias.** Prefer the smallest merge-able step over the impressive
   multi-day branch. A landed 50-line MR beats an open 800-line one.
5. **Honest uncertainty.** For ClickHouse tuning, burn-rate math, and eval statistics:
   if unsure, say so and propose how to verify (benchmark, known-answer test, paper).
   No confident hand-waving on the parts that are the moat.
6. **Consistency check on naming.** Identifiers are ASCII: `urlog`, `rheo`, `chreos`,
   `aitia`, `auga`. Icelandic/Greek spellings live in docs and marketing only.

## SRE translation table (use these framings, they map to Kristin's background)

- Eval suite = test suite + canary analysis
- Quality SLI = golden signal; eval-score time series = latency histogram equivalent
- Error budget on task_completion = error budget on availability — same math, noisier signal
- Judge drift = monitoring-the-monitoring; `evaluator_version` = probe version
- Prompt bundle version = deploy artifact; Chreos gate = progressive delivery analysis step
- Schema package = API contract between services; `buf breaking` = backward-compat CI gate

## Environment notes

- Dev stack: `docker compose -f deploy/docker-compose.dev.yml up` (ClickHouse, Redpanda, otel-collector)
- Design partner #0: retrieval-nord (LangGraph + Qdrant + MCP) — realistic traffic source for end-to-end tests
- Python via venv; pip installs in sandboxed environments may need `--break-system-packages`
- macOS host: MTU is pinned to 1492 for Wi-Fi (PPPoE PMTU); if Node/TLS network calls
  reset mysteriously, check MTU before debugging application code

## Current phase

**Phase 0 — pain harvesting.** MR #1 (scaffold) in flight. Next milestones: OpenLLMetry
instrumentation of retrieval-nord, Langfuse + Phoenix side-by-side, `SPEC.md` from real
unanswered questions. Anything not serving Phase 0 exit criteria gets parked in ROADMAP.

*(Update this section every time a phase closes — a stale "current phase" makes every
session start with wrong context.)*
