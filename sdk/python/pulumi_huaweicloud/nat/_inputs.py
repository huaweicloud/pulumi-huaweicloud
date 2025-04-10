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
    'GatewaySessionConfArgs',
]

@pulumi.input_type
class GatewaySessionConfArgs:
    def __init__(__self__, *,
                 icmp_session_expire_time: Optional[pulumi.Input[int]] = None,
                 tcp_session_expire_time: Optional[pulumi.Input[int]] = None,
                 tcp_time_wait_time: Optional[pulumi.Input[int]] = None,
                 udp_session_expire_time: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] icmp_session_expire_time: Specifies the ICMP session expiration time, in seconds.
               The valid value from `10` to `7,200`, default value is `10`.
        :param pulumi.Input[int] tcp_session_expire_time: Specifies the TCP session expiration time, in seconds.
               The valid value from `40` to `7,200`, default value is `900`.
        :param pulumi.Input[int] tcp_time_wait_time: Specifies the duration of TIME_WAIT state when TCP connection is closed,
               in seconds. The valid value from `0` to `1,800`, default value is `5`.
        :param pulumi.Input[int] udp_session_expire_time: Specifies the UDP session expiration time, in seconds.
               The valid value from `40` to `7,200`, default value is `300`.
        """
        if icmp_session_expire_time is not None:
            pulumi.set(__self__, "icmp_session_expire_time", icmp_session_expire_time)
        if tcp_session_expire_time is not None:
            pulumi.set(__self__, "tcp_session_expire_time", tcp_session_expire_time)
        if tcp_time_wait_time is not None:
            pulumi.set(__self__, "tcp_time_wait_time", tcp_time_wait_time)
        if udp_session_expire_time is not None:
            pulumi.set(__self__, "udp_session_expire_time", udp_session_expire_time)

    @property
    @pulumi.getter(name="icmpSessionExpireTime")
    def icmp_session_expire_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the ICMP session expiration time, in seconds.
        The valid value from `10` to `7,200`, default value is `10`.
        """
        return pulumi.get(self, "icmp_session_expire_time")

    @icmp_session_expire_time.setter
    def icmp_session_expire_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icmp_session_expire_time", value)

    @property
    @pulumi.getter(name="tcpSessionExpireTime")
    def tcp_session_expire_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the TCP session expiration time, in seconds.
        The valid value from `40` to `7,200`, default value is `900`.
        """
        return pulumi.get(self, "tcp_session_expire_time")

    @tcp_session_expire_time.setter
    def tcp_session_expire_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tcp_session_expire_time", value)

    @property
    @pulumi.getter(name="tcpTimeWaitTime")
    def tcp_time_wait_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the duration of TIME_WAIT state when TCP connection is closed,
        in seconds. The valid value from `0` to `1,800`, default value is `5`.
        """
        return pulumi.get(self, "tcp_time_wait_time")

    @tcp_time_wait_time.setter
    def tcp_time_wait_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tcp_time_wait_time", value)

    @property
    @pulumi.getter(name="udpSessionExpireTime")
    def udp_session_expire_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the UDP session expiration time, in seconds.
        The valid value from `40` to `7,200`, default value is `300`.
        """
        return pulumi.get(self, "udp_session_expire_time")

    @udp_session_expire_time.setter
    def udp_session_expire_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "udp_session_expire_time", value)


