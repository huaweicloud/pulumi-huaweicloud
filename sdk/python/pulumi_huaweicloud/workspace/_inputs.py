# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DesktopDataVolumeArgs',
    'DesktopNicArgs',
    'DesktopRootVolumeArgs',
    'ServiceAdDomainArgs',
    'ServiceDesktopSecurityGroupArgs',
    'ServiceInfrastructureSecurityGroupArgs',
]

@pulumi.input_type
class DesktopDataVolumeArgs:
    def __init__(__self__, *,
                 size: pulumi.Input[int],
                 type: pulumi.Input[str],
                 created_at: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] size: Specifies the size of system volume, in GB.
               + For root volume, the valid value is range from `80` to `1,020`.
               + For data volume, the valid value is range from `10` to `8,200`.
        :param pulumi.Input[str] type: Specifies the type of system volume.
               The valid values are as follows:
               + **SAS**: High I/O disk type.
               + **SSD**: Ultra-high I/O disk type.
        :param pulumi.Input[str] created_at: The time that the volume was created.
        :param pulumi.Input[str] device: The device location to which the volume is attached.
        :param pulumi.Input[str] id: The volume ID.
        :param pulumi.Input[str] name: Specifies the desktop name.
               The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
               The name must start with a letter or digit and cannot end with a hyphen.
               Changing this will create a new resource.
        """
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "type", type)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        Specifies the size of system volume, in GB.
        + For root volume, the valid value is range from `80` to `1,020`.
        + For data volume, the valid value is range from `10` to `8,200`.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of system volume.
        The valid values are as follows:
        + **SAS**: High I/O disk type.
        + **SSD**: Ultra-high I/O disk type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the volume was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        The device location to which the volume is attached.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the desktop name.
        The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
        The name must start with a letter or digit and cannot end with a hyphen.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DesktopNicArgs:
    def __init__(__self__, *,
                 network_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] network_id: Specifies the network ID of subnet resource.
               Changing this will create a new resource.
        """
        pulumi.set(__self__, "network_id", network_id)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        Specifies the network ID of subnet resource.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)


