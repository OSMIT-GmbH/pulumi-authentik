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

__all__ = ['SourcePlexArgs', 'SourcePlex']

@pulumi.input_type
class SourcePlexArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 plex_token: pulumi.Input[str],
                 slug: pulumi.Input[str],
                 allow_friends: Optional[pulumi.Input[bool]] = None,
                 allowed_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 enrollment_flow: Optional[pulumi.Input[str]] = None,
                 group_matching_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_engine_mode: Optional[pulumi.Input[str]] = None,
                 user_matching_mode: Optional[pulumi.Input[str]] = None,
                 user_path_template: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SourcePlex resource.
        :param pulumi.Input[str] group_matching_mode: Allowed values: - `identifier` - `name_link` - `name_deny`
        :param pulumi.Input[str] policy_engine_mode: Allowed values: - `all` - `any`
        :param pulumi.Input[str] user_matching_mode: Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "plex_token", plex_token)
        pulumi.set(__self__, "slug", slug)
        if allow_friends is not None:
            pulumi.set(__self__, "allow_friends", allow_friends)
        if allowed_servers is not None:
            pulumi.set(__self__, "allowed_servers", allowed_servers)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if enrollment_flow is not None:
            pulumi.set(__self__, "enrollment_flow", enrollment_flow)
        if group_matching_mode is not None:
            pulumi.set(__self__, "group_matching_mode", group_matching_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_engine_mode is not None:
            pulumi.set(__self__, "policy_engine_mode", policy_engine_mode)
        if user_matching_mode is not None:
            pulumi.set(__self__, "user_matching_mode", user_matching_mode)
        if user_path_template is not None:
            pulumi.set(__self__, "user_path_template", user_path_template)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="plexToken")
    def plex_token(self) -> pulumi.Input[str]:
        return pulumi.get(self, "plex_token")

    @plex_token.setter
    def plex_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "plex_token", value)

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Input[str]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: pulumi.Input[str]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="allowFriends")
    def allow_friends(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_friends")

    @allow_friends.setter
    def allow_friends(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_friends", value)

    @property
    @pulumi.getter(name="allowedServers")
    def allowed_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "allowed_servers")

    @allowed_servers.setter
    def allowed_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_servers", value)

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authentication_flow")

    @authentication_flow.setter
    def authentication_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_flow", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="enrollmentFlow")
    def enrollment_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enrollment_flow")

    @enrollment_flow.setter
    def enrollment_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enrollment_flow", value)

    @property
    @pulumi.getter(name="groupMatchingMode")
    def group_matching_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `identifier` - `name_link` - `name_deny`
        """
        return pulumi.get(self, "group_matching_mode")

    @group_matching_mode.setter
    def group_matching_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_matching_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyEngineMode")
    def policy_engine_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `all` - `any`
        """
        return pulumi.get(self, "policy_engine_mode")

    @policy_engine_mode.setter
    def policy_engine_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_engine_mode", value)

    @property
    @pulumi.getter(name="userMatchingMode")
    def user_matching_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        return pulumi.get(self, "user_matching_mode")

    @user_matching_mode.setter
    def user_matching_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_matching_mode", value)

    @property
    @pulumi.getter(name="userPathTemplate")
    def user_path_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_path_template")

    @user_path_template.setter
    def user_path_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_path_template", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


@pulumi.input_type
class _SourcePlexState:
    def __init__(__self__, *,
                 allow_friends: Optional[pulumi.Input[bool]] = None,
                 allowed_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 enrollment_flow: Optional[pulumi.Input[str]] = None,
                 group_matching_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plex_token: Optional[pulumi.Input[str]] = None,
                 policy_engine_mode: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_matching_mode: Optional[pulumi.Input[str]] = None,
                 user_path_template: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SourcePlex resources.
        :param pulumi.Input[str] group_matching_mode: Allowed values: - `identifier` - `name_link` - `name_deny`
        :param pulumi.Input[str] policy_engine_mode: Allowed values: - `all` - `any`
        :param pulumi.Input[str] user_matching_mode: Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        if allow_friends is not None:
            pulumi.set(__self__, "allow_friends", allow_friends)
        if allowed_servers is not None:
            pulumi.set(__self__, "allowed_servers", allowed_servers)
        if authentication_flow is not None:
            pulumi.set(__self__, "authentication_flow", authentication_flow)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if enrollment_flow is not None:
            pulumi.set(__self__, "enrollment_flow", enrollment_flow)
        if group_matching_mode is not None:
            pulumi.set(__self__, "group_matching_mode", group_matching_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plex_token is not None:
            pulumi.set(__self__, "plex_token", plex_token)
        if policy_engine_mode is not None:
            pulumi.set(__self__, "policy_engine_mode", policy_engine_mode)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if user_matching_mode is not None:
            pulumi.set(__self__, "user_matching_mode", user_matching_mode)
        if user_path_template is not None:
            pulumi.set(__self__, "user_path_template", user_path_template)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="allowFriends")
    def allow_friends(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_friends")

    @allow_friends.setter
    def allow_friends(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_friends", value)

    @property
    @pulumi.getter(name="allowedServers")
    def allowed_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "allowed_servers")

    @allowed_servers.setter
    def allowed_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_servers", value)

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authentication_flow")

    @authentication_flow.setter
    def authentication_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_flow", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="enrollmentFlow")
    def enrollment_flow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enrollment_flow")

    @enrollment_flow.setter
    def enrollment_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enrollment_flow", value)

    @property
    @pulumi.getter(name="groupMatchingMode")
    def group_matching_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `identifier` - `name_link` - `name_deny`
        """
        return pulumi.get(self, "group_matching_mode")

    @group_matching_mode.setter
    def group_matching_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_matching_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="plexToken")
    def plex_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "plex_token")

    @plex_token.setter
    def plex_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plex_token", value)

    @property
    @pulumi.getter(name="policyEngineMode")
    def policy_engine_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `all` - `any`
        """
        return pulumi.get(self, "policy_engine_mode")

    @policy_engine_mode.setter
    def policy_engine_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_engine_mode", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="userMatchingMode")
    def user_matching_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        return pulumi.get(self, "user_matching_mode")

    @user_matching_mode.setter
    def user_matching_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_matching_mode", value)

    @property
    @pulumi.getter(name="userPathTemplate")
    def user_path_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_path_template")

    @user_path_template.setter
    def user_path_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_path_template", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class SourcePlex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_friends: Optional[pulumi.Input[bool]] = None,
                 allowed_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 enrollment_flow: Optional[pulumi.Input[str]] = None,
                 group_matching_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plex_token: Optional[pulumi.Input[str]] = None,
                 policy_engine_mode: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_matching_mode: Optional[pulumi.Input[str]] = None,
                 user_path_template: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        # Create a plex source
        default_authorization_flow = authentik.get_flow(slug="default-provider-authorization-implicit-consent")
        name = authentik.SourcePlex("name",
            name="plex",
            slug="plex",
            authentication_flow=default_authorization_flow.id,
            enrollment_flow=default_authorization_flow.id,
            client_id="foo-bar-baz",
            plex_token="foo")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_matching_mode: Allowed values: - `identifier` - `name_link` - `name_deny`
        :param pulumi.Input[str] policy_engine_mode: Allowed values: - `all` - `any`
        :param pulumi.Input[str] user_matching_mode: Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SourcePlexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik
        import pulumi_authentik as authentik

        # Create a plex source
        default_authorization_flow = authentik.get_flow(slug="default-provider-authorization-implicit-consent")
        name = authentik.SourcePlex("name",
            name="plex",
            slug="plex",
            authentication_flow=default_authorization_flow.id,
            enrollment_flow=default_authorization_flow.id,
            client_id="foo-bar-baz",
            plex_token="foo")
        ```

        :param str resource_name: The name of the resource.
        :param SourcePlexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SourcePlexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_friends: Optional[pulumi.Input[bool]] = None,
                 allowed_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authentication_flow: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 enrollment_flow: Optional[pulumi.Input[str]] = None,
                 group_matching_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plex_token: Optional[pulumi.Input[str]] = None,
                 policy_engine_mode: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_matching_mode: Optional[pulumi.Input[str]] = None,
                 user_path_template: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SourcePlexArgs.__new__(SourcePlexArgs)

            __props__.__dict__["allow_friends"] = allow_friends
            __props__.__dict__["allowed_servers"] = allowed_servers
            __props__.__dict__["authentication_flow"] = authentication_flow
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["enrollment_flow"] = enrollment_flow
            __props__.__dict__["group_matching_mode"] = group_matching_mode
            __props__.__dict__["name"] = name
            if plex_token is None and not opts.urn:
                raise TypeError("Missing required property 'plex_token'")
            __props__.__dict__["plex_token"] = None if plex_token is None else pulumi.Output.secret(plex_token)
            __props__.__dict__["policy_engine_mode"] = policy_engine_mode
            if slug is None and not opts.urn:
                raise TypeError("Missing required property 'slug'")
            __props__.__dict__["slug"] = slug
            __props__.__dict__["user_matching_mode"] = user_matching_mode
            __props__.__dict__["user_path_template"] = user_path_template
            __props__.__dict__["uuid"] = uuid
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["plexToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SourcePlex, __self__).__init__(
            'authentik:index/sourcePlex:SourcePlex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_friends: Optional[pulumi.Input[bool]] = None,
            allowed_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            authentication_flow: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            enrollment_flow: Optional[pulumi.Input[str]] = None,
            group_matching_mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            plex_token: Optional[pulumi.Input[str]] = None,
            policy_engine_mode: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            user_matching_mode: Optional[pulumi.Input[str]] = None,
            user_path_template: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'SourcePlex':
        """
        Get an existing SourcePlex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_matching_mode: Allowed values: - `identifier` - `name_link` - `name_deny`
        :param pulumi.Input[str] policy_engine_mode: Allowed values: - `all` - `any`
        :param pulumi.Input[str] user_matching_mode: Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SourcePlexState.__new__(_SourcePlexState)

        __props__.__dict__["allow_friends"] = allow_friends
        __props__.__dict__["allowed_servers"] = allowed_servers
        __props__.__dict__["authentication_flow"] = authentication_flow
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["enrollment_flow"] = enrollment_flow
        __props__.__dict__["group_matching_mode"] = group_matching_mode
        __props__.__dict__["name"] = name
        __props__.__dict__["plex_token"] = plex_token
        __props__.__dict__["policy_engine_mode"] = policy_engine_mode
        __props__.__dict__["slug"] = slug
        __props__.__dict__["user_matching_mode"] = user_matching_mode
        __props__.__dict__["user_path_template"] = user_path_template
        __props__.__dict__["uuid"] = uuid
        return SourcePlex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowFriends")
    def allow_friends(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allow_friends")

    @property
    @pulumi.getter(name="allowedServers")
    def allowed_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "allowed_servers")

    @property
    @pulumi.getter(name="authenticationFlow")
    def authentication_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authentication_flow")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="enrollmentFlow")
    def enrollment_flow(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrollment_flow")

    @property
    @pulumi.getter(name="groupMatchingMode")
    def group_matching_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `identifier` - `name_link` - `name_deny`
        """
        return pulumi.get(self, "group_matching_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="plexToken")
    def plex_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "plex_token")

    @property
    @pulumi.getter(name="policyEngineMode")
    def policy_engine_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `all` - `any`
        """
        return pulumi.get(self, "policy_engine_mode")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="userMatchingMode")
    def user_matching_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `identifier` - `email_link` - `email_deny` - `username_link` - `username_deny`
        """
        return pulumi.get(self, "user_matching_mode")

    @property
    @pulumi.getter(name="userPathTemplate")
    def user_path_template(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_path_template")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uuid")

