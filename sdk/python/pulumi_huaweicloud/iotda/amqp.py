# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AmqpArgs', 'Amqp']

@pulumi.input_type
class AmqpArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Amqp resource.
        :param pulumi.Input[str] name: Specifies the AMQP queue name, which contains 8 to 128 characters.
               Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA AMQP queue resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AMQP queue name, which contains 8 to 128 characters.
        Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the IoTDA AMQP queue resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _AmqpState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Amqp resources.
        :param pulumi.Input[str] name: Specifies the AMQP queue name, which contains 8 to 128 characters.
               Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA AMQP queue resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AMQP queue name, which contains 8 to 128 characters.
        Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the IoTDA AMQP queue resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Amqp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an IoTDA AMQP queue within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        queue = huaweicloud.io_tda.Amqp("queue")
        ```

        ## Import

        AMQP queues can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:IoTDA/amqp:Amqp test 10022532f4f94f26b01daa1e424853e1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the AMQP queue name, which contains 8 to 128 characters.
               Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA AMQP queue resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AmqpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an IoTDA AMQP queue within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        queue = huaweicloud.io_tda.Amqp("queue")
        ```

        ## Import

        AMQP queues can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:IoTDA/amqp:Amqp test 10022532f4f94f26b01daa1e424853e1
        ```

        :param str resource_name: The name of the resource.
        :param AmqpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AmqpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AmqpArgs.__new__(AmqpArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
        super(Amqp, __self__).__init__(
            'huaweicloud:IoTDA/amqp:Amqp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Amqp':
        """
        Get an existing Amqp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the AMQP queue name, which contains 8 to 128 characters.
               Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA AMQP queue resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AmqpState.__new__(_AmqpState)

        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        return Amqp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the AMQP queue name, which contains 8 to 128 characters.
        Only letters, digits, hyphens (-), underscores (_), dots (.) and colons (:) are allowed.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the IoTDA AMQP queue resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