@pulumi.input_type
class DesktopRootVolumeArgs:
    def __init__(__self__, *,
                 size: pulumi.Input[int],
                 type: pulumi.Input[str],
                 created_at: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] size: Specifies the size of system volume, in GB.
               + For root volume, the valid value is range from `80` to `1,020`.
               + For data volume, the valid value is range from `10` to `8,200`.
        :param pulumi.Input[str] type: Specifies the type of system volume.
               The valid values are as follows:
               + **SAS**: High I/O disk type.
               + **SSD**: Ultra-high I/O disk type.
        :param pulumi.Input[str] created_at: The time that the volume was created.
        :param pulumi.Input[str] device: The device location to which the volume is attached.
        :param pulumi.Input[str] id: The volume ID.
        :param pulumi.Input[str] name: Specifies the desktop name.
               The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
               The name must start with a letter or digit and cannot end with a hyphen.
               Changing this will create a new resource.
        """
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "type", type)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        Specifies the size of system volume, in GB.
        + For root volume, the valid value is range from `80` to `1,020`.
        + For data volume, the valid value is range from `10` to `8,200`.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of system volume.
        The valid values are as follows:
        + **SAS**: High I/O disk type.
        + **SSD**: Ultra-high I/O disk type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the volume was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        The device location to which the volume is attached.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the desktop name.
        The name can contain `1` to `15` characters, only letters, digits and hyphens (-) are allowed.
        The name must start with a letter or digit and cannot end with a hyphen.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ServiceAdDomainArgs:
    def __init__(__self__, *,
                 active_domain_ip: pulumi.Input[str],
                 active_domain_name: pulumi.Input[str],
                 admin_account: pulumi.Input[str],
                 name: pulumi.Input[str],
                 password: pulumi.Input[str],
                 active_dns_ip: Optional[pulumi.Input[str]] = None,
                 delete_computer_object: Optional[pulumi.Input[bool]] = None,
                 standby_dns_ip: Optional[pulumi.Input[str]] = None,
                 standby_domain_ip: Optional[pulumi.Input[str]] = None,
                 standby_domain_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] active_domain_ip: Specifies the IP address of primary domain controller.
        :param pulumi.Input[str] active_domain_name: Specifies the name of primary domain controller.
        :param pulumi.Input[str] admin_account: Specifies the domain administrator account.
               It must be an existing domain administrator account on the AD server.
        :param pulumi.Input[str] name: Specifies the domain name.
               The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        :param pulumi.Input[str] password: Specifies the account password of domain administrator.
        :param pulumi.Input[str] active_dns_ip: Specifies the primary DNS IP address.
        :param pulumi.Input[bool] delete_computer_object: Specifies whether to delete the corresponding computer object on AD
               while deleting the desktop.
        :param pulumi.Input[str] standby_dns_ip: Specifies the standby DNS IP address.
        :param pulumi.Input[str] standby_domain_ip: Specifies the IP address of the standby domain controller.
        :param pulumi.Input[str] standby_domain_name: Specifies the name of the standby domain controller.
        """
        pulumi.set(__self__, "active_domain_ip", active_domain_ip)
        pulumi.set(__self__, "active_domain_name", active_domain_name)
        pulumi.set(__self__, "admin_account", admin_account)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "password", password)
        if active_dns_ip is not None:
            pulumi.set(__self__, "active_dns_ip", active_dns_ip)
        if delete_computer_object is not None:
            pulumi.set(__self__, "delete_computer_object", delete_computer_object)
        if standby_dns_ip is not None:
            pulumi.set(__self__, "standby_dns_ip", standby_dns_ip)
        if standby_domain_ip is not None:
            pulumi.set(__self__, "standby_domain_ip", standby_domain_ip)
        if standby_domain_name is not None:
            pulumi.set(__self__, "standby_domain_name", standby_domain_name)

    @property
    @pulumi.getter(name="activeDomainIp")
    def active_domain_ip(self) -> pulumi.Input[str]:
        """
        Specifies the IP address of primary domain controller.
        """
        return pulumi.get(self, "active_domain_ip")

    @active_domain_ip.setter
    def active_domain_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "active_domain_ip", value)

    @property
    @pulumi.getter(name="activeDomainName")
    def active_domain_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of primary domain controller.
        """
        return pulumi.get(self, "active_domain_name")

    @active_domain_name.setter
    def active_domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "active_domain_name", value)

    @property
    @pulumi.getter(name="adminAccount")
    def admin_account(self) -> pulumi.Input[str]:
        """
        Specifies the domain administrator account.
        It must be an existing domain administrator account on the AD server.
        """
        return pulumi.get(self, "admin_account")

    @admin_account.setter
    def admin_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_account", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the domain name.
        The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Specifies the account password of domain administrator.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="activeDnsIp")
    def active_dns_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the primary DNS IP address.
        """
        return pulumi.get(self, "active_dns_ip")

    @active_dns_ip.setter
    def active_dns_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_dns_ip", value)

    @property
    @pulumi.getter(name="deleteComputerObject")
    def delete_computer_object(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to delete the corresponding computer object on AD
        while deleting the desktop.
        """
        return pulumi.get(self, "delete_computer_object")

    @delete_computer_object.setter
    def delete_computer_object(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_computer_object", value)

    @property
    @pulumi.getter(name="standbyDnsIp")
    def standby_dns_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the standby DNS IP address.
        """
        return pulumi.get(self, "standby_dns_ip")

    @standby_dns_ip.setter
    def standby_dns_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "standby_dns_ip", value)

    @property
    @pulumi.getter(name="standbyDomainIp")
    def standby_domain_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IP address of the standby domain controller.
        """
        return pulumi.get(self, "standby_domain_ip")

    @standby_domain_ip.setter
    def standby_domain_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "standby_domain_ip", value)

    @property
    @pulumi.getter(name="standbyDomainName")
    def standby_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the standby domain controller.
        """
        return pulumi.get(self, "standby_domain_name")

    @standby_domain_name.setter
    def standby_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "standby_domain_name", value)


@pulumi.input_type
class ServiceDesktopSecurityGroupArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: Security group ID.
        :param pulumi.Input[str] name: Specifies the domain name.
               The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Security group ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the domain name.
        The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ServiceInfrastructureSecurityGroupArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: Security group ID.
        :param pulumi.Input[str] name: Specifies the domain name.
               The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Security group ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the domain name.
        The domain name must be an existing domain name on the AD server, and the length cannot exceed `55`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

