/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/valkiriaaquatica/provider-snowflake/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal snowflake credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		ref := mg.GetProviderConfigReference()
		if ref == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		tracker := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
		if err := tracker.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		cfg := map[string]any{}
		// obligatorios
		if v := creds["organization_name"]; v != "" { cfg["organization_name"] = v }
		if v := creds["account_name"]; v != ""      { cfg["account_name"] = v }
		if v := creds["user"]; v != ""              { cfg["user"] = v }
		if v := creds["password"]; v != ""          { cfg["password"] = v }

		// opcionales (si los usas en tu Secret)
		if v := creds["host"]; v != ""              { cfg["host"] = v }
		if v := creds["role"]; v != ""              { cfg["role"] = v }
		if v := creds["authenticator"]; v != ""     { cfg["authenticator"] = v }
		if v := creds["token"]; v != ""             { cfg["token"] = v }                  // PAT
		if v := creds["private_key"]; v != ""       { cfg["private_key"] = v }            // JWT
		if v := creds["private_key_passphrase"]; v != "" { cfg["private_key_passphrase"] = v }

		// autocompletar autenticador si procede
		if cfg["authenticator"] == nil && cfg["token"] != nil {
			cfg["authenticator"] = "PROGRAMMATIC_ACCESS_TOKEN"
		}
		if cfg["authenticator"] == nil && cfg["private_key"] != nil {
			cfg["authenticator"] = "SNOWFLAKE_JWT"
		}

		ps.Configuration = cfg
		return ps, nil
	}
}
