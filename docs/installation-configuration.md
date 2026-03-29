---
title: Phase Provider Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Phase provider for Pulumi.
layout: package
---

The Phase provider for Pulumi allows you to manage secrets in your Phase.dev applications using Pulumi.

## Installation

The Phase provider is available as a package in all supported Pulumi languages:

- JavaScript/TypeScript: [`@jgautheron/pulumi-phase`](https://www.npmjs.com/package/@jgautheron/pulumi-phase)
- Python: [`pulumi_phase`](https://pypi.org/project/pulumi-phase/)
- Go: [`github.com/jgautheron/pulumi-phase/sdk/go/phase`](https://pkg.go.dev/github.com/jgautheron/pulumi-phase/sdk/go/phase)

### Provider Binary

The Phase provider binary is a third-party binary. It can be installed using the `pulumi plugin` command:

```sh
pulumi plugin install resource phase <version> --server github://api.github.com/jgautheron/pulumi-phase
```

Replace the version string with your desired version.

## Configuration

You must configure the Phase provider with a valid token before it can manage secrets.

### Authentication

The provider supports both service tokens and personal access tokens (PATs):

- `phase:phaseToken` (environment: `PHASE_TOKEN`, `PHASE_SERVICE_TOKEN`, or `PHASE_PAT_TOKEN`) - The Phase authentication token.

### Setting Credentials

There are two ways to provide credentials:

1. **Environment variables** (recommended for CI/CD):

    ```sh
    export PHASE_TOKEN=pss_service:v2:...
    ```

2. **Pulumi configuration** (recommended for multi-user access):

    ```sh
    pulumi config set phase:phaseToken --secret
    ```

    Make sure to pass `--secret` so the token is encrypted in the Pulumi state.

### Optional Configuration

| Property | Environment Variable | Description |
|---|---|---|
| `phase:host` | `PHASE_HOST` | Custom Phase API host URL (default: `https://api.phase.dev`) |
| `phase:skipTlsVerification` | - | Skip TLS certificate validation (default: `false`) |
