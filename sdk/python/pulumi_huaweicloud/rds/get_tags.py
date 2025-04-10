# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTagsResult',
    'AwaitableGetTagsResult',
    'get_tags',
    'get_tags_output',
]

@pulumi.output_type
class GetTagsResult:
    """
    A collection of values returned by getTags.
    """
    def __init__(__self__, id=None, region=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetTagsTagResult']:
        """
        Indicates the tag list.
        """
        return pulumi.get(self, "tags")


class AwaitableGetTagsResult(GetTagsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagsResult(
            id=self.id,
            region=self.region,
            tags=self.tags)


def get_tags(region: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagsResult:
    """
    Use this data source to get the project tags.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_tags()
    ```


    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getTags:getTags', __args__, opts=opts, typ=GetTagsResult).value

    return AwaitableGetTagsResult(
        id=__ret__.id,
        region=__ret__.region,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_tags)
def get_tags_output(region: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTagsResult]:
    """
    Use this data source to get the project tags.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_tags()
    ```


    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    ...
