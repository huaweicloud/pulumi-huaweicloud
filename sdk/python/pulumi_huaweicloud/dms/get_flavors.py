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
    'GetFlavorsResult',
    'AwaitableGetFlavorsResult',
    'get_flavors',
    'get_flavors_output',
]

@pulumi.output_type
class GetFlavorsResult:
    """
    A collection of values returned by getFlavors.
    """
    def __init__(__self__, arch_type=None, availability_zones=None, charging_mode=None, flavor_id=None, flavors=None, id=None, region=None, storage_spec_code=None, type=None, versions=None):
        if arch_type and not isinstance(arch_type, str):
            raise TypeError("Expected argument 'arch_type' to be a str")
        pulumi.set(__self__, "arch_type", arch_type)
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError("Expected argument 'availability_zones' to be a list")
        pulumi.set(__self__, "availability_zones", availability_zones)
        if charging_mode and not isinstance(charging_mode, str):
            raise TypeError("Expected argument 'charging_mode' to be a str")
        pulumi.set(__self__, "charging_mode", charging_mode)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavors and not isinstance(flavors, list):
            raise TypeError("Expected argument 'flavors' to be a list")
        pulumi.set(__self__, "flavors", flavors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if storage_spec_code and not isinstance(storage_spec_code, str):
            raise TypeError("Expected argument 'storage_spec_code' to be a str")
        pulumi.set(__self__, "storage_spec_code", storage_spec_code)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)

    @property
    @pulumi.getter(name="archType")
    def arch_type(self) -> Optional[str]:
        return pulumi.get(self, "arch_type")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[Sequence[str]]:
        """
        The list of availability zones with available resources.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="chargingMode")
    def charging_mode(self) -> Optional[str]:
        return pulumi.get(self, "charging_mode")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[str]:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter
    def flavors(self) -> Sequence['outputs.GetFlavorsFlavorResult']:
        """
        The list of flavor details.
        The object structure is documented below.
        """
        return pulumi.get(self, "flavors")

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
    @pulumi.getter(name="storageSpecCode")
    def storage_spec_code(self) -> Optional[str]:
        """
        The disk IO encoding.
        """
        return pulumi.get(self, "storage_spec_code")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The disk type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def versions(self) -> Sequence[str]:
        """
        The supported flavor versions.
        """
        return pulumi.get(self, "versions")


class AwaitableGetFlavorsResult(GetFlavorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorsResult(
            arch_type=self.arch_type,
            availability_zones=self.availability_zones,
            charging_mode=self.charging_mode,
            flavor_id=self.flavor_id,
            flavors=self.flavors,
            id=self.id,
            region=self.region,
            storage_spec_code=self.storage_spec_code,
            type=self.type,
            versions=self.versions)


def get_flavors(arch_type: Optional[str] = None,
                availability_zones: Optional[Sequence[str]] = None,
                charging_mode: Optional[str] = None,
                flavor_id: Optional[str] = None,
                region: Optional[str] = None,
                storage_spec_code: Optional[str] = None,
                type: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlavorsResult:
    """
    Use this data source to get the list of available flavor details within HuaweiCloud.

    ## Example Usage
    ### Query the list of kafka flavors for cluster type

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Dms.get_flavors(type="cluster")
    ```
    ### Query the kafka flavor details of the specified ID

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Dms.get_flavors(flavor_id="c6.2u4g.cluster")
    ```
    ### Query list of kafka flavors that available in the availability zone list

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    az1 = config.require_object("az1")
    az2 = config.require_object("az2")
    test = huaweicloud.Dms.get_flavors(availability_zones=[
        az1,
        az2,
    ])
    ```


    :param str arch_type: Specifies the type of CPU architecture, e.g. **X86**.
    :param Sequence[str] availability_zones: Specifies the list of availability zones with available resources.
    :param str charging_mode: Specifies the flavor billing mode.
           The valid values are **prePaid** and **postPaid**.
    :param str flavor_id: Specifies the DMS flavor ID, e.g. **c6.2u4g.cluster**.
    :param str region: Specifies the region in which to obtain the dms kafka flavors.
           If omitted, the provider-level region will be used.
    :param str storage_spec_code: Specifies the disk IO encoding.
           + **dms.physical.storage.high.v2**: Type of the disk that uses high I/O.
           + **dms.physical.storage.ultra.v2**: Type of the disk that uses ultra-high I/O.
    :param str type: Specifies flavor type. The valid values are **single**, **cluster** and **cluster.small**.
    """
    __args__ = dict()
    __args__['archType'] = arch_type
    __args__['availabilityZones'] = availability_zones
    __args__['chargingMode'] = charging_mode
    __args__['flavorId'] = flavor_id
    __args__['region'] = region
    __args__['storageSpecCode'] = storage_spec_code
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Dms/getFlavors:getFlavors', __args__, opts=opts, typ=GetFlavorsResult).value

    return AwaitableGetFlavorsResult(
        arch_type=__ret__.arch_type,
        availability_zones=__ret__.availability_zones,
        charging_mode=__ret__.charging_mode,
        flavor_id=__ret__.flavor_id,
        flavors=__ret__.flavors,
        id=__ret__.id,
        region=__ret__.region,
        storage_spec_code=__ret__.storage_spec_code,
        type=__ret__.type,
        versions=__ret__.versions)


@_utilities.lift_output_func(get_flavors)
def get_flavors_output(arch_type: Optional[pulumi.Input[Optional[str]]] = None,
                       availability_zones: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       charging_mode: Optional[pulumi.Input[Optional[str]]] = None,
                       flavor_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       storage_spec_code: Optional[pulumi.Input[Optional[str]]] = None,
                       type: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlavorsResult]:
    """
    Use this data source to get the list of available flavor details within HuaweiCloud.

    ## Example Usage
    ### Query the list of kafka flavors for cluster type

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Dms.get_flavors(type="cluster")
    ```
    ### Query the kafka flavor details of the specified ID

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.Dms.get_flavors(flavor_id="c6.2u4g.cluster")
    ```
    ### Query list of kafka flavors that available in the availability zone list

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    az1 = config.require_object("az1")
    az2 = config.require_object("az2")
    test = huaweicloud.Dms.get_flavors(availability_zones=[
        az1,
        az2,
    ])
    ```


    :param str arch_type: Specifies the type of CPU architecture, e.g. **X86**.
    :param Sequence[str] availability_zones: Specifies the list of availability zones with available resources.
    :param str charging_mode: Specifies the flavor billing mode.
           The valid values are **prePaid** and **postPaid**.
    :param str flavor_id: Specifies the DMS flavor ID, e.g. **c6.2u4g.cluster**.
    :param str region: Specifies the region in which to obtain the dms kafka flavors.
           If omitted, the provider-level region will be used.
    :param str storage_spec_code: Specifies the disk IO encoding.
           + **dms.physical.storage.high.v2**: Type of the disk that uses high I/O.
           + **dms.physical.storage.ultra.v2**: Type of the disk that uses ultra-high I/O.
    :param str type: Specifies flavor type. The valid values are **single**, **cluster** and **cluster.small**.
    """
    ...
