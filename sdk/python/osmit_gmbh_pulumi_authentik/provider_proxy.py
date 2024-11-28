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

__all__ = ['ProviderProxyArgs', 'ProviderProxy']

@pulumi.input_type
class ProviderProxyArgs:
    def __init__(__self__, *,
                 authorization_flow: pulumi.Input[str],
                 external_host: pulumi.Input[str],
                 invalidation_flow: pulumi.Input[str],
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password_attribute: Optional[pulumi.Input[str]] = None,
                 basic_auth_username_attribute: Optional[pulumi.Input[str]] = None,
                 cookie_domain: Optional[pulumi.Input[str]] = None,
                 intercept_header_auth: Optional[pulumi.Input[bool]] = None,
                 internal_host: Optional[pulumi.Input[str]] = None,
                 internal_host_ssl_validation: Optional[pulumi.Input[bool]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 skip_path_regex: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProviderProxy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] mode: Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        pulumi.set(__self__, "authorization_flow", authorization_flow)
        pulumi.set(__self__, "external_host", external_host)
        pulumi.set(__self__, "invalidation_flow", invalidation_flow)
        if access_token_validity is not None:
            pulumi.set(__self__, "access_token_validity", access_token_validity)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if basic_auth_enabled is not None:
            pulumi.set(__self__, "basic_auth_enabled", basic_auth_enabled)
        if basic_auth_password_attribute is not None:
            pulumi.set(__self__, "basic_auth_password_attribute", basic_auth_password_attribute)
        if basic_auth_username_attribute is not None:
            pulumi.set(__self__, "basic_auth_username_attribute", basic_auth_username_attribute)
        if cookie_domain is not None:
            pulumi.set(__self__, "cookie_domain", cookie_domain)
        if intercept_header_auth is not None:
            pulumi.set(__self__, "intercept_header_auth", intercept_header_auth)
        if internal_host is not None:
            pulumi.set(__self__, "internal_host", internal_host)
        if internal_host_ssl_validation is not None:
            pulumi.set(__self__, "internal_host_ssl_validation", internal_host_ssl_validation)
        if jwks_sources is not None:
            pulumi.set(__self__, "jwks_sources", jwks_sources)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if property_mappings is not None:
            pulumi.set(__self__, "property_mappings", property_mappings)
        if refresh_token_validity is not None:
            pulumi.set(__self__, "refresh_token_validity", refresh_token_validity)
        if skip_path_regex is not None:
            pulumi.set(__self__, "skip_path_regex", skip_path_regex)

    @property
    @pulumi.getter(name="authorizationFlow")
    def authorization_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authorization_flow")

    @authorization_flow.setter
    def authorization_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_flow", value)

    @property
    @pulumi.getter(name="externalHost")
    def external_host(self) -> pulumi.Input[str]:
        return pulumi.get(self, "external_host")

    @external_host.setter
    def external_host(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_host", value)

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "invalidation_flow")

    @invalidation_flow.setter
    def invalidation_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "invalidation_flow", value)

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_token_validity")

    @access_token_validity.setter
    def access_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_validity", value)

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authentication_flow")

    @authentication_flow.setter
    def authentication_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_flow", value)

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "basic_auth_enabled")

    @basic_auth_enabled.setter
    def basic_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "basic_auth_enabled", value)

    @property
    @pulumi.getter(name="basicAuthPasswordAttribute")
    def basic_auth_password_attribute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "basic_auth_password_attribute")

    @basic_auth_password_attribute.setter
    def basic_auth_password_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_password_attribute", value)

    @property
    @pulumi.getter(name="basicAuthUsernameAttribute")
    def basic_auth_username_attribute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "basic_auth_username_attribute")

    @basic_auth_username_attribute.setter
    def basic_auth_username_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_username_attribute", value)

    @property
    @pulumi.getter(name="cookieDomain")
    def cookie_domain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cookie_domain")

    @cookie_domain.setter
    def cookie_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie_domain", value)

    @property
    @pulumi.getter(name="interceptHeaderAuth")
    def intercept_header_auth(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "intercept_header_auth")

    @intercept_header_auth.setter
    def intercept_header_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "intercept_header_auth", value)

    @property
    @pulumi.getter(name="internalHost")
    def internal_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "internal_host")

    @internal_host.setter
    def internal_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_host", value)

    @property
    @pulumi.getter(name="internalHostSslValidation")
    def internal_host_ssl_validation(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "internal_host_ssl_validation")

    @internal_host_ssl_validation.setter
    def internal_host_ssl_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_host_ssl_validation", value)

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
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

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
    @pulumi.getter(name="skipPathRegex")
    def skip_path_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "skip_path_regex")

    @skip_path_regex.setter
    def skip_path_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "skip_path_regex", value)


