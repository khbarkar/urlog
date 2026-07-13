# Compliance Profiles

ForgeBoard uses compliance profiles to map market and sector expectations to concrete product controls. A profile is not legal advice; it is an implementation checklist for what the autonomous operations layer must collect, enforce, and export.

## Profile Model

Each profile maps:

- Market or sector scope.
- Regulatory or standards anchors.
- Required operational controls.
- Evidence that Urlog/ForgeBoard must retain.
- Controls that block automation or require human approval.

## Profiles

| Profile | Use |
|---|---|
| `profiles/eu.yaml` | EU-first fintech, medical, and green-tech operators |
| `profiles/us.yaml` | US SaaS, fintech, healthcare, public-sector-adjacent buyers |
| `profiles/uk.yaml` | UK regulated SaaS, fintech, health, and public-sector buyers |
| `profiles/global-enterprise.yaml` | Baseline for multinational enterprise procurement |

## Product Rule

The platform should select one primary jurisdiction profile and any number of sector overlays. The strictest applicable control wins.

Example:

```text
primary: eu
sector_overlays: [fintech, healthcare]
deployment_mode: air_gapped
```

This means:

- Accessibility checks are mandatory.
- LLM access is customer supplied.
- Destructive actions require human approval.
- Incident and action evidence must be exportable.
- Data residency and retention policy are explicit deployment inputs.
