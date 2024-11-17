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

__all__ = ['StageAuthenticatorTotpArgs', 'StageAuthenticatorTotp']

@pulumi.input_type
class StageAuthenticatorTotpArgs:
    def __init__(__self__, *,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 digits: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StageAuthenticatorTotp resource.
        :param pulumi.Input[str] digits: Allowed values: - `6` - `8`
        """
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `6` - `8`
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "digits", value)

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


@pulumi.input_type
class _StageAuthenticatorTotpState:
    def __init__(__self__, *,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 digits: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StageAuthenticatorTotp resources.
        :param pulumi.Input[str] digits: Allowed values: - `6` - `8`
        """
        if configure_flow is not None:
            pulumi.set(__self__, "configure_flow", configure_flow)
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configure_flow")

    @configure_flow.setter
    def configure_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configure_flow", value)

    @property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `6` - `8`
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "digits", value)

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


class StageAuthenticatorTotp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 digits: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a TOTP Setup stage
        name = authentik.StageAuthenticatorTotp("name", name="totp-setup")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] digits: Allowed values: - `6` - `8`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StageAuthenticatorTotpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a TOTP Setup stage
        name = authentik.StageAuthenticatorTotp("name", name="totp-setup")
        ```

        :param str resource_name: The name of the resource.
        :param StageAuthenticatorTotpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageAuthenticatorTotpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configure_flow: Optional[pulumi.Input[str]] = None,
                 digits: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageAuthenticatorTotpArgs.__new__(StageAuthenticatorTotpArgs)

            __props__.__dict__["configure_flow"] = configure_flow
            __props__.__dict__["digits"] = digits
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["name"] = name
        super(StageAuthenticatorTotp, __self__).__init__(
            'authentik:index/stageAuthenticatorTotp:StageAuthenticatorTotp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configure_flow: Optional[pulumi.Input[str]] = None,
            digits: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'StageAuthenticatorTotp':
        """
        Get an existing StageAuthenticatorTotp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] digits: Allowed values: - `6` - `8`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageAuthenticatorTotpState.__new__(_StageAuthenticatorTotpState)

        __props__.__dict__["configure_flow"] = configure_flow
        __props__.__dict__["digits"] = digits
        __props__.__dict__["friendly_name"] = friendly_name
        __props__.__dict__["name"] = name
        return StageAuthenticatorTotp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configureFlow")
    def configure_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "configure_flow")

    @property
    @pulumi.getter
    def digits(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `6` - `8`
        """
        return pulumi.get(self, "digits")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

