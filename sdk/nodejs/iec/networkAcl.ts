// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a network ACL resource within HuaweiCloud IEC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const iecVpcId = config.requireObject("iecVpcId");
 * const iecSubnetId = config.requireObject("iecSubnetId");
 * const aclTest = new huaweicloud.iec.NetworkAcl("aclTest", {
 *     description: "This is a network ACL of IEC with networks.",
 *     networks: [{
 *         vpcId: iecVpcId,
 *         subnetId: iecSubnetId,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Network ACL can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iec/networkAcl:NetworkAcl acl_test 40b1159c-4e6e-11eb-a00e-fa165c365e51
 * ```
 */
export class NetworkAcl extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclState, opts?: pulumi.CustomResourceOptions): NetworkAcl {
        return new NetworkAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iec/networkAcl:NetworkAcl';

    /**
     * Returns true if the given object is an instance of NetworkAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAcl.__pulumiType;
    }

    /**
     * Specifies the supplementary information about the iec network ACL. This parameter
     * can contain a maximum of 255 characters and cannot contain angle brackets (< or >).
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of the IDs of ingress rules associated with the iec network ACL.
     */
    public /*out*/ readonly inboundRules!: pulumi.Output<string[]>;
    /**
     * Specifies the iec network ACL name. This parameter can contain a maximum of 64 characters,
     * which may consist of letters, digits, dot (.), underscores (_), and hyphens (-).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies an list of one or more networks. The networks object structure is documented
     * below.
     */
    public readonly networks!: pulumi.Output<outputs.Iec.NetworkAclNetwork[] | undefined>;
    /**
     * A list of the IDs of egress rules associated with the iec network ACL.
     */
    public /*out*/ readonly outboundRules!: pulumi.Output<string[]>;
    /**
     * The status of the iec network ACL.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a NetworkAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclArgs | NetworkAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAclState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["inboundRules"] = state ? state.inboundRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["outboundRules"] = state ? state.outboundRules : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as NetworkAclArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["inboundRules"] = undefined /*out*/;
            resourceInputs["outboundRules"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAcl resources.
 */
export interface NetworkAclState {
    /**
     * Specifies the supplementary information about the iec network ACL. This parameter
     * can contain a maximum of 255 characters and cannot contain angle brackets (< or >).
     */
    description?: pulumi.Input<string>;
    /**
     * A list of the IDs of ingress rules associated with the iec network ACL.
     */
    inboundRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the iec network ACL name. This parameter can contain a maximum of 64 characters,
     * which may consist of letters, digits, dot (.), underscores (_), and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an list of one or more networks. The networks object structure is documented
     * below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.Iec.NetworkAclNetwork>[]>;
    /**
     * A list of the IDs of egress rules associated with the iec network ACL.
     */
    outboundRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the iec network ACL.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAcl resource.
 */
export interface NetworkAclArgs {
    /**
     * Specifies the supplementary information about the iec network ACL. This parameter
     * can contain a maximum of 255 characters and cannot contain angle brackets (< or >).
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the iec network ACL name. This parameter can contain a maximum of 64 characters,
     * which may consist of letters, digits, dot (.), underscores (_), and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an list of one or more networks. The networks object structure is documented
     * below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.Iec.NetworkAclNetwork>[]>;
}