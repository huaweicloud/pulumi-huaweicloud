# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PvcArgs', 'Pvc']

@pulumi.input_type
class PvcArgs:
    def __init__(__self__, *,
                 access_modes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cluster_id: pulumi.Input[str],
                 namespace: pulumi.Input[str],
                 storage: pulumi.Input[str],
                 storage_class_name: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Pvc resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_modes: Specifies the desired access modes the volume should have.
               The valid values are as follows:
               + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
               + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
               + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID to which the CCE PVC belongs.
        :param pulumi.Input[str] namespace: Specifies the namespace to logically divide your containers into different
               group. Changing this will create a new PVC resource.
        :param pulumi.Input[str] storage: Specifies the minimum amount of storage resources required.
               Changing this creates a new PVC resource.
        :param pulumi.Input[str] storage_class_name: Specifies the type of the storage bound to the CCE PVC.
               The valid values are as follows:
               + **csi-disk**: EVS.
               + **csi-obs**: OBS.
               + **csi-nas**: SFS.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Specifies the unstructured key value map for external parameters.
               Changing this will create a new PVC resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Specifies the map of string keys and values for labels.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] name: Specifies the unique name of the PVC resource.  
               This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
               hyphens (-), and must start and end with lowercase letters and digits.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the PVC resource.
               If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        """
        pulumi.set(__self__, "access_modes", access_modes)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "storage", storage)
        pulumi.set(__self__, "storage_class_name", storage_class_name)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accessModes")
    def access_modes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies the desired access modes the volume should have.
        The valid values are as follows:
        + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
        + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
        + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        """
        return pulumi.get(self, "access_modes")

    @access_modes.setter
    def access_modes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "access_modes", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Specifies the cluster ID to which the CCE PVC belongs.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        Specifies the namespace to logically divide your containers into different
        group. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input[str]:
        """
        Specifies the minimum amount of storage resources required.
        Changing this creates a new PVC resource.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter(name="storageClassName")
    def storage_class_name(self) -> pulumi.Input[str]:
        """
        Specifies the type of the storage bound to the CCE PVC.
        The valid values are as follows:
        + **csi-disk**: EVS.
        + **csi-obs**: OBS.
        + **csi-nas**: SFS.
        """
        return pulumi.get(self, "storage_class_name")

    @storage_class_name.setter
    def storage_class_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_class_name", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the unstructured key value map for external parameters.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the map of string keys and values for labels.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the unique name of the PVC resource.  
        This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
        hyphens (-), and must start and end with lowercase letters and digits.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the PVC resource.
        If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _PvcState:
    def __init__(__self__, *,
                 access_modes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 storage_class_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pvc resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_modes: Specifies the desired access modes the volume should have.
               The valid values are as follows:
               + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
               + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
               + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Specifies the unstructured key value map for external parameters.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID to which the CCE PVC belongs.
        :param pulumi.Input[str] creation_timestamp: The server time when PVC was created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Specifies the map of string keys and values for labels.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] name: Specifies the unique name of the PVC resource.  
               This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
               hyphens (-), and must start and end with lowercase letters and digits.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] namespace: Specifies the namespace to logically divide your containers into different
               group. Changing this will create a new PVC resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the PVC resource.
               If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        :param pulumi.Input[str] status: The current phase of the PVC.
               + **Pending**: Not yet bound.
               + **Bound**: Already bound.
        :param pulumi.Input[str] storage: Specifies the minimum amount of storage resources required.
               Changing this creates a new PVC resource.
        :param pulumi.Input[str] storage_class_name: Specifies the type of the storage bound to the CCE PVC.
               The valid values are as follows:
               + **csi-disk**: EVS.
               + **csi-obs**: OBS.
               + **csi-nas**: SFS.
        """
        if access_modes is not None:
            pulumi.set(__self__, "access_modes", access_modes)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)
        if storage_class_name is not None:
            pulumi.set(__self__, "storage_class_name", storage_class_name)

    @property
    @pulumi.getter(name="accessModes")
    def access_modes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the desired access modes the volume should have.
        The valid values are as follows:
        + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
        + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
        + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        """
        return pulumi.get(self, "access_modes")

    @access_modes.setter
    def access_modes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "access_modes", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the unstructured key value map for external parameters.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the cluster ID to which the CCE PVC belongs.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        The server time when PVC was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the map of string keys and values for labels.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the unique name of the PVC resource.  
        This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
        hyphens (-), and must start and end with lowercase letters and digits.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the namespace to logically divide your containers into different
        group. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the PVC resource.
        If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The current phase of the PVC.
        + **Pending**: Not yet bound.
        + **Bound**: Already bound.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the minimum amount of storage resources required.
        Changing this creates a new PVC resource.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter(name="storageClassName")
    def storage_class_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of the storage bound to the CCE PVC.
        The valid values are as follows:
        + **csi-disk**: EVS.
        + **csi-obs**: OBS.
        + **csi-nas**: SFS.
        """
        return pulumi.get(self, "storage_class_name")

    @storage_class_name.setter
    def storage_class_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class_name", value)


