// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages TMS tags resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = new huaweicloud.Tms.Tags("test", {
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 */
export class Tags extends pulumi.CustomResource {
    /**
     * Get an existing Tags resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagsState, opts?: pulumi.CustomResourceOptions): Tags {
        return new Tags(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Tms/tags:Tags';

    /**
     * Returns true if the given object is an instance of Tags.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tags {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tags.__pulumiType;
    }

    /**
     * Specifies an array of one or more predefined tags.
     * The tags structure is documented below.
     */
    public readonly tags!: pulumi.Output<outputs.Tms.TagsTag[]>;

    /**
     * Create a Tags resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagsArgs | TagsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagsState | undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as TagsArgs | undefined;
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tags.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tags resources.
 */
export interface TagsState {
    /**
     * Specifies an array of one or more predefined tags.
     * The tags structure is documented below.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Tms.TagsTag>[]>;
}

/**
 * The set of arguments for constructing a Tags resource.
 */
export interface TagsArgs {
    /**
     * Specifies an array of one or more predefined tags.
     * The tags structure is documented below.
     */
    tags: pulumi.Input<pulumi.Input<inputs.Tms.TagsTag>[]>;
}
