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

__all__ = [
    'ConferenceConfiguration',
    'ConferenceCycleParams',
    'ConferenceJoinPassword',
    'ConferenceParticipant',
    'ConferenceSubconference',
    'ConferenceSubconferenceSubconfiguration',
    'ConferenceSubconferenceSubconfigurationShowAudiencePolicy',
]

@pulumi.output_type
class ConferenceConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowGuestStart":
            suggest = "allow_guest_start"
        elif key == "callinRestriction":
            suggest = "callin_restriction"
        elif key == "guestPassword":
            suggest = "guest_password"
        elif key == "isAutoMute":
            suggest = "is_auto_mute"
        elif key == "isGuestFreePassword":
            suggest = "is_guest_free_password"
        elif key == "isHardTerminalAutoMute":
            suggest = "is_hard_terminal_auto_mute"
        elif key == "isSendCalendar":
            suggest = "is_send_calendar"
        elif key == "isSendNotify":
            suggest = "is_send_notify"
        elif key == "isSendSms":
            suggest = "is_send_sms"
        elif key == "prolongTime":
            suggest = "prolong_time"
        elif key == "waitingRoomEnabled":
            suggest = "waiting_room_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_guest_start: Optional[bool] = None,
                 callin_restriction: Optional[int] = None,
                 guest_password: Optional[str] = None,
                 is_auto_mute: Optional[bool] = None,
                 is_guest_free_password: Optional[bool] = None,
                 is_hard_terminal_auto_mute: Optional[bool] = None,
                 is_send_calendar: Optional[bool] = None,
                 is_send_notify: Optional[bool] = None,
                 is_send_sms: Optional[bool] = None,
                 prolong_time: Optional[int] = None,
                 waiting_room_enabled: Optional[bool] = None):
        """
        :param bool allow_guest_start: Specifies whether to allow guests to start conferences (only valid for random
               ID conferences).
        :param int callin_restriction: Specifies the range to allow incoming calls.
               + **0**: All users.
               + **2**: Users within the enterprise.
               + **3**: The invited user.
        :param str guest_password: Specifies the guest password (pure number which is 4 to 16 digits long).
        :param bool is_auto_mute: Specifies whether the soft terminal is automatically muted when the guest joins the
               conference.
        :param bool is_guest_free_password: Specifies whether the guest is password-free (only valid for random
               conferences).
        :param bool is_hard_terminal_auto_mute: Specifies whether the guest joins the conference, whether the hard
               terminal is automatically muted.
        :param bool is_send_calendar: Specifies whether to send conference calendar notifications.
        :param bool is_send_notify: Specifies whether to send conference email notification.
        :param bool is_send_sms: Specifies whether to send conference SMS notification.
        :param int prolong_time: Specifies the Automatically extend duration, the valid value is range from `0` to
               `60`.
        :param bool waiting_room_enabled: Specifies whether to open the waiting room (only valid for RTC enterprises).
        """
        if allow_guest_start is not None:
            pulumi.set(__self__, "allow_guest_start", allow_guest_start)
        if callin_restriction is not None:
            pulumi.set(__self__, "callin_restriction", callin_restriction)
        if guest_password is not None:
            pulumi.set(__self__, "guest_password", guest_password)
        if is_auto_mute is not None:
            pulumi.set(__self__, "is_auto_mute", is_auto_mute)
        if is_guest_free_password is not None:
            pulumi.set(__self__, "is_guest_free_password", is_guest_free_password)
        if is_hard_terminal_auto_mute is not None:
            pulumi.set(__self__, "is_hard_terminal_auto_mute", is_hard_terminal_auto_mute)
        if is_send_calendar is not None:
            pulumi.set(__self__, "is_send_calendar", is_send_calendar)
        if is_send_notify is not None:
            pulumi.set(__self__, "is_send_notify", is_send_notify)
        if is_send_sms is not None:
            pulumi.set(__self__, "is_send_sms", is_send_sms)
        if prolong_time is not None:
            pulumi.set(__self__, "prolong_time", prolong_time)
        if waiting_room_enabled is not None:
            pulumi.set(__self__, "waiting_room_enabled", waiting_room_enabled)

    @property
    @pulumi.getter(name="allowGuestStart")
    def allow_guest_start(self) -> Optional[bool]:
        """
        Specifies whether to allow guests to start conferences (only valid for random
        ID conferences).
        """
        return pulumi.get(self, "allow_guest_start")

    @property
    @pulumi.getter(name="callinRestriction")
    def callin_restriction(self) -> Optional[int]:
        """
        Specifies the range to allow incoming calls.
        + **0**: All users.
        + **2**: Users within the enterprise.
        + **3**: The invited user.
        """
        return pulumi.get(self, "callin_restriction")

    @property
    @pulumi.getter(name="guestPassword")
    def guest_password(self) -> Optional[str]:
        """
        Specifies the guest password (pure number which is 4 to 16 digits long).
        """
        return pulumi.get(self, "guest_password")

    @property
    @pulumi.getter(name="isAutoMute")
    def is_auto_mute(self) -> Optional[bool]:
        """
        Specifies whether the soft terminal is automatically muted when the guest joins the
        conference.
        """
        return pulumi.get(self, "is_auto_mute")

    @property
    @pulumi.getter(name="isGuestFreePassword")
    def is_guest_free_password(self) -> Optional[bool]:
        """
        Specifies whether the guest is password-free (only valid for random
        conferences).
        """
        return pulumi.get(self, "is_guest_free_password")

    @property
    @pulumi.getter(name="isHardTerminalAutoMute")
    def is_hard_terminal_auto_mute(self) -> Optional[bool]:
        """
        Specifies whether the guest joins the conference, whether the hard
        terminal is automatically muted.
        """
        return pulumi.get(self, "is_hard_terminal_auto_mute")

    @property
    @pulumi.getter(name="isSendCalendar")
    def is_send_calendar(self) -> Optional[bool]:
        """
        Specifies whether to send conference calendar notifications.
        """
        return pulumi.get(self, "is_send_calendar")

    @property
    @pulumi.getter(name="isSendNotify")
    def is_send_notify(self) -> Optional[bool]:
        """
        Specifies whether to send conference email notification.
        """
        return pulumi.get(self, "is_send_notify")

    @property
    @pulumi.getter(name="isSendSms")
    def is_send_sms(self) -> Optional[bool]:
        """
        Specifies whether to send conference SMS notification.
        """
        return pulumi.get(self, "is_send_sms")

    @property
    @pulumi.getter(name="prolongTime")
    def prolong_time(self) -> Optional[int]:
        """
        Specifies the Automatically extend duration, the valid value is range from `0` to
        `60`.
        """
        return pulumi.get(self, "prolong_time")

    @property
    @pulumi.getter(name="waitingRoomEnabled")
    def waiting_room_enabled(self) -> Optional[bool]:
        """
        Specifies whether to open the waiting room (only valid for RTC enterprises).
        """
        return pulumi.get(self, "waiting_room_enabled")


