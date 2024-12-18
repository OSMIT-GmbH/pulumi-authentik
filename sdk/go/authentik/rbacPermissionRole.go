// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentik

import (
	"context"
	"reflect"

	"errors"
	"github.com/OSMIT-GmbH/pulumi-authentik/sdk/v2024/go/authentik/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
type RbacPermissionRole struct {
	pulumi.CustomResourceState

	// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
	// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
	// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
	//   `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
	//   `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
	//   `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
	//   `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
	//   `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
	//   `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
	//   `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
	//   `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
	//   `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
	//   `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
	//   `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
	//   `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
	//   `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
	//   `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
	//   `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
	//   `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
	//   `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
	//   `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
	//   `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
	//   `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
	//   `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
	//   `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
	//   `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
	//   `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
	//   `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
	//   `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
	//   `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
	//   `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
	//   `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
	//   `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
	//   `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovider` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
	//   `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
	//   `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
	//   `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
	//   `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
	Model      pulumi.StringPtrOutput `pulumi:"model"`
	ObjectId   pulumi.StringPtrOutput `pulumi:"objectId"`
	Permission pulumi.StringOutput    `pulumi:"permission"`
	Role       pulumi.StringOutput    `pulumi:"role"`
}

// NewRbacPermissionRole registers a new resource with the given unique name, arguments, and options.
func NewRbacPermissionRole(ctx *pulumi.Context,
	name string, args *RbacPermissionRoleArgs, opts ...pulumi.ResourceOption) (*RbacPermissionRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RbacPermissionRole
	err := ctx.RegisterResource("authentik:index/rbacPermissionRole:RbacPermissionRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRbacPermissionRole gets an existing RbacPermissionRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRbacPermissionRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RbacPermissionRoleState, opts ...pulumi.ResourceOption) (*RbacPermissionRole, error) {
	var resource RbacPermissionRole
	err := ctx.ReadResource("authentik:index/rbacPermissionRole:RbacPermissionRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RbacPermissionRole resources.
type rbacPermissionRoleState struct {
	// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
	// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
	// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
	//   `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
	//   `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
	//   `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
	//   `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
	//   `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
	//   `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
	//   `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
	//   `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
	//   `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
	//   `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
	//   `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
	//   `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
	//   `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
	//   `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
	//   `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
	//   `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
	//   `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
	//   `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
	//   `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
	//   `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
	//   `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
	//   `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
	//   `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
	//   `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
	//   `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
	//   `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
	//   `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
	//   `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
	//   `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
	//   `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
	//   `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovider` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
	//   `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
	//   `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
	//   `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
	//   `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
	Model      *string `pulumi:"model"`
	ObjectId   *string `pulumi:"objectId"`
	Permission *string `pulumi:"permission"`
	Role       *string `pulumi:"role"`
}

type RbacPermissionRoleState struct {
	// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
	// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
	// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
	//   `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
	//   `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
	//   `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
	//   `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
	//   `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
	//   `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
	//   `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
	//   `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
	//   `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
	//   `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
	//   `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
	//   `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
	//   `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
	//   `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
	//   `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
	//   `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
	//   `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
	//   `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
	//   `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
	//   `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
	//   `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
	//   `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
	//   `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
	//   `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
	//   `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
	//   `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
	//   `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
	//   `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
	//   `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
	//   `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
	//   `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovider` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
	//   `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
	//   `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
	//   `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
	//   `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
	Model      pulumi.StringPtrInput
	ObjectId   pulumi.StringPtrInput
	Permission pulumi.StringPtrInput
	Role       pulumi.StringPtrInput
}

func (RbacPermissionRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPermissionRoleState)(nil)).Elem()
}

type rbacPermissionRoleArgs struct {
	// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
	// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
	// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
	//   `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
	//   `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
	//   `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
	//   `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
	//   `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
	//   `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
	//   `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
	//   `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
	//   `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
	//   `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
	//   `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
	//   `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
	//   `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
	//   `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
	//   `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
	//   `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
	//   `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
	//   `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
	//   `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
	//   `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
	//   `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
	//   `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
	//   `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
	//   `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
	//   `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
	//   `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
	//   `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
	//   `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
	//   `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
	//   `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
	//   `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovider` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
	//   `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
	//   `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
	//   `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
	//   `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
	Model      *string `pulumi:"model"`
	ObjectId   *string `pulumi:"objectId"`
	Permission string  `pulumi:"permission"`
	Role       string  `pulumi:"role"`
}

// The set of arguments for constructing a RbacPermissionRole resource.
type RbacPermissionRoleArgs struct {
	// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
	// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
	// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
	//   `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
	//   `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
	//   `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
	//   `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
	//   `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
	//   `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
	//   `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
	//   `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
	//   `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
	//   `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
	//   `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
	//   `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
	//   `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
	//   `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
	//   `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
	//   `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
	//   `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
	//   `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
	//   `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
	//   `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
	//   `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
	//   `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
	//   `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
	//   `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
	//   `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
	//   `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
	//   `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
	//   `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
	//   `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
	//   `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
	//   `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovider` -
	//   `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
	//   `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
	//   `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
	//   `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
	//   `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
	Model      pulumi.StringPtrInput
	ObjectId   pulumi.StringPtrInput
	Permission pulumi.StringInput
	Role       pulumi.StringInput
}

func (RbacPermissionRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPermissionRoleArgs)(nil)).Elem()
}

