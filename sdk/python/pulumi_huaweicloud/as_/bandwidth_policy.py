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
from ._inputs import *

__all__ = ['BandwidthPolicyArgs', 'BandwidthPolicy']

@pulumi.input_type
class BandwidthPolicyArgs:
    def __init__(__self__, *,
                 bandwidth_id: pulumi.Input[str],
                 scaling_policy_name: pulumi.Input[str],
                 scaling_policy_type: pulumi.Input[str],
                 alarm_id: Optional[pulumi.Input[str]] = None,
                 cool_down_time: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scaling_policy_action: Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']] = None,
                 scheduled_policy: Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']] = None):
        """
        The set of arguments for constructing a BandwidthPolicy resource.
        :param pulumi.Input[str] bandwidth_id: Specifies the scaling bandwidth ID.
        :param pulumi.Input[str] scaling_policy_name: Specifies the AS policy name.
               The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        :param pulumi.Input[str] scaling_policy_type: Specifies the AS policy type. The options are as follows:
               - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
               - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
               - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        :param pulumi.Input[str] alarm_id: Specifies the alarm rule ID.
               This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        :param pulumi.Input[int] cool_down_time: Specifies the cooldown period (in seconds).
               The value ranges from 0 to 86400 and is 300 by default.
        :param pulumi.Input[str] description: Specifies the description of the AS policy.
               The value can contain 0 to 256 characters.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input['BandwidthPolicyScalingPolicyActionArgs'] scaling_policy_action: Specifies the scaling action of the AS policy.
               The object structure is documented below.
        :param pulumi.Input['BandwidthPolicyScheduledPolicyArgs'] scheduled_policy: Specifies the periodic or scheduled AS policy.
               This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
               The object structure is documented below.
        """
        pulumi.set(__self__, "bandwidth_id", bandwidth_id)
        pulumi.set(__self__, "scaling_policy_name", scaling_policy_name)
        pulumi.set(__self__, "scaling_policy_type", scaling_policy_type)
        if alarm_id is not None:
            pulumi.set(__self__, "alarm_id", alarm_id)
        if cool_down_time is not None:
            pulumi.set(__self__, "cool_down_time", cool_down_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scaling_policy_action is not None:
            pulumi.set(__self__, "scaling_policy_action", scaling_policy_action)
        if scheduled_policy is not None:
            pulumi.set(__self__, "scheduled_policy", scheduled_policy)

    @property
    @pulumi.getter(name="bandwidthId")
    def bandwidth_id(self) -> pulumi.Input[str]:
        """
        Specifies the scaling bandwidth ID.
        """
        return pulumi.get(self, "bandwidth_id")

    @bandwidth_id.setter
    def bandwidth_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth_id", value)

    @property
    @pulumi.getter(name="scalingPolicyName")
    def scaling_policy_name(self) -> pulumi.Input[str]:
        """
        Specifies the AS policy name.
        The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        """
        return pulumi.get(self, "scaling_policy_name")

    @scaling_policy_name.setter
    def scaling_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_policy_name", value)

    @property
    @pulumi.getter(name="scalingPolicyType")
    def scaling_policy_type(self) -> pulumi.Input[str]:
        """
        Specifies the AS policy type. The options are as follows:
        - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
        - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
        - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        """
        return pulumi.get(self, "scaling_policy_type")

    @scaling_policy_type.setter
    def scaling_policy_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_policy_type", value)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the alarm rule ID.
        This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        """
        return pulumi.get(self, "alarm_id")

    @alarm_id.setter
    def alarm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_id", value)

    @property
    @pulumi.getter(name="coolDownTime")
    def cool_down_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the cooldown period (in seconds).
        The value ranges from 0 to 86400 and is 300 by default.
        """
        return pulumi.get(self, "cool_down_time")

    @cool_down_time.setter
    def cool_down_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cool_down_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the AS policy.
        The value can contain 0 to 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="scalingPolicyAction")
    def scaling_policy_action(self) -> Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']]:
        """
        Specifies the scaling action of the AS policy.
        The object structure is documented below.
        """
        return pulumi.get(self, "scaling_policy_action")

    @scaling_policy_action.setter
    def scaling_policy_action(self, value: Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']]):
        pulumi.set(self, "scaling_policy_action", value)

    @property
    @pulumi.getter(name="scheduledPolicy")
    def scheduled_policy(self) -> Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']]:
        """
        Specifies the periodic or scheduled AS policy.
        This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
        The object structure is documented below.
        """
        return pulumi.get(self, "scheduled_policy")

    @scheduled_policy.setter
    def scheduled_policy(self, value: Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']]):
        pulumi.set(self, "scheduled_policy", value)


@pulumi.input_type
class _BandwidthPolicyState:
    def __init__(__self__, *,
                 alarm_id: Optional[pulumi.Input[str]] = None,
                 bandwidth_id: Optional[pulumi.Input[str]] = None,
                 cool_down_time: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scaling_policy_action: Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']] = None,
                 scaling_policy_name: Optional[pulumi.Input[str]] = None,
                 scaling_policy_type: Optional[pulumi.Input[str]] = None,
                 scaling_resource_type: Optional[pulumi.Input[str]] = None,
                 scheduled_policy: Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BandwidthPolicy resources.
        :param pulumi.Input[str] alarm_id: Specifies the alarm rule ID.
               This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        :param pulumi.Input[str] bandwidth_id: Specifies the scaling bandwidth ID.
        :param pulumi.Input[int] cool_down_time: Specifies the cooldown period (in seconds).
               The value ranges from 0 to 86400 and is 300 by default.
        :param pulumi.Input[str] description: Specifies the description of the AS policy.
               The value can contain 0 to 256 characters.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input['BandwidthPolicyScalingPolicyActionArgs'] scaling_policy_action: Specifies the scaling action of the AS policy.
               The object structure is documented below.
        :param pulumi.Input[str] scaling_policy_name: Specifies the AS policy name.
               The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        :param pulumi.Input[str] scaling_policy_type: Specifies the AS policy type. The options are as follows:
               - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
               - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
               - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        :param pulumi.Input[str] scaling_resource_type: The scaling resource type. The value is fixed to **BANDWIDTH**.
        :param pulumi.Input['BandwidthPolicyScheduledPolicyArgs'] scheduled_policy: Specifies the periodic or scheduled AS policy.
               This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
               The object structure is documented below.
        :param pulumi.Input[str] status: The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
        """
        if alarm_id is not None:
            pulumi.set(__self__, "alarm_id", alarm_id)
        if bandwidth_id is not None:
            pulumi.set(__self__, "bandwidth_id", bandwidth_id)
        if cool_down_time is not None:
            pulumi.set(__self__, "cool_down_time", cool_down_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scaling_policy_action is not None:
            pulumi.set(__self__, "scaling_policy_action", scaling_policy_action)
        if scaling_policy_name is not None:
            pulumi.set(__self__, "scaling_policy_name", scaling_policy_name)
        if scaling_policy_type is not None:
            pulumi.set(__self__, "scaling_policy_type", scaling_policy_type)
        if scaling_resource_type is not None:
            pulumi.set(__self__, "scaling_resource_type", scaling_resource_type)
        if scheduled_policy is not None:
            pulumi.set(__self__, "scheduled_policy", scheduled_policy)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the alarm rule ID.
        This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        """
        return pulumi.get(self, "alarm_id")

    @alarm_id.setter
    def alarm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_id", value)

    @property
    @pulumi.getter(name="bandwidthId")
    def bandwidth_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the scaling bandwidth ID.
        """
        return pulumi.get(self, "bandwidth_id")

    @bandwidth_id.setter
    def bandwidth_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_id", value)

    @property
    @pulumi.getter(name="coolDownTime")
    def cool_down_time(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the cooldown period (in seconds).
        The value ranges from 0 to 86400 and is 300 by default.
        """
        return pulumi.get(self, "cool_down_time")

    @cool_down_time.setter
    def cool_down_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cool_down_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the AS policy.
        The value can contain 0 to 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="scalingPolicyAction")
    def scaling_policy_action(self) -> Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']]:
        """
        Specifies the scaling action of the AS policy.
        The object structure is documented below.
        """
        return pulumi.get(self, "scaling_policy_action")

    @scaling_policy_action.setter
    def scaling_policy_action(self, value: Optional[pulumi.Input['BandwidthPolicyScalingPolicyActionArgs']]):
        pulumi.set(self, "scaling_policy_action", value)

    @property
    @pulumi.getter(name="scalingPolicyName")
    def scaling_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AS policy name.
        The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        """
        return pulumi.get(self, "scaling_policy_name")

    @scaling_policy_name.setter
    def scaling_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_policy_name", value)

    @property
    @pulumi.getter(name="scalingPolicyType")
    def scaling_policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the AS policy type. The options are as follows:
        - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
        - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
        - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        """
        return pulumi.get(self, "scaling_policy_type")

    @scaling_policy_type.setter
    def scaling_policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_policy_type", value)

    @property
    @pulumi.getter(name="scalingResourceType")
    def scaling_resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The scaling resource type. The value is fixed to **BANDWIDTH**.
        """
        return pulumi.get(self, "scaling_resource_type")

    @scaling_resource_type.setter
    def scaling_resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_resource_type", value)

    @property
    @pulumi.getter(name="scheduledPolicy")
    def scheduled_policy(self) -> Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']]:
        """
        Specifies the periodic or scheduled AS policy.
        This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
        The object structure is documented below.
        """
        return pulumi.get(self, "scheduled_policy")

    @scheduled_policy.setter
    def scheduled_policy(self, value: Optional[pulumi.Input['BandwidthPolicyScheduledPolicyArgs']]):
        pulumi.set(self, "scheduled_policy", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BandwidthPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_id: Optional[pulumi.Input[str]] = None,
                 bandwidth_id: Optional[pulumi.Input[str]] = None,
                 cool_down_time: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scaling_policy_action: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScalingPolicyActionArgs']]] = None,
                 scaling_policy_name: Optional[pulumi.Input[str]] = None,
                 scaling_policy_type: Optional[pulumi.Input[str]] = None,
                 scheduled_policy: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScheduledPolicyArgs']]] = None,
                 __props__=None):
        """
        Manages an AS bandwidth scaling policy resource within HuaweiCloud.

        > AS cannot scale yearly/monthly bandwidths.

        ## Example Usage
        ### AS Recurrence Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        bw_policy = huaweicloud.as_.BandwidthPolicy("bwPolicy",
            scaling_policy_name="bw_policy",
            scaling_policy_type="RECURRENCE",
            bandwidth_id=bandwidth_id,
            cool_down_time=600,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
            ),
            scheduled_policy=huaweicloud.as_.BandwidthPolicyScheduledPolicyArgs(
                launch_time="07:00",
                recurrence_type="Weekly",
                recurrence_value="1,3,5",
                start_time="2022-09-30T12:00Z",
                end_time="2022-12-30T12:00Z",
            ))
        ```
        ### AS Scheduled Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        bw_policy = huaweicloud.as_.BandwidthPolicy("bwPolicy",
            scaling_policy_name="bw_policy",
            scaling_policy_type="SCHEDULED",
            bandwidth_id=bandwidth_id,
            cool_down_time=600,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
            ),
            scheduled_policy=huaweicloud.as_.BandwidthPolicyScheduledPolicyArgs(
                launch_time="2022-09-30T12:00Z",
            ))
        ```
        ### AS Alarm Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        alarm_id = config.require_object("alarmId")
        test = huaweicloud.as_.BandwidthPolicy("test",
            scaling_policy_name="bw_policy",
            scaling_policy_type="ALARM",
            bandwidth_id=bandwidth_id,
            alarm_id=alarm_id,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
                limits=300,
            ))
        ```

        ## Import

        The bandwidth scaling policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:As/bandwidthPolicy:BandwidthPolicy test 0ce123456a00f2591fabc00385ff1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_id: Specifies the alarm rule ID.
               This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        :param pulumi.Input[str] bandwidth_id: Specifies the scaling bandwidth ID.
        :param pulumi.Input[int] cool_down_time: Specifies the cooldown period (in seconds).
               The value ranges from 0 to 86400 and is 300 by default.
        :param pulumi.Input[str] description: Specifies the description of the AS policy.
               The value can contain 0 to 256 characters.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[pulumi.InputType['BandwidthPolicyScalingPolicyActionArgs']] scaling_policy_action: Specifies the scaling action of the AS policy.
               The object structure is documented below.
        :param pulumi.Input[str] scaling_policy_name: Specifies the AS policy name.
               The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        :param pulumi.Input[str] scaling_policy_type: Specifies the AS policy type. The options are as follows:
               - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
               - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
               - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        :param pulumi.Input[pulumi.InputType['BandwidthPolicyScheduledPolicyArgs']] scheduled_policy: Specifies the periodic or scheduled AS policy.
               This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
               The object structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BandwidthPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an AS bandwidth scaling policy resource within HuaweiCloud.

        > AS cannot scale yearly/monthly bandwidths.

        ## Example Usage
        ### AS Recurrence Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        bw_policy = huaweicloud.as_.BandwidthPolicy("bwPolicy",
            scaling_policy_name="bw_policy",
            scaling_policy_type="RECURRENCE",
            bandwidth_id=bandwidth_id,
            cool_down_time=600,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
            ),
            scheduled_policy=huaweicloud.as_.BandwidthPolicyScheduledPolicyArgs(
                launch_time="07:00",
                recurrence_type="Weekly",
                recurrence_value="1,3,5",
                start_time="2022-09-30T12:00Z",
                end_time="2022-12-30T12:00Z",
            ))
        ```
        ### AS Scheduled Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        bw_policy = huaweicloud.as_.BandwidthPolicy("bwPolicy",
            scaling_policy_name="bw_policy",
            scaling_policy_type="SCHEDULED",
            bandwidth_id=bandwidth_id,
            cool_down_time=600,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
            ),
            scheduled_policy=huaweicloud.as_.BandwidthPolicyScheduledPolicyArgs(
                launch_time="2022-09-30T12:00Z",
            ))
        ```
        ### AS Alarm Policy

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        bandwidth_id = config.require_object("bandwidthId")
        alarm_id = config.require_object("alarmId")
        test = huaweicloud.as_.BandwidthPolicy("test",
            scaling_policy_name="bw_policy",
            scaling_policy_type="ALARM",
            bandwidth_id=bandwidth_id,
            alarm_id=alarm_id,
            scaling_policy_action=huaweicloud.as_.BandwidthPolicyScalingPolicyActionArgs(
                operation="ADD",
                size=1,
                limits=300,
            ))
        ```

        ## Import

        The bandwidth scaling policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:As/bandwidthPolicy:BandwidthPolicy test 0ce123456a00f2591fabc00385ff1234
        ```

        :param str resource_name: The name of the resource.
        :param BandwidthPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_id: Optional[pulumi.Input[str]] = None,
                 bandwidth_id: Optional[pulumi.Input[str]] = None,
                 cool_down_time: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scaling_policy_action: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScalingPolicyActionArgs']]] = None,
                 scaling_policy_name: Optional[pulumi.Input[str]] = None,
                 scaling_policy_type: Optional[pulumi.Input[str]] = None,
                 scheduled_policy: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScheduledPolicyArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BandwidthPolicyArgs.__new__(BandwidthPolicyArgs)

            __props__.__dict__["alarm_id"] = alarm_id
            if bandwidth_id is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth_id'")
            __props__.__dict__["bandwidth_id"] = bandwidth_id
            __props__.__dict__["cool_down_time"] = cool_down_time
            __props__.__dict__["description"] = description
            __props__.__dict__["region"] = region
            __props__.__dict__["scaling_policy_action"] = scaling_policy_action
            if scaling_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_policy_name'")
            __props__.__dict__["scaling_policy_name"] = scaling_policy_name
            if scaling_policy_type is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_policy_type'")
            __props__.__dict__["scaling_policy_type"] = scaling_policy_type
            __props__.__dict__["scheduled_policy"] = scheduled_policy
            __props__.__dict__["scaling_resource_type"] = None
            __props__.__dict__["status"] = None
        super(BandwidthPolicy, __self__).__init__(
            'huaweicloud:As/bandwidthPolicy:BandwidthPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alarm_id: Optional[pulumi.Input[str]] = None,
            bandwidth_id: Optional[pulumi.Input[str]] = None,
            cool_down_time: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scaling_policy_action: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScalingPolicyActionArgs']]] = None,
            scaling_policy_name: Optional[pulumi.Input[str]] = None,
            scaling_policy_type: Optional[pulumi.Input[str]] = None,
            scaling_resource_type: Optional[pulumi.Input[str]] = None,
            scheduled_policy: Optional[pulumi.Input[pulumi.InputType['BandwidthPolicyScheduledPolicyArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BandwidthPolicy':
        """
        Get an existing BandwidthPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_id: Specifies the alarm rule ID.
               This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        :param pulumi.Input[str] bandwidth_id: Specifies the scaling bandwidth ID.
        :param pulumi.Input[int] cool_down_time: Specifies the cooldown period (in seconds).
               The value ranges from 0 to 86400 and is 300 by default.
        :param pulumi.Input[str] description: Specifies the description of the AS policy.
               The value can contain 0 to 256 characters.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[pulumi.InputType['BandwidthPolicyScalingPolicyActionArgs']] scaling_policy_action: Specifies the scaling action of the AS policy.
               The object structure is documented below.
        :param pulumi.Input[str] scaling_policy_name: Specifies the AS policy name.
               The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        :param pulumi.Input[str] scaling_policy_type: Specifies the AS policy type. The options are as follows:
               - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
               - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
               - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        :param pulumi.Input[str] scaling_resource_type: The scaling resource type. The value is fixed to **BANDWIDTH**.
        :param pulumi.Input[pulumi.InputType['BandwidthPolicyScheduledPolicyArgs']] scheduled_policy: Specifies the periodic or scheduled AS policy.
               This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
               The object structure is documented below.
        :param pulumi.Input[str] status: The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPolicyState.__new__(_BandwidthPolicyState)

        __props__.__dict__["alarm_id"] = alarm_id
        __props__.__dict__["bandwidth_id"] = bandwidth_id
        __props__.__dict__["cool_down_time"] = cool_down_time
        __props__.__dict__["description"] = description
        __props__.__dict__["region"] = region
        __props__.__dict__["scaling_policy_action"] = scaling_policy_action
        __props__.__dict__["scaling_policy_name"] = scaling_policy_name
        __props__.__dict__["scaling_policy_type"] = scaling_policy_type
        __props__.__dict__["scaling_resource_type"] = scaling_resource_type
        __props__.__dict__["scheduled_policy"] = scheduled_policy
        __props__.__dict__["status"] = status
        return BandwidthPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> pulumi.Output[str]:
        """
        Specifies the alarm rule ID.
        This parameter is mandatory when `scaling_policy_type` is set to ALARM.
        """
        return pulumi.get(self, "alarm_id")

    @property
    @pulumi.getter(name="bandwidthId")
    def bandwidth_id(self) -> pulumi.Output[str]:
        """
        Specifies the scaling bandwidth ID.
        """
        return pulumi.get(self, "bandwidth_id")

    @property
    @pulumi.getter(name="coolDownTime")
    def cool_down_time(self) -> pulumi.Output[int]:
        """
        Specifies the cooldown period (in seconds).
        The value ranges from 0 to 86400 and is 300 by default.
        """
        return pulumi.get(self, "cool_down_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Specifies the description of the AS policy.
        The value can contain 0 to 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="scalingPolicyAction")
    def scaling_policy_action(self) -> pulumi.Output['outputs.BandwidthPolicyScalingPolicyAction']:
        """
        Specifies the scaling action of the AS policy.
        The object structure is documented below.
        """
        return pulumi.get(self, "scaling_policy_action")

    @property
    @pulumi.getter(name="scalingPolicyName")
    def scaling_policy_name(self) -> pulumi.Output[str]:
        """
        Specifies the AS policy name.
        The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
        """
        return pulumi.get(self, "scaling_policy_name")

    @property
    @pulumi.getter(name="scalingPolicyType")
    def scaling_policy_type(self) -> pulumi.Output[str]:
        """
        Specifies the AS policy type. The options are as follows:
        - **ALARM** (corresponding to `alarm_id`): indicates that the scaling action is triggered by an alarm.
        - **SCHEDULED** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered as scheduled.
        - **RECURRENCE** (corresponding to `scheduled_policy`): indicates that the scaling action is triggered periodically.
        """
        return pulumi.get(self, "scaling_policy_type")

    @property
    @pulumi.getter(name="scalingResourceType")
    def scaling_resource_type(self) -> pulumi.Output[str]:
        """
        The scaling resource type. The value is fixed to **BANDWIDTH**.
        """
        return pulumi.get(self, "scaling_resource_type")

    @property
    @pulumi.getter(name="scheduledPolicy")
    def scheduled_policy(self) -> pulumi.Output['outputs.BandwidthPolicyScheduledPolicy']:
        """
        Specifies the periodic or scheduled AS policy.
        This parameter is mandatory when `scaling_policy_type` is set to SCHEDULED or RECURRENCE.
        The object structure is documented below.
        """
        return pulumi.get(self, "scheduled_policy")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The AS policy status. The value can be **INSERVICE**, **PAUSED** and **EXECUTING**.
        """
        return pulumi.get(self, "status")
