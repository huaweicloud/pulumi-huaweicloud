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

__all__ = ['CustomAuthorizerArgs', 'CustomAuthorizer']

@pulumi.input_type
class CustomAuthorizerArgs:
    def __init__(__self__, *,
                 function_urn: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 cache_age: Optional[pulumi.Input[int]] = None,
                 function_alias_uri: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 identities: Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]] = None,
                 is_body_send: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_data: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomAuthorizer resource.
        :param pulumi.Input[str] function_urn: Specifies the uniform function URN of the function graph resource.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the
               custom authorizer belongs to.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[int] cache_age: Specifies the maximum cache age.  
               The valid value is range from `1` to `3,600`.
        :param pulumi.Input[str] function_alias_uri: Specifies the version alias URI of the FGS function.
        :param pulumi.Input[str] function_version: Specifies the version of the FGS function.
        :param pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]] identities: Specifies an array of one or more parameter identities of the custom authorizer.
               The object structure is documented below.
        :param pulumi.Input[bool] is_body_send: Specifies whether to send the body.
        :param pulumi.Input[str] name: Specifies the name of the parameter to be verified.
               The parameter includes front-end and back-end parameters.
        :param pulumi.Input[str] network_type: Specifies the framework type of the function.
               + **V1**: Non-VPC network architecture.
               + **V2**: VPC network architecture.
        :param pulumi.Input[str] region: Specifies the region in which to create the custom authorizer resource.
               If omitted, the provider-level region will be used.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] type: Specifies the custom authorize type.
               The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] user_data: Specifies the user data, which can contain a maximum of `2,048` characters.
               The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        pulumi.set(__self__, "function_urn", function_urn)
        pulumi.set(__self__, "instance_id", instance_id)
        if cache_age is not None:
            pulumi.set(__self__, "cache_age", cache_age)
        if function_alias_uri is not None:
            pulumi.set(__self__, "function_alias_uri", function_alias_uri)
        if function_version is not None:
            pulumi.set(__self__, "function_version", function_version)
        if identities is not None:
            pulumi.set(__self__, "identities", identities)
        if is_body_send is not None:
            pulumi.set(__self__, "is_body_send", is_body_send)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_data is not None:
            pulumi.set(__self__, "user_data", user_data)

    @property
    @pulumi.getter(name="functionUrn")
    def function_urn(self) -> pulumi.Input[str]:
        """
        Specifies the uniform function URN of the function graph resource.
        """
        return pulumi.get(self, "function_urn")

    @function_urn.setter
    def function_urn(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_urn", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the
        custom authorizer belongs to.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="cacheAge")
    def cache_age(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum cache age.  
        The valid value is range from `1` to `3,600`.
        """
        return pulumi.get(self, "cache_age")

    @cache_age.setter
    def cache_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cache_age", value)

    @property
    @pulumi.getter(name="functionAliasUri")
    def function_alias_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version alias URI of the FGS function.
        """
        return pulumi.get(self, "function_alias_uri")

    @function_alias_uri.setter
    def function_alias_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_alias_uri", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the FGS function.
        """
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]]:
        """
        Specifies an array of one or more parameter identities of the custom authorizer.
        The object structure is documented below.
        """
        return pulumi.get(self, "identities")

    @identities.setter
    def identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]]):
        pulumi.set(self, "identities", value)

    @property
    @pulumi.getter(name="isBodySend")
    def is_body_send(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to send the body.
        """
        return pulumi.get(self, "is_body_send")

    @is_body_send.setter
    def is_body_send(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_body_send", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the parameter to be verified.
        The parameter includes front-end and back-end parameters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the framework type of the function.
        + **V1**: Non-VPC network architecture.
        + **V2**: VPC network architecture.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the custom authorizer resource.
        If omitted, the provider-level region will be used.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the custom authorize type.
        The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user data, which can contain a maximum of `2,048` characters.
        The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        return pulumi.get(self, "user_data")

    @user_data.setter
    def user_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_data", value)


@pulumi.input_type
class _CustomAuthorizerState:
    def __init__(__self__, *,
                 cache_age: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 function_alias_uri: Optional[pulumi.Input[str]] = None,
                 function_urn: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 identities: Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_body_send: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomAuthorizer resources.
        :param pulumi.Input[int] cache_age: Specifies the maximum cache age.  
               The valid value is range from `1` to `3,600`.
        :param pulumi.Input[str] created_at: The creation time of the custom authorizer.
        :param pulumi.Input[str] function_alias_uri: Specifies the version alias URI of the FGS function.
        :param pulumi.Input[str] function_urn: Specifies the uniform function URN of the function graph resource.
        :param pulumi.Input[str] function_version: Specifies the version of the FGS function.
        :param pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]] identities: Specifies an array of one or more parameter identities of the custom authorizer.
               The object structure is documented below.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the
               custom authorizer belongs to.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[bool] is_body_send: Specifies whether to send the body.
        :param pulumi.Input[str] name: Specifies the name of the parameter to be verified.
               The parameter includes front-end and back-end parameters.
        :param pulumi.Input[str] network_type: Specifies the framework type of the function.
               + **V1**: Non-VPC network architecture.
               + **V2**: VPC network architecture.
        :param pulumi.Input[str] region: Specifies the region in which to create the custom authorizer resource.
               If omitted, the provider-level region will be used.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] type: Specifies the custom authorize type.
               The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] user_data: Specifies the user data, which can contain a maximum of `2,048` characters.
               The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        if cache_age is not None:
            pulumi.set(__self__, "cache_age", cache_age)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if function_alias_uri is not None:
            pulumi.set(__self__, "function_alias_uri", function_alias_uri)
        if function_urn is not None:
            pulumi.set(__self__, "function_urn", function_urn)
        if function_version is not None:
            pulumi.set(__self__, "function_version", function_version)
        if identities is not None:
            pulumi.set(__self__, "identities", identities)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_body_send is not None:
            pulumi.set(__self__, "is_body_send", is_body_send)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_data is not None:
            pulumi.set(__self__, "user_data", user_data)

    @property
    @pulumi.getter(name="cacheAge")
    def cache_age(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum cache age.  
        The valid value is range from `1` to `3,600`.
        """
        return pulumi.get(self, "cache_age")

    @cache_age.setter
    def cache_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cache_age", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the custom authorizer.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="functionAliasUri")
    def function_alias_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version alias URI of the FGS function.
        """
        return pulumi.get(self, "function_alias_uri")

    @function_alias_uri.setter
    def function_alias_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_alias_uri", value)

    @property
    @pulumi.getter(name="functionUrn")
    def function_urn(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the uniform function URN of the function graph resource.
        """
        return pulumi.get(self, "function_urn")

    @function_urn.setter
    def function_urn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_urn", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the FGS function.
        """
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]]:
        """
        Specifies an array of one or more parameter identities of the custom authorizer.
        The object structure is documented below.
        """
        return pulumi.get(self, "identities")

    @identities.setter
    def identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomAuthorizerIdentityArgs']]]]):
        pulumi.set(self, "identities", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies an ID of the APIG dedicated instance to which the
        custom authorizer belongs to.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isBodySend")
    def is_body_send(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to send the body.
        """
        return pulumi.get(self, "is_body_send")

    @is_body_send.setter
    def is_body_send(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_body_send", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the parameter to be verified.
        The parameter includes front-end and back-end parameters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the framework type of the function.
        + **V1**: Non-VPC network architecture.
        + **V2**: VPC network architecture.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the custom authorizer resource.
        If omitted, the provider-level region will be used.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the custom authorize type.
        The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user data, which can contain a maximum of `2,048` characters.
        The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        return pulumi.get(self, "user_data")

    @user_data.setter
    def user_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_data", value)


class CustomAuthorizer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_age: Optional[pulumi.Input[int]] = None,
                 function_alias_uri: Optional[pulumi.Input[str]] = None,
                 function_urn: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomAuthorizerIdentityArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_body_send: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an APIG custom authorizer resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        authorizer_name = config.require_object("authorizerName")
        function_urn = config.require_object("functionUrn")
        test = huaweicloud.dedicated_apig.CustomAuthorizer("test",
            instance_id=instance_id,
            function_urn=function_urn,
            function_version="latest",
            type="FRONTEND",
            cache_age=60,
            identities=[huaweicloud.dedicated_apig.CustomAuthorizerIdentityArgs(
                name="user_name",
                location="QUERY",
            )])
        ```

        ## Import

        Custom Authorizers of the APIG can be imported using their `name` and related dedicated instance IDs, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer test <instance_id>/<name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cache_age: Specifies the maximum cache age.  
               The valid value is range from `1` to `3,600`.
        :param pulumi.Input[str] function_alias_uri: Specifies the version alias URI of the FGS function.
        :param pulumi.Input[str] function_urn: Specifies the uniform function URN of the function graph resource.
        :param pulumi.Input[str] function_version: Specifies the version of the FGS function.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomAuthorizerIdentityArgs']]]] identities: Specifies an array of one or more parameter identities of the custom authorizer.
               The object structure is documented below.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the
               custom authorizer belongs to.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[bool] is_body_send: Specifies whether to send the body.
        :param pulumi.Input[str] name: Specifies the name of the parameter to be verified.
               The parameter includes front-end and back-end parameters.
        :param pulumi.Input[str] network_type: Specifies the framework type of the function.
               + **V1**: Non-VPC network architecture.
               + **V2**: VPC network architecture.
        :param pulumi.Input[str] region: Specifies the region in which to create the custom authorizer resource.
               If omitted, the provider-level region will be used.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] type: Specifies the custom authorize type.
               The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] user_data: Specifies the user data, which can contain a maximum of `2,048` characters.
               The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomAuthorizerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an APIG custom authorizer resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        authorizer_name = config.require_object("authorizerName")
        function_urn = config.require_object("functionUrn")
        test = huaweicloud.dedicated_apig.CustomAuthorizer("test",
            instance_id=instance_id,
            function_urn=function_urn,
            function_version="latest",
            type="FRONTEND",
            cache_age=60,
            identities=[huaweicloud.dedicated_apig.CustomAuthorizerIdentityArgs(
                name="user_name",
                location="QUERY",
            )])
        ```

        ## Import

        Custom Authorizers of the APIG can be imported using their `name` and related dedicated instance IDs, separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer test <instance_id>/<name>
        ```

        :param str resource_name: The name of the resource.
        :param CustomAuthorizerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomAuthorizerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_age: Optional[pulumi.Input[int]] = None,
                 function_alias_uri: Optional[pulumi.Input[str]] = None,
                 function_urn: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomAuthorizerIdentityArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_body_send: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomAuthorizerArgs.__new__(CustomAuthorizerArgs)

            __props__.__dict__["cache_age"] = cache_age
            __props__.__dict__["function_alias_uri"] = function_alias_uri
            if function_urn is None and not opts.urn:
                raise TypeError("Missing required property 'function_urn'")
            __props__.__dict__["function_urn"] = function_urn
            __props__.__dict__["function_version"] = function_version
            __props__.__dict__["identities"] = identities
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["is_body_send"] = is_body_send
            __props__.__dict__["name"] = name
            __props__.__dict__["network_type"] = network_type
            __props__.__dict__["region"] = region
            __props__.__dict__["type"] = type
            __props__.__dict__["user_data"] = user_data
            __props__.__dict__["created_at"] = None
        super(CustomAuthorizer, __self__).__init__(
            'huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_age: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            function_alias_uri: Optional[pulumi.Input[str]] = None,
            function_urn: Optional[pulumi.Input[str]] = None,
            function_version: Optional[pulumi.Input[str]] = None,
            identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomAuthorizerIdentityArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_body_send: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_data: Optional[pulumi.Input[str]] = None) -> 'CustomAuthorizer':
        """
        Get an existing CustomAuthorizer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cache_age: Specifies the maximum cache age.  
               The valid value is range from `1` to `3,600`.
        :param pulumi.Input[str] created_at: The creation time of the custom authorizer.
        :param pulumi.Input[str] function_alias_uri: Specifies the version alias URI of the FGS function.
        :param pulumi.Input[str] function_urn: Specifies the uniform function URN of the function graph resource.
        :param pulumi.Input[str] function_version: Specifies the version of the FGS function.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomAuthorizerIdentityArgs']]]] identities: Specifies an array of one or more parameter identities of the custom authorizer.
               The object structure is documented below.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the
               custom authorizer belongs to.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[bool] is_body_send: Specifies whether to send the body.
        :param pulumi.Input[str] name: Specifies the name of the parameter to be verified.
               The parameter includes front-end and back-end parameters.
        :param pulumi.Input[str] network_type: Specifies the framework type of the function.
               + **V1**: Non-VPC network architecture.
               + **V2**: VPC network architecture.
        :param pulumi.Input[str] region: Specifies the region in which to create the custom authorizer resource.
               If omitted, the provider-level region will be used.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] type: Specifies the custom authorize type.
               The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
               Changing this will create a new custom authorizer resource.
        :param pulumi.Input[str] user_data: Specifies the user data, which can contain a maximum of `2,048` characters.
               The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomAuthorizerState.__new__(_CustomAuthorizerState)

        __props__.__dict__["cache_age"] = cache_age
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["function_alias_uri"] = function_alias_uri
        __props__.__dict__["function_urn"] = function_urn
        __props__.__dict__["function_version"] = function_version
        __props__.__dict__["identities"] = identities
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_body_send"] = is_body_send
        __props__.__dict__["name"] = name
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["region"] = region
        __props__.__dict__["type"] = type
        __props__.__dict__["user_data"] = user_data
        return CustomAuthorizer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cacheAge")
    def cache_age(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum cache age.  
        The valid value is range from `1` to `3,600`.
        """
        return pulumi.get(self, "cache_age")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time of the custom authorizer.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="functionAliasUri")
    def function_alias_uri(self) -> pulumi.Output[str]:
        """
        Specifies the version alias URI of the FGS function.
        """
        return pulumi.get(self, "function_alias_uri")

    @property
    @pulumi.getter(name="functionUrn")
    def function_urn(self) -> pulumi.Output[str]:
        """
        Specifies the uniform function URN of the function graph resource.
        """
        return pulumi.get(self, "function_urn")

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Output[str]:
        """
        Specifies the version of the FGS function.
        """
        return pulumi.get(self, "function_version")

    @property
    @pulumi.getter
    def identities(self) -> pulumi.Output[Optional[Sequence['outputs.CustomAuthorizerIdentity']]]:
        """
        Specifies an array of one or more parameter identities of the custom authorizer.
        The object structure is documented below.
        """
        return pulumi.get(self, "identities")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the
        custom authorizer belongs to.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isBodySend")
    def is_body_send(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to send the body.
        """
        return pulumi.get(self, "is_body_send")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the parameter to be verified.
        The parameter includes front-end and back-end parameters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        Specifies the framework type of the function.
        + **V1**: Non-VPC network architecture.
        + **V2**: VPC network architecture.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the custom authorizer resource.
        If omitted, the provider-level region will be used.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the custom authorize type.
        The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
        Changing this will create a new custom authorizer resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the user data, which can contain a maximum of `2,048` characters.
        The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
        """
        return pulumi.get(self, "user_data")

