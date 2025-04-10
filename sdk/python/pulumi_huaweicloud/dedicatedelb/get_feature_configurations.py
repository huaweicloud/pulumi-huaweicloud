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
    'GetFeatureConfigurationsResult',
    'AwaitableGetFeatureConfigurationsResult',
    'get_feature_configurations',
    'get_feature_configurations_output',
]

@pulumi.output_type
class GetFeatureConfigurationsResult:
    """
    A collection of values returned by getFeatureConfigurations.
    """
    def __init__(__self__, configs=None, feature=None, id=None, region=None):
        if configs and not isinstance(configs, list):
            raise TypeError("Expected argument 'configs' to be a list")
        pulumi.set(__self__, "configs", configs)
        if feature and not isinstance(feature, str):
            raise TypeError("Expected argument 'feature' to be a str")
        pulumi.set(__self__, "feature", feature)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def configs(self) -> Sequence['outputs.GetFeatureConfigurationsConfigResult']:
        """
        Indicates the feature configuration list.
        """
        return pulumi.get(self, "configs")

    @property
    @pulumi.getter
    def feature(self) -> Optional[str]:
        """
        Indicates the feature name.
        """
        return pulumi.get(self, "feature")

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


class AwaitableGetFeatureConfigurationsResult(GetFeatureConfigurationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureConfigurationsResult(
            configs=self.configs,
            feature=self.feature,
            id=self.id,
            region=self.region)


def get_feature_configurations(feature: Optional[str] = None,
                               region: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureConfigurationsResult:
    """
    Use this data source to get the list of feature configurations of ELB of a tenant.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.DedicatedElb.get_feature_configurations()
    ```


    :param str feature: Specifies the feature name.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['feature'] = feature
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:DedicatedElb/getFeatureConfigurations:getFeatureConfigurations', __args__, opts=opts, typ=GetFeatureConfigurationsResult).value

    return AwaitableGetFeatureConfigurationsResult(
        configs=__ret__.configs,
        feature=__ret__.feature,
        id=__ret__.id,
        region=__ret__.region)


@_utilities.lift_output_func(get_feature_configurations)
def get_feature_configurations_output(feature: Optional[pulumi.Input[Optional[str]]] = None,
                                      region: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureConfigurationsResult]:
    """
    Use this data source to get the list of feature configurations of ELB of a tenant.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    test = huaweicloud.DedicatedElb.get_feature_configurations()
    ```


    :param str feature: Specifies the feature name.
    :param str region: Specifies the region in which to query the resource.
           If omitted, the provider-level region will be used.
    """
    ...