@pulumi.output_type
class ConferenceCycleParams(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endDate":
            suggest = "end_date"
        elif key == "preRemind":
            suggest = "pre_remind"
        elif key == "startDate":
            suggest = "start_date"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceCycleParams. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceCycleParams.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceCycleParams.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cycle: str,
                 end_date: str,
                 pre_remind: int,
                 start_date: str,
                 interval: Optional[int] = None,
                 points: Optional[Sequence[int]] = None):
        """
        :param str cycle: Specifies the period type. The valid values are as follows:
               + **Day**
               + **Week**
               + **Month**
        :param str end_date: Specifies the end date of the recurring conference.
               The format is `YYYY-MM-DD`.
        :param int pre_remind: Specifies the number of days for advance conference notice.
               The valid value is range from `0` to `30`, defaults to `1`.
        :param str start_date: Specifies the start date of the recurring conference.
               The format is `YYYY-MM-DD`.
        :param int interval: Specifies the cycle interval.
               For different `cycle` types, the value range of interval are as follows:
               + **Day**: Means that it will be held every few days, and the valid value is range from `1` to `15`.
               + **Week**: Means that it is held every few weeks, and the valid value is range from `1` to `5`.
               + **Month**: Means every few months, the value range is `1` to `3`.
        :param Sequence[int] points: Specifies the conference point in the cycle. Only valid by **Week** and **Month**.
               For different `cycle` types, the value range of elements are as follows:
               + **Week**: The valid value is range from `0` to `6`. The **0** means Sunday, **6** means Saturday.
               + **Month**: The valid range for the elements is `1` to `31`. If the value does not exist in the current month, the
               value means the end of the month.
        """
        pulumi.set(__self__, "cycle", cycle)
        pulumi.set(__self__, "end_date", end_date)
        pulumi.set(__self__, "pre_remind", pre_remind)
        pulumi.set(__self__, "start_date", start_date)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if points is not None:
            pulumi.set(__self__, "points", points)

    @property
    @pulumi.getter
    def cycle(self) -> str:
        """
        Specifies the period type. The valid values are as follows:
        + **Day**
        + **Week**
        + **Month**
        """
        return pulumi.get(self, "cycle")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        """
        Specifies the end date of the recurring conference.
        The format is `YYYY-MM-DD`.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="preRemind")
    def pre_remind(self) -> int:
        """
        Specifies the number of days for advance conference notice.
        The valid value is range from `0` to `30`, defaults to `1`.
        """
        return pulumi.get(self, "pre_remind")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        """
        Specifies the start date of the recurring conference.
        The format is `YYYY-MM-DD`.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        """
        Specifies the cycle interval.
        For different `cycle` types, the value range of interval are as follows:
        + **Day**: Means that it will be held every few days, and the valid value is range from `1` to `15`.
        + **Week**: Means that it is held every few weeks, and the valid value is range from `1` to `5`.
        + **Month**: Means every few months, the value range is `1` to `3`.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def points(self) -> Optional[Sequence[int]]:
        """
        Specifies the conference point in the cycle. Only valid by **Week** and **Month**.
        For different `cycle` types, the value range of elements are as follows:
        + **Week**: The valid value is range from `0` to `6`. The **0** means Sunday, **6** means Saturday.
        + **Month**: The valid range for the elements is `1` to `31`. If the value does not exist in the current month, the
        value means the end of the month.
        """
        return pulumi.get(self, "points")


@pulumi.output_type
class ConferenceJoinPassword(dict):
    def __init__(__self__, *,
                 guest: Optional[str] = None,
                 host: Optional[str] = None):
        """
        :param str guest: The password of the common participant.
        :param str host: The password of the meeting host.
        """
        if guest is not None:
            pulumi.set(__self__, "guest", guest)
        if host is not None:
            pulumi.set(__self__, "host", host)

    @property
    @pulumi.getter
    def guest(self) -> Optional[str]:
        """
        The password of the common participant.
        """
        return pulumi.get(self, "guest")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        The password of the meeting host.
        """
        return pulumi.get(self, "host")


