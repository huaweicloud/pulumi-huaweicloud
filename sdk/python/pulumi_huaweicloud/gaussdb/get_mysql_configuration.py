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
    'GetMysqlConfigurationResult',
    'AwaitableGetMysqlConfigurationResult',
    'get_mysql_configuration',
    'get_mysql_configuration_output',
]

@pulumi.output_type
class GetMysqlConfigurationResult:
    """
    A collection of values returned by getMysqlConfiguration.
    """
    def __init__(__self__, datastore_name=None, datastore_version=None, description=None, id=None, name=None, region=None):
        if datastore_name and not isinstance(datastore_name, str):
            raise TypeError("Expected argument 'datastore_name' to be a str")
        pulumi.set(__self__, "datastore_name", datastore_name)
        if datastore_version and not isinstance(datastore_version, str):
            raise TypeError("Expected argument 'datastore_version' to be a str")
        pulumi.set(__self__, "datastore_version", datastore_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="datastoreName")
    def datastore_name(self) -> str:
        """
        Indicates the datastore name of the configuration.
        """
        return pulumi.get(self, "datastore_name")

    @property
    @pulumi.getter(name="datastoreVersion")
    def datastore_version(self) -> str:
        """
        Indicates the datastore version of the configuration.
        """
        return pulumi.get(self, "datastore_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Indicates the description of the configuration.
        """
        return pulumi.get(self, "description")

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


class AwaitableGetMysqlConfigurationResult(GetMysqlConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMysqlConfigurationResult(
            datastore_name=self.datastore_name,
            datastore_version=self.datastore_version,
            description=self.description,
            id=self.id,
            name=self.name,
            region=self.region)


def get_mysql_configuration(name: Optional[str] = None,
                            region: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMysqlConfigurationResult:
    """
    Use this data source to get available HuaweiCloud gaussdb mysql configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDB.get_mysql_configuration(name="Default-GaussDB-for-MySQL 8.0")
    ```


    :param str name: Specifies the name of the parameter template.
    :param str region: The region in which to obtain the configurations. If omitted, the provider-level region
           will be used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:GaussDB/getMysqlConfiguration:getMysqlConfiguration', __args__, opts=opts, typ=GetMysqlConfigurationResult).value

    return AwaitableGetMysqlConfigurationResult(
        datastore_name=__ret__.datastore_name,
        datastore_version=__ret__.datastore_version,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region)


@_utilities.lift_output_func(get_mysql_configuration)
def get_mysql_configuration_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMysqlConfigurationResult]:
    """
    Use this data source to get available HuaweiCloud gaussdb mysql configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    this = huaweicloud.GaussDB.get_mysql_configuration(name="Default-GaussDB-for-MySQL 8.0")
    ```


    :param str name: Specifies the name of the parameter template.
    :param str region: The region in which to obtain the configurations. If omitted, the provider-level region
           will be used.
    """
    ...
