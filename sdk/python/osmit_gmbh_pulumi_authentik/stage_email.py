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

__all__ = ['StageEmailArgs', 'StageEmail']

@pulumi.input_type
class StageEmailArgs:
    def __init__(__self__, *,
                 activate_user_on_success: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 token_expiry: Optional[pulumi.Input[int]] = None,
                 use_global_settings: Optional[pulumi.Input[bool]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StageEmail resource.
        :param pulumi.Input[bool] activate_user_on_success: Defaults to `false`.
        :param pulumi.Input[str] from_address: Defaults to `system@authentik.local`.
        :param pulumi.Input[str] host: Defaults to `localhost`.
        :param pulumi.Input[int] port: Defaults to `25`.
        :param pulumi.Input[str] subject: Defaults to `authentik`.
        :param pulumi.Input[str] template: Defaults to `email/password_reset.html`.
        :param pulumi.Input[int] timeout: Defaults to `30`.
        :param pulumi.Input[int] token_expiry: Defaults to `30`.
        :param pulumi.Input[bool] use_global_settings: Defaults to `true`.
        """
        if activate_user_on_success is not None:
            pulumi.set(__self__, "activate_user_on_success", activate_user_on_success)
        if from_address is not None:
            pulumi.set(__self__, "from_address", from_address)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if token_expiry is not None:
            pulumi.set(__self__, "token_expiry", token_expiry)
        if use_global_settings is not None:
            pulumi.set(__self__, "use_global_settings", use_global_settings)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if use_tls is not None:
            pulumi.set(__self__, "use_tls", use_tls)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="activateUserOnSuccess")
    def activate_user_on_success(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "activate_user_on_success")

    @activate_user_on_success.setter
    def activate_user_on_success(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "activate_user_on_success", value)

    @property
    @pulumi.getter(name="fromAddress")
    def from_address(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `system@authentik.local`.
        """
        return pulumi.get(self, "from_address")

    @from_address.setter
    def from_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_address", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `25`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `authentik`.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `email/password_reset.html`.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="tokenExpiry")
    def token_expiry(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "token_expiry")

    @token_expiry.setter
    def token_expiry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "token_expiry", value)

    @property
    @pulumi.getter(name="useGlobalSettings")
    def use_global_settings(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "use_global_settings")

    @use_global_settings.setter
    def use_global_settings(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_global_settings", value)

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_tls")

    @use_tls.setter
    def use_tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_tls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _StageEmailState:
    def __init__(__self__, *,
                 activate_user_on_success: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 token_expiry: Optional[pulumi.Input[int]] = None,
                 use_global_settings: Optional[pulumi.Input[bool]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StageEmail resources.
        :param pulumi.Input[bool] activate_user_on_success: Defaults to `false`.
        :param pulumi.Input[str] from_address: Defaults to `system@authentik.local`.
        :param pulumi.Input[str] host: Defaults to `localhost`.
        :param pulumi.Input[int] port: Defaults to `25`.
        :param pulumi.Input[str] subject: Defaults to `authentik`.
        :param pulumi.Input[str] template: Defaults to `email/password_reset.html`.
        :param pulumi.Input[int] timeout: Defaults to `30`.
        :param pulumi.Input[int] token_expiry: Defaults to `30`.
        :param pulumi.Input[bool] use_global_settings: Defaults to `true`.
        """
        if activate_user_on_success is not None:
            pulumi.set(__self__, "activate_user_on_success", activate_user_on_success)
        if from_address is not None:
            pulumi.set(__self__, "from_address", from_address)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if token_expiry is not None:
            pulumi.set(__self__, "token_expiry", token_expiry)
        if use_global_settings is not None:
            pulumi.set(__self__, "use_global_settings", use_global_settings)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if use_tls is not None:
            pulumi.set(__self__, "use_tls", use_tls)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="activateUserOnSuccess")
    def activate_user_on_success(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "activate_user_on_success")

    @activate_user_on_success.setter
    def activate_user_on_success(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "activate_user_on_success", value)

    @property
    @pulumi.getter(name="fromAddress")
    def from_address(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `system@authentik.local`.
        """
        return pulumi.get(self, "from_address")

    @from_address.setter
    def from_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_address", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `25`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `authentik`.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `email/password_reset.html`.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="tokenExpiry")
    def token_expiry(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "token_expiry")

    @token_expiry.setter
    def token_expiry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "token_expiry", value)

    @property
    @pulumi.getter(name="useGlobalSettings")
    def use_global_settings(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "use_global_settings")

    @use_global_settings.setter
    def use_global_settings(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_global_settings", value)

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_tls")

    @use_tls.setter
    def use_tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_tls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class StageEmail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate_user_on_success: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 token_expiry: Optional[pulumi.Input[int]] = None,
                 use_global_settings: Optional[pulumi.Input[bool]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create email stage for email verification, uses global settings by default
        name = authentik.StageEmail("name", name="email-verification")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] activate_user_on_success: Defaults to `false`.
        :param pulumi.Input[str] from_address: Defaults to `system@authentik.local`.
        :param pulumi.Input[str] host: Defaults to `localhost`.
        :param pulumi.Input[int] port: Defaults to `25`.
        :param pulumi.Input[str] subject: Defaults to `authentik`.
        :param pulumi.Input[str] template: Defaults to `email/password_reset.html`.
        :param pulumi.Input[int] timeout: Defaults to `30`.
        :param pulumi.Input[int] token_expiry: Defaults to `30`.
        :param pulumi.Input[bool] use_global_settings: Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StageEmailArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create email stage for email verification, uses global settings by default
        name = authentik.StageEmail("name", name="email-verification")
        ```

        :param str resource_name: The name of the resource.
        :param StageEmailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageEmailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate_user_on_success: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 token_expiry: Optional[pulumi.Input[int]] = None,
                 use_global_settings: Optional[pulumi.Input[bool]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 use_tls: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageEmailArgs.__new__(StageEmailArgs)

            __props__.__dict__["activate_user_on_success"] = activate_user_on_success
            __props__.__dict__["from_address"] = from_address
            __props__.__dict__["host"] = host
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = port
            __props__.__dict__["subject"] = subject
            __props__.__dict__["template"] = template
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["token_expiry"] = token_expiry
            __props__.__dict__["use_global_settings"] = use_global_settings
            __props__.__dict__["use_ssl"] = use_ssl
            __props__.__dict__["use_tls"] = use_tls
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(StageEmail, __self__).__init__(
            'authentik:index/stageEmail:StageEmail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activate_user_on_success: Optional[pulumi.Input[bool]] = None,
            from_address: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            subject: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            token_expiry: Optional[pulumi.Input[int]] = None,
            use_global_settings: Optional[pulumi.Input[bool]] = None,
            use_ssl: Optional[pulumi.Input[bool]] = None,
            use_tls: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'StageEmail':
        """
        Get an existing StageEmail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] activate_user_on_success: Defaults to `false`.
        :param pulumi.Input[str] from_address: Defaults to `system@authentik.local`.
        :param pulumi.Input[str] host: Defaults to `localhost`.
        :param pulumi.Input[int] port: Defaults to `25`.
        :param pulumi.Input[str] subject: Defaults to `authentik`.
        :param pulumi.Input[str] template: Defaults to `email/password_reset.html`.
        :param pulumi.Input[int] timeout: Defaults to `30`.
        :param pulumi.Input[int] token_expiry: Defaults to `30`.
        :param pulumi.Input[bool] use_global_settings: Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageEmailState.__new__(_StageEmailState)

        __props__.__dict__["activate_user_on_success"] = activate_user_on_success
        __props__.__dict__["from_address"] = from_address
        __props__.__dict__["host"] = host
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["subject"] = subject
        __props__.__dict__["template"] = template
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["token_expiry"] = token_expiry
        __props__.__dict__["use_global_settings"] = use_global_settings
        __props__.__dict__["use_ssl"] = use_ssl
        __props__.__dict__["use_tls"] = use_tls
        __props__.__dict__["username"] = username
        return StageEmail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activateUserOnSuccess")
    def activate_user_on_success(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "activate_user_on_success")

    @property
    @pulumi.getter(name="fromAddress")
    def from_address(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `system@authentik.local`.
        """
        return pulumi.get(self, "from_address")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `25`.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `authentik`.
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `email/password_reset.html`.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="tokenExpiry")
    def token_expiry(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "token_expiry")

    @property
    @pulumi.getter(name="useGlobalSettings")
    def use_global_settings(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "use_global_settings")

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "use_ssl")

    @property
    @pulumi.getter(name="useTls")
    def use_tls(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "use_tls")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")
