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

__all__ = ['AddonArgs', 'Addon']

@pulumi.input_type
class AddonArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 template_name: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input['AddonValuesArgs']] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Addon resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] template_name: Specifies the name of the add-on template.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the CCE add-on resource.
               If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        :param pulumi.Input['AddonValuesArgs'] values: Specifies the add-on template installation parameters.
               These parameters vary depending on the add-on. The structure is documented below.
        :param pulumi.Input[str] version: Specifies the version of the add-on.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "template_name", template_name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if values is not None:
            pulumi.set(__self__, "values", values)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Specifies the cluster ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the add-on template.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the CCE add-on resource.
        If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input['AddonValuesArgs']]:
        """
        Specifies the add-on template installation parameters.
        These parameters vary depending on the add-on. The structure is documented below.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input['AddonValuesArgs']]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the add-on.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _AddonState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input['AddonValuesArgs']] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Addon resources.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] description: Description of add-on instance.
        :param pulumi.Input[str] region: Specifies the region in which to create the CCE add-on resource.
               If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        :param pulumi.Input[str] status: Add-on status information.
        :param pulumi.Input[str] template_name: Specifies the name of the add-on template.
               Changing this parameter will create a new resource.
        :param pulumi.Input['AddonValuesArgs'] values: Specifies the add-on template installation parameters.
               These parameters vary depending on the add-on. The structure is documented below.
        :param pulumi.Input[str] version: Specifies the version of the add-on.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if template_name is not None:
            pulumi.set(__self__, "template_name", template_name)
        if values is not None:
            pulumi.set(__self__, "values", values)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the cluster ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of add-on instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the CCE add-on resource.
        If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Add-on status information.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the add-on template.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input['AddonValuesArgs']]:
        """
        Specifies the add-on template installation parameters.
        These parameters vary depending on the add-on. The structure is documented below.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input['AddonValuesArgs']]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the add-on.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Addon(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[pulumi.InputType['AddonValuesArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        CCE add-on can be imported using the cluster ID and add-on ID separated by a slash, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Cce/addon:Addon my_addon <cluster_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the CCE add-on resource.
               If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        :param pulumi.Input[str] template_name: Specifies the name of the add-on template.
               Changing this parameter will create a new resource.
        :param pulumi.Input[pulumi.InputType['AddonValuesArgs']] values: Specifies the add-on template installation parameters.
               These parameters vary depending on the add-on. The structure is documented below.
        :param pulumi.Input[str] version: Specifies the version of the add-on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AddonArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        CCE add-on can be imported using the cluster ID and add-on ID separated by a slash, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Cce/addon:Addon my_addon <cluster_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param AddonArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AddonArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[pulumi.InputType['AddonValuesArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AddonArgs.__new__(AddonArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["region"] = region
            if template_name is None and not opts.urn:
                raise TypeError("Missing required property 'template_name'")
            __props__.__dict__["template_name"] = template_name
            __props__.__dict__["values"] = values
            __props__.__dict__["version"] = version
            __props__.__dict__["description"] = None
            __props__.__dict__["status"] = None
        super(Addon, __self__).__init__(
            'huaweicloud:Cce/addon:Addon',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            template_name: Optional[pulumi.Input[str]] = None,
            values: Optional[pulumi.Input[pulumi.InputType['AddonValuesArgs']]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Addon':
        """
        Get an existing Addon resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] description: Description of add-on instance.
        :param pulumi.Input[str] region: Specifies the region in which to create the CCE add-on resource.
               If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        :param pulumi.Input[str] status: Add-on status information.
        :param pulumi.Input[str] template_name: Specifies the name of the add-on template.
               Changing this parameter will create a new resource.
        :param pulumi.Input[pulumi.InputType['AddonValuesArgs']] values: Specifies the add-on template installation parameters.
               These parameters vary depending on the add-on. The structure is documented below.
        :param pulumi.Input[str] version: Specifies the version of the add-on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AddonState.__new__(_AddonState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["description"] = description
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["template_name"] = template_name
        __props__.__dict__["values"] = values
        __props__.__dict__["version"] = version
        return Addon(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Specifies the cluster ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of add-on instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the CCE add-on resource.
        If omitted, the provider-level region will be used. Changing this creates a new CCE add-on resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Add-on status information.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the add-on template.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter
    def values(self) -> pulumi.Output[Optional['outputs.AddonValues']]:
        """
        Specifies the add-on template installation parameters.
        These parameters vary depending on the add-on. The structure is documented below.
        """
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Specifies the version of the add-on.
        """
        return pulumi.get(self, "version")

