# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['ProviderOauth2Args', 'ProviderOauth2']

@pulumi.input_type
class ProviderOauth2Args:
    def __init__(__self__, *,
                 authorization_flow: pulumi.Input[str],
                 client_id: pulumi.Input[str],
                 invalidation_flow: pulumi.Input[str],
                 access_code_validity: Optional[pulumi.Input[str]] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 allowed_redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_type: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[str]] = None,
                 include_claims_in_id_token: Optional[pulumi.Input[bool]] = None,
                 issuer_mode: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 signing_key: Optional[pulumi.Input[str]] = None,
                 sub_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProviderOauth2 resource.
        :param pulumi.Input[str] client_type: Allowed values: - `confidential` - `public`
        :param pulumi.Input[str] issuer_mode: Allowed values: - `global` - `per_provider`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] sub_mode: Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        pulumi.set(__self__, "authorization_flow", authorization_flow)
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "invalidation_flow", invalidation_flow)
        if access_code_validity is not None:
            pulumi.set(__self__, "access_code_validity", access_code_validity)
        if access_token_validity is not None:
            pulumi.set(__self__, "access_token_validity", access_token_validity)
        if allowed_redirect_uris is not None:
            pulumi.set(__self__, "allowed_redirect_uris", allowed_redirect_uris)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_type is not None:
            pulumi.set(__self__, "client_type", client_type)
        if encryption_key is not None:
            pulumi.set(__self__, "encryption_key", encryption_key)
        if include_claims_in_id_token is not None:
            pulumi.set(__self__, "include_claims_in_id_token", include_claims_in_id_token)
        if issuer_mode is not None:
            pulumi.set(__self__, "issuer_mode", issuer_mode)
        if jwks_sources is not None:
            pulumi.set(__self__, "jwks_sources", jwks_sources)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if property_mappings is not None:
            pulumi.set(__self__, "property_mappings", property_mappings)
        if refresh_token_validity is not None:
            pulumi.set(__self__, "refresh_token_validity", refresh_token_validity)
        if signing_key is not None:
            pulumi.set(__self__, "signing_key", signing_key)
        if sub_mode is not None:
            pulumi.set(__self__, "sub_mode", sub_mode)

    @property
    @pulumi.getter(name="authorizationFlow")
    def authorization_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authorization_flow")

    @authorization_flow.setter
    def authorization_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_flow", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "invalidation_flow")

    @invalidation_flow.setter
    def invalidation_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "invalidation_flow", value)

    @property
    @pulumi.getter(name="accessCodeValidity")
    def access_code_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_code_validity")

    @access_code_validity.setter
    def access_code_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_code_validity", value)

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_token_validity")

    @access_token_validity.setter
    def access_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_validity", value)

    @property
    @pulumi.getter(name="allowedRedirectUris")
    def allowed_redirect_uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        return pulumi.get(self, "allowed_redirect_uris")

    @allowed_redirect_uris.setter
    def allowed_redirect_uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "allowed_redirect_uris", value)

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authentication_flow")

    @authentication_flow.setter
    def authentication_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_flow", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="clientType")
    def client_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `confidential` - `public`
        """
        return pulumi.get(self, "client_type")

    @client_type.setter
    def client_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_type", value)

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encryption_key")

    @encryption_key.setter
    def encryption_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_key", value)

    @property
    @pulumi.getter(name="includeClaimsInIdToken")
    def include_claims_in_id_token(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_claims_in_id_token")

    @include_claims_in_id_token.setter
    def include_claims_in_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_claims_in_id_token", value)

    @property
    @pulumi.getter(name="issuerMode")
    def issuer_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `global` - `per_provider`
        """
        return pulumi.get(self, "issuer_mode")

    @issuer_mode.setter
    def issuer_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_mode", value)

    @property
    @pulumi.getter(name="jwksSources")
    def jwks_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        """
        return pulumi.get(self, "jwks_sources")

    @jwks_sources.setter
    def jwks_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "jwks_sources", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="propertyMappings")
    def property_mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_mappings")

    @property_mappings.setter
    def property_mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_mappings", value)

    @property
    @pulumi.getter(name="refreshTokenValidity")
    def refresh_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "refresh_token_validity")

    @refresh_token_validity.setter
    def refresh_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "refresh_token_validity", value)

    @property
    @pulumi.getter(name="signingKey")
    def signing_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "signing_key")

    @signing_key.setter
    def signing_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_key", value)

    @property
    @pulumi.getter(name="subMode")
    def sub_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        return pulumi.get(self, "sub_mode")

    @sub_mode.setter
    def sub_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_mode", value)


@pulumi.input_type
class _ProviderOauth2State:
    def __init__(__self__, *,
                 access_code_validity: Optional[pulumi.Input[str]] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 allowed_redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_type: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[str]] = None,
                 include_claims_in_id_token: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 issuer_mode: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 signing_key: Optional[pulumi.Input[str]] = None,
                 sub_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProviderOauth2 resources.
        :param pulumi.Input[str] client_type: Allowed values: - `confidential` - `public`
        :param pulumi.Input[str] issuer_mode: Allowed values: - `global` - `per_provider`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] sub_mode: Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        if access_code_validity is not None:
            pulumi.set(__self__, "access_code_validity", access_code_validity)
        if access_token_validity is not None:
            pulumi.set(__self__, "access_token_validity", access_token_validity)
        if allowed_redirect_uris is not None:
            pulumi.set(__self__, "allowed_redirect_uris", allowed_redirect_uris)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if authorization_flow is not None:
            pulumi.set(__self__, "authorization_flow", authorization_flow)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_type is not None:
            pulumi.set(__self__, "client_type", client_type)
        if encryption_key is not None:
            pulumi.set(__self__, "encryption_key", encryption_key)
        if include_claims_in_id_token is not None:
            pulumi.set(__self__, "include_claims_in_id_token", include_claims_in_id_token)
        if invalidation_flow is not None:
            pulumi.set(__self__, "invalidation_flow", invalidation_flow)
        if issuer_mode is not None:
            pulumi.set(__self__, "issuer_mode", issuer_mode)
        if jwks_sources is not None:
            pulumi.set(__self__, "jwks_sources", jwks_sources)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if property_mappings is not None:
            pulumi.set(__self__, "property_mappings", property_mappings)
        if refresh_token_validity is not None:
            pulumi.set(__self__, "refresh_token_validity", refresh_token_validity)
        if signing_key is not None:
            pulumi.set(__self__, "signing_key", signing_key)
        if sub_mode is not None:
            pulumi.set(__self__, "sub_mode", sub_mode)

    @property
    @pulumi.getter(name="accessCodeValidity")
    def access_code_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_code_validity")

    @access_code_validity.setter
    def access_code_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_code_validity", value)

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_token_validity")

    @access_token_validity.setter
    def access_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_validity", value)

    @property
    @pulumi.getter(name="allowedRedirectUris")
    def allowed_redirect_uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        return pulumi.get(self, "allowed_redirect_uris")

    @allowed_redirect_uris.setter
    def allowed_redirect_uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "allowed_redirect_uris", value)

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authentication_flow")

    @authentication_flow.setter
    def authentication_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_flow", value)

    @property
    @pulumi.getter(name="authorizationFlow")
    def authorization_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorization_flow")

    @authorization_flow.setter
    def authorization_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_flow", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="clientType")
    def client_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `confidential` - `public`
        """
        return pulumi.get(self, "client_type")

    @client_type.setter
    def client_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_type", value)

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encryption_key")

    @encryption_key.setter
    def encryption_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_key", value)

    @property
    @pulumi.getter(name="includeClaimsInIdToken")
    def include_claims_in_id_token(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_claims_in_id_token")

    @include_claims_in_id_token.setter
    def include_claims_in_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_claims_in_id_token", value)

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "invalidation_flow")

    @invalidation_flow.setter
    def invalidation_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invalidation_flow", value)

    @property
    @pulumi.getter(name="issuerMode")
    def issuer_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `global` - `per_provider`
        """
        return pulumi.get(self, "issuer_mode")

    @issuer_mode.setter
    def issuer_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_mode", value)

    @property
    @pulumi.getter(name="jwksSources")
    def jwks_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        """
        return pulumi.get(self, "jwks_sources")

    @jwks_sources.setter
    def jwks_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "jwks_sources", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="propertyMappings")
    def property_mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_mappings")

    @property_mappings.setter
    def property_mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_mappings", value)

    @property
    @pulumi.getter(name="refreshTokenValidity")
    def refresh_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "refresh_token_validity")

    @refresh_token_validity.setter
    def refresh_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "refresh_token_validity", value)

    @property
    @pulumi.getter(name="signingKey")
    def signing_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "signing_key")

    @signing_key.setter
    def signing_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_key", value)

    @property
    @pulumi.getter(name="subMode")
    def sub_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        return pulumi.get(self, "sub_mode")

    @sub_mode.setter
    def sub_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_mode", value)