@pulumi.output_type
class ConferenceParticipant(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "isAutoInvite":
            suggest = "is_auto_invite"
        elif key == "isMute":
            suggest = "is_mute"
        elif key == "userId":
            suggest = "user_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceParticipant. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceParticipant.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceParticipant.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_id: Optional[str] = None,
                 email: Optional[str] = None,
                 is_auto_invite: Optional[int] = None,
                 is_mute: Optional[int] = None,
                 name: Optional[str] = None,
                 phone: Optional[str] = None,
                 role: Optional[int] = None,
                 sms: Optional[str] = None,
                 type: Optional[str] = None,
                 user_id: Optional[str] = None):
        """
        :param str account_id: Specifies the account ID of the participant.
        :param str email: Specifies the email address.
        :param int is_auto_invite: Specifies whether to automatically invite this participant when the conference
               starts. The valid values are as follows:
               + **0**: Do not automatically invite.
               + **1**: Automatic invitation.
        :param int is_mute: Specifies whether the user needs to be automatically muted when joining the conference
               (only effective when invited in the conference). The valid values are as follows:
               + **0**: No mute.
               + **1**: Mute.
        :param str name: Specifies the attendee name or nickname.
        :param str phone: Specifies the SIP or TEL number, maximum of 127 characters.
        :param int role: Specifies the role in the conference. The valid values are as follows:
               + **0**: Normal attendee.
               + **1**: The conference chair.
        :param str sms: Specifies the mobile number for SMS notification, maximum of 32 characters.
        :param str type: Specifies the call-in type. The valid values are as follows:
               + **normal**: The soft terminal.
               + **terminal**: The conference room or hard terminal.
               + **outside**: The outside participant.
               + **mobile**: The user's landline phone.
               + **ideahub**: The ideahub.
        :param str user_id: Specifies the user ID of the participant.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if is_auto_invite is not None:
            pulumi.set(__self__, "is_auto_invite", is_auto_invite)
        if is_mute is not None:
            pulumi.set(__self__, "is_mute", is_mute)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if sms is not None:
            pulumi.set(__self__, "sms", sms)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        """
        Specifies the account ID of the participant.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        """
        Specifies the email address.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="isAutoInvite")
    def is_auto_invite(self) -> Optional[int]:
        """
        Specifies whether to automatically invite this participant when the conference
        starts. The valid values are as follows:
        + **0**: Do not automatically invite.
        + **1**: Automatic invitation.
        """
        return pulumi.get(self, "is_auto_invite")

    @property
    @pulumi.getter(name="isMute")
    def is_mute(self) -> Optional[int]:
        """
        Specifies whether the user needs to be automatically muted when joining the conference
        (only effective when invited in the conference). The valid values are as follows:
        + **0**: No mute.
        + **1**: Mute.
        """
        return pulumi.get(self, "is_mute")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Specifies the attendee name or nickname.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def phone(self) -> Optional[str]:
        """
        Specifies the SIP or TEL number, maximum of 127 characters.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def role(self) -> Optional[int]:
        """
        Specifies the role in the conference. The valid values are as follows:
        + **0**: Normal attendee.
        + **1**: The conference chair.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def sms(self) -> Optional[str]:
        """
        Specifies the mobile number for SMS notification, maximum of 32 characters.
        """
        return pulumi.get(self, "sms")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Specifies the call-in type. The valid values are as follows:
        + **normal**: The soft terminal.
        + **terminal**: The conference room or hard terminal.
        + **outside**: The outside participant.
        + **mobile**: The user's landline phone.
        + **ideahub**: The ideahub.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        """
        Specifies the user ID of the participant.
        """
        return pulumi.get(self, "user_id")


@pulumi.output_type
class ConferenceSubconference(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endTime":
            suggest = "end_time"
        elif key == "isAutoRecord":
            suggest = "is_auto_record"
        elif key == "mediaTypes":
            suggest = "media_types"
        elif key == "recordAuthType":
            suggest = "record_auth_type"
        elif key == "startTime":
            suggest = "start_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceSubconference. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceSubconference.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceSubconference.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_time: Optional[str] = None,
                 id: Optional[str] = None,
                 is_auto_record: Optional[int] = None,
                 media_types: Optional[Sequence[str]] = None,
                 record_auth_type: Optional[int] = None,
                 start_time: Optional[str] = None,
                 subconfigurations: Optional[Sequence['outputs.ConferenceSubconferenceSubconfiguration']] = None):
        """
        :param str end_time: The sub-conference end time.
        :param str id: The sub-conference ID.
        :param int is_auto_record: Specifies whether the conference automatically starts recording, it only takes
               effect when the recording type is:
               + **1**: Automatically start recording.
               + **0**: Do not start recording automatically.
        :param Sequence[str] media_types: Specifies the conference media type list.
               It consists of one or more enumerations, and the valid values are as follows:
               + **Voice**: Voice.
               + **Video**: SD video.
               + **HDVideo**: High-definition video (mutually exclusive with Video, if Video and HDVideo are selected at the same
               time, the system will select Video by default).
               + **Data**: Multimedia (If omitted, the system configuration will determines whether to automatically add **Data**).
        :param int record_auth_type: Specifies the recording authentication method.
               **0**: Viewable/downloadable via link.
               **1**: Enterprise users can watch/download.
               **2**: Attendees can watch/download.
        :param str start_time: Specifies the conference start time (UTC time).
               The time format is `YYYY-MM-DD hh:mm`, e.g. `2006-01-02 15:04`.
               There is no need to set if you book a cyclical conference.
        :param Sequence['ConferenceSubconferenceSubconfigurationArgs'] subconfigurations: The other configuration information of periodic subconferences.
               The object structure is documented below.
        """
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_auto_record is not None:
            pulumi.set(__self__, "is_auto_record", is_auto_record)
        if media_types is not None:
            pulumi.set(__self__, "media_types", media_types)
        if record_auth_type is not None:
            pulumi.set(__self__, "record_auth_type", record_auth_type)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if subconfigurations is not None:
            pulumi.set(__self__, "subconfigurations", subconfigurations)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        The sub-conference end time.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The sub-conference ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAutoRecord")
    def is_auto_record(self) -> Optional[int]:
        """
        Specifies whether the conference automatically starts recording, it only takes
        effect when the recording type is:
        + **1**: Automatically start recording.
        + **0**: Do not start recording automatically.
        """
        return pulumi.get(self, "is_auto_record")

    @property
    @pulumi.getter(name="mediaTypes")
    def media_types(self) -> Optional[Sequence[str]]:
        """
        Specifies the conference media type list.
        It consists of one or more enumerations, and the valid values are as follows:
        + **Voice**: Voice.
        + **Video**: SD video.
        + **HDVideo**: High-definition video (mutually exclusive with Video, if Video and HDVideo are selected at the same
        time, the system will select Video by default).
        + **Data**: Multimedia (If omitted, the system configuration will determines whether to automatically add **Data**).
        """
        return pulumi.get(self, "media_types")

    @property
    @pulumi.getter(name="recordAuthType")
    def record_auth_type(self) -> Optional[int]:
        """
        Specifies the recording authentication method.
        **0**: Viewable/downloadable via link.
        **1**: Enterprise users can watch/download.
        **2**: Attendees can watch/download.
        """
        return pulumi.get(self, "record_auth_type")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        """
        Specifies the conference start time (UTC time).
        The time format is `YYYY-MM-DD hh:mm`, e.g. `2006-01-02 15:04`.
        There is no need to set if you book a cyclical conference.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def subconfigurations(self) -> Optional[Sequence['outputs.ConferenceSubconferenceSubconfiguration']]:
        """
        The other configuration information of periodic subconferences.
        The object structure is documented below.
        """
        return pulumi.get(self, "subconfigurations")


