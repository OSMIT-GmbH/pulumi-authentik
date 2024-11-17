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

__all__ = ['PolicyExpiryArgs', 'PolicyExpiry']

@pulumi.input_type
class PolicyExpiryArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int],
                 deny_only: Optional[pulumi.Input[bool]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyExpiry resource.
        :param pulumi.Input[bool] deny_only: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        """
        pulumi.set(__self__, "days", days)
        if deny_only is not None:
            pulumi.set(__self__, "deny_only", deny_only)
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="denyOnly")
    def deny_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "deny_only")

    @deny_only.setter
    def deny_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deny_only", value)

    @property
    @pulumi.getter(name="executionLogging")
    def execution_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "execution_logging")

    @execution_logging.setter
    def execution_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "execution_logging", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PolicyExpiryState:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[int]] = None,
                 deny_only: Optional[pulumi.Input[bool]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyExpiry resources.
        :param pulumi.Input[bool] deny_only: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)
        if deny_only is not None:
            pulumi.set(__self__, "deny_only", deny_only)
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="denyOnly")
    def deny_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "deny_only")

    @deny_only.setter
    def deny_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deny_only", value)

    @property
    @pulumi.getter(name="executionLogging")
    def execution_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "execution_logging")

    @execution_logging.setter
    def execution_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "execution_logging", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class PolicyExpiry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 days: Optional[pulumi.Input[int]] = None,
                 deny_only: Optional[pulumi.Input[bool]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create expiry policy
        name = authentik.PolicyExpiry("name",
            name="expiry",
            days=3)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deny_only: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyExpiryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create expiry policy
        name = authentik.PolicyExpiry("name",
            name="expiry",
            days=3)
        ```

        :param str resource_name: The name of the resource.
        :param PolicyExpiryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyExpiryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 days: Optional[pulumi.Input[int]] = None,
                 deny_only: Optional[pulumi.Input[bool]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyExpiryArgs.__new__(PolicyExpiryArgs)

            if days is None and not opts.urn:
                raise TypeError("Missing required property 'days'")
            __props__.__dict__["days"] = days
            __props__.__dict__["deny_only"] = deny_only
            __props__.__dict__["execution_logging"] = execution_logging
            __props__.__dict__["name"] = name
        super(PolicyExpiry, __self__).__init__(
            'authentik:index/policyExpiry:PolicyExpiry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            days: Optional[pulumi.Input[int]] = None,
            deny_only: Optional[pulumi.Input[bool]] = None,
            execution_logging: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'PolicyExpiry':
        """
        Get an existing PolicyExpiry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deny_only: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyExpiryState.__new__(_PolicyExpiryState)

        __props__.__dict__["days"] = days
        __props__.__dict__["deny_only"] = deny_only
        __props__.__dict__["execution_logging"] = execution_logging
        __props__.__dict__["name"] = name
        return PolicyExpiry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Output[int]:
        return pulumi.get(self, "days")

    @property
    @pulumi.getter(name="denyOnly")
    def deny_only(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "deny_only")

    @property
    @pulumi.getter(name="executionLogging")
    def execution_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "execution_logging")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

