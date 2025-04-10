# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InterfaceAttachArgs', 'InterfaceAttach']

@pulumi.input_type
class InterfaceAttachArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 ipv6_bandwidth_id: Optional[pulumi.Input[str]] = None,
                 ipv6_enable: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_dest_check: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a InterfaceAttach resource.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] fixed_ip: An IP address to associate with the port.
        :param pulumi.Input[str] ipv6_bandwidth_id: Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        :param pulumi.Input[bool] ipv6_enable: Specifies if the NIC supporting IPv6 or not.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created
               automatically.
               This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the network interface attache resource. If
               omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies the list of security group IDs bound to the specified port.  
               Defaults to the default security group.
        :param pulumi.Input[bool] source_dest_check: Specifies whether the ECS processes only traffic that is destined specifically
               for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
               virtual IP address bound to it.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if ipv6_bandwidth_id is not None:
            pulumi.set(__self__, "ipv6_bandwidth_id", ipv6_bandwidth_id)
        if ipv6_enable is not None:
            pulumi.set(__self__, "ipv6_enable", ipv6_enable)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if source_dest_check is not None:
            pulumi.set(__self__, "source_dest_check", source_dest_check)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address to associate with the port.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="ipv6BandwidthId")
    def ipv6_bandwidth_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        """
        return pulumi.get(self, "ipv6_bandwidth_id")

    @ipv6_bandwidth_id.setter
    def ipv6_bandwidth_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_bandwidth_id", value)

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the NIC supporting IPv6 or not.
        """
        return pulumi.get(self, "ipv6_enable")

    @ipv6_enable.setter
    def ipv6_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ipv6_enable", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Network to attach to an Instance. A port will be created
        automatically.
        This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Port to attach to an Instance.
        This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the network interface attache resource. If
        omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of security group IDs bound to the specified port.  
        Defaults to the default security group.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="sourceDestCheck")
    def source_dest_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the ECS processes only traffic that is destined specifically
        for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
        virtual IP address bound to it.
        """
        return pulumi.get(self, "source_dest_check")

    @source_dest_check.setter
    def source_dest_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "source_dest_check", value)


@pulumi.input_type
class _InterfaceAttachState:
    def __init__(__self__, *,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 fixed_ipv6: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ipv6_bandwidth_id: Optional[pulumi.Input[str]] = None,
                 ipv6_enable: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_dest_check: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering InterfaceAttach resources.
        :param pulumi.Input[str] fixed_ip: An IP address to associate with the port.
        :param pulumi.Input[str] fixed_ipv6: The IPv6 address.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] ipv6_bandwidth_id: Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        :param pulumi.Input[bool] ipv6_enable: Specifies if the NIC supporting IPv6 or not.
        :param pulumi.Input[str] mac: The MAC address of the NIC.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created
               automatically.
               This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the network interface attache resource. If
               omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies the list of security group IDs bound to the specified port.  
               Defaults to the default security group.
        :param pulumi.Input[bool] source_dest_check: Specifies whether the ECS processes only traffic that is destined specifically
               for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
               virtual IP address bound to it.
        """
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if fixed_ipv6 is not None:
            pulumi.set(__self__, "fixed_ipv6", fixed_ipv6)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if ipv6_bandwidth_id is not None:
            pulumi.set(__self__, "ipv6_bandwidth_id", ipv6_bandwidth_id)
        if ipv6_enable is not None:
            pulumi.set(__self__, "ipv6_enable", ipv6_enable)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if source_dest_check is not None:
            pulumi.set(__self__, "source_dest_check", source_dest_check)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address to associate with the port.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="fixedIpv6")
    def fixed_ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 address.
        """
        return pulumi.get(self, "fixed_ipv6")

    @fixed_ipv6.setter
    def fixed_ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ipv6", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="ipv6BandwidthId")
    def ipv6_bandwidth_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        """
        return pulumi.get(self, "ipv6_bandwidth_id")

    @ipv6_bandwidth_id.setter
    def ipv6_bandwidth_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_bandwidth_id", value)

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the NIC supporting IPv6 or not.
        """
        return pulumi.get(self, "ipv6_enable")

    @ipv6_enable.setter
    def ipv6_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ipv6_enable", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC address of the NIC.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Network to attach to an Instance. A port will be created
        automatically.
        This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Port to attach to an Instance.
        This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the network interface attache resource. If
        omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of security group IDs bound to the specified port.  
        Defaults to the default security group.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="sourceDestCheck")
    def source_dest_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the ECS processes only traffic that is destined specifically
        for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
        virtual IP address bound to it.
        """
        return pulumi.get(self, "source_dest_check")

    @source_dest_check.setter
    def source_dest_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "source_dest_check", value)


class InterfaceAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ipv6_bandwidth_id: Optional[pulumi.Input[str]] = None,
                 ipv6_enable: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_dest_check: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Attaches a Network Interface to an Instance.

        ## Example Usage
        ### Attach a port (under the specified network) to the ECS instance and generate a random IP address

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        network_id = config.require_object("networkId")
        test = huaweicloud.ecs.InterfaceAttach("test",
            instance_id=instance_id,
            network_id=network_id)
        ```
        ### Attach a port (under the specified network) to the ECS instance and use the custom security groups

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        network_id = config.require_object("networkId")
        security_group_ids = config.require_object("securityGroupIds")
        test = huaweicloud.ecs.InterfaceAttach("test",
            instance_id=instance_id,
            network_id=network_id,
            security_group_ids=security_group_ids)
        ```
        ### Attach a custom port to the ECS instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        mynet = huaweicloud.Vpc.get_subnet(name="subnet-default")
        myport = huaweicloud.Vpc.get_port(network_id=mynet.id,
            fixed_ip="192.168.0.100")
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a",
            networks=[huaweicloud.ecs.InstanceNetworkArgs(
                uuid="55534eaa-533a-419d-9b40-ec427ea7195a",
            )])
        attached = huaweicloud.ecs.InterfaceAttach("attached",
            instance_id=myinstance.id,
            port_id=myport.id)
        ```

        ## Import

        Interface Attachments can be imported using the Instance ID and Port ID separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:Ecs/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: An IP address to associate with the port.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] ipv6_bandwidth_id: Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        :param pulumi.Input[bool] ipv6_enable: Specifies if the NIC supporting IPv6 or not.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created
               automatically.
               This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the network interface attache resource. If
               omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies the list of security group IDs bound to the specified port.  
               Defaults to the default security group.
        :param pulumi.Input[bool] source_dest_check: Specifies whether the ECS processes only traffic that is destined specifically
               for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
               virtual IP address bound to it.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InterfaceAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches a Network Interface to an Instance.

        ## Example Usage
        ### Attach a port (under the specified network) to the ECS instance and generate a random IP address

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        network_id = config.require_object("networkId")
        test = huaweicloud.ecs.InterfaceAttach("test",
            instance_id=instance_id,
            network_id=network_id)
        ```
        ### Attach a port (under the specified network) to the ECS instance and use the custom security groups

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        network_id = config.require_object("networkId")
        security_group_ids = config.require_object("securityGroupIds")
        test = huaweicloud.ecs.InterfaceAttach("test",
            instance_id=instance_id,
            network_id=network_id,
            security_group_ids=security_group_ids)
        ```
        ### Attach a custom port to the ECS instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        mynet = huaweicloud.Vpc.get_subnet(name="subnet-default")
        myport = huaweicloud.Vpc.get_port(network_id=mynet.id,
            fixed_ip="192.168.0.100")
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a",
            networks=[huaweicloud.ecs.InstanceNetworkArgs(
                uuid="55534eaa-533a-419d-9b40-ec427ea7195a",
            )])
        attached = huaweicloud.ecs.InterfaceAttach("attached",
            instance_id=myinstance.id,
            port_id=myport.id)
        ```

        ## Import

        Interface Attachments can be imported using the Instance ID and Port ID separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:Ecs/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param InterfaceAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InterfaceAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ipv6_bandwidth_id: Optional[pulumi.Input[str]] = None,
                 ipv6_enable: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_dest_check: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InterfaceAttachArgs.__new__(InterfaceAttachArgs)

            __props__.__dict__["fixed_ip"] = fixed_ip
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["ipv6_bandwidth_id"] = ipv6_bandwidth_id
            __props__.__dict__["ipv6_enable"] = ipv6_enable
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["region"] = region
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["source_dest_check"] = source_dest_check
            __props__.__dict__["fixed_ipv6"] = None
            __props__.__dict__["mac"] = None
        super(InterfaceAttach, __self__).__init__(
            'huaweicloud:Ecs/interfaceAttach:InterfaceAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fixed_ip: Optional[pulumi.Input[str]] = None,
            fixed_ipv6: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ipv6_bandwidth_id: Optional[pulumi.Input[str]] = None,
            ipv6_enable: Optional[pulumi.Input[bool]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            source_dest_check: Optional[pulumi.Input[bool]] = None) -> 'InterfaceAttach':
        """
        Get an existing InterfaceAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: An IP address to associate with the port.
        :param pulumi.Input[str] fixed_ipv6: The IPv6 address.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] ipv6_bandwidth_id: Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        :param pulumi.Input[bool] ipv6_enable: Specifies if the NIC supporting IPv6 or not.
        :param pulumi.Input[str] mac: The MAC address of the NIC.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created
               automatically.
               This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the network interface attache resource. If
               omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies the list of security group IDs bound to the specified port.  
               Defaults to the default security group.
        :param pulumi.Input[bool] source_dest_check: Specifies whether the ECS processes only traffic that is destined specifically
               for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
               virtual IP address bound to it.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InterfaceAttachState.__new__(_InterfaceAttachState)

        __props__.__dict__["fixed_ip"] = fixed_ip
        __props__.__dict__["fixed_ipv6"] = fixed_ipv6
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ipv6_bandwidth_id"] = ipv6_bandwidth_id
        __props__.__dict__["ipv6_enable"] = ipv6_enable
        __props__.__dict__["mac"] = mac
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["region"] = region
        __props__.__dict__["security_group_ids"] = security_group_ids
        __props__.__dict__["source_dest_check"] = source_dest_check
        return InterfaceAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[str]:
        """
        An IP address to associate with the port.
        """
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter(name="fixedIpv6")
    def fixed_ipv6(self) -> pulumi.Output[str]:
        """
        The IPv6 address.
        """
        return pulumi.get(self, "fixed_ipv6")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="ipv6BandwidthId")
    def ipv6_bandwidth_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the shared bandwidth ID to which the IPv6 NIC attaches.
        """
        return pulumi.get(self, "ipv6_bandwidth_id")

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> pulumi.Output[bool]:
        """
        Specifies if the NIC supporting IPv6 or not.
        """
        return pulumi.get(self, "ipv6_enable")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        The MAC address of the NIC.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The ID of the Network to attach to an Instance. A port will be created
        automatically.
        This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        """
        The ID of the Port to attach to an Instance.
        This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the network interface attache resource. If
        omitted, the provider-level region will be used. Changing this creates a new network interface attache resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the list of security group IDs bound to the specified port.  
        Defaults to the default security group.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="sourceDestCheck")
    def source_dest_check(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the ECS processes only traffic that is destined specifically
        for it. This function is enabled by default but should be disabled if the ECS functions as a SNAT server or has a
        virtual IP address bound to it.
        """
        return pulumi.get(self, "source_dest_check")

