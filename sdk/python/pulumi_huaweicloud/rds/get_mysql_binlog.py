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
    'GetMysqlBinlogResult',
    'AwaitableGetMysqlBinlogResult',
    'get_mysql_binlog',
    'get_mysql_binlog_output',
]

@pulumi.output_type
class GetMysqlBinlogResult:
    """
    A collection of values returned by getMysqlBinlog.
    """
    def __init__(__self__, binlog_retention_hours=None, id=None, instance_id=None, region=None):
        if binlog_retention_hours and not isinstance(binlog_retention_hours, int):
            raise TypeError("Expected argument 'binlog_retention_hours' to be a int")
        pulumi.set(__self__, "binlog_retention_hours", binlog_retention_hours)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="binlogRetentionHours")
    def binlog_retention_hours(self) -> int:
        """
        The binlog retention period. Value range: 0 to 168 (7x24).
        """
        return pulumi.get(self, "binlog_retention_hours")

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


class AwaitableGetMysqlBinlogResult(GetMysqlBinlogResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMysqlBinlogResult(
            binlog_retention_hours=self.binlog_retention_hours,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region)


def get_mysql_binlog(instance_id: Optional[str] = None,
                     region: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMysqlBinlogResult:
    """
    Use this data source to get the binlog retention hours of RDS MySQL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_mysql_binlog(instance_id=var["instance_id"])
    ```


    :param str instance_id: Specifies the ID of the RDS MySQL instance.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getMysqlBinlog:getMysqlBinlog', __args__, opts=opts, typ=GetMysqlBinlogResult).value

    return AwaitableGetMysqlBinlogResult(
        binlog_retention_hours=__ret__.binlog_retention_hours,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        region=__ret__.region)


@_utilities.lift_output_func(get_mysql_binlog)
def get_mysql_binlog_output(instance_id: Optional[pulumi.Input[str]] = None,
                            region: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMysqlBinlogResult]:
    """
    Use this data source to get the binlog retention hours of RDS MySQL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Rds.get_mysql_binlog(instance_id=var["instance_id"])
    ```


    :param str instance_id: Specifies the ID of the RDS MySQL instance.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    ...
