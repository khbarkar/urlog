# Eye Reporting Architecture

Eye turns Urlog operational metadata and evidence references into reports.

Eye is not a raw log collector. Raw application logs, full LLM payloads, and large command transcripts should stay in dedicated log systems or object storage. Eye stores compact metadata and links to evidence.

## Flow

```text
Redpanda topics
  -> Eye report workers
  -> ClickHouse canonical metadata/evidence tables
  -> object storage for PDFs/docs/transcripts/artifacts
  -> optional OpenSearch report/evidence index
  -> optional Loki or customer log system for raw logs
```

## Inputs

- Integration readiness scores.
- Delivery gate decisions.
- Debt incident and action records.
- SecFlow findings.
- Compliance profile evidence.
- Accessibility evidence.
- Kubernetes and CI metadata events.
- Artifact references for command output and raw logs.

## Outputs

- Live readiness reports.
- Release decision packets.
- Incident summaries.
- Compliance evidence packets.
- Security and dependency summaries.
- PDF or document exports.
- Searchable report/evidence index.

## Backend Choice

Default:

- ClickHouse for canonical analytical metadata.
- Object storage for generated documents, transcripts, and large artifacts.

Optional:

- OpenSearch for full-text, semantic, and report search.
- Loki, Elasticsearch, cloud logging, or another customer system for raw logs.

Not default:

- Elastic default distribution, because licensing is more ambiguous for an open-source-first, self-hosted, air-gapped product story.
- Using Urlog as the customer's primary log collector.
