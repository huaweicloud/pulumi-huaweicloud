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
    'AlarmruleAlarmActionArgs',
    'AlarmruleConditionArgs',
    'AlarmruleInsufficientdataActionArgs',
    'AlarmruleMetricArgs',
    'AlarmruleMetricDimensionArgs',
    'AlarmruleOkActionArgs',
    'AlarmruleResourceArgs',
    'AlarmruleResourceDimensionArgs',
    'MicroserviceEngineConfigCenterAddressArgs',
    'MicroserviceEngineServiceRegistryAddressArgs',
    'MicroserviceInstanceDataCenterArgs',
    'MicroserviceInstanceHealthCheckArgs',
]

@pulumi.input_type
class AlarmruleAlarmActionArgs:
    def __init__(__self__, *,
                 notification_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_lists: specifies the list of objects to be notified if the alarm status changes, the
               maximum length is 5.
        :param pulumi.Input[str] type: Specifies the type of action triggered by an alarm. the value is notification.
               notification: indicates that a notification will be sent to the user.
        """
        pulumi.set(__self__, "notification_lists", notification_lists)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="notificationLists")
    def notification_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        specifies the list of objects to be notified if the alarm status changes, the
        maximum length is 5.
        """
        return pulumi.get(self, "notification_lists")

    @notification_lists.setter
    def notification_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_lists", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of action triggered by an alarm. the value is notification.
        notification: indicates that a notification will be sent to the user.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class AlarmruleConditionArgs:
    def __init__(__self__, *,
                 comparison_operator: pulumi.Input[str],
                 count: pulumi.Input[int],
                 filter: pulumi.Input[str],
                 period: pulumi.Input[int],
                 value: pulumi.Input[float],
                 alarm_level: Optional[pulumi.Input[int]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 suppress_duration: Optional[pulumi.Input[int]] = None,
                 unit: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] comparison_operator: Specifies the comparison condition of alarm thresholds. The value can be >,
               =, <, >=, or <=.
        :param pulumi.Input[int] count: Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.
        :param pulumi.Input[str] filter: Specifies the data rollup methods. The value can be max, min, average, sum, and variance.
        :param pulumi.Input[int] period: Specifies the alarm checking period in seconds. The value can be 0, 1, 300, 1200, 3600, 14400,
               and 86400.
        :param pulumi.Input[float] value: Specifies the alarm threshold. The value ranges from 0 to Number of
               1.7976931348623157e+108.
        :param pulumi.Input[int] alarm_level: Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
               which indicates *critical*, *major*, *minor*, and *informational*, respectively.
               The default value is 2.
        :param pulumi.Input[str] metric_name: Specifies the metric name of the condition. The value can be a string of
               1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
               For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        :param pulumi.Input[int] suppress_duration: Specifies the interval for triggering an alarm if the alarm persists.
               Possible values are as follows:
               + **0**: Cloud Eye triggers the alarm only once;
               + **300**: Cloud Eye triggers the alarm every 5 minutes;
               + **600**: Cloud Eye triggers the alarm every 10 minutes;
               + **900**: Cloud Eye triggers the alarm every 15 minutes;
               + **1800**: Cloud Eye triggers the alarm every 30 minutes;
               + **3600**: Cloud Eye triggers the alarm every hour;
               + **10800**: Cloud Eye triggers the alarm every 3 hours;
               + **21600**: Cloud Eye triggers the alarm every 6 hours;
               + **43200**: Cloud Eye triggers the alarm every 12 hour;
               + **86400**: Cloud Eye triggers the alarm every day.
        :param pulumi.Input[str] unit: Specifies the data unit.
               For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "period", period)
        pulumi.set(__self__, "value", value)
        if alarm_level is not None:
            pulumi.set(__self__, "alarm_level", alarm_level)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)
        if suppress_duration is not None:
            pulumi.set(__self__, "suppress_duration", suppress_duration)
        if unit is not None:
            pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> pulumi.Input[str]:
        """
        Specifies the comparison condition of alarm thresholds. The value can be >,
        =, <, >=, or <=.
        """
        return pulumi.get(self, "comparison_operator")

    @comparison_operator.setter
    def comparison_operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison_operator", value)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[int]:
        """
        Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[int]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        Specifies the data rollup methods. The value can be max, min, average, sum, and variance.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[int]:
        """
        Specifies the alarm checking period in seconds. The value can be 0, 1, 300, 1200, 3600, 14400,
        and 86400.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[int]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[float]:
        """
        Specifies the alarm threshold. The value ranges from 0 to Number of
        1.7976931348623157e+108.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[float]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="alarmLevel")
    def alarm_level(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the alarm severity of the condition. The value can be 1, 2, 3 or 4,
        which indicates *critical*, *major*, *minor*, and *informational*, respectively.
        The default value is 2.
        """
        return pulumi.get(self, "alarm_level")

    @alarm_level.setter
    def alarm_level(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alarm_level", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the metric name of the condition. The value can be a string of
        1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
        For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="suppressDuration")
    def suppress_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the interval for triggering an alarm if the alarm persists.
        Possible values are as follows:
        + **0**: Cloud Eye triggers the alarm only once;
        + **300**: Cloud Eye triggers the alarm every 5 minutes;
        + **600**: Cloud Eye triggers the alarm every 10 minutes;
        + **900**: Cloud Eye triggers the alarm every 15 minutes;
        + **1800**: Cloud Eye triggers the alarm every 30 minutes;
        + **3600**: Cloud Eye triggers the alarm every hour;
        + **10800**: Cloud Eye triggers the alarm every 3 hours;
        + **21600**: Cloud Eye triggers the alarm every 6 hours;
        + **43200**: Cloud Eye triggers the alarm every 12 hour;
        + **86400**: Cloud Eye triggers the alarm every day.
        """
        return pulumi.get(self, "suppress_duration")

    @suppress_duration.setter
    def suppress_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "suppress_duration", value)

    @property
    @pulumi.getter
    def unit(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the data unit.
        For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit", value)


@pulumi.input_type
class AlarmruleInsufficientdataActionArgs:
    def __init__(__self__, *,
                 notification_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_lists: specifies the list of objects to be notified if the alarm status changes, the
               maximum length is 5.
        :param pulumi.Input[str] type: Specifies the type of action triggered by an alarm. the value is notification.
               notification: indicates that a notification will be sent to the user.
        """
        pulumi.set(__self__, "notification_lists", notification_lists)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="notificationLists")
    def notification_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        specifies the list of objects to be notified if the alarm status changes, the
        maximum length is 5.
        """
        return pulumi.get(self, "notification_lists")

    @notification_lists.setter
    def notification_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_lists", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of action triggered by an alarm. the value is notification.
        notification: indicates that a notification will be sent to the user.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class AlarmruleMetricArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleMetricDimensionArgs']]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] namespace: Specifies the namespace in **service.item** format. **service** and **item**
               each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
               Changing this creates a new resource.
               For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        :param pulumi.Input[Sequence[pulumi.Input['AlarmruleMetricDimensionArgs']]] dimensions: Specifies the list of metric dimensions. The structure is described below.
        :param pulumi.Input[str] metric_name: Specifies the metric name of the condition. The value can be a string of
               1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
               For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        pulumi.set(__self__, "namespace", namespace)
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        Specifies the namespace in **service.item** format. **service** and **item**
        each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
        Changing this creates a new resource.
        For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleMetricDimensionArgs']]]]:
        """
        Specifies the list of metric dimensions. The structure is described below.
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleMetricDimensionArgs']]]]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the metric name of the condition. The value can be a string of
        1 to 64 characters that must start with a letter and contain only letters, digits, and underscores (_).
        For details, see [Services Interconnected with Cloud Eye](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html).
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)


