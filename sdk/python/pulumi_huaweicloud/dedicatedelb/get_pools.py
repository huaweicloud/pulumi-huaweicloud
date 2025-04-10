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
    'GetPoolsResult',
    'AwaitableGetPoolsResult',
    'get_pools',
    'get_pools_output',
]

@pulumi.output_type
class GetPoolsResult:
    """
    A collection of values returned by getPools.
    """
    def __init__(__self__, description=None, healthmonitor_id=None, id=None, lb_method=None, listener_id=None, loadbalancer_id=None, name=None, pool_id=None, pools=None, protection_status=None, protocol=None, region=None, type=None, vpc_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if healthmonitor_id and not isinstance(healthmonitor_id, str):
            raise TypeError("Expected argument 'healthmonitor_id' to be a str")
        pulumi.set(__self__, "healthmonitor_id", healthmonitor_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lb_method and not isinstance(lb_method, str):
            raise TypeError("Expected argument 'lb_method' to be a str")
        pulumi.set(__self__, "lb_method", lb_method)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        pulumi.set(__self__, "pool_id", pool_id)
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)
        if protection_status and not isinstance(protection_status, str):
            raise TypeError("Expected argument 'protection_status' to be a str")
        pulumi.set(__self__, "protection_status", protection_status)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of pool.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="healthmonitorId")
    def healthmonitor_id(self) -> Optional[str]:
        """
        The health monitor ID of the LB pool.
        """
        return pulumi.get(self, "healthmonitor_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> Optional[str]:
        """
        The load balancing algorithm to distribute traffic to the pool's members.
        """
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[str]:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The pool name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[str]:
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def pools(self) -> Sequence['outputs.GetPoolsPoolResult']:
        """
        Pool list. For details, see data structure of the pool field.
        The object structure is documented below.
        """
        return pulumi.get(self, "pools")

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> Optional[str]:
        """
        The protection status for update.
        """
        return pulumi.get(self, "protection_status")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        The protocol of pool.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of persistence mode.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The ID of the VPC where the backend server group works.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetPoolsResult(GetPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolsResult(
            description=self.description,
            healthmonitor_id=self.healthmonitor_id,
            id=self.id,
            lb_method=self.lb_method,
            listener_id=self.listener_id,
            loadbalancer_id=self.loadbalancer_id,
            name=self.name,
            pool_id=self.pool_id,
            pools=self.pools,
            protection_status=self.protection_status,
            protocol=self.protocol,
            region=self.region,
            type=self.type,
            vpc_id=self.vpc_id)


def get_pools(description: Optional[str] = None,
              healthmonitor_id: Optional[str] = None,
              lb_method: Optional[str] = None,
              listener_id: Optional[str] = None,
              loadbalancer_id: Optional[str] = None,
              name: Optional[str] = None,
              pool_id: Optional[str] = None,
              protection_status: Optional[str] = None,
              protocol: Optional[str] = None,
              region: Optional[str] = None,
              type: Optional[str] = None,
              vpc_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoolsResult:
    """
    Use this data source to get the list of ELB pools.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    pool_name = config.require_object("poolName")
    test = huaweicloud.DedicatedElb.get_pools(name=pool_name)
    ```


    :param str description: Specifies the description of the ELB pool.
    :param str healthmonitor_id: Specifies the health monitor ID of the ELB pool.
    :param str lb_method: Specifies the method of the ELB pool. Value options: **ROUND_ROBIN**,
           **LEAST_CONNECTIONS**, **SOURCE_IP** or **QUIC_CID**.
    :param str listener_id: Specifies the listener ID of the ELB pool.
    :param str loadbalancer_id: Specifies the loadbalancer ID of the ELB pool.
    :param str name: Specifies the name of the ELB pool.
    :param str pool_id: Specifies the ID of the ELB pool.
    :param str protection_status: Specifies the protection status for update.
           Value options: **nonProtection**, **consoleProtection**.
    :param str protocol: Specifies the protocol of the ELB pool. Value options: **TCP**, **UDP**, **HTTP**,
           **HTTPS**, **QUIC**, **GRPC** or **TLS**.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param str type: Specifies the type of the backend server group. Value options: **instance**, **ip**.
    :param str vpc_id: Specifies the ID of the VPC where the backend server group works.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['healthmonitorId'] = healthmonitor_id
    __args__['lbMethod'] = lb_method
    __args__['listenerId'] = listener_id
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    __args__['poolId'] = pool_id
    __args__['protectionStatus'] = protection_status
    __args__['protocol'] = protocol
    __args__['region'] = region
    __args__['type'] = type
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:DedicatedElb/getPools:getPools', __args__, opts=opts, typ=GetPoolsResult).value

    return AwaitableGetPoolsResult(
        description=__ret__.description,
        healthmonitor_id=__ret__.healthmonitor_id,
        id=__ret__.id,
        lb_method=__ret__.lb_method,
        listener_id=__ret__.listener_id,
        loadbalancer_id=__ret__.loadbalancer_id,
        name=__ret__.name,
        pool_id=__ret__.pool_id,
        pools=__ret__.pools,
        protection_status=__ret__.protection_status,
        protocol=__ret__.protocol,
        region=__ret__.region,
        type=__ret__.type,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_pools)
def get_pools_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                     healthmonitor_id: Optional[pulumi.Input[Optional[str]]] = None,
                     lb_method: Optional[pulumi.Input[Optional[str]]] = None,
                     listener_id: Optional[pulumi.Input[Optional[str]]] = None,
                     loadbalancer_id: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     pool_id: Optional[pulumi.Input[Optional[str]]] = None,
                     protection_status: Optional[pulumi.Input[Optional[str]]] = None,
                     protocol: Optional[pulumi.Input[Optional[str]]] = None,
                     region: Optional[pulumi.Input[Optional[str]]] = None,
                     type: Optional[pulumi.Input[Optional[str]]] = None,
                     vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoolsResult]:
    """
    Use this data source to get the list of ELB pools.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    pool_name = config.require_object("poolName")
    test = huaweicloud.DedicatedElb.get_pools(name=pool_name)
    ```


    :param str description: Specifies the description of the ELB pool.
    :param str healthmonitor_id: Specifies the health monitor ID of the ELB pool.
    :param str lb_method: Specifies the method of the ELB pool. Value options: **ROUND_ROBIN**,
           **LEAST_CONNECTIONS**, **SOURCE_IP** or **QUIC_CID**.
    :param str listener_id: Specifies the listener ID of the ELB pool.
    :param str loadbalancer_id: Specifies the loadbalancer ID of the ELB pool.
    :param str name: Specifies the name of the ELB pool.
    :param str pool_id: Specifies the ID of the ELB pool.
    :param str protection_status: Specifies the protection status for update.
           Value options: **nonProtection**, **consoleProtection**.
    :param str protocol: Specifies the protocol of the ELB pool. Value options: **TCP**, **UDP**, **HTTP**,
           **HTTPS**, **QUIC**, **GRPC** or **TLS**.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param str type: Specifies the type of the backend server group. Value options: **instance**, **ip**.
    :param str vpc_id: Specifies the ID of the VPC where the backend server group works.
    """
    ...
