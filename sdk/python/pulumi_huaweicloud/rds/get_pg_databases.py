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
    'GetPgDatabasesResult',
    'AwaitableGetPgDatabasesResult',
    'get_pg_databases',
    'get_pg_databases_output',
]

@pulumi.output_type
class GetPgDatabasesResult:
    """
    A collection of values returned by getPgDatabases.
    """
    def __init__(__self__, character_set=None, databases=None, id=None, instance_id=None, lc_collate=None, name=None, owner=None, region=None, size=None):
        if character_set and not isinstance(character_set, str):
            raise TypeError("Expected argument 'character_set' to be a str")
        pulumi.set(__self__, "character_set", character_set)
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lc_collate and not isinstance(lc_collate, str):
            raise TypeError("Expected argument 'lc_collate' to be a str")
        pulumi.set(__self__, "lc_collate", lc_collate)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter(name="characterSet")
    def character_set(self) -> Optional[str]:
        """
        Indicates the character set used by the database.
        """
        return pulumi.get(self, "character_set")

    @property
    @pulumi.getter
    def databases(self) -> Sequence['outputs.GetPgDatabasesDatabaseResult']:
        """
        Indicates the database list.
        The databases structure is documented below.
        """
        return pulumi.get(self, "databases")

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
    @pulumi.getter(name="lcCollate")
    def lc_collate(self) -> Optional[str]:
        """
        Indicates the database collation.
        """
        return pulumi.get(self, "lc_collate")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Indicates the database name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        """
        Indicates the database owner.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        Indicates the database size, in bytes.
        """
        return pulumi.get(self, "size")


class AwaitableGetPgDatabasesResult(GetPgDatabasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPgDatabasesResult(
            character_set=self.character_set,
            databases=self.databases,
            id=self.id,
            instance_id=self.instance_id,
            lc_collate=self.lc_collate,
            name=self.name,
            owner=self.owner,
            region=self.region,
            size=self.size)


def get_pg_databases(character_set: Optional[str] = None,
                     instance_id: Optional[str] = None,
                     lc_collate: Optional[str] = None,
                     name: Optional[str] = None,
                     owner: Optional[str] = None,
                     region: Optional[str] = None,
                     size: Optional[int] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPgDatabasesResult:
    """
    Use this data source to get the list of RDS PostgreSQL databases.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_pg_databases(instance_id=instance_id)
    ```


    :param str character_set: Specifies the character set used by the database.
           For details, see [documentation](https://www.postgresql.org/docs/16/infoschema-character-sets.html).
    :param str instance_id: Specifies the PostgreSQL instance ID.
    :param str lc_collate: Specifies the database collation.
           For details, see [documentation](https://support.huaweicloud.com/intl/en-us/bestpractice-rds/rds_pg_0017.html).
    :param str name: Specifies the database name.
    :param str owner: Specifies the database owner.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param int size: Specifies the database size, in bytes.
    """
    __args__ = dict()
    __args__['characterSet'] = character_set
    __args__['instanceId'] = instance_id
    __args__['lcCollate'] = lc_collate
    __args__['name'] = name
    __args__['owner'] = owner
    __args__['region'] = region
    __args__['size'] = size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Rds/getPgDatabases:getPgDatabases', __args__, opts=opts, typ=GetPgDatabasesResult).value

    return AwaitableGetPgDatabasesResult(
        character_set=__ret__.character_set,
        databases=__ret__.databases,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        lc_collate=__ret__.lc_collate,
        name=__ret__.name,
        owner=__ret__.owner,
        region=__ret__.region,
        size=__ret__.size)


@_utilities.lift_output_func(get_pg_databases)
def get_pg_databases_output(character_set: Optional[pulumi.Input[Optional[str]]] = None,
                            instance_id: Optional[pulumi.Input[str]] = None,
                            lc_collate: Optional[pulumi.Input[Optional[str]]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            owner: Optional[pulumi.Input[Optional[str]]] = None,
                            region: Optional[pulumi.Input[Optional[str]]] = None,
                            size: Optional[pulumi.Input[Optional[int]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPgDatabasesResult]:
    """
    Use this data source to get the list of RDS PostgreSQL databases.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    instance_id = config.require_object("instanceId")
    test = huaweicloud.Rds.get_pg_databases(instance_id=instance_id)
    ```


    :param str character_set: Specifies the character set used by the database.
           For details, see [documentation](https://www.postgresql.org/docs/16/infoschema-character-sets.html).
    :param str instance_id: Specifies the PostgreSQL instance ID.
    :param str lc_collate: Specifies the database collation.
           For details, see [documentation](https://support.huaweicloud.com/intl/en-us/bestpractice-rds/rds_pg_0017.html).
    :param str name: Specifies the database name.
    :param str owner: Specifies the database owner.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param int size: Specifies the database size, in bytes.
    """
    ...
