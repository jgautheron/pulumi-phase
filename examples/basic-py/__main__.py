import pulumi
import pulumi_phase as phase

config = pulumi.Config()
app_id = config.require("appId")
secret_value = config.require_secret("secretValue")

secret = phase.Secret("example-secret",
    app_id=app_id,
    env="development",
    key="EXAMPLE_KEY",
    value=secret_value,
    path="/",
)

secrets = phase.get_secrets_output(
    app_id=app_id,
    env="development",
    path="/",
)

pulumi.export("secret_key", secret.key)
pulumi.export("all_secrets", secrets.secrets)
