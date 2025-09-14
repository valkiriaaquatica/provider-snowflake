package snowflakegrantapplicationrole

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_grant_application_role", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeGrantApplicationRole"
	})
}
