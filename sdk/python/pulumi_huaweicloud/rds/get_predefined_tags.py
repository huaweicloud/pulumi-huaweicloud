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
    'GetPredefinedTagsResult',
    'AwaitableGetPredefinedTagsResult',
    'get_predefined_tags',
    'get_predefined_tags_output',
]

@pulumi.output_type
class GetPredefinedTagsResult:
    """
    A collection of values returned by getPredefinedTags.
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
    def tags(self) -> Sequence['outputs.GetPredefinedTagsTagResult']:
        """
        Indicates the list of predefined tags.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPredefinedTagsResult(GetPredefinedTagsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPredefinedTagsResult(
            id=self.id,
            region=self.region,
            tags=self.tags)


def get_predefined_tags(region: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPredefinedTagsResult:
    """
    Use this data source to get the predefined tags.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_predefined_tags()
    ```


    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getPredefinedTags:getPredefinedTags', __args__, opts=opts, typ=GetPredefinedTagsResult).value

    return AwaitableGetPredefinedTagsResult(
        id=__ret__.id,
        region=__ret__.region,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_predefined_tags)
def get_predefined_tags_output(region: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPredefinedTagsResult]:
    """
    Use this data source to get the predefined tags.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_predefined_tags()
    ```


    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    ...
