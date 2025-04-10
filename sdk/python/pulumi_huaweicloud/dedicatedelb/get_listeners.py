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
    'GetListenersResult',
    'AwaitableGetListenersResult',
    'get_listeners',
    'get_listeners_output',
]

@pulumi.output_type
class GetListenersResult:
    """
    A collection of values returned by getListeners.
    """
    def __init__(__self__, description=None, enterprise_project_id=None, id=None, listener_id=None, listeners=None, loadbalancer_id=None, name=None, protocol=None, protocol_port=None, region=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if listeners and not isinstance(listeners, list):
            raise TypeError("Expected argument 'listeners' to be a list")
        pulumi.set(__self__, "listeners", listeners)
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if protocol_port and not isinstance(protocol_port, int):
            raise TypeError("Expected argument 'protocol_port' to be a int")
        pulumi.set(__self__, "protocol_port", protocol_port)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the listener.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[str]:
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.GetListenersListenerResult']:
        """
        Lists the listeners.
        The listeners structure is documented below.
        """
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[str]:
        """
        The ID of the load balancer that the listener is added to.
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The listener name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        The protocol used by the listener.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> Optional[int]:
        """
        The port used by the listener.
        """
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetListenersResult(GetListenersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetListenersResult(
            description=self.description,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            listener_id=self.listener_id,
            listeners=self.listeners,
            loadbalancer_id=self.loadbalancer_id,
            name=self.name,
            protocol=self.protocol,
            protocol_port=self.protocol_port,
            region=self.region)


def get_listeners(description: Optional[str] = None,
                  enterprise_project_id: Optional[str] = None,
                  listener_id: Optional[str] = None,
                  loadbalancer_id: Optional[str] = None,
                  name: Optional[str] = None,
                  protocol: Optional[str] = None,
                  protocol_port: Optional[int] = None,
                  region: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetListenersResult:
    """
    Use this data source to get the list of ELB listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    listener_name = config.require_object("listenerName")
    test = huaweicloud.DedicatedElb.get_listeners(name=listener_name)
    ```


    :param str description: Specifies the description of the ELB listener.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str listener_id: Specifies the ID of the ELB listener.
    :param str loadbalancer_id: Specifies the ID of the load balancer that the listener is added to.
    :param str name: Specifies the name of the ELB listener.
    :param str protocol: Specifies the protocol of the ELB listener. Value options:
           **TCP**, **UDP**, **HTTP**, **HTTPS** or **QUIC**.
    :param int protocol_port: Specifies the port used by the listener.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['listenerId'] = listener_id
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    __args__['protocol'] = protocol
    __args__['protocolPort'] = protocol_port
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:DedicatedElb/getListeners:getListeners', __args__, opts=opts, typ=GetListenersResult).value

    return AwaitableGetListenersResult(
        description=__ret__.description,
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        listener_id=__ret__.listener_id,
        listeners=__ret__.listeners,
        loadbalancer_id=__ret__.loadbalancer_id,
        name=__ret__.name,
        protocol=__ret__.protocol,
        protocol_port=__ret__.protocol_port,
        region=__ret__.region)


@_utilities.lift_output_func(get_listeners)
def get_listeners_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                         enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                         listener_id: Optional[pulumi.Input[Optional[str]]] = None,
                         loadbalancer_id: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         protocol: Optional[pulumi.Input[Optional[str]]] = None,
                         protocol_port: Optional[pulumi.Input[Optional[int]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetListenersResult]:
    """
    Use this data source to get the list of ELB listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    listener_name = config.require_object("listenerName")
    test = huaweicloud.DedicatedElb.get_listeners(name=listener_name)
    ```


    :param str description: Specifies the description of the ELB listener.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str listener_id: Specifies the ID of the ELB listener.
    :param str loadbalancer_id: Specifies the ID of the load balancer that the listener is added to.
    :param str name: Specifies the name of the ELB listener.
    :param str protocol: Specifies the protocol of the ELB listener. Value options:
           **TCP**, **UDP**, **HTTP**, **HTTPS** or **QUIC**.
    :param int protocol_port: Specifies the port used by the listener.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    ...
