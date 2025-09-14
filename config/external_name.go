/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"snowflake_warehouse":             config.IdentifierFromProvider,
	"snowflake_database":              config.IdentifierFromProvider,
	"snowflake_authentication_policy": config.IdentifierFromProvider,
	"snowflake_account_role":          config.IdentifierFromProvider,
	"snowflake_account":               config.IdentifierFromProvider,
	"snowflake_account_parameter":     config.IdentifierFromProvider,
	"snowflake_api_authentication_integration_with_authorization_code_grant": config.IdentifierFromProvider,
	"snowflake_api_authentication_integration_with_client_credentials":       config.IdentifierFromProvider,
	"snowflake_api_authentication_integration_with_jwt_bearer":               config.IdentifierFromProvider,
	"snowflake_database_role":                              config.IdentifierFromProvider,
	"snowflake_execute":                                    config.IdentifierFromProvider,
	"snowflake_external_oauth_integration":                 config.IdentifierFromProvider,
	"snowflake_grant_account_role":                         config.IdentifierFromProvider,
	"snowflake_grant_database_role ":                       config.IdentifierFromProvider,
	"snowflake_grant_application_role":                     config.IdentifierFromProvider,
	"snowflake_grant_ownership":                            config.IdentifierFromProvider,
	"snowflake_grant_privileges_to_account_role":           config.IdentifierFromProvider,
	"snowflake_grant_privileges_to_database_role":          config.IdentifierFromProvider,
	"snowflake_grant_privileges_to_share":                  config.IdentifierFromProvider,
	"snowflake_legacy_service_user":                        config.IdentifierFromProvider,
	"snowflake_masking_policy":                             config.IdentifierFromProvider,
	"snowflake_network_policy":                             config.IdentifierFromProvider,
	"snowflake_oauth_integration_for_custom_clients":       config.IdentifierFromProvider,
	"snowflake_oauth_integration_for_partner_applications": config.IdentifierFromProvider,
	"snowflake_primary_connection":                         config.IdentifierFromProvider,
	"snowflake_resource_monitor":                           config.IdentifierFromProvider,
	"snowflake_row_access_policy":                          config.IdentifierFromProvider,
	"snowflake_saml2_integration":                          config.IdentifierFromProvider,
	"snowflake_schema":                                     config.IdentifierFromProvider,
	"snowflake_scim_integration":                           config.IdentifierFromProvider,
	"snowflake_secondary_connection":                       config.IdentifierFromProvider,
	"snowflake_secondary_database":                         config.IdentifierFromProvider,
	"snowflake_secret_with_authorization_code_grant":       config.IdentifierFromProvider,
	"snowflake_secret_with_basic_authentication ":          config.IdentifierFromProvider,
	"snowflake_secret_with_client_credentials":             config.IdentifierFromProvider,
	"snowflake_secret_with_generic_string":                 config.IdentifierFromProvider,
	"snowflake_service_user":                               config.IdentifierFromProvider,
	"snowflake_shared_database":                            config.IdentifierFromProvider,
	"snowflake_stream_on_directory_table":                  config.IdentifierFromProvider,
	"snowflake_stream_on_external_table":                   config.IdentifierFromProvider,
	"snowflake_stream_on_table":                            config.IdentifierFromProvider,
	"snowflake_stream_on_view":                             config.IdentifierFromProvider,
	"snowflake_streamlit":                                  config.IdentifierFromProvider,
	"snowflake_tag":                                        config.IdentifierFromProvider,
	"snowflake_tag_association":                            config.IdentifierFromProvider,
	"snowflake_task":                                       config.IdentifierFromProvider,
	"snowflake_user":                                       config.IdentifierFromProvider,
	"snowflake_view":                                       config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
