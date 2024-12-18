# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .application import *
from .blueprint import *
from .brand import *
from .certificate_key_pair import *
from .enterprise_license import *
from .event_rule import *
from .event_transport import *
from .flow import *
from .flow_stage_binding import *
from .get_brand import *
from .get_certificate_key_pair import *
from .get_flow import *
from .get_group import *
from .get_groups import *
from .get_property_mapping_provider_rac import *
from .get_property_mapping_provider_radius import *
from .get_property_mapping_provider_saml import *
from .get_property_mapping_provider_scim import *
from .get_property_mapping_provider_scope import *
from .get_property_mapping_source_ldap import *
from .get_provider_oauth2_config import *
from .get_provider_saml_metadata import *
from .get_source import *
from .get_stage import *
from .get_user import *
from .get_users import *
from .get_webauthn_device_type import *
from .group import *
from .outpost import *
from .policy_binding import *
from .policy_dummy import *
from .policy_event_matcher import *
from .policy_expiry import *
from .policy_expression import *
from .policy_geoip import *
from .policy_password import *
from .policy_reputation import *
from .property_mapping_google_workspace import *
from .property_mapping_ldap import *
from .property_mapping_microsoft_entra import *
from .property_mapping_notification import *
from .property_mapping_provider_google_workspace import *
from .property_mapping_provider_microsoft_entra import *
from .property_mapping_provider_rac import *
from .property_mapping_provider_radius import *
from .property_mapping_provider_saml import *
from .property_mapping_provider_scim import *
from .property_mapping_provider_scope import *
from .property_mapping_rac import *
from .property_mapping_radius import *
from .property_mapping_saml import *
from .property_mapping_scim import *
from .property_mapping_source_kerberos import *
from .property_mapping_source_ldap import *
from .property_mapping_source_oauth import *
from .property_mapping_source_plex import *
from .property_mapping_source_saml import *
from .property_mapping_source_scim import *
from .provider import *
from .provider_google_workspace import *
from .provider_ldap import *
from .provider_microsoft_entra import *
from .provider_oauth2 import *
from .provider_proxy import *
from .provider_rac import *
from .provider_radius import *
from .provider_saml import *
from .provider_scim import *
from .rac_endpoint import *
from .rbac_permission_role import *
from .rbac_permission_user import *
from .rbac_role import *
from .scope_mapping import *
from .service_connection_docker import *
from .service_connection_kubernetes import *
from .source_kerberos import *
from .source_ldap import *
from .source_oauth import *
from .source_plex import *
from .source_saml import *
from .stage_authenticator_duo import *
from .stage_authenticator_endpoint_gdtc import *
from .stage_authenticator_sms import *
from .stage_authenticator_static import *
from .stage_authenticator_totp import *
from .stage_authenticator_validate import *
from .stage_authenticator_webauthn import *
from .stage_captcha import *
from .stage_consent import *
from .stage_deny import *
from .stage_dummy import *
from .stage_email import *
from .stage_identification import *
from .stage_invitation import *
from .stage_password import *
from .stage_prompt import *
from .stage_prompt_field import *
from .stage_source import *
from .stage_user_delete import *
from .stage_user_login import *
from .stage_user_logout import *
from .stage_user_write import *
from .system_settings import *
from .token import *
from .user import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import osmit_gmbh_pulumi_authentik.config as __config
    config = __config
