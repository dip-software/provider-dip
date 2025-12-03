package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/philips-software/terraform-provider-hsdp/hsdp"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/dip-software/provider-dip/config/connect/dbs"
	"github.com/dip-software/provider-dip/config/connect/mdm"
	"github.com/dip-software/provider-dip/config/iam"
	"github.com/dip-software/provider-dip/config/tenant"
)

const (
	resourcePrefix = "hsdp"
	modulePath     = "github.com/dip-software/provider-dip"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("hsdp"),
		ujconfig.WithRootGroup("dip.crossplane.io"),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(hsdp.Provider("v0.71.1")),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		iam.Configure,
		dbs.Configure,
		mdm.Configure,
		tenant.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("hsdp"),
		ujconfig.WithRootGroup("dip.m.crossplane.io"),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(hsdp.Provider("v0.71.1")),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		iam.Configure,
		dbs.Configure,
		mdm.Configure,
		tenant.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
