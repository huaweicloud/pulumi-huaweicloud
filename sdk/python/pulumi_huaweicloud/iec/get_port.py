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
    'GetPortResult',
    'AwaitableGetPortResult',
    'get_port',
    'get_port_output',
]

@pulumi.output_type
class GetPortResult:
    """
    A collection of values returned by getPort.
    """
    def __init__(__self__, fixed_ip=None, id=None, mac_address=None, region=None, security_groups=None, site_id=None, status=None, subnet_id=None):
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        pulumi.set(__self__, "fixed_ip", fixed_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if site_id and not isinstance(site_id, str):
            raise TypeError("Expected argument 'site_id' to be a str")
        pulumi.set(__self__, "site_id", site_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> str:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Specifies a data source ID in UUID format.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        """
        Indicates the list of security group IDs applied on the port.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> str:
        """
        Indicates the ID of the IEC site.
        """
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Indicates the status of the port.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")


class AwaitableGetPortResult(GetPortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortResult(
            fixed_ip=self.fixed_ip,
            id=self.id,
            mac_address=self.mac_address,
            region=self.region,
            security_groups=self.security_groups,
            site_id=self.site_id,
            status=self.status,
            subnet_id=self.subnet_id)


def get_port(fixed_ip: Optional[str] = None,
             id: Optional[str] = None,
             mac_address: Optional[str] = None,
             region: Optional[str] = None,
             subnet_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortResult:
    """
    Use this data source to get the details of a specific IEC subnet port.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    port1 = huaweicloud.Iec.get_port(subnet_id=subnet_id,
        fixed_ip="192.168.1.123")
    ```


    :param str fixed_ip: The IP address of the port.
    :param str id: The ID of the port.
    :param str mac_address: The MAC address of the port.
    :param str region: The region in which to obtain the port. If omitted, the provider-level region will be
           used.
    :param str subnet_id: The ID of the subnet which the port belongs to.
    """
    __args__ = dict()
    __args__['fixedIp'] = fixed_ip
    __args__['id'] = id
    __args__['macAddress'] = mac_address
    __args__['region'] = region
    __args__['subnetId'] = subnet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Iec/getPort:getPort', __args__, opts=opts, typ=GetPortResult).value

    return AwaitableGetPortResult(
        fixed_ip=__ret__.fixed_ip,
        id=__ret__.id,
        mac_address=__ret__.mac_address,
        region=__ret__.region,
        security_groups=__ret__.security_groups,
        site_id=__ret__.site_id,
        status=__ret__.status,
        subnet_id=__ret__.subnet_id)


@_utilities.lift_output_func(get_port)
def get_port_output(fixed_ip: Optional[pulumi.Input[Optional[str]]] = None,
                    id: Optional[pulumi.Input[Optional[str]]] = None,
                    mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortResult]:
    """
    Use this data source to get the details of a specific IEC subnet port.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    port1 = huaweicloud.Iec.get_port(subnet_id=subnet_id,
        fixed_ip="192.168.1.123")
    ```


    :param str fixed_ip: The IP address of the port.
    :param str id: The ID of the port.
    :param str mac_address: The MAC address of the port.
    :param str region: The region in which to obtain the port. If omitted, the provider-level region will be
           used.
    :param str subnet_id: The ID of the subnet which the port belongs to.
    """
    ...