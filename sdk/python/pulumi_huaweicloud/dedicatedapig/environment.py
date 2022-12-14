# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API
               environment belongs to. Changing this will create a new APIG environment resource.
        :param pulumi.Input[str] description: Specifies the description about the API environment. The description contain a
               maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
               Unicode format.
        :param pulumi.Input[str] name: Specifies the name of the API environment. The API environment name consists of 3 to 64
               characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region in which to create the APIG environment resource. If
               omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the API
        environment belongs to. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description about the API environment. The description contain a
        maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
        Unicode format.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API environment. The API environment name consists of 3 to 64
        characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the APIG environment resource. If
        omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _EnvironmentState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Environment resources.
        :param pulumi.Input[str] create_time: Time when the APIG environment was created, in RFC-3339 format.
        :param pulumi.Input[str] description: Specifies the description about the API environment. The description contain a
               maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
               Unicode format.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API
               environment belongs to. Changing this will create a new APIG environment resource.
        :param pulumi.Input[str] name: Specifies the name of the API environment. The API environment name consists of 3 to 64
               characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region in which to create the APIG environment resource. If
               omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time when the APIG environment was created, in RFC-3339 format.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description about the API environment. The description contain a
        maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
        Unicode format.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies an ID of the APIG dedicated instance to which the API
        environment belongs to. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API environment. The API environment name consists of 3 to 64
        characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the APIG environment resource. If
        omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an APIG environment resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        environment_name = config.require_object("environmentName")
        description = config.require_object("description")
        test = huaweicloud.dedicated_apig.Environment("test",
            instance_id=instance_id,
            description=description)
        ```

        ## Import

        Environments can be imported using their `id` and the ID of the APIG instance to which the environment belongs, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/environment:Environment test <instance ID>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description about the API environment. The description contain a
               maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
               Unicode format.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API
               environment belongs to. Changing this will create a new APIG environment resource.
        :param pulumi.Input[str] name: Specifies the name of the API environment. The API environment name consists of 3 to 64
               characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region in which to create the APIG environment resource. If
               omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an APIG environment resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        environment_name = config.require_object("environmentName")
        description = config.require_object("description")
        test = huaweicloud.dedicated_apig.Environment("test",
            instance_id=instance_id,
            description=description)
        ```

        ## Import

        Environments can be imported using their `id` and the ID of the APIG instance to which the environment belongs, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/environment:Environment test <instance ID>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["create_time"] = None
        super(Environment, __self__).__init__(
            'huaweicloud:DedicatedApig/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time when the APIG environment was created, in RFC-3339 format.
        :param pulumi.Input[str] description: Specifies the description about the API environment. The description contain a
               maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
               Unicode format.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API
               environment belongs to. Changing this will create a new APIG environment resource.
        :param pulumi.Input[str] name: Specifies the name of the API environment. The API environment name consists of 3 to 64
               characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region in which to create the APIG environment resource. If
               omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentState.__new__(_EnvironmentState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time when the APIG environment was created, in RFC-3339 format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description about the API environment. The description contain a
        maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
        Unicode format.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the API
        environment belongs to. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the API environment. The API environment name consists of 3 to 64
        characters, starting with a letter. Only letters, digits and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the APIG environment resource. If
        omitted, the provider-level region will be used. Changing this will create a new APIG environment resource.
        """
        return pulumi.get(self, "region")

