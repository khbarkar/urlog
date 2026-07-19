# Repository Layout

Keep the repository boring and predictable.

## Root

The root should contain only project-level entry points:

- `README.md`
- `ROADMAP.md`
- `go.mod`
- `Makefile`
- top-level product areas such as `bootstrap/`, `deploy/`, `docs/`, `examples/`, `internal/`, `learning/`, and `modules/`
- static sample entry point `index.html`

Do not add one-off scripts, generated reports, screenshots, temporary files, or implementation experiments to the root.

## Go Code

Production Go code lives under:

```text
internal/
```

Command entry points will later live under:

```text
cmd/urlog
cmd/urlogd
cmd/urlog-operator
```

Small, testable packages are preferred. A package should have a narrow responsibility and table-driven tests before it becomes a dependency of higher-level systems.

## Docs

Design notes that are not the main README or roadmap live under:

```text
docs/
```

Shared image assets live under:

```text
docs/assets/
```

## Modules

Product module docs, contracts, and static concept pages live under:

```text
modules/
```

Current modules:

- `modules/integration/`
- `modules/delivery/`
- `modules/debt/`
- `modules/eye/`
- `modules/secflow/`

## Bootstrap And Deploy

First-install contracts live under:

```text
bootstrap/
```

Kubernetes deployment assets live under:

```text
deploy/kubernetes/
```

## Examples

Customer/sample systems live under:

```text
examples/
```

ForgeBoard is the first sample target.

## Secrets

Secret values do not belong anywhere in the repository.

Allowed:

- `secret://...` references
- role references
- path allowlists
- non-secret examples

Forbidden:

- passwords
- API keys
- private keys
- kubeconfig contents
- bearer tokens
