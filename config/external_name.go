package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Fabric resources
	"equinix_fabric_connection":        config.IdentifierFromProvider,
	"equinix_fabric_cloud_router":      config.IdentifierFromProvider,
	"equinix_fabric_network":           config.IdentifierFromProvider,
	"equinix_fabric_routing_protocol":  config.IdentifierFromProvider,
	"equinix_fabric_service_profile":   config.IdentifierFromProvider,
	
	// Network Edge resources
	"equinix_network_acl_template":     config.IdentifierFromProvider,
	"equinix_network_bgp":              config.IdentifierFromProvider,
	"equinix_network_device":           config.IdentifierFromProvider,
	"equinix_network_device_link":      config.IdentifierFromProvider,
	"equinix_network_ssh_key":          config.IdentifierFromProvider,
	"equinix_network_ssh_user":         config.IdentifierFromProvider,
	
	// Note: We exclude all equinix_metal_* resources as they are deprecated
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
