<p align="center">
  <img src="assets/phase-logo.svg" alt="Phase" width="64">
</p>

# Pulumi Phase Provider

A Pulumi provider for managing secrets with [Phase.dev](https://phase.dev), bridged from the [Phase Terraform Provider](https://github.com/phasehq/terraform-provider-phase).

## Installation

### Node.js (TypeScript/JavaScript)

```bash
npm install @jgautheron/pulumi-phase
```

### Python

```bash
pip install pulumi_phase
```

### Go

```go
import "github.com/jgautheron/pulumi-phase/sdk/go/phase"
```

## Configuration

The provider requires a Phase token for authentication:

```bash
pulumi config set phase:phaseToken --secret <your-token>
```

Or set one of these environment variables:
- `PHASE_TOKEN`
- `PHASE_SERVICE_TOKEN`
- `PHASE_PAT_TOKEN`

### Optional configuration

| Property | Environment Variable | Description |
|---|---|---|
| `phase:host` | `PHASE_HOST` | Custom Phase API host URL |
| `phase:skipTlsVerification` | — | Skip TLS certificate validation (default: `false`) |

## Resources

### `phase.Secret`

Manages a single secret in Phase.

```typescript
import * as phase from "@jgautheron/pulumi-phase";

const secret = new phase.Secret("my-secret", {
    appId: "your-app-id",
    env: "production",
    key: "DATABASE_URL",
    value: pulumi.secret("postgres://..."),
    path: "/",
    tags: ["database"],
});
```

#### Properties

| Input | Type | Required | Description |
|---|---|---|---|
| `appId` | `string` | Yes | Application ID |
| `env` | `string` | Yes | Environment name |
| `key` | `string` | Yes | Secret key |
| `value` | `string` (secret) | Yes | Secret value |
| `comment` | `string` | No | Comment |
| `path` | `string` | No | Path (default: `/`) |
| `tags` | `string[]` | No | Tags |
| `override` | `SecretOverride` | No | Personal override |

## Data Sources

### `phase.getSecrets`

Retrieves secrets from Phase.

```typescript
const secrets = await phase.getSecrets({
    appId: "your-app-id",
    env: "production",
    path: "/",
});
```

#### Properties

| Input | Type | Required | Description |
|---|---|---|---|
| `appId` | `string` | Yes | Application ID |
| `env` | `string` | Yes | Environment name |
| `path` | `string` | No | Path filter |
| `key` | `string` | No | Key filter |
| `tags` | `string[]` | No | Tag filter (OR logic) |

## Development

### Prerequisites

- Go 1.22+
- Node.js 18+
- Python 3.8+
- Pulumi CLI

### Build

```bash
make tfgen       # Generate schema
make provider    # Build provider binary
make build_sdks  # Generate and build all SDKs
```
