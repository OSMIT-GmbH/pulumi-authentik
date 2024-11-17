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

__all__ = ['PolicyDummyArgs', 'PolicyDummy']

@pulumi.input_type
class PolicyDummyArgs:
    def __init__(__self__, *,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[bool]] = None,
                 wait_max: Optional[pulumi.Input[int]] = None,
                 wait_min: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PolicyDummy resource.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[bool] result: Defaults to `false`.
        :param pulumi.Input[int] wait_max: Defaults to `30`.
        :param pulumi.Input[int] wait_min: Defaults to `5`.
        """
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if result is not None:
            pulumi.set(__self__, "result", result)
        if wait_max is not None:
            pulumi.set(__self__, "wait_max", wait_max)
        if wait_min is not None:
            pulumi.set(__self__, "wait_min", wait_min)

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

    @property
    @pulumi.getter
    def result(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "result")

    @result.setter
    def result(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "result", value)

    @property
    @pulumi.getter(name="waitMax")
    def wait_max(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "wait_max")

    @wait_max.setter
    def wait_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_max", value)

    @property
    @pulumi.getter(name="waitMin")
    def wait_min(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `5`.
        """
        return pulumi.get(self, "wait_min")

    @wait_min.setter
    def wait_min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_min", value)


@pulumi.input_type
class _PolicyDummyState:
    def __init__(__self__, *,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[bool]] = None,
                 wait_max: Optional[pulumi.Input[int]] = None,
                 wait_min: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PolicyDummy resources.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[bool] result: Defaults to `false`.
        :param pulumi.Input[int] wait_max: Defaults to `30`.
        :param pulumi.Input[int] wait_min: Defaults to `5`.
        """
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if result is not None:
            pulumi.set(__self__, "result", result)
        if wait_max is not None:
            pulumi.set(__self__, "wait_max", wait_max)
        if wait_min is not None:
            pulumi.set(__self__, "wait_min", wait_min)

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

    @property
    @pulumi.getter
    def result(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "result")

    @result.setter
    def result(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "result", value)

    @property
    @pulumi.getter(name="waitMax")
    def wait_max(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "wait_max")

    @wait_max.setter
    def wait_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_max", value)

    @property
    @pulumi.getter(name="waitMin")
    def wait_min(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `5`.
        """
        return pulumi.get(self, "wait_min")

    @wait_min.setter
    def wait_min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_min", value)


class PolicyDummy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[bool]] = None,
                 wait_max: Optional[pulumi.Input[int]] = None,
                 wait_min: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        name = authentik.PolicyDummy("name")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[bool] result: Defaults to `false`.
        :param pulumi.Input[int] wait_max: Defaults to `30`.
        :param pulumi.Input[int] wait_min: Defaults to `5`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyDummyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        name = authentik.PolicyDummy("name")
        ```

        :param str resource_name: The name of the resource.
        :param PolicyDummyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyDummyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[bool]] = None,
                 wait_max: Optional[pulumi.Input[int]] = None,
                 wait_min: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyDummyArgs.__new__(PolicyDummyArgs)

            __props__.__dict__["execution_logging"] = execution_logging
            __props__.__dict__["name"] = name
            __props__.__dict__["result"] = result
            __props__.__dict__["wait_max"] = wait_max
            __props__.__dict__["wait_min"] = wait_min
        super(PolicyDummy, __self__).__init__(
            'authentik:index/policyDummy:PolicyDummy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            execution_logging: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            result: Optional[pulumi.Input[bool]] = None,
            wait_max: Optional[pulumi.Input[int]] = None,
            wait_min: Optional[pulumi.Input[int]] = None) -> 'PolicyDummy':
        """
        Get an existing PolicyDummy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[bool] result: Defaults to `false`.
        :param pulumi.Input[int] wait_max: Defaults to `30`.
        :param pulumi.Input[int] wait_min: Defaults to `5`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyDummyState.__new__(_PolicyDummyState)

        __props__.__dict__["execution_logging"] = execution_logging
        __props__.__dict__["name"] = name
        __props__.__dict__["result"] = result
        __props__.__dict__["wait_max"] = wait_max
        __props__.__dict__["wait_min"] = wait_min
        return PolicyDummy(resource_name, opts=opts, __props__=__props__)

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

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="waitMax")
    def wait_max(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `30`.
        """
        return pulumi.get(self, "wait_max")

    @property
    @pulumi.getter(name="waitMin")
    def wait_min(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `5`.
        """
        return pulumi.get(self, "wait_min")

