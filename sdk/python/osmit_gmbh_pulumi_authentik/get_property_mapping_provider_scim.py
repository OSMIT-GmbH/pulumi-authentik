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

__all__ = [
    'GetPropertyMappingProviderScimResult',
    'AwaitableGetPropertyMappingProviderScimResult',
    'get_property_mapping_provider_scim',
    'get_property_mapping_provider_scim_output',
]

@pulumi.output_type
class GetPropertyMappingProviderScimResult:
    """
    A collection of values returned by getPropertyMappingProviderScim.
    """
    def __init__(__self__, expression=None, id=None, ids=None, managed=None, managed_lists=None, name=None):
        if expression and not isinstance(expression, str):
            raise TypeError("Expected argument 'expression' to be a str")
        pulumi.set(__self__, "expression", expression)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if managed and not isinstance(managed, str):
            raise TypeError("Expected argument 'managed' to be a str")
        pulumi.set(__self__, "managed", managed)
        if managed_lists and not isinstance(managed_lists, list):
            raise TypeError("Expected argument 'managed_lists' to be a list")
        pulumi.set(__self__, "managed_lists", managed_lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        List of ids when `managed_list` is set. Generated.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def managed(self) -> Optional[str]:
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter(name="managedLists")
    def managed_lists(self) -> Optional[Sequence[str]]:
        """
        Retrieve multiple property mappings
        """
        return pulumi.get(self, "managed_lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


class AwaitableGetPropertyMappingProviderScimResult(GetPropertyMappingProviderScimResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPropertyMappingProviderScimResult(
            expression=self.expression,
            id=self.id,
            ids=self.ids,
            managed=self.managed,
            managed_lists=self.managed_lists,
            name=self.name)


def get_property_mapping_provider_scim(ids: Optional[Sequence[str]] = None,
                                       managed: Optional[str] = None,
                                       managed_lists: Optional[Sequence[str]] = None,
                                       name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPropertyMappingProviderScimResult:
    """
    Get SCIM Provider Property mappings

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    test = authentik.get_property_mapping_provider_scim(managed_lists=[
        "goauthentik.io/providers/scim/user",
        "goauthentik.io/providers/scim/group",
    ])
    ```


    :param Sequence[str] ids: List of ids when `managed_list` is set. Generated.
    :param Sequence[str] managed_lists: Retrieve multiple property mappings
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['managed'] = managed
    __args__['managedLists'] = managed_lists
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('authentik:index/getPropertyMappingProviderScim:getPropertyMappingProviderScim', __args__, opts=opts, typ=GetPropertyMappingProviderScimResult).value

    return AwaitableGetPropertyMappingProviderScimResult(
        expression=pulumi.get(__ret__, 'expression'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        managed=pulumi.get(__ret__, 'managed'),
        managed_lists=pulumi.get(__ret__, 'managed_lists'),
        name=pulumi.get(__ret__, 'name'))
def get_property_mapping_provider_scim_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              managed: Optional[pulumi.Input[Optional[str]]] = None,
                                              managed_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              name: Optional[pulumi.Input[Optional[str]]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPropertyMappingProviderScimResult]:
    """
    Get SCIM Provider Property mappings

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    test = authentik.get_property_mapping_provider_scim(managed_lists=[
        "goauthentik.io/providers/scim/user",
        "goauthentik.io/providers/scim/group",
    ])
    ```


    :param Sequence[str] ids: List of ids when `managed_list` is set. Generated.
    :param Sequence[str] managed_lists: Retrieve multiple property mappings
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['managed'] = managed
    __args__['managedLists'] = managed_lists
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('authentik:index/getPropertyMappingProviderScim:getPropertyMappingProviderScim', __args__, opts=opts, typ=GetPropertyMappingProviderScimResult)
    return __ret__.apply(lambda __response__: GetPropertyMappingProviderScimResult(
        expression=pulumi.get(__response__, 'expression'),
        id=pulumi.get(__response__, 'id'),
        ids=pulumi.get(__response__, 'ids'),
        managed=pulumi.get(__response__, 'managed'),
        managed_lists=pulumi.get(__response__, 'managed_lists'),
        name=pulumi.get(__response__, 'name')))
