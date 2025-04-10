# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AdminAssignmentArgs', 'AdminAssignment']

@pulumi.input_type
class AdminAssignmentArgs:
    def __init__(__self__, *,
                 account: pulumi.Input[str],
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 app_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AdminAssignment resource.
        :param pulumi.Input[str] account: Specifies the user account to be assigned the administrator role.
               The value can contain `1` to `64` characters.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_name: Specifies the (HUAWEI Cloud meeting) user account name to which the
               default administrator belongs. Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_password: Specifies the user password.
               Required if `account_name` is set. Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_id: Specifies the ID of the Third-party application.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_key: Specifies the Key information of the Third-party APP.
               Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "account", account)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if app_key is not None:
            pulumi.set(__self__, "app_key", app_key)

    @property
    @pulumi.getter
    def account(self) -> pulumi.Input[str]:
        """
        Specifies the user account to be assigned the administrator role.
        The value can contain `1` to `64` characters.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account")

    @account.setter
    def account(self, value: pulumi.Input[str]):
        pulumi.set(self, "account", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the (HUAWEI Cloud meeting) user account name to which the
        default administrator belongs. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user password.
        Required if `account_name` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the Third-party application.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Key information of the Third-party APP.
        Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_key")

    @app_key.setter
    def app_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_key", value)


@pulumi.input_type
class _AdminAssignmentState:
    def __init__(__self__, *,
                 account: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 app_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AdminAssignment resources.
        :param pulumi.Input[str] account: Specifies the user account to be assigned the administrator role.
               The value can contain `1` to `64` characters.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_name: Specifies the (HUAWEI Cloud meeting) user account name to which the
               default administrator belongs. Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_password: Specifies the user password.
               Required if `account_name` is set. Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_id: Specifies the ID of the Third-party application.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_key: Specifies the Key information of the Third-party APP.
               Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        if account is not None:
            pulumi.set(__self__, "account", account)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if app_key is not None:
            pulumi.set(__self__, "app_key", app_key)

    @property
    @pulumi.getter
    def account(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user account to be assigned the administrator role.
        The value can contain `1` to `64` characters.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account")

    @account.setter
    def account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the (HUAWEI Cloud meeting) user account name to which the
        default administrator belongs. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user password.
        Required if `account_name` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the Third-party application.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Key information of the Third-party APP.
        Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_key")

    @app_key.setter
    def app_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_key", value)


class AdminAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 app_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Using this resource to assign an administrator role to a user within HuaweiCloud.

        ## Example Usage
        ### Assign an administrator role to a user

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        app_id = config.require_object("appId")
        app_key = config.require_object("appKey")
        user_account = config.require_object("userAccount")
        test = huaweicloud.meeting.AdminAssignment("test",
            app_id=app_id,
            app_key=app_key,
            account=user_account)
        ```

        ## Import

        The assignment relationships can be imported using their `id` and authorization parameters, separated by slashes, e.g. Import an administrator assignment and authenticated by account. bash

        ```sh
         $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test <id>/<account_name>/<account_password>
        ```

         Import an administrator assignment and authenticated by `APP ID`/`APP Key`. bash

        ```sh
         $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test <id>/<app_id>/<app_key>/<corp_id>/<user_id>
        ```

         For this resource, the `corp_id` and `user_id` are never used, you can omit them but the slashes cannot be missing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account: Specifies the user account to be assigned the administrator role.
               The value can contain `1` to `64` characters.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_name: Specifies the (HUAWEI Cloud meeting) user account name to which the
               default administrator belongs. Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_password: Specifies the user password.
               Required if `account_name` is set. Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_id: Specifies the ID of the Third-party application.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_key: Specifies the Key information of the Third-party APP.
               Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AdminAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Using this resource to assign an administrator role to a user within HuaweiCloud.

        ## Example Usage
        ### Assign an administrator role to a user

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        app_id = config.require_object("appId")
        app_key = config.require_object("appKey")
        user_account = config.require_object("userAccount")
        test = huaweicloud.meeting.AdminAssignment("test",
            app_id=app_id,
            app_key=app_key,
            account=user_account)
        ```

        ## Import

        The assignment relationships can be imported using their `id` and authorization parameters, separated by slashes, e.g. Import an administrator assignment and authenticated by account. bash

        ```sh
         $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test <id>/<account_name>/<account_password>
        ```

         Import an administrator assignment and authenticated by `APP ID`/`APP Key`. bash

        ```sh
         $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test <id>/<app_id>/<app_key>/<corp_id>/<user_id>
        ```

         For this resource, the `corp_id` and `user_id` are never used, you can omit them but the slashes cannot be missing.

        :param str resource_name: The name of the resource.
        :param AdminAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AdminAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 app_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AdminAssignmentArgs.__new__(AdminAssignmentArgs)

            if account is None and not opts.urn:
                raise TypeError("Missing required property 'account'")
            __props__.__dict__["account"] = account
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["account_password"] = account_password
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["app_key"] = app_key
        super(AdminAssignment, __self__).__init__(
            'huaweicloud:Meeting/adminAssignment:AdminAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account: Optional[pulumi.Input[str]] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            account_password: Optional[pulumi.Input[str]] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            app_key: Optional[pulumi.Input[str]] = None) -> 'AdminAssignment':
        """
        Get an existing AdminAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account: Specifies the user account to be assigned the administrator role.
               The value can contain `1` to `64` characters.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_name: Specifies the (HUAWEI Cloud meeting) user account name to which the
               default administrator belongs. Changing this parameter will create a new resource.
        :param pulumi.Input[str] account_password: Specifies the user password.
               Required if `account_name` is set. Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_id: Specifies the ID of the Third-party application.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] app_key: Specifies the Key information of the Third-party APP.
               Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AdminAssignmentState.__new__(_AdminAssignmentState)

        __props__.__dict__["account"] = account
        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["account_password"] = account_password
        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["app_key"] = app_key
        return AdminAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def account(self) -> pulumi.Output[str]:
        """
        Specifies the user account to be assigned the administrator role.
        The value can contain `1` to `64` characters.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the (HUAWEI Cloud meeting) user account name to which the
        default administrator belongs. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the user password.
        Required if `account_name` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "account_password")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of the Third-party application.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Key information of the Third-party APP.
        Required if `app_id` is set. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "app_key")

