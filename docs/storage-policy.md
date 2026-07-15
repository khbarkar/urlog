# Storage Policy

Urlog is not a log collector.

Urlog stores operational metadata, evidence records, scores, and references. It does not try to replace Loki, OpenSearch, Elasticsearch, CloudWatch Logs, Azure Monitor, GCP Logging, Splunk, or any other log platform.

## Default Rule

```text
Observe every operation.
Store compact metadata by default.
Store large or sensitive payloads only by reference.
Use dedicated log systems for raw log volume.
```

## What Urlog Stores

Urlog stores structured metadata such as:

- Run ID, target ID, module, phase, action, status.
- Timestamps, durations, exit codes, and result summaries.
- Tool names and versions.
- MCP server names and operation names.
- Secret references used, never secret values.
- Evidence references.
- Readiness scores and report summaries.
- Redaction status and policy decisions.

## What Urlog Does Not Store By Default

Urlog does not store high-volume raw logs by default:

- Full application logs.
- Full LLM prompts and completions.
- Full tool arguments/results.
- Full command stdout/stderr.
- Full retrieved documents.
- Full customer payloads.

Those belong in purpose-built systems or object storage, with Urlog storing references.

## Storage Roles

Default MVP shape:

```text
Redpanda/Kafka-compatible transport:
  short-lived operational event movement

ClickHouse:
  compact append-only metadata and evidence facts

Object storage:
  command transcripts, report artifacts, PDFs, bundles, sampled payloads

Eye:
  report and query surface over metadata plus selected artifact refs
```

Optional customer sinks:

```text
Loki:
  live tail and raw operational logs

OpenSearch / Elasticsearch:
  full-text search and report/evidence indexing

Cloud log systems:
  customer-owned raw log retention
```

## LLM Capture Policy

Default LLM capture should be tiered:

```yaml
capture_policy:
  metadata: 100%
  errors: 100%
  tool_calls: 100%
  prompts: sampled
  completions: sampled
  payload_sample_rate: 0.01
  payload_retention_days: 7
  metadata_retention_days: 90
```

Security or incident policies may temporarily increase capture for a target, model, prompt version, or incident scope. Even then, secrets remain forbidden and large payloads should be stored as artifact references.

## Product Boundary

Urlog can integrate with log collectors, but it is not one.

Eye reports what happened, where evidence lives, and which policy decisions were made. It should not become a raw log lake.