@pulumi.input_type
class AlarmruleMetricDimensionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the dimension name. The value can be a string of 1 to 32 characters
               that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
        :param pulumi.Input[str] value: Specifies the alarm threshold. The value ranges from 0 to Number of
               1.7976931348623157e+108.
        """
        pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the dimension name. The value can be a string of 1 to 32 characters
        that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the alarm threshold. The value ranges from 0 to Number of
        1.7976931348623157e+108.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class AlarmruleOkActionArgs:
    def __init__(__self__, *,
                 notification_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_lists: specifies the list of objects to be notified if the alarm status changes, the
               maximum length is 5.
        :param pulumi.Input[str] type: Specifies the type of action triggered by an alarm. the value is notification.
               notification: indicates that a notification will be sent to the user.
        """
        pulumi.set(__self__, "notification_lists", notification_lists)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="notificationLists")
    def notification_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        specifies the list of objects to be notified if the alarm status changes, the
        maximum length is 5.
        """
        return pulumi.get(self, "notification_lists")

    @notification_lists.setter
    def notification_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_lists", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of action triggered by an alarm. the value is notification.
        notification: indicates that a notification will be sent to the user.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class AlarmruleResourceArgs:
    def __init__(__self__, *,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleResourceDimensionArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['AlarmruleResourceDimensionArgs']]] dimensions: Specifies the list of metric dimensions. The structure is described below.
        """
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleResourceDimensionArgs']]]]:
        """
        Specifies the list of metric dimensions. The structure is described below.
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmruleResourceDimensionArgs']]]]):
        pulumi.set(self, "dimensions", value)


@pulumi.input_type
class AlarmruleResourceDimensionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the dimension name. The value can be a string of 1 to 32 characters
               that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
        :param pulumi.Input[str] value: Specifies the alarm threshold. The value ranges from 0 to Number of
               1.7976931348623157e+108.
        """
        pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the dimension name. The value can be a string of 1 to 32 characters
        that must start with a letter and contain only letters, digits, underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the alarm threshold. The value ranges from 0 to Number of
        1.7976931348623157e+108.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class MicroserviceEngineConfigCenterAddressArgs:
    def __init__(__self__, *,
                 private: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] private: The internal access address.
        :param pulumi.Input[str] public: The public access address. This address is only set when EIP is bound.
        """
        if private is not None:
            pulumi.set(__self__, "private", private)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def private(self) -> Optional[pulumi.Input[str]]:
        """
        The internal access address.
        """
        return pulumi.get(self, "private")

    @private.setter
    def private(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[str]]:
        """
        The public access address. This address is only set when EIP is bound.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public", value)


