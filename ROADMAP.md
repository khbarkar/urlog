# Urlog — Roadmap

> **Urlog** (ór + lǫg, "the primordial law") — the SLO engine and incident layer for production AI systems. Four modules: **Rheo** (integration — everything flows), **Chreos** (deployment — the debt), **Aitia** (troubleshooting — the cause), **Auga** (security — the eye in the well).

**Working rule:** every day, one entry in `LOG.md`: date · built/learned · next smallest step. If a day allows only 20 minutes, do the "next smallest step" and log it. The streak matters more than the size of the increment.

---

## MR #1 — `chore: scaffold urlog monorepo`

**Branch:** `init/scaffold` → `main`

### Description

Establishes the Urlog monorepo as a modular monolith: four modules with hard internal
boundaries, one versioned schema package as the only permitted interface between them.
No business logic in this MR — it exists to make every future MR small.

**Why monorepo / modular monolith:** one engineer, four modules. Separate repos now =
process overhead with no isolation benefit. The `schema/` package is versioned like a
public API from commit one, so extracting any module into its own repo later is a
`git filter-repo` operation, not a rewrite.

**Why the schema is the contract:** Rheo writes it, Chreos reads it to gate, Aitia
queries it forever, Auga annotates it. If a change touches two modules without going
through `schema/`, the change is wrong.

### File tree

```
urlog/
├── README.md                 # mission, module map, name lore (one paragraph each)
├── ROADMAP.md                # this file
├── LOG.md                    # daily log — one line per day, never skip "Next:"
├── CLAUDE.md                 # project context for Claude Code sessions
├── schema/                   # THE contract — protobuf, versioned, changes need own MR
│   ├── buf.yaml
│   └── proto/urlog/v0/
│       ├── session.proto     # goal-level outcome: the unit of analysis
│       ├── span.proto        # OTel GenAI semconv-aligned nested spans
│       └── eval.proto        # eval scores as separate stream, joined at query time
├── rheo/                     # integration: OTLP collector cfg, Redpanda consumers, eval workers
│   └── README.md             # module scope + explicit non-goals
├── chreos/                   # deployment: SLO engine, error budgets, burn-rate, release gates
│   └── README.md
├── aitia/                    # troubleshooting: trace/session queries, incidents, audit record
│   └── README.md
├── auga/                     # security: injection detection, security deploy gates, threat forensics
│   └── README.md
├── assets/brand/             # logos: SVG source + PNG exports per module
├── deploy/
│   └── docker-compose.dev.yml  # ClickHouse + Redpanda + otel-collector, one command up
└── .github/workflows/ci.yml    # buf lint + buf breaking + compose config validation
```

### schema/proto/urlog/v0 — v0 sketch (full schema is its own MR)

```protobuf
message Session {          // the unit of analysis — NOT the request
  bytes  session_id = 1;
  string tenant_id  = 2;
  string agent_name = 3;   // e.g. "retrieval-nord"
  Outcome outcome   = 4;   // goal-level: SUCCESS / FAILURE / ABANDONED / UNKNOWN
  int64  started_unix_nano = 5;
  int64  ended_unix_nano   = 6;
  uint64 total_input_tokens  = 7;
  uint64 total_output_tokens = 8;
  string prompt_bundle_version = 9;  // what Chreos gates on
}

message Span {             // OTel GenAI semconv-aligned; nests via parent_span_id
  bytes  span_id = 1;
  bytes  parent_span_id = 2;
  bytes  session_id = 3;
  SpanKind kind = 4;       // LLM_CALL / TOOL_CALL / RETRIEVAL / AGENT_STEP
  map<string,string> gen_ai_attributes = 5;  // gen_ai.* passthrough
  bytes  payload_ref = 6;  // pointer into blob store, not inline (COGS)
}

message EvalScore {        // separate stream — joined to sessions/spans at query time
  bytes  target_id = 1;    // session_id or span_id
  string metric = 2;       // "task_completion", "faithfulness", "injection_risk"
  double score = 3;
  Tier   tier = 4;         // CLASSIFIER (100%) / JUDGE (sampled) / HUMAN
  string evaluator_version = 5;   // judges drift too — version them
}
```

### Definition of done for MR #1

- [ ] `docker compose -f deploy/docker-compose.dev.yml up` brings up ClickHouse, Redpanda, otel-collector, all healthy
- [ ] `buf lint` and `buf breaking` pass in CI; proto compiles to Go + Python stubs
- [ ] Each module README states its scope AND its non-goals (one paragraph each)
- [ ] `CLAUDE.md` contains: mission, current phase, standing decisions table, module boundaries
- [ ] First `LOG.md` line written
- [ ] Logos committed to `assets/brand/` (SVG source of record)

---

## Positioning (read this when motivation drifts)

- Market is crowded on **tracing** (Langfuse/ClickHouse, LangSmith, Phoenix, Braintrust, Helicone, Laminar). Do not compete there.
- Nobody has built the **ops loop**: quality SLOs, error budgets, multi-window burn-rate alerting, incident lifecycle, audit-grade logging.
- Structural moats: (1) statistical alerting depth — SRE math applied to noisy eval-score time series (Chreos); (2) EU sovereignty + AI Act compliance — Article 12 record-keeping, high-risk obligations from Aug 2026 (Aitia); (3) OTel-native backend — ingest standard OTLP, be swappable-in behind competitors' instrumentation (Rheo).
- **Anti-goals:** no proprietary SDK, no prompt playground, no general APM, no training/fine-tuning observability.

---

## Phase 0 — Pain harvesting (≈2 weeks)

Goal: a spec derived from real pain, not imagination. Design partner #0 = retrieval-nord.

