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
    'GetRecyclingInstancesResult',
    'AwaitableGetRecyclingInstancesResult',
    'get_recycling_instances',
    'get_recycling_instances_output',
]

@pulumi.output_type
class GetRecyclingInstancesResult:
    """
    A collection of values returned by getRecyclingInstances.
    """
    def __init__(__self__, data_vip=None, engine_name=None, engine_version=None, enterprise_project_id=None, ha_mode=None, id=None, instance_id=None, instances=None, is_serverless=None, name=None, pay_model=None, recycle_backup_id=None, recycle_status=None, region=None, volume_size=None, volume_type=None):
        if data_vip and not isinstance(data_vip, str):
            raise TypeError("Expected argument 'data_vip' to be a str")
        pulumi.set(__self__, "data_vip", data_vip)
        if engine_name and not isinstance(engine_name, str):
            raise TypeError("Expected argument 'engine_name' to be a str")
        pulumi.set(__self__, "engine_name", engine_name)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if ha_mode and not isinstance(ha_mode, str):
            raise TypeError("Expected argument 'ha_mode' to be a str")
        pulumi.set(__self__, "ha_mode", ha_mode)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if is_serverless and not isinstance(is_serverless, str):
            raise TypeError("Expected argument 'is_serverless' to be a str")
        pulumi.set(__self__, "is_serverless", is_serverless)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pay_model and not isinstance(pay_model, str):
            raise TypeError("Expected argument 'pay_model' to be a str")
        pulumi.set(__self__, "pay_model", pay_model)
        if recycle_backup_id and not isinstance(recycle_backup_id, str):
            raise TypeError("Expected argument 'recycle_backup_id' to be a str")
        pulumi.set(__self__, "recycle_backup_id", recycle_backup_id)
        if recycle_status and not isinstance(recycle_status, str):
            raise TypeError("Expected argument 'recycle_status' to be a str")
        pulumi.set(__self__, "recycle_status", recycle_status)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if volume_size and not isinstance(volume_size, str):
            raise TypeError("Expected argument 'volume_size' to be a str")
        pulumi.set(__self__, "volume_size", volume_size)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="dataVip")
    def data_vip(self) -> Optional[str]:
        """
        Indicates the floating IP address.
        """
        return pulumi.get(self, "data_vip")

    @property
    @pulumi.getter(name="engineName")
    def engine_name(self) -> Optional[str]:
        """
        Indicates the instance DB engine name.
        """
        return pulumi.get(self, "engine_name")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        """
        Indicates the instance DB engine version.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[str]:
        """
        Indicates the enterprise project ID.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="haMode")
    def ha_mode(self) -> Optional[str]:
        """
        Indicates the instance type.
        """
        return pulumi.get(self, "ha_mode")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetRecyclingInstancesInstanceResult']:
        """
        Indicates the list of recycling RDS instances.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="isServerless")
    def is_serverless(self) -> Optional[str]:
        """
        Indicates whether the instance is a serverless instance.
        """
        return pulumi.get(self, "is_serverless")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Indicates the instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="payModel")
    def pay_model(self) -> Optional[str]:
        """
        Indicates the instance billing mode.
        """
        return pulumi.get(self, "pay_model")

    @property
    @pulumi.getter(name="recycleBackupId")
    def recycle_backup_id(self) -> Optional[str]:
        """
        Indicates the backup ID.
        """
        return pulumi.get(self, "recycle_backup_id")

    @property
    @pulumi.getter(name="recycleStatus")
    def recycle_status(self) -> Optional[str]:
        """
        Indicates the backup status.
        """
        return pulumi.get(self, "recycle_status")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[str]:
        """
        Indicates the storage space in **GB**.
        """
        return pulumi.get(self, "volume_size")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[str]:
        """
        Indicates the storage type.
        """
        return pulumi.get(self, "volume_type")


class AwaitableGetRecyclingInstancesResult(GetRecyclingInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRecyclingInstancesResult(
            data_vip=self.data_vip,
            engine_name=self.engine_name,
            engine_version=self.engine_version,
            enterprise_project_id=self.enterprise_project_id,
            ha_mode=self.ha_mode,
            id=self.id,
            instance_id=self.instance_id,
            instances=self.instances,
            is_serverless=self.is_serverless,
            name=self.name,
            pay_model=self.pay_model,
            recycle_backup_id=self.recycle_backup_id,
            recycle_status=self.recycle_status,
            region=self.region,
            volume_size=self.volume_size,
            volume_type=self.volume_type)


def get_recycling_instances(data_vip: Optional[str] = None,
                            engine_name: Optional[str] = None,
                            engine_version: Optional[str] = None,
                            enterprise_project_id: Optional[str] = None,
                            ha_mode: Optional[str] = None,
                            instance_id: Optional[str] = None,
                            is_serverless: Optional[str] = None,
                            name: Optional[str] = None,
                            pay_model: Optional[str] = None,
                            recycle_backup_id: Optional[str] = None,
                            recycle_status: Optional[str] = None,
                            region: Optional[str] = None,
                            volume_size: Optional[str] = None,
                            volume_type: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRecyclingInstancesResult:
    """
    Use this data source to get the list of RDS recycling instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_recycling_instances(instance_id=instance_id)
    ```


    :param str data_vip: Specifies the floating IP address.
    :param str engine_name: Specifies the DB engine name.
    :param str engine_version: Specifies the DB engine version.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str ha_mode: Specifies the instance type.
           Value options: **Ha**, **Single**.
    :param str instance_id: Specifies the instance ID.
    :param str is_serverless: Specifies whether the instance is a serverless instance.
           Value options: **true**, **false**.
    :param str name: Specifies the instance name.
    :param str pay_model: Specifies the billing mode.
           Value options: **0** (pay-per-use), **1** (yearly/monthly).
    :param str recycle_backup_id: Specifies the backup ID.
    :param str recycle_status: Specifies the backup status.
           Value options:
           + **BUILDING**: The instance is being backed up and cannot be rebuilt.
           + **COMPLETED**: The backup is complete and the instance can be rebuilt.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str volume_size: Specifies the storage space in **GB**.
           The value must be a multiple of **10** and the value range is from **40 GB** to **4,000 GB**.
    :param str volume_type: Specifies the storage type.
           Value options:
           + **ULTRAHIGH**: ultra-high I/O storage.
           + **ULTRAHIGHPRO**: ultra-high I/O (advanced) storage.
           + **CLOUDSSD**: cloud SSD storage.
           + **LOCALSSD**: local SSD storage.
    """
    __args__ = dict()
    __args__['dataVip'] = data_vip
    __args__['engineName'] = engine_name
    __args__['engineVersion'] = engine_version
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['haMode'] = ha_mode
    __args__['instanceId'] = instance_id
    __args__['isServerless'] = is_serverless
    __args__['name'] = name
    __args__['payModel'] = pay_model
    __args__['recycleBackupId'] = recycle_backup_id
    __args__['recycleStatus'] = recycle_status
    __args__['region'] = region
    __args__['volumeSize'] = volume_size
    __args__['volumeType'] = volume_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getRecyclingInstances:getRecyclingInstances', __args__, opts=opts, typ=GetRecyclingInstancesResult).value

    return AwaitableGetRecyclingInstancesResult(
        data_vip=__ret__.data_vip,
        engine_name=__ret__.engine_name,
        engine_version=__ret__.engine_version,
        enterprise_project_id=__ret__.enterprise_project_id,
        ha_mode=__ret__.ha_mode,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instances=__ret__.instances,
        is_serverless=__ret__.is_serverless,
        name=__ret__.name,
        pay_model=__ret__.pay_model,
        recycle_backup_id=__ret__.recycle_backup_id,
        recycle_status=__ret__.recycle_status,
        region=__ret__.region,
        volume_size=__ret__.volume_size,
        volume_type=__ret__.volume_type)


@_utilities.lift_output_func(get_recycling_instances)
def get_recycling_instances_output(data_vip: Optional[pulumi.Input[Optional[str]]] = None,
                                   engine_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   engine_version: Optional[pulumi.Input[Optional[str]]] = None,
                                   enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   ha_mode: Optional[pulumi.Input[Optional[str]]] = None,
                                   instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   is_serverless: Optional[pulumi.Input[Optional[str]]] = None,
                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                   pay_model: Optional[pulumi.Input[Optional[str]]] = None,
                                   recycle_backup_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   recycle_status: Optional[pulumi.Input[Optional[str]]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   volume_size: Optional[pulumi.Input[Optional[str]]] = None,
                                   volume_type: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRecyclingInstancesResult]:
    """
    Use this data source to get the list of RDS recycling instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_recycling_instances(instance_id=instance_id)
    ```


    :param str data_vip: Specifies the floating IP address.
    :param str engine_name: Specifies the DB engine name.
    :param str engine_version: Specifies the DB engine version.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str ha_mode: Specifies the instance type.
           Value options: **Ha**, **Single**.
    :param str instance_id: Specifies the instance ID.
    :param str is_serverless: Specifies whether the instance is a serverless instance.
           Value options: **true**, **false**.
    :param str name: Specifies the instance name.
    :param str pay_model: Specifies the billing mode.
           Value options: **0** (pay-per-use), **1** (yearly/monthly).
    :param str recycle_backup_id: Specifies the backup ID.
    :param str recycle_status: Specifies the backup status.
           Value options:
           + **BUILDING**: The instance is being backed up and cannot be rebuilt.
           + **COMPLETED**: The backup is complete and the instance can be rebuilt.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str volume_size: Specifies the storage space in **GB**.
           The value must be a multiple of **10** and the value range is from **40 GB** to **4,000 GB**.
    :param str volume_type: Specifies the storage type.
           Value options:
           + **ULTRAHIGH**: ultra-high I/O storage.
           + **ULTRAHIGHPRO**: ultra-high I/O (advanced) storage.
           + **CLOUDSSD**: cloud SSD storage.
           + **LOCALSSD**: local SSD storage.
    """
    ...
