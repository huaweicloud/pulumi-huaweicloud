// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ## Appendix
 *
 * <a name="timeZoneMapping"></a>
 * The time zone mapping relationship supports:
 *
 * | Timezone ID | Timezone |
 * | ---- | ---- |
 * | 1 | (GMT-12:00) Eniwetok, Kwajalein |
 * | 2 | (GMT-11:00) Midway Island, Samoa |
 * | 3 | (GMT-10:00) Hawaii |
 * | 4 | (GMT-09:00) Alaska |
 * | 5 | (GMT-08:00) Pacific Time(US&Canada);Tijuana |
 * | 6 | (GMT-07:00) Arizona |
 * | 7 | (GMT-07:00) Mountain Time(US&Canada) |
 * | 8 | (GMT-06:00) Central America |
 * | 9 | (GMT-06:00) Central Time(US&Canada) |
 * | 10 | (GMT-06:00) Mexico City |
 * | 11 | (GMT-06:00) Saskatchewan |
 * | 12 | (GMT-05:00) Bogota, Lima, Quito |
 * | 13 | (GMT-05:00) Eastern Time(US&Canada) |
 * | 14 | (GMT-05:00) Indiana(East) |
 * | 15 | (GMT-04:00) Atlantic time(Canada) |
 * | 16 | (GMT-04:00) Caracas, La Paz |
 * | 17 | (GMT-04:00) Santiago |
 * | 18 | (GMT-03:30) Newfoundland |
 * | 19 | (GMT-03:00) Brasilia |
 * | 20 | (GMT-03:00) Buenos Aires, Georgetown |
 * | 21 | (GMT-03:00) Greenland |
 * | 22 | (GMT-02:00) Mid-Atlantic |
 * | 23 | (GMT-01:00) Azores |
 * | 24 | (GMT-01:00) Cape Verde Is. |
 * | 25 | (GMT) Casablanca, Monrovia |
 * | 26 | (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London |
 * | 27 | (GMT+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna |
 * | 28 | (GMT+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague |
 * | 29 | (GMT+01:00) Brussels, Copenhagen, Madrid, Paris |
 * | 30 | (GMT+01:00) Sarajevo, Skopje, Sofija, Warsaw, Zagreb |
 * | 31 | (GMT+01:00) West Central Africa |
 * | 32 | (GMT+02:00) Athens, Istanbul, Vilnius |
 * | 33 | (GMT+02:00) Bucharest |
 * | 34 | (GMT+02:00) Cairo |
 * | 35 | (GMT+02:00) Harare, Pretoria |
 * | 36 | (GMT+02:00) Helsinki, Riga, Tallinn |
 * | 37 | (GMT+02:00) Jerusalem |
 * | 38 | (GMT+03:00) Baghdad, Minsk |
 * | 39 | (GMT+03:00) Kuwait, Riyadh |
 * | 40 | (GMT+03:00) Moscow, St. Petersburg, Volgograd |
 * | 41 | (GMT+03:00) Nairobi |
 * | 42 | (GMT+03:30) Tehran |
 * | 43 | (GMT+04:00) Abu Dhabi, Muscat |
 * | 44 | (GMT+04:00) Baku, Tbilisi, Yerevan |
 * | 45 | (GMT+04:30) Kabul |
 * | 46 | (GMT+05:00) Ekaterinburg |
 * | 47 | (GMT+05:00) Islamabad, Karachi, Tashkent |
 * | 48 | (GMT+05:30) Calcutta, Chennai, Mumbai, New Delhi |
 * | 49 | (GMT+05:45) Kathmandu |
 * | 50 | (GMT+06:00) Almaty, Novosibirsk |
 * | 51 | (GMT+06:00) Astana, Dhaka |
 * | 52 | (GMT+06:00) Sri Jayawardenepura |
 * | 53 | (GMT+06:30) Rangoon |
 * | 54 | (GMT+07:00) Bangkok, Hanoi, Jakarta |
 * | 55 | (GMT+07:00) Krasnoyarsk |
 * | 56 | (GMT+08:00) Beijing, Chongqing, Hong Kong, Urumqi, Taipei |
 * | 57 | (GMT+08:00) Irkutsk, Ulaan Bataar |
 * | 58 | (GMT+08:00) Kuala Lumpur, Singapore |
 * | 59 | (GMT+08:00) Perth |
 * | 60 | (GMT+09:00) Osaka, Sapporo, Tokyo |
 * | 61 | (GMT+09:00) Seoul |
 * | 62 | (GMT+09:00) Yakutsk |
 * | 63 | (GMT+09:30) Adelaide |
 * | 64 | (GMT+09:30) Darwin |
 * | 65 | (GMT+10:00) Brisbane |
 * | 66 | (GMT+10:00) Canberra, Melbourne, Sydney |
 * | 67 | (GMT+10:00) Guam, Port Moresby |
 * | 68 | (GMT+10:00) Hobart |
 * | 69 | (GMT+10:00) Vladivostok |
 * | 70 | (GMT+11:00) Magadan, Solomon Is., New Caledonia |
 * | 71 | (GMT+12:00) Auckland, Welington |
 * | 72 | (GMT+12:00) Fiji |
 * | 73 | (GMT+13:00) Nuku'alofa |
 * | 74 | (GMT+09:00) Irkutsk |
 * | 75 | (GMT) Casablanca |
 * | 76 | (GMT+04:00) Baku |
 * | 77 | (GMT+12:00) Kamchatka, Marshall Is. |
 *
 * ## Import
 *
 * Conferences (only scheduled conference and progressing conference) can be imported using their `id` and authorization parameters, separated by slashes, e.g. Import a conference and authenticated by account. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/conference:Conference test <id>/<account_name>/<account_password>
 * ```
 *
 *  Import a conference and authenticated by `APP ID`/`APP Key`. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/conference:Conference test <id>/<app_id>/<app_key>/<corp_id>/<user_id>
 * ```
 *
 *  The slashes cannot be missing even corporation ID and user ID are empty. Note that importing is not supported for expired conferences and the start time of the meeting is not imported along with it. You can ignore this change as below. hcl resource "huaweicloud_meeting_conference" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  start_time,
 *
 *  ]
 *
 *  } }
 */