@pulumi.output_type
class ConferenceSubconferenceSubconfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowGuestStart":
            suggest = "allow_guest_start"
        elif key == "audienceCallinRestriction":
            suggest = "audience_callin_restriction"
        elif key == "callinRestriction":
            suggest = "callin_restriction"
        elif key == "showAudiencePolicies":
            suggest = "show_audience_policies"
        elif key == "waitingRoomEnabled":
            suggest = "waiting_room_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceSubconferenceSubconfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceSubconferenceSubconfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceSubconferenceSubconfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_guest_start: Optional[bool] = None,
                 audience_callin_restriction: Optional[int] = None,
                 callin_restriction: Optional[int] = None,
                 show_audience_policies: Optional[Sequence['outputs.ConferenceSubconferenceSubconfigurationShowAudiencePolicy']] = None,
                 waiting_room_enabled: Optional[bool] = None):
        """
        :param bool allow_guest_start: Specifies whether to allow guests to start conferences (only valid for random
               ID conferences).
        :param int audience_callin_restriction: The range that the webinar audience is allowed to call in.
               The valid values are as follows:
               + **0**: All users.
               + **2**: Users within the enterprise.
        :param int callin_restriction: Specifies the range to allow incoming calls.
               + **0**: All users.
               + **2**: Users within the enterprise.
               + **3**: The invited user.
        :param Sequence['ConferenceSubconferenceSubconfigurationShowAudiencePolicyArgs'] show_audience_policies: The webinar Audience Display Strategy.
               The object structure is documented below.
        :param bool waiting_room_enabled: Specifies whether to open the waiting room (only valid for RTC enterprises).
        """
        if allow_guest_start is not None:
            pulumi.set(__self__, "allow_guest_start", allow_guest_start)
        if audience_callin_restriction is not None:
            pulumi.set(__self__, "audience_callin_restriction", audience_callin_restriction)
        if callin_restriction is not None:
            pulumi.set(__self__, "callin_restriction", callin_restriction)
        if show_audience_policies is not None:
            pulumi.set(__self__, "show_audience_policies", show_audience_policies)
        if waiting_room_enabled is not None:
            pulumi.set(__self__, "waiting_room_enabled", waiting_room_enabled)

    @property
    @pulumi.getter(name="allowGuestStart")
    def allow_guest_start(self) -> Optional[bool]:
        """
        Specifies whether to allow guests to start conferences (only valid for random
        ID conferences).
        """
        return pulumi.get(self, "allow_guest_start")

    @property
    @pulumi.getter(name="audienceCallinRestriction")
    def audience_callin_restriction(self) -> Optional[int]:
        """
        The range that the webinar audience is allowed to call in.
        The valid values are as follows:
        + **0**: All users.
        + **2**: Users within the enterprise.
        """
        return pulumi.get(self, "audience_callin_restriction")

    @property
    @pulumi.getter(name="callinRestriction")
    def callin_restriction(self) -> Optional[int]:
        """
        Specifies the range to allow incoming calls.
        + **0**: All users.
        + **2**: Users within the enterprise.
        + **3**: The invited user.
        """
        return pulumi.get(self, "callin_restriction")

    @property
    @pulumi.getter(name="showAudiencePolicies")
    def show_audience_policies(self) -> Optional[Sequence['outputs.ConferenceSubconferenceSubconfigurationShowAudiencePolicy']]:
        """
        The webinar Audience Display Strategy.
        The object structure is documented below.
        """
        return pulumi.get(self, "show_audience_policies")

    @property
    @pulumi.getter(name="waitingRoomEnabled")
    def waiting_room_enabled(self) -> Optional[bool]:
        """
        Specifies whether to open the waiting room (only valid for RTC enterprises).
        """
        return pulumi.get(self, "waiting_room_enabled")