type RbacPermissionRoleInput interface {
	pulumi.Input

	ToRbacPermissionRoleOutput() RbacPermissionRoleOutput
	ToRbacPermissionRoleOutputWithContext(ctx context.Context) RbacPermissionRoleOutput
}

func (*RbacPermissionRole) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacPermissionRole)(nil)).Elem()
}

func (i *RbacPermissionRole) ToRbacPermissionRoleOutput() RbacPermissionRoleOutput {
	return i.ToRbacPermissionRoleOutputWithContext(context.Background())
}

func (i *RbacPermissionRole) ToRbacPermissionRoleOutputWithContext(ctx context.Context) RbacPermissionRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPermissionRoleOutput)
}

// RbacPermissionRoleArrayInput is an input type that accepts RbacPermissionRoleArray and RbacPermissionRoleArrayOutput values.
// You can construct a concrete instance of `RbacPermissionRoleArrayInput` via:
//
//	RbacPermissionRoleArray{ RbacPermissionRoleArgs{...} }
type RbacPermissionRoleArrayInput interface {
	pulumi.Input

	ToRbacPermissionRoleArrayOutput() RbacPermissionRoleArrayOutput
	ToRbacPermissionRoleArrayOutputWithContext(context.Context) RbacPermissionRoleArrayOutput
}

type RbacPermissionRoleArray []RbacPermissionRoleInput

func (RbacPermissionRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacPermissionRole)(nil)).Elem()
}

func (i RbacPermissionRoleArray) ToRbacPermissionRoleArrayOutput() RbacPermissionRoleArrayOutput {
	return i.ToRbacPermissionRoleArrayOutputWithContext(context.Background())
}

func (i RbacPermissionRoleArray) ToRbacPermissionRoleArrayOutputWithContext(ctx context.Context) RbacPermissionRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPermissionRoleArrayOutput)
}

// RbacPermissionRoleMapInput is an input type that accepts RbacPermissionRoleMap and RbacPermissionRoleMapOutput values.
// You can construct a concrete instance of `RbacPermissionRoleMapInput` via:
//
//	RbacPermissionRoleMap{ "key": RbacPermissionRoleArgs{...} }
type RbacPermissionRoleMapInput interface {
	pulumi.Input

	ToRbacPermissionRoleMapOutput() RbacPermissionRoleMapOutput
	ToRbacPermissionRoleMapOutputWithContext(context.Context) RbacPermissionRoleMapOutput
}

type RbacPermissionRoleMap map[string]RbacPermissionRoleInput

func (RbacPermissionRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacPermissionRole)(nil)).Elem()
}

