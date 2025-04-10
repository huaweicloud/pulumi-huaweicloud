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
    'GetL7rulesResult',
    'AwaitableGetL7rulesResult',
    'get_l7rules',
    'get_l7rules_output',
]

@pulumi.output_type
class GetL7rulesResult:
    """
    A collection of values returned by getL7rules.
    """
    def __init__(__self__, compare_type=None, enterprise_project_id=None, id=None, l7policy_id=None, l7rule_id=None, l7rules=None, region=None, type=None, value=None):
        if compare_type and not isinstance(compare_type, str):
            raise TypeError("Expected argument 'compare_type' to be a str")
        pulumi.set(__self__, "compare_type", compare_type)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if l7policy_id and not isinstance(l7policy_id, str):
            raise TypeError("Expected argument 'l7policy_id' to be a str")
        pulumi.set(__self__, "l7policy_id", l7policy_id)
        if l7rule_id and not isinstance(l7rule_id, str):
            raise TypeError("Expected argument 'l7rule_id' to be a str")
        pulumi.set(__self__, "l7rule_id", l7rule_id)
        if l7rules and not isinstance(l7rules, list):
            raise TypeError("Expected argument 'l7rules' to be a list")
        pulumi.set(__self__, "l7rules", l7rules)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="compareType")
    def compare_type(self) -> Optional[str]:
        """
        How the requests are matched with the domain name or URL.
        """
        return pulumi.get(self, "compare_type")

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
    @pulumi.getter(name="l7policyId")
    def l7policy_id(self) -> str:
        return pulumi.get(self, "l7policy_id")

    @property
    @pulumi.getter(name="l7ruleId")
    def l7rule_id(self) -> Optional[str]:
        return pulumi.get(self, "l7rule_id")

    @property
    @pulumi.getter
    def l7rules(self) -> Sequence['outputs.GetL7rulesL7ruleResult']:
        """
        Lists the L7 rules.
        The l7rules structure is documented below.
        """
        return pulumi.get(self, "l7rules")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the forwarding rule.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of the match item.
        """
        return pulumi.get(self, "value")


class AwaitableGetL7rulesResult(GetL7rulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetL7rulesResult(
            compare_type=self.compare_type,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            l7policy_id=self.l7policy_id,
            l7rule_id=self.l7rule_id,
            l7rules=self.l7rules,
            region=self.region,
            type=self.type,
            value=self.value)


def get_l7rules(compare_type: Optional[str] = None,
                enterprise_project_id: Optional[str] = None,
                l7policy_id: Optional[str] = None,
                l7rule_id: Optional[str] = None,
                region: Optional[str] = None,
                type: Optional[str] = None,
                value: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetL7rulesResult:
    """
    Use this data source to get the list of ELB L7 rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    compare_type = config.require_object("compareType")
    test = huaweicloud.DedicatedElb.get_l7rules(compare_type=compare_type)
    ```


    :param str compare_type: Specifies how requests are matched with the domain names or URL. Values options:
           **EQUAL_TO**, **REGEX**, **STARTS_WITH**.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str l7policy_id: Specifies the forwarding policy ID.
    :param str l7rule_id: Specifies the forwarding rule ID.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param str type: Specifies the match type. Value options: **HOST_NAME**, **PATH**, **METHOD**, **HEADER**,
           **QUERY_STRING**, **SOURCE_IP**, **COOKIE**.
    :param str value: Specifies the value of the match content.
    """
    __args__ = dict()
    __args__['compareType'] = compare_type
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['l7policyId'] = l7policy_id
    __args__['l7ruleId'] = l7rule_id
    __args__['region'] = region
    __args__['type'] = type
    __args__['value'] = value
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:DedicatedElb/getL7rules:getL7rules', __args__, opts=opts, typ=GetL7rulesResult).value

    return AwaitableGetL7rulesResult(
        compare_type=__ret__.compare_type,
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        l7policy_id=__ret__.l7policy_id,
        l7rule_id=__ret__.l7rule_id,
        l7rules=__ret__.l7rules,
        region=__ret__.region,
        type=__ret__.type,
        value=__ret__.value)


@_utilities.lift_output_func(get_l7rules)
def get_l7rules_output(compare_type: Optional[pulumi.Input[Optional[str]]] = None,
                       enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       l7policy_id: Optional[pulumi.Input[str]] = None,
                       l7rule_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       type: Optional[pulumi.Input[Optional[str]]] = None,
                       value: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetL7rulesResult]:
    """
    Use this data source to get the list of ELB L7 rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    compare_type = config.require_object("compareType")
    test = huaweicloud.DedicatedElb.get_l7rules(compare_type=compare_type)
    ```


    :param str compare_type: Specifies how requests are matched with the domain names or URL. Values options:
           **EQUAL_TO**, **REGEX**, **STARTS_WITH**.
    :param str enterprise_project_id: Specifies the enterprise project ID.
    :param str l7policy_id: Specifies the forwarding policy ID.
    :param str l7rule_id: Specifies the forwarding rule ID.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param str type: Specifies the match type. Value options: **HOST_NAME**, **PATH**, **METHOD**, **HEADER**,
           **QUERY_STRING**, **SOURCE_IP**, **COOKIE**.
    :param str value: Specifies the value of the match content.
    """
    ...
