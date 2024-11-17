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
from . import outputs

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, attributes=None, id=None, include_users=None, is_superuser=None, name=None, num_pk=None, parent=None, parent_name=None, pk=None, users=None, users_objs=None):
        if attributes and not isinstance(attributes, str):
            raise TypeError("Expected argument 'attributes' to be a str")
        pulumi.set(__self__, "attributes", attributes)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_users and not isinstance(include_users, bool):
            raise TypeError("Expected argument 'include_users' to be a bool")
        pulumi.set(__self__, "include_users", include_users)
        if is_superuser and not isinstance(is_superuser, bool):
            raise TypeError("Expected argument 'is_superuser' to be a bool")
        pulumi.set(__self__, "is_superuser", is_superuser)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if num_pk and not isinstance(num_pk, int):
            raise TypeError("Expected argument 'num_pk' to be a int")
        pulumi.set(__self__, "num_pk", num_pk)
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        pulumi.set(__self__, "parent", parent)
        if parent_name and not isinstance(parent_name, str):
            raise TypeError("Expected argument 'parent_name' to be a str")
        pulumi.set(__self__, "parent_name", parent_name)
        if pk and not isinstance(pk, str):
            raise TypeError("Expected argument 'pk' to be a str")
        pulumi.set(__self__, "pk", pk)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if users_objs and not isinstance(users_objs, list):
            raise TypeError("Expected argument 'users_objs' to be a list")
        pulumi.set(__self__, "users_objs", users_objs)

    @property
    @pulumi.getter
    def attributes(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeUsers")
    def include_users(self) -> Optional[bool]:
        return pulumi.get(self, "include_users")

    @property
    @pulumi.getter(name="isSuperuser")
    def is_superuser(self) -> bool:
        """
        Generated.
        """
        return pulumi.get(self, "is_superuser")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numPk")
    def num_pk(self) -> int:
        """
        Generated.
        """
        return pulumi.get(self, "num_pk")

    @property
    @pulumi.getter
    def parent(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="parentName")
    def parent_name(self) -> str:
        """
        Generated.
        """
        return pulumi.get(self, "parent_name")

    @property
    @pulumi.getter
    def pk(self) -> Optional[str]:
        return pulumi.get(self, "pk")

    @property
    @pulumi.getter
    def users(self) -> Sequence[int]:
        """
        Generated.
        """
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="usersObjs")
    def users_objs(self) -> Sequence['outputs.GetGroupUsersObjResult']:
        """
        Generated.
        """
        return pulumi.get(self, "users_objs")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            attributes=self.attributes,
            id=self.id,
            include_users=self.include_users,
            is_superuser=self.is_superuser,
            name=self.name,
            num_pk=self.num_pk,
            parent=self.parent,
            parent_name=self.parent_name,
            pk=self.pk,
            users=self.users,
            users_objs=self.users_objs)


def get_group(include_users: Optional[bool] = None,
              name: Optional[str] = None,
              pk: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Get groups by pk or name

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    admins = authentik.get_group(name="authentik Admins")
    ```
    """
    __args__ = dict()
    __args__['includeUsers'] = include_users
    __args__['name'] = name
    __args__['pk'] = pk
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('authentik:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        id=pulumi.get(__ret__, 'id'),
        include_users=pulumi.get(__ret__, 'include_users'),
        is_superuser=pulumi.get(__ret__, 'is_superuser'),
        name=pulumi.get(__ret__, 'name'),
        num_pk=pulumi.get(__ret__, 'num_pk'),
        parent=pulumi.get(__ret__, 'parent'),
        parent_name=pulumi.get(__ret__, 'parent_name'),
        pk=pulumi.get(__ret__, 'pk'),
        users=pulumi.get(__ret__, 'users'),
        users_objs=pulumi.get(__ret__, 'users_objs'))
def get_group_output(include_users: Optional[pulumi.Input[Optional[bool]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     pk: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    Get groups by pk or name

    ## Example Usage

    ```python
    import pulumi
    import pulumi_authentik as authentik

    admins = authentik.get_group(name="authentik Admins")
    ```
    """
    __args__ = dict()
    __args__['includeUsers'] = include_users
    __args__['name'] = name
    __args__['pk'] = pk
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('authentik:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult)
    return __ret__.apply(lambda __response__: GetGroupResult(
        attributes=pulumi.get(__response__, 'attributes'),
        id=pulumi.get(__response__, 'id'),
        include_users=pulumi.get(__response__, 'include_users'),
        is_superuser=pulumi.get(__response__, 'is_superuser'),
        name=pulumi.get(__response__, 'name'),
        num_pk=pulumi.get(__response__, 'num_pk'),
        parent=pulumi.get(__response__, 'parent'),
        parent_name=pulumi.get(__response__, 'parent_name'),
        pk=pulumi.get(__response__, 'pk'),
        users=pulumi.get(__response__, 'users'),
        users_objs=pulumi.get(__response__, 'users_objs')))
