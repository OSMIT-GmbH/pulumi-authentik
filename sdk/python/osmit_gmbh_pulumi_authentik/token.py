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

__all__ = ['TokenArgs', 'Token']

@pulumi.input_type
class TokenArgs:
    def __init__(__self__, *,
                 identifier: pulumi.Input[str],
                 user: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 expiring: Optional[pulumi.Input[bool]] = None,
                 intent: Optional[pulumi.Input[str]] = None,
                 retrieve_key: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Token resource.
        :param pulumi.Input[str] intent: Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        pulumi.set(__self__, "identifier", identifier)
        pulumi.set(__self__, "user", user)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if expiring is not None:
            pulumi.set(__self__, "expiring", expiring)
        if intent is not None:
            pulumi.set(__self__, "intent", intent)
        if retrieve_key is not None:
            pulumi.set(__self__, "retrieve_key", retrieve_key)

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[int]:
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[int]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter
    def expiring(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "expiring")

    @expiring.setter
    def expiring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "expiring", value)

    @property
    @pulumi.getter
    def intent(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        return pulumi.get(self, "intent")

    @intent.setter
    def intent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intent", value)

    @property
    @pulumi.getter(name="retrieveKey")
    def retrieve_key(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retrieve_key")

    @retrieve_key.setter
    def retrieve_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retrieve_key", value)


@pulumi.input_type
class _TokenState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 expires_in: Optional[pulumi.Input[int]] = None,
                 expiring: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 intent: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retrieve_key: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Token resources.
        :param pulumi.Input[str] intent: Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if expires_in is not None:
            pulumi.set(__self__, "expires_in", expires_in)
        if expiring is not None:
            pulumi.set(__self__, "expiring", expiring)
        if identifier is not None:
            pulumi.set(__self__, "identifier", identifier)
        if intent is not None:
            pulumi.set(__self__, "intent", intent)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if retrieve_key is not None:
            pulumi.set(__self__, "retrieve_key", retrieve_key)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter(name="expiresIn")
    def expires_in(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "expires_in")

    @expires_in.setter
    def expires_in(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expires_in", value)

    @property
    @pulumi.getter
    def expiring(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "expiring")

    @expiring.setter
    def expiring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "expiring", value)

    @property
    @pulumi.getter
    def identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter
    def intent(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        return pulumi.get(self, "intent")

    @intent.setter
    def intent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intent", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="retrieveKey")
    def retrieve_key(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retrieve_key")

    @retrieve_key.setter
    def retrieve_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retrieve_key", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user", value)


class Token(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 expiring: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 intent: Optional[pulumi.Input[str]] = None,
                 retrieve_key: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a token for a user
        name = authentik.User("name", username="user")
        default = authentik.Token("default",
            identifier="my-token",
            user=name.id,
            description="My secret token",
            expires="2025-01-01T15:04:05Z")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] intent: Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a token for a user
        name = authentik.User("name", username="user")
        default = authentik.Token("default",
            identifier="my-token",
            user=name.id,
            description="My secret token",
            expires="2025-01-01T15:04:05Z")
        ```

        :param str resource_name: The name of the resource.
        :param TokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 expiring: Optional[pulumi.Input[bool]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 intent: Optional[pulumi.Input[str]] = None,
                 retrieve_key: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TokenArgs.__new__(TokenArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["expires"] = expires
            __props__.__dict__["expiring"] = expiring
            if identifier is None and not opts.urn:
                raise TypeError("Missing required property 'identifier'")
            __props__.__dict__["identifier"] = identifier
            __props__.__dict__["intent"] = intent
            __props__.__dict__["retrieve_key"] = retrieve_key
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
            __props__.__dict__["expires_in"] = None
            __props__.__dict__["key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["key"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Token, __self__).__init__(
            'authentik:index/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            expires: Optional[pulumi.Input[str]] = None,
            expires_in: Optional[pulumi.Input[int]] = None,
            expiring: Optional[pulumi.Input[bool]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            intent: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            retrieve_key: Optional[pulumi.Input[bool]] = None,
            user: Optional[pulumi.Input[int]] = None) -> 'Token':
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] intent: Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TokenState.__new__(_TokenState)

        __props__.__dict__["description"] = description
        __props__.__dict__["expires"] = expires
        __props__.__dict__["expires_in"] = expires_in
        __props__.__dict__["expiring"] = expiring
        __props__.__dict__["identifier"] = identifier
        __props__.__dict__["intent"] = intent
        __props__.__dict__["key"] = key
        __props__.__dict__["retrieve_key"] = retrieve_key
        __props__.__dict__["user"] = user
        return Token(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter(name="expiresIn")
    def expires_in(self) -> pulumi.Output[int]:
        return pulumi.get(self, "expires_in")

    @property
    @pulumi.getter
    def expiring(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "expiring")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter
    def intent(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `verification` - `api` - `recovery` - `app_password`
        """
        return pulumi.get(self, "intent")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="retrieveKey")
    def retrieve_key(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "retrieve_key")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[int]:
        return pulumi.get(self, "user")

