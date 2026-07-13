# Company Profile

## ForgeBoard

ForgeBoard sells an event-driven transaction operations PaaS to EU mid-market operators. Customers upload policy bundles that validate transactions, enrich incoming events, and route outcomes to downstream systems.

## Operating Profile

- Staff: 180
- Engineering: 32
- SRE / platform: 4
- Security: 2
- Customers: 80 B2B tenants across fintech, medical operations, and green technology
- Runtime: Kubernetes, initially k3s and Docker Desktop Kubernetes
- Target buyers: companies up to 500 staff that need strong operational automation without a large SRE team

## Product Risk

ForgeBoard is multi-tenant. Customer uploads are useful but dangerous: a single bad bundle can increase stream lag, break schema compatibility, overrun resources, create bad transaction outcomes, or introduce adversarial instructions into downstream AI workflows.

The autonomous ops system must make routine safe actions cheap while keeping destructive actions behind explicit approval.

## Example Tenant Workloads

- Fintech: merchant settlement controls, sanctions-screening handoff, payment exception routing.
- Medical startups: reimbursement workflows, device telemetry billing evidence, appointment no-show charge policies.
- Green technology: carbon-credit transaction evidence, grid-battery settlement events, energy-savings subsidy claims.
