# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogtankArgs', 'Logtank']

@pulumi.input_type
class LogtankArgs:
    def __init__(__self__, *,
                 loadbalancer_id: pulumi.Input[str],
                 log_group_id: pulumi.Input[str],
                 log_topic_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Logtank resource.
        :param pulumi.Input[str] loadbalancer_id: Specifies the ID of a loadbalancer. Changing this
               creates a new logtank
        :param pulumi.Input[str] log_group_id: Specifies the ID of a log group. It is provided by other service.
        :param pulumi.Input[str] log_topic_id: Specifies the ID of the subscribe topic.
        :param pulumi.Input[str] region: The region in which to create the logtank resource.
               If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        pulumi.set(__self__, "log_group_id", log_group_id)
        pulumi.set(__self__, "log_topic_id", log_topic_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of a loadbalancer. Changing this
        creates a new logtank
        """
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter(name="logGroupId")
    def log_group_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of a log group. It is provided by other service.
        """
        return pulumi.get(self, "log_group_id")

    @log_group_id.setter
    def log_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_group_id", value)

    @property
    @pulumi.getter(name="logTopicId")
    def log_topic_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the subscribe topic.
        """
        return pulumi.get(self, "log_topic_id")

    @log_topic_id.setter
    def log_topic_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_topic_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the logtank resource.
        If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _LogtankState:
    def __init__(__self__, *,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 log_group_id: Optional[pulumi.Input[str]] = None,
                 log_topic_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Logtank resources.
        :param pulumi.Input[str] loadbalancer_id: Specifies the ID of a loadbalancer. Changing this
               creates a new logtank
        :param pulumi.Input[str] log_group_id: Specifies the ID of a log group. It is provided by other service.
        :param pulumi.Input[str] log_topic_id: Specifies the ID of the subscribe topic.
        :param pulumi.Input[str] region: The region in which to create the logtank resource.
               If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if log_group_id is not None:
            pulumi.set(__self__, "log_group_id", log_group_id)
        if log_topic_id is not None:
            pulumi.set(__self__, "log_topic_id", log_topic_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of a loadbalancer. Changing this
        creates a new logtank
        """
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter(name="logGroupId")
    def log_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of a log group. It is provided by other service.
        """
        return pulumi.get(self, "log_group_id")

    @log_group_id.setter
    def log_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group_id", value)

    @property
    @pulumi.getter(name="logTopicId")
    def log_topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the subscribe topic.
        """
        return pulumi.get(self, "log_topic_id")

    @log_topic_id.setter
    def log_topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_topic_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the logtank resource.
        If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Logtank(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 log_group_id: Optional[pulumi.Input[str]] = None,
                 log_topic_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage an ELB logtank resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        loadbalancer_id = config.require_object("loadbalancerId")
        group_id = config.require_object("groupId")
        topic_id = config.require_object("topicId")
        test = huaweicloud.dedicated_elb.Logtank("test",
            loadbalancer_id=loadbalancer_id,
            log_group_id=group_id,
            log_topic_id=topic_id)
        ```

        ## Import

        ELB logtank can be imported using the logtank ID, e.g. bash

        ```sh
         $ pulumi import huaweicloud:DedicatedElb/logtank:Logtank test 2f148a75-acd3-4ce7-8f63-d5c9fadab3a0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] loadbalancer_id: Specifies the ID of a loadbalancer. Changing this
               creates a new logtank
        :param pulumi.Input[str] log_group_id: Specifies the ID of a log group. It is provided by other service.
        :param pulumi.Input[str] log_topic_id: Specifies the ID of the subscribe topic.
        :param pulumi.Input[str] region: The region in which to create the logtank resource.
               If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogtankArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage an ELB logtank resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        loadbalancer_id = config.require_object("loadbalancerId")
        group_id = config.require_object("groupId")
        topic_id = config.require_object("topicId")
        test = huaweicloud.dedicated_elb.Logtank("test",
            loadbalancer_id=loadbalancer_id,
            log_group_id=group_id,
            log_topic_id=topic_id)
        ```

        ## Import

        ELB logtank can be imported using the logtank ID, e.g. bash

        ```sh
         $ pulumi import huaweicloud:DedicatedElb/logtank:Logtank test 2f148a75-acd3-4ce7-8f63-d5c9fadab3a0
        ```

        :param str resource_name: The name of the resource.
        :param LogtankArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogtankArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 log_group_id: Optional[pulumi.Input[str]] = None,
                 log_topic_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogtankArgs.__new__(LogtankArgs)

            if loadbalancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__.__dict__["loadbalancer_id"] = loadbalancer_id
            if log_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_group_id'")
            __props__.__dict__["log_group_id"] = log_group_id
            if log_topic_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_topic_id'")
            __props__.__dict__["log_topic_id"] = log_topic_id
            __props__.__dict__["region"] = region
        super(Logtank, __self__).__init__(
            'huaweicloud:DedicatedElb/logtank:Logtank',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            log_group_id: Optional[pulumi.Input[str]] = None,
            log_topic_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Logtank':
        """
        Get an existing Logtank resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] loadbalancer_id: Specifies the ID of a loadbalancer. Changing this
               creates a new logtank
        :param pulumi.Input[str] log_group_id: Specifies the ID of a log group. It is provided by other service.
        :param pulumi.Input[str] log_topic_id: Specifies the ID of the subscribe topic.
        :param pulumi.Input[str] region: The region in which to create the logtank resource.
               If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogtankState.__new__(_LogtankState)

        __props__.__dict__["loadbalancer_id"] = loadbalancer_id
        __props__.__dict__["log_group_id"] = log_group_id
        __props__.__dict__["log_topic_id"] = log_topic_id
        __props__.__dict__["region"] = region
        return Logtank(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of a loadbalancer. Changing this
        creates a new logtank
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter(name="logGroupId")
    def log_group_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of a log group. It is provided by other service.
        """
        return pulumi.get(self, "log_group_id")

    @property
    @pulumi.getter(name="logTopicId")
    def log_topic_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the subscribe topic.
        """
        return pulumi.get(self, "log_topic_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the logtank resource.
        If omitted, the provider-level region will be used. Changing this creates a new logtank.
        """
        return pulumi.get(self, "region")

