# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DnatRuleArgs', 'DnatRule']

@pulumi.input_type
class DnatRuleArgs:
    def __init__(__self__, *,
                 external_service_port: pulumi.Input[int],
                 floating_ip_id: pulumi.Input[str],
                 internal_service_port: pulumi.Input[int],
                 nat_gateway_id: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DnatRule resource.
        :param pulumi.Input[int] external_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] floating_ip_id: Specifies the ID of the floating IP address. Changing this creates a
               new dnat rule.
        :param pulumi.Input[int] internal_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] nat_gateway_id: ID of the nat gateway this dnat rule belongs to. Changing this creates
               a new dnat rule.
        :param pulumi.Input[str] protocol: Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
               Changing this creates a new dnat rule.
        :param pulumi.Input[str] description: Specifies the description of the dnat rule.
               The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
               Changing this creates a new dnat rule.
        :param pulumi.Input[str] port_id: Specifies the port ID of network. This parameter is mandatory in VPC
               scenario. Use Vpc.Port to get the port if just know a fixed IP
               addresses on the port. Changing this creates a new dnat rule.
        :param pulumi.Input[str] private_ip: Specifies the private IP address of a user. This parameter is mandatory in
               Direct Connect scenario. Changing this creates a new dnat rule.
        :param pulumi.Input[str] region: The region in which to create the dnat rule resource. If omitted, the
               provider-level region will be used. Changing this creates a new dnat rule.
        """
        pulumi.set(__self__, "external_service_port", external_service_port)
        pulumi.set(__self__, "floating_ip_id", floating_ip_id)
        pulumi.set(__self__, "internal_service_port", internal_service_port)
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        pulumi.set(__self__, "protocol", protocol)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="externalServicePort")
    def external_service_port(self) -> pulumi.Input[int]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "external_service_port")

    @external_service_port.setter
    def external_service_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "external_service_port", value)

    @property
    @pulumi.getter(name="floatingIpId")
    def floating_ip_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the floating IP address. Changing this creates a
        new dnat rule.
        """
        return pulumi.get(self, "floating_ip_id")

    @floating_ip_id.setter
    def floating_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "floating_ip_id", value)

    @property
    @pulumi.getter(name="internalServicePort")
    def internal_service_port(self) -> pulumi.Input[int]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "internal_service_port")

    @internal_service_port.setter
    def internal_service_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "internal_service_port", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Input[str]:
        """
        ID of the nat gateway this dnat rule belongs to. Changing this creates
        a new dnat rule.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the dnat rule.
        The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the port ID of network. This parameter is mandatory in VPC
        scenario. Use Vpc.Port to get the port if just know a fixed IP
        addresses on the port. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the private IP address of a user. This parameter is mandatory in
        Direct Connect scenario. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the dnat rule resource. If omitted, the
        provider-level region will be used. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DnatRuleState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_service_port: Optional[pulumi.Input[int]] = None,
                 floating_ip_address: Optional[pulumi.Input[str]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 internal_service_port: Optional[pulumi.Input[int]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DnatRule resources.
        :param pulumi.Input[str] created_at: Dnat rule creation time.
        :param pulumi.Input[str] description: Specifies the description of the dnat rule.
               The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
               Changing this creates a new dnat rule.
        :param pulumi.Input[int] external_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] floating_ip_address: The actual floating IP address.
        :param pulumi.Input[str] floating_ip_id: Specifies the ID of the floating IP address. Changing this creates a
               new dnat rule.
        :param pulumi.Input[int] internal_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] nat_gateway_id: ID of the nat gateway this dnat rule belongs to. Changing this creates
               a new dnat rule.
        :param pulumi.Input[str] port_id: Specifies the port ID of network. This parameter is mandatory in VPC
               scenario. Use Vpc.Port to get the port if just know a fixed IP
               addresses on the port. Changing this creates a new dnat rule.
        :param pulumi.Input[str] private_ip: Specifies the private IP address of a user. This parameter is mandatory in
               Direct Connect scenario. Changing this creates a new dnat rule.
        :param pulumi.Input[str] protocol: Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
               Changing this creates a new dnat rule.
        :param pulumi.Input[str] region: The region in which to create the dnat rule resource. If omitted, the
               provider-level region will be used. Changing this creates a new dnat rule.
        :param pulumi.Input[str] status: Dnat rule status.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_service_port is not None:
            pulumi.set(__self__, "external_service_port", external_service_port)
        if floating_ip_address is not None:
            pulumi.set(__self__, "floating_ip_address", floating_ip_address)
        if floating_ip_id is not None:
            pulumi.set(__self__, "floating_ip_id", floating_ip_id)
        if internal_service_port is not None:
            pulumi.set(__self__, "internal_service_port", internal_service_port)
        if nat_gateway_id is not None:
            pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Dnat rule creation time.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the dnat rule.
        The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalServicePort")
    def external_service_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "external_service_port")

    @external_service_port.setter
    def external_service_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "external_service_port", value)

    @property
    @pulumi.getter(name="floatingIpAddress")
    def floating_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The actual floating IP address.
        """
        return pulumi.get(self, "floating_ip_address")

    @floating_ip_address.setter
    def floating_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip_address", value)

    @property
    @pulumi.getter(name="floatingIpId")
    def floating_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the floating IP address. Changing this creates a
        new dnat rule.
        """
        return pulumi.get(self, "floating_ip_id")

    @floating_ip_id.setter
    def floating_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip_id", value)

    @property
    @pulumi.getter(name="internalServicePort")
    def internal_service_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "internal_service_port")

    @internal_service_port.setter
    def internal_service_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internal_service_port", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the nat gateway this dnat rule belongs to. Changing this creates
        a new dnat rule.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the port ID of network. This parameter is mandatory in VPC
        scenario. Use Vpc.Port to get the port if just know a fixed IP
        addresses on the port. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the private IP address of a user. This parameter is mandatory in
        Direct Connect scenario. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the dnat rule resource. If omitted, the
        provider-level region will be used. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Dnat rule status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class DnatRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_service_port: Optional[pulumi.Input[int]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 internal_service_port: Optional[pulumi.Input[int]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a DNAT rule resource within HuaweiCloud.

        ## Example Usage
        ### DNAT rule in Direct Connect scenario

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        dnat2 = huaweicloud.nat.DnatRule("dnat2",
            nat_gateway_id=var["natgw_id"],
            floating_ip_id=var["publicip_id"],
            private_ip="10.0.0.12",
            protocol="tcp",
            internal_service_port=80,
            external_service_port=8080)
        ```

        ## Import

        DNAT rules can be imported using the following format

        ```sh
         $ pulumi import huaweicloud:Nat/dnatRule:DnatRule dnat_1 f4f783a7-b908-4215-b018-724960e5df4a
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of the dnat rule.
               The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
               Changing this creates a new dnat rule.
        :param pulumi.Input[int] external_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] floating_ip_id: Specifies the ID of the floating IP address. Changing this creates a
               new dnat rule.
        :param pulumi.Input[int] internal_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] nat_gateway_id: ID of the nat gateway this dnat rule belongs to. Changing this creates
               a new dnat rule.
        :param pulumi.Input[str] port_id: Specifies the port ID of network. This parameter is mandatory in VPC
               scenario. Use Vpc.Port to get the port if just know a fixed IP
               addresses on the port. Changing this creates a new dnat rule.
        :param pulumi.Input[str] private_ip: Specifies the private IP address of a user. This parameter is mandatory in
               Direct Connect scenario. Changing this creates a new dnat rule.
        :param pulumi.Input[str] protocol: Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
               Changing this creates a new dnat rule.
        :param pulumi.Input[str] region: The region in which to create the dnat rule resource. If omitted, the
               provider-level region will be used. Changing this creates a new dnat rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnatRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a DNAT rule resource within HuaweiCloud.

        ## Example Usage
        ### DNAT rule in Direct Connect scenario

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        dnat2 = huaweicloud.nat.DnatRule("dnat2",
            nat_gateway_id=var["natgw_id"],
            floating_ip_id=var["publicip_id"],
            private_ip="10.0.0.12",
            protocol="tcp",
            internal_service_port=80,
            external_service_port=8080)
        ```

        ## Import

        DNAT rules can be imported using the following format

        ```sh
         $ pulumi import huaweicloud:Nat/dnatRule:DnatRule dnat_1 f4f783a7-b908-4215-b018-724960e5df4a
        ```

        :param str resource_name: The name of the resource.
        :param DnatRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnatRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_service_port: Optional[pulumi.Input[int]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 internal_service_port: Optional[pulumi.Input[int]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DnatRuleArgs.__new__(DnatRuleArgs)

            __props__.__dict__["description"] = description
            if external_service_port is None and not opts.urn:
                raise TypeError("Missing required property 'external_service_port'")
            __props__.__dict__["external_service_port"] = external_service_port
            if floating_ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'floating_ip_id'")
            __props__.__dict__["floating_ip_id"] = floating_ip_id
            if internal_service_port is None and not opts.urn:
                raise TypeError("Missing required property 'internal_service_port'")
            __props__.__dict__["internal_service_port"] = internal_service_port
            if nat_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'nat_gateway_id'")
            __props__.__dict__["nat_gateway_id"] = nat_gateway_id
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["private_ip"] = private_ip
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["created_at"] = None
            __props__.__dict__["floating_ip_address"] = None
            __props__.__dict__["status"] = None
        super(DnatRule, __self__).__init__(
            'huaweicloud:Nat/dnatRule:DnatRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_service_port: Optional[pulumi.Input[int]] = None,
            floating_ip_address: Optional[pulumi.Input[str]] = None,
            floating_ip_id: Optional[pulumi.Input[str]] = None,
            internal_service_port: Optional[pulumi.Input[int]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'DnatRule':
        """
        Get an existing DnatRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Dnat rule creation time.
        :param pulumi.Input[str] description: Specifies the description of the dnat rule.
               The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
               Changing this creates a new dnat rule.
        :param pulumi.Input[int] external_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] floating_ip_address: The actual floating IP address.
        :param pulumi.Input[str] floating_ip_id: Specifies the ID of the floating IP address. Changing this creates a
               new dnat rule.
        :param pulumi.Input[int] internal_service_port: Specifies port used by ECSs or BMSs to provide services for
               external systems. Changing this creates a new dnat rule.
        :param pulumi.Input[str] nat_gateway_id: ID of the nat gateway this dnat rule belongs to. Changing this creates
               a new dnat rule.
        :param pulumi.Input[str] port_id: Specifies the port ID of network. This parameter is mandatory in VPC
               scenario. Use Vpc.Port to get the port if just know a fixed IP
               addresses on the port. Changing this creates a new dnat rule.
        :param pulumi.Input[str] private_ip: Specifies the private IP address of a user. This parameter is mandatory in
               Direct Connect scenario. Changing this creates a new dnat rule.
        :param pulumi.Input[str] protocol: Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
               Changing this creates a new dnat rule.
        :param pulumi.Input[str] region: The region in which to create the dnat rule resource. If omitted, the
               provider-level region will be used. Changing this creates a new dnat rule.
        :param pulumi.Input[str] status: Dnat rule status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnatRuleState.__new__(_DnatRuleState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["external_service_port"] = external_service_port
        __props__.__dict__["floating_ip_address"] = floating_ip_address
        __props__.__dict__["floating_ip_id"] = floating_ip_id
        __props__.__dict__["internal_service_port"] = internal_service_port
        __props__.__dict__["nat_gateway_id"] = nat_gateway_id
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["private_ip"] = private_ip
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        return DnatRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Dnat rule creation time.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Specifies the description of the dnat rule.
        The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalServicePort")
    def external_service_port(self) -> pulumi.Output[int]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "external_service_port")

    @property
    @pulumi.getter(name="floatingIpAddress")
    def floating_ip_address(self) -> pulumi.Output[str]:
        """
        The actual floating IP address.
        """
        return pulumi.get(self, "floating_ip_address")

    @property
    @pulumi.getter(name="floatingIpId")
    def floating_ip_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the floating IP address. Changing this creates a
        new dnat rule.
        """
        return pulumi.get(self, "floating_ip_id")

    @property
    @pulumi.getter(name="internalServicePort")
    def internal_service_port(self) -> pulumi.Output[int]:
        """
        Specifies port used by ECSs or BMSs to provide services for
        external systems. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "internal_service_port")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[str]:
        """
        ID of the nat gateway this dnat rule belongs to. Changing this creates
        a new dnat rule.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the port ID of network. This parameter is mandatory in VPC
        scenario. Use Vpc.Port to get the port if just know a fixed IP
        addresses on the port. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the private IP address of a user. This parameter is mandatory in
        Direct Connect scenario. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
        Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the dnat rule resource. If omitted, the
        provider-level region will be used. Changing this creates a new dnat rule.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Dnat rule status.
        """
        return pulumi.get(self, "status")

