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
    'GetPropertyMappingProviderSamlResult',
    'AwaitableGetPropertyMappingProviderSamlResult',
    'get_property_mapping_provider_saml',
    'get_property_mapping_provider_saml_output',
]

@pulumi.output_type
class GetPropertyMappingProviderSamlResult:
    """
    A collection of values returned by getPropertyMappingProviderSaml.
    """
    def __init__(__self__, expression=None, friendly_name=None, id=None, ids=None, managed=None, managed_lists=None, name=None, saml_name=None):
        if expression and not isinstance(expression, str):
            raise TypeError("Expected argument 'expression' to be a str")
        pulumi.set(__self__, "expression", expression)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
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
        if saml_name and not isinstance(saml_name, str):
            raise TypeError("Expected argument 'saml_name' to be a str")
        pulumi.set(__self__, "saml_name", saml_name)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "friendly_name")

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

    @property
    @pulumi.getter(name="samlName")
    def saml_name(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "saml_name")


class AwaitableGetPropertyMappingProviderSamlResult(GetPropertyMappingProviderSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPropertyMappingProviderSamlResult(
            expression=self.expression,
            friendly_name=self.friendly_name,
            id=self.id,
            ids=self.ids,
            managed=self.managed,
            managed_lists=self.managed_lists,
            name=self.name,
            saml_name=self.saml_name)


def get_property_mapping_provider_saml(expression: Optional[str] = None,
                                       friendly_name: Optional[str] = None,
                                       ids: Optional[Sequence[str]] = None,
                                       managed: Optional[str] = None,
                                       managed_lists: Optional[Sequence[str]] = None,
                                       name: Optional[str] = None,
                                       saml_name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPropertyMappingProviderSamlResult:
    """
    Get SAML Provider Property mappings


    :param str expression: Generated.
    :param str friendly_name: Generated.
    :param Sequence[str] ids: List of ids when `managed_list` is set. Generated.
    :param Sequence[str] managed_lists: Retrieve multiple property mappings
    :param str saml_name: Generated.
    """
    __args__ = dict()
    __args__['expression'] = expression
    __args__['friendlyName'] = friendly_name
    __args__['ids'] = ids
    __args__['managed'] = managed
    __args__['managedLists'] = managed_lists
    __args__['name'] = name
    __args__['samlName'] = saml_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('authentik:index/getPropertyMappingProviderSaml:getPropertyMappingProviderSaml', __args__, opts=opts, typ=GetPropertyMappingProviderSamlResult).value

    return AwaitableGetPropertyMappingProviderSamlResult(
        expression=pulumi.get(__ret__, 'expression'),
        friendly_name=pulumi.get(__ret__, 'friendly_name'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        managed=pulumi.get(__ret__, 'managed'),
        managed_lists=pulumi.get(__ret__, 'managed_lists'),
        name=pulumi.get(__ret__, 'name'),
        saml_name=pulumi.get(__ret__, 'saml_name'))
def get_property_mapping_provider_saml_output(expression: Optional[pulumi.Input[Optional[str]]] = None,
                                              friendly_name: Optional[pulumi.Input[Optional[str]]] = None,
                                              ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              managed: Optional[pulumi.Input[Optional[str]]] = None,
                                              managed_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              name: Optional[pulumi.Input[Optional[str]]] = None,
                                              saml_name: Optional[pulumi.Input[Optional[str]]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPropertyMappingProviderSamlResult]:
    """
    Get SAML Provider Property mappings


    :param str expression: Generated.
    :param str friendly_name: Generated.
    :param Sequence[str] ids: List of ids when `managed_list` is set. Generated.
    :param Sequence[str] managed_lists: Retrieve multiple property mappings
    :param str saml_name: Generated.
    """
    __args__ = dict()
    __args__['expression'] = expression
    __args__['friendlyName'] = friendly_name
    __args__['ids'] = ids
    __args__['managed'] = managed
    __args__['managedLists'] = managed_lists
    __args__['name'] = name
    __args__['samlName'] = saml_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('authentik:index/getPropertyMappingProviderSaml:getPropertyMappingProviderSaml', __args__, opts=opts, typ=GetPropertyMappingProviderSamlResult)
    return __ret__.apply(lambda __response__: GetPropertyMappingProviderSamlResult(
        expression=pulumi.get(__response__, 'expression'),
        friendly_name=pulumi.get(__response__, 'friendly_name'),
        id=pulumi.get(__response__, 'id'),
        ids=pulumi.get(__response__, 'ids'),
        managed=pulumi.get(__response__, 'managed'),
        managed_lists=pulumi.get(__response__, 'managed_lists'),
        name=pulumi.get(__response__, 'name'),
        saml_name=pulumi.get(__response__, 'saml_name')))
