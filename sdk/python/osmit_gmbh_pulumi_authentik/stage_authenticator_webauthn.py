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

__all__ = ['StageAuthenticatorWebauthnArgs', 'StageAuthenticatorWebauthn']

@pulumi.input_type
class StageAuthenticatorWebauthnArgs:
    def __init__(__self__, *,
                 authenticator_attachment: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 device_type_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resident_key_requirement: Optional[pulumi.Input[str]] = None,
                 user_verification: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StageAuthenticatorWebauthn resource.
        :param pulumi.Input[str] authenticator_attachment: Allowed values: - `platform` - `cross-platform`
        :param pulumi.Input[str] resident_key_requirement: Allowed values: - `discouraged` - `preferred` - `required`
        :param pulumi.Input[str] user_verification: Allowed values: - `required` - `preferred` - `discouraged`
        """
        if authenticator_attachment is not None:
            pulumi.set(__self__, "authenticator_attachment", authenticator_attachment)
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if device_type_restrictions is not None:
            pulumi.set(__self__, "device_type_restrictions", device_type_restrictions)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resident_key_requirement is not None:
            pulumi.set(__self__, "resident_key_requirement", resident_key_requirement)
        if user_verification is not None:
            pulumi.set(__self__, "user_verification", user_verification)

    @property
    @pulumi.getter(name="authenticatorAttachment")
    def authenticator_attachment(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `platform` - `cross-platform`
        """
        return pulumi.get(self, "authenticator_attachment")

    @authenticator_attachment.setter
    def authenticator_attachment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticator_attachment", value)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter(name="deviceTypeRestrictions")
    def device_type_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "device_type_restrictions")

    @device_type_restrictions.setter
    def device_type_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "device_type_restrictions", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="residentKeyRequirement")
    def resident_key_requirement(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `discouraged` - `preferred` - `required`
        """
        return pulumi.get(self, "resident_key_requirement")

    @resident_key_requirement.setter
    def resident_key_requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resident_key_requirement", value)

    @property
    @pulumi.getter(name="userVerification")
    def user_verification(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `required` - `preferred` - `discouraged`
        """
        return pulumi.get(self, "user_verification")

    @user_verification.setter
    def user_verification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_verification", value)


@pulumi.input_type
class _StageAuthenticatorWebauthnState:
    def __init__(__self__, *,
                 authenticator_attachment: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 device_type_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resident_key_requirement: Optional[pulumi.Input[str]] = None,
                 user_verification: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StageAuthenticatorWebauthn resources.
        :param pulumi.Input[str] authenticator_attachment: Allowed values: - `platform` - `cross-platform`
        :param pulumi.Input[str] resident_key_requirement: Allowed values: - `discouraged` - `preferred` - `required`
        :param pulumi.Input[str] user_verification: Allowed values: - `required` - `preferred` - `discouraged`
        """
        if authenticator_attachment is not None:
            pulumi.set(__self__, "authenticator_attachment", authenticator_attachment)
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if device_type_restrictions is not None:
            pulumi.set(__self__, "device_type_restrictions", device_type_restrictions)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resident_key_requirement is not None:
            pulumi.set(__self__, "resident_key_requirement", resident_key_requirement)
        if user_verification is not None:
            pulumi.set(__self__, "user_verification", user_verification)

    @property
    @pulumi.getter(name="authenticatorAttachment")
    def authenticator_attachment(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `platform` - `cross-platform`
        """
        return pulumi.get(self, "authenticator_attachment")

    @authenticator_attachment.setter
    def authenticator_attachment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticator_attachment", value)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter(name="deviceTypeRestrictions")
    def device_type_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "device_type_restrictions")

    @device_type_restrictions.setter
    def device_type_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "device_type_restrictions", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="residentKeyRequirement")
    def resident_key_requirement(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `discouraged` - `preferred` - `required`
        """
        return pulumi.get(self, "resident_key_requirement")

    @resident_key_requirement.setter
    def resident_key_requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resident_key_requirement", value)

    @property
    @pulumi.getter(name="userVerification")
    def user_verification(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `required` - `preferred` - `discouraged`
        """
        return pulumi.get(self, "user_verification")

    @user_verification.setter
    def user_verification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_verification", value)


class StageAuthenticatorWebauthn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticator_attachment: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 device_type_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resident_key_requirement: Optional[pulumi.Input[str]] = None,
                 user_verification: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator_attachment: Allowed values: - `platform` - `cross-platform`
        :param pulumi.Input[str] resident_key_requirement: Allowed values: - `discouraged` - `preferred` - `required`
        :param pulumi.Input[str] user_verification: Allowed values: - `required` - `preferred` - `discouraged`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StageAuthenticatorWebauthnArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param StageAuthenticatorWebauthnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageAuthenticatorWebauthnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticator_attachment: Optional[pulumi.Input[str]] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 device_type_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resident_key_requirement: Optional[pulumi.Input[str]] = None,
                 user_verification: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageAuthenticatorWebauthnArgs.__new__(StageAuthenticatorWebauthnArgs)

            __props__.__dict__["authenticator_attachment"] = authenticator_attachment
            __props__.__dict__["configure_flow"] = configure_flow
            __props__.__dict__["device_type_restrictions"] = device_type_restrictions
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["name"] = name
            __props__.__dict__["resident_key_requirement"] = resident_key_requirement
            __props__.__dict__["user_verification"] = user_verification
        super(StageAuthenticatorWebauthn, __self__).__init__(
            'authentik:index/stageAuthenticatorWebauthn:StageAuthenticatorWebauthn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authenticator_attachment: Optional[pulumi.Input[str]] = None,
            configure_flow: Optional[pulumi.Input[str]] = None,
            device_type_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resident_key_requirement: Optional[pulumi.Input[str]] = None,
            user_verification: Optional[pulumi.Input[str]] = None) -> 'StageAuthenticatorWebauthn':
        """
        Get an existing StageAuthenticatorWebauthn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator_attachment: Allowed values: - `platform` - `cross-platform`
        :param pulumi.Input[str] resident_key_requirement: Allowed values: - `discouraged` - `preferred` - `required`
        :param pulumi.Input[str] user_verification: Allowed values: - `required` - `preferred` - `discouraged`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageAuthenticatorWebauthnState.__new__(_StageAuthenticatorWebauthnState)

        __props__.__dict__["authenticator_attachment"] = authenticator_attachment
        __props__.__dict__["configure_flow"] = configure_flow
        __props__.__dict__["device_type_restrictions"] = device_type_restrictions
        __props__.__dict__["friendly_name"] = friendly_name
        __props__.__dict__["name"] = name
        __props__.__dict__["resident_key_requirement"] = resident_key_requirement
        __props__.__dict__["user_verification"] = user_verification
        return StageAuthenticatorWebauthn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticatorAttachment")
    def authenticator_attachment(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `platform` - `cross-platform`
        """
        return pulumi.get(self, "authenticator_attachment")

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "configure_flow")

    @property
    @pulumi.getter(name="deviceTypeRestrictions")
    def device_type_restrictions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "device_type_restrictions")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="residentKeyRequirement")
    def resident_key_requirement(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `discouraged` - `preferred` - `required`
        """
        return pulumi.get(self, "resident_key_requirement")

    @property
    @pulumi.getter(name="userVerification")
    def user_verification(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `required` - `preferred` - `discouraged`
        """
        return pulumi.get(self, "user_verification")

