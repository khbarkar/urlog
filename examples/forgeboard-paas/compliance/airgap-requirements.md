# Air-Gap Requirements

ForgeBoard must run in high-security environments where the customer supplies all external dependencies.

## Required Customer Inputs

- Private container registry.
- Approved base images.
- Offline vulnerability feeds.
- Offline SBOM import/export.
- Customer-managed keys.
- Customer-provided LLM endpoint or deterministic local stub.
- Offline policy bundle update path.

## Forbidden Defaults

- Public LLM provider calls.
- Public image pulls.
- Runtime dependency downloads.
- Telemetry export outside the customer boundary.
- Automatic destructive actions.

## Evidence

The system must record:

- Active compliance profile.
- Active deployment mode.
- Model endpoint identity without leaking credentials.
- Image digest and SBOM ID.
- Offline update bundle ID.
- Human approval for any destructive action.
