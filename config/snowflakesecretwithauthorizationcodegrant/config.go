package snowflakesecretwithauthorizationcodegrant

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_secret_with_authorization_code_grant", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeSecretWithAuthorizationCodeGrant"
	})
}
