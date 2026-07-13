# Eye

Small name: **Monitor**

Eye is the optional monitoring and security plugin layer. It can assist Integration, Delivery, and Debt by supplying security, vulnerability, deprecation, obsolete dependency, and prompt-injection findings.

## Owns

- Optional DevSecOps plugins.
- Prompt-injection and jailbreak detection on live streams.
- Security deploy-gate signals for Delivery.
- Threat and security evidence for Debt.
- Dependency vulnerability, deprecation, and obsolete-component findings.

## Explicitly Out Of Scope

- General code scanning as a product category.
- Owning CI/CD systems.
- Owning source-control systems.
- Owning storage, transport, or release policy.

Eye may inspect downloaded source artifacts and dependency manifests to produce findings, but it does not become a SAST platform. Findings are scored and handed to Integration/Delivery/Debt.
