# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeypairArgs', 'Keypair']

@pulumi.input_type
class KeypairArgs:
    def __init__(__self__, *,
                 key_file: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Keypair resource.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               By default, the private key file will be created in the same folder as the current script file.
               If you need to create it in another folder, please specify the path for `key_file`.
               Changing this creates a new keypair.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. Changing this creates a new keypair.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key. Changing this creates
               a new keypair.
               This parameter and `key_file` are alternative.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this creates a new keypair.
        """
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        By default, the private key file will be created in the same folder as the current script file.
        If you need to create it in another folder, please specify the path for `key_file`.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the keypair. Changing this creates a new keypair.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the imported OpenSSH-formatted public key. Changing this creates
        a new keypair.
        This parameter and `key_file` are alternative.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _KeypairState:
    def __init__(__self__, *,
                 key_file: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Keypair resources.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               By default, the private key file will be created in the same folder as the current script file.
               If you need to create it in another folder, please specify the path for `key_file`.
               Changing this creates a new keypair.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. Changing this creates a new keypair.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key. Changing this creates
               a new keypair.
               This parameter and `key_file` are alternative.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this creates a new keypair.
        """
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        By default, the private key file will be created in the same folder as the current script file.
        If you need to create it in another folder, please specify the path for `key_file`.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique name for the keypair. Changing this creates a new keypair.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the imported OpenSSH-formatted public key. Changing this creates
        a new keypair.
        This parameter and `key_file` are alternative.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Keypair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        !> **WARNING:** It has been deprecated, use `Dew.Keypair` instead.

        Manages a keypair resource within HuaweiCloud.

        ## Example Usage
        ### Create a new keypair and export private key to current folder

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.ecs.Keypair("test-keypair", key_file="private_key.pem")
        ```
        ### Import an exist keypair

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.ecs.Keypair("test-keypair", public_key="ssh-rsa AAAAB3NzaC1yc2EAAAlJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcAT... root@terra-dev")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Ecs/keypair:Keypair my-keypair test-keypair
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               By default, the private key file will be created in the same folder as the current script file.
               If you need to create it in another folder, please specify the path for `key_file`.
               Changing this creates a new keypair.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. Changing this creates a new keypair.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key. Changing this creates
               a new keypair.
               This parameter and `key_file` are alternative.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this creates a new keypair.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeypairArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        !> **WARNING:** It has been deprecated, use `Dew.Keypair` instead.

        Manages a keypair resource within HuaweiCloud.

        ## Example Usage
        ### Create a new keypair and export private key to current folder

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.ecs.Keypair("test-keypair", key_file="private_key.pem")
        ```
        ### Import an exist keypair

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test_keypair = huaweicloud.ecs.Keypair("test-keypair", public_key="ssh-rsa AAAAB3NzaC1yc2EAAAlJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcAT... root@terra-dev")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Ecs/keypair:Keypair my-keypair test-keypair
        ```

        :param str resource_name: The name of the resource.
        :param KeypairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeypairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeypairArgs.__new__(KeypairArgs)

            __props__.__dict__["key_file"] = key_file
            __props__.__dict__["name"] = name
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["region"] = region
        super(Keypair, __self__).__init__(
            'huaweicloud:Ecs/keypair:Keypair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_file: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Keypair':
        """
        Get an existing Keypair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_file: Specifies the path of the created private key.
               The private key file (**.pem**) is created only after the resource is created.
               By default, the private key file will be created in the same folder as the current script file.
               If you need to create it in another folder, please specify the path for `key_file`.
               Changing this creates a new keypair.
        :param pulumi.Input[str] name: Specifies a unique name for the keypair. Changing this creates a new keypair.
        :param pulumi.Input[str] public_key: Specifies the imported OpenSSH-formatted public key. Changing this creates
               a new keypair.
               This parameter and `key_file` are alternative.
        :param pulumi.Input[str] region: Specifies the region in which to create the keypair resource. If omitted, the
               provider-level region will be used. Changing this creates a new keypair.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeypairState.__new__(_KeypairState)

        __props__.__dict__["key_file"] = key_file
        __props__.__dict__["name"] = name
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["region"] = region
        return Keypair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> pulumi.Output[str]:
        """
        Specifies the path of the created private key.
        The private key file (**.pem**) is created only after the resource is created.
        By default, the private key file will be created in the same folder as the current script file.
        If you need to create it in another folder, please specify the path for `key_file`.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "key_file")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies a unique name for the keypair. Changing this creates a new keypair.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        Specifies the imported OpenSSH-formatted public key. Changing this creates
        a new keypair.
        This parameter and `key_file` are alternative.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the keypair resource. If omitted, the
        provider-level region will be used. Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

