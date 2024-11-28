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

__all__ = ['StagePromptFieldArgs', 'StagePromptField']

@pulumi.input_type
class StagePromptFieldArgs:
    def __init__(__self__, *,
                 field_key: pulumi.Input[str],
                 label: pulumi.Input[str],
                 type: pulumi.Input[str],
                 initial_value: Optional[pulumi.Input[str]] = None,
                 initial_value_expression: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 placeholder: Optional[pulumi.Input[str]] = None,
                 placeholder_expression: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sub_text: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StagePromptField resource.
        :param pulumi.Input[str] type: Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
               `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
               `static` - `ak-locale`
        """
        pulumi.set(__self__, "field_key", field_key)
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "type", type)
        if initial_value is not None:
            pulumi.set(__self__, "initial_value", initial_value)
        if initial_value_expression is not None:
            pulumi.set(__self__, "initial_value_expression", initial_value_expression)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if placeholder is not None:
            pulumi.set(__self__, "placeholder", placeholder)
        if placeholder_expression is not None:
            pulumi.set(__self__, "placeholder_expression", placeholder_expression)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sub_text is not None:
            pulumi.set(__self__, "sub_text", sub_text)

    @property
    @pulumi.getter(name="fieldKey")
    def field_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "field_key")

    @field_key.setter
    def field_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "field_key", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
        `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
        `static` - `ak-locale`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="initialValue")
    def initial_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "initial_value")

    @initial_value.setter
    def initial_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial_value", value)

    @property
    @pulumi.getter(name="initialValueExpression")
    def initial_value_expression(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "initial_value_expression")

    @initial_value_expression.setter
    def initial_value_expression(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "initial_value_expression", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def placeholder(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "placeholder")

    @placeholder.setter
    def placeholder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "placeholder", value)

    @property
    @pulumi.getter(name="placeholderExpression")
    def placeholder_expression(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "placeholder_expression")

    @placeholder_expression.setter
    def placeholder_expression(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "placeholder_expression", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter(name="subText")
    def sub_text(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sub_text")

    @sub_text.setter
    def sub_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_text", value)


@pulumi.input_type
class _StagePromptFieldState:
    def __init__(__self__, *,
                 field_key: Optional[pulumi.Input[str]] = None,
                 initial_value: Optional[pulumi.Input[str]] = None,
                 initial_value_expression: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 placeholder: Optional[pulumi.Input[str]] = None,
                 placeholder_expression: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sub_text: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StagePromptField resources.
        :param pulumi.Input[str] type: Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
               `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
               `static` - `ak-locale`
        """
        if field_key is not None:
            pulumi.set(__self__, "field_key", field_key)
        if initial_value is not None:
            pulumi.set(__self__, "initial_value", initial_value)
        if initial_value_expression is not None:
            pulumi.set(__self__, "initial_value_expression", initial_value_expression)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if placeholder is not None:
            pulumi.set(__self__, "placeholder", placeholder)
        if placeholder_expression is not None:
            pulumi.set(__self__, "placeholder_expression", placeholder_expression)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sub_text is not None:
            pulumi.set(__self__, "sub_text", sub_text)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="fieldKey")
    def field_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "field_key")

    @field_key.setter
    def field_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_key", value)

    @property
    @pulumi.getter(name="initialValue")
    def initial_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "initial_value")

    @initial_value.setter
    def initial_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial_value", value)

    @property
    @pulumi.getter(name="initialValueExpression")
    def initial_value_expression(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "initial_value_expression")

    @initial_value_expression.setter
    def initial_value_expression(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "initial_value_expression", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def placeholder(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "placeholder")

    @placeholder.setter
    def placeholder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "placeholder", value)

    @property
    @pulumi.getter(name="placeholderExpression")
    def placeholder_expression(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "placeholder_expression")

    @placeholder_expression.setter
    def placeholder_expression(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "placeholder_expression", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter(name="subText")
    def sub_text(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sub_text")

    @sub_text.setter
    def sub_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_text", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
        `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
        `static` - `ak-locale`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class StagePromptField(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 field_key: Optional[pulumi.Input[str]] = None,
                 initial_value: Optional[pulumi.Input[str]] = None,
                 initial_value_expression: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 placeholder: Optional[pulumi.Input[str]] = None,
                 placeholder_expression: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sub_text: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a prompt field
        field = authentik.StagePromptField("field",
            field_key="username",
            label="Username",
            type="username")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] type: Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
               `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
               `static` - `ak-locale`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StagePromptFieldArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a prompt field
        field = authentik.StagePromptField("field",
            field_key="username",
            label="Username",
            type="username")
        ```

        :param str resource_name: The name of the resource.
        :param StagePromptFieldArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StagePromptFieldArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 field_key: Optional[pulumi.Input[str]] = None,
                 initial_value: Optional[pulumi.Input[str]] = None,
                 initial_value_expression: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 placeholder: Optional[pulumi.Input[str]] = None,
                 placeholder_expression: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sub_text: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StagePromptFieldArgs.__new__(StagePromptFieldArgs)

            if field_key is None and not opts.urn:
                raise TypeError("Missing required property 'field_key'")
            __props__.__dict__["field_key"] = field_key
            __props__.__dict__["initial_value"] = initial_value
            __props__.__dict__["initial_value_expression"] = initial_value_expression
            if label is None and not opts.urn:
                raise TypeError("Missing required property 'label'")
            __props__.__dict__["label"] = label
            __props__.__dict__["name"] = name
            __props__.__dict__["order"] = order
            __props__.__dict__["placeholder"] = placeholder
            __props__.__dict__["placeholder_expression"] = placeholder_expression
            __props__.__dict__["required"] = required
            __props__.__dict__["sub_text"] = sub_text
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(StagePromptField, __self__).__init__(
            'authentik:index/stagePromptField:StagePromptField',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            field_key: Optional[pulumi.Input[str]] = None,
            initial_value: Optional[pulumi.Input[str]] = None,
            initial_value_expression: Optional[pulumi.Input[bool]] = None,
            label: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            order: Optional[pulumi.Input[int]] = None,
            placeholder: Optional[pulumi.Input[str]] = None,
            placeholder_expression: Optional[pulumi.Input[bool]] = None,
            required: Optional[pulumi.Input[bool]] = None,
            sub_text: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'StagePromptField':
        """
        Get an existing StagePromptField resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] type: Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
               `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
               `static` - `ak-locale`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StagePromptFieldState.__new__(_StagePromptFieldState)

        __props__.__dict__["field_key"] = field_key
        __props__.__dict__["initial_value"] = initial_value
        __props__.__dict__["initial_value_expression"] = initial_value_expression
        __props__.__dict__["label"] = label
        __props__.__dict__["name"] = name
        __props__.__dict__["order"] = order
        __props__.__dict__["placeholder"] = placeholder
        __props__.__dict__["placeholder_expression"] = placeholder_expression
        __props__.__dict__["required"] = required
        __props__.__dict__["sub_text"] = sub_text
        __props__.__dict__["type"] = type
        return StagePromptField(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fieldKey")
    def field_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "field_key")

    @property
    @pulumi.getter(name="initialValue")
    def initial_value(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "initial_value")

    @property
    @pulumi.getter(name="initialValueExpression")
    def initial_value_expression(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "initial_value_expression")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def placeholder(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "placeholder")

    @property
    @pulumi.getter(name="placeholderExpression")
    def placeholder_expression(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "placeholder_expression")

    @property
    @pulumi.getter
    def required(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "required")

    @property
    @pulumi.getter(name="subText")
    def sub_text(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sub_text")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Allowed values: - `text` - `text_area` - `text_read_only` - `text_area_read_only` - `username` - `email` - `password` -
        `number` - `checkbox` - `radio-button-group` - `dropdown` - `date` - `date-time` - `file` - `separator` - `hidden` -
        `static` - `ak-locale`
        """
        return pulumi.get(self, "type")
