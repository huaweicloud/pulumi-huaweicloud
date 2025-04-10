// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Dedicated ELB Ip Group resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const basic = new huaweicloud.DedicatedElb.Ipgroup("basic", {
 *     description: "basic example",
 *     ipLists: [{
 *         description: "ECS01",
 *         ip: "192.168.10.10",
 *     }],
 * });
 * ```
 */
export class Ipgroup extends pulumi.CustomResource {
    /**
     * Get an existing Ipgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpgroupState, opts?: pulumi.CustomResourceOptions): Ipgroup {
        return new Ipgroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedElb/ipgroup:Ipgroup';

    /**
     * Returns true if the given object is an instance of Ipgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipgroup.__pulumiType;
    }

    /**
     * The create time of the ip group.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Human-readable description for the ip.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The enterprise project id of the ip group. Changing this
     * creates a new ip group.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more ip addresses. The ipList object structure is
     * documented below.
     */
    public readonly ipLists!: pulumi.Output<outputs.DedicatedElb.IpgroupIpList[]>;
    /**
     * The listener IDs which the ip group associated with.
     */
    public /*out*/ readonly listenerIds!: pulumi.Output<string[]>;
    /**
     * Human-readable name for the ip group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the ip group resource. If omitted, the
     * provider-level region will be used. Changing this creates a new ip group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The update time of the ip group.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Ipgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpgroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpgroupArgs | IpgroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpgroupState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["ipLists"] = state ? state.ipLists : undefined;
            resourceInputs["listenerIds"] = state ? state.listenerIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IpgroupArgs | undefined;
            if ((!args || args.ipLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipLists'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["ipLists"] = args ? args.ipLists : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["listenerIds"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipgroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipgroup resources.
 */
export interface IpgroupState {
    /**
     * The create time of the ip group.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Human-readable description for the ip.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the ip group. Changing this
     * creates a new ip group.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more ip addresses. The ipList object structure is
     * documented below.
     */
    ipLists?: pulumi.Input<pulumi.Input<inputs.DedicatedElb.IpgroupIpList>[]>;
    /**
     * The listener IDs which the ip group associated with.
     */
    listenerIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable name for the ip group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the ip group resource. If omitted, the
     * provider-level region will be used. Changing this creates a new ip group.
     */
    region?: pulumi.Input<string>;
    /**
     * The update time of the ip group.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipgroup resource.
 */
export interface IpgroupArgs {
    /**
     * Human-readable description for the ip.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the ip group. Changing this
     * creates a new ip group.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more ip addresses. The ipList object structure is
     * documented below.
     */
    ipLists: pulumi.Input<pulumi.Input<inputs.DedicatedElb.IpgroupIpList>[]>;
    /**
     * Human-readable name for the ip group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the ip group resource. If omitted, the
     * provider-level region will be used. Changing this creates a new ip group.
     */
    region?: pulumi.Input<string>;
}