func (i RbacPermissionRoleMap) ToRbacPermissionRoleMapOutput() RbacPermissionRoleMapOutput {
	return i.ToRbacPermissionRoleMapOutputWithContext(context.Background())
}

func (i RbacPermissionRoleMap) ToRbacPermissionRoleMapOutputWithContext(ctx context.Context) RbacPermissionRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPermissionRoleMapOutput)
}

type RbacPermissionRoleOutput struct{ *pulumi.OutputState }

func (RbacPermissionRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacPermissionRole)(nil)).Elem()
}

func (o RbacPermissionRoleOutput) ToRbacPermissionRoleOutput() RbacPermissionRoleOutput {
	return o
}

func (o RbacPermissionRoleOutput) ToRbacPermissionRoleOutputWithContext(ctx context.Context) RbacPermissionRoleOutput {
	return o
}

// Allowed values: - `authentik_tenants.domain` - `authentik_crypto.certificatekeypair` - `authentik_flows.flow` -
// `authentik_flows.flowstagebinding` - `authentik_outposts.dockerserviceconnection` -
// `authentik_outposts.kubernetesserviceconnection` - `authentik_outposts.outpost` - `authentik_policies_dummy.dummypolicy`
//   - `authentik_policies_event_matcher.eventmatcherpolicy` - `authentik_policies_expiry.passwordexpirypolicy` -
//     `authentik_policies_expression.expressionpolicy` - `authentik_policies_geoip.geoippolicy` -
//     `authentik_policies_password.passwordpolicy` - `authentik_policies_reputation.reputationpolicy` -
//     `authentik_policies.policybinding` - `authentik_providers_ldap.ldapprovider` - `authentik_providers_oauth2.scopemapping`
//   - `authentik_providers_oauth2.oauth2provider` - `authentik_providers_proxy.proxyprovider` -
//     `authentik_providers_radius.radiusprovider` - `authentik_providers_radius.radiusproviderpropertymapping` -
//     `authentik_providers_saml.samlprovider` - `authentik_providers_saml.samlpropertymapping` -
//     `authentik_providers_scim.scimprovider` - `authentik_providers_scim.scimmapping` - `authentik_rbac.role` -
//     `authentik_sources_kerberos.kerberossource` - `authentik_sources_kerberos.kerberossourcepropertymapping` -
//     `authentik_sources_kerberos.userkerberossourceconnection` - `authentik_sources_kerberos.groupkerberossourceconnection` -
//     `authentik_sources_ldap.ldapsource` - `authentik_sources_ldap.ldapsourcepropertymapping` -
//     `authentik_sources_oauth.oauthsource` - `authentik_sources_oauth.oauthsourcepropertymapping` -
//     `authentik_sources_oauth.useroauthsourceconnection` - `authentik_sources_oauth.groupoauthsourceconnection` -
//     `authentik_sources_plex.plexsource` - `authentik_sources_plex.plexsourcepropertymapping` -
//     `authentik_sources_plex.userplexsourceconnection` - `authentik_sources_plex.groupplexsourceconnection` -
//     `authentik_sources_saml.samlsource` - `authentik_sources_saml.samlsourcepropertymapping` -
//     `authentik_sources_saml.usersamlsourceconnection` - `authentik_sources_saml.groupsamlsourceconnection` -
//     `authentik_sources_scim.scimsource` - `authentik_sources_scim.scimsourcepropertymapping` -
//     `authentik_stages_authenticator_duo.authenticatorduostage` - `authentik_stages_authenticator_duo.duodevice` -
//     `authentik_stages_authenticator_sms.authenticatorsmsstage` - `authentik_stages_authenticator_sms.smsdevice` -
//     `authentik_stages_authenticator_static.authenticatorstaticstage` - `authentik_stages_authenticator_static.staticdevice`
//   - `authentik_stages_authenticator_totp.authenticatortotpstage` - `authentik_stages_authenticator_totp.totpdevice` -
//     `authentik_stages_authenticator_validate.authenticatorvalidatestage` -
//     `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage` -
//     `authentik_stages_authenticator_webauthn.webauthndevice` - `authentik_stages_captcha.captchastage` -
//     `authentik_stages_consent.consentstage` - `authentik_stages_consent.userconsent` - `authentik_stages_deny.denystage` -
//     `authentik_stages_dummy.dummystage` - `authentik_stages_email.emailstage` -
//     `authentik_stages_identification.identificationstage` - `authentik_stages_invitation.invitationstage` -
//     `authentik_stages_invitation.invitation` - `authentik_stages_password.passwordstage` - `authentik_stages_prompt.prompt`
//   - `authentik_stages_prompt.promptstage` - `authentik_stages_user_delete.userdeletestage` -
//     `authentik_stages_user_login.userloginstage` - `authentik_stages_user_logout.userlogoutstage` -
//     `authentik_stages_user_write.userwritestage` - `authentik_brands.brand` - `authentik_blueprints.blueprintinstance` -
//     `authentik_core.group` - `authentik_core.user` - `authentik_core.application` - `authentik_core.token` -
//     `authentik_enterprise.license` - `authentik_providers_google_workspace.googleworkspaceprovider` -
//     `authentik_providers_google_workspace.googleworkspaceprovidermapping` -
//     `authentik_providers_microsoft_entra.microsoftentraprovider` -
//     `authentik_providers_microsoft_entra.microsoftentraprovidermapping` - `authentik_providers_rac.racprovider` -
//     `authentik_providers_rac.endpoint` - `authentik_providers_rac.racpropertymapping` -
//     `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage` - `authentik_stages_source.sourcestage` -
//     `authentik_events.event` - `authentik_events.notificationtransport` - `authentik_events.notification` -
//     `authentik_events.notificationrule` - `authentik_events.notificationwebhookmapping`
func (o RbacPermissionRoleOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RbacPermissionRole) pulumi.StringPtrOutput { return v.Model }).(pulumi.StringPtrOutput)
}

