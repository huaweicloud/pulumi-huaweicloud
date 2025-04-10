# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['L7policyArgs', 'L7policy']

@pulumi.input_type
class L7policyArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 listener_id: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 redirect_listener_id: Optional[pulumi.Input[str]] = None,
                 redirect_pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a L7policy resource.
        :param pulumi.Input[str] action: Specifies whether requests are forwarded to another backend server group
               or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
               + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
               + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
               HTTPS listener specified by `redirect_listener_id`.
        :param pulumi.Input[str] listener_id: Specifies the ID of the listener for which the forwarding policy is added.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] description: Human-readable description for the L7 Policy.
        :param pulumi.Input[str] name: Human-readable name for the L7 Policy. Does not have to be unique.
        :param pulumi.Input[int] position: The position of this policy on the listener. Positions start at 1.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] redirect_listener_id: Specifies the ID of the listener to which the traffic is redirected.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
               following requirements:
               + Can only be an HTTPS listener.
               + Can only be a listener of the same load balancer.
        :param pulumi.Input[str] redirect_pool_id: Specifies the ID of the backend server group to which traffic is forwarded.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
               following requirements:
               + Cannot be the default backend server group of the listener.
               + Cannot be the backend server group used by forwarding policies of other listeners.
        :param pulumi.Input[str] region: The region in which to create the L7 Policy resource. If omitted, the
               provider-level region will be used. Changing this creates a new L7 Policy.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "listener_id", listener_id)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if redirect_listener_id is not None:
            pulumi.set(__self__, "redirect_listener_id", redirect_listener_id)
        if redirect_pool_id is not None:
            pulumi.set(__self__, "redirect_pool_id", redirect_pool_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Specifies whether requests are forwarded to another backend server group
        or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
        + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
        + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
        HTTPS listener specified by `redirect_listener_id`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the listener for which the forwarding policy is added.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable description for the L7 Policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the L7 Policy. Does not have to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        The position of this policy on the listener. Positions start at 1.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="redirectListenerId")
    def redirect_listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the listener to which the traffic is redirected.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
        following requirements:
        + Can only be an HTTPS listener.
        + Can only be a listener of the same load balancer.
        """
        return pulumi.get(self, "redirect_listener_id")

    @redirect_listener_id.setter
    def redirect_listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_listener_id", value)

    @property
    @pulumi.getter(name="redirectPoolId")
    def redirect_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the backend server group to which traffic is forwarded.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
        following requirements:
        + Cannot be the default backend server group of the listener.
        + Cannot be the backend server group used by forwarding policies of other listeners.
        """
        return pulumi.get(self, "redirect_pool_id")

    @redirect_pool_id.setter
    def redirect_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_pool_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the L7 Policy resource. If omitted, the
        provider-level region will be used. Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _L7policyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 redirect_listener_id: Optional[pulumi.Input[str]] = None,
                 redirect_pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering L7policy resources.
        :param pulumi.Input[str] action: Specifies whether requests are forwarded to another backend server group
               or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
               + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
               + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
               HTTPS listener specified by `redirect_listener_id`.
        :param pulumi.Input[str] description: Human-readable description for the L7 Policy.
        :param pulumi.Input[str] listener_id: Specifies the ID of the listener for which the forwarding policy is added.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] name: Human-readable name for the L7 Policy. Does not have to be unique.
        :param pulumi.Input[int] position: The position of this policy on the listener. Positions start at 1.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] redirect_listener_id: Specifies the ID of the listener to which the traffic is redirected.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
               following requirements:
               + Can only be an HTTPS listener.
               + Can only be a listener of the same load balancer.
        :param pulumi.Input[str] redirect_pool_id: Specifies the ID of the backend server group to which traffic is forwarded.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
               following requirements:
               + Cannot be the default backend server group of the listener.
               + Cannot be the backend server group used by forwarding policies of other listeners.
        :param pulumi.Input[str] region: The region in which to create the L7 Policy resource. If omitted, the
               provider-level region will be used. Changing this creates a new L7 Policy.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if redirect_listener_id is not None:
            pulumi.set(__self__, "redirect_listener_id", redirect_listener_id)
        if redirect_pool_id is not None:
            pulumi.set(__self__, "redirect_pool_id", redirect_pool_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
            pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether requests are forwarded to another backend server group
        or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
        + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
        + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
        HTTPS listener specified by `redirect_listener_id`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable description for the L7 Policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the listener for which the forwarding policy is added.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the L7 Policy. Does not have to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        The position of this policy on the listener. Positions start at 1.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="redirectListenerId")
    def redirect_listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the listener to which the traffic is redirected.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
        following requirements:
        + Can only be an HTTPS listener.
        + Can only be a listener of the same load balancer.
        """
        return pulumi.get(self, "redirect_listener_id")

    @redirect_listener_id.setter
    def redirect_listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_listener_id", value)

    @property
    @pulumi.getter(name="redirectPoolId")
    def redirect_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the backend server group to which traffic is forwarded.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
        following requirements:
        + Cannot be the default backend server group of the listener.
        + Cannot be the backend server group used by forwarding policies of other listeners.
        """
        return pulumi.get(self, "redirect_pool_id")

    @redirect_pool_id.setter
    def redirect_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_pool_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the L7 Policy resource. If omitted, the
        provider-level region will be used. Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class L7policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 redirect_listener_id: Optional[pulumi.Input[str]] = None,
                 redirect_pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an ELB L7 Policy resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        loadbalancer1 = huaweicloud.elb.Loadbalancer("loadbalancer1", vip_subnet_id=var["subnet_id"])
        listener1 = huaweicloud.elb.Listener("listener1",
            protocol="HTTP",
            protocol_port=8080,
            loadbalancer_id=loadbalancer1.id)
        pool1 = huaweicloud.elb.Pool("pool1",
            protocol="HTTP",
            lb_method="ROUND_ROBIN",
            loadbalancer_id=loadbalancer1.id)
        l7policy1 = huaweicloud.elb.L7policy("l7policy1",
            action="REDIRECT_TO_POOL",
            description="test l7 policy",
            position=1,
            listener_id=listener1.id,
            redirect_pool_id=pool1.id)
        ```

        ## Import

        Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Elb/l7policy:L7policy l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies whether requests are forwarded to another backend server group
               or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
               + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
               + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
               HTTPS listener specified by `redirect_listener_id`.
        :param pulumi.Input[str] description: Human-readable description for the L7 Policy.
        :param pulumi.Input[str] listener_id: Specifies the ID of the listener for which the forwarding policy is added.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] name: Human-readable name for the L7 Policy. Does not have to be unique.
        :param pulumi.Input[int] position: The position of this policy on the listener. Positions start at 1.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] redirect_listener_id: Specifies the ID of the listener to which the traffic is redirected.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
               following requirements:
               + Can only be an HTTPS listener.
               + Can only be a listener of the same load balancer.
        :param pulumi.Input[str] redirect_pool_id: Specifies the ID of the backend server group to which traffic is forwarded.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
               following requirements:
               + Cannot be the default backend server group of the listener.
               + Cannot be the backend server group used by forwarding policies of other listeners.
        :param pulumi.Input[str] region: The region in which to create the L7 Policy resource. If omitted, the
               provider-level region will be used. Changing this creates a new L7 Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: L7policyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an ELB L7 Policy resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        loadbalancer1 = huaweicloud.elb.Loadbalancer("loadbalancer1", vip_subnet_id=var["subnet_id"])
        listener1 = huaweicloud.elb.Listener("listener1",
            protocol="HTTP",
            protocol_port=8080,
            loadbalancer_id=loadbalancer1.id)
        pool1 = huaweicloud.elb.Pool("pool1",
            protocol="HTTP",
            lb_method="ROUND_ROBIN",
            loadbalancer_id=loadbalancer1.id)
        l7policy1 = huaweicloud.elb.L7policy("l7policy1",
            action="REDIRECT_TO_POOL",
            description="test l7 policy",
            position=1,
            listener_id=listener1.id,
            redirect_pool_id=pool1.id)
        ```

        ## Import

        Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.bash

        ```sh
         $ pulumi import huaweicloud:Elb/l7policy:L7policy l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
        ```

        :param str resource_name: The name of the resource.
        :param L7policyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(L7policyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 redirect_listener_id: Optional[pulumi.Input[str]] = None,
                 redirect_pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = L7policyArgs.__new__(L7policyArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["description"] = description
            if listener_id is None and not opts.urn:
                raise TypeError("Missing required property 'listener_id'")
            __props__.__dict__["listener_id"] = listener_id
            __props__.__dict__["name"] = name
            __props__.__dict__["position"] = position
            __props__.__dict__["redirect_listener_id"] = redirect_listener_id
            __props__.__dict__["redirect_pool_id"] = redirect_pool_id
            __props__.__dict__["region"] = region
            if tenant_id is not None and not opts.urn:
                warnings.warn("""tenant_id is deprecated""", DeprecationWarning)
                pulumi.log.warn("""tenant_id is deprecated: tenant_id is deprecated""")
            __props__.__dict__["tenant_id"] = tenant_id
        super(L7policy, __self__).__init__(
            'huaweicloud:Elb/l7policy:L7policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            position: Optional[pulumi.Input[int]] = None,
            redirect_listener_id: Optional[pulumi.Input[str]] = None,
            redirect_pool_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'L7policy':
        """
        Get an existing L7policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies whether requests are forwarded to another backend server group
               or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
               + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
               + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
               HTTPS listener specified by `redirect_listener_id`.
        :param pulumi.Input[str] description: Human-readable description for the L7 Policy.
        :param pulumi.Input[str] listener_id: Specifies the ID of the listener for which the forwarding policy is added.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] name: Human-readable name for the L7 Policy. Does not have to be unique.
        :param pulumi.Input[int] position: The position of this policy on the listener. Positions start at 1.
               Changing this creates a new L7 Policy.
        :param pulumi.Input[str] redirect_listener_id: Specifies the ID of the listener to which the traffic is redirected.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
               following requirements:
               + Can only be an HTTPS listener.
               + Can only be a listener of the same load balancer.
        :param pulumi.Input[str] redirect_pool_id: Specifies the ID of the backend server group to which traffic is forwarded.
               This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
               following requirements:
               + Cannot be the default backend server group of the listener.
               + Cannot be the backend server group used by forwarding policies of other listeners.
        :param pulumi.Input[str] region: The region in which to create the L7 Policy resource. If omitted, the
               provider-level region will be used. Changing this creates a new L7 Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _L7policyState.__new__(_L7policyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["description"] = description
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["name"] = name
        __props__.__dict__["position"] = position
        __props__.__dict__["redirect_listener_id"] = redirect_listener_id
        __props__.__dict__["redirect_pool_id"] = redirect_pool_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        return L7policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Specifies whether requests are forwarded to another backend server group
        or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
        + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirect_pool_id`.
        + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listener_id` to the
        HTTPS listener specified by `redirect_listener_id`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-readable description for the L7 Policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the listener for which the forwarding policy is added.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human-readable name for the L7 Policy. Does not have to be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[int]:
        """
        The position of this policy on the listener. Positions start at 1.
        Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="redirectListenerId")
    def redirect_listener_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of the listener to which the traffic is redirected.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
        following requirements:
        + Can only be an HTTPS listener.
        + Can only be a listener of the same load balancer.
        """
        return pulumi.get(self, "redirect_listener_id")

    @property
    @pulumi.getter(name="redirectPoolId")
    def redirect_pool_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of the backend server group to which traffic is forwarded.
        This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
        following requirements:
        + Cannot be the default backend server group of the listener.
        + Cannot be the backend server group used by forwarding policies of other listeners.
        """
        return pulumi.get(self, "redirect_pool_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the L7 Policy resource. If omitted, the
        provider-level region will be used. Changing this creates a new L7 Policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tenant_id")

