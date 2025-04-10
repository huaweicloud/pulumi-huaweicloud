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
    def __init__(__self__, engine_name=None, engine_version=None, flavors=None, id=None, memory=None, region=None, type=None, vcpus=None):
        if engine_name and not isinstance(engine_name, str):
            raise TypeError("Expected argument 'engine_name' to be a str")
        pulumi.set(__self__, "engine_name", engine_name)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if flavors and not isinstance(flavors, list):
            raise TypeError("Expected argument 'flavors' to be a list")
        pulumi.set(__self__, "flavors", flavors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if memory and not isinstance(memory, str):
            raise TypeError("Expected argument 'memory' to be a str")
        pulumi.set(__self__, "memory", memory)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vcpus and not isinstance(vcpus, str):
            raise TypeError("Expected argument 'vcpus' to be a str")
        pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter(name="engineName")
    def engine_name(self) -> str:
        """
        Indicates the engine name.
        """
        return pulumi.get(self, "engine_name")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def flavors(self) -> Sequence['outputs.GetFlavorsFlavorResult']:
        """
        Indicates the flavors information. Structure is documented below.
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
    def memory(self) -> Optional[str]:
        """
        Indicates the memory size in GB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Indicates the type of the flavor.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[str]:
        """
        Indicates the number of vCPUs.
        """
        return pulumi.get(self, "vcpus")


class AwaitableGetFlavorsResult(GetFlavorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorsResult(
            engine_name=self.engine_name,
            engine_version=self.engine_version,
            flavors=self.flavors,
            id=self.id,
            memory=self.memory,
            region=self.region,
            type=self.type,
            vcpus=self.vcpus)


def get_flavors(engine_name: Optional[str] = None,
                engine_version: Optional[str] = None,
                memory: Optional[str] = None,
                region: Optional[str] = None,
                type: Optional[str] = None,
                vcpus: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlavorsResult:
    """
    Use this data source to get the details of available DDS flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavor = huaweicloud.Dds.get_flavors(engine_name="DDS-Community",
        vcpus="8")
    ```


    :param str engine_name: Specifies the engine name. Value options: **DDS-Community** and **DDS-Enhanced**.
    :param str engine_version: Specifies the DB version number. Value options: **3.4**, **4.0**, **4.2** and **4.4**.
    :param str memory: Specifies the memory size in GB.
    :param str region: Specifies the region in which to obtain the flavors. If omitted,
           the provider-level region will be used.
    :param str type: Specifies the type of the flavor. Value options: **mongos**, **shard**, **config**,
           **replica**, **single** and **readonly**.
    :param str vcpus: Specifies the number of vCPUs.
    """
    __args__ = dict()
    __args__['engineName'] = engine_name
    __args__['engineVersion'] = engine_version
    __args__['memory'] = memory
    __args__['region'] = region
    __args__['type'] = type
    __args__['vcpus'] = vcpus
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Dds/getFlavors:getFlavors', __args__, opts=opts, typ=GetFlavorsResult).value

    return AwaitableGetFlavorsResult(
        engine_name=__ret__.engine_name,
        engine_version=__ret__.engine_version,
        flavors=__ret__.flavors,
        id=__ret__.id,
        memory=__ret__.memory,
        region=__ret__.region,
        type=__ret__.type,
        vcpus=__ret__.vcpus)


@_utilities.lift_output_func(get_flavors)
def get_flavors_output(engine_name: Optional[pulumi.Input[str]] = None,
                       engine_version: Optional[pulumi.Input[Optional[str]]] = None,
                       memory: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       type: Optional[pulumi.Input[Optional[str]]] = None,
                       vcpus: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlavorsResult]:
    """
    Use this data source to get the details of available DDS flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavor = huaweicloud.Dds.get_flavors(engine_name="DDS-Community",
        vcpus="8")
    ```


    :param str engine_name: Specifies the engine name. Value options: **DDS-Community** and **DDS-Enhanced**.
    :param str engine_version: Specifies the DB version number. Value options: **3.4**, **4.0**, **4.2** and **4.4**.
    :param str memory: Specifies the memory size in GB.
    :param str region: Specifies the region in which to obtain the flavors. If omitted,
           the provider-level region will be used.
    :param str type: Specifies the type of the flavor. Value options: **mongos**, **shard**, **config**,
           **replica**, **single** and **readonly**.
    :param str vcpus: Specifies the number of vCPUs.
    """
    ...
