package snowflakerowaccesspolicy

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_row_access_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeRowAccessPolicy"
	})
}
