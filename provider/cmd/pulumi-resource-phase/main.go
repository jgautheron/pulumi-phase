package main

import (
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"

	phase "github.com/jgautheron/pulumi-phase/provider"
	"github.com/jgautheron/pulumi-phase/provider/pkg/version"
)

//go:embed schema.json
var pulumiSchema []byte

func main() {
	tfbridge.Main("phase", version.Version, phase.Provider(), pulumiSchema)
}
