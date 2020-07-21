package terraform

import (
	"fmt"

	"os/user"

	"github.com/gruntwork-io/terratest/modules/testing"
)

// Init calls terraform init and return stdout/stderr.
func Init(t testing.TestingT, options *Options) string {
	out, err := InitE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// InitE calls terraform init and return stdout/stderr.
func InitE(t testing.TestingT, options *Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommandE(t, options, args...)
}

// CustomInit function that calls terraform init with custom azurerm plugin and returns stdout/stderr
func CustomInit(t testing.TestingT, options *Options) string {
	out, err := CustomInitE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// CustomInitE function that calls terraform init with custom azurerm plugin and returns stdout/stderr.
// For now, manually change <user> in the directory path defined for the -plugin-dir flag below, in the args string.
func CustomInitE(t testing.TestingT, options *Options) (string, error) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := []string{"init", fmt.Sprintf("-plugin-dir=/home/%s/go/bin", user.Username), fmt.Sprintf("-plugin-dir=%s", "/usr/local/bin"), fmt.Sprintf("-plugin-dir=%s", ".terraform/plugins/linux_amd64"), fmt.Sprintf("-plugin-dir=/home/%s/.terraform.d/plugins", user.Username), fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommandE(t, options, args...)
}
