// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-snowflake/internal/controller/providerconfig"
	account "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccount/account"
	parameter "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccountparameter/parameter"
	role "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccountrole/role"
	authenticationintegrationwithauthorizationcodegrant "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithauthorizationcodegrant/authenticationintegrationwithauthorizationcodegrant"
	authenticationintegrationwithclientcredentials "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithclientcredentials/authenticationintegrationwithclientcredentials"
	authenticationintegrationwithjwtbearer "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithjwtbearer/authenticationintegrationwithjwtbearer"
	policy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeauthenticationpolicy/policy"
	database "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakedatabase/database"
	rolesnowflakedatabaserole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakedatabaserole/role"
	execute "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeexecute/execute"
	oauthintegration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeexternaloauthintegration/oauthintegration"
	accountrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantaccountrole/accountrole"
	applicationrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantapplicationrole/applicationrole"
	ownership "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantownership/ownership"
	privilegestoaccountrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestoaccountrole/privilegestoaccountrole"
	privilegestodatabaserole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestodatabaserole/privilegestodatabaserole"
	privilegestoshare "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestoshare/privilegestoshare"
	serviceuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakelegacyserviceuser/serviceuser"
	policysnowflakemaskingpolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakemaskingpolicy/policy"
	policysnowflakenetworkpolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakenetworkpolicy/policy"
	integrationforcustomclients "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeoauthintegrationforcustomclients/integrationforcustomclients"
	integrationforpartnerapplications "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeoauthintegrationforpartnerapplications/integrationforpartnerapplications"
	connection "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeprimaryconnection/connection"
	monitor "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeresourcemonitor/monitor"
	accesspolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakerowaccesspolicy/accesspolicy"
	integration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesaml2integration/integration"
	schema "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeschema/schema"
	integrationsnowflakescimintegration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakescimintegration/integration"
	connectionsnowflakesecondaryconnection "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecondaryconnection/connection"
	databasesnowflakesecondarydatabase "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecondarydatabase/database"
	withauthorizationcodegrant "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithauthorizationcodegrant/withauthorizationcodegrant"
	withclientcredentials "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithclientcredentials/withclientcredentials"
	withgenericstring "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithgenericstring/withgenericstring"
	user "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeserviceuser/user"
	databasesnowflakeshareddatabase "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeshareddatabase/database"
	streamlit "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamlit/streamlit"
	ondirectorytable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamondirectorytable/ondirectorytable"
	onexternaltable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamonexternaltable/onexternaltable"
	ontable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamontable/ontable"
	onview "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamonview/onview"
	tag "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketag/tag"
	association "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketagassociation/association"
	task "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketask/task"
	usersnowflakeuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeuser/user"
	view "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeview/view"
	warehouse "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakewarehouse/warehouse"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		account.Setup,
		parameter.Setup,
		role.Setup,
		authenticationintegrationwithauthorizationcodegrant.Setup,
		authenticationintegrationwithclientcredentials.Setup,
		authenticationintegrationwithjwtbearer.Setup,
		policy.Setup,
		database.Setup,
		rolesnowflakedatabaserole.Setup,
		execute.Setup,
		oauthintegration.Setup,
		accountrole.Setup,
		applicationrole.Setup,
		ownership.Setup,
		privilegestoaccountrole.Setup,
		privilegestodatabaserole.Setup,
		privilegestoshare.Setup,
		serviceuser.Setup,
		policysnowflakemaskingpolicy.Setup,
		policysnowflakenetworkpolicy.Setup,
		integrationforcustomclients.Setup,
		integrationforpartnerapplications.Setup,
		connection.Setup,
		monitor.Setup,
		accesspolicy.Setup,
		integration.Setup,
		schema.Setup,
		integrationsnowflakescimintegration.Setup,
		connectionsnowflakesecondaryconnection.Setup,
		databasesnowflakesecondarydatabase.Setup,
		withauthorizationcodegrant.Setup,
		withclientcredentials.Setup,
		withgenericstring.Setup,
		user.Setup,
		databasesnowflakeshareddatabase.Setup,
		streamlit.Setup,
		ondirectorytable.Setup,
		onexternaltable.Setup,
		ontable.Setup,
		onview.Setup,
		tag.Setup,
		association.Setup,
		task.Setup,
		usersnowflakeuser.Setup,
		view.Setup,
		warehouse.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
