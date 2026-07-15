# Eye

Small name: **Monitor**

Eye is the live reporting and evidence layer for Urlog. It watches the event stream, gathers metadata and evidence references from Integration, Delivery, Debt, and SecFlow, generates reports, and publishes those reports to humans, machines, and customer document systems.

Eye is **not** a raw log collector. It stores compact operational metadata, policy decisions, scores, and references to evidence. Raw application logs, full LLM payloads, full command output, and large transcripts belong in dedicated log systems or object storage.

## Owns

- Live report generation from Redpanda/Kafka-compatible metadata streams.
- Human-readable reports: readiness, release, incident, compliance, and security summaries.
- Machine-readable reports for gates, audits, and downstream tools.
- PDF/document generation.
- Report upload to customer document systems where the plugged LLM/tool provider supports it.
- Metadata and evidence-reference indexing.
- Operator dashboards and report APIs.

## Default Storage Shape

Eye uses by default:

- **ClickHouse** as the canonical analytical metadata/evidence store.
- **Object storage** for generated PDFs, document packets, command transcripts, sampled payloads, and large report artifacts.
- **OpenSearch** optionally as the default open-source report and evidence search index.
- **Loki** optionally for live tail/raw logs when a Prometheus/Grafana stack is already present.

OpenSearch is preferred over a default "ELK" stack for optional search because OpenSearch is a cleaner open-source default for self-hosted and air-gapped buyers, while still giving Elasticsearch/Kibana-style search and dashboards.

Customer-provided backends such as Elasticsearch, cloud logging, Loki, Kafka, or managed ClickHouse can be used through adapters. The Eye event schema remains Urlog-owned.

## Capture Policy

Default capture is metadata-first:

- 100% operation metadata.
- 100% errors, security findings, policy decisions, and action records.
- Sampled prompts/completions and large tool payloads.
- Large command stdout/stderr stored by artifact reference.
- Secret values never stored, logged, traced, reported, or sent to the LLM.

## Does Not Own

- Build/test/repo integration. That is Integration.
- Release decisions. That is Delivery.
- Incident truth and action scoring. That is Debt.
- Security scanner implementation. That is SecFlow.
- Raw log collection. Use Loki, OpenSearch, Elasticsearch, cloud logging, or another dedicated log platform.

## Myth Frame

Eye watches the tree and writes what it sees. SecFlow provides the well water: optional security and dependency signals poured into the system so reports show whether the operational tree is healthy.
