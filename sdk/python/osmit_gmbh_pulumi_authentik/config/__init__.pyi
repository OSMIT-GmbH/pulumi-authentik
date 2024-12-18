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
from .. import _utilities

headers: Optional[str]
"""
Optional HTTP headers sent with every request
"""

insecure: Optional[bool]
"""
Whether to skip TLS verification, can optionally be passed as `AUTHENTIK_INSECURE` environmental variable
"""

token: Optional[str]
"""
The authentik API token, can optionally be passed as `AUTHENTIK_TOKEN` environmental variable
"""

url: Optional[str]
"""
The authentik API endpoint, can optionally be passed as `AUTHENTIK_URL` environmental variable
"""

