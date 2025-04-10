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
    'GetCassandraDedicatedResourceResult',
    'AwaitableGetCassandraDedicatedResourceResult',
    'get_cassandra_dedicated_resource',
    'get_cassandra_dedicated_resource_output',
]

@pulumi.output_type
class GetCassandraDedicatedResourceResult:
    """
    A collection of values returned by getCassandraDedicatedResource.
    """
    def __init__(__self__, architecture=None, availability_zone=None, id=None, ram=None, region=None, resource_name=None, status=None, vcpus=None, volume=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ram and not isinstance(ram, int):
            raise TypeError("Expected argument 'ram' to be a int")
        pulumi.set(__self__, "ram", ram)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if resource_name and not isinstance(resource_name, str):
            raise TypeError("Expected argument 'resource_name' to be a str")
        pulumi.set(__self__, "resource_name", resource_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vcpus and not isinstance(vcpus, int):
            raise TypeError("Expected argument 'vcpus' to be a int")
        pulumi.set(__self__, "vcpus", vcpus)
        if volume and not isinstance(volume, int):
            raise TypeError("Expected argument 'volume' to be a int")
        pulumi.set(__self__, "volume", volume)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        """
        Indicates the architecture of the dedicated resource.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Indicates the availability zone of the dedicated resource.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ram(self) -> int:
        """
        Indicates the ram size of the dedicated resource.
        """
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> str:
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Indicates the status of the dedicated resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vcpus(self) -> int:
        """
        Indicates the vcpus count of the dedicated resource.
        """
        return pulumi.get(self, "vcpus")

    @property
    @pulumi.getter
    def volume(self) -> int:
        """
        Indicates the volume size of the dedicated resource.
        """
        return pulumi.get(self, "volume")


class AwaitableGetCassandraDedicatedResourceResult(GetCassandraDedicatedResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCassandraDedicatedResourceResult(
            architecture=self.architecture,
            availability_zone=self.availability_zone,
            id=self.id,
            ram=self.ram,
            region=self.region,
            resource_name=self.resource_name,
            status=self.status,
            vcpus=self.vcpus,
            volume=self.volume)


def get_cassandra_dedicated_resource(region: Optional[str] = None,
                                     resource_name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCassandraDedicatedResourceResult:
    """
    Use this data source to get available HuaweiCloud GeminiDB Cassandra dedicated resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_dedicated_resource(resource_name="test")
    ```


    :param str region: The region in which to obtain the dedicated resource. If omitted, the provider-level
           region will be used.
    :param str resource_name: Specifies the dedicated resource name.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:GaussDBforNoSQL/getCassandraDedicatedResource:getCassandraDedicatedResource', __args__, opts=opts, typ=GetCassandraDedicatedResourceResult).value

    return AwaitableGetCassandraDedicatedResourceResult(
        architecture=__ret__.architecture,
        availability_zone=__ret__.availability_zone,
        id=__ret__.id,
        ram=__ret__.ram,
        region=__ret__.region,
        resource_name=__ret__.resource_name,
        status=__ret__.status,
        vcpus=__ret__.vcpus,
        volume=__ret__.volume)


@_utilities.lift_output_func(get_cassandra_dedicated_resource)
def get_cassandra_dedicated_resource_output(region: Optional[pulumi.Input[Optional[str]]] = None,
                                            resource_name: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCassandraDedicatedResourceResult]:
    """
    Use this data source to get available HuaweiCloud GeminiDB Cassandra dedicated resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_dedicated_resource(resource_name="test")
    ```


    :param str region: The region in which to obtain the dedicated resource. If omitted, the provider-level
           region will be used.
    :param str resource_name: Specifies the dedicated resource name.
    """
    ...
