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

__all__ = ['ResponseArgs', 'Response']

@pulumi.input_type
class ResponseArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]] = None):
        """
        The set of arguments for constructing a Response resource.
        :param pulumi.Input[str] group_id: Specifies the ID of the API group to which the API custom response
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the API group and the
               API custom response belong.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the API custom response.  
               The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region where the API custom response is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]] rules: Specifies the API custom response rules definition.  
               The object structure is documented below.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the API group to which the API custom response
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the dedicated instance to which the API group and the
        API custom response belong.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API custom response.  
        The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the API custom response is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]]:
        """
        Specifies the API custom response rules definition.  
        The object structure is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class _ResponseState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Response resources.
        :param pulumi.Input[str] created_at: The creation time of the API custom response.
        :param pulumi.Input[str] group_id: Specifies the ID of the API group to which the API custom response
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the API group and the
               API custom response belong.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the API custom response.  
               The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region where the API custom response is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]] rules: Specifies the API custom response rules definition.  
               The object structure is documented below.
        :param pulumi.Input[str] updated_at: The latest update time of the API custom response.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the API custom response.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the API group to which the API custom response
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the dedicated instance to which the API group and the
        API custom response belong.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API custom response.  
        The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the API custom response is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]]:
        """
        Specifies the API custom response rules definition.  
        The object structure is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponseRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The latest update time of the API custom response.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Response(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponseRuleArgs']]]]] = None,
                 __props__=None):
        """
        Manages an APIG (API) custom response resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        group_id = config.require_object("groupId")
        response_name = config.require_object("responseName")
        test = huaweicloud.dedicated_apig.Response("test",
            instance_id=instance_id,
            group_id=group_id,
            rules=[huaweicloud.dedicated_apig.ResponseRuleArgs(
                error_type="AUTHORIZER_FAILURE",
                body="{\\"code\\":\\"$context.authorizer.frontend.code\\",\\"message\\":\\"$context.authorizer.frontend.message\\"}",
                status_code=401,
            )])
        ```

        ## Import

        API Responses can be imported using their `name` and IDs of the APIG dedicated instances and API groups to which the API response belongs, separated by slashes, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/response:Response test <instance_id>/<group_id>/<name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: Specifies the ID of the API group to which the API custom response
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the API group and the
               API custom response belong.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the API custom response.  
               The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region where the API custom response is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponseRuleArgs']]]] rules: Specifies the API custom response rules definition.  
               The object structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResponseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an APIG (API) custom response resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        group_id = config.require_object("groupId")
        response_name = config.require_object("responseName")
        test = huaweicloud.dedicated_apig.Response("test",
            instance_id=instance_id,
            group_id=group_id,
            rules=[huaweicloud.dedicated_apig.ResponseRuleArgs(
                error_type="AUTHORIZER_FAILURE",
                body="{\\"code\\":\\"$context.authorizer.frontend.code\\",\\"message\\":\\"$context.authorizer.frontend.message\\"}",
                status_code=401,
            )])
        ```

        ## Import

        API Responses can be imported using their `name` and IDs of the APIG dedicated instances and API groups to which the API response belongs, separated by slashes, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/response:Response test <instance_id>/<group_id>/<name>
        ```

        :param str resource_name: The name of the resource.
        :param ResponseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResponseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponseRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResponseArgs.__new__(ResponseArgs)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(Response, __self__).__init__(
            'huaweicloud:DedicatedApig/response:Response',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponseRuleArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Response':
        """
        Get an existing Response resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The creation time of the API custom response.
        :param pulumi.Input[str] group_id: Specifies the ID of the API group to which the API custom response
               belongs.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the dedicated instance to which the API group and the
               API custom response belong.
               Changing this will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the API custom response.  
               The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] region: Specifies the region where the API custom response is located.  
               If omitted, the provider-level region will be used. Changing this will create a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponseRuleArgs']]]] rules: Specifies the API custom response rules definition.  
               The object structure is documented below.
        :param pulumi.Input[str] updated_at: The latest update time of the API custom response.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResponseState.__new__(_ResponseState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        __props__.__dict__["updated_at"] = updated_at
        return Response(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time of the API custom response.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the API group to which the API custom response
        belongs.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the dedicated instance to which the API group and the
        API custom response belong.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the API custom response.  
        The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region where the API custom response is located.  
        If omitted, the provider-level region will be used. Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.ResponseRule']]:
        """
        Specifies the API custom response rules definition.  
        The object structure is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The latest update time of the API custom response.
        """
        return pulumi.get(self, "updated_at")
