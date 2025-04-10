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
    'GetMysqlProxyFlavorsResult',
    'AwaitableGetMysqlProxyFlavorsResult',
    'get_mysql_proxy_flavors',
    'get_mysql_proxy_flavors_output',
]

@pulumi.output_type
class GetMysqlProxyFlavorsResult:
    """
    A collection of values returned by getMysqlProxyFlavors.
    """
    def __init__(__self__, flavor_groups=None, id=None, instance_id=None, region=None):
        if flavor_groups and not isinstance(flavor_groups, list):
            raise TypeError("Expected argument 'flavor_groups' to be a list")
        pulumi.set(__self__, "flavor_groups", flavor_groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="flavorGroups")
    def flavor_groups(self) -> Sequence['outputs.GetMysqlProxyFlavorsFlavorGroupResult']:
        """
        Indicates the list of flavor groups.
        The flavor_groups structure is documented below.
        """
        return pulumi.get(self, "flavor_groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetMysqlProxyFlavorsResult(GetMysqlProxyFlavorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMysqlProxyFlavorsResult(
            flavor_groups=self.flavor_groups,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region)


def get_mysql_proxy_flavors(instance_id: Optional[str] = None,
                            region: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMysqlProxyFlavorsResult:
    """
    Use this data source to get the list of RDS MySQL proxy flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    flavor = huaweicloud.Rds.get_mysql_proxy_flavors(instance_id=instance_id)
    ```


    :param str instance_id: Specifies the ID of RDS MySQL instance.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getMysqlProxyFlavors:getMysqlProxyFlavors', __args__, opts=opts, typ=GetMysqlProxyFlavorsResult).value

    return AwaitableGetMysqlProxyFlavorsResult(
        flavor_groups=__ret__.flavor_groups,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        region=__ret__.region)


@_utilities.lift_output_func(get_mysql_proxy_flavors)
def get_mysql_proxy_flavors_output(instance_id: Optional[pulumi.Input[str]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMysqlProxyFlavorsResult]:
    """
    Use this data source to get the list of RDS MySQL proxy flavors.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    flavor = huaweicloud.Rds.get_mysql_proxy_flavors(instance_id=instance_id)
    ```


    :param str instance_id: Specifies the ID of RDS MySQL instance.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    ...