class Pvc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_modes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 storage_class_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a CCE Persistent Volume Claim resource within HuaweiCloud.

        ## Example Usage
        ### Create PVC with EVS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            annotations={
                "everest.io/disk-volume-type": "SSD",
            },
            storage_class_name="csi-disk",
            access_modes=["ReadWriteOnce"],
            storage="10Gi")
        ```
        ### Create PVC with OBS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            annotations={
                "everest.io/obs-volume-type": "STANDARD",
                "csi.storage.k8s.io/fstype": "obsfs",
            },
            storage_class_name="csi-obs",
            access_modes=["ReadWriteMany"],
            storage="1Gi")
        ```
        ### Create PVC with SFS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            storage_class_name="csi-nas",
            access_modes=["ReadWriteMany"],
            storage="10Gi")
        ```

        ## Import

        CCE PVC can be imported using the cluster ID, namespace and ID separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Cce/pvc:Pvc test 5c20fdad-7288-11eb-b817-0255ac10158b/default/fa540f3b-12d9-40e5-8268-04bcfed95a46
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`annotations`. It is generally recommended running `terraform plan` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. hcl resource "huaweicloud_cce_pvc" "test" {

         ...

         lifecycle {

         ignore_changes = [

         annotations,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_modes: Specifies the desired access modes the volume should have.
               The valid values are as follows:
               + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
               + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
               + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Specifies the unstructured key value map for external parameters.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID to which the CCE PVC belongs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Specifies the map of string keys and values for labels.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] name: Specifies the unique name of the PVC resource.  
               This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
               hyphens (-), and must start and end with lowercase letters and digits.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] namespace: Specifies the namespace to logically divide your containers into different
               group. Changing this will create a new PVC resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the PVC resource.
               If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        :param pulumi.Input[str] storage: Specifies the minimum amount of storage resources required.
               Changing this creates a new PVC resource.
        :param pulumi.Input[str] storage_class_name: Specifies the type of the storage bound to the CCE PVC.
               The valid values are as follows:
               + **csi-disk**: EVS.
               + **csi-obs**: OBS.
               + **csi-nas**: SFS.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PvcArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a CCE Persistent Volume Claim resource within HuaweiCloud.

        ## Example Usage
        ### Create PVC with EVS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            annotations={
                "everest.io/disk-volume-type": "SSD",
            },
            storage_class_name="csi-disk",
            access_modes=["ReadWriteOnce"],
            storage="10Gi")
        ```
        ### Create PVC with OBS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            annotations={
                "everest.io/obs-volume-type": "STANDARD",
                "csi.storage.k8s.io/fstype": "obsfs",
            },
            storage_class_name="csi-obs",
            access_modes=["ReadWriteMany"],
            storage="1Gi")
        ```
        ### Create PVC with SFS

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        cluster_id = config.require_object("clusterId")
        namespace = config.require_object("namespace")
        pvc_name = config.require_object("pvcName")
        test = huaweicloud.cce.Pvc("test",
            cluster_id=cluster_id,
            namespace=namespace,
            storage_class_name="csi-nas",
            access_modes=["ReadWriteMany"],
            storage="10Gi")
        ```

        ## Import

        CCE PVC can be imported using the cluster ID, namespace and ID separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Cce/pvc:Pvc test 5c20fdad-7288-11eb-b817-0255ac10158b/default/fa540f3b-12d9-40e5-8268-04bcfed95a46
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`annotations`. It is generally recommended running `terraform plan` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. hcl resource "huaweicloud_cce_pvc" "test" {

         ...

         lifecycle {

         ignore_changes = [

         annotations,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param PvcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PvcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_modes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 storage_class_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PvcArgs.__new__(PvcArgs)

            if access_modes is None and not opts.urn:
                raise TypeError("Missing required property 'access_modes'")
            __props__.__dict__["access_modes"] = access_modes
            __props__.__dict__["annotations"] = annotations
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["region"] = region
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__.__dict__["storage"] = storage
            if storage_class_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_class_name'")
            __props__.__dict__["storage_class_name"] = storage_class_name
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["status"] = None
        super(Pvc, __self__).__init__(
            'huaweicloud:Cce/pvc:Pvc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_modes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            storage: Optional[pulumi.Input[str]] = None,
            storage_class_name: Optional[pulumi.Input[str]] = None) -> 'Pvc':
        """
        Get an existing Pvc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_modes: Specifies the desired access modes the volume should have.
               The valid values are as follows:
               + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
               + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
               + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Specifies the unstructured key value map for external parameters.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] cluster_id: Specifies the cluster ID to which the CCE PVC belongs.
        :param pulumi.Input[str] creation_timestamp: The server time when PVC was created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Specifies the map of string keys and values for labels.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] name: Specifies the unique name of the PVC resource.  
               This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
               hyphens (-), and must start and end with lowercase letters and digits.
               Changing this will create a new PVC resource.
        :param pulumi.Input[str] namespace: Specifies the namespace to logically divide your containers into different
               group. Changing this will create a new PVC resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the PVC resource.
               If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        :param pulumi.Input[str] status: The current phase of the PVC.
               + **Pending**: Not yet bound.
               + **Bound**: Already bound.
        :param pulumi.Input[str] storage: Specifies the minimum amount of storage resources required.
               Changing this creates a new PVC resource.
        :param pulumi.Input[str] storage_class_name: Specifies the type of the storage bound to the CCE PVC.
               The valid values are as follows:
               + **csi-disk**: EVS.
               + **csi-obs**: OBS.
               + **csi-nas**: SFS.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PvcState.__new__(_PvcState)

        __props__.__dict__["access_modes"] = access_modes
        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["storage"] = storage
        __props__.__dict__["storage_class_name"] = storage_class_name
        return Pvc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessModes")
    def access_modes(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the desired access modes the volume should have.
        The valid values are as follows:
        + **ReadWriteOnce**: The volume can be mounted as read-write by a single node.
        + **ReadOnlyMany**: The volume can be mounted as read-only by many nodes.
        + **ReadWriteMany**: The volume can be mounted as read-write by many nodes.
        """
        return pulumi.get(self, "access_modes")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the unstructured key value map for external parameters.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Specifies the cluster ID to which the CCE PVC belongs.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        The server time when PVC was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the map of string keys and values for labels.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the unique name of the PVC resource.  
        This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
        hyphens (-), and must start and end with lowercase letters and digits.
        Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        Specifies the namespace to logically divide your containers into different
        group. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the PVC resource.
        If omitted, the provider-level region will be used. Changing this will create a new PVC resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current phase of the PVC.
        + **Pending**: Not yet bound.
        + **Bound**: Already bound.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output[str]:
        """
        Specifies the minimum amount of storage resources required.
        Changing this creates a new PVC resource.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter(name="storageClassName")
    def storage_class_name(self) -> pulumi.Output[str]:
        """
        Specifies the type of the storage bound to the CCE PVC.
        The valid values are as follows:
        + **csi-disk**: EVS.
        + **csi-obs**: OBS.
        + **csi-nas**: SFS.
        """
        return pulumi.get(self, "storage_class_name")

