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
    'GetReferenceTablesResult',
    'AwaitableGetReferenceTablesResult',
    'get_reference_tables',
    'get_reference_tables_output',
]

@pulumi.output_type
class GetReferenceTablesResult:
    """
    A collection of values returned by getReferenceTables.
    """
    def __init__(__self__, enterprise_project_id=None, id=None, name=None, region=None, tables=None):
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
        if tables and not isinstance(tables, list):
            raise TypeError("Expected argument 'tables' to be a list")
        pulumi.set(__self__, "tables", tables)

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
    def name(self) -> Optional[str]:
        """
        The name of the reference table. The maximum length is 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tables(self) -> Sequence['outputs.GetReferenceTablesTableResult']:
        """
        A list of WAF reference tables.
        """
        return pulumi.get(self, "tables")


class AwaitableGetReferenceTablesResult(GetReferenceTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReferenceTablesResult(
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            region=self.region,
            tables=self.tables)


def get_reference_tables(enterprise_project_id: Optional[str] = None,
                         name: Optional[str] = None,
                         region: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReferenceTablesResult:
    """
    Use this data source to get a list of WAF reference tables.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    enterprise_project_id = config.require_object("enterpriseProjectId")
    reftables = huaweicloud.Waf.get_reference_tables(name="reference_table_name",
        enterprise_project_id=enterprise_project_id)
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of WAF reference tables.
    :param str name: The name of the reference table. The value is case sensitive and matches exactly.
    :param str region: The region in which to create the WAF reference table resource.
           If omitted, the provider-level region will be used.
    """
    __args__ = dict()
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Waf/getReferenceTables:getReferenceTables', __args__, opts=opts, typ=GetReferenceTablesResult).value

    return AwaitableGetReferenceTablesResult(
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region,
        tables=__ret__.tables)


@_utilities.lift_output_func(get_reference_tables)
def get_reference_tables_output(enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                region: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReferenceTablesResult]:
    """
    Use this data source to get a list of WAF reference tables.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    enterprise_project_id = config.require_object("enterpriseProjectId")
    reftables = huaweicloud.Waf.get_reference_tables(name="reference_table_name",
        enterprise_project_id=enterprise_project_id)
    ```


    :param str enterprise_project_id: Specifies the enterprise project ID of WAF reference tables.
    :param str name: The name of the reference table. The value is case sensitive and matches exactly.
    :param str region: The region in which to create the WAF reference table resource.
           If omitted, the provider-level region will be used.
    """
    ...