# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CrossRegionBackupStrategyArgs', 'CrossRegionBackupStrategy']

@pulumi.input_type
class CrossRegionBackupStrategyArgs:
    def __init__(__self__, *,
                 backup_type: pulumi.Input[str],
                 destination_project_id: pulumi.Input[str],
                 destination_region: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 keep_days: pulumi.Input[int],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CrossRegionBackupStrategy resource.
        :param pulumi.Input[str] backup_type: Specifies the backup type. Value options:
               + **auto**: open automated full backup.
               + **all**: open both automated full backup and automated incremental backup.
        :param pulumi.Input[str] destination_project_id: Specifies the target project ID for the cross-region backup
               policy.
        :param pulumi.Input[str] destination_region: Specifies the target region ID for the cross-region backup policy.
        :param pulumi.Input[str] instance_id: Specifies the ID of the RDS instance.
        :param pulumi.Input[int] keep_days: Specifies the number of days to retain the generated backup files.
               Value ranges from `1` to `1,825`.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "destination_project_id", destination_project_id)
        pulumi.set(__self__, "destination_region", destination_region)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "keep_days", keep_days)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Input[str]:
        """
        Specifies the backup type. Value options:
        + **auto**: open automated full backup.
        + **all**: open both automated full backup and automated incremental backup.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> pulumi.Input[str]:
        """
        Specifies the target project ID for the cross-region backup
        policy.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Input[str]:
        """
        Specifies the target region ID for the cross-region backup policy.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="keepDays")
    def keep_days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days to retain the generated backup files.
        Value ranges from `1` to `1,825`.
        """
        return pulumi.get(self, "keep_days")

    @keep_days.setter
    def keep_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "keep_days", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _CrossRegionBackupStrategyState:
    def __init__(__self__, *,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_days: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CrossRegionBackupStrategy resources.
        :param pulumi.Input[str] backup_type: Specifies the backup type. Value options:
               + **auto**: open automated full backup.
               + **all**: open both automated full backup and automated incremental backup.
        :param pulumi.Input[str] destination_project_id: Specifies the target project ID for the cross-region backup
               policy.
        :param pulumi.Input[str] destination_region: Specifies the target region ID for the cross-region backup policy.
        :param pulumi.Input[str] instance_id: Specifies the ID of the RDS instance.
        :param pulumi.Input[int] keep_days: Specifies the number of days to retain the generated backup files.
               Value ranges from `1` to `1,825`.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if destination_project_id is not None:
            pulumi.set(__self__, "destination_project_id", destination_project_id)
        if destination_region is not None:
            pulumi.set(__self__, "destination_region", destination_region)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if keep_days is not None:
            pulumi.set(__self__, "keep_days", keep_days)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the backup type. Value options:
        + **auto**: open automated full backup.
        + **all**: open both automated full backup and automated incremental backup.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the target project ID for the cross-region backup
        policy.
        """
        return pulumi.get(self, "destination_project_id")

    @destination_project_id.setter
    def destination_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_project_id", value)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the target region ID for the cross-region backup policy.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="keepDays")
    def keep_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days to retain the generated backup files.
        Value ranges from `1` to `1,825`.
        """
        return pulumi.get(self, "keep_days")

    @keep_days.setter
    def keep_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keep_days", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class CrossRegionBackupStrategy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_days: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages RDS cross-region backup strategy resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        destination_region = config.require_object("destinationRegion")
        destination_project_id = config.require_object("destinationProjectId")
        test = huaweicloud.rds.CrossRegionBackupStrategy("test",
            instance_id=instance_id,
            backup_type="all",
            keep_days=5,
            destination_region=destination_region,
            destination_project_id=destination_project_id)
        ```

        ## Import

        The RDS cross-region backup strategy can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Rds/crossRegionBackupStrategy:CrossRegionBackupStrategy test <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Specifies the backup type. Value options:
               + **auto**: open automated full backup.
               + **all**: open both automated full backup and automated incremental backup.
        :param pulumi.Input[str] destination_project_id: Specifies the target project ID for the cross-region backup
               policy.
        :param pulumi.Input[str] destination_region: Specifies the target region ID for the cross-region backup policy.
        :param pulumi.Input[str] instance_id: Specifies the ID of the RDS instance.
        :param pulumi.Input[int] keep_days: Specifies the number of days to retain the generated backup files.
               Value ranges from `1` to `1,825`.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CrossRegionBackupStrategyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages RDS cross-region backup strategy resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        destination_region = config.require_object("destinationRegion")
        destination_project_id = config.require_object("destinationProjectId")
        test = huaweicloud.rds.CrossRegionBackupStrategy("test",
            instance_id=instance_id,
            backup_type="all",
            keep_days=5,
            destination_region=destination_region,
            destination_project_id=destination_project_id)
        ```

        ## Import

        The RDS cross-region backup strategy can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Rds/crossRegionBackupStrategy:CrossRegionBackupStrategy test <id>
        ```

        :param str resource_name: The name of the resource.
        :param CrossRegionBackupStrategyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CrossRegionBackupStrategyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 destination_project_id: Optional[pulumi.Input[str]] = None,
                 destination_region: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_days: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CrossRegionBackupStrategyArgs.__new__(CrossRegionBackupStrategyArgs)

            if backup_type is None and not opts.urn:
                raise TypeError("Missing required property 'backup_type'")
            __props__.__dict__["backup_type"] = backup_type
            if destination_project_id is None and not opts.urn:
                raise TypeError("Missing required property 'destination_project_id'")
            __props__.__dict__["destination_project_id"] = destination_project_id
            if destination_region is None and not opts.urn:
                raise TypeError("Missing required property 'destination_region'")
            __props__.__dict__["destination_region"] = destination_region
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if keep_days is None and not opts.urn:
                raise TypeError("Missing required property 'keep_days'")
            __props__.__dict__["keep_days"] = keep_days
            __props__.__dict__["region"] = region
        super(CrossRegionBackupStrategy, __self__).__init__(
            'huaweicloud:Rds/crossRegionBackupStrategy:CrossRegionBackupStrategy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            destination_project_id: Optional[pulumi.Input[str]] = None,
            destination_region: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            keep_days: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'CrossRegionBackupStrategy':
        """
        Get an existing CrossRegionBackupStrategy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Specifies the backup type. Value options:
               + **auto**: open automated full backup.
               + **all**: open both automated full backup and automated incremental backup.
        :param pulumi.Input[str] destination_project_id: Specifies the target project ID for the cross-region backup
               policy.
        :param pulumi.Input[str] destination_region: Specifies the target region ID for the cross-region backup policy.
        :param pulumi.Input[str] instance_id: Specifies the ID of the RDS instance.
        :param pulumi.Input[int] keep_days: Specifies the number of days to retain the generated backup files.
               Value ranges from `1` to `1,825`.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CrossRegionBackupStrategyState.__new__(_CrossRegionBackupStrategyState)

        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["destination_project_id"] = destination_project_id
        __props__.__dict__["destination_region"] = destination_region
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["keep_days"] = keep_days
        __props__.__dict__["region"] = region
        return CrossRegionBackupStrategy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        Specifies the backup type. Value options:
        + **auto**: open automated full backup.
        + **all**: open both automated full backup and automated incremental backup.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="destinationProjectId")
    def destination_project_id(self) -> pulumi.Output[str]:
        """
        Specifies the target project ID for the cross-region backup
        policy.
        """
        return pulumi.get(self, "destination_project_id")

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Output[str]:
        """
        Specifies the target region ID for the cross-region backup policy.
        """
        return pulumi.get(self, "destination_region")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="keepDays")
    def keep_days(self) -> pulumi.Output[int]:
        """
        Specifies the number of days to retain the generated backup files.
        Value ranges from `1` to `1,825`.
        """
        return pulumi.get(self, "keep_days")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

