// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an IoTDA device linkage rule within HuaweiCloud.
 *
 * > When accessing an IoTDA **standard** or **enterprise** edition instance, you need to specify the IoTDA service
 * endpoint in `provider` block.
 * You can login to the IoTDA console, choose the instance **Overview** and click **Access Details**
 * to view the HTTPS application access address. An example of the access address might be
 * **9bc34xxxxx.st1.iotda-app.ap-southeast-1.myhuaweicloud.com**, then you need to configure the
 * `provider` block as follows:
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const spaceId = config.requireObject("spaceId");
 * const triggerDeviceId = config.requireObject("triggerDeviceId");
 * const actionDeviceId = config.requireObject("actionDeviceId");
 * const topic = new huaweicloud.smn.Topic("topic", {});
 * const test = new huaweicloud.iotda.DeviceLinkageRule("test", {
 *     spaceId: spaceId,
 *     triggers: [
 *         {
 *             type: "SIMPLE_TIMER",
 *             simpleTimerCondition: {
 *                 startTime: "20220622T160000Z",
 *                 repeatInterval: 2,
 *                 repeatCount: 2,
 *             },
 *         },
 *         {
 *             type: "DEVICE_DATA",
 *             deviceDataCondition: {
 *                 deviceId: triggerDeviceId,
 *                 path: "service_id/propertyName_1",
 *                 operator: "=",
 *                 value: "5",
 *                 triggerStrategy: "pulse",
 *                 dataValidatiyPeriod: 300,
 *             },
 *         },
 *         {
 *             type: "DAILY_TIMER",
 *             dailyTimerCondition: {
 *                 startTime: "19:02",
 *             },
 *         },
 *     ],
 *     actions: [
 *         {
 *             type: "SMN_FORWARDING",
 *             smnForwarding: {
 *                 region: topic.region,
 *                 topicName: topic.name,
 *                 topicUrn: topic.topicUrn,
 *                 messageTitle: "message_title",
 *                 messageContent: "message_content",
 *             },
 *         },
 *         {
 *             type: "DEVICE_CMD",
 *             deviceCommand: {
 *                 deviceId: actionDeviceId,
 *                 serviceId: "service_id",
 *                 commandName: "cmd_name",
 *                 commandBody: "{\"cmd_parameter_1\":\"3\"}",
 *             },
 *         },
 *     ],
 *     effectivePeriod: {
 *         startTime: "00:00",
 *         endTime: "23:59",
 *         daysOfWeek: "1,2,3",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Device linkage rules can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:IoTDA/deviceLinkageRule:DeviceLinkageRule test 62b6cc5aa367f403fea86127
 * ```
 */
export class DeviceLinkageRule extends pulumi.CustomResource {
    /**
     * Get an existing DeviceLinkageRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceLinkageRuleState, opts?: pulumi.CustomResourceOptions): DeviceLinkageRule {
        return new DeviceLinkageRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:IoTDA/deviceLinkageRule:DeviceLinkageRule';

    /**
     * Returns true if the given object is an instance of DeviceLinkageRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceLinkageRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceLinkageRule.__pulumiType;
    }

    /**
     * Specifies the list of the actions, at most 10 actions.
     * The actions structure is documented below.
     */
    public readonly actions!: pulumi.Output<outputs.IoTDA.DeviceLinkageRuleAction[]>;
    /**
     * Specifies the description of the alarm.  
     * The value can contain a maximum of `256` characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the effective period of the device linkage rule. Always effectives
     * by default. The effectivePeriod structure is documented below.
     */
    public readonly effectivePeriod!: pulumi.Output<outputs.IoTDA.DeviceLinkageRuleEffectivePeriod | undefined>;
    /**
     * Specifies whether to enable the device linkage rule. Defaults to **true**.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the alarm.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region to which the SMN belongs.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the resource space ID to which the device linkage rule belongs.
     * Changing this parameter will create a new resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * Specifies the logical relationship between multiple triggers.
     * The options are as follows:
     * + **and**: All of the triggers are met.
     * + **or**: Any of the triggers are met.
     */
    public readonly triggerLogic!: pulumi.Output<string | undefined>;
    /**
     * Specifies the list of the triggers, at most 10 triggers.
     * The triggers structure is documented below.
     */
    public readonly triggers!: pulumi.Output<outputs.IoTDA.DeviceLinkageRuleTrigger[]>;

    /**
     * Create a DeviceLinkageRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceLinkageRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceLinkageRuleArgs | DeviceLinkageRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceLinkageRuleState | undefined;
            resourceInputs["actions"] = state ? state.actions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectivePeriod"] = state ? state.effectivePeriod : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["triggerLogic"] = state ? state.triggerLogic : undefined;
            resourceInputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as DeviceLinkageRuleArgs | undefined;
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.spaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spaceId'");
            }
            if ((!args || args.triggers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggers'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["effectivePeriod"] = args ? args.effectivePeriod : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["triggerLogic"] = args ? args.triggerLogic : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceLinkageRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceLinkageRule resources.
 */
export interface DeviceLinkageRuleState {
    /**
     * Specifies the list of the actions, at most 10 actions.
     * The actions structure is documented below.
     */
    actions?: pulumi.Input<pulumi.Input<inputs.IoTDA.DeviceLinkageRuleAction>[]>;
    /**
     * Specifies the description of the alarm.  
     * The value can contain a maximum of `256` characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the effective period of the device linkage rule. Always effectives
     * by default. The effectivePeriod structure is documented below.
     */
    effectivePeriod?: pulumi.Input<inputs.IoTDA.DeviceLinkageRuleEffectivePeriod>;
    /**
     * Specifies whether to enable the device linkage rule. Defaults to **true**.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the alarm.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region to which the SMN belongs.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID to which the device linkage rule belongs.
     * Changing this parameter will create a new resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * Specifies the logical relationship between multiple triggers.
     * The options are as follows:
     * + **and**: All of the triggers are met.
     * + **or**: Any of the triggers are met.
     */
    triggerLogic?: pulumi.Input<string>;
    /**
     * Specifies the list of the triggers, at most 10 triggers.
     * The triggers structure is documented below.
     */
    triggers?: pulumi.Input<pulumi.Input<inputs.IoTDA.DeviceLinkageRuleTrigger>[]>;
}

/**
 * The set of arguments for constructing a DeviceLinkageRule resource.
 */
export interface DeviceLinkageRuleArgs {
    /**
     * Specifies the list of the actions, at most 10 actions.
     * The actions structure is documented below.
     */
    actions: pulumi.Input<pulumi.Input<inputs.IoTDA.DeviceLinkageRuleAction>[]>;
    /**
     * Specifies the description of the alarm.  
     * The value can contain a maximum of `256` characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the effective period of the device linkage rule. Always effectives
     * by default. The effectivePeriod structure is documented below.
     */
    effectivePeriod?: pulumi.Input<inputs.IoTDA.DeviceLinkageRuleEffectivePeriod>;
    /**
     * Specifies whether to enable the device linkage rule. Defaults to **true**.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the alarm.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region to which the SMN belongs.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the resource space ID to which the device linkage rule belongs.
     * Changing this parameter will create a new resource.
     */
    spaceId: pulumi.Input<string>;
    /**
     * Specifies the logical relationship between multiple triggers.
     * The options are as follows:
     * + **and**: All of the triggers are met.
     * + **or**: Any of the triggers are met.
     */
    triggerLogic?: pulumi.Input<string>;
    /**
     * Specifies the list of the triggers, at most 10 triggers.
     * The triggers structure is documented below.
     */
    triggers: pulumi.Input<pulumi.Input<inputs.IoTDA.DeviceLinkageRuleTrigger>[]>;
}
