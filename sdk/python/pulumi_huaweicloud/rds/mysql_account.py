# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MysqlAccountArgs', 'MysqlAccount']

@pulumi.input_type
class MysqlAccountArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MysqlAccount resource.
        :param pulumi.Input[str] instance_id: Specifies the rds instance id. Changing this will create a new resource.
        :param pulumi.Input[str] password: Specifies the password of the db account. The parameter must be 8 to 32 characters
               long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
               different from `name` or `name` spelled backwards.
        :param pulumi.Input[str] description: Specifies remarks of the database account. The parameter must be 1 to 512
               characters long and is supported only for MySQL 8.0.25 and later versions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: Specifies the IP addresses that are allowed to access your DB instance.
               + If the IP address is set to %, all IP addresses are allowed to access your instance.
               + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
               your instance.
               + Multiple IP addresses can be added.
        :param pulumi.Input[str] name: Specifies the username of the db account. Only lowercase letters, digits,
               hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
               + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
               + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        :param pulumi.Input[str] region: The region in which to create the rds account resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "password", password)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the rds instance id. Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Specifies the password of the db account. The parameter must be 8 to 32 characters
        long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
        different from `name` or `name` spelled backwards.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies remarks of the database account. The parameter must be 1 to 512
        characters long and is supported only for MySQL 8.0.25 and later versions.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the IP addresses that are allowed to access your DB instance.
        + If the IP address is set to %, all IP addresses are allowed to access your instance.
        + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
        your instance.
        + Multiple IP addresses can be added.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the username of the db account. Only lowercase letters, digits,
        hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
        + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
        + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the rds account resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MysqlAccountState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MysqlAccount resources.
        :param pulumi.Input[str] description: Specifies remarks of the database account. The parameter must be 1 to 512
               characters long and is supported only for MySQL 8.0.25 and later versions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: Specifies the IP addresses that are allowed to access your DB instance.
               + If the IP address is set to %, all IP addresses are allowed to access your instance.
               + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
               your instance.
               + Multiple IP addresses can be added.
        :param pulumi.Input[str] instance_id: Specifies the rds instance id. Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the username of the db account. Only lowercase letters, digits,
               hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
               + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
               + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        :param pulumi.Input[str] password: Specifies the password of the db account. The parameter must be 8 to 32 characters
               long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
               different from `name` or `name` spelled backwards.
        :param pulumi.Input[str] region: The region in which to create the rds account resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies remarks of the database account. The parameter must be 1 to 512
        characters long and is supported only for MySQL 8.0.25 and later versions.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the IP addresses that are allowed to access your DB instance.
        + If the IP address is set to %, all IP addresses are allowed to access your instance.
        + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
        your instance.
        + Multiple IP addresses can be added.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the rds instance id. Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the username of the db account. Only lowercase letters, digits,
        hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
        + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
        + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the password of the db account. The parameter must be 8 to 32 characters
        long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
        different from `name` or `name` spelled backwards.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the rds account resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class MysqlAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages RDS Mysql account resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        account_password = config.require_object("accountPassword")
        test = huaweicloud.rds.MysqlAccount("test",
            instance_id=instance_id,
            password=account_password)
        ```

        ## Import

        RDS account can be imported using the `instance_id` and `name` separated by a slash, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Rds/mysqlAccount:MysqlAccount account_1 <instance_id>/<name>
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`. It is generally recommended running `terraform plan` after importing the RDS Mysql account. You can then decide if changes should be applied to the RDS Mysql account, or the resource definition should be updated to align with the RDS Mysql account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_mysql_account" "account_1" {

         ...

         lifecycle {

         ignore_changes = [

         password

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies remarks of the database account. The parameter must be 1 to 512
               characters long and is supported only for MySQL 8.0.25 and later versions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: Specifies the IP addresses that are allowed to access your DB instance.
               + If the IP address is set to %, all IP addresses are allowed to access your instance.
               + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
               your instance.
               + Multiple IP addresses can be added.
        :param pulumi.Input[str] instance_id: Specifies the rds instance id. Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the username of the db account. Only lowercase letters, digits,
               hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
               + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
               + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        :param pulumi.Input[str] password: Specifies the password of the db account. The parameter must be 8 to 32 characters
               long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
               different from `name` or `name` spelled backwards.
        :param pulumi.Input[str] region: The region in which to create the rds account resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MysqlAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages RDS Mysql account resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        account_password = config.require_object("accountPassword")
        test = huaweicloud.rds.MysqlAccount("test",
            instance_id=instance_id,
            password=account_password)
        ```

        ## Import

        RDS account can be imported using the `instance_id` and `name` separated by a slash, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Rds/mysqlAccount:MysqlAccount account_1 <instance_id>/<name>
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`. It is generally recommended running `terraform plan` after importing the RDS Mysql account. You can then decide if changes should be applied to the RDS Mysql account, or the resource definition should be updated to align with the RDS Mysql account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_mysql_account" "account_1" {

         ...

         lifecycle {

         ignore_changes = [

         password

         ]

         } }

        :param str resource_name: The name of the resource.
        :param MysqlAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MysqlAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MysqlAccountArgs.__new__(MysqlAccountArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["hosts"] = hosts
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["region"] = region
        super(MysqlAccount, __self__).__init__(
            'huaweicloud:Rds/mysqlAccount:MysqlAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'MysqlAccount':
        """
        Get an existing MysqlAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies remarks of the database account. The parameter must be 1 to 512
               characters long and is supported only for MySQL 8.0.25 and later versions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: Specifies the IP addresses that are allowed to access your DB instance.
               + If the IP address is set to %, all IP addresses are allowed to access your instance.
               + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
               your instance.
               + Multiple IP addresses can be added.
        :param pulumi.Input[str] instance_id: Specifies the rds instance id. Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the username of the db account. Only lowercase letters, digits,
               hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
               + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
               + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        :param pulumi.Input[str] password: Specifies the password of the db account. The parameter must be 8 to 32 characters
               long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
               different from `name` or `name` spelled backwards.
        :param pulumi.Input[str] region: The region in which to create the rds account resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MysqlAccountState.__new__(_MysqlAccountState)

        __props__.__dict__["description"] = description
        __props__.__dict__["hosts"] = hosts
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["region"] = region
        return MysqlAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies remarks of the database account. The parameter must be 1 to 512
        characters long and is supported only for MySQL 8.0.25 and later versions.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the IP addresses that are allowed to access your DB instance.
        + If the IP address is set to %, all IP addresses are allowed to access your instance.
        + If the IP address is set to 10.10.10.%, all IP addresses in the subnet 10.10.10.X are allowed to access
        your instance.
        + Multiple IP addresses can be added.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the rds instance id. Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the username of the db account. Only lowercase letters, digits,
        hyphens (-), and underscores (_) are allowed. Changing this will create a new resource.
        + If the database version is MySQL 5.6, the username consists of 1 to 16 characters.
        + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Specifies the password of the db account. The parameter must be 8 to 32 characters
        long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^*-_=+?,()&). The value must be
        different from `name` or `name` spelled backwards.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the rds account resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

