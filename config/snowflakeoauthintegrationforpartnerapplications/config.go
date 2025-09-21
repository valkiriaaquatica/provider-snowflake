package snowflakeoauthintegrationforpartnerapplications

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_oauth_integration_for_partner_applications", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeOAuthIntegrationForPartnerApplications"
	})
}
