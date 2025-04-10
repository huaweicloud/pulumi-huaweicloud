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
    'GetRestoreTimeRangesResult',
    'AwaitableGetRestoreTimeRangesResult',
    'get_restore_time_ranges',
    'get_restore_time_ranges_output',
]

@pulumi.output_type
class GetRestoreTimeRangesResult:
    """
    A collection of values returned by getRestoreTimeRanges.
    """
    def __init__(__self__, date=None, id=None, instance_id=None, region=None, restore_times=None):
        if date and not isinstance(date, str):
            raise TypeError("Expected argument 'date' to be a str")
        pulumi.set(__self__, "date", date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if restore_times and not isinstance(restore_times, list):
            raise TypeError("Expected argument 'restore_times' to be a list")
        pulumi.set(__self__, "restore_times", restore_times)

    @property
    @pulumi.getter
    def date(self) -> Optional[str]:
        return pulumi.get(self, "date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="restoreTimes")
    def restore_times(self) -> Sequence['outputs.GetRestoreTimeRangesRestoreTimeResult']:
        """
        Indicates the list of restoration time ranges.
        """
        return pulumi.get(self, "restore_times")


class AwaitableGetRestoreTimeRangesResult(GetRestoreTimeRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRestoreTimeRangesResult(
            date=self.date,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region,
            restore_times=self.restore_times)


def get_restore_time_ranges(date: Optional[str] = None,
                            instance_id: Optional[str] = None,
                            region: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRestoreTimeRangesResult:
    """
    Use this data source to get the list of RDS restore time ranges.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_restore_time_ranges(instance_id=instance_id,
        date="2024-05-15")
    ```


    :param str date: Specifies the date to be queried.
           The value is in the **yyyy-mm-dd** format, and the time zone is UTC.
    :param str instance_id: Specifies the ID of RDS instance.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['date'] = date
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getRestoreTimeRanges:getRestoreTimeRanges', __args__, opts=opts, typ=GetRestoreTimeRangesResult).value

    return AwaitableGetRestoreTimeRangesResult(
        date=__ret__.date,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        region=__ret__.region,
        restore_times=__ret__.restore_times)


@_utilities.lift_output_func(get_restore_time_ranges)
def get_restore_time_ranges_output(date: Optional[pulumi.Input[Optional[str]]] = None,
                                   instance_id: Optional[pulumi.Input[str]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRestoreTimeRangesResult]:
    """
    Use this data source to get the list of RDS restore time ranges.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_restore_time_ranges(instance_id=instance_id,
        date="2024-05-15")
    ```


    :param str date: Specifies the date to be queried.
           The value is in the **yyyy-mm-dd** format, and the time zone is UTC.
    :param str instance_id: Specifies the ID of RDS instance.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    ...
