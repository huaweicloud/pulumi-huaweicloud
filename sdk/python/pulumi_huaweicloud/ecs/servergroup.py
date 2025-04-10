# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServergroupArgs', 'Servergroup']

@pulumi.input_type
class ServergroupArgs:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Servergroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies an array of one or more instance ID to attach server group.
        :param pulumi.Input[str] name: Specifies a unique name for the server group. This parameter can contain a
               maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
               creates a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Specifies the set of policies for the server group. Only *anti-affinity*
               policies are supported.
        :param pulumi.Input[str] region: Specifies the region in which to create the server group resource. If omitted,
               the provider-level region will be used. Changing this creates a new server group.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more instance ID to attach server group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the server group. This parameter can contain a
        maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
        creates a new server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the set of policies for the server group. Only *anti-affinity*
        policies are supported.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the server group resource. If omitted,
        the provider-level region will be used. Changing this creates a new server group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ServergroupState:
    def __init__(__self__, *,
                 fault_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Servergroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fault_domains: schema: Internal
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies an array of one or more instance ID to attach server group.
        :param pulumi.Input[str] name: Specifies a unique name for the server group. This parameter can contain a
               maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
               creates a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Specifies the set of policies for the server group. Only *anti-affinity*
               policies are supported.
        :param pulumi.Input[str] region: Specifies the region in which to create the server group resource. If omitted,
               the provider-level region will be used. Changing this creates a new server group.
        """
        if fault_domains is not None:
            pulumi.set(__self__, "fault_domains", fault_domains)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="faultDomains")
    def fault_domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        schema: Internal
        """
        return pulumi.get(self, "fault_domains")

    @fault_domains.setter
    def fault_domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "fault_domains", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more instance ID to attach server group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the server group. This parameter can contain a
        maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
        creates a new server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the set of policies for the server group. Only *anti-affinity*
        policies are supported.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the server group resource. If omitted,
        the provider-level region will be used. Changing this creates a new server group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Servergroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Server Group resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        instance_demo = huaweicloud.Ecs.get_instance(name="ecs-servergroup-demo")
        test_sg = huaweicloud.ecs.Servergroup("test-sg",
            policies=["anti-affinity"],
            members=[instance_demo.id])
        ```

        ## Import

        Server Groups can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Ecs/servergroup:Servergroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies an array of one or more instance ID to attach server group.
        :param pulumi.Input[str] name: Specifies a unique name for the server group. This parameter can contain a
               maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
               creates a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Specifies the set of policies for the server group. Only *anti-affinity*
               policies are supported.
        :param pulumi.Input[str] region: Specifies the region in which to create the server group resource. If omitted,
               the provider-level region will be used. Changing this creates a new server group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServergroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Server Group resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        instance_demo = huaweicloud.Ecs.get_instance(name="ecs-servergroup-demo")
        test_sg = huaweicloud.ecs.Servergroup("test-sg",
            policies=["anti-affinity"],
            members=[instance_demo.id])
        ```

        ## Import

        Server Groups can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Ecs/servergroup:Servergroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
        ```

        :param str resource_name: The name of the resource.
        :param ServergroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServergroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServergroupArgs.__new__(ServergroupArgs)

            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["policies"] = policies
            __props__.__dict__["region"] = region
            __props__.__dict__["fault_domains"] = None
        super(Servergroup, __self__).__init__(
            'huaweicloud:Ecs/servergroup:Servergroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fault_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Servergroup':
        """
        Get an existing Servergroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fault_domains: schema: Internal
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies an array of one or more instance ID to attach server group.
        :param pulumi.Input[str] name: Specifies a unique name for the server group. This parameter can contain a
               maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
               creates a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Specifies the set of policies for the server group. Only *anti-affinity*
               policies are supported.
        :param pulumi.Input[str] region: Specifies the region in which to create the server group resource. If omitted,
               the provider-level region will be used. Changing this creates a new server group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServergroupState.__new__(_ServergroupState)

        __props__.__dict__["fault_domains"] = fault_domains
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["policies"] = policies
        __props__.__dict__["region"] = region
        return Servergroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="faultDomains")
    def fault_domains(self) -> pulumi.Output[Sequence[str]]:
        """
        schema: Internal
        """
        return pulumi.get(self, "fault_domains")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies an array of one or more instance ID to attach server group.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies a unique name for the server group. This parameter can contain a
        maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
        creates a new server group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the set of policies for the server group. Only *anti-affinity*
        policies are supported.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the server group resource. If omitted,
        the provider-level region will be used. Changing this creates a new server group.
        """
        return pulumi.get(self, "region")