**Build track**
- [ ] Land MR #1 (scaffold above)
- [ ] Instrument retrieval-nord with OpenLLMetry (OTel GenAI semconv)
- [ ] Self-host Langfuse (docker compose) — ship retrieval-nord spans to it
- [ ] Self-host Arize Phoenix side by side — same spans, compare data models
- [ ] Run both for 2 weeks against real retrieval-nord traffic/tests

**Learn track**
- [ ] OTel GenAI semantic conventions spec — read in full, note `gen_ai.*` attributes
- [ ] OpenInference spec — note where it diverges from OTel semconv
- [ ] Read OpenLLMetry source (instrumentation patterns, span shaping)
- [ ] Langfuse ClickHouse schema (open source — read their migrations)

**Exit artifact:** `SPEC.md` — every question the incumbent tools could not answer, ranked. ("Is quality degrading? Should someone be paged? Which prompt version caused this? What did this session cost and why?")

---

## Phase 1 — Vertical slice (≈8–10 weeks)

Goal: one signal, end to end, through all four modules. Demo: "provider silently degraded → paged in 4 minutes with offending trace cluster attached."

**Build track (in dependency order)**
- [ ] `schema/` v0 → v1: full Session/Span/EvalScore contract (own MR, reviewed hard)
- [ ] **Rheo:** OTLP gRPC collector → Redpanda → batch inserts to ClickHouse
- [ ] **Aitia:** ClickHouse tables (sessions, spans, eval_scores), TTL tiering hot → S3/Parquet
- [ ] **Rheo:** stratified sampler (oversample errors, long sessions, new prompt versions)
- [ ] **Rheo:** tier-1 eval — cheap classifier on 100% of traffic (heuristics + small model)
- [ ] **Rheo:** tier-2 eval — LLM-as-judge on stratified sample, one metric (task completion)
- [ ] **Chreos:** rolling eval-score SLI → error budget → multi-window multi-burn-rate alerts (fast-burn pages, slow-burn tickets)
- [ ] **Chreos:** alert sink — webhook → phone. Grafana dashboard only, no product UI yet
- [ ] Chaos test: deliberately degrade retrieval-nord (weaker model, corrupted retrieval), measure time-to-page

**Learn track**
- [ ] ClickHouse deep dive: MergeTree engines, materialized views, ZSTD on JSON payloads, TTL
- [ ] Google SRE Workbook ch. 5 (alerting on SLOs) — transfer multi-window burn-rate math to quality SLIs
- [ ] Judge calibration: position bias, verbosity bias, self-preference; Zheng et al. "Judging LLM-as-a-Judge"; Prometheus 2 paper
- [ ] Meta-eval: judge-vs-human agreement, Cohen's kappa
- [ ] Change-point detection: CUSUM, EWMA on eval time series
- [ ] Drift: PSI / KL divergence on embedding distributions

**Exit artifact:** recorded chaos-test demo + writeup. This is the fundraising / design-partner asset.

---

## Phase 2 — Compliance + the eye (≈6–8 weeks)

Goal: convert "another observability tool" into "the thing the compliance officer asks for" — and open Auga's eye.

**Build track**
- [ ] **Aitia:** immutable audit log — append-only object storage, hash-chained
- [ ] **Aitia:** retention policies mapped to AI Act Article 12 record-keeping
- [ ] **Aitia:** RBAC + access audit trail
- [ ] **Aitia:** incident lifecycle — quality incident → owner → runbook → postmortem template
- [ ] **Auga v0:** prompt-injection / jailbreak classifier on Rheo's live stream, scores into `eval_scores` (metric: `injection_risk`)
- [ ] **Auga:** security gate in Chreos — block release if injection-susceptibility evals regress
- [ ] EU-hosted deployment (Hetzner/OVH) + self-host Helm chart

**Learn track**
- [ ] AI Act Articles 12 (logging), 26 (deployer obligations); high-risk timeline from Aug 2026
- [ ] Prompt-injection taxonomy + detection literature (direct, indirect/retrieval-borne — retrieval-nord is a perfect testbed)
- [ ] What auditors actually ask for (interview 2–3 compliance people — retrieval-nord's EU legal network helps here)

**Exit artifact:** compliance mapping doc (control → Article) + 2 design-partner conversations booked.

---

## Standing decisions (do not relitigate daily)

| Decision | Choice | Why |
|---|---|---|
| Umbrella + modules | Urlog → Rheo / Chreos / Aitia / Auga | Ur-log pun is etymologically legitimate; Greek modules carry function without labeling it |
| Repo shape | Modular monolith, `schema/` as sole inter-module interface | One engineer; extraction later is cheap if the contract is respected |
| Instrumentation | Ingest standard OTLP, no own SDK | Swappable-in behind any competitor's instrumentation |
| Store | ClickHouse | Industry standard for this workload; compression = COGS |
| Unit of analysis | Session with goal-level outcome | Agent failures are multi-step causal chains |
| Eval strategy | Tiered: classifiers 100%, judge sampled, human on disagreement | Eval bill must stay below customer inference bill |
| Alerting | Multi-window multi-burn-rate on quality SLOs | The moat; SRE math nobody else ships |
| Hosting | EU-first, self-hostable | Sovereignty + AI Act wedge |
| Name verification | EUIPO class 42+9, GitHub org, .io/.dev/.is for all five names | Before anything is printed or registered |

## Daily log convention (`LOG.md`)

```
2026-07-13 · B: MR #1 scaffold drafted, names locked · Next: create GitHub org, land MR #1
2026-07-14 · L: OTel GenAI semconv §1-3 · Next: finish semconv, note tool-call attrs
```

`B:` built, `L:` learned, `B/L:` both. One line. Never skip the "Next:" field.
