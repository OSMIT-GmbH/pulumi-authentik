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

__all__ = ['StageAuthenticatorSmsArgs', 'StageAuthenticatorSms']

@pulumi.input_type
class StageAuthenticatorSmsArgs:
    def __init__(__self__, *,
                 account_sid: pulumi.Input[str],
                 auth: pulumi.Input[str],
                 from_number: pulumi.Input[str],
                 auth_password: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 mapping: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sms_provider: Optional[pulumi.Input[str]] = None,
                 verify_only: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a StageAuthenticatorSms resource.
        :param pulumi.Input[str] auth_type: Allowed values: - `basic` - `bearer`
        :param pulumi.Input[str] sms_provider: Allowed values: - `twilio` - `generic`
        """
        pulumi.set(__self__, "account_sid", account_sid)
        pulumi.set(__self__, "auth", auth)
        pulumi.set(__self__, "from_number", from_number)
        if auth_password is not None:
            pulumi.set(__self__, "auth_password", auth_password)
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if mapping is not None:
            pulumi.set(__self__, "mapping", mapping)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sms_provider is not None:
            pulumi.set(__self__, "sms_provider", sms_provider)
        if verify_only is not None:
            pulumi.set(__self__, "verify_only", verify_only)

    @property
    @pulumi.getter(name="accountSid")
    def account_sid(self) -> pulumi.Input[str]:
        return pulumi.get(self, "account_sid")

    @account_sid.setter
    def account_sid(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_sid", value)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="fromNumber")
    def from_number(self) -> pulumi.Input[str]:
        return pulumi.get(self, "from_number")

    @from_number.setter
    def from_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "from_number", value)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_password")

    @auth_password.setter
    def auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_password", value)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `basic` - `bearer`
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def mapping(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mapping")

    @mapping.setter
    def mapping(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapping", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="smsProvider")
    def sms_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `twilio` - `generic`
        """
        return pulumi.get(self, "sms_provider")

    @sms_provider.setter
    def sms_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sms_provider", value)

    @property
    @pulumi.getter(name="verifyOnly")
    def verify_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verify_only")

    @verify_only.setter
    def verify_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_only", value)


@pulumi.input_type
class _StageAuthenticatorSmsState:
    def __init__(__self__, *,
                 account_sid: Optional[pulumi.Input[str]] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 from_number: Optional[pulumi.Input[str]] = None,
                 mapping: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sms_provider: Optional[pulumi.Input[str]] = None,
                 verify_only: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering StageAuthenticatorSms resources.
        :param pulumi.Input[str] auth_type: Allowed values: - `basic` - `bearer`
        :param pulumi.Input[str] sms_provider: Allowed values: - `twilio` - `generic`
        """
        if account_sid is not None:
            pulumi.set(__self__, "account_sid", account_sid)
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if auth_password is not None:
            pulumi.set(__self__, "auth_password", auth_password)
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if from_number is not None:
            pulumi.set(__self__, "from_number", from_number)
        if mapping is not None:
            pulumi.set(__self__, "mapping", mapping)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sms_provider is not None:
            pulumi.set(__self__, "sms_provider", sms_provider)
        if verify_only is not None:
            pulumi.set(__self__, "verify_only", verify_only)

    @property
    @pulumi.getter(name="accountSid")
    def account_sid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_sid")

    @account_sid.setter
    def account_sid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_sid", value)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_password")

    @auth_password.setter
    def auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_password", value)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `basic` - `bearer`
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter(name="fromNumber")
    def from_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "from_number")

    @from_number.setter
    def from_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_number", value)

    @property
    @pulumi.getter
    def mapping(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mapping")

    @mapping.setter
    def mapping(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapping", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="smsProvider")
    def sms_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `twilio` - `generic`
        """
        return pulumi.get(self, "sms_provider")

    @sms_provider.setter
    def sms_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sms_provider", value)

    @property
    @pulumi.getter(name="verifyOnly")
    def verify_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verify_only")

    @verify_only.setter
    def verify_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_only", value)


class StageAuthenticatorSms(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_sid: Optional[pulumi.Input[str]] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 from_number: Optional[pulumi.Input[str]] = None,
                 mapping: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sms_provider: Optional[pulumi.Input[str]] = None,
                 verify_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a StageAuthenticatorSms resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: Allowed values: - `basic` - `bearer`
        :param pulumi.Input[str] sms_provider: Allowed values: - `twilio` - `generic`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StageAuthenticatorSmsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a StageAuthenticatorSms resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param StageAuthenticatorSmsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageAuthenticatorSmsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_sid: Optional[pulumi.Input[str]] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 from_number: Optional[pulumi.Input[str]] = None,
                 mapping: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sms_provider: Optional[pulumi.Input[str]] = None,
                 verify_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageAuthenticatorSmsArgs.__new__(StageAuthenticatorSmsArgs)

            if account_sid is None and not opts.urn:
                raise TypeError("Missing required property 'account_sid'")
            __props__.__dict__["account_sid"] = None if account_sid is None else pulumi.Output.secret(account_sid)
            if auth is None and not opts.urn:
                raise TypeError("Missing required property 'auth'")
            __props__.__dict__["auth"] = None if auth is None else pulumi.Output.secret(auth)
            __props__.__dict__["auth_password"] = None if auth_password is None else pulumi.Output.secret(auth_password)
            __props__.__dict__["auth_type"] = auth_type
            __props__.__dict__["configure_flow"] = configure_flow
            __props__.__dict__["friendly_name"] = friendly_name
            if from_number is None and not opts.urn:
                raise TypeError("Missing required property 'from_number'")
            __props__.__dict__["from_number"] = from_number
            __props__.__dict__["mapping"] = mapping
            __props__.__dict__["name"] = name
            __props__.__dict__["sms_provider"] = sms_provider
            __props__.__dict__["verify_only"] = verify_only
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accountSid", "auth", "authPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(StageAuthenticatorSms, __self__).__init__(
            'authentik:index/stageAuthenticatorSms:StageAuthenticatorSms',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_sid: Optional[pulumi.Input[str]] = None,
            auth: Optional[pulumi.Input[str]] = None,
            auth_password: Optional[pulumi.Input[str]] = None,
            auth_type: Optional[pulumi.Input[str]] = None,
            configure_flow: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            from_number: Optional[pulumi.Input[str]] = None,
            mapping: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            sms_provider: Optional[pulumi.Input[str]] = None,
            verify_only: Optional[pulumi.Input[bool]] = None) -> 'StageAuthenticatorSms':
        """
        Get an existing StageAuthenticatorSms resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: Allowed values: - `basic` - `bearer`
        :param pulumi.Input[str] sms_provider: Allowed values: - `twilio` - `generic`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageAuthenticatorSmsState.__new__(_StageAuthenticatorSmsState)

        __props__.__dict__["account_sid"] = account_sid
        __props__.__dict__["auth"] = auth
        __props__.__dict__["auth_password"] = auth_password
        __props__.__dict__["auth_type"] = auth_type
        __props__.__dict__["configure_flow"] = configure_flow
        __props__.__dict__["friendly_name"] = friendly_name
        __props__.__dict__["from_number"] = from_number
        __props__.__dict__["mapping"] = mapping
        __props__.__dict__["name"] = name
        __props__.__dict__["sms_provider"] = sms_provider
        __props__.__dict__["verify_only"] = verify_only
        return StageAuthenticatorSms(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountSid")
    def account_sid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "account_sid")

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "auth_password")

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `basic` - `bearer`
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "configure_flow")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="fromNumber")
    def from_number(self) -> pulumi.Output[str]:
        return pulumi.get(self, "from_number")

    @property
    @pulumi.getter
    def mapping(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mapping")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="smsProvider")
    def sms_provider(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `twilio` - `generic`
        """
        return pulumi.get(self, "sms_provider")

    @property
    @pulumi.getter(name="verifyOnly")
    def verify_only(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "verify_only")

