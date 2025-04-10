# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StreamArgs', 'Stream']

@pulumi.input_type
class StreamArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 stream_name: pulumi.Input[str],
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl_in_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Stream resource.
        :param pulumi.Input[str] group_id: Specifies the ID of a created log group. Changing this parameter will create
               a new resource.
        :param pulumi.Input[str] stream_name: Specifies the log stream name. Changing this parameter will create a new
               resource.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the log stream resource. If omitted, the
               provider-level region will be used. Changing this creates a new log stream resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs of the log stream.
        :param pulumi.Input[int] ttl_in_days: Specifies the log expiration time (days).
               The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "stream_name", stream_name)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl_in_days is not None:
            pulumi.set(__self__, "ttl_in_days", ttl_in_days)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of a created log group. Changing this parameter will create
        a new resource.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Input[str]:
        """
        Specifies the log stream name. Changing this parameter will create a new
        resource.
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the log stream resource. If omitted, the
        provider-level region will be used. Changing this creates a new log stream resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs of the log stream.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="ttlInDays")
    def ttl_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the log expiration time (days).
        The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        return pulumi.get(self, "ttl_in_days")

    @ttl_in_days.setter
    def ttl_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl_in_days", value)


@pulumi.input_type
class _StreamState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 filter_count: Optional[pulumi.Input[int]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl_in_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Stream resources.
        :param pulumi.Input[str] created_at: The creation time of the log stream.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] filter_count: Number of log stream filters.
        :param pulumi.Input[str] group_id: Specifies the ID of a created log group. Changing this parameter will create
               a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the log stream resource. If omitted, the
               provider-level region will be used. Changing this creates a new log stream resource.
        :param pulumi.Input[str] stream_name: Specifies the log stream name. Changing this parameter will create a new
               resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs of the log stream.
        :param pulumi.Input[int] ttl_in_days: Specifies the log expiration time (days).
               The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if filter_count is not None:
            pulumi.set(__self__, "filter_count", filter_count)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl_in_days is not None:
            pulumi.set(__self__, "ttl_in_days", ttl_in_days)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the log stream.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter(name="filterCount")
    def filter_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of log stream filters.
        """
        return pulumi.get(self, "filter_count")

    @filter_count.setter
    def filter_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "filter_count", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of a created log group. Changing this parameter will create
        a new resource.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the log stream resource. If omitted, the
        provider-level region will be used. Changing this creates a new log stream resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the log stream name. Changing this parameter will create a new
        resource.
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs of the log stream.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="ttlInDays")
    def ttl_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the log expiration time (days).
        The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        return pulumi.get(self, "ttl_in_days")

    @ttl_in_days.setter
    def ttl_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl_in_days", value)


class Stream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl_in_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manage a log stream resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        group_id = config.require_object("groupId")
        test = huaweicloud.lts.Stream("test",
            group_id=group_id,
            stream_name="testacc_stream")
        ```

        ## Import

        The log stream can be imported using the group ID and stream ID separated by a slash, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Lts/stream:Stream test <group_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] group_id: Specifies the ID of a created log group. Changing this parameter will create
               a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the log stream resource. If omitted, the
               provider-level region will be used. Changing this creates a new log stream resource.
        :param pulumi.Input[str] stream_name: Specifies the log stream name. Changing this parameter will create a new
               resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs of the log stream.
        :param pulumi.Input[int] ttl_in_days: Specifies the log expiration time (days).
               The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a log stream resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        group_id = config.require_object("groupId")
        test = huaweicloud.lts.Stream("test",
            group_id=group_id,
            stream_name="testacc_stream")
        ```

        ## Import

        The log stream can be imported using the group ID and stream ID separated by a slash, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Lts/stream:Stream test <group_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param StreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl_in_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamArgs.__new__(StreamArgs)

            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["region"] = region
            if stream_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_name'")
            __props__.__dict__["stream_name"] = stream_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["ttl_in_days"] = ttl_in_days
            __props__.__dict__["created_at"] = None
            __props__.__dict__["filter_count"] = None
        super(Stream, __self__).__init__(
            'huaweicloud:Lts/stream:Stream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            enterprise_project_id: Optional[pulumi.Input[str]] = None,
            filter_count: Optional[pulumi.Input[int]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            stream_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            ttl_in_days: Optional[pulumi.Input[int]] = None) -> 'Stream':
        """
        Get an existing Stream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The creation time of the log stream.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] filter_count: Number of log stream filters.
        :param pulumi.Input[str] group_id: Specifies the ID of a created log group. Changing this parameter will create
               a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the log stream resource. If omitted, the
               provider-level region will be used. Changing this creates a new log stream resource.
        :param pulumi.Input[str] stream_name: Specifies the log stream name. Changing this parameter will create a new
               resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs of the log stream.
        :param pulumi.Input[int] ttl_in_days: Specifies the log expiration time (days).
               The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StreamState.__new__(_StreamState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["filter_count"] = filter_count
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["region"] = region
        __props__.__dict__["stream_name"] = stream_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["ttl_in_days"] = ttl_in_days
        return Stream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time of the log stream.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[str]:
        """
        Specifies the enterprise project ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter(name="filterCount")
    def filter_count(self) -> pulumi.Output[int]:
        """
        Number of log stream filters.
        """
        return pulumi.get(self, "filter_count")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of a created log group. Changing this parameter will create
        a new resource.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the log stream resource. If omitted, the
        provider-level region will be used. Changing this creates a new log stream resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Output[str]:
        """
        Specifies the log stream name. Changing this parameter will create a new
        resource.
        """
        return pulumi.get(self, "stream_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the key/value pairs of the log stream.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="ttlInDays")
    def ttl_in_days(self) -> pulumi.Output[int]:
        """
        Specifies the log expiration time (days).
        The valid value is a non-zero integer from `-1` to `365`, defaults to `-1` which means inherit the log group settings.
        """
        return pulumi.get(self, "ttl_in_days")

