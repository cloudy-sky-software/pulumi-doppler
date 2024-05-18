# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'AuthenticationProperties',
]

@pulumi.output_type
class AuthenticationProperties(dict):
    def __init__(__self__, *,
                 password: Optional[str] = None,
                 token: Optional[str] = None,
                 type: Optional['WebhooksAuthenticationPropertiesType'] = None,
                 username: Optional[str] = None):
        """
        :param str password: Used when type = Basic
        :param str token: Used when type = Bearer
        :param str username: Used when type = Basic
        """
        if password is not None:
            pulumi.set(__self__, "password", password)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        Used when type = Basic
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Used when type = Bearer
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def type(self) -> Optional['WebhooksAuthenticationPropertiesType']:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        Used when type = Basic
        """
        return pulumi.get(self, "username")

