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
    'PoolPersistence',
    'GetListenersListenerResult',
    'GetListenersListenerLoadbalancerResult',
    'GetPoolsPoolResult',
    'GetPoolsPoolListenerResult',
    'GetPoolsPoolLoadbalancerResult',
    'GetPoolsPoolMemberResult',
    'GetPoolsPoolPersistenceResult',
]

@pulumi.output_type
class PoolPersistence(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cookieName":
            suggest = "cookie_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PoolPersistence. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PoolPersistence.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PoolPersistence.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 cookie_name: Optional[str] = None,
                 timeout: Optional[int] = None):
        """
        :param str type: The type of persistence mode. The current specification supports SOURCE_IP,
               HTTP_COOKIE, and APP_COOKIE.
        :param str cookie_name: The name of the cookie if persistence mode is set appropriately. Required
               if `type = APP_COOKIE`.
        :param int timeout: Specifies the sticky session timeout duration in minutes. This parameter is
               invalid when type is set to APP_COOKIE. The value range varies depending on the protocol of the backend server group:
               + When the protocol of the backend server group is TCP or UDP, the value ranges from 1 to 60.
               + When the protocol of the backend server group is HTTP or HTTPS, the value ranges from 1 to 1440.
        """
        pulumi.set(__self__, "type", type)
        if cookie_name is not None:
            pulumi.set(__self__, "cookie_name", cookie_name)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of persistence mode. The current specification supports SOURCE_IP,
        HTTP_COOKIE, and APP_COOKIE.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> Optional[str]:
        """
        The name of the cookie if persistence mode is set appropriately. Required
        if `type = APP_COOKIE`.
        """
        return pulumi.get(self, "cookie_name")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        """
        Specifies the sticky session timeout duration in minutes. This parameter is
        invalid when type is set to APP_COOKIE. The value range varies depending on the protocol of the backend server group:
        + When the protocol of the backend server group is TCP or UDP, the value ranges from 1 to 60.
        + When the protocol of the backend server group is HTTP or HTTPS, the value ranges from 1 to 1440.
        """
        return pulumi.get(self, "timeout")