@pulumi.output_type
class ConferenceSubconferenceSubconfigurationShowAudiencePolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "baseAudienceCount":
            suggest = "base_audience_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConferenceSubconferenceSubconfigurationShowAudiencePolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConferenceSubconferenceSubconfigurationShowAudiencePolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConferenceSubconferenceSubconfigurationShowAudiencePolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 base_audience_count: Optional[int] = None,
                 mode: Optional[int] = None,
                 multiple: Optional[float] = None):
        """
        :param int base_audience_count: Specifies the basic number of people, the valid values is range from `0` to `10,000`.
        :param int mode: Audience display strategy: The server is used to calculate the number of audiences and send it to the client
               to control the audience display.
               + **0**: Do not display.
               + **1**: Multiply display the number of participants, based on the real-time number of participants or the cumulative
               number of participants, the multiplication setting can be performed.
        :param float multiple: Specifies the multiplier. The valid values is range from `0` to `10`, it can be set to 1 decimal place.
        """
        if base_audience_count is not None:
            pulumi.set(__self__, "base_audience_count", base_audience_count)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if multiple is not None:
            pulumi.set(__self__, "multiple", multiple)

    @property
    @pulumi.getter(name="baseAudienceCount")
    def base_audience_count(self) -> Optional[int]:
        """
        Specifies the basic number of people, the valid values is range from `0` to `10,000`.
        """
        return pulumi.get(self, "base_audience_count")

    @property
    @pulumi.getter
    def mode(self) -> Optional[int]:
        """
        Audience display strategy: The server is used to calculate the number of audiences and send it to the client
        to control the audience display.
        + **0**: Do not display.
        + **1**: Multiply display the number of participants, based on the real-time number of participants or the cumulative
        number of participants, the multiplication setting can be performed.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def multiple(self) -> Optional[float]:
        """
        Specifies the multiplier. The valid values is range from `0` to `10`, it can be set to 1 decimal place.
        """
        return pulumi.get(self, "multiple")

