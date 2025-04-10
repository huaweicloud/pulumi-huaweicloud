// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a WAF reference table resource within HuaweiCloud.
 *
 * > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
 * used. The reference table resource can be used in Cloud Mode and Dedicated Mode.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const name = config.requireObject("name");
 * const enterpriseProjectId = config.requireObject("enterpriseProjectId");
 * const test = new huaweicloud.waf.ReferenceTable("test", {
 *     type: "url",
 *     enterpriseProjectId: enterpriseProjectId,
 *     conditions: [
 *         "/admin",
 *         "/manage",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * There are two ways to import WAF reference table state. * Using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>
 * ```
 *
 *  * Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/referenceTable:ReferenceTable test <id>/<enterprise_project_id>
 * ```
 */
export class ReferenceTable extends pulumi.CustomResource {
    /**
     * Get an existing ReferenceTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReferenceTableState, opts?: pulumi.CustomResourceOptions): ReferenceTable {
        return new ReferenceTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Waf/referenceTable:ReferenceTable';

    /**
     * Returns true if the given object is an instance of ReferenceTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReferenceTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReferenceTable.__pulumiType;
    }

    /**
     * Specifies the conditions of the reference table.
     */
    public readonly conditions!: pulumi.Output<string[] | undefined>;
    /**
     * The creation time of the reference table, in UTC format.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Specifies the description of the reference table.
     * The maximum length is `128` characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the enterprise project ID to which the reference
     * table belongs. For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the reference table. Only letters, digits, hyphens (-),
     * underscores(_) and dots(.) are allowed. The maximum length is `64` characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the WAF reference table resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the type of the reference table.
     * The valid values are **url**, **user-agent**, **ip**, **params**, **cookie**, **referer** and **header**.
     * Changing this parameter will create a new resource.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ReferenceTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReferenceTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReferenceTableArgs | ReferenceTableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReferenceTableState | undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ReferenceTableArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReferenceTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReferenceTable resources.
 */
export interface ReferenceTableState {
    /**
     * Specifies the conditions of the reference table.
     */
    conditions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The creation time of the reference table, in UTC format.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * Specifies the description of the reference table.
     * The maximum length is `128` characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID to which the reference
     * table belongs. For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the name of the reference table. Only letters, digits, hyphens (-),
     * underscores(_) and dots(.) are allowed. The maximum length is `64` characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the WAF reference table resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the type of the reference table.
     * The valid values are **url**, **user-agent**, **ip**, **params**, **cookie**, **referer** and **header**.
     * Changing this parameter will create a new resource.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReferenceTable resource.
 */
export interface ReferenceTableArgs {
    /**
     * Specifies the conditions of the reference table.
     */
    conditions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the description of the reference table.
     * The maximum length is `128` characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID to which the reference
     * table belongs. For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the name of the reference table. Only letters, digits, hyphens (-),
     * underscores(_) and dots(.) are allowed. The maximum length is `64` characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the WAF reference table resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the type of the reference table.
     * The valid values are **url**, **user-agent**, **ip**, **params**, **cookie**, **referer** and **header**.
     * Changing this parameter will create a new resource.
     */
    type: pulumi.Input<string>;
}
