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

__all__ = ['StageUserLoginArgs', 'StageUserLogin']

@pulumi.input_type
class StageUserLoginArgs:
    def __init__(__self__, *,
                 geoip_binding: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_binding: Optional[pulumi.Input[str]] = None,
                 remember_me_offset: Optional[pulumi.Input[str]] = None,
                 session_duration: Optional[pulumi.Input[str]] = None,
                 terminate_other_sessions: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a StageUserLogin resource.
        :param pulumi.Input[str] geoip_binding: Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        :param pulumi.Input[str] network_binding: Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        if geoip_binding is not None:
            pulumi.set(__self__, "geoip_binding", geoip_binding)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_binding is not None:
            pulumi.set(__self__, "network_binding", network_binding)
        if remember_me_offset is not None:
            pulumi.set(__self__, "remember_me_offset", remember_me_offset)
        if session_duration is not None:
            pulumi.set(__self__, "session_duration", session_duration)
        if terminate_other_sessions is not None:
            pulumi.set(__self__, "terminate_other_sessions", terminate_other_sessions)

    @property
    @pulumi.getter(name="geoipBinding")
    def geoip_binding(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        """
        return pulumi.get(self, "geoip_binding")

    @geoip_binding.setter
    def geoip_binding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geoip_binding", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkBinding")
    def network_binding(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        return pulumi.get(self, "network_binding")

    @network_binding.setter
    def network_binding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_binding", value)

    @property
    @pulumi.getter(name="rememberMeOffset")
    def remember_me_offset(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "remember_me_offset")

    @remember_me_offset.setter
    def remember_me_offset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remember_me_offset", value)

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "session_duration")

    @session_duration.setter
    def session_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_duration", value)

    @property
    @pulumi.getter(name="terminateOtherSessions")
    def terminate_other_sessions(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "terminate_other_sessions")

    @terminate_other_sessions.setter
    def terminate_other_sessions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "terminate_other_sessions", value)


@pulumi.input_type
class _StageUserLoginState:
    def __init__(__self__, *,
                 geoip_binding: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_binding: Optional[pulumi.Input[str]] = None,
                 remember_me_offset: Optional[pulumi.Input[str]] = None,
                 session_duration: Optional[pulumi.Input[str]] = None,
                 terminate_other_sessions: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering StageUserLogin resources.
        :param pulumi.Input[str] geoip_binding: Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        :param pulumi.Input[str] network_binding: Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        if geoip_binding is not None:
            pulumi.set(__self__, "geoip_binding", geoip_binding)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_binding is not None:
            pulumi.set(__self__, "network_binding", network_binding)
        if remember_me_offset is not None:
            pulumi.set(__self__, "remember_me_offset", remember_me_offset)
        if session_duration is not None:
            pulumi.set(__self__, "session_duration", session_duration)
        if terminate_other_sessions is not None:
            pulumi.set(__self__, "terminate_other_sessions", terminate_other_sessions)

    @property
    @pulumi.getter(name="geoipBinding")
    def geoip_binding(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        """
        return pulumi.get(self, "geoip_binding")

    @geoip_binding.setter
    def geoip_binding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geoip_binding", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkBinding")
    def network_binding(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        return pulumi.get(self, "network_binding")

    @network_binding.setter
    def network_binding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_binding", value)

    @property
    @pulumi.getter(name="rememberMeOffset")
    def remember_me_offset(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "remember_me_offset")

    @remember_me_offset.setter
    def remember_me_offset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remember_me_offset", value)

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "session_duration")

    @session_duration.setter
    def session_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_duration", value)

    @property
    @pulumi.getter(name="terminateOtherSessions")
    def terminate_other_sessions(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "terminate_other_sessions")

    @terminate_other_sessions.setter
    def terminate_other_sessions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "terminate_other_sessions", value)


class StageUserLogin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geoip_binding: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_binding: Optional[pulumi.Input[str]] = None,
                 remember_me_offset: Optional[pulumi.Input[str]] = None,
                 session_duration: Optional[pulumi.Input[str]] = None,
                 terminate_other_sessions: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a user login stage
        name = authentik.StageUserLogin("name", name="user-login")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] geoip_binding: Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        :param pulumi.Input[str] network_binding: Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StageUserLoginArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import osmit_gmbh_pulumi_authentik as authentik

        # Create a user login stage
        name = authentik.StageUserLogin("name", name="user-login")
        ```

        :param str resource_name: The name of the resource.
        :param StageUserLoginArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageUserLoginArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geoip_binding: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_binding: Optional[pulumi.Input[str]] = None,
                 remember_me_offset: Optional[pulumi.Input[str]] = None,
                 session_duration: Optional[pulumi.Input[str]] = None,
                 terminate_other_sessions: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StageUserLoginArgs.__new__(StageUserLoginArgs)

            __props__.__dict__["geoip_binding"] = geoip_binding
            __props__.__dict__["name"] = name
            __props__.__dict__["network_binding"] = network_binding
            __props__.__dict__["remember_me_offset"] = remember_me_offset
            __props__.__dict__["session_duration"] = session_duration
            __props__.__dict__["terminate_other_sessions"] = terminate_other_sessions
        super(StageUserLogin, __self__).__init__(
            'authentik:index/stageUserLogin:StageUserLogin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            geoip_binding: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_binding: Optional[pulumi.Input[str]] = None,
            remember_me_offset: Optional[pulumi.Input[str]] = None,
            session_duration: Optional[pulumi.Input[str]] = None,
            terminate_other_sessions: Optional[pulumi.Input[bool]] = None) -> 'StageUserLogin':
        """
        Get an existing StageUserLogin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] geoip_binding: Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        :param pulumi.Input[str] network_binding: Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StageUserLoginState.__new__(_StageUserLoginState)

        __props__.__dict__["geoip_binding"] = geoip_binding
        __props__.__dict__["name"] = name
        __props__.__dict__["network_binding"] = network_binding
        __props__.__dict__["remember_me_offset"] = remember_me_offset
        __props__.__dict__["session_duration"] = session_duration
        __props__.__dict__["terminate_other_sessions"] = terminate_other_sessions
        return StageUserLogin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="geoipBinding")
    def geoip_binding(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `no_binding` - `bind_continent` - `bind_continent_country` - `bind_continent_country_city`
        """
        return pulumi.get(self, "geoip_binding")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkBinding")
    def network_binding(self) -> pulumi.Output[Optional[str]]:
        """
        Allowed values: - `no_binding` - `bind_asn` - `bind_asn_network` - `bind_asn_network_ip`
        """
        return pulumi.get(self, "network_binding")

    @property
    @pulumi.getter(name="rememberMeOffset")
    def remember_me_offset(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "remember_me_offset")

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "session_duration")

    @property
    @pulumi.getter(name="terminateOtherSessions")
    def terminate_other_sessions(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "terminate_other_sessions")