@pulumi.input_type
class MicroserviceEngineServiceRegistryAddressArgs:
    def __init__(__self__, *,
                 private: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] private: The internal access address.
        :param pulumi.Input[str] public: The public access address. This address is only set when EIP is bound.
        """
        if private is not None:
            pulumi.set(__self__, "private", private)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def private(self) -> Optional[pulumi.Input[str]]:
        """
        The internal access address.
        """
        return pulumi.get(self, "private")

    @private.setter
    def private(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[str]]:
        """
        The public access address. This address is only set when EIP is bound.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public", value)


@pulumi.input_type
class MicroserviceInstanceDataCenterArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 name: pulumi.Input[str],
                 region: pulumi.Input[str]):
        """
        :param pulumi.Input[str] availability_zone: Specifies the custom availability zone name of the data center.
               Changing this will create a new microservice instance.
        :param pulumi.Input[str] name: Specifies the data center name.
               Changing this will create a new microservice instance.
        :param pulumi.Input[str] region: Specifies the custom region name of the data center.
               Changing this will create a new microservice instance.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        Specifies the custom availability zone name of the data center.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the data center name.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Specifies the custom region name of the data center.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class MicroserviceInstanceHealthCheckArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[int],
                 max_retries: pulumi.Input[int],
                 mode: pulumi.Input[str],
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] interval: Specifies the heartbeat interval. The unit is **s** (second).
               Changing this will create a new microservice instance.
        :param pulumi.Input[int] max_retries: Specifies the maximum retries.
               Changing this will create a new microservice instance.
        :param pulumi.Input[str] mode: Specifies the heartbeat mode. The valid values are **push** and **pull**.
               Changing this will create a new microservice instance.
        :param pulumi.Input[int] port: Specifies the port number.
               Changing this will create a new microservice instance.
        """
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "max_retries", max_retries)
        pulumi.set(__self__, "mode", mode)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[int]:
        """
        Specifies the heartbeat interval. The unit is **s** (second).
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Input[int]:
        """
        Specifies the maximum retries.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        Specifies the heartbeat mode. The valid values are **push** and **pull**.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the port number.
        Changing this will create a new microservice instance.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


