package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	t.Helper()
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	return integration.ProgramTestOptions{
		LocalProviders: []integration.LocalDependency{
			{
				Package: "phase",
				Path:    binPath,
			},
		},
	}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	t.Helper()
	base := getBaseOptions(t)
	return base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@jgautheron/pulumi-phase",
		},
	})
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	t.Helper()
	base := getBaseOptions(t)
	return base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	t.Helper()
	goDepRoot := os.Getenv("PULUMI_GO_DEP_ROOT")
	if goDepRoot == "" {
		var err error
		goDepRoot, err = filepath.Abs("../..")
		if err != nil {
			t.Fatal(err)
		}
	}
	rootSdkPath, err := filepath.Abs("../sdk")
	if err != nil {
		t.Fatal(err)
	}

	base := getBaseOptions(t)
	return base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/jgautheron/pulumi-phase/sdk=" + rootSdkPath,
		},
		Env: []string{
			"PULUMI_GO_DEP_ROOT=" + goDepRoot,
		},
	})
}

func getCwd(t *testing.T) string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}
	return cwd
}
