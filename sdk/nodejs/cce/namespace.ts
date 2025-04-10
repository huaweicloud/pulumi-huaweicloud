// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a CCE namespace resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.requireObject("clusterId");
 * const test = new huaweicloud.cce.Namespace("test", {clusterId: clusterId});
 * ```
 *
 * ## Import
 *
 * CCE namespace can be imported using the cluster ID and namespace name separated by a slash, e.g.bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cce/namespace:Namespace test bb6923e4-b16e-11eb-b0cd-0255ac101da1/test-namespace
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cce/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Specifies an unstructured key value map for external parameters.
     * Changing this will create a new namespace resource.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies the cluster ID to which the CCE namespace belongs.
     * Changing this will create a new namespace resource.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The server time when namespace was created.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new namespace resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies the unique name of the namespace.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies a prefix used by the server to generate a unique name.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * Specifies the region in which to create the namespace resource.
     * If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current phase of the namespace.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["prefix"] = args ? args.prefix : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Specifies an unstructured key value map for external parameters.
     * Changing this will create a new namespace resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the cluster ID to which the CCE namespace belongs.
     * Changing this will create a new namespace resource.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The server time when namespace was created.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new namespace resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the unique name of the namespace.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a prefix used by the server to generate a unique name.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    prefix?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the namespace resource.
     * If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The current phase of the namespace.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Specifies an unstructured key value map for external parameters.
     * Changing this will create a new namespace resource.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the cluster ID to which the CCE namespace belongs.
     * Changing this will create a new namespace resource.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Specifies the map of string keys and values for labels.
     * Changing this will create a new namespace resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the unique name of the namespace.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a prefix used by the server to generate a unique name.  
     * This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
     * hyphens (-), and must start and end with lowercase letters and digits.
     * Changing this will create a new namespace resource.
     */
    prefix?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the namespace resource.
     * If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
     */
    region?: pulumi.Input<string>;
}
