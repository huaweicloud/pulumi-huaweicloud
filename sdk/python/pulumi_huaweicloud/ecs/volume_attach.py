# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VolumeAttachArgs', 'VolumeAttach']

@pulumi.input_type
class VolumeAttachArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 volume_id: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VolumeAttach resource.
        :param pulumi.Input[str] instance_id: Specifies the ID of the Instance to attach the Volume to.
        :param pulumi.Input[str] volume_id: Specifies the ID of the Volume to attach to an Instance.
        :param pulumi.Input[str] device: Specifies the device of the volume attachment (ex: `/dev/vdc`).
        :param pulumi.Input[str] region: Specifies the region in which to create the volume resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "volume_id", volume_id)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Instance to attach the Volume to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the Volume to attach to an Instance.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_id", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the device of the volume attachment (ex: `/dev/vdc`).
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the volume resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _VolumeAttachState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 pci_address: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VolumeAttach resources.
        :param pulumi.Input[str] device: Specifies the device of the volume attachment (ex: `/dev/vdc`).
        :param pulumi.Input[str] instance_id: Specifies the ID of the Instance to attach the Volume to.
        :param pulumi.Input[str] pci_address: PCI address of the block device.
        :param pulumi.Input[str] region: Specifies the region in which to create the volume resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] volume_id: Specifies the ID of the Volume to attach to an Instance.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if pci_address is not None:
            pulumi.set(__self__, "pci_address", pci_address)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the device of the volume attachment (ex: `/dev/vdc`).
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the Instance to attach the Volume to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="pciAddress")
    def pci_address(self) -> Optional[pulumi.Input[str]]:
        """
        PCI address of the block device.
        """
        return pulumi.get(self, "pci_address")

    @pci_address.setter
    def pci_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pci_address", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the volume resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the Volume to attach to an Instance.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_id", value)


class VolumeAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attaches a volume to an ECS Instance.

        ## Example Usage
        ### Basic attachment of a single volume to a single instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        myvol = huaweicloud.evs.Volume("myvol",
            availability_zone="cn-north-4a",
            volume_type="SAS",
            size=10)
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a",
            networks=[huaweicloud.ecs.InstanceNetworkArgs(
                uuid="55534eaa-533a-419d-9b40-ec427ea7195a",
            )])
        attached = huaweicloud.ecs.VolumeAttach("attached",
            instance_id=myinstance.id,
            volume_id=myvol.id)
        ```
        ### Attaching multiple volumes to a single instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        myvol = []
        for range in [{"value": i} for i in range(0, 2)]:
            myvol.append(huaweicloud.evs.Volume(f"myvol-{range['value']}",
                availability_zone="cn-north-4a",
                volume_type="SAS",
                size=10))
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a")
        attachments = []
        for range in [{"value": i} for i in range(0, 2)]:
            attachments.append(huaweicloud.ecs.VolumeAttach(f"attachments-{range['value']}",
                instance_id=myinstance.id,
                volume_id=[__item.id for __item in myvol][range["value"]]))
        pulumi.export("volume devices", [__item.device for __item in attachments])
        ```

        ## Import

        Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:Ecs/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Specifies the device of the volume attachment (ex: `/dev/vdc`).
        :param pulumi.Input[str] instance_id: Specifies the ID of the Instance to attach the Volume to.
        :param pulumi.Input[str] region: Specifies the region in which to create the volume resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] volume_id: Specifies the ID of the Volume to attach to an Instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches a volume to an ECS Instance.

        ## Example Usage
        ### Basic attachment of a single volume to a single instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        myvol = huaweicloud.evs.Volume("myvol",
            availability_zone="cn-north-4a",
            volume_type="SAS",
            size=10)
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a",
            networks=[huaweicloud.ecs.InstanceNetworkArgs(
                uuid="55534eaa-533a-419d-9b40-ec427ea7195a",
            )])
        attached = huaweicloud.ecs.VolumeAttach("attached",
            instance_id=myinstance.id,
            volume_id=myvol.id)
        ```
        ### Attaching multiple volumes to a single instance

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        security_group_id = config.require_object("securityGroupId")
        myvol = []
        for range in [{"value": i} for i in range(0, 2)]:
            myvol.append(huaweicloud.evs.Volume(f"myvol-{range['value']}",
                availability_zone="cn-north-4a",
                volume_type="SAS",
                size=10))
        myinstance = huaweicloud.ecs.Instance("myinstance",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="s6.small.1",
            key_pair="my_key_pair_name",
            security_group_ids=[security_group_id],
            availability_zone="cn-north-4a")
        attachments = []
        for range in [{"value": i} for i in range(0, 2)]:
            attachments.append(huaweicloud.ecs.VolumeAttach(f"attachments-{range['value']}",
                instance_id=myinstance.id,
                volume_id=[__item.id for __item in myvol][range["value"]]))
        pulumi.export("volume devices", [__item.device for __item in attachments])
        ```

        ## Import

        Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g.

        ```sh
         $ pulumi import huaweicloud:Ecs/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param VolumeAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeAttachArgs.__new__(VolumeAttachArgs)

            __props__.__dict__["device"] = device
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            if volume_id is None and not opts.urn:
                raise TypeError("Missing required property 'volume_id'")
            __props__.__dict__["volume_id"] = volume_id
            __props__.__dict__["pci_address"] = None
        super(VolumeAttach, __self__).__init__(
            'huaweicloud:Ecs/volumeAttach:VolumeAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            pci_address: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            volume_id: Optional[pulumi.Input[str]] = None) -> 'VolumeAttach':
        """
        Get an existing VolumeAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Specifies the device of the volume attachment (ex: `/dev/vdc`).
        :param pulumi.Input[str] instance_id: Specifies the ID of the Instance to attach the Volume to.
        :param pulumi.Input[str] pci_address: PCI address of the block device.
        :param pulumi.Input[str] region: Specifies the region in which to create the volume resource. If omitted, the
               provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[str] volume_id: Specifies the ID of the Volume to attach to an Instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeAttachState.__new__(_VolumeAttachState)

        __props__.__dict__["device"] = device
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["pci_address"] = pci_address
        __props__.__dict__["region"] = region
        __props__.__dict__["volume_id"] = volume_id
        return VolumeAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        """
        Specifies the device of the volume attachment (ex: `/dev/vdc`).
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the Instance to attach the Volume to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="pciAddress")
    def pci_address(self) -> pulumi.Output[str]:
        """
        PCI address of the block device.
        """
        return pulumi.get(self, "pci_address")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the volume resource. If omitted, the
        provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the Volume to attach to an Instance.
        """
        return pulumi.get(self, "volume_id")

