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
    'GetSecgroupResult',
    'AwaitableGetSecgroupResult',
    'get_secgroup',
    'get_secgroup_output',
]

@pulumi.output_type
class GetSecgroupResult:
    """
    A collection of values returned by getSecgroup.
    """
    def __init__(__self__, created_at=None, description=None, enterprise_project_id=None, id=None, name=None, region=None, rules=None, secgroup_id=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if secgroup_id and not isinstance(secgroup_id, str):
            raise TypeError("Expected argument 'secgroup_id' to be a str")
        pulumi.set(__self__, "secgroup_id", secgroup_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The creation time, in UTC format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The supplementary information about the security group rule.
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
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetSecgroupRuleResult']:
        """
        The array of security group rules associating with the security group.
        The rule object is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="secgroupId")
    def secgroup_id(self) -> Optional[str]:
        return pulumi.get(self, "secgroup_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The last update time, in UTC format.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetSecgroupResult(GetSecgroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecgroupResult(
            created_at=self.created_at,
            description=self.description,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            region=self.region,
            rules=self.rules,
            secgroup_id=self.secgroup_id,
            updated_at=self.updated_at)


def get_secgroup(enterprise_project_id: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 secgroup_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecgroupResult:
    """
    Use this data source to get the ID of an available HuaweiCloud security group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    secgroup = huaweicloud.Vpc.get_secgroup(name="tf_test_secgroup")
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of the security group.
    :param str name: Specifies the name of the security group.
    :param str region: Specifies the region in which to obtain the security group. If omitted, the
           provider-level region will be used.
    :param str secgroup_id: Specifiest he ID of the security group.
    """
    __args__ = dict()
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['secgroupId'] = secgroup_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Vpc/getSecgroup:getSecgroup', __args__, opts=opts, typ=GetSecgroupResult).value

    return AwaitableGetSecgroupResult(
        created_at=__ret__.created_at,
        description=__ret__.description,
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region,
        rules=__ret__.rules,
        secgroup_id=__ret__.secgroup_id,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_secgroup)
def get_secgroup_output(enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        secgroup_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecgroupResult]:
    """
    Use this data source to get the ID of an available HuaweiCloud security group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    secgroup = huaweicloud.Vpc.get_secgroup(name="tf_test_secgroup")
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of the security group.
    :param str name: Specifies the name of the security group.
    :param str region: Specifies the region in which to obtain the security group. If omitted, the
           provider-level region will be used.
    :param str secgroup_id: Specifiest he ID of the security group.
    """
    ...