func (o RbacPermissionRoleOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RbacPermissionRole) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o RbacPermissionRoleOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPermissionRole) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

func (o RbacPermissionRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPermissionRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type RbacPermissionRoleArrayOutput struct{ *pulumi.OutputState }

func (RbacPermissionRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacPermissionRole)(nil)).Elem()
}

func (o RbacPermissionRoleArrayOutput) ToRbacPermissionRoleArrayOutput() RbacPermissionRoleArrayOutput {
	return o
}

func (o RbacPermissionRoleArrayOutput) ToRbacPermissionRoleArrayOutputWithContext(ctx context.Context) RbacPermissionRoleArrayOutput {
	return o
}

func (o RbacPermissionRoleArrayOutput) Index(i pulumi.IntInput) RbacPermissionRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RbacPermissionRole {
		return vs[0].([]*RbacPermissionRole)[vs[1].(int)]
	}).(RbacPermissionRoleOutput)
}

type RbacPermissionRoleMapOutput struct{ *pulumi.OutputState }

func (RbacPermissionRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacPermissionRole)(nil)).Elem()
}

func (o RbacPermissionRoleMapOutput) ToRbacPermissionRoleMapOutput() RbacPermissionRoleMapOutput {
	return o
}

func (o RbacPermissionRoleMapOutput) ToRbacPermissionRoleMapOutputWithContext(ctx context.Context) RbacPermissionRoleMapOutput {
	return o
}

func (o RbacPermissionRoleMapOutput) MapIndex(k pulumi.StringInput) RbacPermissionRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RbacPermissionRole {
		return vs[0].(map[string]*RbacPermissionRole)[vs[1].(string)]
	}).(RbacPermissionRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPermissionRoleInput)(nil)).Elem(), &RbacPermissionRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPermissionRoleArrayInput)(nil)).Elem(), RbacPermissionRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPermissionRoleMapInput)(nil)).Elem(), RbacPermissionRoleMap{})
	pulumi.RegisterOutputType(RbacPermissionRoleOutput{})
	pulumi.RegisterOutputType(RbacPermissionRoleArrayOutput{})
	pulumi.RegisterOutputType(RbacPermissionRoleMapOutput{})
}
