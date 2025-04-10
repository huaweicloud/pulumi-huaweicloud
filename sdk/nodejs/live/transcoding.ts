// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Live transcoding resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a transcoding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const ingestDomainName = config.requireObject("ingestDomainName");
 * const appName = config.requireObject("appName");
 * const videoEncoding = config.requireObject("videoEncoding");
 * const templateName = config.requireObject("templateName");
 * const test = new huaweicloud.live.Transcoding("test", {
 *     domainName: ingestDomainName,
 *     appName: appName,
 *     videoEncoding: videoEncoding,
 *     templates: [{
 *         name: templateName,
 *         width: 300,
 *         height: 400,
 *         bitrate: 300,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the `domain_name` and `app_name`, separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Live/transcoding:Transcoding test <domian_name>/<app_name>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`trans_type`. It is generally recommended running `terraform plan` after importing the resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the resource. Also, you can ignore changes as below. hcl resource "huaweicloud_live_transcoding" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  trans_type,
 *
 *  ]
 *
 *  } }
 */
export class Transcoding extends pulumi.CustomResource {
    /**
     * Get an existing Transcoding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TranscodingState, opts?: pulumi.CustomResourceOptions): Transcoding {
        return new Transcoding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Live/transcoding:Transcoding';

    /**
     * Returns true if the given object is an instance of Transcoding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Transcoding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Transcoding.__pulumiType;
    }

    /**
     * Specifies the application name.
     * Changing this parameter will create a new resource.
     */
    public readonly appName!: pulumi.Output<string>;
    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Specifies whether to enable low bitrate HD rates. If enabled
     * the output media will have a lower bitrate with the same image quality. Defaults to **false**.
     */
    public readonly lowBitrateHd!: pulumi.Output<boolean>;
    /**
     * Specifies the region in which to create this resource.
     * If omitted, the provider-level region will be used.
     * Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the video quality templates. A maximum of `4` templates can be added.
     * The templates structure is documented below.
     * For resolution and bitrate settings in the presets,
     * please refer to the [document](https://support.huaweicloud.com/intl/en-us/usermanual-live/live01000802.html).
     */
    public readonly templates!: pulumi.Output<outputs.Live.TranscodingTemplate[]>;
    /**
     * Specifies the transcoding stream trigger mode.
     * The valid values are as follows:
     * + **play**: Pull stream triggers transcoding.
     * + **publish**: Push stream triggers transcoding.
     */
    public readonly transType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the video codec. The valid values are **H264** and **H265**.
     */
    public readonly videoEncoding!: pulumi.Output<string>;

    /**
     * Create a Transcoding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TranscodingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TranscodingArgs | TranscodingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TranscodingState | undefined;
            resourceInputs["appName"] = state ? state.appName : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["lowBitrateHd"] = state ? state.lowBitrateHd : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["templates"] = state ? state.templates : undefined;
            resourceInputs["transType"] = state ? state.transType : undefined;
            resourceInputs["videoEncoding"] = state ? state.videoEncoding : undefined;
        } else {
            const args = argsOrState as TranscodingArgs | undefined;
            if ((!args || args.appName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appName'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.templates === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templates'");
            }
            if ((!args || args.videoEncoding === undefined) && !opts.urn) {
                throw new Error("Missing required property 'videoEncoding'");
            }
            resourceInputs["appName"] = args ? args.appName : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["lowBitrateHd"] = args ? args.lowBitrateHd : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["templates"] = args ? args.templates : undefined;
            resourceInputs["transType"] = args ? args.transType : undefined;
            resourceInputs["videoEncoding"] = args ? args.videoEncoding : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Transcoding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Transcoding resources.
 */
export interface TranscodingState {
    /**
     * Specifies the application name.
     * Changing this parameter will create a new resource.
     */
    appName?: pulumi.Input<string>;
    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Specifies whether to enable low bitrate HD rates. If enabled
     * the output media will have a lower bitrate with the same image quality. Defaults to **false**.
     */
    lowBitrateHd?: pulumi.Input<boolean>;
    /**
     * Specifies the region in which to create this resource.
     * If omitted, the provider-level region will be used.
     * Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the video quality templates. A maximum of `4` templates can be added.
     * The templates structure is documented below.
     * For resolution and bitrate settings in the presets,
     * please refer to the [document](https://support.huaweicloud.com/intl/en-us/usermanual-live/live01000802.html).
     */
    templates?: pulumi.Input<pulumi.Input<inputs.Live.TranscodingTemplate>[]>;
    /**
     * Specifies the transcoding stream trigger mode.
     * The valid values are as follows:
     * + **play**: Pull stream triggers transcoding.
     * + **publish**: Push stream triggers transcoding.
     */
    transType?: pulumi.Input<string>;
    /**
     * Specifies the video codec. The valid values are **H264** and **H265**.
     */
    videoEncoding?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Transcoding resource.
 */
export interface TranscodingArgs {
    /**
     * Specifies the application name.
     * Changing this parameter will create a new resource.
     */
    appName: pulumi.Input<string>;
    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    domainName: pulumi.Input<string>;
    /**
     * Specifies whether to enable low bitrate HD rates. If enabled
     * the output media will have a lower bitrate with the same image quality. Defaults to **false**.
     */
    lowBitrateHd?: pulumi.Input<boolean>;
    /**
     * Specifies the region in which to create this resource.
     * If omitted, the provider-level region will be used.
     * Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the video quality templates. A maximum of `4` templates can be added.
     * The templates structure is documented below.
     * For resolution and bitrate settings in the presets,
     * please refer to the [document](https://support.huaweicloud.com/intl/en-us/usermanual-live/live01000802.html).
     */
    templates: pulumi.Input<pulumi.Input<inputs.Live.TranscodingTemplate>[]>;
    /**
     * Specifies the transcoding stream trigger mode.
     * The valid values are as follows:
     * + **play**: Pull stream triggers transcoding.
     * + **publish**: Push stream triggers transcoding.
     */
    transType?: pulumi.Input<string>;
    /**
     * Specifies the video codec. The valid values are **H264** and **H265**.
     */
    videoEncoding: pulumi.Input<string>;
}
