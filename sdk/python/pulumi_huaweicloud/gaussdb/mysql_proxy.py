# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MysqlProxyArgs', 'MysqlProxy']

@pulumi.input_type
class MysqlProxyArgs:
    def __init__(__self__, *,
                 flavor: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 node_num: pulumi.Input[int],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MysqlProxy resource.
        :param pulumi.Input[str] flavor: Specifies the flavor of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the instance ID of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] node_num: Specifies the node count of the proxy.
        :param pulumi.Input[str] region: The region in which to create the GaussDB mysql proxy resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        """
        pulumi.set(__self__, "flavor", flavor)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "node_num", node_num)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Input[str]:
        """
        Specifies the flavor of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: pulumi.Input[str]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the instance ID of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Input[int]:
        """
        Specifies the node count of the proxy.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the GaussDB mysql proxy resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MysqlProxyState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MysqlProxy resources.
        :param pulumi.Input[str] address: Indicates the address of the proxy.
        :param pulumi.Input[str] flavor: Specifies the flavor of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the instance ID of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] node_num: Specifies the node count of the proxy.
        :param pulumi.Input[int] port: Indicates the port of the proxy.
        :param pulumi.Input[str] region: The region in which to create the GaussDB mysql proxy resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if node_num is not None:
            pulumi.set(__self__, "node_num", node_num)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the address of the proxy.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the flavor of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance ID of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the node count of the proxy.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Indicates the port of the proxy.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the GaussDB mysql proxy resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class MysqlProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        GaussDB mysql proxy management within HuaweiCoud.

        ## Example Usage
        ### create a proxy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        proxy1 = huaweicloud.gauss_db.MysqlProxy("proxy1",
            instance_id=instance_id,
            flavor="gaussdb.proxy.xlarge.arm.2",
            node_num=3)
        ```

        ## Import

        GaussDB instance can be imported using the instance `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:GaussDB/mysqlProxy:MysqlProxy proxy_1 ee678f40-ce8e-4d0c-8221-38dead426f06
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flavor: Specifies the flavor of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the instance ID of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] node_num: Specifies the node count of the proxy.
        :param pulumi.Input[str] region: The region in which to create the GaussDB mysql proxy resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MysqlProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        GaussDB mysql proxy management within HuaweiCoud.

        ## Example Usage
        ### create a proxy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        proxy1 = huaweicloud.gauss_db.MysqlProxy("proxy1",
            instance_id=instance_id,
            flavor="gaussdb.proxy.xlarge.arm.2",
            node_num=3)
        ```

        ## Import

        GaussDB instance can be imported using the instance `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:GaussDB/mysqlProxy:MysqlProxy proxy_1 ee678f40-ce8e-4d0c-8221-38dead426f06
        ```

        :param str resource_name: The name of the resource.
        :param MysqlProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MysqlProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MysqlProxyArgs.__new__(MysqlProxyArgs)

            if flavor is None and not opts.urn:
                raise TypeError("Missing required property 'flavor'")
            __props__.__dict__["flavor"] = flavor
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if node_num is None and not opts.urn:
                raise TypeError("Missing required property 'node_num'")
            __props__.__dict__["node_num"] = node_num
            __props__.__dict__["region"] = region
            __props__.__dict__["address"] = None
            __props__.__dict__["port"] = None
        super(MysqlProxy, __self__).__init__(
            'huaweicloud:GaussDB/mysqlProxy:MysqlProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            flavor: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            node_num: Optional[pulumi.Input[int]] = None,
            port: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'MysqlProxy':
        """
        Get an existing MysqlProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Indicates the address of the proxy.
        :param pulumi.Input[str] flavor: Specifies the flavor of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the instance ID of the proxy.
               Changing this parameter will create a new resource.
        :param pulumi.Input[int] node_num: Specifies the node count of the proxy.
        :param pulumi.Input[int] port: Indicates the port of the proxy.
        :param pulumi.Input[str] region: The region in which to create the GaussDB mysql proxy resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MysqlProxyState.__new__(_MysqlProxyState)

        __props__.__dict__["address"] = address
        __props__.__dict__["flavor"] = flavor
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["node_num"] = node_num
        __props__.__dict__["port"] = port
        __props__.__dict__["region"] = region
        return MysqlProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Indicates the address of the proxy.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Output[str]:
        """
        Specifies the flavor of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the instance ID of the proxy.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Output[int]:
        """
        Specifies the node count of the proxy.
        """
        return pulumi.get(self, "node_num")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Indicates the port of the proxy.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the GaussDB mysql proxy resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

