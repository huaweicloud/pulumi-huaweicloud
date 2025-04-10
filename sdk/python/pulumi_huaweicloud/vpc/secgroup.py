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

__all__ = ['SecgroupArgs', 'Secgroup']

@pulumi.input_type
class SecgroupArgs:
    def __init__(__self__, *,
                 delete_default_rules: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Secgroup resource.
        :param pulumi.Input[bool] delete_default_rules: Specifies whether or not to delete the default security rules.
               This is `false` by default.
        :param pulumi.Input[str] description: Specifies the description for the security group.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project id of the security group.
               Changing this creates a new security group.
        :param pulumi.Input[str] name: Specifies a unique name for the security group.
        :param pulumi.Input[str] region: The region in which to create the security group resource. If omitted, the
               provider-level region will be used. Changing this creates a new security group resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the security group.
        """
        if delete_default_rules is not None:
            pulumi.set(__self__, "delete_default_rules", delete_default_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deleteDefaultRules")
    def delete_default_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether or not to delete the default security rules.
        This is `false` by default.
        """
        return pulumi.get(self, "delete_default_rules")

    @delete_default_rules.setter
    def delete_default_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_default_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description for the security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project id of the security group.
        Changing this creates a new security group.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the security group resource. If omitted, the
        provider-level region will be used. Changing this creates a new security group resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs to associate with the security group.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecgroupState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 delete_default_rules: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['SecgroupRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Secgroup resources.
        :param pulumi.Input[str] created_at: The creation time, in UTC format.
        :param pulumi.Input[bool] delete_default_rules: Specifies whether or not to delete the default security rules.
               This is `false` by default.
        :param pulumi.Input[str] description: Specifies the description for the security group.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project id of the security group.
               Changing this creates a new security group.
        :param pulumi.Input[str] name: Specifies a unique name for the security group.
        :param pulumi.Input[str] region: The region in which to create the security group resource. If omitted, the
               provider-level region will be used. Changing this creates a new security group resource.
        :param pulumi.Input[Sequence[pulumi.Input['SecgroupRuleArgs']]] rules: The array of security group rules associating with the security group.
               The rule object is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the security group.
        :param pulumi.Input[str] updated_at: The last update time, in UTC format.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if delete_default_rules is not None:
            pulumi.set(__self__, "delete_default_rules", delete_default_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enterprise_project_id is not None:
            pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time, in UTC format.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deleteDefaultRules")
    def delete_default_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether or not to delete the default security rules.
        This is `false` by default.
        """
        return pulumi.get(self, "delete_default_rules")

    @delete_default_rules.setter
    def delete_default_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_default_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description for the security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the enterprise project id of the security group.
        Changing this creates a new security group.
        """
        return pulumi.get(self, "enterprise_project_id")

    @enterprise_project_id.setter
    def enterprise_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the security group resource. If omitted, the
        provider-level region will be used. Changing this creates a new security group resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecgroupRuleArgs']]]]:
        """
        The array of security group rules associating with the security group.
        The rule object is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecgroupRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs to associate with the security group.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The last update time, in UTC format.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Secgroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_default_rules: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Import

        Security Groups can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Vpc/secgroup:Secgroup secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_default_rules: Specifies whether or not to delete the default security rules.
               This is `false` by default.
        :param pulumi.Input[str] description: Specifies the description for the security group.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project id of the security group.
               Changing this creates a new security group.
        :param pulumi.Input[str] name: Specifies a unique name for the security group.
        :param pulumi.Input[str] region: The region in which to create the security group resource. If omitted, the
               provider-level region will be used. Changing this creates a new security group resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the security group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecgroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Security Groups can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Vpc/secgroup:Secgroup secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
        ```

        :param str resource_name: The name of the resource.
        :param SecgroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecgroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_default_rules: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enterprise_project_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecgroupArgs.__new__(SecgroupArgs)

            __props__.__dict__["delete_default_rules"] = delete_default_rules
            __props__.__dict__["description"] = description
            __props__.__dict__["enterprise_project_id"] = enterprise_project_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["rules"] = None
            __props__.__dict__["updated_at"] = None
        super(Secgroup, __self__).__init__(
            'huaweicloud:Vpc/secgroup:Secgroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            delete_default_rules: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enterprise_project_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecgroupRuleArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Secgroup':
        """
        Get an existing Secgroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The creation time, in UTC format.
        :param pulumi.Input[bool] delete_default_rules: Specifies whether or not to delete the default security rules.
               This is `false` by default.
        :param pulumi.Input[str] description: Specifies the description for the security group.
        :param pulumi.Input[str] enterprise_project_id: Specifies the enterprise project id of the security group.
               Changing this creates a new security group.
        :param pulumi.Input[str] name: Specifies a unique name for the security group.
        :param pulumi.Input[str] region: The region in which to create the security group resource. If omitted, the
               provider-level region will be used. Changing this creates a new security group resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecgroupRuleArgs']]]] rules: The array of security group rules associating with the security group.
               The rule object is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the security group.
        :param pulumi.Input[str] updated_at: The last update time, in UTC format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecgroupState.__new__(_SecgroupState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["delete_default_rules"] = delete_default_rules
        __props__.__dict__["description"] = description
        __props__.__dict__["enterprise_project_id"] = enterprise_project_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return Secgroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time, in UTC format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deleteDefaultRules")
    def delete_default_rules(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether or not to delete the default security rules.
        This is `false` by default.
        """
        return pulumi.get(self, "delete_default_rules")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description for the security group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> pulumi.Output[str]:
        """
        Specifies the enterprise project id of the security group.
        Changing this creates a new security group.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies a unique name for the security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the security group resource. If omitted, the
        provider-level region will be used. Changing this creates a new security group resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.SecgroupRule']]:
        """
        The array of security group rules associating with the security group.
        The rule object is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the key/value pairs to associate with the security group.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The last update time, in UTC format.
        """
        return pulumi.get(self, "updated_at")

