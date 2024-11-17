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
    'GetStageResult',
    'AwaitableGetStageResult',
    'get_stage',
    'get_stage_output',
]

@pulumi.output_type
class GetStageResult:
    """
    A collection of values returned by getStage.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "name")


class AwaitableGetStageResult(GetStageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStageResult(
            id=self.id,
            name=self.name)


def get_stage(name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStageResult:
    """
    Get stages by name

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    default_authentication_identification = authentik.get_stage(name="default-authentication-identification")
    ```


    :param str name: Generated.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('authentik:index/getStage:getStage', __args__, opts=opts, typ=GetStageResult).value

    return AwaitableGetStageResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_stage_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStageResult]:
    """
    Get stages by name

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    default_authentication_identification = authentik.get_stage(name="default-authentication-identification")
    ```


    :param str name: Generated.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('authentik:index/getStage:getStage', __args__, opts=opts, typ=GetStageResult)
    return __ret__.apply(lambda __response__: GetStageResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
