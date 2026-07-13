# Phase 0 Observability Harness

This folder contains local learning scaffolding for the roadmap's Langfuse/Phoenix comparison.

It intentionally does not vendor Langfuse or Phoenix compose files. Those projects move quickly; use their official docs for the running services and keep this folder as the Urlog-side routing and note-taking layer.

## Files

- `otel-collector.yaml`: local OTLP routing sketch.
- `.env.example`: endpoints and project identifiers to fill in.

## Expected Shape

```text
retrieval-nord -> OTLP collector -> Langfuse
                              \-> Phoenix
```

The exact exporter endpoints may change with upstream versions. Confirm them against the official docs before running.
