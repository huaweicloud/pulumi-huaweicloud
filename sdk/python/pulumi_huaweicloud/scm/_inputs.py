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
    'CertificateAuthentificationArgs',
    'CertificateTargetArgs',
]

@pulumi.input_type
class CertificateAuthentificationArgs:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 record_name: Optional[pulumi.Input[str]] = None,
                 record_type: Optional[pulumi.Input[str]] = None,
                 record_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] domain: Domain name mapping to the verification value
        :param pulumi.Input[str] record_name: Name of a domain ownership verification value.
        :param pulumi.Input[str] record_type: Type of the domain name verification value.
        :param pulumi.Input[str] record_value: Domain verification value.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if record_name is not None:
            pulumi.set(__self__, "record_name", record_name)
        if record_type is not None:
            pulumi.set(__self__, "record_type", record_type)
        if record_value is not None:
            pulumi.set(__self__, "record_value", record_value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name mapping to the verification value
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="recordName")
    def record_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a domain ownership verification value.
        """
        return pulumi.get(self, "record_name")

    @record_name.setter
    def record_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "record_name", value)

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the domain name verification value.
        """
        return pulumi.get(self, "record_type")

    @record_type.setter
    def record_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "record_type", value)

    @property
    @pulumi.getter(name="recordValue")
    def record_value(self) -> Optional[pulumi.Input[str]]:
        """
        Domain verification value.
        """
        return pulumi.get(self, "record_value")

    @record_value.setter
    def record_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "record_value", value)


@pulumi.input_type
class CertificateTargetArgs:
    def __init__(__self__, *,
                 service: pulumi.Input[str],
                 projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] service: Service to which the certificate is pushed. The options include `CDN`,`WAF`
               and `Enhance_ELB`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] projects: The project where the service you want to push a certificate to. The same certificate
               can be pushed repeatedly to the same WAF or ELB service in the same `project`, but the CDN service can only be pushed
               once.
        """
        pulumi.set(__self__, "service", service)
        if projects is not None:
            pulumi.set(__self__, "projects", projects)

    @property
    @pulumi.getter
    def service(self) -> pulumi.Input[str]:
        """
        Service to which the certificate is pushed. The options include `CDN`,`WAF`
        and `Enhance_ELB`.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: pulumi.Input[str]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def projects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The project where the service you want to push a certificate to. The same certificate
        can be pushed repeatedly to the same WAF or ELB service in the same `project`, but the CDN service can only be pushed
        once.
        """
        return pulumi.get(self, "projects")

    @projects.setter
    def projects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "projects", value)

