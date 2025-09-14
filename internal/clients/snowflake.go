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

const (
	keyOrganizationName     = "organization_name"
	keyAccountName          = "account_name"
	keyUser                 = "user"
	keyPassword             = "password"
	keyHost                 = "host"
	keyRole                 = "role"
	keyAuthenticator        = "authenticator"
	keyToken                = "token"
	keyPrivateKey           = "private_key"
	keyPrivateKeyPassphrase = "private_key_passphrase"
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
		if err := populateProviderConfig(ctx, c, mg, &ps); err != nil {
			return ps, err
		}
		return ps, nil
	}
}

func populateProviderConfig(ctx context.Context, c client.Client, mg resource.Managed, ps *terraform.Setup) error {
	ref := mg.GetProviderConfigReference()
	if ref == nil {
		return errors.New(errNoProviderConfig)
	}

	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
		return errors.Wrap(err, errGetProviderConfig)
	}

	if err := trackProviderUsage(ctx, c, mg); err != nil {
		return err
	}

	creds, err := loadCredentials(ctx, c, pc)
	if err != nil {
		return err
	}

	ps.Configuration = buildConfiguration(creds)
	return nil
}

func trackProviderUsage(ctx context.Context, c client.Client, mg resource.Managed) error {
	tracker := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	return errors.Wrap(tracker.Track(ctx, mg), errTrackUsage)
}

func loadCredentials(ctx context.Context, c client.Client, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}
	return creds, nil
}

func buildConfiguration(creds map[string]string) map[string]any {
	cfg := map[string]any{}

	add := func(k string) {
		if v := creds[k]; v != "" {
			cfg[k] = v
		}
	}

	for _, k := range []string{
		keyOrganizationName, keyAccountName, keyUser, keyPassword,
	} {
		add(k)
	}

	for _, k := range []string{
		keyHost, keyRole, keyAuthenticator, keyToken,
		keyPrivateKey, keyPrivateKeyPassphrase,
	} {
		add(k)
	}

	if _, ok := cfg[keyAuthenticator]; !ok {
		if _, ok := cfg[keyToken]; ok {
			cfg[keyAuthenticator] = "PROGRAMMATIC_ACCESS_TOKEN"
		}
		if _, ok := cfg[keyPrivateKey]; ok {
			cfg[keyAuthenticator] = "SNOWFLAKE_JWT"
		}
	}

	return cfg
}
