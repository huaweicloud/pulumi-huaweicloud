// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a RDS LTS log resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const ltsGroupId = config.requireObject("ltsGroupId");
 * const ltsStreamId = config.requireObject("ltsStreamId");
 * const test = new huaweicloud.rds.LtsLog("test", {
 *     instanceId: instanceId,
 *     engine: "mysql",
 *     logType: "error_log",
 *     ltsGroupId: ltsGroupId,
 *     ltsStreamId: ltsStreamId,
 * });
 * ```
 *
 * ## Import
 *
 * The RDS LTS log can be imported using `instance_id` and `log_type` separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Rds/ltsLog:LtsLog test <instance_id>/<log_type>
 * ```
 */
export class LtsLog extends pulumi.CustomResource {
    /**
     * Get an existing LtsLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LtsLogState, opts?: pulumi.CustomResourceOptions): LtsLog {
        return new LtsLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Rds/ltsLog:LtsLog';

    /**
     * Returns true if the given object is an instance of LtsLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LtsLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LtsLog.__pulumiType;
    }

    /**
     * Specifies the engine of the RDS instance.
     * Value options: **mysql**, **postgresql**, **sqlserver**. Changing this creates a new resource.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Specifies the ID of the RDS instance.
     * Changing this creates a new resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the type of the LTS log.
     * Value options: **error_log**, **slow_log**, **audit_log**. Changing this creates a new resource.
     */
    public readonly logType!: pulumi.Output<string>;
    /**
     * Specifies the ID of the LTS log group.
     */
    public readonly ltsGroupId!: pulumi.Output<string>;
    /**
     * Specifies the ID of the LTS log stream.
     */
    public readonly ltsStreamId!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a LtsLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LtsLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LtsLogArgs | LtsLogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LtsLogState | undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["ltsGroupId"] = state ? state.ltsGroupId : undefined;
            resourceInputs["ltsStreamId"] = state ? state.ltsStreamId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as LtsLogArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.logType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logType'");
            }
            if ((!args || args.ltsGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ltsGroupId'");
            }
            if ((!args || args.ltsStreamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ltsStreamId'");
            }
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["ltsGroupId"] = args ? args.ltsGroupId : undefined;
            resourceInputs["ltsStreamId"] = args ? args.ltsStreamId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LtsLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LtsLog resources.
 */
export interface LtsLogState {
    /**
     * Specifies the engine of the RDS instance.
     * Value options: **mysql**, **postgresql**, **sqlserver**. Changing this creates a new resource.
     */
    engine?: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS instance.
     * Changing this creates a new resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the type of the LTS log.
     * Value options: **error_log**, **slow_log**, **audit_log**. Changing this creates a new resource.
     */
    logType?: pulumi.Input<string>;
    /**
     * Specifies the ID of the LTS log group.
     */
    ltsGroupId?: pulumi.Input<string>;
    /**
     * Specifies the ID of the LTS log stream.
     */
    ltsStreamId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LtsLog resource.
 */
export interface LtsLogArgs {
    /**
     * Specifies the engine of the RDS instance.
     * Value options: **mysql**, **postgresql**, **sqlserver**. Changing this creates a new resource.
     */
    engine: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS instance.
     * Changing this creates a new resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the type of the LTS log.
     * Value options: **error_log**, **slow_log**, **audit_log**. Changing this creates a new resource.
     */
    logType: pulumi.Input<string>;
    /**
     * Specifies the ID of the LTS log group.
     */
    ltsGroupId: pulumi.Input<string>;
    /**
     * Specifies the ID of the LTS log stream.
     */
    ltsStreamId: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
}
