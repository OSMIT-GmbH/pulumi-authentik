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
    'GetGroupUsersObjResult',
    'GetGroupsGroupResult',
    'GetGroupsGroupUsersObjResult',
    'GetUsersUserResult',
]

@pulumi.output_type
class GetGroupUsersObjResult(dict):
    def __init__(__self__, *,
                 attributes: str,
                 email: str,
                 is_active: bool,
                 last_login: str,
                 name: str,
                 pk: int,
                 uid: str,
                 username: str):
        pulumi.set(__self__, "attributes", attributes)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "is_active", is_active)
        pulumi.set(__self__, "last_login", last_login)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "pk", pk)
        pulumi.set(__self__, "uid", uid)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def attributes(self) -> str:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> str:
        return pulumi.get(self, "last_login")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pk(self) -> int:
        return pulumi.get(self, "pk")

    @property
    @pulumi.getter
    def uid(self) -> str:
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


@pulumi.output_type
class GetGroupsGroupResult(dict):
    def __init__(__self__, *,
                 attributes: str,
                 is_superuser: bool,
                 name: str,
                 num_pk: int,
                 parent: str,
                 parent_name: str,
                 pk: str,
                 users: Sequence[int],
                 users_objs: Sequence['outputs.GetGroupsGroupUsersObjResult']):
        pulumi.set(__self__, "attributes", attributes)
        pulumi.set(__self__, "is_superuser", is_superuser)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "num_pk", num_pk)
        pulumi.set(__self__, "parent", parent)
        pulumi.set(__self__, "parent_name", parent_name)
        pulumi.set(__self__, "pk", pk)
        pulumi.set(__self__, "users", users)
        pulumi.set(__self__, "users_objs", users_objs)

    @property
    @pulumi.getter
    def attributes(self) -> str:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="isSuperuser")
    def is_superuser(self) -> bool:
        return pulumi.get(self, "is_superuser")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numPk")
    def num_pk(self) -> int:
        return pulumi.get(self, "num_pk")

    @property
    @pulumi.getter
    def parent(self) -> str:
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="parentName")
    def parent_name(self) -> str:
        return pulumi.get(self, "parent_name")

    @property
    @pulumi.getter
    def pk(self) -> str:
        return pulumi.get(self, "pk")

    @property
    @pulumi.getter
    def users(self) -> Sequence[int]:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="usersObjs")
    def users_objs(self) -> Sequence['outputs.GetGroupsGroupUsersObjResult']:
        return pulumi.get(self, "users_objs")


@pulumi.output_type
class GetGroupsGroupUsersObjResult(dict):
    def __init__(__self__, *,
                 attributes: str,
                 email: str,
                 is_active: bool,
                 last_login: str,
                 name: str,
                 pk: int,
                 uid: str,
                 username: str):
        pulumi.set(__self__, "attributes", attributes)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "is_active", is_active)
        pulumi.set(__self__, "last_login", last_login)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "pk", pk)
        pulumi.set(__self__, "uid", uid)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def attributes(self) -> str:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> str:
        return pulumi.get(self, "last_login")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pk(self) -> int:
        return pulumi.get(self, "pk")

    @property
    @pulumi.getter
    def uid(self) -> str:
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


@pulumi.output_type
class GetUsersUserResult(dict):
    def __init__(__self__, *,
                 attributes: str,
                 avatar: str,
                 email: str,
                 groups: Sequence[str],
                 is_active: bool,
                 is_superuser: bool,
                 last_login: str,
                 name: str,
                 path: str,
                 pk: int,
                 type: str,
                 uid: str,
                 username: str,
                 uuid: str):
        pulumi.set(__self__, "attributes", attributes)
        pulumi.set(__self__, "avatar", avatar)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "is_active", is_active)
        pulumi.set(__self__, "is_superuser", is_superuser)
        pulumi.set(__self__, "last_login", last_login)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "pk", pk)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "uid", uid)
        pulumi.set(__self__, "username", username)
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def attributes(self) -> str:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def avatar(self) -> str:
        return pulumi.get(self, "avatar")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def groups(self) -> Sequence[str]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="isSuperuser")
    def is_superuser(self) -> bool:
        return pulumi.get(self, "is_superuser")

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> str:
        return pulumi.get(self, "last_login")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def pk(self) -> int:
        return pulumi.get(self, "pk")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> str:
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        return pulumi.get(self, "uuid")


