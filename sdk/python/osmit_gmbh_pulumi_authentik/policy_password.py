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

__all__ = ['PolicyPasswordArgs', 'PolicyPassword']

@pulumi.input_type
class PolicyPasswordArgs:
    def __init__(__self__, *,
                 error_message: pulumi.Input[str],
                 amount_digits: Optional[pulumi.Input[int]] = None,
                 amount_lowercase: Optional[pulumi.Input[int]] = None,
                 amount_symbols: Optional[pulumi.Input[int]] = None,
                 amount_uppercase: Optional[pulumi.Input[int]] = None,
                 check_have_i_been_pwned: Optional[pulumi.Input[bool]] = None,
                 check_static_rules: Optional[pulumi.Input[bool]] = None,
                 check_zxcvbn: Optional[pulumi.Input[bool]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 hibp_allowed_count: Optional[pulumi.Input[int]] = None,
                 length_min: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password_field: Optional[pulumi.Input[str]] = None,
                 symbol_charset: Optional[pulumi.Input[str]] = None,
                 zxcvbn_score_threshold: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PolicyPassword resource.
        :param pulumi.Input[bool] check_have_i_been_pwned: Defaults to `false`.
        :param pulumi.Input[bool] check_static_rules: Defaults to `true`.
        :param pulumi.Input[bool] check_zxcvbn: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[int] hibp_allowed_count: Defaults to `1`.
        :param pulumi.Input[str] password_field: Defaults to `password`.
        :param pulumi.Input[str] symbol_charset: Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        :param pulumi.Input[int] zxcvbn_score_threshold: Defaults to `2`.
        """
        pulumi.set(__self__, "error_message", error_message)
        if amount_digits is not None:
            pulumi.set(__self__, "amount_digits", amount_digits)
        if amount_lowercase is not None:
            pulumi.set(__self__, "amount_lowercase", amount_lowercase)
        if amount_symbols is not None:
            pulumi.set(__self__, "amount_symbols", amount_symbols)
        if amount_uppercase is not None:
            pulumi.set(__self__, "amount_uppercase", amount_uppercase)
        if check_have_i_been_pwned is not None:
            pulumi.set(__self__, "check_have_i_been_pwned", check_have_i_been_pwned)
        if check_static_rules is not None:
            pulumi.set(__self__, "check_static_rules", check_static_rules)
        if check_zxcvbn is not None:
            pulumi.set(__self__, "check_zxcvbn", check_zxcvbn)
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if hibp_allowed_count is not None:
            pulumi.set(__self__, "hibp_allowed_count", hibp_allowed_count)
        if length_min is not None:
            pulumi.set(__self__, "length_min", length_min)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password_field is not None:
            pulumi.set(__self__, "password_field", password_field)
        if symbol_charset is not None:
            pulumi.set(__self__, "symbol_charset", symbol_charset)
        if zxcvbn_score_threshold is not None:
            pulumi.set(__self__, "zxcvbn_score_threshold", zxcvbn_score_threshold)

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> pulumi.Input[str]:
        return pulumi.get(self, "error_message")

    @error_message.setter
    def error_message(self, value: pulumi.Input[str]):
        pulumi.set(self, "error_message", value)

    @property
    @pulumi.getter(name="amountDigits")
    def amount_digits(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_digits")

    @amount_digits.setter
    def amount_digits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_digits", value)

    @property
    @pulumi.getter(name="amountLowercase")
    def amount_lowercase(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_lowercase")

    @amount_lowercase.setter
    def amount_lowercase(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_lowercase", value)

    @property
    @pulumi.getter(name="amountSymbols")
    def amount_symbols(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_symbols")

    @amount_symbols.setter
    def amount_symbols(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_symbols", value)

    @property
    @pulumi.getter(name="amountUppercase")
    def amount_uppercase(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_uppercase")

    @amount_uppercase.setter
    def amount_uppercase(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_uppercase", value)

    @property
    @pulumi.getter(name="checkHaveIBeenPwned")
    def check_have_i_been_pwned(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_have_i_been_pwned")

    @check_have_i_been_pwned.setter
    def check_have_i_been_pwned(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_have_i_been_pwned", value)

    @property
    @pulumi.getter(name="checkStaticRules")
    def check_static_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "check_static_rules")

    @check_static_rules.setter
    def check_static_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_static_rules", value)

    @property
    @pulumi.getter(name="checkZxcvbn")
    def check_zxcvbn(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_zxcvbn")

    @check_zxcvbn.setter
    def check_zxcvbn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_zxcvbn", value)

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
    @pulumi.getter(name="hibpAllowedCount")
    def hibp_allowed_count(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "hibp_allowed_count")

    @hibp_allowed_count.setter
    def hibp_allowed_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hibp_allowed_count", value)

    @property
    @pulumi.getter(name="lengthMin")
    def length_min(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "length_min")

    @length_min.setter
    def length_min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "length_min", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="passwordField")
    def password_field(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `password`.
        """
        return pulumi.get(self, "password_field")

    @password_field.setter
    def password_field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_field", value)

    @property
    @pulumi.getter(name="symbolCharset")
    def symbol_charset(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        """
        return pulumi.get(self, "symbol_charset")

    @symbol_charset.setter
    def symbol_charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "symbol_charset", value)

    @property
    @pulumi.getter(name="zxcvbnScoreThreshold")
    def zxcvbn_score_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `2`.
        """
        return pulumi.get(self, "zxcvbn_score_threshold")

    @zxcvbn_score_threshold.setter
    def zxcvbn_score_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "zxcvbn_score_threshold", value)


@pulumi.input_type
class _PolicyPasswordState:
    def __init__(__self__, *,
                 amount_digits: Optional[pulumi.Input[int]] = None,
                 amount_lowercase: Optional[pulumi.Input[int]] = None,
                 amount_symbols: Optional[pulumi.Input[int]] = None,
                 amount_uppercase: Optional[pulumi.Input[int]] = None,
                 check_have_i_been_pwned: Optional[pulumi.Input[bool]] = None,
                 check_static_rules: Optional[pulumi.Input[bool]] = None,
                 check_zxcvbn: Optional[pulumi.Input[bool]] = None,
                 error_message: Optional[pulumi.Input[str]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 hibp_allowed_count: Optional[pulumi.Input[int]] = None,
                 length_min: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password_field: Optional[pulumi.Input[str]] = None,
                 symbol_charset: Optional[pulumi.Input[str]] = None,
                 zxcvbn_score_threshold: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PolicyPassword resources.
        :param pulumi.Input[bool] check_have_i_been_pwned: Defaults to `false`.
        :param pulumi.Input[bool] check_static_rules: Defaults to `true`.
        :param pulumi.Input[bool] check_zxcvbn: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[int] hibp_allowed_count: Defaults to `1`.
        :param pulumi.Input[str] password_field: Defaults to `password`.
        :param pulumi.Input[str] symbol_charset: Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        :param pulumi.Input[int] zxcvbn_score_threshold: Defaults to `2`.
        """
        if amount_digits is not None:
            pulumi.set(__self__, "amount_digits", amount_digits)
        if amount_lowercase is not None:
            pulumi.set(__self__, "amount_lowercase", amount_lowercase)
        if amount_symbols is not None:
            pulumi.set(__self__, "amount_symbols", amount_symbols)
        if amount_uppercase is not None:
            pulumi.set(__self__, "amount_uppercase", amount_uppercase)
        if check_have_i_been_pwned is not None:
            pulumi.set(__self__, "check_have_i_been_pwned", check_have_i_been_pwned)
        if check_static_rules is not None:
            pulumi.set(__self__, "check_static_rules", check_static_rules)
        if check_zxcvbn is not None:
            pulumi.set(__self__, "check_zxcvbn", check_zxcvbn)
        if error_message is not None:
            pulumi.set(__self__, "error_message", error_message)
        if execution_logging is not None:
            pulumi.set(__self__, "execution_logging", execution_logging)
        if hibp_allowed_count is not None:
            pulumi.set(__self__, "hibp_allowed_count", hibp_allowed_count)
        if length_min is not None:
            pulumi.set(__self__, "length_min", length_min)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password_field is not None:
            pulumi.set(__self__, "password_field", password_field)
        if symbol_charset is not None:
            pulumi.set(__self__, "symbol_charset", symbol_charset)
        if zxcvbn_score_threshold is not None:
            pulumi.set(__self__, "zxcvbn_score_threshold", zxcvbn_score_threshold)

    @property
    @pulumi.getter(name="amountDigits")
    def amount_digits(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_digits")

    @amount_digits.setter
    def amount_digits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_digits", value)

    @property
    @pulumi.getter(name="amountLowercase")
    def amount_lowercase(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_lowercase")

    @amount_lowercase.setter
    def amount_lowercase(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_lowercase", value)

    @property
    @pulumi.getter(name="amountSymbols")
    def amount_symbols(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_symbols")

    @amount_symbols.setter
    def amount_symbols(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_symbols", value)

    @property
    @pulumi.getter(name="amountUppercase")
    def amount_uppercase(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount_uppercase")

    @amount_uppercase.setter
    def amount_uppercase(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount_uppercase", value)

    @property
    @pulumi.getter(name="checkHaveIBeenPwned")
    def check_have_i_been_pwned(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_have_i_been_pwned")

    @check_have_i_been_pwned.setter
    def check_have_i_been_pwned(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_have_i_been_pwned", value)

    @property
    @pulumi.getter(name="checkStaticRules")
    def check_static_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "check_static_rules")

    @check_static_rules.setter
    def check_static_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_static_rules", value)

    @property
    @pulumi.getter(name="checkZxcvbn")
    def check_zxcvbn(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_zxcvbn")

    @check_zxcvbn.setter
    def check_zxcvbn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_zxcvbn", value)

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "error_message")

    @error_message.setter
    def error_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "error_message", value)

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
    @pulumi.getter(name="hibpAllowedCount")
    def hibp_allowed_count(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "hibp_allowed_count")

    @hibp_allowed_count.setter
    def hibp_allowed_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hibp_allowed_count", value)

    @property
    @pulumi.getter(name="lengthMin")
    def length_min(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "length_min")

    @length_min.setter
    def length_min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "length_min", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="passwordField")
    def password_field(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `password`.
        """
        return pulumi.get(self, "password_field")

    @password_field.setter
    def password_field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_field", value)

    @property
    @pulumi.getter(name="symbolCharset")
    def symbol_charset(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        """
        return pulumi.get(self, "symbol_charset")

    @symbol_charset.setter
    def symbol_charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "symbol_charset", value)

    @property
    @pulumi.getter(name="zxcvbnScoreThreshold")
    def zxcvbn_score_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Defaults to `2`.
        """
        return pulumi.get(self, "zxcvbn_score_threshold")

    @zxcvbn_score_threshold.setter
    def zxcvbn_score_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "zxcvbn_score_threshold", value)


class PolicyPassword(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount_digits: Optional[pulumi.Input[int]] = None,
                 amount_lowercase: Optional[pulumi.Input[int]] = None,
                 amount_symbols: Optional[pulumi.Input[int]] = None,
                 amount_uppercase: Optional[pulumi.Input[int]] = None,
                 check_have_i_been_pwned: Optional[pulumi.Input[bool]] = None,
                 check_static_rules: Optional[pulumi.Input[bool]] = None,
                 check_zxcvbn: Optional[pulumi.Input[bool]] = None,
                 error_message: Optional[pulumi.Input[str]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 hibp_allowed_count: Optional[pulumi.Input[int]] = None,
                 length_min: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password_field: Optional[pulumi.Input[str]] = None,
                 symbol_charset: Optional[pulumi.Input[str]] = None,
                 zxcvbn_score_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        name = authentik.PolicyPassword("name",
            error_message="foo",
            length_min=8)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] check_have_i_been_pwned: Defaults to `false`.
        :param pulumi.Input[bool] check_static_rules: Defaults to `true`.
        :param pulumi.Input[bool] check_zxcvbn: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[int] hibp_allowed_count: Defaults to `1`.
        :param pulumi.Input[str] password_field: Defaults to `password`.
        :param pulumi.Input[str] symbol_charset: Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        :param pulumi.Input[int] zxcvbn_score_threshold: Defaults to `2`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyPasswordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        name = authentik.PolicyPassword("name",
            error_message="foo",
            length_min=8)
        ```

        :param str resource_name: The name of the resource.
        :param PolicyPasswordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyPasswordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount_digits: Optional[pulumi.Input[int]] = None,
                 amount_lowercase: Optional[pulumi.Input[int]] = None,
                 amount_symbols: Optional[pulumi.Input[int]] = None,
                 amount_uppercase: Optional[pulumi.Input[int]] = None,
                 check_have_i_been_pwned: Optional[pulumi.Input[bool]] = None,
                 check_static_rules: Optional[pulumi.Input[bool]] = None,
                 check_zxcvbn: Optional[pulumi.Input[bool]] = None,
                 error_message: Optional[pulumi.Input[str]] = None,
                 execution_logging: Optional[pulumi.Input[bool]] = None,
                 hibp_allowed_count: Optional[pulumi.Input[int]] = None,
                 length_min: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password_field: Optional[pulumi.Input[str]] = None,
                 symbol_charset: Optional[pulumi.Input[str]] = None,
                 zxcvbn_score_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyPasswordArgs.__new__(PolicyPasswordArgs)

            __props__.__dict__["amount_digits"] = amount_digits
            __props__.__dict__["amount_lowercase"] = amount_lowercase
            __props__.__dict__["amount_symbols"] = amount_symbols
            __props__.__dict__["amount_uppercase"] = amount_uppercase
            __props__.__dict__["check_have_i_been_pwned"] = check_have_i_been_pwned
            __props__.__dict__["check_static_rules"] = check_static_rules
            __props__.__dict__["check_zxcvbn"] = check_zxcvbn
            if error_message is None and not opts.urn:
                raise TypeError("Missing required property 'error_message'")
            __props__.__dict__["error_message"] = error_message
            __props__.__dict__["execution_logging"] = execution_logging
            __props__.__dict__["hibp_allowed_count"] = hibp_allowed_count
            __props__.__dict__["length_min"] = length_min
            __props__.__dict__["name"] = name
            __props__.__dict__["password_field"] = password_field
            __props__.__dict__["symbol_charset"] = symbol_charset
            __props__.__dict__["zxcvbn_score_threshold"] = zxcvbn_score_threshold
        super(PolicyPassword, __self__).__init__(
            'authentik:index/policyPassword:PolicyPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            amount_digits: Optional[pulumi.Input[int]] = None,
            amount_lowercase: Optional[pulumi.Input[int]] = None,
            amount_symbols: Optional[pulumi.Input[int]] = None,
            amount_uppercase: Optional[pulumi.Input[int]] = None,
            check_have_i_been_pwned: Optional[pulumi.Input[bool]] = None,
            check_static_rules: Optional[pulumi.Input[bool]] = None,
            check_zxcvbn: Optional[pulumi.Input[bool]] = None,
            error_message: Optional[pulumi.Input[str]] = None,
            execution_logging: Optional[pulumi.Input[bool]] = None,
            hibp_allowed_count: Optional[pulumi.Input[int]] = None,
            length_min: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password_field: Optional[pulumi.Input[str]] = None,
            symbol_charset: Optional[pulumi.Input[str]] = None,
            zxcvbn_score_threshold: Optional[pulumi.Input[int]] = None) -> 'PolicyPassword':
        """
        Get an existing PolicyPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] check_have_i_been_pwned: Defaults to `false`.
        :param pulumi.Input[bool] check_static_rules: Defaults to `true`.
        :param pulumi.Input[bool] check_zxcvbn: Defaults to `false`.
        :param pulumi.Input[bool] execution_logging: Defaults to `false`.
        :param pulumi.Input[int] hibp_allowed_count: Defaults to `1`.
        :param pulumi.Input[str] password_field: Defaults to `password`.
        :param pulumi.Input[str] symbol_charset: Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        :param pulumi.Input[int] zxcvbn_score_threshold: Defaults to `2`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyPasswordState.__new__(_PolicyPasswordState)

        __props__.__dict__["amount_digits"] = amount_digits
        __props__.__dict__["amount_lowercase"] = amount_lowercase
        __props__.__dict__["amount_symbols"] = amount_symbols
        __props__.__dict__["amount_uppercase"] = amount_uppercase
        __props__.__dict__["check_have_i_been_pwned"] = check_have_i_been_pwned
        __props__.__dict__["check_static_rules"] = check_static_rules
        __props__.__dict__["check_zxcvbn"] = check_zxcvbn
        __props__.__dict__["error_message"] = error_message
        __props__.__dict__["execution_logging"] = execution_logging
        __props__.__dict__["hibp_allowed_count"] = hibp_allowed_count
        __props__.__dict__["length_min"] = length_min
        __props__.__dict__["name"] = name
        __props__.__dict__["password_field"] = password_field
        __props__.__dict__["symbol_charset"] = symbol_charset
        __props__.__dict__["zxcvbn_score_threshold"] = zxcvbn_score_threshold
        return PolicyPassword(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="amountDigits")
    def amount_digits(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "amount_digits")

    @property
    @pulumi.getter(name="amountLowercase")
    def amount_lowercase(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "amount_lowercase")

    @property
    @pulumi.getter(name="amountSymbols")
    def amount_symbols(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "amount_symbols")

    @property
    @pulumi.getter(name="amountUppercase")
    def amount_uppercase(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "amount_uppercase")

    @property
    @pulumi.getter(name="checkHaveIBeenPwned")
    def check_have_i_been_pwned(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_have_i_been_pwned")

    @property
    @pulumi.getter(name="checkStaticRules")
    def check_static_rules(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "check_static_rules")

    @property
    @pulumi.getter(name="checkZxcvbn")
    def check_zxcvbn(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "check_zxcvbn")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> pulumi.Output[str]:
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="executionLogging")
    def execution_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "execution_logging")

    @property
    @pulumi.getter(name="hibpAllowedCount")
    def hibp_allowed_count(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "hibp_allowed_count")

    @property
    @pulumi.getter(name="lengthMin")
    def length_min(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "length_min")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passwordField")
    def password_field(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `password`.
        """
        return pulumi.get(self, "password_field")

    @property
    @pulumi.getter(name="symbolCharset")
    def symbol_charset(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `!\\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~`.
        """
        return pulumi.get(self, "symbol_charset")

    @property
    @pulumi.getter(name="zxcvbnScoreThreshold")
    def zxcvbn_score_threshold(self) -> pulumi.Output[Optional[int]]:
        """
        Defaults to `2`.
        """
        return pulumi.get(self, "zxcvbn_score_threshold")

