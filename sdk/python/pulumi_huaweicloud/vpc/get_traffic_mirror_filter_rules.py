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
    'GetTrafficMirrorFilterRulesResult',
    'AwaitableGetTrafficMirrorFilterRulesResult',
    'get_traffic_mirror_filter_rules',
    'get_traffic_mirror_filter_rules_output',
]

@pulumi.output_type
class GetTrafficMirrorFilterRulesResult:
    """
    A collection of values returned by getTrafficMirrorFilterRules.
    """
    def __init__(__self__, action=None, destination_cidr_block=None, destination_port_range=None, direction=None, id=None, priority=None, protocol=None, region=None, source_cidr_block=None, source_port_range=None, traffic_mirror_filter_id=None, traffic_mirror_filter_rule_id=None, traffic_mirror_filter_rules=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if destination_cidr_block and not isinstance(destination_cidr_block, str):
            raise TypeError("Expected argument 'destination_cidr_block' to be a str")
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if destination_port_range and not isinstance(destination_port_range, str):
            raise TypeError("Expected argument 'destination_port_range' to be a str")
        pulumi.set(__self__, "destination_port_range", destination_port_range)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if source_cidr_block and not isinstance(source_cidr_block, str):
            raise TypeError("Expected argument 'source_cidr_block' to be a str")
        pulumi.set(__self__, "source_cidr_block", source_cidr_block)
        if source_port_range and not isinstance(source_port_range, str):
            raise TypeError("Expected argument 'source_port_range' to be a str")
        pulumi.set(__self__, "source_port_range", source_port_range)
        if traffic_mirror_filter_id and not isinstance(traffic_mirror_filter_id, str):
            raise TypeError("Expected argument 'traffic_mirror_filter_id' to be a str")
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if traffic_mirror_filter_rule_id and not isinstance(traffic_mirror_filter_rule_id, str):
            raise TypeError("Expected argument 'traffic_mirror_filter_rule_id' to be a str")
        pulumi.set(__self__, "traffic_mirror_filter_rule_id", traffic_mirror_filter_rule_id)
        if traffic_mirror_filter_rules and not isinstance(traffic_mirror_filter_rules, list):
            raise TypeError("Expected argument 'traffic_mirror_filter_rules' to be a list")
        pulumi.set(__self__, "traffic_mirror_filter_rules", traffic_mirror_filter_rules)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Whether to accept or reject traffic.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[str]:
        """
        Destination CIDR block of the mirrored traffic.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> Optional[str]:
        """
        Source port range.
        """
        return pulumi.get(self, "destination_port_range")

    @property
    @pulumi.getter
    def direction(self) -> Optional[str]:
        """
        Traffic direction.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def priority(self) -> Optional[str]:
        """
        Mirror filter rule priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        Protocol of the mirrored traffic.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> Optional[str]:
        """
        Source CIDR block of the mirrored traffic.
        """
        return pulumi.get(self, "source_cidr_block")

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> Optional[str]:
        """
        Source port range.
        """
        return pulumi.get(self, "source_port_range")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> Optional[str]:
        """
        Traffic mirror filter ID.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorFilterRuleId")
    def traffic_mirror_filter_rule_id(self) -> Optional[str]:
        return pulumi.get(self, "traffic_mirror_filter_rule_id")

    @property
    @pulumi.getter(name="trafficMirrorFilterRules")
    def traffic_mirror_filter_rules(self) -> Sequence['outputs.GetTrafficMirrorFilterRulesTrafficMirrorFilterRuleResult']:
        """
        List of traffic mirror filter rules.
        """
        return pulumi.get(self, "traffic_mirror_filter_rules")


class AwaitableGetTrafficMirrorFilterRulesResult(GetTrafficMirrorFilterRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficMirrorFilterRulesResult(
            action=self.action,
            destination_cidr_block=self.destination_cidr_block,
            destination_port_range=self.destination_port_range,
            direction=self.direction,
            id=self.id,
            priority=self.priority,
            protocol=self.protocol,
            region=self.region,
            source_cidr_block=self.source_cidr_block,
            source_port_range=self.source_port_range,
            traffic_mirror_filter_id=self.traffic_mirror_filter_id,
            traffic_mirror_filter_rule_id=self.traffic_mirror_filter_rule_id,
            traffic_mirror_filter_rules=self.traffic_mirror_filter_rules)


def get_traffic_mirror_filter_rules(action: Optional[str] = None,
                                    destination_cidr_block: Optional[str] = None,
                                    destination_port_range: Optional[str] = None,
                                    direction: Optional[str] = None,
                                    priority: Optional[str] = None,
                                    protocol: Optional[str] = None,
                                    region: Optional[str] = None,
                                    source_cidr_block: Optional[str] = None,
                                    source_port_range: Optional[str] = None,
                                    traffic_mirror_filter_id: Optional[str] = None,
                                    traffic_mirror_filter_rule_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficMirrorFilterRulesResult:
    """
    Use this data source to get the traffic mirror filter rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test1 = huaweicloud.Vpc.get_traffic_mirror_filter_rules(protocol="udp")
    ```


    :param str action: The policy of in the traffic mirror filter rule.
           Valid values are **accept** or **reject**.
    :param str destination_cidr_block: The destination IP address of the traffic mirror filter rule.
    :param str destination_port_range: The destination port number range of the traffic mirror filter rule.
           The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
    :param str direction: The direction of the traffic mirror filter rule.
           Valid values are **ingress** or **egress**.
    :param str priority: The priority number of the traffic mirror filter rule.
           Valid value ranges from `1` to `65,535`.
    :param str protocol: The protocol of the traffic mirror filter rule.
           Valid value are **tcp**, **udp**, **icmp**, **icmpv6**, **all**.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str source_cidr_block: The source IP address of the traffic mirror filter rule.
    :param str source_port_range: The source port number range of the traffic mirror filter rule.
           The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
    :param str traffic_mirror_filter_id: The traffic mirror filter ID used as the query filter.
    :param str traffic_mirror_filter_rule_id: The traffic mirror filter rule ID used as the query filter.
    """
    __args__ = dict()
    __args__['action'] = action
    __args__['destinationCidrBlock'] = destination_cidr_block
    __args__['destinationPortRange'] = destination_port_range
    __args__['direction'] = direction
    __args__['priority'] = priority
    __args__['protocol'] = protocol
    __args__['region'] = region
    __args__['sourceCidrBlock'] = source_cidr_block
    __args__['sourcePortRange'] = source_port_range
    __args__['trafficMirrorFilterId'] = traffic_mirror_filter_id
    __args__['trafficMirrorFilterRuleId'] = traffic_mirror_filter_rule_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Vpc/getTrafficMirrorFilterRules:getTrafficMirrorFilterRules', __args__, opts=opts, typ=GetTrafficMirrorFilterRulesResult).value

    return AwaitableGetTrafficMirrorFilterRulesResult(
        action=__ret__.action,
        destination_cidr_block=__ret__.destination_cidr_block,
        destination_port_range=__ret__.destination_port_range,
        direction=__ret__.direction,
        id=__ret__.id,
        priority=__ret__.priority,
        protocol=__ret__.protocol,
        region=__ret__.region,
        source_cidr_block=__ret__.source_cidr_block,
        source_port_range=__ret__.source_port_range,
        traffic_mirror_filter_id=__ret__.traffic_mirror_filter_id,
        traffic_mirror_filter_rule_id=__ret__.traffic_mirror_filter_rule_id,
        traffic_mirror_filter_rules=__ret__.traffic_mirror_filter_rules)


@_utilities.lift_output_func(get_traffic_mirror_filter_rules)
def get_traffic_mirror_filter_rules_output(action: Optional[pulumi.Input[Optional[str]]] = None,
                                           destination_cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                                           destination_port_range: Optional[pulumi.Input[Optional[str]]] = None,
                                           direction: Optional[pulumi.Input[Optional[str]]] = None,
                                           priority: Optional[pulumi.Input[Optional[str]]] = None,
                                           protocol: Optional[pulumi.Input[Optional[str]]] = None,
                                           region: Optional[pulumi.Input[Optional[str]]] = None,
                                           source_cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                                           source_port_range: Optional[pulumi.Input[Optional[str]]] = None,
                                           traffic_mirror_filter_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           traffic_mirror_filter_rule_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrafficMirrorFilterRulesResult]:
    """
    Use this data source to get the traffic mirror filter rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test1 = huaweicloud.Vpc.get_traffic_mirror_filter_rules(protocol="udp")
    ```


    :param str action: The policy of in the traffic mirror filter rule.
           Valid values are **accept** or **reject**.
    :param str destination_cidr_block: The destination IP address of the traffic mirror filter rule.
    :param str destination_port_range: The destination port number range of the traffic mirror filter rule.
           The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
    :param str direction: The direction of the traffic mirror filter rule.
           Valid values are **ingress** or **egress**.
    :param str priority: The priority number of the traffic mirror filter rule.
           Valid value ranges from `1` to `65,535`.
    :param str protocol: The protocol of the traffic mirror filter rule.
           Valid value are **tcp**, **udp**, **icmp**, **icmpv6**, **all**.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    :param str source_cidr_block: The source IP address of the traffic mirror filter rule.
    :param str source_port_range: The source port number range of the traffic mirror filter rule.
           The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
    :param str traffic_mirror_filter_id: The traffic mirror filter ID used as the query filter.
    :param str traffic_mirror_filter_rule_id: The traffic mirror filter rule ID used as the query filter.
    """
    ...
