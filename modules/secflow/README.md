# SecFlow

SecFlow is the optional security-flow module for Urlog.

It hosts third-party security tools and turns their output into normalized findings that Integration, Delivery, Debt, and Eye can use.

## Owns

- Optional DevSecOps/security plugins.
- Prompt-injection and jailbreak detection on live streams.
- Dependency vulnerability findings.
- Deprecated dependency and obsolete runtime findings.
- Container image and SBOM findings.
- CI/red-team/security regression evidence.
- Security gate recommendations for Delivery.

## Candidate Tools

- **LlamaFirewall / PromptGuard 2** for prompt-injection and jailbreak detection.
- **garak** for CI, red-team, and security regression evidence.
- Customer-provided vulnerability, dependency, image, SBOM, and compliance scanners.

## Explicitly Out Of Scope

- General code scanning as a product category.
- Owning CI/CD execution.
- Owning source-control systems.
- Owning release policy.
- Owning incident truth or audit records.

SecFlow may inspect downloaded source artifacts, dependency manifests, SBOMs, image metadata, and security reports to produce findings. It does not become a SAST platform.

## Myth Frame

SecFlow is the water from the well: optional outside signals poured into the system to keep the operational tree from drying out. Eye watches and reports what that water changes.
