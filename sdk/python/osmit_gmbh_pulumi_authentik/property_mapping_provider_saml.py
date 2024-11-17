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

__all__ = ['PropertyMappingProviderSamlArgs', 'PropertyMappingProviderSaml']

@pulumi.input_type
class PropertyMappingProviderSamlArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 saml_name: pulumi.Input[str],
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PropertyMappingProviderSaml resource.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "saml_name", saml_name)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter(name="samlName")
    def saml_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "saml_name")

    @saml_name.setter
    def saml_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "saml_name", value)

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
class _PropertyMappingProviderSamlState:
    def __init__(__self__, *,
                 expression: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 saml_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PropertyMappingProviderSaml resources.
        """
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if saml_name is not None:
            pulumi.set(__self__, "saml_name", saml_name)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

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
    @pulumi.getter(name="samlName")
    def saml_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "saml_name")

    @saml_name.setter
    def saml_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "saml_name", value)


class PropertyMappingProviderSaml(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 saml_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage SAML Provider Property mappings

        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        saml_aws_rolessessionname = authentik.PropertyMappingProviderSaml("saml-aws-rolessessionname",
            expression="return user.email",
            saml_name="https://aws.amazon.com/SAML/Attributes/RoleSessionName")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PropertyMappingProviderSamlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage SAML Provider Property mappings

        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        saml_aws_rolessessionname = authentik.PropertyMappingProviderSaml("saml-aws-rolessessionname",
            expression="return user.email",
            saml_name="https://aws.amazon.com/SAML/Attributes/RoleSessionName")
        ```

        :param str resource_name: The name of the resource.
        :param PropertyMappingProviderSamlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PropertyMappingProviderSamlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 saml_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PropertyMappingProviderSamlArgs.__new__(PropertyMappingProviderSamlArgs)

            if expression is None and not opts.urn:
                raise TypeError("Missing required property 'expression'")
            __props__.__dict__["expression"] = expression
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["name"] = name
            if saml_name is None and not opts.urn:
                raise TypeError("Missing required property 'saml_name'")
            __props__.__dict__["saml_name"] = saml_name
        super(PropertyMappingProviderSaml, __self__).__init__(
            'authentik:index/propertyMappingProviderSaml:PropertyMappingProviderSaml',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expression: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            saml_name: Optional[pulumi.Input[str]] = None) -> 'PropertyMappingProviderSaml':
        """
        Get an existing PropertyMappingProviderSaml resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PropertyMappingProviderSamlState.__new__(_PropertyMappingProviderSamlState)

        __props__.__dict__["expression"] = expression
        __props__.__dict__["friendly_name"] = friendly_name
        __props__.__dict__["name"] = name
        __props__.__dict__["saml_name"] = saml_name
        return PropertyMappingProviderSaml(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Output[str]:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="samlName")
    def saml_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "saml_name")