else:
    config = _utilities.lazy_import('osmit_gmbh_pulumi_authentik.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "authentik",
  "mod": "index/application",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/application:Application": "Application"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/blueprint",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/blueprint:Blueprint": "Blueprint"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/brand",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/brand:Brand": "Brand"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/certificateKeyPair",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/certificateKeyPair:CertificateKeyPair": "CertificateKeyPair"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/enterpriseLicense",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/enterpriseLicense:EnterpriseLicense": "EnterpriseLicense"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/eventRule",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/eventRule:EventRule": "EventRule"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/eventTransport",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/eventTransport:EventTransport": "EventTransport"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/flow",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/flow:Flow": "Flow"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/flowStageBinding",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/flowStageBinding:FlowStageBinding": "FlowStageBinding"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/group",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/group:Group": "Group"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/outpost",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/outpost:Outpost": "Outpost"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyBinding",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyBinding:PolicyBinding": "PolicyBinding"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyDummy",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyDummy:PolicyDummy": "PolicyDummy"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyEventMatcher",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyEventMatcher:PolicyEventMatcher": "PolicyEventMatcher"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyExpiry",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyExpiry:PolicyExpiry": "PolicyExpiry"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyExpression",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyExpression:PolicyExpression": "PolicyExpression"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyGeoip",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyGeoip:PolicyGeoip": "PolicyGeoip"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyPassword",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyPassword:PolicyPassword": "PolicyPassword"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/policyReputation",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/policyReputation:PolicyReputation": "PolicyReputation"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingGoogleWorkspace",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingGoogleWorkspace:PropertyMappingGoogleWorkspace": "PropertyMappingGoogleWorkspace"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingLdap",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingLdap:PropertyMappingLdap": "PropertyMappingLdap"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingMicrosoftEntra",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingMicrosoftEntra:PropertyMappingMicrosoftEntra": "PropertyMappingMicrosoftEntra"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingNotification",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingNotification:PropertyMappingNotification": "PropertyMappingNotification"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderGoogleWorkspace",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderGoogleWorkspace:PropertyMappingProviderGoogleWorkspace": "PropertyMappingProviderGoogleWorkspace"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderMicrosoftEntra",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderMicrosoftEntra:PropertyMappingProviderMicrosoftEntra": "PropertyMappingProviderMicrosoftEntra"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderRac",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderRac:PropertyMappingProviderRac": "PropertyMappingProviderRac"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderRadius",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderRadius:PropertyMappingProviderRadius": "PropertyMappingProviderRadius"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderSaml",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderSaml:PropertyMappingProviderSaml": "PropertyMappingProviderSaml"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderScim",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderScim:PropertyMappingProviderScim": "PropertyMappingProviderScim"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingProviderScope",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingProviderScope:PropertyMappingProviderScope": "PropertyMappingProviderScope"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingRac",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingRac:PropertyMappingRac": "PropertyMappingRac"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingRadius",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingRadius:PropertyMappingRadius": "PropertyMappingRadius"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSaml",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSaml:PropertyMappingSaml": "PropertyMappingSaml"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingScim",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingScim:PropertyMappingScim": "PropertyMappingScim"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourceKerberos",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourceKerberos:PropertyMappingSourceKerberos": "PropertyMappingSourceKerberos"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourceLdap",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourceLdap:PropertyMappingSourceLdap": "PropertyMappingSourceLdap"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourceOauth",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourceOauth:PropertyMappingSourceOauth": "PropertyMappingSourceOauth"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourcePlex",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourcePlex:PropertyMappingSourcePlex": "PropertyMappingSourcePlex"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourceSaml",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourceSaml:PropertyMappingSourceSaml": "PropertyMappingSourceSaml"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/propertyMappingSourceScim",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/propertyMappingSourceScim:PropertyMappingSourceScim": "PropertyMappingSourceScim"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerGoogleWorkspace",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerGoogleWorkspace:ProviderGoogleWorkspace": "ProviderGoogleWorkspace"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerLdap",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerLdap:ProviderLdap": "ProviderLdap"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerMicrosoftEntra",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerMicrosoftEntra:ProviderMicrosoftEntra": "ProviderMicrosoftEntra"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerOauth2",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerOauth2:ProviderOauth2": "ProviderOauth2"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerProxy",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerProxy:ProviderProxy": "ProviderProxy"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerRac",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerRac:ProviderRac": "ProviderRac"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerRadius",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerRadius:ProviderRadius": "ProviderRadius"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerSaml",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerSaml:ProviderSaml": "ProviderSaml"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/providerScim",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/providerScim:ProviderScim": "ProviderScim"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/racEndpoint",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/racEndpoint:RacEndpoint": "RacEndpoint"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/rbacPermissionRole",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/rbacPermissionRole:RbacPermissionRole": "RbacPermissionRole"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/rbacPermissionUser",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/rbacPermissionUser:RbacPermissionUser": "RbacPermissionUser"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/rbacRole",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/rbacRole:RbacRole": "RbacRole"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/scopeMapping",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/scopeMapping:ScopeMapping": "ScopeMapping"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/serviceConnectionDocker",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/serviceConnectionDocker:ServiceConnectionDocker": "ServiceConnectionDocker"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/serviceConnectionKubernetes",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/serviceConnectionKubernetes:ServiceConnectionKubernetes": "ServiceConnectionKubernetes"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/sourceKerberos",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/sourceKerberos:SourceKerberos": "SourceKerberos"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/sourceLdap",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/sourceLdap:SourceLdap": "SourceLdap"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/sourceOauth",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/sourceOauth:SourceOauth": "SourceOauth"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/sourcePlex",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/sourcePlex:SourcePlex": "SourcePlex"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/sourceSaml",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/sourceSaml:SourceSaml": "SourceSaml"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorDuo",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorDuo:StageAuthenticatorDuo": "StageAuthenticatorDuo"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorEndpointGdtc",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorEndpointGdtc:StageAuthenticatorEndpointGdtc": "StageAuthenticatorEndpointGdtc"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorSms",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorSms:StageAuthenticatorSms": "StageAuthenticatorSms"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorStatic",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorStatic:StageAuthenticatorStatic": "StageAuthenticatorStatic"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorTotp",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorTotp:StageAuthenticatorTotp": "StageAuthenticatorTotp"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorValidate",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorValidate:StageAuthenticatorValidate": "StageAuthenticatorValidate"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageAuthenticatorWebauthn",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn": "StageAuthenticatorWebauthn"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageCaptcha",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageCaptcha:StageCaptcha": "StageCaptcha"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageConsent",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageConsent:StageConsent": "StageConsent"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageDeny",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageDeny:StageDeny": "StageDeny"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageDummy",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageDummy:StageDummy": "StageDummy"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageEmail",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageEmail:StageEmail": "StageEmail"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageIdentification",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageIdentification:StageIdentification": "StageIdentification"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageInvitation",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageInvitation:StageInvitation": "StageInvitation"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stagePassword",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stagePassword:StagePassword": "StagePassword"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stagePrompt",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stagePrompt:StagePrompt": "StagePrompt"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stagePromptField",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stagePromptField:StagePromptField": "StagePromptField"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageSource",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageSource:StageSource": "StageSource"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageUserDelete",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageUserDelete:StageUserDelete": "StageUserDelete"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageUserLogin",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageUserLogin:StageUserLogin": "StageUserLogin"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageUserLogout",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageUserLogout:StageUserLogout": "StageUserLogout"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/stageUserWrite",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/stageUserWrite:StageUserWrite": "StageUserWrite"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/systemSettings",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/systemSettings:SystemSettings": "SystemSettings"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/token",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/token:Token": "Token"
  }
 },
 {
  "pkg": "authentik",
  "mod": "index/user",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "classes": {
   "authentik:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "authentik",
  "token": "pulumi:providers:authentik",
  "fqn": "osmit_gmbh_pulumi_authentik",
  "class": "Provider"
 }
]
"""
)
