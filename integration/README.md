# Integration

Norn name: **Verdandi**  
Small name: **Plan**

Integration owns the live connection between source systems, runtime signals, and the Urlog contract. It is responsible for OTLP ingestion, Redpanda consumers, tiered eval workers, source-repository metadata, build/test signal ingestion, and the operational stream that later powers autonomous cloudops/devops actions.

## Owns

- OTLP ingestion for OTel GenAI / OpenInference / OpenLLMetry traffic.
- Redpanda consumers.
- Tiered eval workers: classifiers on 100%, LLM judge on stratified samples.
- Centralized version-control metadata from connected repos.
- Build, test, dependency, and CI feedback ingestion.
- Security/deprecation/obsolete dependency findings from Eye plugins.
- Readiness scores emitted for Delivery.

## Does Not Own

- Quality SLO math, error budgets, or release policy. That is Delivery.
- Session forensics, incident lifecycle, and audit evidence. That is Debt.
- Security scanner implementation. Eye supplies optional plugins.
- Proprietary SDKs. Integration uses standard instrumentation and existing CI/source APIs.

## ForgeBoard Role

For `examples/forgeboard-paas`, Integration is the module that connects to the repository, builds the frontend and middleware containers, records test feedback, imports Eye plugin findings, and emits a score that Delivery can use to decide whether the system is ready.
