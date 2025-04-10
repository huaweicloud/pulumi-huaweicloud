# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VipAssociateArgs', 'VipAssociate']

@pulumi.input_type
class VipAssociateArgs:
    def __init__(__self__, *,
                 port_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vip_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VipAssociate resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] port_ids: An array of one or more IDs of the ports to attach the vip to.
        :param pulumi.Input[str] vip_id: The ID of vip to attach the ports to.
        :param pulumi.Input[str] region: The region in which to create the vip associate resource. If omitted, the
               provider-level region will be used.
        """
        pulumi.set(__self__, "port_ids", port_ids)
        pulumi.set(__self__, "vip_id", vip_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="portIds")
    def port_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of one or more IDs of the ports to attach the vip to.
        """
        return pulumi.get(self, "port_ids")

    @port_ids.setter
    def port_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "port_ids", value)

    @property
    @pulumi.getter(name="vipId")
    def vip_id(self) -> pulumi.Input[str]:
        """
        The ID of vip to attach the ports to.
        """
        return pulumi.get(self, "vip_id")

    @vip_id.setter
    def vip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vip_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the vip associate resource. If omitted, the
        provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _VipAssociateState:
    def __init__(__self__, *,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 port_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vip_id: Optional[pulumi.Input[str]] = None,
                 vip_ip_address: Optional[pulumi.Input[str]] = None,
                 vip_subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VipAssociate resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: The IP addresses of ports to attach the vip to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] port_ids: An array of one or more IDs of the ports to attach the vip to.
        :param pulumi.Input[str] region: The region in which to create the vip associate resource. If omitted, the
               provider-level region will be used.
        :param pulumi.Input[str] vip_id: The ID of vip to attach the ports to.
        :param pulumi.Input[str] vip_ip_address: The IP address in the subnet for this vip.
        :param pulumi.Input[str] vip_subnet_id: The ID of the subnet this vip connects to.
        """
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if port_ids is not None:
            pulumi.set(__self__, "port_ids", port_ids)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if vip_id is not None:
            pulumi.set(__self__, "vip_id", vip_id)
        if vip_ip_address is not None:
            pulumi.set(__self__, "vip_ip_address", vip_ip_address)
        if vip_subnet_id is not None:
            pulumi.set(__self__, "vip_subnet_id", vip_subnet_id)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IP addresses of ports to attach the vip to.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="portIds")
    def port_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of one or more IDs of the ports to attach the vip to.
        """
        return pulumi.get(self, "port_ids")

    @port_ids.setter
    def port_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "port_ids", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the vip associate resource. If omitted, the
        provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="vipId")
    def vip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of vip to attach the ports to.
        """
        return pulumi.get(self, "vip_id")

    @vip_id.setter
    def vip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_id", value)

    @property
    @pulumi.getter(name="vipIpAddress")
    def vip_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address in the subnet for this vip.
        """
        return pulumi.get(self, "vip_ip_address")

    @vip_ip_address.setter
    def vip_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_ip_address", value)

    @property
    @pulumi.getter(name="vipSubnetId")
    def vip_subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the subnet this vip connects to.
        """
        return pulumi.get(self, "vip_subnet_id")

    @vip_subnet_id.setter
    def vip_subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_subnet_id", value)


class VipAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 port_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vip_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Using this resource, one or more NICs (to which the ECS instance belongs) can be bound to the VIP.

        > A VIP can only have one resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        vip_id = config.require_object("vipId")
        nic_port_ids = config.require_object("nicPortIds")
        vip_associated = huaweicloud.vpc.VipAssociate("vipAssociated",
            vip_id=vip_id,
            port_ids=nic_port_ids)
        ```

        ## Import

        Vip associate can be imported using the `vip_id` and port IDs separated by slashes (no limit on the number of port IDs), e.g. bash

        ```sh
         $ pulumi import huaweicloud:Vpc/vipAssociate:VipAssociate vip_associated vip_id/port1_id/port2_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] port_ids: An array of one or more IDs of the ports to attach the vip to.
        :param pulumi.Input[str] region: The region in which to create the vip associate resource. If omitted, the
               provider-level region will be used.
        :param pulumi.Input[str] vip_id: The ID of vip to attach the ports to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VipAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Using this resource, one or more NICs (to which the ECS instance belongs) can be bound to the VIP.

        > A VIP can only have one resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        vip_id = config.require_object("vipId")
        nic_port_ids = config.require_object("nicPortIds")
        vip_associated = huaweicloud.vpc.VipAssociate("vipAssociated",
            vip_id=vip_id,
            port_ids=nic_port_ids)
        ```

        ## Import

        Vip associate can be imported using the `vip_id` and port IDs separated by slashes (no limit on the number of port IDs), e.g. bash

        ```sh
         $ pulumi import huaweicloud:Vpc/vipAssociate:VipAssociate vip_associated vip_id/port1_id/port2_id
        ```

        :param str resource_name: The name of the resource.
        :param VipAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VipAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 port_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 vip_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VipAssociateArgs.__new__(VipAssociateArgs)

            if port_ids is None and not opts.urn:
                raise TypeError("Missing required property 'port_ids'")
            __props__.__dict__["port_ids"] = port_ids
            __props__.__dict__["region"] = region
            if vip_id is None and not opts.urn:
                raise TypeError("Missing required property 'vip_id'")
            __props__.__dict__["vip_id"] = vip_id
            __props__.__dict__["ip_addresses"] = None
            __props__.__dict__["vip_ip_address"] = None
            __props__.__dict__["vip_subnet_id"] = None
        super(VipAssociate, __self__).__init__(
            'huaweicloud:Vpc/vipAssociate:VipAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            port_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            vip_id: Optional[pulumi.Input[str]] = None,
            vip_ip_address: Optional[pulumi.Input[str]] = None,
            vip_subnet_id: Optional[pulumi.Input[str]] = None) -> 'VipAssociate':
        """
        Get an existing VipAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: The IP addresses of ports to attach the vip to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] port_ids: An array of one or more IDs of the ports to attach the vip to.
        :param pulumi.Input[str] region: The region in which to create the vip associate resource. If omitted, the
               provider-level region will be used.
        :param pulumi.Input[str] vip_id: The ID of vip to attach the ports to.
        :param pulumi.Input[str] vip_ip_address: The IP address in the subnet for this vip.
        :param pulumi.Input[str] vip_subnet_id: The ID of the subnet this vip connects to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VipAssociateState.__new__(_VipAssociateState)

        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["port_ids"] = port_ids
        __props__.__dict__["region"] = region
        __props__.__dict__["vip_id"] = vip_id
        __props__.__dict__["vip_ip_address"] = vip_ip_address
        __props__.__dict__["vip_subnet_id"] = vip_subnet_id
        return VipAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The IP addresses of ports to attach the vip to.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="portIds")
    def port_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of one or more IDs of the ports to attach the vip to.
        """
        return pulumi.get(self, "port_ids")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the vip associate resource. If omitted, the
        provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="vipId")
    def vip_id(self) -> pulumi.Output[str]:
        """
        The ID of vip to attach the ports to.
        """
        return pulumi.get(self, "vip_id")

    @property
    @pulumi.getter(name="vipIpAddress")
    def vip_ip_address(self) -> pulumi.Output[str]:
        """
        The IP address in the subnet for this vip.
        """
        return pulumi.get(self, "vip_ip_address")

    @property
    @pulumi.getter(name="vipSubnetId")
    def vip_subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet this vip connects to.
        """
        return pulumi.get(self, "vip_subnet_id")

