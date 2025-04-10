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
    'GetCassandraInstancesResult',
    'AwaitableGetCassandraInstancesResult',
    'get_cassandra_instances',
    'get_cassandra_instances_output',
]

@pulumi.output_type
class GetCassandraInstancesResult:
    """
    A collection of values returned by getCassandraInstances.
    """
    def __init__(__self__, id=None, instances=None, name=None, region=None, subnet_id=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetCassandraInstancesInstanceResult']:
        """
        An array of available instances.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Indicates the node name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region of the instance.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Indicates the network ID of a subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        Indicates the VPC ID.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetCassandraInstancesResult(GetCassandraInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCassandraInstancesResult(
            id=self.id,
            instances=self.instances,
            name=self.name,
            region=self.region,
            subnet_id=self.subnet_id,
            vpc_id=self.vpc_id)


def get_cassandra_instances(name: Optional[str] = None,
                            region: Optional[str] = None,
                            subnet_id: Optional[str] = None,
                            vpc_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCassandraInstancesResult:
    """
    Use this data source to get available HuaweiCloud GeminiDB Cassandra instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_instances(name="gaussdb-instance")
    ```


    :param str name: Specifies the name of the instance.
    :param str region: The region in which to obtain the instance. If omitted, the provider-level region will
           be used.
    :param str subnet_id: Specifies the network ID of a subnet.
    :param str vpc_id: Specifies the VPC ID.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:GaussDBforNoSQL/getCassandraInstances:getCassandraInstances', __args__, opts=opts, typ=GetCassandraInstancesResult).value

    return AwaitableGetCassandraInstancesResult(
        id=__ret__.id,
        instances=__ret__.instances,
        name=__ret__.name,
        region=__ret__.region,
        subnet_id=__ret__.subnet_id,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_cassandra_instances)
def get_cassandra_instances_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCassandraInstancesResult]:
    """
    Use this data source to get available HuaweiCloud GeminiDB Cassandra instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDBforNoSQL.get_cassandra_instances(name="gaussdb-instance")
    ```


    :param str name: Specifies the name of the instance.
    :param str region: The region in which to obtain the instance. If omitted, the provider-level region will
           be used.
    :param str subnet_id: Specifies the network ID of a subnet.
    :param str vpc_id: Specifies the VPC ID.
    """
    ...
