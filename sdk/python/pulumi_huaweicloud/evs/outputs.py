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

__all__ = [
    'VolumeAttachment',
    'GetVolumesVolumeResult',
    'GetVolumesVolumeAttachmentResult',
]

@pulumi.output_type
class VolumeAttachment(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "instanceId":
            suggest = "instance_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachment. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachment.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachment.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 device: Optional[str] = None,
                 id: Optional[str] = None,
                 instance_id: Optional[str] = None):
        """
        :param str device: The device name.
        :param str id: The ID of the attachment information.
        :param str instance_id: The ID of the server to which the disk is attached.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        The device name.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the attachment information.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        The ID of the server to which the disk is attached.
        """
        return pulumi.get(self, "instance_id")


@pulumi.output_type
class GetVolumesVolumeResult(dict):
    def __init__(__self__, *,
                 attachments: Sequence['outputs.GetVolumesVolumeAttachmentResult'],
                 availability_zone: str,
                 bootable: bool,
                 create_at: str,
                 description: str,
                 enterprise_project_id: str,
                 id: str,
                 name: str,
                 service_type: str,
                 shareable: bool,
                 size: int,
                 status: str,
                 tags: Mapping[str, str],
                 update_at: str,
                 wwn: str):
        """
        :param Sequence['GetVolumesVolumeAttachmentArgs'] attachments: The disk attachment information. Structure is documented below.
        :param str availability_zone: Specifies the availability zone for the disks.
        :param bool bootable: Whether the disk is bootable.
        :param str create_at: The time when the disk was created.
        :param str description: The disk description.
        :param str enterprise_project_id: Specifies the enterprise project ID for filtering.
        :param str id: The ID of the attached resource in UUID format.
        :param str name: The disk name.
        :param str service_type: The service type, such as EVS, DSS or DESS.
        :param bool shareable: Specifies whether the disk is shareable.
        :param int size: The disk size, in GB.
        :param str status: Specifies the disk status. The valid values are as following:
               + **FREEZED**
               + **BIND_ERROR**
               + **BINDING**
               + **PENDING_DELETE**
               + **PENDING_CREATE**
               + **NOTIFYING**
               + **NOTIFY_DELETE**
               + **PENDING_UPDATE**
               + **DOWN**
               + **ACTIVE**
               + **ELB**
               + **ERROR**
               + **VPN**
        :param Mapping[str, str] tags: Specifies the included key/value pairs which associated with the desired disk.
        :param str update_at: The time when the disk was updated.
        :param str wwn: The unique identifier used when attaching the disk.
        """
        pulumi.set(__self__, "attachments", attachments)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "bootable", bootable)
        pulumi.set(__self__, "create_at", create_at)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "service_type", service_type)
        pulumi.set(__self__, "shareable", shareable)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "update_at", update_at)
        pulumi.set(__self__, "wwn", wwn)

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.GetVolumesVolumeAttachmentResult']:
        """
        The disk attachment information. Structure is documented below.
        """
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Specifies the availability zone for the disks.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def bootable(self) -> bool:
        """
        Whether the disk is bootable.
        """
        return pulumi.get(self, "bootable")

    @property
    @pulumi.getter(name="createAt")
    def create_at(self) -> str:
        """
        The time when the disk was created.
        """
        return pulumi.get(self, "create_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The disk description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> str:
        """
        Specifies the enterprise project ID for filtering.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the attached resource in UUID format.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The disk name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        The service type, such as EVS, DSS or DESS.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def shareable(self) -> bool:
        """
        Specifies whether the disk is shareable.
        """
        return pulumi.get(self, "shareable")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The disk size, in GB.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Specifies the disk status. The valid values are as following:
        + **FREEZED**
        + **BIND_ERROR**
        + **BINDING**
        + **PENDING_DELETE**
        + **PENDING_CREATE**
        + **NOTIFYING**
        + **NOTIFY_DELETE**
        + **PENDING_UPDATE**
        + **DOWN**
        + **ACTIVE**
        + **ELB**
        + **ERROR**
        + **VPN**
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Specifies the included key/value pairs which associated with the desired disk.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateAt")
    def update_at(self) -> str:
        """
        The time when the disk was updated.
        """
        return pulumi.get(self, "update_at")

    @property
    @pulumi.getter
    def wwn(self) -> str:
        """
        The unique identifier used when attaching the disk.
        """
        return pulumi.get(self, "wwn")


@pulumi.output_type
class GetVolumesVolumeAttachmentResult(dict):
    def __init__(__self__, *,
                 attached_at: str,
                 attached_mode: str,
                 device_name: str,
                 id: str,
                 server_id: str):
        """
        :param str attached_at: The time when the disk was attached.
        :param str attached_mode: The ID of the attachment information.
        :param str device_name: The device name to which the disk is attached.
        :param str id: The ID of the attached resource in UUID format.
        :param str server_id: Specifies the server ID to which the disks are attached.
        """
        pulumi.set(__self__, "attached_at", attached_at)
        pulumi.set(__self__, "attached_mode", attached_mode)
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter(name="attachedAt")
    def attached_at(self) -> str:
        """
        The time when the disk was attached.
        """
        return pulumi.get(self, "attached_at")

    @property
    @pulumi.getter(name="attachedMode")
    def attached_mode(self) -> str:
        """
        The ID of the attachment information.
        """
        return pulumi.get(self, "attached_mode")

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> str:
        """
        The device name to which the disk is attached.
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the attached resource in UUID format.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> str:
        """
        Specifies the server ID to which the disks are attached.
        """
        return pulumi.get(self, "server_id")

