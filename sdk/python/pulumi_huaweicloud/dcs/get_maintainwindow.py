# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetMaintainwindowResult',
    'AwaitableGetMaintainwindowResult',
    'get_maintainwindow',
    'get_maintainwindow_output',
]

@pulumi.output_type
class GetMaintainwindowResult:
    """
    A collection of values returned by getMaintainwindow.
    """
    def __init__(__self__, begin=None, default=None, end=None, id=None, region=None, seq=None):
        if begin and not isinstance(begin, str):
            raise TypeError("Expected argument 'begin' to be a str")
        pulumi.set(__self__, "begin", begin)
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if end and not isinstance(end, str):
            raise TypeError("Expected argument 'end' to be a str")
        pulumi.set(__self__, "end", end)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if seq and not isinstance(seq, int):
            raise TypeError("Expected argument 'seq' to be a int")
        pulumi.set(__self__, "seq", seq)

    @property
    @pulumi.getter
    def begin(self) -> str:
        return pulumi.get(self, "begin")

    @property
    @pulumi.getter
    def default(self) -> bool:
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def end(self) -> str:
        return pulumi.get(self, "end")

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
    def seq(self) -> int:
        return pulumi.get(self, "seq")


class AwaitableGetMaintainwindowResult(GetMaintainwindowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMaintainwindowResult(
            begin=self.begin,
            default=self.default,
            end=self.end,
            id=self.id,
            region=self.region,
            seq=self.seq)


def get_maintainwindow(begin: Optional[str] = None,
                       default: Optional[bool] = None,
                       end: Optional[str] = None,
                       region: Optional[str] = None,
                       seq: Optional[int] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMaintainwindowResult:
    """
    Use this data source to get the ID of an available DCS maintainwindow.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    maintainwindow1 = huaweicloud.Dcs.get_maintainwindow(seq=1)
    ```


    :param str begin: Specifies the time at which a maintenance time window starts.
    :param bool default: Specifies whether a maintenance time window is set to the default time segment.
    :param str end: Specifies the time at which a maintenance time window ends.
    :param str region: The region in which to obtain the dcs maintainwindows. If omitted, the provider-level
           region will be used.
    :param int seq: Specifies the sequential number of a maintenance time window.
    """
    __args__ = dict()
    __args__['begin'] = begin
    __args__['default'] = default
    __args__['end'] = end
    __args__['region'] = region
    __args__['seq'] = seq
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Dcs/getMaintainwindow:getMaintainwindow', __args__, opts=opts, typ=GetMaintainwindowResult).value

    return AwaitableGetMaintainwindowResult(
        begin=__ret__.begin,
        default=__ret__.default,
        end=__ret__.end,
        id=__ret__.id,
        region=__ret__.region,
        seq=__ret__.seq)


@_utilities.lift_output_func(get_maintainwindow)
def get_maintainwindow_output(begin: Optional[pulumi.Input[Optional[str]]] = None,
                              default: Optional[pulumi.Input[Optional[bool]]] = None,
                              end: Optional[pulumi.Input[Optional[str]]] = None,
                              region: Optional[pulumi.Input[Optional[str]]] = None,
                              seq: Optional[pulumi.Input[Optional[int]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMaintainwindowResult]:
    """
    Use this data source to get the ID of an available DCS maintainwindow.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    maintainwindow1 = huaweicloud.Dcs.get_maintainwindow(seq=1)
    ```


    :param str begin: Specifies the time at which a maintenance time window starts.
    :param bool default: Specifies whether a maintenance time window is set to the default time segment.
    :param str end: Specifies the time at which a maintenance time window ends.
    :param str region: The region in which to obtain the dcs maintainwindows. If omitted, the provider-level
           region will be used.
    :param int seq: Specifies the sequential number of a maintenance time window.
    """
    ...