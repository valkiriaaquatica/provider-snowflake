package snowflakesecretwithgenericstring

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_secret_with_generic_string", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeSecretWithGenericString"
	})
}
