# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetServiceAccountTokenResult',
    'AwaitableGetServiceAccountTokenResult',
    'get_service_account_token',
    'get_service_account_token_output',
]

@pulumi.output_type
class GetServiceAccountTokenResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetServiceAccountTokenProperties':
        return pulumi.get(self, "items")


class AwaitableGetServiceAccountTokenResult(GetServiceAccountTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountTokenResult(
            items=self.items)


def get_service_account_token(api_token: Optional[str] = None,
                              service_account: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountTokenResult:
    """
    Use this data source to access information about an existing resource.

    :param str api_token: Slug of the API token
    :param str service_account: Slug of the service account
    """
    __args__ = dict()
    __args__['apiToken'] = api_token
    __args__['serviceAccount'] = service_account
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:workplace/v3:getServiceAccountToken', __args__, opts=opts, typ=GetServiceAccountTokenResult).value

    return AwaitableGetServiceAccountTokenResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_service_account_token)
def get_service_account_token_output(api_token: Optional[pulumi.Input[str]] = None,
                                     service_account: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceAccountTokenResult]:
    """
    Use this data source to access information about an existing resource.

    :param str api_token: Slug of the API token
    :param str service_account: Slug of the service account
    """
    ...
