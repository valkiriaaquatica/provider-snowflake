package snowflakeapiauthenticationintegrationwithjwtbearer

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_api_authentication_integration_with_jwt_bearer", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeApiAuthenticationIntegrationWithJwtBearer"
	})
}
