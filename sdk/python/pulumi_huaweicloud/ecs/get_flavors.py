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
    def __init__(__self__, availability_zone=None, cpu_core_count=None, generation=None, id=None, ids=None, memory_size=None, performance_type=None, region=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if cpu_core_count and not isinstance(cpu_core_count, int):
            raise TypeError("Expected argument 'cpu_core_count' to be a int")
        pulumi.set(__self__, "cpu_core_count", cpu_core_count)
        if generation and not isinstance(generation, str):
            raise TypeError("Expected argument 'generation' to be a str")
        pulumi.set(__self__, "generation", generation)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if memory_size and not isinstance(memory_size, int):
            raise TypeError("Expected argument 'memory_size' to be a int")
        pulumi.set(__self__, "memory_size", memory_size)
        if performance_type and not isinstance(performance_type, str):
            raise TypeError("Expected argument 'performance_type' to be a str")
        pulumi.set(__self__, "performance_type", performance_type)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cpuCoreCount")
    def cpu_core_count(self) -> Optional[int]:
        return pulumi.get(self, "cpu_core_count")

    @property
    @pulumi.getter
    def generation(self) -> Optional[str]:
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        A list of flavor IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> Optional[int]:
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter(name="performanceType")
    def performance_type(self) -> Optional[str]:
        return pulumi.get(self, "performance_type")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetFlavorsResult(GetFlavorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorsResult(
            availability_zone=self.availability_zone,
            cpu_core_count=self.cpu_core_count,
            generation=self.generation,
            id=self.id,
            ids=self.ids,
            memory_size=self.memory_size,
            performance_type=self.performance_type,
            region=self.region)


def get_flavors(availability_zone: Optional[str] = None,
                cpu_core_count: Optional[int] = None,
                generation: Optional[str] = None,
                memory_size: Optional[int] = None,
                performance_type: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlavorsResult:
    """
    Use this data source to get the available Compute Flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavors = huaweicloud.Ecs.get_flavors(availability_zone="cn-north-1a",
        performance_type="normal",
        cpu_core_count=2,
        memory_size=4)
    # Create ECS instance with the first matched flavor
    instance = huaweicloud.ecs.Instance("instance", flavor_id=flavors.ids[0])
    # Other properties...
    ```


    :param str availability_zone: Specifies the AZ name.
    :param int cpu_core_count: Specifies the number of vCPUs in the ECS flavor.
    :param str generation: Specifies the generation of an ECS type. For example, **s3** indicates
           the general-purpose third-generation ECSs. For details, see
           [ECS Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-ecs/ecs_01_0014.html).
    :param int memory_size: Specifies the memory size(GB) in the ECS flavor.
    :param str performance_type: Specifies the ECS flavor type. Possible values are as follows:
           + **normal**: General computing
           + **computingv3**: General computing-plus
           + **highmem**: Memory-optimized
           + **saphana**: Large-memory HANA ECS
           + **diskintensive**: Disk-intensive
    :param str region: The region in which to obtain the flavors.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['cpuCoreCount'] = cpu_core_count
    __args__['generation'] = generation
    __args__['memorySize'] = memory_size
    __args__['performanceType'] = performance_type
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Ecs/getFlavors:getFlavors', __args__, opts=opts, typ=GetFlavorsResult).value

    return AwaitableGetFlavorsResult(
        availability_zone=__ret__.availability_zone,
        cpu_core_count=__ret__.cpu_core_count,
        generation=__ret__.generation,
        id=__ret__.id,
        ids=__ret__.ids,
        memory_size=__ret__.memory_size,
        performance_type=__ret__.performance_type,
        region=__ret__.region)


@_utilities.lift_output_func(get_flavors)
def get_flavors_output(availability_zone: Optional[pulumi.Input[Optional[str]]] = None,
                       cpu_core_count: Optional[pulumi.Input[Optional[int]]] = None,
                       generation: Optional[pulumi.Input[Optional[str]]] = None,
                       memory_size: Optional[pulumi.Input[Optional[int]]] = None,
                       performance_type: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlavorsResult]:
    """
    Use this data source to get the available Compute Flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavors = huaweicloud.Ecs.get_flavors(availability_zone="cn-north-1a",
        performance_type="normal",
        cpu_core_count=2,
        memory_size=4)
    # Create ECS instance with the first matched flavor
    instance = huaweicloud.ecs.Instance("instance", flavor_id=flavors.ids[0])
    # Other properties...
    ```


    :param str availability_zone: Specifies the AZ name.
    :param int cpu_core_count: Specifies the number of vCPUs in the ECS flavor.
    :param str generation: Specifies the generation of an ECS type. For example, **s3** indicates
           the general-purpose third-generation ECSs. For details, see
           [ECS Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-ecs/ecs_01_0014.html).
    :param int memory_size: Specifies the memory size(GB) in the ECS flavor.
    :param str performance_type: Specifies the ECS flavor type. Possible values are as follows:
           + **normal**: General computing
           + **computingv3**: General computing-plus
           + **highmem**: Memory-optimized
           + **saphana**: Large-memory HANA ECS
           + **diskintensive**: Disk-intensive
    :param str region: The region in which to obtain the flavors.
           If omitted, the provider-level region will be used.
    """
    ...