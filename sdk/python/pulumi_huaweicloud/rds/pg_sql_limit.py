# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PgSqlLimitArgs', 'PgSqlLimit']

@pulumi.input_type
class PgSqlLimitArgs:
    def __init__(__self__, *,
                 db_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 max_concurrency: pulumi.Input[int],
                 max_waiting: pulumi.Input[int],
                 query_id: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 search_path: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PgSqlLimit resource.
        :param pulumi.Input[str] db_name: Specifies the name of the database.
        :param pulumi.Input[str] instance_id: Specifies the ID of RDS PostgreSQL instance.
        :param pulumi.Input[int] max_concurrency: Specifies the number of SQL statements executed simultaneously.
               Value ranges from `0` to `50000`. `0` means no limit.
        :param pulumi.Input[int] max_waiting: Specifies the max waiting time in seconds.
        :param pulumi.Input[str] query_id: Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        :param pulumi.Input[str] query_string: Specifies the text form of SQL statement.
        :param pulumi.Input[str] region: The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] search_path: Specifies the query order for names that are not schema qualified.
               Defaults to **public**,
        :param pulumi.Input[str] switch: Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "max_concurrency", max_concurrency)
        pulumi.set(__self__, "max_waiting", max_waiting)
        if query_id is not None:
            pulumi.set(__self__, "query_id", query_id)
        if query_string is not None:
            pulumi.set(__self__, "query_string", query_string)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if search_path is not None:
            pulumi.set(__self__, "search_path", search_path)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the database.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of RDS PostgreSQL instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> pulumi.Input[int]:
        """
        Specifies the number of SQL statements executed simultaneously.
        Value ranges from `0` to `50000`. `0` means no limit.
        """
        return pulumi.get(self, "max_concurrency")

    @max_concurrency.setter
    def max_concurrency(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_concurrency", value)

    @property
    @pulumi.getter(name="maxWaiting")
    def max_waiting(self) -> pulumi.Input[int]:
        """
        Specifies the max waiting time in seconds.
        """
        return pulumi.get(self, "max_waiting")

    @max_waiting.setter
    def max_waiting(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_waiting", value)

    @property
    @pulumi.getter(name="queryId")
    def query_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        """
        return pulumi.get(self, "query_id")

    @query_id.setter
    def query_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_id", value)

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the text form of SQL statement.
        """
        return pulumi.get(self, "query_string")

    @query_string.setter
    def query_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_string", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="searchPath")
    def search_path(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the query order for names that are not schema qualified.
        Defaults to **public**,
        """
        return pulumi.get(self, "search_path")

    @search_path.setter
    def search_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_path", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class _PgSqlLimitState:
    def __init__(__self__, *,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_effective: Optional[pulumi.Input[bool]] = None,
                 max_concurrency: Optional[pulumi.Input[int]] = None,
                 max_waiting: Optional[pulumi.Input[int]] = None,
                 query_id: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 search_path: Optional[pulumi.Input[str]] = None,
                 sql_limit_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PgSqlLimit resources.
        :param pulumi.Input[str] db_name: Specifies the name of the database.
        :param pulumi.Input[str] instance_id: Specifies the ID of RDS PostgreSQL instance.
        :param pulumi.Input[bool] is_effective: Indicates whether the SQL limit is effective.
        :param pulumi.Input[int] max_concurrency: Specifies the number of SQL statements executed simultaneously.
               Value ranges from `0` to `50000`. `0` means no limit.
        :param pulumi.Input[int] max_waiting: Specifies the max waiting time in seconds.
        :param pulumi.Input[str] query_id: Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        :param pulumi.Input[str] query_string: Specifies the text form of SQL statement.
        :param pulumi.Input[str] region: The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] search_path: Specifies the query order for names that are not schema qualified.
               Defaults to **public**,
        :param pulumi.Input[str] sql_limit_id: Indicates the ID of SQL limit.
        :param pulumi.Input[str] switch: Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_effective is not None:
            pulumi.set(__self__, "is_effective", is_effective)
        if max_concurrency is not None:
            pulumi.set(__self__, "max_concurrency", max_concurrency)
        if max_waiting is not None:
            pulumi.set(__self__, "max_waiting", max_waiting)
        if query_id is not None:
            pulumi.set(__self__, "query_id", query_id)
        if query_string is not None:
            pulumi.set(__self__, "query_string", query_string)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if search_path is not None:
            pulumi.set(__self__, "search_path", search_path)
        if sql_limit_id is not None:
            pulumi.set(__self__, "sql_limit_id", sql_limit_id)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the database.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of RDS PostgreSQL instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isEffective")
    def is_effective(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the SQL limit is effective.
        """
        return pulumi.get(self, "is_effective")

    @is_effective.setter
    def is_effective(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_effective", value)

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of SQL statements executed simultaneously.
        Value ranges from `0` to `50000`. `0` means no limit.
        """
        return pulumi.get(self, "max_concurrency")

    @max_concurrency.setter
    def max_concurrency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrency", value)

    @property
    @pulumi.getter(name="maxWaiting")
    def max_waiting(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the max waiting time in seconds.
        """
        return pulumi.get(self, "max_waiting")

    @max_waiting.setter
    def max_waiting(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_waiting", value)

    @property
    @pulumi.getter(name="queryId")
    def query_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        """
        return pulumi.get(self, "query_id")

    @query_id.setter
    def query_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_id", value)

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the text form of SQL statement.
        """
        return pulumi.get(self, "query_string")

    @query_string.setter
    def query_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_string", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="searchPath")
    def search_path(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the query order for names that are not schema qualified.
        Defaults to **public**,
        """
        return pulumi.get(self, "search_path")

    @search_path.setter
    def search_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_path", value)

    @property
    @pulumi.getter(name="sqlLimitId")
    def sql_limit_id(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the ID of SQL limit.
        """
        return pulumi.get(self, "sql_limit_id")

    @sql_limit_id.setter
    def sql_limit_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_limit_id", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


class PgSqlLimit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[int]] = None,
                 max_waiting: Optional[pulumi.Input[int]] = None,
                 query_id: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 search_path: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage an RDS PostgreSQL SQL limit resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        db_name = config.require_object("dbName")
        instance = huaweicloud.rds.PgSqlLimit("instance",
            instance_id=instance_id,
            db_name=db_name,
            query_id="5",
            max_concurrency=20,
            max_waiting=5)
        ```

        ## Import

        The SQL limit can be imported using the `instance_id`, `db_name` and `sql_limit_id`, separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Rds/pgSqlLimit:PgSqlLimit test <instance_id>/<db_name>/<sql_limit_id>
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`query_id`, `query_string`. It is generally recommended running `terraform plan` after importing an RDS PostgreSQL SQL limit. You can then decide if changes should be applied to the RDS PostgreSQL SQL limit, or the resource definition should be updated to align with the RDS PostgreSQL SQL limit. Also, you can ignore changes as below. hcl resource "huaweicloud_rds_pg_sql_limit" "test" {

         ...

         lifecycle {

         ignore_changes = [

         "query_id", "query_string",

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_name: Specifies the name of the database.
        :param pulumi.Input[str] instance_id: Specifies the ID of RDS PostgreSQL instance.
        :param pulumi.Input[int] max_concurrency: Specifies the number of SQL statements executed simultaneously.
               Value ranges from `0` to `50000`. `0` means no limit.
        :param pulumi.Input[int] max_waiting: Specifies the max waiting time in seconds.
        :param pulumi.Input[str] query_id: Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        :param pulumi.Input[str] query_string: Specifies the text form of SQL statement.
        :param pulumi.Input[str] region: The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] search_path: Specifies the query order for names that are not schema qualified.
               Defaults to **public**,
        :param pulumi.Input[str] switch: Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PgSqlLimitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage an RDS PostgreSQL SQL limit resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        db_name = config.require_object("dbName")
        instance = huaweicloud.rds.PgSqlLimit("instance",
            instance_id=instance_id,
            db_name=db_name,
            query_id="5",
            max_concurrency=20,
            max_waiting=5)
        ```

        ## Import

        The SQL limit can be imported using the `instance_id`, `db_name` and `sql_limit_id`, separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Rds/pgSqlLimit:PgSqlLimit test <instance_id>/<db_name>/<sql_limit_id>
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`query_id`, `query_string`. It is generally recommended running `terraform plan` after importing an RDS PostgreSQL SQL limit. You can then decide if changes should be applied to the RDS PostgreSQL SQL limit, or the resource definition should be updated to align with the RDS PostgreSQL SQL limit. Also, you can ignore changes as below. hcl resource "huaweicloud_rds_pg_sql_limit" "test" {

         ...

         lifecycle {

         ignore_changes = [

         "query_id", "query_string",

         ]

         } }

        :param str resource_name: The name of the resource.
        :param PgSqlLimitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PgSqlLimitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[int]] = None,
                 max_waiting: Optional[pulumi.Input[int]] = None,
                 query_id: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 search_path: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PgSqlLimitArgs.__new__(PgSqlLimitArgs)

            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if max_concurrency is None and not opts.urn:
                raise TypeError("Missing required property 'max_concurrency'")
            __props__.__dict__["max_concurrency"] = max_concurrency
            if max_waiting is None and not opts.urn:
                raise TypeError("Missing required property 'max_waiting'")
            __props__.__dict__["max_waiting"] = max_waiting
            __props__.__dict__["query_id"] = query_id
            __props__.__dict__["query_string"] = query_string
            __props__.__dict__["region"] = region
            __props__.__dict__["search_path"] = search_path
            __props__.__dict__["switch"] = switch
            __props__.__dict__["is_effective"] = None
            __props__.__dict__["sql_limit_id"] = None
        super(PgSqlLimit, __self__).__init__(
            'huaweicloud:Rds/pgSqlLimit:PgSqlLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_effective: Optional[pulumi.Input[bool]] = None,
            max_concurrency: Optional[pulumi.Input[int]] = None,
            max_waiting: Optional[pulumi.Input[int]] = None,
            query_id: Optional[pulumi.Input[str]] = None,
            query_string: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            search_path: Optional[pulumi.Input[str]] = None,
            sql_limit_id: Optional[pulumi.Input[str]] = None,
            switch: Optional[pulumi.Input[str]] = None) -> 'PgSqlLimit':
        """
        Get an existing PgSqlLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_name: Specifies the name of the database.
        :param pulumi.Input[str] instance_id: Specifies the ID of RDS PostgreSQL instance.
        :param pulumi.Input[bool] is_effective: Indicates whether the SQL limit is effective.
        :param pulumi.Input[int] max_concurrency: Specifies the number of SQL statements executed simultaneously.
               Value ranges from `0` to `50000`. `0` means no limit.
        :param pulumi.Input[int] max_waiting: Specifies the max waiting time in seconds.
        :param pulumi.Input[str] query_id: Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        :param pulumi.Input[str] query_string: Specifies the text form of SQL statement.
        :param pulumi.Input[str] region: The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] search_path: Specifies the query order for names that are not schema qualified.
               Defaults to **public**,
        :param pulumi.Input[str] sql_limit_id: Indicates the ID of SQL limit.
        :param pulumi.Input[str] switch: Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PgSqlLimitState.__new__(_PgSqlLimitState)

        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_effective"] = is_effective
        __props__.__dict__["max_concurrency"] = max_concurrency
        __props__.__dict__["max_waiting"] = max_waiting
        __props__.__dict__["query_id"] = query_id
        __props__.__dict__["query_string"] = query_string
        __props__.__dict__["region"] = region
        __props__.__dict__["search_path"] = search_path
        __props__.__dict__["sql_limit_id"] = sql_limit_id
        __props__.__dict__["switch"] = switch
        return PgSqlLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the database.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of RDS PostgreSQL instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isEffective")
    def is_effective(self) -> pulumi.Output[bool]:
        """
        Indicates whether the SQL limit is effective.
        """
        return pulumi.get(self, "is_effective")

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> pulumi.Output[int]:
        """
        Specifies the number of SQL statements executed simultaneously.
        Value ranges from `0` to `50000`. `0` means no limit.
        """
        return pulumi.get(self, "max_concurrency")

    @property
    @pulumi.getter(name="maxWaiting")
    def max_waiting(self) -> pulumi.Output[int]:
        """
        Specifies the max waiting time in seconds.
        """
        return pulumi.get(self, "max_waiting")

    @property
    @pulumi.getter(name="queryId")
    def query_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the query ID. Value ranges: **-9223372036854775808~9223372036854775807**.
        """
        return pulumi.get(self, "query_id")

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the text form of SQL statement.
        """
        return pulumi.get(self, "query_string")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the rds PostgreSQL SQL limit resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="searchPath")
    def search_path(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the query order for names that are not schema qualified.
        Defaults to **public**,
        """
        return pulumi.get(self, "search_path")

    @property
    @pulumi.getter(name="sqlLimitId")
    def sql_limit_id(self) -> pulumi.Output[str]:
        """
        Indicates the ID of SQL limit.
        """
        return pulumi.get(self, "sql_limit_id")

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Output[str]:
        """
        Specifies the SQL limit switch. Value options: **open**, **close**.
        """
        return pulumi.get(self, "switch")

