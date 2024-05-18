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
    'ListProjectMembersResult',
    'AwaitableListProjectMembersResult',
    'list_project_members',
    'list_project_members_output',
]

@pulumi.output_type
class ListProjectMembersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListProjectMembersProperties':
        return pulumi.get(self, "items")


class AwaitableListProjectMembersResult(ListProjectMembersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectMembersResult(
            items=self.items)


def list_project_members(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectMembersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('doppler-native:projects/v3:listProjectMembers', __args__, opts=opts, typ=ListProjectMembersResult).value

    return AwaitableListProjectMembersResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_project_members)
def list_project_members_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListProjectMembersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
