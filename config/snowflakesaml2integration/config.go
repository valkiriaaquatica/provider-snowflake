package snowflakesaml2integration

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_saml2_integration", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeSAML2Integration"
	})
}
