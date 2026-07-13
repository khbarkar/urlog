# Eye Reporting Architecture

Eye turns Urlog operational events into reports and searchable evidence.

## Flow

```text
Redpanda topics
  -> Eye report workers
  -> ClickHouse canonical tables
  -> OpenSearch report/evidence index
  -> object storage for PDFs/docs
  -> optional Loki raw-log retention
```

## Inputs

- Integration readiness scores.
- Delivery gate decisions.
- Debt incident and action records.
- SecFlow findings.
- Compliance profile evidence.
- Accessibility evidence.
- Kubernetes and CI events.

## Outputs

- Live readiness reports.
- Release decision packets.
- Incident summaries.
- Compliance evidence packets.
- Security and dependency summaries.
- PDF or document exports.
- Searchable report index.

## Backend Choice

Default:

- ClickHouse for canonical analytical data.
- OpenSearch for full-text, semantic, and report search.
- Object storage for generated documents.

Optional:

- Loki for raw logs when the customer already runs the Grafana/Prometheus stack.

Not default:

- Elastic default distribution, because licensing is more ambiguous for an open-source-first, self-hosted, air-gapped product story.
