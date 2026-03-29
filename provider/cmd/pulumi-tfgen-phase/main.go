package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"

	phase "github.com/jgautheron/pulumi-phase/provider"
	"github.com/jgautheron/pulumi-phase/provider/pkg/version"
)

func main() {
	tfgen.Main("phase", version.Version, phase.Provider())
}
