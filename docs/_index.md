---
title: Phase
meta_desc: Provides an overview of the Phase Provider for Pulumi.
layout: package
---

The Phase provider for Pulumi allows you to manage secrets in [Phase.dev](https://phase.dev) using infrastructure as code.

## Example

{{< chooser language "typescript,python,go" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as phase from "@jgautheron/pulumi-phase";

const secret = new phase.Secret("my-secret", {
    appId: "your-app-id",
    env: "production",
    key: "DATABASE_URL",
    value: pulumi.secret("postgres://..."),
    path: "/",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_phase as phase

secret = phase.Secret("my-secret",
    app_id="your-app-id",
    env="production",
    key="DATABASE_URL",
    value=pulumi.Output.secret("postgres://..."),
    path="/",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "github.com/jgautheron/pulumi-phase/sdk/go/phase"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := phase.NewSecret(ctx, "my-secret", &phase.SecretArgs{
            AppId: pulumi.String("your-app-id"),
            Env:   pulumi.String("production"),
            Key:   pulumi.String("DATABASE_URL"),
            Value: pulumi.ToSecret(pulumi.String("postgres://...")).(pulumi.StringOutput),
            Path:  pulumi.String("/"),
        })
        return err
    })
}
```

{{% /choosable %}}

{{< /chooser >}}
