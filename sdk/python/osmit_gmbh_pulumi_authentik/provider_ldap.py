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

__all__ = ['ProviderLdapArgs', 'ProviderLdap']

@pulumi.input_type
class ProviderLdapArgs:
    def __init__(__self__, *,
                 base_dn: pulumi.Input[str],
                 bind_flow: pulumi.Input[str],
                 unbind_flow: pulumi.Input[str],
                 bind_mode: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 gid_start_number: Optional[pulumi.Input[int]] = None,
                 mfa_support: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_mode: Optional[pulumi.Input[str]] = None,
                 tls_server_name: Optional[pulumi.Input[str]] = None,
                 uid_start_number: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ProviderLdap resource.
        :param pulumi.Input[str] bind_mode: Defaults to `direct`.
        :param pulumi.Input[int] gid_start_number: Defaults to `4000`.
        :param pulumi.Input[bool] mfa_support: Defaults to `true`.
        :param pulumi.Input[str] search_mode: Defaults to `direct`.
        :param pulumi.Input[int] uid_start_number: Defaults to `2000`.
        """
        pulumi.set(__self__, "base_dn", base_dn)
        pulumi.set(__self__, "bind_flow", bind_flow)
        pulumi.set(__self__, "unbind_flow", unbind_flow)
        if bind_mode is not None:
            pulumi.set(__self__, "bind_mode", bind_mode)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if gid_start_number is not None:
            pulumi.set(__self__, "gid_start_number", gid_start_number)
        if mfa_support is not None:
            pulumi.set(__self__, "mfa_support", mfa_support)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if search_mode is not None:
            pulumi.set(__self__, "search_mode", search_mode)
        if tls_server_name is not None:
            pulumi.set(__self__, "tls_server_name", tls_server_name)
        if uid_start_number is not None:
            pulumi.set(__self__, "uid_start_number", uid_start_number)

    @property
    @pulumi.getter(name="baseDn")
    def base_dn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "base_dn")

    @base_dn.setter
    def base_dn(self, value: pulumi.Input[str]):
        pulumi.set(self, "base_dn", value)

    @property
    @pulumi.getter(name="bindFlow")
    def bind_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bind_flow")

    @bind_flow.setter
    def bind_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_flow", value)

    @property
    @pulumi.getter(name="unbindFlow")
    def unbind_flow(self) -> pulumi.Input[str]:
        return pulumi.get(self, "unbind_flow")

    @unbind_flow.setter
    def unbind_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "unbind_flow", value)

    @property
    @pulumi.getter(name="bindMode")
    def bind_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "bind_mode")

    @bind_mode.setter
    def bind_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_mode", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="gidStartNumber")
    def gid_start_number(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `4000`.
        """
        return pulumi.get(self, "gid_start_number")

    @gid_start_number.setter
    def gid_start_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gid_start_number", value)

    @property
    @pulumi.getter(name="mfaSupport")
    def mfa_support(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "mfa_support")

    @mfa_support.setter
    def mfa_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mfa_support", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "search_mode")

    @search_mode.setter
    def search_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_mode", value)

    @property
    @pulumi.getter(name="tlsServerName")
    def tls_server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_server_name")

    @tls_server_name.setter
    def tls_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_server_name", value)

    @property
    @pulumi.getter(name="uidStartNumber")
    def uid_start_number(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `2000`.
        """
        return pulumi.get(self, "uid_start_number")

    @uid_start_number.setter
    def uid_start_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uid_start_number", value)


@pulumi.input_type
class _ProviderLdapState:
    def __init__(__self__, *,
                 base_dn: Optional[pulumi.Input[str]] = None,
                 bind_flow: Optional[pulumi.Input[str]] = None,
                 bind_mode: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 gid_start_number: Optional[pulumi.Input[int]] = None,
                 mfa_support: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_mode: Optional[pulumi.Input[str]] = None,
                 tls_server_name: Optional[pulumi.Input[str]] = None,
                 uid_start_number: Optional[pulumi.Input[int]] = None,
                 unbind_flow: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProviderLdap resources.
        :param pulumi.Input[str] bind_mode: Defaults to `direct`.
        :param pulumi.Input[int] gid_start_number: Defaults to `4000`.
        :param pulumi.Input[bool] mfa_support: Defaults to `true`.
        :param pulumi.Input[str] search_mode: Defaults to `direct`.
        :param pulumi.Input[int] uid_start_number: Defaults to `2000`.
        """
        if base_dn is not None:
            pulumi.set(__self__, "base_dn", base_dn)
        if bind_flow is not None:
            pulumi.set(__self__, "bind_flow", bind_flow)
        if bind_mode is not None:
            pulumi.set(__self__, "bind_mode", bind_mode)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if gid_start_number is not None:
            pulumi.set(__self__, "gid_start_number", gid_start_number)
        if mfa_support is not None:
            pulumi.set(__self__, "mfa_support", mfa_support)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if search_mode is not None:
            pulumi.set(__self__, "search_mode", search_mode)
        if tls_server_name is not None:
            pulumi.set(__self__, "tls_server_name", tls_server_name)
        if uid_start_number is not None:
            pulumi.set(__self__, "uid_start_number", uid_start_number)
        if unbind_flow is not None:
            pulumi.set(__self__, "unbind_flow", unbind_flow)

    @property
    @pulumi.getter(name="baseDn")
    def base_dn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "base_dn")

    @base_dn.setter
    def base_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_dn", value)

    @property
    @pulumi.getter(name="bindFlow")
    def bind_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bind_flow")

    @bind_flow.setter
    def bind_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_flow", value)

    @property
    @pulumi.getter(name="bindMode")
    def bind_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "bind_mode")

    @bind_mode.setter
    def bind_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_mode", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="gidStartNumber")
    def gid_start_number(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `4000`.
        """
        return pulumi.get(self, "gid_start_number")

    @gid_start_number.setter
    def gid_start_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gid_start_number", value)

    @property
    @pulumi.getter(name="mfaSupport")
    def mfa_support(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "mfa_support")

    @mfa_support.setter
    def mfa_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mfa_support", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "search_mode")

    @search_mode.setter
    def search_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_mode", value)

    @property
    @pulumi.getter(name="tlsServerName")
    def tls_server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_server_name")

    @tls_server_name.setter
    def tls_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_server_name", value)

    @property
    @pulumi.getter(name="uidStartNumber")
    def uid_start_number(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `2000`.
        """
        return pulumi.get(self, "uid_start_number")

    @uid_start_number.setter
    def uid_start_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uid_start_number", value)

    @property
    @pulumi.getter(name="unbindFlow")
    def unbind_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unbind_flow")

    @unbind_flow.setter
    def unbind_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unbind_flow", value)


class ProviderLdap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_dn: Optional[pulumi.Input[str]] = None,
                 bind_flow: Optional[pulumi.Input[str]] = None,
                 bind_mode: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 gid_start_number: Optional[pulumi.Input[int]] = None,
                 mfa_support: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_mode: Optional[pulumi.Input[str]] = None,
                 tls_server_name: Optional[pulumi.Input[str]] = None,
                 uid_start_number: Optional[pulumi.Input[int]] = None,
                 unbind_flow: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        default_authentication_flow = authentik.get_flow(slug="default-authentication-flow")
        name_provider_ldap = authentik.ProviderLdap("nameProviderLdap",
            base_dn="dc=ldap,dc=goauthentik,dc=io",
            bind_flow=default_authentication_flow.id)
        name_application = authentik.Application("nameApplication",
            slug="ldap-app",
            protocol_provider=name_provider_ldap.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bind_mode: Defaults to `direct`.
        :param pulumi.Input[int] gid_start_number: Defaults to `4000`.
        :param pulumi.Input[bool] mfa_support: Defaults to `true`.
        :param pulumi.Input[str] search_mode: Defaults to `direct`.
        :param pulumi.Input[int] uid_start_number: Defaults to `2000`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderLdapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        default_authentication_flow = authentik.get_flow(slug="default-authentication-flow")
        name_provider_ldap = authentik.ProviderLdap("nameProviderLdap",
            base_dn="dc=ldap,dc=goauthentik,dc=io",
            bind_flow=default_authentication_flow.id)
        name_application = authentik.Application("nameApplication",
            slug="ldap-app",
            protocol_provider=name_provider_ldap.id)
        ```

        :param str resource_name: The name of the resource.
        :param ProviderLdapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderLdapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_dn: Optional[pulumi.Input[str]] = None,
                 bind_flow: Optional[pulumi.Input[str]] = None,
                 bind_mode: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 gid_start_number: Optional[pulumi.Input[int]] = None,
                 mfa_support: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_mode: Optional[pulumi.Input[str]] = None,
                 tls_server_name: Optional[pulumi.Input[str]] = None,
                 uid_start_number: Optional[pulumi.Input[int]] = None,
                 unbind_flow: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderLdapArgs.__new__(ProviderLdapArgs)

            if base_dn is None and not opts.urn:
                raise TypeError("Missing required property 'base_dn'")
            __props__.__dict__["base_dn"] = base_dn
            if bind_flow is None and not opts.urn:
                raise TypeError("Missing required property 'bind_flow'")
            __props__.__dict__["bind_flow"] = bind_flow
            __props__.__dict__["bind_mode"] = bind_mode
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["gid_start_number"] = gid_start_number
            __props__.__dict__["mfa_support"] = mfa_support
            __props__.__dict__["name"] = name
            __props__.__dict__["search_mode"] = search_mode
            __props__.__dict__["tls_server_name"] = tls_server_name
            __props__.__dict__["uid_start_number"] = uid_start_number
            if unbind_flow is None and not opts.urn:
                raise TypeError("Missing required property 'unbind_flow'")
            __props__.__dict__["unbind_flow"] = unbind_flow
        super(ProviderLdap, __self__).__init__(
            'authentik:index/providerLdap:ProviderLdap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            base_dn: Optional[pulumi.Input[str]] = None,
            bind_flow: Optional[pulumi.Input[str]] = None,
            bind_mode: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            gid_start_number: Optional[pulumi.Input[int]] = None,
            mfa_support: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            search_mode: Optional[pulumi.Input[str]] = None,
            tls_server_name: Optional[pulumi.Input[str]] = None,
            uid_start_number: Optional[pulumi.Input[int]] = None,
            unbind_flow: Optional[pulumi.Input[str]] = None) -> 'ProviderLdap':
        """
        Get an existing ProviderLdap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bind_mode: Defaults to `direct`.
        :param pulumi.Input[int] gid_start_number: Defaults to `4000`.
        :param pulumi.Input[bool] mfa_support: Defaults to `true`.
        :param pulumi.Input[str] search_mode: Defaults to `direct`.
        :param pulumi.Input[int] uid_start_number: Defaults to `2000`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProviderLdapState.__new__(_ProviderLdapState)

        __props__.__dict__["base_dn"] = base_dn
        __props__.__dict__["bind_flow"] = bind_flow
        __props__.__dict__["bind_mode"] = bind_mode
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["gid_start_number"] = gid_start_number
        __props__.__dict__["mfa_support"] = mfa_support
        __props__.__dict__["name"] = name
        __props__.__dict__["search_mode"] = search_mode
        __props__.__dict__["tls_server_name"] = tls_server_name
        __props__.__dict__["uid_start_number"] = uid_start_number
        __props__.__dict__["unbind_flow"] = unbind_flow
        return ProviderLdap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="baseDn")
    def base_dn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "base_dn")

    @property
    @pulumi.getter(name="bindFlow")
    def bind_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bind_flow")

    @property
    @pulumi.getter(name="bindMode")
    def bind_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "bind_mode")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="gidStartNumber")
    def gid_start_number(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `4000`.
        """
        return pulumi.get(self, "gid_start_number")

    @property
    @pulumi.getter(name="mfaSupport")
    def mfa_support(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "mfa_support")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `direct`.
        """
        return pulumi.get(self, "search_mode")

    @property
    @pulumi.getter(name="tlsServerName")
    def tls_server_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tls_server_name")

    @property
    @pulumi.getter(name="uidStartNumber")
    def uid_start_number(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `2000`.
        """
        return pulumi.get(self, "uid_start_number")

    @property
    @pulumi.getter(name="unbindFlow")
    def unbind_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "unbind_flow")

