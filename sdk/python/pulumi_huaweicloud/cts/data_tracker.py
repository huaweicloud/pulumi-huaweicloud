# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DataTrackerArgs', 'DataTracker']

@pulumi.input_type
class DataTrackerArgs:
    def __init__(__self__, *,
                 data_bucket: pulumi.Input[str],
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 file_prefix: Optional[pulumi.Input[str]] = None,
                 lts_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obs_retention_period: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 validate_file: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DataTracker resource.
        :param pulumi.Input[str] data_bucket: Specifies the OBS bucket tracked by the data tracker.
               Changing this creates a new resource.
        :param pulumi.Input[str] bucket_name: Specifies the OBS bucket to which traces will be transferred.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_operations: Specifies an array of operation types tracked by the data tracker,
               the value of operation can be **WRITE** or **READ**.
        :param pulumi.Input[bool] enabled: Specifies whether tracker is enabled.
        :param pulumi.Input[str] file_prefix: Specifies the file name prefix to mark trace files that need to be stored
               in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
               and periods (.) are allowed.
        :param pulumi.Input[bool] lts_enabled: Specifies whether trace analysis is enabled.
        :param pulumi.Input[str] name: Specifies the data tracker name. The name cannot be system or ststem-trace.
               Changing this creates a new resource.
        :param pulumi.Input[int] obs_retention_period: Specifies the retention period that traces are stored in `bucket_name`,
               the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        :param pulumi.Input[str] region: Specifies the region in which to manage the CTS data tracker resource.
               If omitted, the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[bool] validate_file: Specifies whether trace file verification is enabled during trace transfer.
        """
        pulumi.set(__self__, "data_bucket", data_bucket)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if data_operations is not None:
            pulumi.set(__self__, "data_operations", data_operations)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if file_prefix is not None:
            pulumi.set(__self__, "file_prefix", file_prefix)
        if lts_enabled is not None:
            pulumi.set(__self__, "lts_enabled", lts_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if obs_retention_period is not None:
            pulumi.set(__self__, "obs_retention_period", obs_retention_period)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if validate_file is not None:
            pulumi.set(__self__, "validate_file", validate_file)

    @property
    @pulumi.getter(name="dataBucket")
    def data_bucket(self) -> pulumi.Input[str]:
        """
        Specifies the OBS bucket tracked by the data tracker.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "data_bucket")

    @data_bucket.setter
    def data_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_bucket", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the OBS bucket to which traces will be transferred.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="dataOperations")
    def data_operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of operation types tracked by the data tracker,
        the value of operation can be **WRITE** or **READ**.
        """
        return pulumi.get(self, "data_operations")

    @data_operations.setter
    def data_operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_operations", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether tracker is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="filePrefix")
    def file_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the file name prefix to mark trace files that need to be stored
        in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
        and periods (.) are allowed.
        """
        return pulumi.get(self, "file_prefix")

    @file_prefix.setter
    def file_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_prefix", value)

    @property
    @pulumi.getter(name="ltsEnabled")
    def lts_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether trace analysis is enabled.
        """
        return pulumi.get(self, "lts_enabled")

    @lts_enabled.setter
    def lts_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lts_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the data tracker name. The name cannot be system or ststem-trace.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="obsRetentionPeriod")
    def obs_retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the retention period that traces are stored in `bucket_name`,
        the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        """
        return pulumi.get(self, "obs_retention_period")

    @obs_retention_period.setter
    def obs_retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "obs_retention_period", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to manage the CTS data tracker resource.
        If omitted, the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="validateFile")
    def validate_file(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether trace file verification is enabled during trace transfer.
        """
        return pulumi.get(self, "validate_file")

    @validate_file.setter
    def validate_file(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_file", value)


@pulumi.input_type
class _DataTrackerState:
    def __init__(__self__, *,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_bucket: Optional[pulumi.Input[str]] = None,
                 data_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 file_prefix: Optional[pulumi.Input[str]] = None,
                 lts_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obs_retention_period: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transfer_enabled: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 validate_file: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DataTracker resources.
        :param pulumi.Input[str] bucket_name: Specifies the OBS bucket to which traces will be transferred.
        :param pulumi.Input[str] data_bucket: Specifies the OBS bucket tracked by the data tracker.
               Changing this creates a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_operations: Specifies an array of operation types tracked by the data tracker,
               the value of operation can be **WRITE** or **READ**.
        :param pulumi.Input[bool] enabled: Specifies whether tracker is enabled.
        :param pulumi.Input[str] file_prefix: Specifies the file name prefix to mark trace files that need to be stored
               in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
               and periods (.) are allowed.
        :param pulumi.Input[bool] lts_enabled: Specifies whether trace analysis is enabled.
        :param pulumi.Input[str] name: Specifies the data tracker name. The name cannot be system or ststem-trace.
               Changing this creates a new resource.
        :param pulumi.Input[int] obs_retention_period: Specifies the retention period that traces are stored in `bucket_name`,
               the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        :param pulumi.Input[str] region: Specifies the region in which to manage the CTS data tracker resource.
               If omitted, the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] status: The tracker status, the value can be **enabled**, **disabled** or **error**.
        :param pulumi.Input[bool] transfer_enabled: Whether traces will be transferred.
        :param pulumi.Input[str] type: The tracker type, only **data** is available.
        :param pulumi.Input[bool] validate_file: Specifies whether trace file verification is enabled during trace transfer.
        """
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if data_bucket is not None:
            pulumi.set(__self__, "data_bucket", data_bucket)
        if data_operations is not None:
            pulumi.set(__self__, "data_operations", data_operations)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if file_prefix is not None:
            pulumi.set(__self__, "file_prefix", file_prefix)
        if lts_enabled is not None:
            pulumi.set(__self__, "lts_enabled", lts_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if obs_retention_period is not None:
            pulumi.set(__self__, "obs_retention_period", obs_retention_period)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transfer_enabled is not None:
            pulumi.set(__self__, "transfer_enabled", transfer_enabled)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if validate_file is not None:
            pulumi.set(__self__, "validate_file", validate_file)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the OBS bucket to which traces will be transferred.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="dataBucket")
    def data_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the OBS bucket tracked by the data tracker.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "data_bucket")

    @data_bucket.setter
    def data_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_bucket", value)

    @property
    @pulumi.getter(name="dataOperations")
    def data_operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of operation types tracked by the data tracker,
        the value of operation can be **WRITE** or **READ**.
        """
        return pulumi.get(self, "data_operations")

    @data_operations.setter
    def data_operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_operations", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether tracker is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="filePrefix")
    def file_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the file name prefix to mark trace files that need to be stored
        in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
        and periods (.) are allowed.
        """
        return pulumi.get(self, "file_prefix")

    @file_prefix.setter
    def file_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_prefix", value)

    @property
    @pulumi.getter(name="ltsEnabled")
    def lts_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether trace analysis is enabled.
        """
        return pulumi.get(self, "lts_enabled")

    @lts_enabled.setter
    def lts_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lts_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the data tracker name. The name cannot be system or ststem-trace.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="obsRetentionPeriod")
    def obs_retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the retention period that traces are stored in `bucket_name`,
        the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        """
        return pulumi.get(self, "obs_retention_period")

    @obs_retention_period.setter
    def obs_retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "obs_retention_period", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to manage the CTS data tracker resource.
        If omitted, the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The tracker status, the value can be **enabled**, **disabled** or **error**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transferEnabled")
    def transfer_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether traces will be transferred.
        """
        return pulumi.get(self, "transfer_enabled")

    @transfer_enabled.setter
    def transfer_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transfer_enabled", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The tracker type, only **data** is available.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="validateFile")
    def validate_file(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether trace file verification is enabled during trace transfer.
        """
        return pulumi.get(self, "validate_file")

    @validate_file.setter
    def validate_file(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_file", value)


class DataTracker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_bucket: Optional[pulumi.Input[str]] = None,
                 data_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 file_prefix: Optional[pulumi.Input[str]] = None,
                 lts_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obs_retention_period: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 validate_file: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages CTS **data** tracker resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        data_bucket = config.require_object("dataBucket")
        transfer_bucket = config.require_object("transferBucket")
        tracker = huaweicloud.cts.DataTracker("tracker",
            data_bucket=data_bucket,
            bucket_name=transfer_bucket,
            file_prefix="cloudTrace",
            lts_enabled=True)
        ```

        ## Import

        CTS data tracker can be imported using `name`, e.g.

        ```sh
         $ pulumi import huaweicloud:Cts/dataTracker:DataTracker tracker your_tracker_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: Specifies the OBS bucket to which traces will be transferred.
        :param pulumi.Input[str] data_bucket: Specifies the OBS bucket tracked by the data tracker.
               Changing this creates a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_operations: Specifies an array of operation types tracked by the data tracker,
               the value of operation can be **WRITE** or **READ**.
        :param pulumi.Input[bool] enabled: Specifies whether tracker is enabled.
        :param pulumi.Input[str] file_prefix: Specifies the file name prefix to mark trace files that need to be stored
               in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
               and periods (.) are allowed.
        :param pulumi.Input[bool] lts_enabled: Specifies whether trace analysis is enabled.
        :param pulumi.Input[str] name: Specifies the data tracker name. The name cannot be system or ststem-trace.
               Changing this creates a new resource.
        :param pulumi.Input[int] obs_retention_period: Specifies the retention period that traces are stored in `bucket_name`,
               the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        :param pulumi.Input[str] region: Specifies the region in which to manage the CTS data tracker resource.
               If omitted, the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[bool] validate_file: Specifies whether trace file verification is enabled during trace transfer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataTrackerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages CTS **data** tracker resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        data_bucket = config.require_object("dataBucket")
        transfer_bucket = config.require_object("transferBucket")
        tracker = huaweicloud.cts.DataTracker("tracker",
            data_bucket=data_bucket,
            bucket_name=transfer_bucket,
            file_prefix="cloudTrace",
            lts_enabled=True)
        ```

        ## Import

        CTS data tracker can be imported using `name`, e.g.

        ```sh
         $ pulumi import huaweicloud:Cts/dataTracker:DataTracker tracker your_tracker_name
        ```

        :param str resource_name: The name of the resource.
        :param DataTrackerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataTrackerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_bucket: Optional[pulumi.Input[str]] = None,
                 data_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 file_prefix: Optional[pulumi.Input[str]] = None,
                 lts_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obs_retention_period: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 validate_file: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataTrackerArgs.__new__(DataTrackerArgs)

            __props__.__dict__["bucket_name"] = bucket_name
            if data_bucket is None and not opts.urn:
                raise TypeError("Missing required property 'data_bucket'")
            __props__.__dict__["data_bucket"] = data_bucket
            __props__.__dict__["data_operations"] = data_operations
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["file_prefix"] = file_prefix
            __props__.__dict__["lts_enabled"] = lts_enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["obs_retention_period"] = obs_retention_period
            __props__.__dict__["region"] = region
            __props__.__dict__["validate_file"] = validate_file
            __props__.__dict__["status"] = None
            __props__.__dict__["transfer_enabled"] = None
            __props__.__dict__["type"] = None
        super(DataTracker, __self__).__init__(
            'huaweicloud:Cts/dataTracker:DataTracker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            data_bucket: Optional[pulumi.Input[str]] = None,
            data_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            file_prefix: Optional[pulumi.Input[str]] = None,
            lts_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            obs_retention_period: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transfer_enabled: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None,
            validate_file: Optional[pulumi.Input[bool]] = None) -> 'DataTracker':
        """
        Get an existing DataTracker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: Specifies the OBS bucket to which traces will be transferred.
        :param pulumi.Input[str] data_bucket: Specifies the OBS bucket tracked by the data tracker.
               Changing this creates a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_operations: Specifies an array of operation types tracked by the data tracker,
               the value of operation can be **WRITE** or **READ**.
        :param pulumi.Input[bool] enabled: Specifies whether tracker is enabled.
        :param pulumi.Input[str] file_prefix: Specifies the file name prefix to mark trace files that need to be stored
               in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
               and periods (.) are allowed.
        :param pulumi.Input[bool] lts_enabled: Specifies whether trace analysis is enabled.
        :param pulumi.Input[str] name: Specifies the data tracker name. The name cannot be system or ststem-trace.
               Changing this creates a new resource.
        :param pulumi.Input[int] obs_retention_period: Specifies the retention period that traces are stored in `bucket_name`,
               the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        :param pulumi.Input[str] region: Specifies the region in which to manage the CTS data tracker resource.
               If omitted, the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] status: The tracker status, the value can be **enabled**, **disabled** or **error**.
        :param pulumi.Input[bool] transfer_enabled: Whether traces will be transferred.
        :param pulumi.Input[str] type: The tracker type, only **data** is available.
        :param pulumi.Input[bool] validate_file: Specifies whether trace file verification is enabled during trace transfer.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataTrackerState.__new__(_DataTrackerState)

        __props__.__dict__["bucket_name"] = bucket_name
        __props__.__dict__["data_bucket"] = data_bucket
        __props__.__dict__["data_operations"] = data_operations
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["file_prefix"] = file_prefix
        __props__.__dict__["lts_enabled"] = lts_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["obs_retention_period"] = obs_retention_period
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["transfer_enabled"] = transfer_enabled
        __props__.__dict__["type"] = type
        __props__.__dict__["validate_file"] = validate_file
        return DataTracker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the OBS bucket to which traces will be transferred.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="dataBucket")
    def data_bucket(self) -> pulumi.Output[str]:
        """
        Specifies the OBS bucket tracked by the data tracker.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "data_bucket")

    @property
    @pulumi.getter(name="dataOperations")
    def data_operations(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies an array of operation types tracked by the data tracker,
        the value of operation can be **WRITE** or **READ**.
        """
        return pulumi.get(self, "data_operations")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether tracker is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="filePrefix")
    def file_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the file name prefix to mark trace files that need to be stored
        in an OBS bucket. The value contains 0 to 64 characters. Only letters, numbers, hyphens (-), underscores (_),
        and periods (.) are allowed.
        """
        return pulumi.get(self, "file_prefix")

    @property
    @pulumi.getter(name="ltsEnabled")
    def lts_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether trace analysis is enabled.
        """
        return pulumi.get(self, "lts_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the data tracker name. The name cannot be system or ststem-trace.
        Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="obsRetentionPeriod")
    def obs_retention_period(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the retention period that traces are stored in `bucket_name`,
        the value can be **0**(permanent), **30**, **60**, **90**, **180** or **1095**.
        """
        return pulumi.get(self, "obs_retention_period")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to manage the CTS data tracker resource.
        If omitted, the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The tracker status, the value can be **enabled**, **disabled** or **error**.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transferEnabled")
    def transfer_enabled(self) -> pulumi.Output[bool]:
        """
        Whether traces will be transferred.
        """
        return pulumi.get(self, "transfer_enabled")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The tracker type, only **data** is available.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validateFile")
    def validate_file(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether trace file verification is enabled during trace transfer.
        """
        return pulumi.get(self, "validate_file")