@pulumi.input_type
class _ProviderProxyState:
    def __init__(__self__, *,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password_attribute: Optional[pulumi.Input[str]] = None,
                 basic_auth_username_attribute: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 cookie_domain: Optional[pulumi.Input[str]] = None,
                 external_host: Optional[pulumi.Input[str]] = None,
                 intercept_header_auth: Optional[pulumi.Input[bool]] = None,
                 internal_host: Optional[pulumi.Input[str]] = None,
                 internal_host_ssl_validation: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 skip_path_regex: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProviderProxy resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] mode: Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        if access_token_validity is not None:
            pulumi.set(__self__, "access_token_validity", access_token_validity)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if authorization_flow is not None:
            pulumi.set(__self__, "authorization_flow", authorization_flow)
        if basic_auth_enabled is not None:
            pulumi.set(__self__, "basic_auth_enabled", basic_auth_enabled)
        if basic_auth_password_attribute is not None:
            pulumi.set(__self__, "basic_auth_password_attribute", basic_auth_password_attribute)
        if basic_auth_username_attribute is not None:
            pulumi.set(__self__, "basic_auth_username_attribute", basic_auth_username_attribute)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if cookie_domain is not None:
            pulumi.set(__self__, "cookie_domain", cookie_domain)
        if external_host is not None:
            pulumi.set(__self__, "external_host", external_host)
        if intercept_header_auth is not None:
            pulumi.set(__self__, "intercept_header_auth", intercept_header_auth)
        if internal_host is not None:
            pulumi.set(__self__, "internal_host", internal_host)
        if internal_host_ssl_validation is not None:
            pulumi.set(__self__, "internal_host_ssl_validation", internal_host_ssl_validation)
        if invalidation_flow is not None:
            pulumi.set(__self__, "invalidation_flow", invalidation_flow)
        if jwks_sources is not None:
            pulumi.set(__self__, "jwks_sources", jwks_sources)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if property_mappings is not None:
            pulumi.set(__self__, "property_mappings", property_mappings)
        if refresh_token_validity is not None:
            pulumi.set(__self__, "refresh_token_validity", refresh_token_validity)
        if skip_path_regex is not None:
            pulumi.set(__self__, "skip_path_regex", skip_path_regex)

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_token_validity")

    @access_token_validity.setter
    def access_token_validity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_validity", value)

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
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "basic_auth_enabled")

    @basic_auth_enabled.setter
    def basic_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "basic_auth_enabled", value)

    @property
    @pulumi.getter(name="basicAuthPasswordAttribute")
    def basic_auth_password_attribute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "basic_auth_password_attribute")

    @basic_auth_password_attribute.setter
    def basic_auth_password_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_password_attribute", value)

    @property
    @pulumi.getter(name="basicAuthUsernameAttribute")
    def basic_auth_username_attribute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "basic_auth_username_attribute")

    @basic_auth_username_attribute.setter
    def basic_auth_username_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_auth_username_attribute", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="cookieDomain")
    def cookie_domain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cookie_domain")

    @cookie_domain.setter
    def cookie_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie_domain", value)

    @property
    @pulumi.getter(name="externalHost")
    def external_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "external_host")

    @external_host.setter
    def external_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_host", value)

    @property
    @pulumi.getter(name="interceptHeaderAuth")
    def intercept_header_auth(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "intercept_header_auth")

    @intercept_header_auth.setter
    def intercept_header_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "intercept_header_auth", value)

    @property
    @pulumi.getter(name="internalHost")
    def internal_host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "internal_host")

    @internal_host.setter
    def internal_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_host", value)

    @property
    @pulumi.getter(name="internalHostSslValidation")
    def internal_host_ssl_validation(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "internal_host_ssl_validation")

    @internal_host_ssl_validation.setter
    def internal_host_ssl_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_host_ssl_validation", value)

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "invalidation_flow")

    @invalidation_flow.setter
    def invalidation_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invalidation_flow", value)

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
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

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
    @pulumi.getter(name="skipPathRegex")
    def skip_path_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "skip_path_regex")

    @skip_path_regex.setter
    def skip_path_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "skip_path_regex", value)


class ProviderProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password_attribute: Optional[pulumi.Input[str]] = None,
                 basic_auth_username_attribute: Optional[pulumi.Input[str]] = None,
                 cookie_domain: Optional[pulumi.Input[str]] = None,
                 external_host: Optional[pulumi.Input[str]] = None,
                 intercept_header_auth: Optional[pulumi.Input[bool]] = None,
                 internal_host: Optional[pulumi.Input[str]] = None,
                 internal_host_ssl_validation: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 skip_path_regex: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        # Create a proxy provider
        default_authorization_flow = authentik.get_flow(slug="default-provider-authorization-implicit-consent")
        name = authentik.ProviderProxy("name",
            name="test-app",
            internal_host="http://foo.bar.baz",
            external_host="http://internal.service",
            authorization_flow=default_authorization_flow.id)
        name_application = authentik.Application("name",
            name="test-app",
            slug="test-app",
            protocol_provider=name.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] mode: Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        # Create a proxy provider
        default_authorization_flow = authentik.get_flow(slug="default-provider-authorization-implicit-consent")
        name = authentik.ProviderProxy("name",
            name="test-app",
            internal_host="http://foo.bar.baz",
            external_host="http://internal.service",
            authorization_flow=default_authorization_flow.id)
        name_application = authentik.Application("name",
            name="test-app",
            slug="test-app",
            protocol_provider=name.id)
        ```

        :param str resource_name: The name of the resource.
        :param ProviderProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token_validity: Optional[pulumi.Input[str]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 authorization_flow: Optional[pulumi.Input[str]] = None,
                 basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 basic_auth_password_attribute: Optional[pulumi.Input[str]] = None,
                 basic_auth_username_attribute: Optional[pulumi.Input[str]] = None,
                 cookie_domain: Optional[pulumi.Input[str]] = None,
                 external_host: Optional[pulumi.Input[str]] = None,
                 intercept_header_auth: Optional[pulumi.Input[bool]] = None,
                 internal_host: Optional[pulumi.Input[str]] = None,
                 internal_host_ssl_validation: Optional[pulumi.Input[bool]] = None,
                 invalidation_flow: Optional[pulumi.Input[str]] = None,
                 jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token_validity: Optional[pulumi.Input[str]] = None,
                 skip_path_regex: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderProxyArgs.__new__(ProviderProxyArgs)

            __props__.__dict__["access_token_validity"] = access_token_validity
            __props__.__dict__["authentication_flow"] = authentication_flow
            if authorization_flow is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_flow'")
            __props__.__dict__["authorization_flow"] = authorization_flow
            __props__.__dict__["basic_auth_enabled"] = basic_auth_enabled
            __props__.__dict__["basic_auth_password_attribute"] = basic_auth_password_attribute
            __props__.__dict__["basic_auth_username_attribute"] = basic_auth_username_attribute
            __props__.__dict__["cookie_domain"] = cookie_domain
            if external_host is None and not opts.urn:
                raise TypeError("Missing required property 'external_host'")
            __props__.__dict__["external_host"] = external_host
            __props__.__dict__["intercept_header_auth"] = intercept_header_auth
            __props__.__dict__["internal_host"] = internal_host
            __props__.__dict__["internal_host_ssl_validation"] = internal_host_ssl_validation
            if invalidation_flow is None and not opts.urn:
                raise TypeError("Missing required property 'invalidation_flow'")
            __props__.__dict__["invalidation_flow"] = invalidation_flow
            __props__.__dict__["jwks_sources"] = jwks_sources
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            __props__.__dict__["property_mappings"] = property_mappings
            __props__.__dict__["refresh_token_validity"] = refresh_token_validity
            __props__.__dict__["skip_path_regex"] = skip_path_regex
            __props__.__dict__["client_id"] = None
        super(ProviderProxy, __self__).__init__(
            'authentik:index/providerProxy:ProviderProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token_validity: Optional[pulumi.Input[str]] = None,
            authentication_flow: Optional[pulumi.Input[str]] = None,
            authorization_flow: Optional[pulumi.Input[str]] = None,
            basic_auth_enabled: Optional[pulumi.Input[bool]] = None,
            basic_auth_password_attribute: Optional[pulumi.Input[str]] = None,
            basic_auth_username_attribute: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            cookie_domain: Optional[pulumi.Input[str]] = None,
            external_host: Optional[pulumi.Input[str]] = None,
            intercept_header_auth: Optional[pulumi.Input[bool]] = None,
            internal_host: Optional[pulumi.Input[str]] = None,
            internal_host_ssl_validation: Optional[pulumi.Input[bool]] = None,
            invalidation_flow: Optional[pulumi.Input[str]] = None,
            jwks_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            property_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            refresh_token_validity: Optional[pulumi.Input[str]] = None,
            skip_path_regex: Optional[pulumi.Input[str]] = None) -> 'ProviderProxy':
        """
        Get an existing ProviderProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] jwks_sources: JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        :param pulumi.Input[str] mode: Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProviderProxyState.__new__(_ProviderProxyState)

        __props__.__dict__["access_token_validity"] = access_token_validity
        __props__.__dict__["authentication_flow"] = authentication_flow
        __props__.__dict__["authorization_flow"] = authorization_flow
        __props__.__dict__["basic_auth_enabled"] = basic_auth_enabled
        __props__.__dict__["basic_auth_password_attribute"] = basic_auth_password_attribute
        __props__.__dict__["basic_auth_username_attribute"] = basic_auth_username_attribute
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["cookie_domain"] = cookie_domain
        __props__.__dict__["external_host"] = external_host
        __props__.__dict__["intercept_header_auth"] = intercept_header_auth
        __props__.__dict__["internal_host"] = internal_host
        __props__.__dict__["internal_host_ssl_validation"] = internal_host_ssl_validation
        __props__.__dict__["invalidation_flow"] = invalidation_flow
        __props__.__dict__["jwks_sources"] = jwks_sources
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["property_mappings"] = property_mappings
        __props__.__dict__["refresh_token_validity"] = refresh_token_validity
        __props__.__dict__["skip_path_regex"] = skip_path_regex
        return ProviderProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessTokenValidity")
    def access_token_validity(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_token_validity")

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authentication_flow")

    @property
    @pulumi.getter(name="authorizationFlow")
    def authorization_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authorization_flow")

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "basic_auth_enabled")

    @property
    @pulumi.getter(name="basicAuthPasswordAttribute")
    def basic_auth_password_attribute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "basic_auth_password_attribute")

    @property
    @pulumi.getter(name="basicAuthUsernameAttribute")
    def basic_auth_username_attribute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "basic_auth_username_attribute")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="cookieDomain")
    def cookie_domain(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cookie_domain")

    @property
    @pulumi.getter(name="externalHost")
    def external_host(self) -> pulumi.Output[str]:
        return pulumi.get(self, "external_host")

    @property
    @pulumi.getter(name="interceptHeaderAuth")
    def intercept_header_auth(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "intercept_header_auth")

    @property
    @pulumi.getter(name="internalHost")
    def internal_host(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "internal_host")

    @property
    @pulumi.getter(name="internalHostSslValidation")
    def internal_host_ssl_validation(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "internal_host_ssl_validation")

    @property
    @pulumi.getter(name="invalidationFlow")
    def invalidation_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "invalidation_flow")

    @property
    @pulumi.getter(name="jwksSources")
    def jwks_sources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        JWTs issued by keys configured in any of the selected sources can be used to authenticate on behalf of this provider.
        """
        return pulumi.get(self, "jwks_sources")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `proxy` - `forward_single` - `forward_domain`
        """
        return pulumi.get(self, "mode")

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
    @pulumi.getter(name="skipPathRegex")
    def skip_path_regex(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "skip_path_regex")
