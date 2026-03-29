import * as pulumi from "@pulumi/pulumi";
import * as phase from "@jgautheron/pulumi-phase";

const config = new pulumi.Config();
const appId = config.require("appId");
const secretValue = config.requireSecret("secretValue");

// Create a secret in Phase
const secret = new phase.Secret("example-secret", {
    appId: appId,
    env: "development",
    key: "EXAMPLE_KEY",
    value: secretValue,
    path: "/",
});

// Read secrets from Phase
const secrets = phase.getSecretsOutput({
    appId: appId,
    env: "development",
    path: "/",
});

export const secretKey = secret.key;
export const allSecrets = secrets.secrets;
