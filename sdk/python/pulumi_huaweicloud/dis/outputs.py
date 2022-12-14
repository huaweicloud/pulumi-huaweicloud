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
    'StreamPartition',
]

@pulumi.output_type
class StreamPartition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hashRange":
            suggest = "hash_range"
        elif key == "sequenceNumberRange":
            suggest = "sequence_number_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamPartition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamPartition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamPartition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hash_range: Optional[str] = None,
                 id: Optional[str] = None,
                 sequence_number_range: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str hash_range: Possible value range of the hash key used by each partition.
        :param str id: The ID of the partition.
        :param str sequence_number_range: Sequence number range of each partition.
        :param str status: The status of the partition.
        """
        if hash_range is not None:
            pulumi.set(__self__, "hash_range", hash_range)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if sequence_number_range is not None:
            pulumi.set(__self__, "sequence_number_range", sequence_number_range)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="hashRange")
    def hash_range(self) -> Optional[str]:
        """
        Possible value range of the hash key used by each partition.
        """
        return pulumi.get(self, "hash_range")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the partition.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="sequenceNumberRange")
    def sequence_number_range(self) -> Optional[str]:
        """
        Sequence number range of each partition.
        """
        return pulumi.get(self, "sequence_number_range")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the partition.
        """
        return pulumi.get(self, "status")


