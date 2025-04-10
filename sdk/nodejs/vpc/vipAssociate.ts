// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this resource, one or more NICs (to which the ECS instance belongs) can be bound to the VIP.
 *
 * > A VIP can only have one resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vipId = config.requireObject("vipId");
 * const nicPortIds = config.requireObject("nicPortIds");
 * const vipAssociated = new huaweicloud.vpc.VipAssociate("vipAssociated", {
 *     vipId: vipId,
 *     portIds: nicPortIds,
 * });
 * ```
 *
 * ## Import
 *
 * Vip associate can be imported using the `vip_id` and port IDs separated by slashes (no limit on the number of port IDs), e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Vpc/vipAssociate:VipAssociate vip_associated vip_id/port1_id/port2_id
 * ```
 */
export class VipAssociate extends pulumi.CustomResource {
    /**
     * Get an existing VipAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VipAssociateState, opts?: pulumi.CustomResourceOptions): VipAssociate {
        return new VipAssociate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Vpc/vipAssociate:VipAssociate';

    /**
     * Returns true if the given object is an instance of VipAssociate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VipAssociate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VipAssociate.__pulumiType;
    }

    /**
     * The IP addresses of ports to attach the vip to.
     */
    public /*out*/ readonly ipAddresses!: pulumi.Output<string[]>;
    /**
     * An array of one or more IDs of the ports to attach the vip to.
     */
    public readonly portIds!: pulumi.Output<string[]>;
    /**
     * The region in which to create the vip associate resource. If omitted, the
     * provider-level region will be used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The ID of vip to attach the ports to.
     */
    public readonly vipId!: pulumi.Output<string>;
    /**
     * The IP address in the subnet for this vip.
     */
    public /*out*/ readonly vipIpAddress!: pulumi.Output<string>;
    /**
     * The ID of the subnet this vip connects to.
     */
    public /*out*/ readonly vipSubnetId!: pulumi.Output<string>;

    /**
     * Create a VipAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VipAssociateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VipAssociateArgs | VipAssociateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VipAssociateState | undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["portIds"] = state ? state.portIds : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["vipId"] = state ? state.vipId : undefined;
            resourceInputs["vipIpAddress"] = state ? state.vipIpAddress : undefined;
            resourceInputs["vipSubnetId"] = state ? state.vipSubnetId : undefined;
        } else {
            const args = argsOrState as VipAssociateArgs | undefined;
            if ((!args || args.portIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portIds'");
            }
            if ((!args || args.vipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vipId'");
            }
            resourceInputs["portIds"] = args ? args.portIds : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["vipId"] = args ? args.vipId : undefined;
            resourceInputs["ipAddresses"] = undefined /*out*/;
            resourceInputs["vipIpAddress"] = undefined /*out*/;
            resourceInputs["vipSubnetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VipAssociate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VipAssociate resources.
 */
export interface VipAssociateState {
    /**
     * The IP addresses of ports to attach the vip to.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of one or more IDs of the ports to attach the vip to.
     */
    portIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to create the vip associate resource. If omitted, the
     * provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of vip to attach the ports to.
     */
    vipId?: pulumi.Input<string>;
    /**
     * The IP address in the subnet for this vip.
     */
    vipIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the subnet this vip connects to.
     */
    vipSubnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VipAssociate resource.
 */
export interface VipAssociateArgs {
    /**
     * An array of one or more IDs of the ports to attach the vip to.
     */
    portIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to create the vip associate resource. If omitted, the
     * provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of vip to attach the ports to.
     */
    vipId: pulumi.Input<string>;
}
