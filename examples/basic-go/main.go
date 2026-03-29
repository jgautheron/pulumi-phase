package main

import (
	"github.com/jgautheron/pulumi-phase/sdk/go/phase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		appID := cfg.Require("appId")
		secretValue := cfg.RequireSecret("secretValue")

		secret, err := phase.NewSecret(ctx, "example-secret", &phase.SecretArgs{
			AppId: pulumi.String(appID),
			Env:   pulumi.String("development"),
			Key:   pulumi.String("EXAMPLE_KEY"),
			Value: secretValue,
			Path:  pulumi.String("/"),
		})
		if err != nil {
			return err
		}

		secrets, err := phase.GetSecretsOutput(ctx, phase.GetSecretsOutputArgs{
			AppId: pulumi.String(appID),
			Env:   pulumi.String("development"),
			Path:  pulumi.StringRef("/"),
		}, nil)
		if err != nil {
			return err
		}

		ctx.Export("secretKey", secret.Key)
		ctx.Export("allSecrets", secrets.Secrets())

		return nil
	})
}
