# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AclArgs', 'Acl']

@pulumi.input_type
class AclArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]] = None):
        """
        The set of arguments for constructing a Acl resource.
        :param pulumi.Input[str] type: Specifies the ACL is created through the Console or API.
               Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        :param pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]] ip_cidrs: Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
               The `ip_cidrs` cannot repeat. The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]] ip_ranges: Specifies the IP address ranges from which console access or api access is allowed.
               The `ip_ranges` cannot repeat. The object structure is documented below.
        """
        pulumi.set(__self__, "type", type)
        if ip_cidrs is not None:
            pulumi.set(__self__, "ip_cidrs", ip_cidrs)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the ACL is created through the Console or API.
        Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]]:
        """
        Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
        The `ip_cidrs` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_cidrs")

    @ip_cidrs.setter
    def ip_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]]):
        pulumi.set(self, "ip_cidrs", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]]:
        """
        Specifies the IP address ranges from which console access or api access is allowed.
        The `ip_ranges` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)


@pulumi.input_type
class _AclState:
    def __init__(__self__, *,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Acl resources.
        :param pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]] ip_cidrs: Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
               The `ip_cidrs` cannot repeat. The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]] ip_ranges: Specifies the IP address ranges from which console access or api access is allowed.
               The `ip_ranges` cannot repeat. The object structure is documented below.
        :param pulumi.Input[str] type: Specifies the ACL is created through the Console or API.
               Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        if ip_cidrs is not None:
            pulumi.set(__self__, "ip_cidrs", ip_cidrs)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]]:
        """
        Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
        The `ip_cidrs` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_cidrs")

    @ip_cidrs.setter
    def ip_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpCidrArgs']]]]):
        pulumi.set(self, "ip_cidrs", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]]:
        """
        Specifies the IP address ranges from which console access or api access is allowed.
        The `ip_ranges` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ACL is created through the Console or API.
        Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Acl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpCidrArgs']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpRangeArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an ACL resource within HuaweiCloud IAM service. The ACL allowing user access only from specified IP address
        ranges and IPv4 CIDR blocks. The ACL take effect for IAM users under the Domain account rather than the account itself.

        > **NOTE:** You *must* have admin privileges to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        acl = huaweicloud.iam.Acl("acl",
            ip_cidrs=[huaweicloud.iam.AclIpCidrArgs(
                cidr="159.138.39.192/32",
                description="This is a test ip address",
            )],
            ip_ranges=[huaweicloud.iam.AclIpRangeArgs(
                description="This is a test ip range",
                range="0.0.0.0-255.255.255.0",
            )],
            type="console")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpCidrArgs']]]] ip_cidrs: Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
               The `ip_cidrs` cannot repeat. The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpRangeArgs']]]] ip_ranges: Specifies the IP address ranges from which console access or api access is allowed.
               The `ip_ranges` cannot repeat. The object structure is documented below.
        :param pulumi.Input[str] type: Specifies the ACL is created through the Console or API.
               Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an ACL resource within HuaweiCloud IAM service. The ACL allowing user access only from specified IP address
        ranges and IPv4 CIDR blocks. The ACL take effect for IAM users under the Domain account rather than the account itself.

        > **NOTE:** You *must* have admin privileges to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        acl = huaweicloud.iam.Acl("acl",
            ip_cidrs=[huaweicloud.iam.AclIpCidrArgs(
                cidr="159.138.39.192/32",
                description="This is a test ip address",
            )],
            ip_ranges=[huaweicloud.iam.AclIpRangeArgs(
                description="This is a test ip range",
                range="0.0.0.0-255.255.255.0",
            )],
            type="console")
        ```

        :param str resource_name: The name of the resource.
        :param AclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpCidrArgs']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpRangeArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AclArgs.__new__(AclArgs)

            __props__.__dict__["ip_cidrs"] = ip_cidrs
            __props__.__dict__["ip_ranges"] = ip_ranges
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Acl, __self__).__init__(
            'huaweicloud:Iam/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpCidrArgs']]]]] = None,
            ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpRangeArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Acl':
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpCidrArgs']]]] ip_cidrs: Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
               The `ip_cidrs` cannot repeat. The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclIpRangeArgs']]]] ip_ranges: Specifies the IP address ranges from which console access or api access is allowed.
               The `ip_ranges` cannot repeat. The object structure is documented below.
        :param pulumi.Input[str] type: Specifies the ACL is created through the Console or API.
               Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AclState.__new__(_AclState)

        __props__.__dict__["ip_cidrs"] = ip_cidrs
        __props__.__dict__["ip_ranges"] = ip_ranges
        __props__.__dict__["type"] = type
        return Acl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipCidrs")
    def ip_cidrs(self) -> pulumi.Output[Optional[Sequence['outputs.AclIpCidr']]]:
        """
        Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
        The `ip_cidrs` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_cidrs")

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> pulumi.Output[Optional[Sequence['outputs.AclIpRange']]]:
        """
        Specifies the IP address ranges from which console access or api access is allowed.
        The `ip_ranges` cannot repeat. The object structure is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the ACL is created through the Console or API.
        Valid values are **console** and **api**. Changing this parameter will create a new ACL.
        """
        return pulumi.get(self, "type")

