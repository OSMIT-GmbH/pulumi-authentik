// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"fmt"

	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/go/authentik/internal"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "authentik:index/application:Application":
		r = &Application{}
	case "authentik:index/blueprint:Blueprint":
		r = &Blueprint{}
	case "authentik:index/brand:Brand":
		r = &Brand{}
	case "authentik:index/certificateKeyPair:CertificateKeyPair":
		r = &CertificateKeyPair{}
	case "authentik:index/enterpriseLicense:EnterpriseLicense":
		r = &EnterpriseLicense{}
	case "authentik:index/eventRule:EventRule":
		r = &EventRule{}
	case "authentik:index/eventTransport:EventTransport":
		r = &EventTransport{}
	case "authentik:index/flow:Flow":
		r = &Flow{}
	case "authentik:index/flowStageBinding:FlowStageBinding":
		r = &FlowStageBinding{}
	case "authentik:index/group:Group":
		r = &Group{}
	case "authentik:index/outpost:Outpost":
		r = &Outpost{}
	case "authentik:index/policyBinding:PolicyBinding":
		r = &PolicyBinding{}
	case "authentik:index/policyDummy:PolicyDummy":
		r = &PolicyDummy{}
	case "authentik:index/policyEventMatcher:PolicyEventMatcher":
		r = &PolicyEventMatcher{}
	case "authentik:index/policyExpiry:PolicyExpiry":
		r = &PolicyExpiry{}
	case "authentik:index/policyExpression:PolicyExpression":
		r = &PolicyExpression{}
	case "authentik:index/policyGeoip:PolicyGeoip":
		r = &PolicyGeoip{}
	case "authentik:index/policyPassword:PolicyPassword":
		r = &PolicyPassword{}
	case "authentik:index/policyReputation:PolicyReputation":
		r = &PolicyReputation{}
	case "authentik:index/propertyMappingGoogleWorkspace:PropertyMappingGoogleWorkspace":
		r = &PropertyMappingGoogleWorkspace{}
	case "authentik:index/propertyMappingLdap:PropertyMappingLdap":
		r = &PropertyMappingLdap{}
	case "authentik:index/propertyMappingMicrosoftEntra:PropertyMappingMicrosoftEntra":
		r = &PropertyMappingMicrosoftEntra{}
	case "authentik:index/propertyMappingNotification:PropertyMappingNotification":
		r = &PropertyMappingNotification{}
	case "authentik:index/propertyMappingProviderGoogleWorkspace:PropertyMappingProviderGoogleWorkspace":
		r = &PropertyMappingProviderGoogleWorkspace{}
	case "authentik:index/propertyMappingProviderMicrosoftEntra:PropertyMappingProviderMicrosoftEntra":
		r = &PropertyMappingProviderMicrosoftEntra{}
	case "authentik:index/propertyMappingProviderRac:PropertyMappingProviderRac":
		r = &PropertyMappingProviderRac{}
	case "authentik:index/propertyMappingProviderRadius:PropertyMappingProviderRadius":
		r = &PropertyMappingProviderRadius{}
	case "authentik:index/propertyMappingProviderSaml:PropertyMappingProviderSaml":
		r = &PropertyMappingProviderSaml{}
	case "authentik:index/propertyMappingProviderScim:PropertyMappingProviderScim":
		r = &PropertyMappingProviderScim{}
	case "authentik:index/propertyMappingProviderScope:PropertyMappingProviderScope":
		r = &PropertyMappingProviderScope{}
	case "authentik:index/propertyMappingRac:PropertyMappingRac":
		r = &PropertyMappingRac{}
	case "authentik:index/propertyMappingRadius:PropertyMappingRadius":
		r = &PropertyMappingRadius{}
	case "authentik:index/propertyMappingSaml:PropertyMappingSaml":
		r = &PropertyMappingSaml{}
	case "authentik:index/propertyMappingScim:PropertyMappingScim":
		r = &PropertyMappingScim{}
	case "authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos":
		r = &PropertyMappingSourceKerberos{}
	case "authentik:index/propertyMappingSourceLdap:PropertyMappingSourceLdap":
		r = &PropertyMappingSourceLdap{}
	case "authentik:index/propertyMappingSourceOauth:PropertyMappingSourceOauth":
		r = &PropertyMappingSourceOauth{}
	case "authentik:index/propertyMappingSourcePlex:PropertyMappingSourcePlex":
		r = &PropertyMappingSourcePlex{}
	case "authentik:index/propertyMappingSourceSaml:PropertyMappingSourceSaml":
		r = &PropertyMappingSourceSaml{}
	case "authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim":
		r = &PropertyMappingSourceScim{}
	case "authentik:index/providerGoogleWorkspace:ProviderGoogleWorkspace":
		r = &ProviderGoogleWorkspace{}
	case "authentik:index/providerLdap:ProviderLdap":
		r = &ProviderLdap{}
	case "authentik:index/providerMicrosoftEntra:ProviderMicrosoftEntra":
		r = &ProviderMicrosoftEntra{}
	case "authentik:index/providerOauth2:ProviderOauth2":
		r = &ProviderOauth2{}
	case "authentik:index/providerProxy:ProviderProxy":
		r = &ProviderProxy{}
	case "authentik:index/providerRac:ProviderRac":
		r = &ProviderRac{}
	case "authentik:index/providerRadius:ProviderRadius":
		r = &ProviderRadius{}
	case "authentik:index/providerSaml:ProviderSaml":
		r = &ProviderSaml{}
	case "authentik:index/providerScim:ProviderScim":
		r = &ProviderScim{}
	case "authentik:index/racEndpoint:RacEndpoint":
		r = &RacEndpoint{}
	case "authentik:index/rbacPermissionRole:RbacPermissionRole":
		r = &RbacPermissionRole{}
	case "authentik:index/rbacPermissionUser:RbacPermissionUser":
		r = &RbacPermissionUser{}
	case "authentik:index/rbacRole:RbacRole":
		r = &RbacRole{}
	case "authentik:index/scopeMapping:ScopeMapping":
		r = &ScopeMapping{}
	case "authentik:index/serviceConnectionDocker:ServiceConnectionDocker":
		r = &ServiceConnectionDocker{}
	case "authentik:index/serviceConnectionKubernetes:ServiceConnectionKubernetes":
		r = &ServiceConnectionKubernetes{}
	case "authentik:index/sourceKerberos:SourceKerberos":
		r = &SourceKerberos{}
	case "authentik:index/sourceLdap:SourceLdap":
		r = &SourceLdap{}
	case "authentik:index/sourceOauth:SourceOauth":
		r = &SourceOauth{}
	case "authentik:index/sourcePlex:SourcePlex":
		r = &SourcePlex{}
	case "authentik:index/sourceSaml:SourceSaml":
		r = &SourceSaml{}
	case "authentik:index/stageAuthenticatorDuo:StageAuthenticatorDuo":
		r = &StageAuthenticatorDuo{}
	case "authentik:index/stageAuthenticatorEndpointGdtc:StageAuthenticatorEndpointGdtc":
		r = &StageAuthenticatorEndpointGdtc{}
	case "authentik:index/stageAuthenticatorSms:StageAuthenticatorSms":
		r = &StageAuthenticatorSms{}
	case "authentik:index/stageAuthenticatorStatic:StageAuthenticatorStatic":
		r = &StageAuthenticatorStatic{}
	case "authentik:index/stageAuthenticatorTotp:StageAuthenticatorTotp":
		r = &StageAuthenticatorTotp{}
	case "authentik:index/stageAuthenticatorValidate:StageAuthenticatorValidate":
		r = &StageAuthenticatorValidate{}
	case "authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn":
		r = &StageAuthenticatorWebauthn{}
	case "authentik:index/stageCaptcha:StageCaptcha":
		r = &StageCaptcha{}
	case "authentik:index/stageConsent:StageConsent":
		r = &StageConsent{}
	case "authentik:index/stageDeny:StageDeny":
		r = &StageDeny{}
	case "authentik:index/stageDummy:StageDummy":
		r = &StageDummy{}
	case "authentik:index/stageEmail:StageEmail":
		r = &StageEmail{}
	case "authentik:index/stageIdentification:StageIdentification":
		r = &StageIdentification{}
	case "authentik:index/stageInvitation:StageInvitation":
		r = &StageInvitation{}
	case "authentik:index/stagePassword:StagePassword":
		r = &StagePassword{}
	case "authentik:index/stagePrompt:StagePrompt":
		r = &StagePrompt{}
	case "authentik:index/stagePromptField:StagePromptField":
		r = &StagePromptField{}
	case "authentik:index/stageSource:StageSource":
		r = &StageSource{}
	case "authentik:index/stageUserDelete:StageUserDelete":
		r = &StageUserDelete{}
	case "authentik:index/stageUserLogin:StageUserLogin":
		r = &StageUserLogin{}
	case "authentik:index/stageUserLogout:StageUserLogout":
		r = &StageUserLogout{}
	case "authentik:index/stageUserWrite:StageUserWrite":
		r = &StageUserWrite{}
	case "authentik:index/systemSettings:SystemSettings":
		r = &SystemSettings{}
	case "authentik:index/token:Token":
		r = &Token{}
	case "authentik:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:authentik" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"authentik",
		"index/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/blueprint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/brand",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/certificateKeyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/enterpriseLicense",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/eventRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/eventTransport",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/flow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/flowStageBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/outpost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyDummy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyEventMatcher",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyExpiry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyExpression",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyGeoip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/policyReputation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingGoogleWorkspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingLdap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingMicrosoftEntra",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingNotification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderGoogleWorkspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderMicrosoftEntra",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderRac",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderRadius",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderSaml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderScim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingProviderScope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingRac",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingRadius",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSaml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingScim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourceKerberos",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourceLdap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourceOauth",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourcePlex",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourceSaml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/propertyMappingSourceScim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerGoogleWorkspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerLdap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerMicrosoftEntra",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerOauth2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerProxy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerRac",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerRadius",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerSaml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/providerScim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/racEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/rbacPermissionRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/rbacPermissionUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/rbacRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/scopeMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/serviceConnectionDocker",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/serviceConnectionKubernetes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/sourceKerberos",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/sourceLdap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/sourceOauth",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/sourcePlex",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/sourceSaml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorDuo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorEndpointGdtc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorSms",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorTotp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorValidate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageAuthenticatorWebauthn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageCaptcha",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageConsent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageDeny",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageDummy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageIdentification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageInvitation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stagePassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stagePrompt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stagePromptField",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageUserDelete",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageUserLogin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageUserLogout",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/stageUserWrite",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/systemSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/token",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"authentik",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"authentik",
		&pkg{version},
	)
}
