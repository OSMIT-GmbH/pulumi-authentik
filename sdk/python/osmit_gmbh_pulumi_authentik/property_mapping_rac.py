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

__all__ = ['PropertyMappingRacArgs', 'PropertyMappingRac']

@pulumi.input_type
class PropertyMappingRacArgs:
    def __init__(__self__, *,
                 expression: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PropertyMappingRac resource.
        :param pulumi.Input[str] settings: JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[str]]:
        """
        JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "settings", value)


@pulumi.input_type
class _PropertyMappingRacState:
    def __init__(__self__, *,
                 expression: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PropertyMappingRac resources.
        :param pulumi.Input[str] settings: JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[str]]:
        """
        JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "settings", value)


class PropertyMappingRac(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage RAC Provider Property mappings

        > This resource is deprecated. Migrate to `PropertyMappingProviderRac`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] settings: JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PropertyMappingRacArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage RAC Provider Property mappings

        > This resource is deprecated. Migrate to `PropertyMappingProviderRac`.

        :param str resource_name: The name of the resource.
        :param PropertyMappingRacArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PropertyMappingRacArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PropertyMappingRacArgs.__new__(PropertyMappingRacArgs)

            __props__.__dict__["expression"] = expression
            __props__.__dict__["name"] = name
            __props__.__dict__["settings"] = settings
        super(PropertyMappingRac, __self__).__init__(
            'authentik:index/propertyMappingRac:PropertyMappingRac',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expression: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            settings: Optional[pulumi.Input[str]] = None) -> 'PropertyMappingRac':
        """
        Get an existing PropertyMappingRac resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] settings: JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PropertyMappingRacState.__new__(_PropertyMappingRacState)

        __props__.__dict__["expression"] = expression
        __props__.__dict__["name"] = name
        __props__.__dict__["settings"] = settings
        return PropertyMappingRac(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output[Optional[str]]:
        """
        JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
        """
        return pulumi.get(self, "settings")

