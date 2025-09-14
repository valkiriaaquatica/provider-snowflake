/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeaccount"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeaccountparameter"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeaccountrole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeapiauthenticationintegrationwithauthorizationcodegrant"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeapiauthenticationintegrationwithclientcredentials"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeapiauthenticationintegrationwithjwtbearer"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeauthenticationpolicy"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakedatabase"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakedatabaserole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeexecute"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeexternaloauthintegration"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantaccountrole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantapplicationrole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantownership"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantprivilegestoaccountrole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantprivilegestodatabaserole"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakegrantprivilegestoshare"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakelegacyserviceuser"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakemaskingpolicy"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakenetworkpolicy"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeoauthintegrationforcustomclients"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeoauthintegrationforpartnerapplications"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeprimaryconnection"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeresourcemonitor"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakerowaccesspolicy"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesaml2integration"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeschema"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakescimintegration"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecondaryconnection"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecondarydatabase"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecretwithauthorizationcodegrant"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecretwithbasicauthentication"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecretwithclientcredentials"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakesecretwithgenericstring"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeserviceuser"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeshareddatabase"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakestreamlit"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakestreamondirectorytable"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakestreamonexternaltable"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakestreamontable"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakestreamonview"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflaketag"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflaketagassociation"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflaketask"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeuser"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakeview"
	"github.com/valkiriaaquatica/provider-snowflake/config/snowflakewarehouse"
)

const (
	resourcePrefix = "snowflake"
	modulePath     = "github.com/valkiriaaquatica/provider-snowflake"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		snowflakewarehouse.Configure,
		snowflakedatabase.Configure,
		snowflakeauthenticationpolicy.Configure,
		snowflakeaccountrole.Configure,
		snowflakeaccount.Configure,
		snowflakeaccountparameter.Configure,
		snowflakeapiauthenticationintegrationwithauthorizationcodegrant.Configure,
		snowflakeapiauthenticationintegrationwithclientcredentials.Configure,
		snowflakeapiauthenticationintegrationwithjwtbearer.Configure,
		snowflakedatabaserole.Configure,
		snowflakeexecute.Configure,
		snowflakeexternaloauthintegration.Configure,
		snowflakegrantaccountrole.Configure,
		snowflakegrantapplicationrole.Configure,
		snowflakegrantownership.Configure,
		snowflakegrantprivilegestoaccountrole.Configure,
		snowflakegrantprivilegestodatabaserole.Configure,
		snowflakegrantprivilegestoshare.Configure,
		snowflakelegacyserviceuser.Configure,
		snowflakemaskingpolicy.Configure,
		snowflakenetworkpolicy.Configure,
		snowflakeoauthintegrationforcustomclients.Configure,
		snowflakeoauthintegrationforpartnerapplications.Configure,
		snowflakeprimaryconnection.Configure,
		snowflakeresourcemonitor.Configure,
		snowflakerowaccesspolicy.Configure,
		snowflakesaml2integration.Configure,
		snowflakeschema.Configure,
		snowflakescimintegration.Configure,
		snowflakesecondaryconnection.Configure,
		snowflakesecondarydatabase.Configure,
		snowflakesecretwithauthorizationcodegrant.Configure,
		snowflakesecretwithbasicauthentication.Configure,
		snowflakesecretwithclientcredentials.Configure,
		snowflakesecretwithgenericstring.Configure,
		snowflakeserviceuser.Configure,
		snowflakeshareddatabase.Configure,
		snowflakestreamondirectorytable.Configure,
		snowflakestreamonexternaltable.Configure,
		snowflakestreamontable.Configure,
		snowflakestreamonview.Configure,
		snowflakestreamlit.Configure,
		snowflaketag.Configure,
		snowflaketagassociation.Configure,
		snowflaketask.Configure,
		snowflakeuser.Configure,
		snowflakeview.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
