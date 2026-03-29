package phase

import (
	_ "embed"
	"path"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/jgautheron/pulumi-phase/provider/pkg/version"
	"github.com/phasehq/terraform-provider/shim"
)

const (
	mainPkg = "phase"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-phase/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 shimv2.NewProvider(shim.Provider()),
		Name:              "phase",
		Version:           version.Version,
		DisplayName:       "Phase",
		Publisher:         "jgautheron",
		LogoURL:           "https://raw.githubusercontent.com/jgautheron/pulumi-phase/main/assets/phase-logo.png",
		PluginDownloadURL: "github://api.github.com/jgautheron/pulumi-phase",
		Description:       "A Pulumi provider for managing secrets with Phase.dev.",
		Keywords:          []string{"phase", "secrets", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://phase.dev",
		Repository:        "https://github.com/jgautheron/pulumi-phase",
		GitHubOrg:         "phasehq",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"phase_token": {
				Secret: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PHASE_TOKEN", "PHASE_SERVICE_TOKEN", "PHASE_PAT_TOKEN"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@jgautheron/pulumi-phase",
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName:          "pulumi_phase",
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/jgautheron/pulumi-phase/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			RespectSchemaVersion:           true,
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("phase_", mainMod, tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
