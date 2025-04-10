# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnatRuleArgs', 'SnatRule']

@pulumi.input_type
class SnatRuleArgs:
    def __init__(__self__, *,
                 nat_gateway_id: pulumi.Input[str],
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 global_eip_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SnatRule resource.
        :param pulumi.Input[str] nat_gateway_id: Specifies the ID of the gateway to which the SNAT rule belongs.  
               Changing this will create a new resource.
        :param pulumi.Input[str] cidr: Specifies the CIDR block connected by SNAT rule (DC side).  
               This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the SNAT rule.
               The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        :param pulumi.Input[str] floating_ip_id: Specifies the IDs of floating IPs connected by SNAT rule.  
               Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        :param pulumi.Input[str] global_eip_id: Specifies the IDs of global EIPs connected by SNAT rule.  
               Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        :param pulumi.Input[str] region: Specifies the region where the SNAT rule is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[int] source_type: Specifies the resource scenario.  
               The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
               Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        :param pulumi.Input[str] subnet_id: Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
               This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if floating_ip_id is not None:
            pulumi.set(__self__, "floating_ip_id", floating_ip_id)
        if global_eip_id is not None:
            pulumi.set(__self__, "global_eip_id", global_eip_id)
        if network_id is not None:
            warnings.warn("""schema: Deprecated; Use 'subnet_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""network_id is deprecated: schema: Deprecated; Use 'subnet_id' instead.""")
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the gateway to which the SNAT rule belongs.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the CIDR block connected by SNAT rule (DC side).  
        This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the SNAT rule.
        The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="floatingIpId")
    def floating_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IDs of floating IPs connected by SNAT rule.  
        Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "floating_ip_id")

    @floating_ip_id.setter
    def floating_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip_id", value)

    @property
    @pulumi.getter(name="globalEipId")
    def global_eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IDs of global EIPs connected by SNAT rule.  
        Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "global_eip_id")

    @global_eip_id.setter
    def global_eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_eip_id", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the SNAT rule is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the resource scenario.  
        The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
        Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
        This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class _SnatRuleState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 floating_ip_address: Optional[pulumi.Input[str]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 freezed_ip_address: Optional[pulumi.Input[str]] = None,
                 global_eip_address: Optional[pulumi.Input[str]] = None,
                 global_eip_id: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnatRule resources.
        :param pulumi.Input[str] cidr: Specifies the CIDR block connected by SNAT rule (DC side).  
               This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        :param pulumi.Input[str] created_at: The creation time of the SNAT rule.
        :param pulumi.Input[str] description: Specifies the description of the SNAT rule.
               The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        :param pulumi.Input[str] floating_ip_address: The actual floating IP address.
        :param pulumi.Input[str] floating_ip_id: Specifies the IDs of floating IPs connected by SNAT rule.  
               Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        :param pulumi.Input[str] freezed_ip_address: The frozen EIP associated with the SNAT rule.
        :param pulumi.Input[str] global_eip_address: The global EIP addresses (separated by commas) connected by SNAT rule.
        :param pulumi.Input[str] global_eip_id: Specifies the IDs of global EIPs connected by SNAT rule.  
               Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        :param pulumi.Input[str] nat_gateway_id: Specifies the ID of the gateway to which the SNAT rule belongs.  
               Changing this will create a new resource.
        :param pulumi.Input[str] region: Specifies the region where the SNAT rule is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[int] source_type: Specifies the resource scenario.  
               The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
               Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        :param pulumi.Input[str] status: The status of the SNAT rule.
        :param pulumi.Input[str] subnet_id: Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
               This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if floating_ip_address is not None:
            pulumi.set(__self__, "floating_ip_address", floating_ip_address)
        if floating_ip_id is not None:
            pulumi.set(__self__, "floating_ip_id", floating_ip_id)
        if freezed_ip_address is not None:
            pulumi.set(__self__, "freezed_ip_address", freezed_ip_address)
        if global_eip_address is not None:
            pulumi.set(__self__, "global_eip_address", global_eip_address)
        if global_eip_id is not None:
            pulumi.set(__self__, "global_eip_id", global_eip_id)
        if nat_gateway_id is not None:
            pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if network_id is not None:
            warnings.warn("""schema: Deprecated; Use 'subnet_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""network_id is deprecated: schema: Deprecated; Use 'subnet_id' instead.""")
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the CIDR block connected by SNAT rule (DC side).  
        This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the SNAT rule.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the SNAT rule.
        The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
        Specifies the IDs of floating IPs connected by SNAT rule.  
        Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "floating_ip_id")

    @floating_ip_id.setter
    def floating_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip_id", value)

    @property
    @pulumi.getter(name="freezedIpAddress")
    def freezed_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The frozen EIP associated with the SNAT rule.
        """
        return pulumi.get(self, "freezed_ip_address")

    @freezed_ip_address.setter
    def freezed_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "freezed_ip_address", value)

    @property
    @pulumi.getter(name="globalEipAddress")
    def global_eip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The global EIP addresses (separated by commas) connected by SNAT rule.
        """
        return pulumi.get(self, "global_eip_address")

    @global_eip_address.setter
    def global_eip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_eip_address", value)

    @property
    @pulumi.getter(name="globalEipId")
    def global_eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IDs of global EIPs connected by SNAT rule.  
        Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "global_eip_id")

    @global_eip_id.setter
    def global_eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_eip_id", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the gateway to which the SNAT rule belongs.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the SNAT rule is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the resource scenario.  
        The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
        Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the SNAT rule.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
        This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class SnatRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 global_eip_id: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an SNAT rule resource of the **public** NAT within HuaweiCloud.

        ## Example Usage

        ## Import

        The SNAT rule can be imported using `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Nat/snatRule:SnatRule test <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: Specifies the CIDR block connected by SNAT rule (DC side).  
               This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the SNAT rule.
               The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        :param pulumi.Input[str] floating_ip_id: Specifies the IDs of floating IPs connected by SNAT rule.  
               Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        :param pulumi.Input[str] global_eip_id: Specifies the IDs of global EIPs connected by SNAT rule.  
               Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        :param pulumi.Input[str] nat_gateway_id: Specifies the ID of the gateway to which the SNAT rule belongs.  
               Changing this will create a new resource.
        :param pulumi.Input[str] region: Specifies the region where the SNAT rule is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[int] source_type: Specifies the resource scenario.  
               The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
               Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        :param pulumi.Input[str] subnet_id: Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
               This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnatRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an SNAT rule resource of the **public** NAT within HuaweiCloud.

        ## Example Usage

        ## Import

        The SNAT rule can be imported using `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Nat/snatRule:SnatRule test <id>
        ```

        :param str resource_name: The name of the resource.
        :param SnatRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnatRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 floating_ip_id: Optional[pulumi.Input[str]] = None,
                 global_eip_id: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnatRuleArgs.__new__(SnatRuleArgs)

            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["description"] = description
            __props__.__dict__["floating_ip_id"] = floating_ip_id
            __props__.__dict__["global_eip_id"] = global_eip_id
            if nat_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'nat_gateway_id'")
            __props__.__dict__["nat_gateway_id"] = nat_gateway_id
            if network_id is not None and not opts.urn:
                warnings.warn("""schema: Deprecated; Use 'subnet_id' instead.""", DeprecationWarning)
                pulumi.log.warn("""network_id is deprecated: schema: Deprecated; Use 'subnet_id' instead.""")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["region"] = region
            __props__.__dict__["source_type"] = source_type
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["floating_ip_address"] = None
            __props__.__dict__["freezed_ip_address"] = None
            __props__.__dict__["global_eip_address"] = None
            __props__.__dict__["status"] = None
        super(SnatRule, __self__).__init__(
            'huaweicloud:Nat/snatRule:SnatRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            floating_ip_address: Optional[pulumi.Input[str]] = None,
            floating_ip_id: Optional[pulumi.Input[str]] = None,
            freezed_ip_address: Optional[pulumi.Input[str]] = None,
            global_eip_address: Optional[pulumi.Input[str]] = None,
            global_eip_id: Optional[pulumi.Input[str]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            source_type: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'SnatRule':
        """
        Get an existing SnatRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: Specifies the CIDR block connected by SNAT rule (DC side).  
               This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        :param pulumi.Input[str] created_at: The creation time of the SNAT rule.
        :param pulumi.Input[str] description: Specifies the description of the SNAT rule.
               The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        :param pulumi.Input[str] floating_ip_address: The actual floating IP address.
        :param pulumi.Input[str] floating_ip_id: Specifies the IDs of floating IPs connected by SNAT rule.  
               Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        :param pulumi.Input[str] freezed_ip_address: The frozen EIP associated with the SNAT rule.
        :param pulumi.Input[str] global_eip_address: The global EIP addresses (separated by commas) connected by SNAT rule.
        :param pulumi.Input[str] global_eip_id: Specifies the IDs of global EIPs connected by SNAT rule.  
               Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        :param pulumi.Input[str] nat_gateway_id: Specifies the ID of the gateway to which the SNAT rule belongs.  
               Changing this will create a new resource.
        :param pulumi.Input[str] region: Specifies the region where the SNAT rule is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[int] source_type: Specifies the resource scenario.  
               The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
               Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        :param pulumi.Input[str] status: The status of the SNAT rule.
        :param pulumi.Input[str] subnet_id: Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
               This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnatRuleState.__new__(_SnatRuleState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["floating_ip_address"] = floating_ip_address
        __props__.__dict__["floating_ip_id"] = floating_ip_id
        __props__.__dict__["freezed_ip_address"] = freezed_ip_address
        __props__.__dict__["global_eip_address"] = global_eip_address
        __props__.__dict__["global_eip_id"] = global_eip_id
        __props__.__dict__["nat_gateway_id"] = nat_gateway_id
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["region"] = region
        __props__.__dict__["source_type"] = source_type
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        return SnatRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the CIDR block connected by SNAT rule (DC side).  
        This parameter and `subnet_id` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time of the SNAT rule.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description of the SNAT rule.
        The value is a string of no more than `255` characters, and angle brackets (<>) are not allowed.
        """
        return pulumi.get(self, "description")

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
        Specifies the IDs of floating IPs connected by SNAT rule.  
        Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "floating_ip_id")

    @property
    @pulumi.getter(name="freezedIpAddress")
    def freezed_ip_address(self) -> pulumi.Output[str]:
        """
        The frozen EIP associated with the SNAT rule.
        """
        return pulumi.get(self, "freezed_ip_address")

    @property
    @pulumi.getter(name="globalEipAddress")
    def global_eip_address(self) -> pulumi.Output[str]:
        """
        The global EIP addresses (separated by commas) connected by SNAT rule.
        """
        return pulumi.get(self, "global_eip_address")

    @property
    @pulumi.getter(name="globalEipId")
    def global_eip_id(self) -> pulumi.Output[str]:
        """
        Specifies the IDs of global EIPs connected by SNAT rule.  
        Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed `20`.
        """
        return pulumi.get(self, "global_eip_id")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the gateway to which the SNAT rule belongs.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region where the SNAT rule is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the resource scenario.  
        The valid values are `0` (VPC scenario) and `1` (Direct Connect scenario), and the default value is `0`.
        Only `cidr` can be specified over a Direct Connect connection. Changing this will create a new resource.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the SNAT rule.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Specifies the network IDs of subnet connected by SNAT rule (VPC side).  
        This parameter and `cidr` are alternative. Changing this will create a new resource.
        """
        return pulumi.get(self, "subnet_id")