export class Conference extends pulumi.CustomResource {
    /**
     * Get an existing Conference resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConferenceState, opts?: pulumi.CustomResourceOptions): Conference {
        return new Conference(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Meeting/conference:Conference';

    /**
     * Returns true if the given object is an instance of Conference.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Conference {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Conference.__pulumiType;
    }

    /**
     * The access number of the conference.
     */
    public /*out*/ readonly accessNumber!: pulumi.Output<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * meeting initiator belongs. Changing this parameter will create a new resource.
     */
    public readonly accountName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    public readonly appId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    public readonly appKey!: pulumi.Output<string | undefined>;
    /**
     * The audience meeting link address.
     */
    public /*out*/ readonly audienceJoinUri!: pulumi.Output<string>;
    /**
     * Specifies the auxiliary streaming address, the maximum length is 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    public readonly auxAddress!: pulumi.Output<string>;
    /**
     * The host meeting link address.
     */
    public /*out*/ readonly chairJoinUri!: pulumi.Output<string>;
    /**
     * The conference type, the valid values are as follows:
     * + **FUTURE**
     * + **IMMEDIATELY**
     * + **CYCLE**
     */
    public /*out*/ readonly conferenceType!: pulumi.Output<string>;
    /**
     * The conference UUID.
     */
    public /*out*/ readonly conferenceUuid!: pulumi.Output<string>;
    /**
     * Specifies the other conference configurations.
     * The object structure is documented below.
     */
    public readonly configuration!: pulumi.Output<outputs.Meeting.ConferenceConfiguration | undefined>;
    /**
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    public readonly corpId!: pulumi.Output<string>;
    /**
     * Specifies the configurations of the cyclical conference.
     * The object structure is documented below.
     */
    public readonly cycleParams!: pulumi.Output<outputs.Meeting.ConferenceCycleParams>;
    /**
     * Specifies the duration of the conference, in minutes.
     * The valid value is range from `15` to `1,440`, defaults to `30`.
     */
    public readonly duration!: pulumi.Output<number>;
    /**
     * Specifies the conference media encryption mode.
     * + **0**: Adaptive encryption.
     * + **1**: Force encryption.
     * + **2**: Do not encrypt.
     */
    public readonly encryptMode!: pulumi.Output<number>;
    /**
     * The common attendee meeting link address.
     */
    public /*out*/ readonly guestJoinUri!: pulumi.Output<string>;
    /**
     * Specifies whether the conference automatically starts recording, it only takes
     * effect when the recording type is:
     * + **1**: Automatically start recording.
     * + **0**: Do not start recording automatically.
     */
    public readonly isAutoRecord!: pulumi.Output<number>;
    /**
     * Specifies whether to record auxiliary stream.
     * + **0**: Do not record.
     * + **1**: Record.
     */
    public readonly isRecordAuxStream!: pulumi.Output<number>;
    /**
     * The meeting password.
     * The joinPassword structure is documented below.
     */
    public /*out*/ readonly joinPasswords!: pulumi.Output<outputs.Meeting.ConferenceJoinPassword[]>;
    /**
     * Specifies the default language of the conference, the default value is defined by the
     * conference cloud service. For languages supported by the system, it is passed according to the RFC3066 specification.
     * The valid values are as follows:
     * + **zh-CN**: Simplified Chinese.
     * + **en-US**: US English.
     */
    public readonly language!: pulumi.Output<string>;
    /**
     * Specifies the mainstream live broadcast address, with a maximum of 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    public readonly liveAddress!: pulumi.Output<string>;
    /**
     * Specifies the conference media type list.
     * It consists of one or more enumerations, and the valid values are as follows:
     * + **Voice**: Voice.
     * + **Video**: SD video.
     * + **HDVideo**: High-definition video (mutually exclusive with Video, if Video and HDVideo are selected at the same
     * time, the system will select Video by default).
     * + **Data**: Multimedia (If omitted, the system configuration will determines whether to automatically add **Data**).
     */
    public readonly mediaTypes!: pulumi.Output<string[]>;
    /**
     * Specifies the cloud meeting room ID.
     */
    public readonly meetingRoomId!: pulumi.Output<string>;
    /**
     * Specifies the number of parties in the conference, the maximum number of
     * participants in the conference. Defaults to `0` (Unlimited).
     */
    public readonly participantNumber!: pulumi.Output<number | undefined>;
    /**
     * Specifies the attendee list.
     * The object structure is documented below.
     */
    public readonly participants!: pulumi.Output<outputs.Meeting.ConferenceParticipant[] | undefined>;
    /**
     * Specifies the recording authentication method.
     * + **0**: Viewable/downloadable via link.
     * + **1**: Enterprise users can watch/download.
     * + **2**: Attendees can watch/download.
     */
    public readonly recordAuthType!: pulumi.Output<number>;
    /**
     * Specifies the recording type.
     * + **0**: Disabled.
     * + **1**: Live broadcast.
     * + **2**: Record and broadcast.
     * + **3**: Live + Recording.
     */
    public readonly recordType!: pulumi.Output<number>;
    /**
     * Specifies the conference start time (UTC time).
     * The time format is `YYYY-MM-DD hh:mm`, e.g. `2006-01-02 15:04`.
     * There is no need to set if you book a cyclical conference.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * The conference status, the valid values are as follows:
     * + **Schedule**: the conference is in schedule.
     * + **Created**: The conference is in progress.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The list of periodic sub-conferences.
     * The object structure is documented below.
     */
    public /*out*/ readonly subconferences!: pulumi.Output<outputs.Meeting.ConferenceSubconference[]>;
    /**
     * Specifies the time zone information of the conference time in the conference
     * notification. For time zone information, refer to the time zone mapping relationship.
     */
    public readonly timezoneId!: pulumi.Output<number>;
    /**
     * Specifies the conference topic. The topic can contain `1` to `128` characters.
     */
    public readonly topic!: pulumi.Output<string>;
    /**
     * Specifies the user ID of the participant.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a Conference resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConferenceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConferenceArgs | ConferenceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConferenceState | undefined;
            resourceInputs["accessNumber"] = state ? state.accessNumber : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["appKey"] = state ? state.appKey : undefined;
            resourceInputs["audienceJoinUri"] = state ? state.audienceJoinUri : undefined;
            resourceInputs["auxAddress"] = state ? state.auxAddress : undefined;
            resourceInputs["chairJoinUri"] = state ? state.chairJoinUri : undefined;
            resourceInputs["conferenceType"] = state ? state.conferenceType : undefined;
            resourceInputs["conferenceUuid"] = state ? state.conferenceUuid : undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["corpId"] = state ? state.corpId : undefined;
            resourceInputs["cycleParams"] = state ? state.cycleParams : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["encryptMode"] = state ? state.encryptMode : undefined;
            resourceInputs["guestJoinUri"] = state ? state.guestJoinUri : undefined;
            resourceInputs["isAutoRecord"] = state ? state.isAutoRecord : undefined;
            resourceInputs["isRecordAuxStream"] = state ? state.isRecordAuxStream : undefined;
            resourceInputs["joinPasswords"] = state ? state.joinPasswords : undefined;
            resourceInputs["language"] = state ? state.language : undefined;
            resourceInputs["liveAddress"] = state ? state.liveAddress : undefined;
            resourceInputs["mediaTypes"] = state ? state.mediaTypes : undefined;
            resourceInputs["meetingRoomId"] = state ? state.meetingRoomId : undefined;
            resourceInputs["participantNumber"] = state ? state.participantNumber : undefined;
            resourceInputs["participants"] = state ? state.participants : undefined;
            resourceInputs["recordAuthType"] = state ? state.recordAuthType : undefined;
            resourceInputs["recordType"] = state ? state.recordType : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subconferences"] = state ? state.subconferences : undefined;
            resourceInputs["timezoneId"] = state ? state.timezoneId : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as ConferenceArgs | undefined;
            if ((!args || args.duration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'duration'");
            }
            if ((!args || args.meetingRoomId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'meetingRoomId'");
            }
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args ? args.accountPassword : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["appKey"] = args ? args.appKey : undefined;
            resourceInputs["auxAddress"] = args ? args.auxAddress : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["corpId"] = args ? args.corpId : undefined;
            resourceInputs["cycleParams"] = args ? args.cycleParams : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["encryptMode"] = args ? args.encryptMode : undefined;
            resourceInputs["isAutoRecord"] = args ? args.isAutoRecord : undefined;
            resourceInputs["isRecordAuxStream"] = args ? args.isRecordAuxStream : undefined;
            resourceInputs["language"] = args ? args.language : undefined;
            resourceInputs["liveAddress"] = args ? args.liveAddress : undefined;
            resourceInputs["mediaTypes"] = args ? args.mediaTypes : undefined;
            resourceInputs["meetingRoomId"] = args ? args.meetingRoomId : undefined;
            resourceInputs["participantNumber"] = args ? args.participantNumber : undefined;
            resourceInputs["participants"] = args ? args.participants : undefined;
            resourceInputs["recordAuthType"] = args ? args.recordAuthType : undefined;
            resourceInputs["recordType"] = args ? args.recordType : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["timezoneId"] = args ? args.timezoneId : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["accessNumber"] = undefined /*out*/;
            resourceInputs["audienceJoinUri"] = undefined /*out*/;
            resourceInputs["chairJoinUri"] = undefined /*out*/;
            resourceInputs["conferenceType"] = undefined /*out*/;
            resourceInputs["conferenceUuid"] = undefined /*out*/;
            resourceInputs["guestJoinUri"] = undefined /*out*/;
            resourceInputs["joinPasswords"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subconferences"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Conference.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Conference resources.
 */
export interface ConferenceState {
    /**
     * The access number of the conference.
     */
    accessNumber?: pulumi.Input<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * meeting initiator belongs. Changing this parameter will create a new resource.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    appId?: pulumi.Input<string>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    appKey?: pulumi.Input<string>;
    /**
     * The audience meeting link address.
     */
    audienceJoinUri?: pulumi.Input<string>;
    /**
     * Specifies the auxiliary streaming address, the maximum length is 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    auxAddress?: pulumi.Input<string>;
    /**
     * The host meeting link address.
     */
    chairJoinUri?: pulumi.Input<string>;
    /**
     * The conference type, the valid values are as follows:
     * + **FUTURE**
     * + **IMMEDIATELY**
     * + **CYCLE**
     */
    conferenceType?: pulumi.Input<string>;
    /**
     * The conference UUID.
     */
    conferenceUuid?: pulumi.Input<string>;
    /**
     * Specifies the other conference configurations.
     * The object structure is documented below.
     */
    configuration?: pulumi.Input<inputs.Meeting.ConferenceConfiguration>;
    /**
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    corpId?: pulumi.Input<string>;
    /**
     * Specifies the configurations of the cyclical conference.
     * The object structure is documented below.
     */
    cycleParams?: pulumi.Input<inputs.Meeting.ConferenceCycleParams>;
    /**
     * Specifies the duration of the conference, in minutes.
     * The valid value is range from `15` to `1,440`, defaults to `30`.
     */
    duration?: pulumi.Input<number>;
    /**
     * Specifies the conference media encryption mode.
     * + **0**: Adaptive encryption.
     * + **1**: Force encryption.
     * + **2**: Do not encrypt.
     */
    encryptMode?: pulumi.Input<number>;
    /**
     * The common attendee meeting link address.
     */
    guestJoinUri?: pulumi.Input<string>;
    /**
     * Specifies whether the conference automatically starts recording, it only takes
     * effect when the recording type is:
     * + **1**: Automatically start recording.
     * + **0**: Do not start recording automatically.
     */
    isAutoRecord?: pulumi.Input<number>;
    /**
     * Specifies whether to record auxiliary stream.
     * + **0**: Do not record.
     * + **1**: Record.
     */
    isRecordAuxStream?: pulumi.Input<number>;
    /**
     * The meeting password.
     * The joinPassword structure is documented below.
     */
    joinPasswords?: pulumi.Input<pulumi.Input<inputs.Meeting.ConferenceJoinPassword>[]>;
    /**
     * Specifies the default language of the conference, the default value is defined by the
     * conference cloud service. For languages supported by the system, it is passed according to the RFC3066 specification.
     * The valid values are as follows:
     * + **zh-CN**: Simplified Chinese.
     * + **en-US**: US English.
     */
    language?: pulumi.Input<string>;
    /**
     * Specifies the mainstream live broadcast address, with a maximum of 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    liveAddress?: pulumi.Input<string>;
    /**
     * Specifies the conference media type list.
     * It consists of one or more enumerations, and the valid values are as follows:
     * + **Voice**: Voice.
     * + **Video**: SD video.
     * + **HDVideo**: High-definition video (mutually exclusive with Video, if Video and HDVideo are selected at the same
     * time, the system will select Video by default).
     * + **Data**: Multimedia (If omitted, the system configuration will determines whether to automatically add **Data**).
     */
    mediaTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the cloud meeting room ID.
     */
    meetingRoomId?: pulumi.Input<string>;
    /**
     * Specifies the number of parties in the conference, the maximum number of
     * participants in the conference. Defaults to `0` (Unlimited).
     */
    participantNumber?: pulumi.Input<number>;
    /**
     * Specifies the attendee list.
     * The object structure is documented below.
     */
    participants?: pulumi.Input<pulumi.Input<inputs.Meeting.ConferenceParticipant>[]>;
    /**
     * Specifies the recording authentication method.
     * + **0**: Viewable/downloadable via link.
     * + **1**: Enterprise users can watch/download.
     * + **2**: Attendees can watch/download.
     */
    recordAuthType?: pulumi.Input<number>;
    /**
     * Specifies the recording type.
     * + **0**: Disabled.
     * + **1**: Live broadcast.
     * + **2**: Record and broadcast.
     * + **3**: Live + Recording.
     */
    recordType?: pulumi.Input<number>;
    /**
     * Specifies the conference start time (UTC time).
     * The time format is `YYYY-MM-DD hh:mm`, e.g. `2006-01-02 15:04`.
     * There is no need to set if you book a cyclical conference.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The conference status, the valid values are as follows:
     * + **Schedule**: the conference is in schedule.
     * + **Created**: The conference is in progress.
     */
    status?: pulumi.Input<string>;
    /**
     * The list of periodic sub-conferences.
     * The object structure is documented below.
     */
    subconferences?: pulumi.Input<pulumi.Input<inputs.Meeting.ConferenceSubconference>[]>;
    /**
     * Specifies the time zone information of the conference time in the conference
     * notification. For time zone information, refer to the time zone mapping relationship.
     */
    timezoneId?: pulumi.Input<number>;
    /**
     * Specifies the conference topic. The topic can contain `1` to `128` characters.
     */
    topic?: pulumi.Input<string>;
    /**
     * Specifies the user ID of the participant.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Conference resource.
 */
export interface ConferenceArgs {
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * meeting initiator belongs. Changing this parameter will create a new resource.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    appId?: pulumi.Input<string>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    appKey?: pulumi.Input<string>;
    /**
     * Specifies the auxiliary streaming address, the maximum length is 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    auxAddress?: pulumi.Input<string>;
    /**
     * Specifies the other conference configurations.
     * The object structure is documented below.
     */
    configuration?: pulumi.Input<inputs.Meeting.ConferenceConfiguration>;
    /**
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    corpId?: pulumi.Input<string>;
    /**
     * Specifies the configurations of the cyclical conference.
     * The object structure is documented below.
     */
    cycleParams?: pulumi.Input<inputs.Meeting.ConferenceCycleParams>;
    /**
     * Specifies the duration of the conference, in minutes.
     * The valid value is range from `15` to `1,440`, defaults to `30`.
     */
    duration: pulumi.Input<number>;
    /**
     * Specifies the conference media encryption mode.
     * + **0**: Adaptive encryption.
     * + **1**: Force encryption.
     * + **2**: Do not encrypt.
     */
    encryptMode?: pulumi.Input<number>;
    /**
     * Specifies whether the conference automatically starts recording, it only takes
     * effect when the recording type is:
     * + **1**: Automatically start recording.
     * + **0**: Do not start recording automatically.
     */
    isAutoRecord?: pulumi.Input<number>;
    /**
     * Specifies whether to record auxiliary stream.
     * + **0**: Do not record.
     * + **1**: Record.
     */
    isRecordAuxStream?: pulumi.Input<number>;
    /**
     * Specifies the default language of the conference, the default value is defined by the
     * conference cloud service. For languages supported by the system, it is passed according to the RFC3066 specification.
     * The valid values are as follows:
     * + **zh-CN**: Simplified Chinese.
     * + **en-US**: US English.
     */
    language?: pulumi.Input<string>;
    /**
     * Specifies the mainstream live broadcast address, with a maximum of 255 characters.
     * Only available if `recordType` is `2` or `3`.
     */
    liveAddress?: pulumi.Input<string>;
    /**
     * Specifies the conference media type list.
     * It consists of one or more enumerations, and the valid values are as follows:
     * + **Voice**: Voice.
     * + **Video**: SD video.
     * + **HDVideo**: High-definition video (mutually exclusive with Video, if Video and HDVideo are selected at the same
     * time, the system will select Video by default).
     * + **Data**: Multimedia (If omitted, the system configuration will determines whether to automatically add **Data**).
     */
    mediaTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the cloud meeting room ID.
     */
    meetingRoomId: pulumi.Input<string>;
    /**
     * Specifies the number of parties in the conference, the maximum number of
     * participants in the conference. Defaults to `0` (Unlimited).
     */
    participantNumber?: pulumi.Input<number>;
    /**
     * Specifies the attendee list.
     * The object structure is documented below.
     */
    participants?: pulumi.Input<pulumi.Input<inputs.Meeting.ConferenceParticipant>[]>;
    /**
     * Specifies the recording authentication method.
     * + **0**: Viewable/downloadable via link.
     * + **1**: Enterprise users can watch/download.
     * + **2**: Attendees can watch/download.
     */
    recordAuthType?: pulumi.Input<number>;
    /**
     * Specifies the recording type.
     * + **0**: Disabled.
     * + **1**: Live broadcast.
     * + **2**: Record and broadcast.
     * + **3**: Live + Recording.
     */
    recordType?: pulumi.Input<number>;
    /**
     * Specifies the conference start time (UTC time).
     * The time format is `YYYY-MM-DD hh:mm`, e.g. `2006-01-02 15:04`.
     * There is no need to set if you book a cyclical conference.
     */
    startTime?: pulumi.Input<string>;
    /**
     * Specifies the time zone information of the conference time in the conference
     * notification. For time zone information, refer to the time zone mapping relationship.
     */
    timezoneId?: pulumi.Input<number>;
    /**
     * Specifies the conference topic. The topic can contain `1` to `128` characters.
     */
    topic: pulumi.Input<string>;
    /**
     * Specifies the user ID of the participant.
     */
    userId?: pulumi.Input<string>;
}
