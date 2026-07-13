# Urlog First-Time Bootstrap

Bootstrap is the first trust handoff. It tells Urlog where it may run, which model provider it may use, which repository it may inspect, which infrastructure it may touch, and whether it should install an integration system or connect to an existing one.

Bootstrap does **not** store cleartext passwords in Urlog config. It stores references, roles, allowed paths, and approval boundaries.

## Bootstrap Operator

The first actor is the **Bootstrap Operator**.

This can be:

- A human platform/security operator.
- A corporate automation identity.
- A break-glass bootstrap role with short-lived credentials.

The Bootstrap Operator grants Urlog scoped access to:

- An LLM provider or customer-hosted LLM endpoint.
- Source control.
- A secret backend.
- Infrastructure needed for integration work.
- Existing CI/CD systems, or permission to install one.

After bootstrap, normal Integration runs are started by Urlog's service identity and pre-approval policies.

## Required Inputs

### 1. LLM Access

Urlog needs an LLM provider, but the LLM is not the authority.

The LLM may:

- Inspect non-secret repository metadata and manifests.
- Propose target config.
- Propose run plans.
- Produce intents that match Urlog schemas.

The LLM may not:

- Receive secret values.
- Execute actions directly.
- Override policy.
- Invent action IDs.
- Approve destructive actions.

### 2. Repository Access

Urlog needs access to the application repository or artifact source.

Supported first shapes:

- Local path.
- GitHub.
- GitLab.
- Generic Git over SSH/HTTPS.

Future shapes:

- Jenkins source jobs.
- Perforce.
- Artifact-only onboarding.

Repository credentials must be secret references only.

### 3. Secret Backend

Urlog supports a secret adapter model.

Initial backends:

- `local-sops-age` for offline and air-gapped installs.
- `cloud-secret-manager` for AWS, Azure, GCP, and other cloud providers through their secret manager APIs or MCP servers.

Future backends:

- Vault/OpenBao.
- 1Password.
- Bitwarden.
- Customer-specific secret manager adapters.

Config may contain:

- `secret://...` references.
- Cloud role ARNs/resource IDs.
- MCP server names.
- Allowed secret paths.

Config may not contain:

- Passwords.
- API keys.
- OAuth tokens.
- Private keys.
- Kubeconfig content.

### 4. Infrastructure Access

Urlog needs access to the infrastructure required for the integration job.

For Kubernetes-first bootstrap this normally means:

- Kubernetes API access.
- Namespace or cluster-scope depending on install mode.
- Permission to create ephemeral build/test namespaces if selected.
- Access to container registry credentials by reference.
- Access to Redpanda, ClickHouse, OpenSearch, object storage, or managed equivalents.

For non-Kubernetes systems later:

- KVM access through MCP or a constrained runner.
- Bare-metal runner access.
- Remote worker registration.

ESXi is out of scope for the first implementation.

## CI/CD Decision

Bootstrap must answer one question explicitly:

```text
Should Urlog install the integration system, or use an existing one?
```

### Install Mode

Use this when the customer wants Urlog to set up the build/test integration layer.

Urlog needs permission to install selected components, for example:

- Tekton.
- Argo Workflows.
- GitHub Actions runner controller.
- GitLab runner.
- BuildKit.
- Kaniko or equivalent image builder.
- Registry credentials by secret reference.

Install mode requires a declared tool choice. Urlog should not guess which CI/CD system to install.

### Existing Mode

Use this when the customer already has CI/CD.

Urlog needs access to:

- Read pipeline/job definitions.
- Trigger allowed jobs.
- Read logs and artifacts.
- Read test results.
- Read image digests and SBOMs.
- Attach readiness evidence.

Urlog does not own the existing CI/CD system. It consumes and scores evidence.

### Observe-Only Mode

Use this for the safest first onboarding.

Urlog may:

- Inspect repo metadata.
- Inspect manifests.
- Read CI/CD results.
- Read Kubernetes state.
- Generate a readiness packet.

Urlog may not:

- Trigger builds.
- Create namespaces.
- Install components.
- Mutate infrastructure.

## First Bootstrap Flow

```text
Bootstrap Operator chooses install mode
        ↓
Bootstrap Operator provides LLM config
        ↓
Bootstrap Operator provides repo access refs
        ↓
Bootstrap Operator provides secret backend refs
        ↓
Bootstrap Operator provides infrastructure access refs
        ↓
Urlog validates the bootstrap contract
        ↓
Urlog writes non-secret config
        ↓
Urlog starts Integration in observe-only or approved mode
        ↓
Eye and Debt record bootstrap evidence
```

## First Implementation Target

The first implementation should support:

- Urlog running as containers.
- Kubernetes Deployment for `urlog-api`, `urlog-worker`, and later `urlog-operator`.
- Stateless Urlog services.
- State in Redpanda, ClickHouse, OpenSearch, object storage, and the chosen secret backend.
- `local-sops-age` and cloud secret-manager references in config.
- Observe-only mode for ForgeBoard.

## Files

- [`bootstrap-contract.yaml`](bootstrap-contract.yaml): machine-readable first-install contract shape.
- [`secret-backends.yaml`](secret-backends.yaml): supported secret backend types and rules.
- [`examples/forgeboard-bootstrap.yaml`](examples/forgeboard-bootstrap.yaml): example first target for ForgeBoard.

