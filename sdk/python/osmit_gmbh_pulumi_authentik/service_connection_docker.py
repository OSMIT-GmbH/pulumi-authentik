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

__all__ = ['ServiceConnectionDockerArgs', 'ServiceConnectionDocker']

@pulumi.input_type
class ServiceConnectionDockerArgs:
    def __init__(__self__, *,
                 local: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tls_authentication: Optional[pulumi.Input[str]] = None,
                 tls_verification: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceConnectionDocker resource.
        :param pulumi.Input[bool] local: Defaults to `false`.
        :param pulumi.Input[str] url: Defaults to `http+unix:///var/run/docker.sock`.
        """
        if local is not None:
            pulumi.set(__self__, "local", local)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tls_authentication is not None:
            pulumi.set(__self__, "tls_authentication", tls_authentication)
        if tls_verification is not None:
            pulumi.set(__self__, "tls_verification", tls_verification)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tlsAuthentication")
    def tls_authentication(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_authentication")

    @tls_authentication.setter
    def tls_authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_authentication", value)

    @property
    @pulumi.getter(name="tlsVerification")
    def tls_verification(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_verification")

    @tls_verification.setter
    def tls_verification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_verification", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `http+unix:///var/run/docker.sock`.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _ServiceConnectionDockerState:
    def __init__(__self__, *,
                 local: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tls_authentication: Optional[pulumi.Input[str]] = None,
                 tls_verification: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceConnectionDocker resources.
        :param pulumi.Input[bool] local: Defaults to `false`.
        :param pulumi.Input[str] url: Defaults to `http+unix:///var/run/docker.sock`.
        """
        if local is not None:
            pulumi.set(__self__, "local", local)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tls_authentication is not None:
            pulumi.set(__self__, "tls_authentication", tls_authentication)
        if tls_verification is not None:
            pulumi.set(__self__, "tls_verification", tls_verification)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tlsAuthentication")
    def tls_authentication(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_authentication")

    @tls_authentication.setter
    def tls_authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_authentication", value)

    @property
    @pulumi.getter(name="tlsVerification")
    def tls_verification(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_verification")

    @tls_verification.setter
    def tls_verification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_verification", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `http+unix:///var/run/docker.sock`.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceConnectionDocker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tls_authentication: Optional[pulumi.Input[str]] = None,
                 tls_verification: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a local docker connection
        local = authentik.ServiceConnectionDocker("local",
            name="local",
            local=True)
        # Create a remote docker connection
        tls_auth = authentik.CertificateKeyPair("tls-auth",
            name="docker-tls-auth",
            certificate_data="...",
            key_data="...")
        tls_verification = authentik.CertificateKeyPair("tls-verification",
            name="docker-tls-verification",
            certificate_data="...")
        remote_host = authentik.ServiceConnectionDocker("remote-host",
            name="remote-host",
            url="http://1.2.3.4:2368",
            tls_verification=tls_auth.id,
            tls_authentication=tls_verification.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] local: Defaults to `false`.
        :param pulumi.Input[str] url: Defaults to `http+unix:///var/run/docker.sock`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServiceConnectionDockerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a local docker connection
        local = authentik.ServiceConnectionDocker("local",
            name="local",
            local=True)
        # Create a remote docker connection
        tls_auth = authentik.CertificateKeyPair("tls-auth",
            name="docker-tls-auth",
            certificate_data="...",
            key_data="...")
        tls_verification = authentik.CertificateKeyPair("tls-verification",
            name="docker-tls-verification",
            certificate_data="...")
        remote_host = authentik.ServiceConnectionDocker("remote-host",
            name="remote-host",
            url="http://1.2.3.4:2368",
            tls_verification=tls_auth.id,
            tls_authentication=tls_verification.id)
        ```

        :param str resource_name: The name of the resource.
        :param ServiceConnectionDockerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceConnectionDockerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tls_authentication: Optional[pulumi.Input[str]] = None,
                 tls_verification: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceConnectionDockerArgs.__new__(ServiceConnectionDockerArgs)

            __props__.__dict__["local"] = local
            __props__.__dict__["name"] = name
            __props__.__dict__["tls_authentication"] = tls_authentication
            __props__.__dict__["tls_verification"] = tls_verification
            __props__.__dict__["url"] = url
        super(ServiceConnectionDocker, __self__).__init__(
            'authentik:index/serviceConnectionDocker:ServiceConnectionDocker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            local: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tls_authentication: Optional[pulumi.Input[str]] = None,
            tls_verification: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceConnectionDocker':
        """
        Get an existing ServiceConnectionDocker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] local: Defaults to `false`.
        :param pulumi.Input[str] url: Defaults to `http+unix:///var/run/docker.sock`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceConnectionDockerState.__new__(_ServiceConnectionDockerState)

        __props__.__dict__["local"] = local
        __props__.__dict__["name"] = name
        __props__.__dict__["tls_authentication"] = tls_authentication
        __props__.__dict__["tls_verification"] = tls_verification
        __props__.__dict__["url"] = url
        return ServiceConnectionDocker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def local(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tlsAuthentication")
    def tls_authentication(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tls_authentication")

    @property
    @pulumi.getter(name="tlsVerification")
    def tls_verification(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tls_verification")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `http+unix:///var/run/docker.sock`.
        """
        return pulumi.get(self, "url")

