// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * !> **WARNING:** It has been deprecated, use `huaweicloud.Vpc.NetworkAcl` instead.
 *
 * Manages a network ACL rule resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const rule1 = new huaweicloud.NetworkAcl.AclRule("rule_1", {
 *     action: "deny",
 *     destinationIpAddress: "4.3.2.0/24",
 *     destinationPort: "555",
 *     protocol: "udp",
 *     sourceIpAddress: "1.2.3.4",
 *     sourcePort: "444",
 * });
 * ```
 *
 * ## Import
 *
 * network ACL rules can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:NetworkAcl/aclRule:AclRule rule_1 89a84b28-4cc2-4859-9885-c67e802a46a3
 * ```
 */
export class AclRule extends pulumi.CustomResource {
    /**
     * Get an existing AclRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclRuleState, opts?: pulumi.CustomResourceOptions): AclRule {
        return new AclRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:NetworkAcl/aclRule:AclRule';

    /**
     * Returns true if the given object is an instance of AclRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AclRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AclRule.__pulumiType;
    }

    /**
     * Specifies the action in the network ACL rule. Currently, the value can be *allow* or
     * *deny*.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Specifies the description for the network ACL rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the destination IP address to which the traffic is allowed.
     * The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    public readonly destinationIpAddress!: pulumi.Output<string | undefined>;
    /**
     * Specifies the destination port number or port number range. The value ranges
     * from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    public readonly destinationPort!: pulumi.Output<string | undefined>;
    /**
     * Enabled status for the network ACL rule. Defaults to true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the IP version, either 4 (default) or 6. This parameter is available after
     * the IPv6 function is enabled.
     */
    public readonly ipVersion!: pulumi.Output<number | undefined>;
    /**
     * Specifies a unique name for the network ACL rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
     * *udp*, *icmp* and *any*.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to create the network ACL rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new network ACL rule resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the source IP address that the traffic is allowed from. The default
     * value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    public readonly sourceIpAddress!: pulumi.Output<string | undefined>;
    /**
     * Specifies the source port number or port number range. The value ranges from 1 to
     * 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    public readonly sourcePort!: pulumi.Output<string | undefined>;

    /**
     * Create a AclRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclRuleArgs | AclRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationIpAddress"] = state ? state.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = state ? state.destinationPort : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sourceIpAddress"] = state ? state.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
        } else {
            const args = argsOrState as AclRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationIpAddress"] = args ? args.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = args ? args.destinationPort : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sourceIpAddress"] = args ? args.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AclRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AclRule resources.
 */
export interface AclRuleState {
    /**
     * Specifies the action in the network ACL rule. Currently, the value can be *allow* or
     * *deny*.
     */
    action?: pulumi.Input<string>;
    /**
     * Specifies the description for the network ACL rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the destination IP address to which the traffic is allowed.
     * The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    destinationIpAddress?: pulumi.Input<string>;
    /**
     * Specifies the destination port number or port number range. The value ranges
     * from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the network ACL rule. Defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the IP version, either 4 (default) or 6. This parameter is available after
     * the IPv6 function is enabled.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies a unique name for the network ACL rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
     * *udp*, *icmp* and *any*.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to create the network ACL rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new network ACL rule resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the source IP address that the traffic is allowed from. The default
     * value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    sourceIpAddress?: pulumi.Input<string>;
    /**
     * Specifies the source port number or port number range. The value ranges from 1 to
     * 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    sourcePort?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AclRule resource.
 */
export interface AclRuleArgs {
    /**
     * Specifies the action in the network ACL rule. Currently, the value can be *allow* or
     * *deny*.
     */
    action: pulumi.Input<string>;
    /**
     * Specifies the description for the network ACL rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the destination IP address to which the traffic is allowed.
     * The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    destinationIpAddress?: pulumi.Input<string>;
    /**
     * Specifies the destination port number or port number range. The value ranges
     * from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the network ACL rule. Defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies the IP version, either 4 (default) or 6. This parameter is available after
     * the IPv6 function is enabled.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Specifies a unique name for the network ACL rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
     * *udp*, *icmp* and *any*.
     */
    protocol: pulumi.Input<string>;
    /**
     * The region in which to create the network ACL rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new network ACL rule resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the source IP address that the traffic is allowed from. The default
     * value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
     */
    sourceIpAddress?: pulumi.Input<string>;
    /**
     * Specifies the source port number or port number range. The value ranges from 1 to
     * 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
     */
    sourcePort?: pulumi.Input<string>;
}
