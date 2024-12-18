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

__all__ = ['StageCaptchaArgs', 'StageCaptcha']

@pulumi.input_type
class StageCaptchaArgs:
    def __init__(__self__, *,
                 private_key: pulumi.Input[str],
                 public_key: pulumi.Input[str],
                 api_url: Optional[pulumi.Input[str]] = None,
                 error_on_invalid_score: Optional[pulumi.Input[bool]] = None,
                 interactive: Optional[pulumi.Input[bool]] = None,
                 js_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 score_max_threshold: Optional[pulumi.Input[float]] = None,
                 score_min_threshold: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a StageCaptcha resource.
        :param pulumi.Input[str] api_url: Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        :param pulumi.Input[bool] error_on_invalid_score: Defaults to `true`.
        :param pulumi.Input[bool] interactive: Defaults to `false`.
        :param pulumi.Input[str] js_url: Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        :param pulumi.Input[float] score_max_threshold: Defaults to `0.5`.
        :param pulumi.Input[float] score_min_threshold: Defaults to `1`.
        """
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "public_key", public_key)
        if api_url is not None:
            pulumi.set(__self__, "api_url", api_url)
        if error_on_invalid_score is not None:
            pulumi.set(__self__, "error_on_invalid_score", error_on_invalid_score)
        if interactive is not None:
            pulumi.set(__self__, "interactive", interactive)
        if js_url is not None:
            pulumi.set(__self__, "js_url", js_url)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if score_max_threshold is not None:
            pulumi.set(__self__, "score_max_threshold", score_max_threshold)
        if score_min_threshold is not None:
            pulumi.set(__self__, "score_min_threshold", score_min_threshold)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        """
        return pulumi.get(self, "api_url")

    @api_url.setter
    def api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_url", value)

    @property
    @pulumi.getter(name="errorOnInvalidScore")
    def error_on_invalid_score(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "error_on_invalid_score")

    @error_on_invalid_score.setter
    def error_on_invalid_score(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "error_on_invalid_score", value)

    @property
    @pulumi.getter
    def interactive(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "interactive")

    @interactive.setter
    def interactive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "interactive", value)

    @property
    @pulumi.getter(name="jsUrl")
    def js_url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        """
        return pulumi.get(self, "js_url")

    @js_url.setter
    def js_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "js_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="scoreMaxThreshold")
    def score_max_threshold(self) -> Optional[pulumi.Input[float]]:
        """
        Defaults to `0.5`.
        """
        return pulumi.get(self, "score_max_threshold")

    @score_max_threshold.setter
    def score_max_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "score_max_threshold", value)

    @property
    @pulumi.getter(name="scoreMinThreshold")
    def score_min_threshold(self) -> Optional[pulumi.Input[float]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "score_min_threshold")

    @score_min_threshold.setter
    def score_min_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "score_min_threshold", value)


@pulumi.input_type
class _StageCaptchaState:
    def __init__(__self__, *,
                 api_url: Optional[pulumi.Input[str]] = None,
                 error_on_invalid_score: Optional[pulumi.Input[bool]] = None,
                 interactive: Optional[pulumi.Input[bool]] = None,
                 js_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 score_max_threshold: Optional[pulumi.Input[float]] = None,
                 score_min_threshold: Optional[pulumi.Input[float]] = None):
        """
        Input properties used for looking up and filtering StageCaptcha resources.
        :param pulumi.Input[str] api_url: Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        :param pulumi.Input[bool] error_on_invalid_score: Defaults to `true`.
        :param pulumi.Input[bool] interactive: Defaults to `false`.
        :param pulumi.Input[str] js_url: Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        :param pulumi.Input[float] score_max_threshold: Defaults to `0.5`.
        :param pulumi.Input[float] score_min_threshold: Defaults to `1`.
        """
        if api_url is not None:
            pulumi.set(__self__, "api_url", api_url)
        if error_on_invalid_score is not None:
            pulumi.set(__self__, "error_on_invalid_score", error_on_invalid_score)
        if interactive is not None:
            pulumi.set(__self__, "interactive", interactive)
        if js_url is not None:
            pulumi.set(__self__, "js_url", js_url)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if score_max_threshold is not None:
            pulumi.set(__self__, "score_max_threshold", score_max_threshold)
        if score_min_threshold is not None:
            pulumi.set(__self__, "score_min_threshold", score_min_threshold)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        """
        return pulumi.get(self, "api_url")

    @api_url.setter
    def api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_url", value)

    @property
    @pulumi.getter(name="errorOnInvalidScore")
    def error_on_invalid_score(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "error_on_invalid_score")

    @error_on_invalid_score.setter
    def error_on_invalid_score(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "error_on_invalid_score", value)

    @property
    @pulumi.getter
    def interactive(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "interactive")

    @interactive.setter
    def interactive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "interactive", value)

    @property
    @pulumi.getter(name="jsUrl")
    def js_url(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        """
        return pulumi.get(self, "js_url")

    @js_url.setter
    def js_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "js_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="scoreMaxThreshold")
    def score_max_threshold(self) -> Optional[pulumi.Input[float]]:
        """
        Defaults to `0.5`.
        """
        return pulumi.get(self, "score_max_threshold")

    @score_max_threshold.setter
    def score_max_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "score_max_threshold", value)

    @property
    @pulumi.getter(name="scoreMinThreshold")
    def score_min_threshold(self) -> Optional[pulumi.Input[float]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "score_min_threshold")

    @score_min_threshold.setter
    def score_min_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "score_min_threshold", value)


class StageCaptcha(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_url: Optional[pulumi.Input[str]] = None,
                 error_on_invalid_score: Optional[pulumi.Input[bool]] = None,
                 interactive: Optional[pulumi.Input[bool]] = None,
                 js_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 score_max_threshold: Optional[pulumi.Input[float]] = None,
                 score_min_threshold: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a captcha stage
        name = authentik.StageCaptcha("name",
            name="captcha",
            private_key="foo",
            public_key="bar")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_url: Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        :param pulumi.Input[bool] error_on_invalid_score: Defaults to `true`.
        :param pulumi.Input[bool] interactive: Defaults to `false`.
        :param pulumi.Input[str] js_url: Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        :param pulumi.Input[float] score_max_threshold: Defaults to `0.5`.
        :param pulumi.Input[float] score_min_threshold: Defaults to `1`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StageCaptchaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a captcha stage
        name = authentik.StageCaptcha("name",
            name="captcha",
            private_key="foo",
            public_key="bar")
        ```

        :param str resource_name: The name of the resource.
        :param StageCaptchaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageCaptchaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_url: Optional[pulumi.Input[str]] = None,
                 error_on_invalid_score: Optional[pulumi.Input[bool]] = None,
                 interactive: Optional[pulumi.Input[bool]] = None,
                 js_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 score_max_threshold: Optional[pulumi.Input[float]] = None,
                 score_min_threshold: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageCaptchaArgs.__new__(StageCaptchaArgs)

            __props__.__dict__["api_url"] = api_url
            __props__.__dict__["error_on_invalid_score"] = error_on_invalid_score
            __props__.__dict__["interactive"] = interactive
            __props__.__dict__["js_url"] = js_url
            __props__.__dict__["name"] = name
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["score_max_threshold"] = score_max_threshold
            __props__.__dict__["score_min_threshold"] = score_min_threshold
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(StageCaptcha, __self__).__init__(
            'authentik:index/stageCaptcha:StageCaptcha',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_url: Optional[pulumi.Input[str]] = None,
            error_on_invalid_score: Optional[pulumi.Input[bool]] = None,
            interactive: Optional[pulumi.Input[bool]] = None,
            js_url: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            score_max_threshold: Optional[pulumi.Input[float]] = None,
            score_min_threshold: Optional[pulumi.Input[float]] = None) -> 'StageCaptcha':
        """
        Get an existing StageCaptcha resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_url: Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        :param pulumi.Input[bool] error_on_invalid_score: Defaults to `true`.
        :param pulumi.Input[bool] interactive: Defaults to `false`.
        :param pulumi.Input[str] js_url: Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        :param pulumi.Input[float] score_max_threshold: Defaults to `0.5`.
        :param pulumi.Input[float] score_min_threshold: Defaults to `1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageCaptchaState.__new__(_StageCaptchaState)

        __props__.__dict__["api_url"] = api_url
        __props__.__dict__["error_on_invalid_score"] = error_on_invalid_score
        __props__.__dict__["interactive"] = interactive
        __props__.__dict__["js_url"] = js_url
        __props__.__dict__["name"] = name
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["score_max_threshold"] = score_max_threshold
        __props__.__dict__["score_min_threshold"] = score_min_threshold
        return StageCaptcha(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
        """
        return pulumi.get(self, "api_url")

    @property
    @pulumi.getter(name="errorOnInvalidScore")
    def error_on_invalid_score(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "error_on_invalid_score")

    @property
    @pulumi.getter
    def interactive(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "interactive")

    @property
    @pulumi.getter(name="jsUrl")
    def js_url(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
        """
        return pulumi.get(self, "js_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="scoreMaxThreshold")
    def score_max_threshold(self) -> pulumi.Output[Optional[float]]:
        """
        Defaults to `0.5`.
        """
        return pulumi.get(self, "score_max_threshold")

    @property
    @pulumi.getter(name="scoreMinThreshold")
    def score_min_threshold(self) -> pulumi.Output[Optional[float]]:
        """
        Defaults to `1`.
        """
        return pulumi.get(self, "score_min_threshold")

