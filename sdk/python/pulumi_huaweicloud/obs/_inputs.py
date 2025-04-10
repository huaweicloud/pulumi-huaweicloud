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
    'BucketCorsRuleArgs',
    'BucketLifecycleRuleArgs',
    'BucketLifecycleRuleAbortIncompleteMultipartUploadArgs',
    'BucketLifecycleRuleExpirationArgs',
    'BucketLifecycleRuleNoncurrentVersionExpirationArgs',
    'BucketLifecycleRuleNoncurrentVersionTransitionArgs',
    'BucketLifecycleRuleTransitionArgs',
    'BucketLoggingArgs',
    'BucketStorageInfoArgs',
    'BucketWebsiteArgs',
]

@pulumi.input_type
class BucketCorsRuleArgs:
    def __init__(__self__, *,
                 allowed_methods: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_origins: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expose_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_age_seconds: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_methods: Specifies the acceptable operation type of buckets and objects. The methods
               include `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_origins: Requests from this origin can access the bucket. Multiple matching rules are
               allowed. One rule occupies one line, and allows one wildcard character (*) at most.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_headers: Specifies the allowed header of cross-origin requests. Only CORS requests
               matching the allowed header are valid.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] expose_headers: Specifies the exposed header in CORS responses, providing additional information
               for clients.
        :param pulumi.Input[int] max_age_seconds: Specifies the duration that your browser can cache CORS responses, expressed in
               seconds. The default value is 100.
        """
        pulumi.set(__self__, "allowed_methods", allowed_methods)
        pulumi.set(__self__, "allowed_origins", allowed_origins)
        if allowed_headers is not None:
            pulumi.set(__self__, "allowed_headers", allowed_headers)
        if expose_headers is not None:
            pulumi.set(__self__, "expose_headers", expose_headers)
        if max_age_seconds is not None:
            pulumi.set(__self__, "max_age_seconds", max_age_seconds)

    @property
    @pulumi.getter(name="allowedMethods")
    def allowed_methods(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies the acceptable operation type of buckets and objects. The methods
        include `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
        """
        return pulumi.get(self, "allowed_methods")

    @allowed_methods.setter
    def allowed_methods(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_methods", value)

    @property
    @pulumi.getter(name="allowedOrigins")
    def allowed_origins(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Requests from this origin can access the bucket. Multiple matching rules are
        allowed. One rule occupies one line, and allows one wildcard character (*) at most.
        """
        return pulumi.get(self, "allowed_origins")

    @allowed_origins.setter
    def allowed_origins(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_origins", value)

    @property
    @pulumi.getter(name="allowedHeaders")
    def allowed_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the allowed header of cross-origin requests. Only CORS requests
        matching the allowed header are valid.
        """
        return pulumi.get(self, "allowed_headers")

    @allowed_headers.setter
    def allowed_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_headers", value)

    @property
    @pulumi.getter(name="exposeHeaders")
    def expose_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the exposed header in CORS responses, providing additional information
        for clients.
        """
        return pulumi.get(self, "expose_headers")

    @expose_headers.setter
    def expose_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "expose_headers", value)

    @property
    @pulumi.getter(name="maxAgeSeconds")
    def max_age_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the duration that your browser can cache CORS responses, expressed in
        seconds. The default value is 100.
        """
        return pulumi.get(self, "max_age_seconds")

    @max_age_seconds.setter
    def max_age_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age_seconds", value)


@pulumi.input_type
class BucketLifecycleRuleArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 abort_incomplete_multipart_uploads: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleAbortIncompleteMultipartUploadArgs']]]] = None,
                 expirations: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleExpirationArgs']]]] = None,
                 noncurrent_version_expirations: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionExpirationArgs']]]] = None,
                 noncurrent_version_transitions: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionTransitionArgs']]]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]] = None):
        """
        :param pulumi.Input[bool] enabled: Specifies lifecycle rule status.
        :param pulumi.Input[str] name: Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleAbortIncompleteMultipartUploadArgs']]] abort_incomplete_multipart_uploads: Specifies a period when the not merged parts (fragments) in an
               incomplete upload are automatically deleted. (documented below).
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleExpirationArgs']]] expirations: Specifies a period when objects that have been last updated are automatically
               deleted. (documented below).
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionExpirationArgs']]] noncurrent_version_expirations: Specifies a period when noncurrent object versions are
               automatically deleted. (documented below).
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionTransitionArgs']]] noncurrent_version_transitions: Specifies a period when noncurrent object versions are
               automatically transitioned to `WARM` or `COLD` storage class (documented below).
        :param pulumi.Input[str] prefix: Object key prefix identifying one or more objects to which the rule applies. If omitted,
               all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start or end with a slash (/),
               cannot have consecutive slashes (/), and cannot contain the following special characters: \\:*?"<>|.
               When configuring multiple `lifecycle_rule`, field `prefix` in multiple `lifecycle_rule` cannot have an inclusive
               relationship.
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]] transitions: Specifies a period when objects that have been last updated are automatically
               transitioned to `WARM` or `COLD` storage class (documented below).
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "name", name)
        if abort_incomplete_multipart_uploads is not None:
            pulumi.set(__self__, "abort_incomplete_multipart_uploads", abort_incomplete_multipart_uploads)
        if expirations is not None:
            pulumi.set(__self__, "expirations", expirations)
        if noncurrent_version_expirations is not None:
            pulumi.set(__self__, "noncurrent_version_expirations", noncurrent_version_expirations)
        if noncurrent_version_transitions is not None:
            pulumi.set(__self__, "noncurrent_version_transitions", noncurrent_version_transitions)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Specifies lifecycle rule status.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="abortIncompleteMultipartUploads")
    def abort_incomplete_multipart_uploads(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleAbortIncompleteMultipartUploadArgs']]]]:
        """
        Specifies a period when the not merged parts (fragments) in an
        incomplete upload are automatically deleted. (documented below).
        """
        return pulumi.get(self, "abort_incomplete_multipart_uploads")

    @abort_incomplete_multipart_uploads.setter
    def abort_incomplete_multipart_uploads(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleAbortIncompleteMultipartUploadArgs']]]]):
        pulumi.set(self, "abort_incomplete_multipart_uploads", value)

    @property
    @pulumi.getter
    def expirations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleExpirationArgs']]]]:
        """
        Specifies a period when objects that have been last updated are automatically
        deleted. (documented below).
        """
        return pulumi.get(self, "expirations")

    @expirations.setter
    def expirations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleExpirationArgs']]]]):
        pulumi.set(self, "expirations", value)

    @property
    @pulumi.getter(name="noncurrentVersionExpirations")
    def noncurrent_version_expirations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionExpirationArgs']]]]:
        """
        Specifies a period when noncurrent object versions are
        automatically deleted. (documented below).
        """
        return pulumi.get(self, "noncurrent_version_expirations")

    @noncurrent_version_expirations.setter
    def noncurrent_version_expirations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionExpirationArgs']]]]):
        pulumi.set(self, "noncurrent_version_expirations", value)

    @property
    @pulumi.getter(name="noncurrentVersionTransitions")
    def noncurrent_version_transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionTransitionArgs']]]]:
        """
        Specifies a period when noncurrent object versions are
        automatically transitioned to `WARM` or `COLD` storage class (documented below).
        """
        return pulumi.get(self, "noncurrent_version_transitions")

    @noncurrent_version_transitions.setter
    def noncurrent_version_transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleNoncurrentVersionTransitionArgs']]]]):
        pulumi.set(self, "noncurrent_version_transitions", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Object key prefix identifying one or more objects to which the rule applies. If omitted,
        all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start or end with a slash (/),
        cannot have consecutive slashes (/), and cannot contain the following special characters: \\:*?"<>|.
        When configuring multiple `lifecycle_rule`, field `prefix` in multiple `lifecycle_rule` cannot have an inclusive
        relationship.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]]:
        """
        Specifies a period when objects that have been last updated are automatically
        transitioned to `WARM` or `COLD` storage class (documented below).
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]]):
        pulumi.set(self, "transitions", value)


@pulumi.input_type
class BucketLifecycleRuleAbortIncompleteMultipartUploadArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] days: Specifies the number of days since the initiation of an incomplete multipart upload that OBS
               will wait before deleting the not merged parts (fragments) of the upload.
               The valid value ranges from 1 to 2,147,483,647.
        """
        pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days since the initiation of an incomplete multipart upload that OBS
        will wait before deleting the not merged parts (fragments) of the upload.
        The valid value ranges from 1 to 2,147,483,647.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class BucketLifecycleRuleExpirationArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] days: Specifies the number of days when objects that have been last updated are automatically
               deleted. The expiration time must be greater than the transition times.
        """
        pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days when objects that have been last updated are automatically
        deleted. The expiration time must be greater than the transition times.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class BucketLifecycleRuleNoncurrentVersionExpirationArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] days: Specifies the number of days when noncurrent object versions are automatically deleted.
        """
        pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days when noncurrent object versions are automatically deleted.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class BucketLifecycleRuleNoncurrentVersionTransitionArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int],
                 storage_class: pulumi.Input[str]):
        """
        :param pulumi.Input[int] days: Specifies the number of days when noncurrent object versions are automatically transitioned
               to the specified storage class.
        :param pulumi.Input[str] storage_class: The class of storage used to store the object. Only `WARM` and `COLD` are
               supported.
        """
        pulumi.set(__self__, "days", days)
        pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days when noncurrent object versions are automatically transitioned
        to the specified storage class.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Input[str]:
        """
        The class of storage used to store the object. Only `WARM` and `COLD` are
        supported.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_class", value)


@pulumi.input_type
class BucketLifecycleRuleTransitionArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int],
                 storage_class: pulumi.Input[str]):
        """
        :param pulumi.Input[int] days: Specifies the number of days when objects that have been last updated are automatically
               transitioned to the specified storage class.
        :param pulumi.Input[str] storage_class: The class of storage used to store the object. Only `WARM` and `COLD` are
               supported.
        """
        pulumi.set(__self__, "days", days)
        pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days when objects that have been last updated are automatically
        transitioned to the specified storage class.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Input[str]:
        """
        The class of storage used to store the object. Only `WARM` and `COLD` are
        supported.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_class", value)