class ProviderOauth2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_code_validity: Optional[pulumi.Input[str]] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 allowed_redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_type: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[str]] = None,
                 include_claims_in_id_token: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 issuer_mode: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 signing_key: Optional[pulumi.Input[str]] = None,
                 sub_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create an OAuth2 Provider
        name = authentik.ProviderOauth2("name",
            name="grafana",
            client_id="grafana",
            allowed_redirect_uris=[{
                "matching_mode": "strict",
                "url": "http://localhost",
            }])
        name_application = authentik.Application("name",
            name="test app",
            slug="test-app",
            protocol_provider=name.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_type: Allowed values: - `confidential` - `public`
        :param pulumi.Input[str] issuer_mode: Allowed values: - `global` - `per_provider`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] sub_mode: Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderOauth2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create an OAuth2 Provider
        name = authentik.ProviderOauth2("name",
            name="grafana",
            client_id="grafana",
            allowed_redirect_uris=[{
                "matching_mode": "strict",
                "url": "http://localhost",
            }])
        name_application = authentik.Application("name",
            name="test app",
            slug="test-app",
            protocol_provider=name.id)
        ```

        :param str resource_name: The name of the resource.
        :param ProviderOauth2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderOauth2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_code_validity: Optional[pulumi.Input[str]] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 allowed_redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_type: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[str]] = None,
                 include_claims_in_id_token: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 issuer_mode: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 signing_key: Optional[pulumi.Input[str]] = None,
                 sub_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderOauth2Args.__new__(ProviderOauth2Args)

            __props__.__dict__["access_code_validity"] = access_code_validity
            __props__.__dict__["access_token_validity"] = access_token_validity
            __props__.__dict__["allowed_redirect_uris"] = allowed_redirect_uris
            __props__.__dict__["authentication_flow"] = authentication_flow
            if authorization_flow is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_flow'")
            __props__.__dict__["authorization_flow"] = authorization_flow
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = None if client_secret is None else pulumi.Output.secret(client_secret)
            __props__.__dict__["client_type"] = client_type
            __props__.__dict__["encryption_key"] = encryption_key
            __props__.__dict__["include_claims_in_id_token"] = include_claims_in_id_token
            if invalidation_flow is None and not opts.urn:
                raise TypeError("Missing required property 'invalidation_flow'")
            __props__.__dict__["invalidation_flow"] = invalidation_flow
            __props__.__dict__["issuer_mode"] = issuer_mode
            __props__.__dict__["jwks_sources"] = jwks_sources
            __props__.__dict__["name"] = name
            __props__.__dict__["property_mappings"] = property_mappings
            __props__.__dict__["refresh_token_validity"] = refresh_token_validity
            __props__.__dict__["signing_key"] = signing_key
            __props__.__dict__["sub_mode"] = sub_mode
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ProviderOauth2, __self__).__init__(
            'authentik:index/providerOauth2:ProviderOauth2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_code_validity: Optional[pulumi.Input[str]] = None,
            access_token_validity: Optional[pulumi.Input[str]] = None,
            allowed_redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
            authentication_flow: Optional[pulumi.Input[str]] = None,
            authorization_flow: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            client_type: Optional[pulumi.Input[str]] = None,
            encryption_key: Optional[pulumi.Input[str]] = None,
            include_claims_in_id_token: Optional[pulumi.Input[bool]] = None,
            invalidation_flow: Optional[pulumi.Input[str]] = None,
            issuer_mode: Optional[pulumi.Input[str]] = None,
            jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            refresh_token_validity: Optional[pulumi.Input[str]] = None,
            signing_key: Optional[pulumi.Input[str]] = None,
            sub_mode: Optional[pulumi.Input[str]] = None) -> 'ProviderOauth2':
        """
        Get an existing ProviderOauth2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_type: Allowed values: - `confidential` - `public`
        :param pulumi.Input[str] issuer_mode: Allowed values: - `global` - `per_provider`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] sub_mode: Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProviderOauth2State.__new__(_ProviderOauth2State)

        __props__.__dict__["access_code_validity"] = access_code_validity
        __props__.__dict__["access_token_validity"] = access_token_validity
        __props__.__dict__["allowed_redirect_uris"] = allowed_redirect_uris
        __props__.__dict__["authentication_flow"] = authentication_flow
        __props__.__dict__["authorization_flow"] = authorization_flow
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["client_type"] = client_type
        __props__.__dict__["encryption_key"] = encryption_key
        __props__.__dict__["include_claims_in_id_token"] = include_claims_in_id_token
        __props__.__dict__["invalidation_flow"] = invalidation_flow
        __props__.__dict__["issuer_mode"] = issuer_mode
        __props__.__dict__["jwks_sources"] = jwks_sources
        __props__.__dict__["name"] = name
        __props__.__dict__["property_mappings"] = property_mappings
        __props__.__dict__["refresh_token_validity"] = refresh_token_validity
        __props__.__dict__["signing_key"] = signing_key
        __props__.__dict__["sub_mode"] = sub_mode
        return ProviderOauth2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessCodeValidity")
    def access_code_validity(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_code_validity")

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_token_validity")

    @property
    @pulumi.getter(name="allowedRedirectUris")
    def allowed_redirect_uris(self) -> pulumi.Output[Optional[Sequence[Mapping[str, str]]]]:
        return pulumi.get(self, "allowed_redirect_uris")

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authentication_flow")

    @property
    @pulumi.getter(name="authorizationFlow")
    def authorization_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authorization_flow")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientType")
    def client_type(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `confidential` - `public`
        """
        return pulumi.get(self, "client_type")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="includeClaimsInIdToken")
    def include_claims_in_id_token(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "include_claims_in_id_token")

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "invalidation_flow")

    @property
    @pulumi.getter(name="issuerMode")
    def issuer_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `global` - `per_provider`
        """
        return pulumi.get(self, "issuer_mode")

    @property
    @pulumi.getter(name="jwksSources")
    def jwks_sources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        """
        return pulumi.get(self, "jwks_sources")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="propertyMappings")
    def property_mappings(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "property_mappings")

    @property
    @pulumi.getter(name="refreshTokenValidity")
    def refresh_token_validity(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "refresh_token_validity")

    @property
    @pulumi.getter(name="signingKey")
    def signing_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "signing_key")

    @property
    @pulumi.getter(name="subMode")
    def sub_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `hashed_user_id` - `user_id` - `user_uuid` - `user_username` - `user_email` - `user_upn`
        """
        return pulumi.get(self, "sub_mode")

