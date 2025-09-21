package snowflakeoauthintegrationforcustomclients

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_oauth_integration_for_custom_clients", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeOAuthIntegrationForCustomClients"
	})
}
