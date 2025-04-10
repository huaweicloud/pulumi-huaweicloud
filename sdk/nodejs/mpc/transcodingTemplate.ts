// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an MPC transcoding template resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = new huaweicloud.Mpc.TranscodingTemplate("test", {
 *     audio: {
 *         bitrate: 0,
 *         channels: 2,
 *         codec: 2,
 *         outputPolicy: "transcode",
 *         sampleRate: 1,
 *     },
 *     dashSegmentDuration: 5,
 *     hlsSegmentDuration: 5,
 *     lowBitrateHd: true,
 *     outputFormat: 1,
 *     video: {
 *         bitrate: 0,
 *         blackBarRemoval: 0,
 *         codec: 2,
 *         fps: 0,
 *         height: 0,
 *         level: 15,
 *         maxConsecutiveBframes: 7,
 *         maxIframesInterval: 5,
 *         outputPolicy: "transcode",
 *         profile: 4,
 *         quality: 1,
 *         width: 0,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * MPC transcoding templates can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Mpc/transcodingTemplate:TranscodingTemplate test 542899
 * ```
 */
export class TranscodingTemplate extends pulumi.CustomResource {
    /**
     * Get an existing TranscodingTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TranscodingTemplateState, opts?: pulumi.CustomResourceOptions): TranscodingTemplate {
        return new TranscodingTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Mpc/transcodingTemplate:TranscodingTemplate';

    /**
     * Returns true if the given object is an instance of TranscodingTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TranscodingTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TranscodingTemplate.__pulumiType;
    }

    /**
     * Specifies the audio parameters. The object structure is documented below.
     */
    public readonly audio!: pulumi.Output<outputs.Mpc.TranscodingTemplateAudio | undefined>;
    /**
     * Specifies the dash segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    public readonly dashSegmentDuration!: pulumi.Output<number | undefined>;
    /**
     * Specifies the HLS segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    public readonly hlsSegmentDuration!: pulumi.Output<number | undefined>;
    /**
     * Specifies Whether to enable low bitrate HD. The default value is false.
     */
    public readonly lowBitrateHd!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of a transcoding template.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the packaging type. Possible values are:
     * + **1**: HLS
     * + **2**: DASH
     * + **3**: HLS+DASH
     * + **4**: MP4
     * + **5**: MP3
     * + **6**: ADTS
     */
    public readonly outputFormat!: pulumi.Output<number>;
    /**
     * The region in which to create the transcoding template resource. If omitted,
     * the provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the video parameters. The object structure is documented below.
     */
    public readonly video!: pulumi.Output<outputs.Mpc.TranscodingTemplateVideo | undefined>;

    /**
     * Create a TranscodingTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TranscodingTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TranscodingTemplateArgs | TranscodingTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TranscodingTemplateState | undefined;
            resourceInputs["audio"] = state ? state.audio : undefined;
            resourceInputs["dashSegmentDuration"] = state ? state.dashSegmentDuration : undefined;
            resourceInputs["hlsSegmentDuration"] = state ? state.hlsSegmentDuration : undefined;
            resourceInputs["lowBitrateHd"] = state ? state.lowBitrateHd : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputFormat"] = state ? state.outputFormat : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["video"] = state ? state.video : undefined;
        } else {
            const args = argsOrState as TranscodingTemplateArgs | undefined;
            if ((!args || args.outputFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outputFormat'");
            }
            resourceInputs["audio"] = args ? args.audio : undefined;
            resourceInputs["dashSegmentDuration"] = args ? args.dashSegmentDuration : undefined;
            resourceInputs["hlsSegmentDuration"] = args ? args.hlsSegmentDuration : undefined;
            resourceInputs["lowBitrateHd"] = args ? args.lowBitrateHd : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputFormat"] = args ? args.outputFormat : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["video"] = args ? args.video : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TranscodingTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TranscodingTemplate resources.
 */
export interface TranscodingTemplateState {
    /**
     * Specifies the audio parameters. The object structure is documented below.
     */
    audio?: pulumi.Input<inputs.Mpc.TranscodingTemplateAudio>;
    /**
     * Specifies the dash segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    dashSegmentDuration?: pulumi.Input<number>;
    /**
     * Specifies the HLS segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    hlsSegmentDuration?: pulumi.Input<number>;
    /**
     * Specifies Whether to enable low bitrate HD. The default value is false.
     */
    lowBitrateHd?: pulumi.Input<boolean>;
    /**
     * Specifies the name of a transcoding template.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the packaging type. Possible values are:
     * + **1**: HLS
     * + **2**: DASH
     * + **3**: HLS+DASH
     * + **4**: MP4
     * + **5**: MP3
     * + **6**: ADTS
     */
    outputFormat?: pulumi.Input<number>;
    /**
     * The region in which to create the transcoding template resource. If omitted,
     * the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the video parameters. The object structure is documented below.
     */
    video?: pulumi.Input<inputs.Mpc.TranscodingTemplateVideo>;
}

/**
 * The set of arguments for constructing a TranscodingTemplate resource.
 */
export interface TranscodingTemplateArgs {
    /**
     * Specifies the audio parameters. The object structure is documented below.
     */
    audio?: pulumi.Input<inputs.Mpc.TranscodingTemplateAudio>;
    /**
     * Specifies the dash segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    dashSegmentDuration?: pulumi.Input<number>;
    /**
     * Specifies the HLS segment duration, in second.  
     * The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
     * The default value is `5`.
     */
    hlsSegmentDuration?: pulumi.Input<number>;
    /**
     * Specifies Whether to enable low bitrate HD. The default value is false.
     */
    lowBitrateHd?: pulumi.Input<boolean>;
    /**
     * Specifies the name of a transcoding template.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the packaging type. Possible values are:
     * + **1**: HLS
     * + **2**: DASH
     * + **3**: HLS+DASH
     * + **4**: MP4
     * + **5**: MP3
     * + **6**: ADTS
     */
    outputFormat: pulumi.Input<number>;
    /**
     * The region in which to create the transcoding template resource. If omitted,
     * the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the video parameters. The object structure is documented below.
     */
    video?: pulumi.Input<inputs.Mpc.TranscodingTemplateVideo>;
}
