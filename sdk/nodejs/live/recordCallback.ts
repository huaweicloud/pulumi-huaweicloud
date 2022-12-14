// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a callback configuration within HuaweiCloud Live.
 *
 * > Only one callback configuration can be created for an ingestion domain name.
 *
 * ## Example Usage
 * ### Create a callback configuration for an ingest domain name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const ingestDomainName = config.requireObject("ingestDomainName");
 * const ingestDomain = new huaweicloud.live.Domain("ingestDomain", {type: "push"});
 * const callback = new huaweicloud.live.RecordCallback("callback", {
 *     domainName: ingestDomainName,
 *     url: "http://mycallback.com.cn/record_notify",
 *     types: ["RECORD_NEW_FILE_START"],
 * });
 * ```
 *
 * ## Import
 *
 * Callback configurations can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Live/recordCallback:RecordCallback test 55534eaa-533a-419d-9b40-ec427ea7195a
 * ```
 */
export class RecordCallback extends pulumi.CustomResource {
    /**
     * Get an existing RecordCallback resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordCallbackState, opts?: pulumi.CustomResourceOptions): RecordCallback {
        return new RecordCallback(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Live/recordCallback:RecordCallback';

    /**
     * Returns true if the given object is an instance of RecordCallback.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordCallback {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordCallback.__pulumiType;
    }

    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the types of recording notifications. The options are as follows:
     * + **RECORD_NEW_FILE_START**: Recording started.
     * + **RECORD_FILE_COMPLETE**: Recording file generated.
     * + **RECORD_OVER**: Recording completed.
     * + **RECORD_FAILED**: Recording failed.
     */
    public readonly types!: pulumi.Output<string[]>;
    /**
     * Specifies the callback URL for sending recording notifications, which must start with
     * `http://` or `https://`, and cannot contain message headers or parameters.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a RecordCallback resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordCallbackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordCallbackArgs | RecordCallbackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordCallbackState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["types"] = state ? state.types : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as RecordCallbackArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.types === undefined) && !opts.urn) {
                throw new Error("Missing required property 'types'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["types"] = args ? args.types : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecordCallback.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecordCallback resources.
 */
export interface RecordCallbackState {
    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the types of recording notifications. The options are as follows:
     * + **RECORD_NEW_FILE_START**: Recording started.
     * + **RECORD_FILE_COMPLETE**: Recording file generated.
     * + **RECORD_OVER**: Recording completed.
     * + **RECORD_FAILED**: Recording failed.
     */
    types?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the callback URL for sending recording notifications, which must start with
     * `http://` or `https://`, and cannot contain message headers or parameters.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecordCallback resource.
 */
export interface RecordCallbackArgs {
    /**
     * Specifies the ingest domain name.
     * Changing this parameter will create a new resource.
     */
    domainName: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the types of recording notifications. The options are as follows:
     * + **RECORD_NEW_FILE_START**: Recording started.
     * + **RECORD_FILE_COMPLETE**: Recording file generated.
     * + **RECORD_OVER**: Recording completed.
     * + **RECORD_FAILED**: Recording failed.
     */
    types: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the callback URL for sending recording notifications, which must start with
     * `http://` or `https://`, and cannot contain message headers or parameters.
     */
    url: pulumi.Input<string>;
}
