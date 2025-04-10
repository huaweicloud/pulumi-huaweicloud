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
    'GetSubNetworkInterfacesResult',
    'AwaitableGetSubNetworkInterfacesResult',
    'get_sub_network_interfaces',
    'get_sub_network_interfaces_output',
]

@pulumi.output_type
class GetSubNetworkInterfacesResult:
    """
    A collection of values returned by getSubNetworkInterfaces.
    """
    def __init__(__self__, descriptions=None, id=None, interface_id=None, ip_address=None, mac_address=None, parent_id=None, region=None, sub_network_interfaces=None, subnet_id=None, vpc_id=None):
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        pulumi.set(__self__, "descriptions", descriptions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if sub_network_interfaces and not isinstance(sub_network_interfaces, list):
            raise TypeError("Expected argument 'sub_network_interfaces' to be a list")
        pulumi.set(__self__, "sub_network_interfaces", sub_network_interfaces)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def descriptions(self) -> Optional[Sequence[str]]:
        """
        The description of the supplementary network interface.
        """
        return pulumi.get(self, "descriptions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> Optional[str]:
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The private IPv4 address of the supplementary network interface.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        """
        The MAC address of the supplementary network interface.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[str]:
        """
        The ID of the elastic network interface to which the supplementary network interface belongs.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subNetworkInterfaces")
    def sub_network_interfaces(self) -> Sequence['outputs.GetSubNetworkInterfacesSubNetworkInterfaceResult']:
        """
        The list of supplementary network interfaces.
        """
        return pulumi.get(self, "sub_network_interfaces")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The ID of the subnet to which the supplementary network interface belongs.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The ID of the VPC to which the supplementary network interface belongs.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetSubNetworkInterfacesResult(GetSubNetworkInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubNetworkInterfacesResult(
            descriptions=self.descriptions,
            id=self.id,
            interface_id=self.interface_id,
            ip_address=self.ip_address,
            mac_address=self.mac_address,
            parent_id=self.parent_id,
            region=self.region,
            sub_network_interfaces=self.sub_network_interfaces,
            subnet_id=self.subnet_id,
            vpc_id=self.vpc_id)


def get_sub_network_interfaces(descriptions: Optional[Sequence[str]] = None,
                               interface_id: Optional[str] = None,
                               ip_address: Optional[str] = None,
                               mac_address: Optional[str] = None,
                               parent_id: Optional[str] = None,
                               region: Optional[str] = None,
                               subnet_id: Optional[str] = None,
                               vpc_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubNetworkInterfacesResult:
    """
    Use this data source to get a list of VPC supplementary network interfaces.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    parent_id = config.require_object("parentId")
    test = huaweicloud.Vpc.get_sub_network_interfaces(subnet_id=subnet_id,
        parent_id=parent_id)
    ```


    :param Sequence[str] descriptions: Specifies the description of the supplementary network interface.
    :param str interface_id: Specifies the ID of the supplementary network interface.
    :param str ip_address: Specifies the private IPv4 address of the supplementary network interface.
    :param str mac_address: Specifies the MAC address of the supplementary network interface.
    :param str parent_id: Specifies the ID of the elastic network interface
           to which the supplementary network interface belongs.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str subnet_id: Specifies the ID of the subnet to which the supplementary network interface belongs.
    :param str vpc_id: Specifies the ID of the VPC to which the supplementary network interface belongs.
    """
    __args__ = dict()
    __args__['descriptions'] = descriptions
    __args__['interfaceId'] = interface_id
    __args__['ipAddress'] = ip_address
    __args__['macAddress'] = mac_address
    __args__['parentId'] = parent_id
    __args__['region'] = region
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Vpc/getSubNetworkInterfaces:getSubNetworkInterfaces', __args__, opts=opts, typ=GetSubNetworkInterfacesResult).value

    return AwaitableGetSubNetworkInterfacesResult(
        descriptions=__ret__.descriptions,
        id=__ret__.id,
        interface_id=__ret__.interface_id,
        ip_address=__ret__.ip_address,
        mac_address=__ret__.mac_address,
        parent_id=__ret__.parent_id,
        region=__ret__.region,
        sub_network_interfaces=__ret__.sub_network_interfaces,
        subnet_id=__ret__.subnet_id,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_sub_network_interfaces)
def get_sub_network_interfaces_output(descriptions: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      interface_id: Optional[pulumi.Input[Optional[str]]] = None,
                                      ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                                      mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                                      parent_id: Optional[pulumi.Input[Optional[str]]] = None,
                                      region: Optional[pulumi.Input[Optional[str]]] = None,
                                      subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                                      vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubNetworkInterfacesResult]:
    """
    Use this data source to get a list of VPC supplementary network interfaces.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    parent_id = config.require_object("parentId")
    test = huaweicloud.Vpc.get_sub_network_interfaces(subnet_id=subnet_id,
        parent_id=parent_id)
    ```


    :param Sequence[str] descriptions: Specifies the description of the supplementary network interface.
    :param str interface_id: Specifies the ID of the supplementary network interface.
    :param str ip_address: Specifies the private IPv4 address of the supplementary network interface.
    :param str mac_address: Specifies the MAC address of the supplementary network interface.
    :param str parent_id: Specifies the ID of the elastic network interface
           to which the supplementary network interface belongs.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str subnet_id: Specifies the ID of the subnet to which the supplementary network interface belongs.
    :param str vpc_id: Specifies the ID of the VPC to which the supplementary network interface belongs.
    """
    ...
