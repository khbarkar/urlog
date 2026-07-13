# Eye

Small name: **Monitor**

Eye is the live reporting and evidence layer for Urlog. It watches the event stream, gathers evidence from Integration, Delivery, Debt, and SecFlow, generates reports, and publishes those reports to humans, machines, and customer document systems.

## Owns

- Live report generation from Redpanda streams.
- Human-readable reports: readiness, release, incident, compliance, and security summaries.
- Machine-readable reports for gates, audits, and downstream tools.
- PDF/document generation.
- Report upload to customer document systems where the plugged LLM/tool provider supports it.
- Report and evidence indexing.
- Operator dashboards and report APIs.

## Default Storage Shape

Eye uses:

- **ClickHouse** as the canonical analytical/evidence store.
- **OpenSearch** as the default open-source report and evidence search index.
- **Object storage** for generated PDFs, document packets, and large report artifacts.
- **Loki** optionally for raw logs when a Prometheus/Grafana stack is already present.

OpenSearch is preferred over a default "ELK" stack because OpenSearch is a cleaner open-source default for self-hosted and air-gapped buyers, while still giving Elasticsearch/Kibana-style search and dashboards.

## Does Not Own

- Build/test/repo integration. That is Integration.
- Release decisions. That is Delivery.
- Incident truth and action scoring. That is Debt.
- Security scanner implementation. That is SecFlow.

## Myth Frame

Eye watches the tree and writes what it sees. SecFlow provides the well water: optional security and dependency signals poured into the system so reports show whether the operational tree is healthy.
