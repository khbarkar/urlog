# 04 - Dual Shipping

Goal: route the same telemetry to Langfuse and Phoenix for comparison.

## Shape

```text
retrieval-nord
  -> OpenLLMetry / OTel SDK
  -> OTLP collector
  -> Langfuse
  -> Phoenix
```

Use `deploy/phase0-observability/otel-collector.yaml` as the local routing sketch.

## Learn

Record:

- Which exporter path works cleanly.
- Whether trace IDs remain stable.
- Whether either tool mutates or drops attributes.
- Which fields you need for Urlog's `Session`, `Span`, and `EvalScore`.

## Done

One test run appears in both tools with enough shared identity to compare data models.
