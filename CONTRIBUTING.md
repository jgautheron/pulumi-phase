# Contributing to pulumi-phase

We appreciate contributions to the Phase Pulumi provider.

## Code of Conduct

Please be respectful and constructive in all interactions.

## Setting Up Your Development Environment

### Prerequisites

- Go 1.22+
- Node.js 18+
- Python 3.8+
- Pulumi CLI
- golangci-lint

You can install all tools at once using [mise](https://mise.jdx.dev/):

```sh
mise install
```

### Building

```sh
make build
```

This will build the provider binary and all SDKs.

## Committing Generated Code

You must generate and check in the SDKs on each pull request containing a code change (e.g. adding a new resource to `resources.go`).

1. Run `make build_sdks` from the root of this repository
2. Open a pull request containing all changes

If a large number of seemingly-unrelated diffs are produced by `make build_sdks`, ensure that the latest dependencies for the provider are installed by running `go mod tidy` in the `provider/` directory.

## Running Integration Tests

The examples and integration tests in this repository will create and destroy real Phase secrets while running. Before running these tests, make sure you have a valid Phase token configured:

```sh
export PHASE_TOKEN=pss_service:v2:...
```

Then run:

```sh
make test
```

## Linting

```sh
make lint_provider
```

## Regenerating CI Configuration

If you modify `.ci-mgmt.yaml`, regenerate the CI workflows:

```sh
make ci-mgmt
```
