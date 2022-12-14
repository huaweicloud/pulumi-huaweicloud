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
    'GetVpcSubnetsResult',
    'AwaitableGetVpcSubnetsResult',
    'get_vpc_subnets',
    'get_vpc_subnets_output',
]

@pulumi.output_type
class GetVpcSubnetsResult:
    """
    A collection of values returned by getVpcSubnets.
    """
    def __init__(__self__, id=None, region=None, site_id=None, subnets=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if site_id and not isinstance(site_id, str):
            raise TypeError("Expected argument 'site_id' to be a str")
        pulumi.set(__self__, "site_id", site_id)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> Optional[str]:
        """
        Indicates the ID of the IEC site.
        """
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence['outputs.GetVpcSubnetsSubnetResult']:
        """
        A list of all the subnets found. The object is documented below.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcSubnetsResult(GetVpcSubnetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcSubnetsResult(
            id=self.id,
            region=self.region,
            site_id=self.site_id,
            subnets=self.subnets,
            vpc_id=self.vpc_id)


def get_vpc_subnets(region: Optional[str] = None,
                    site_id: Optional[str] = None,
                    vpc_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcSubnetsResult:
    """
    Use this data source to get a list of subnets belong to a specific IEC VPC.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    vpc_id = config.require_object("vpcId")
    all_subnets = huaweicloud.Iec.get_vpc_subnets(vpc_id=vpc_id)
    ```


    :param str region: The region in which to obtain the subnets. If omitted, the provider-level region will be
           used.
    :param str site_id: Specifies the ID of the IEC site.
    :param str vpc_id: Specifies the ID of the IEC VPC.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['siteId'] = site_id
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Iec/getVpcSubnets:getVpcSubnets', __args__, opts=opts, typ=GetVpcSubnetsResult).value

    return AwaitableGetVpcSubnetsResult(
        id=__ret__.id,
        region=__ret__.region,
        site_id=__ret__.site_id,
        subnets=__ret__.subnets,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_vpc_subnets)
def get_vpc_subnets_output(region: Optional[pulumi.Input[Optional[str]]] = None,
                           site_id: Optional[pulumi.Input[Optional[str]]] = None,
                           vpc_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcSubnetsResult]:
    """
    Use this data source to get a list of subnets belong to a specific IEC VPC.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    vpc_id = config.require_object("vpcId")
    all_subnets = huaweicloud.Iec.get_vpc_subnets(vpc_id=vpc_id)
    ```


    :param str region: The region in which to obtain the subnets. If omitted, the provider-level region will be
           used.
    :param str site_id: Specifies the ID of the IEC site.
    :param str vpc_id: Specifies the ID of the IEC VPC.
    """
    ...
