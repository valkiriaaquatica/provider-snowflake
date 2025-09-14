package snowflakemaskingpolicy

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("snowflake_masking_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = "SnowflakeMaskingPolicy"
	})
}