@pulumi.output_type
class GetListenersListenerResult(dict):
    def __init__(__self__, *,
                 connection_limit: int,
                 default_pool_id: str,
                 default_tls_container_ref: str,
                 description: str,
                 http2_enable: bool,
                 id: str,
                 loadbalancers: Sequence['outputs.GetListenersListenerLoadbalancerResult'],
                 name: str,
                 protocol: str,
                 protocol_port: int,
                 sni_container_refs: Sequence[str]):
        """
        :param int connection_limit: The maximum number of connections allowed for the listener.
        :param str default_pool_id: The ID of the default pool with which the ELB listener is associated.
        :param str default_tls_container_ref: The ID of the server certificate used by the listener.
        :param str description: The description of the ELB listener.
        :param bool http2_enable: Whether the ELB listener uses HTTP/2.
        :param str id: The ELB loadbalancer ID.
        :param Sequence['GetListenersListenerLoadbalancerArgs'] loadbalancers: Listener list.
               The object structure is documented below.
        :param str name: The listener name.
        :param str protocol: The listener protocol.  
               The valid values are **TCP**, **UDP**, **HTTP** and **TERMINATED_HTTPS**.
        :param int protocol_port: The front-end listening port of the listener.  
               The valid value is range from `1` to `65535`.
        :param Sequence[str] sni_container_refs: List of the SNI certificate (server certificates with a domain name) IDs used by the listener.
        """
        pulumi.set(__self__, "connection_limit", connection_limit)
        pulumi.set(__self__, "default_pool_id", default_pool_id)
        pulumi.set(__self__, "default_tls_container_ref", default_tls_container_ref)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "http2_enable", http2_enable)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "loadbalancers", loadbalancers)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "protocol_port", protocol_port)
        pulumi.set(__self__, "sni_container_refs", sni_container_refs)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> int:
        """
        The maximum number of connections allowed for the listener.
        """
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter(name="defaultPoolId")
    def default_pool_id(self) -> str:
        """
        The ID of the default pool with which the ELB listener is associated.
        """
        return pulumi.get(self, "default_pool_id")

    @property
    @pulumi.getter(name="defaultTlsContainerRef")
    def default_tls_container_ref(self) -> str:
        """
        The ID of the server certificate used by the listener.
        """
        return pulumi.get(self, "default_tls_container_ref")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the ELB listener.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="http2Enable")
    def http2_enable(self) -> bool:
        """
        Whether the ELB listener uses HTTP/2.
        """
        return pulumi.get(self, "http2_enable")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ELB loadbalancer ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def loadbalancers(self) -> Sequence['outputs.GetListenersListenerLoadbalancerResult']:
        """
        Listener list.
        The object structure is documented below.
        """
        return pulumi.get(self, "loadbalancers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The listener name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The listener protocol.  
        The valid values are **TCP**, **UDP**, **HTTP** and **TERMINATED_HTTPS**.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> int:
        """
        The front-end listening port of the listener.  
        The valid value is range from `1` to `65535`.
        """
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter(name="sniContainerRefs")
    def sni_container_refs(self) -> Sequence[str]:
        """
        List of the SNI certificate (server certificates with a domain name) IDs used by the listener.
        """
        return pulumi.get(self, "sni_container_refs")


@pulumi.output_type
class GetListenersListenerLoadbalancerResult(dict):
    def __init__(__self__, *,
                 id: str):
        """
        :param str id: The ELB loadbalancer ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ELB loadbalancer ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPoolsPoolResult(dict):
    def __init__(__self__, *,
                 description: str,
                 healthmonitor_id: str,
                 id: str,
                 lb_method: str,
                 listeners: Sequence['outputs.GetPoolsPoolListenerResult'],
                 loadbalancers: Sequence['outputs.GetPoolsPoolLoadbalancerResult'],
                 members: Sequence['outputs.GetPoolsPoolMemberResult'],
                 name: str,
                 persistences: Sequence['outputs.GetPoolsPoolPersistenceResult'],
                 protocol: str):
        """
        :param str description: Specifies the description of the ELB pool.
        :param str healthmonitor_id: Specifies the health monitor ID of the ELB pool.
        :param str id: The listener, loadbalancer or member ID.
        :param str lb_method: Specifies the method of the ELB pool. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
               or SOURCE_IP.
        :param Sequence['GetPoolsPoolListenerArgs'] listeners: The listener list. The object structure is documented below.
        :param Sequence['GetPoolsPoolLoadbalancerArgs'] loadbalancers: The loadbalancer list. The object structure is documented below.
        :param Sequence['GetPoolsPoolMemberArgs'] members: The member list. The object structure is documented below.
        :param str name: Specifies the name of the ELB pool.
        :param Sequence['GetPoolsPoolPersistenceArgs'] persistences: Indicates whether connections in the same session will be processed by the same pool member or not.
               The object structure is documented below.
        :param str protocol: Specifies the protocol of the ELB pool. This can either be TCP, UDP or HTTP.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "healthmonitor_id", healthmonitor_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "lb_method", lb_method)
        pulumi.set(__self__, "listeners", listeners)
        pulumi.set(__self__, "loadbalancers", loadbalancers)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "persistences", persistences)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Specifies the description of the ELB pool.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="healthmonitorId")
    def healthmonitor_id(self) -> str:
        """
        Specifies the health monitor ID of the ELB pool.
        """
        return pulumi.get(self, "healthmonitor_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The listener, loadbalancer or member ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> str:
        """
        Specifies the method of the ELB pool. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
        or SOURCE_IP.
        """
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.GetPoolsPoolListenerResult']:
        """
        The listener list. The object structure is documented below.
        """
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter
    def loadbalancers(self) -> Sequence['outputs.GetPoolsPoolLoadbalancerResult']:
        """
        The loadbalancer list. The object structure is documented below.
        """
        return pulumi.get(self, "loadbalancers")

    @property
    @pulumi.getter
    def members(self) -> Sequence['outputs.GetPoolsPoolMemberResult']:
        """
        The member list. The object structure is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the ELB pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def persistences(self) -> Sequence['outputs.GetPoolsPoolPersistenceResult']:
        """
        Indicates whether connections in the same session will be processed by the same pool member or not.
        The object structure is documented below.
        """
        return pulumi.get(self, "persistences")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Specifies the protocol of the ELB pool. This can either be TCP, UDP or HTTP.
        """
        return pulumi.get(self, "protocol")


@pulumi.output_type
class GetPoolsPoolListenerResult(dict):
    def __init__(__self__, *,
                 id: str):
        """
        :param str id: The listener, loadbalancer or member ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The listener, loadbalancer or member ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPoolsPoolLoadbalancerResult(dict):
    def __init__(__self__, *,
                 id: str):
        """
        :param str id: The listener, loadbalancer or member ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The listener, loadbalancer or member ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPoolsPoolMemberResult(dict):
    def __init__(__self__, *,
                 id: str):
        """
        :param str id: The listener, loadbalancer or member ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The listener, loadbalancer or member ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPoolsPoolPersistenceResult(dict):
    def __init__(__self__, *,
                 cookie_name: str,
                 type: str):
        """
        :param str cookie_name: The name of the cookie if persistence mode is set appropriately.
        :param str type: The type of persistence mode.
        """
        pulumi.set(__self__, "cookie_name", cookie_name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> str:
        """
        The name of the cookie if persistence mode is set appropriately.
        """
        return pulumi.get(self, "cookie_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of persistence mode.
        """
        return pulumi.get(self, "type")


