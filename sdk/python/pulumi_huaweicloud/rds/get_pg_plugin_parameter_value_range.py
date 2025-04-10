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
    'GetPgPluginParameterValueRangeResult',
    'AwaitableGetPgPluginParameterValueRangeResult',
    'get_pg_plugin_parameter_value_range',
    'get_pg_plugin_parameter_value_range_output',
]

@pulumi.output_type
class GetPgPluginParameterValueRangeResult:
    """
    A collection of values returned by getPgPluginParameterValueRange.
    """
    def __init__(__self__, default_values=None, id=None, instance_id=None, name=None, region=None, restart_required=None, values=None):
        if default_values and not isinstance(default_values, list):
            raise TypeError("Expected argument 'default_values' to be a list")
        pulumi.set(__self__, "default_values", default_values)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if restart_required and not isinstance(restart_required, bool):
            raise TypeError("Expected argument 'restart_required' to be a bool")
        pulumi.set(__self__, "restart_required", restart_required)
        if values and not isinstance(values, list):
            raise TypeError("Expected argument 'values' to be a list")
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="defaultValues")
    def default_values(self) -> Sequence[str]:
        """
        Indicates the list of default parameter values. The default plugin will be loaded when the instance
        is created and can not be unloaded.
        """
        return pulumi.get(self, "default_values")

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
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="restartRequired")
    def restart_required(self) -> bool:
        """
        Indicates whether a reboot is required.
        """
        return pulumi.get(self, "restart_required")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Indicates the list of parameter values.
        """
        return pulumi.get(self, "values")


class AwaitableGetPgPluginParameterValueRangeResult(GetPgPluginParameterValueRangeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPgPluginParameterValueRangeResult(
            default_values=self.default_values,
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            region=self.region,
            restart_required=self.restart_required,
            values=self.values)


def get_pg_plugin_parameter_value_range(instance_id: Optional[str] = None,
                                        name: Optional[str] = None,
                                        region: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPgPluginParameterValueRangeResult:
    """
    Use this data source to get the list of RDS PostgreSQL plugin parameter value range.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_pg_plugin_parameter_value_range(instance_id=instance_id,
        name="shared_preload_libraries")
    ```


    :param str instance_id: Specifies the ID of the RDS instance.
    :param str name: Specifies the parameter name. Currently only **shared_preload_libraries** is supported.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getPgPluginParameterValueRange:getPgPluginParameterValueRange', __args__, opts=opts, typ=GetPgPluginParameterValueRangeResult).value

    return AwaitableGetPgPluginParameterValueRangeResult(
        default_values=__ret__.default_values,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        name=__ret__.name,
        region=__ret__.region,
        restart_required=__ret__.restart_required,
        values=__ret__.values)


@_utilities.lift_output_func(get_pg_plugin_parameter_value_range)
def get_pg_plugin_parameter_value_range_output(instance_id: Optional[pulumi.Input[str]] = None,
                                               name: Optional[pulumi.Input[str]] = None,
                                               region: Optional[pulumi.Input[Optional[str]]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPgPluginParameterValueRangeResult]:
    """
    Use this data source to get the list of RDS PostgreSQL plugin parameter value range.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_pg_plugin_parameter_value_range(instance_id=instance_id,
        name="shared_preload_libraries")
    ```


    :param str instance_id: Specifies the ID of the RDS instance.
    :param str name: Specifies the parameter name. Currently only **shared_preload_libraries** is supported.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    """
    ...