@pulumi.input_type
class BucketLoggingArgs:
    def __init__(__self__, *,
                 target_bucket: pulumi.Input[str],
                 agency: Optional[pulumi.Input[str]] = None,
                 target_prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] target_bucket: The name of the bucket that will receive the log objects. The acl policy of the
               target bucket should be `log-delivery-write`.
        :param pulumi.Input[str] agency: Specifies the IAM agency of OBS cloud service.
        :param pulumi.Input[str] target_prefix: To specify a key prefix for log objects.
        """
        pulumi.set(__self__, "target_bucket", target_bucket)
        if agency is not None:
            pulumi.set(__self__, "agency", agency)
        if target_prefix is not None:
            pulumi.set(__self__, "target_prefix", target_prefix)

    @property
    @pulumi.getter(name="targetBucket")
    def target_bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket that will receive the log objects. The acl policy of the
        target bucket should be `log-delivery-write`.
        """
        return pulumi.get(self, "target_bucket")

    @target_bucket.setter
    def target_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_bucket", value)

    @property
    @pulumi.getter
    def agency(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IAM agency of OBS cloud service.
        """
        return pulumi.get(self, "agency")

    @agency.setter
    def agency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agency", value)

    @property
    @pulumi.getter(name="targetPrefix")
    def target_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        To specify a key prefix for log objects.
        """
        return pulumi.get(self, "target_prefix")

    @target_prefix.setter
    def target_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_prefix", value)


@pulumi.input_type
class BucketStorageInfoArgs:
    def __init__(__self__, *,
                 object_number: Optional[pulumi.Input[int]] = None,
                 size: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] object_number: The number of objects stored in the bucket.
        :param pulumi.Input[int] size: The stored size of the bucket.
        """
        if object_number is not None:
            pulumi.set(__self__, "object_number", object_number)
        if size is not None:
            pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter(name="objectNumber")
    def object_number(self) -> Optional[pulumi.Input[int]]:
        """
        The number of objects stored in the bucket.
        """
        return pulumi.get(self, "object_number")

    @object_number.setter
    def object_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "object_number", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The stored size of the bucket.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)


@pulumi.input_type
class BucketWebsiteArgs:
    def __init__(__self__, *,
                 error_document: Optional[pulumi.Input[str]] = None,
                 index_document: Optional[pulumi.Input[str]] = None,
                 redirect_all_requests_to: Optional[pulumi.Input[str]] = None,
                 routing_rules: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] error_document: Specifies the error page returned when an error occurs during static website
               access. Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
        :param pulumi.Input[str] index_document: Unless using `redirect_all_requests_to`. Specifies the default homepage of the
               static website, only HTML web pages are supported. OBS only allows files such as `index.html` in the root directory of
               a bucket to function as the default homepage. That is to say, do not set the default homepage with a multi-level
               directory structure (for example, /page/index.html).
        :param pulumi.Input[str] redirect_all_requests_to: A hostname to redirect all website requests for this bucket to.
               Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The
               default is the protocol that is used in the original request.
        :param pulumi.Input[str] routing_rules: A JSON or XML format containing routing rules describing redirect behavior and
               when redirects are applied. Each rule contains a `Condition` and a `Redirect` as shown in the following table:
        """
        if error_document is not None:
            pulumi.set(__self__, "error_document", error_document)
        if index_document is not None:
            pulumi.set(__self__, "index_document", index_document)
        if redirect_all_requests_to is not None:
            pulumi.set(__self__, "redirect_all_requests_to", redirect_all_requests_to)
        if routing_rules is not None:
            pulumi.set(__self__, "routing_rules", routing_rules)

    @property
    @pulumi.getter(name="errorDocument")
    def error_document(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the error page returned when an error occurs during static website
        access. Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
        """
        return pulumi.get(self, "error_document")

    @error_document.setter
    def error_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "error_document", value)

    @property
    @pulumi.getter(name="indexDocument")
    def index_document(self) -> Optional[pulumi.Input[str]]:
        """
        Unless using `redirect_all_requests_to`. Specifies the default homepage of the
        static website, only HTML web pages are supported. OBS only allows files such as `index.html` in the root directory of
        a bucket to function as the default homepage. That is to say, do not set the default homepage with a multi-level
        directory structure (for example, /page/index.html).
        """
        return pulumi.get(self, "index_document")

    @index_document.setter
    def index_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index_document", value)

    @property
    @pulumi.getter(name="redirectAllRequestsTo")
    def redirect_all_requests_to(self) -> Optional[pulumi.Input[str]]:
        """
        A hostname to redirect all website requests for this bucket to.
        Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The
        default is the protocol that is used in the original request.
        """
        return pulumi.get(self, "redirect_all_requests_to")

    @redirect_all_requests_to.setter
    def redirect_all_requests_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_all_requests_to", value)

    @property
    @pulumi.getter(name="routingRules")
    def routing_rules(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON or XML format containing routing rules describing redirect behavior and
        when redirects are applied. Each rule contains a `Condition` and a `Redirect` as shown in the following table:
        """
        return pulumi.get(self, "routing_rules")

    @routing_rules.setter
    def routing_rules(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_rules", value)


