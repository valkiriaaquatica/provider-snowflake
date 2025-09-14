package snowflakesecondaryconnection

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_secondary_connection", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeSecondaryConnection"
	})
}
