// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB L7 Rule resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const loadbalancer1 = new huaweicloud.elb.Loadbalancer("loadbalancer1", {vipSubnetId: _var.subnet_id});
 * const listener1 = new huaweicloud.elb.Listener("listener1", {
 *     protocol: "HTTP",
 *     protocolPort: 8080,
 *     loadbalancerId: loadbalancer1.id,
 * });
 * const pool1 = new huaweicloud.elb.Pool("pool1", {
 *     protocol: "HTTP",
 *     lbMethod: "ROUND_ROBIN",
 *     loadbalancerId: loadbalancer1.id,
 * });
 * const l7policy1 = new huaweicloud.elb.L7policy("l7policy1", {
 *     action: "REDIRECT_TO_POOL",
 *     description: "test l7 policy",
 *     position: 1,
 *     listenerId: listener1.id,
 *     redirectPoolId: pool1.id,
 * });
 * const l7rule1 = new huaweicloud.elb.L7rule("l7rule1", {
 *     l7policyId: l7policy1.id,
 *     type: "PATH",
 *     compareType: "EQUAL_TO",
 *     value: "/api",
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Elb/l7rule:L7rule l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
 * ```
 */
export class L7rule extends pulumi.CustomResource {
    /**
     * Get an existing L7rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L7ruleState, opts?: pulumi.CustomResourceOptions): L7rule {
        return new L7rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Elb/l7rule:L7rule';

    /**
     * Returns true if the given object is an instance of L7rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L7rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L7rule.__pulumiType;
    }

    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The comparison type for the L7 rule - can either be STARTS_WITH, EQUAL_TO or REGEX
     */
    public readonly compareType!: pulumi.Output<string>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER. Changing this creates a new L7 Rule.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new L7 Rule.
     */
    public readonly l7policyId!: pulumi.Output<string>;
    /**
     * The ID of the Listener owning this resource.
     */
    public /*out*/ readonly listenerId!: pulumi.Output<string>;
    /**
     * The region in which to create the L7 Rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Rule.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The L7 Rule type - can either be HOST_NAME or PATH. Changing this creates a new
     * L7 Rule.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value to use for the comparison. For example, the file type to compare.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a L7rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L7ruleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L7ruleArgs | L7ruleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L7ruleState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["compareType"] = state ? state.compareType : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["l7policyId"] = state ? state.l7policyId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as L7ruleArgs | undefined;
            if ((!args || args.compareType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compareType'");
            }
            if ((!args || args.l7policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'l7policyId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["compareType"] = args ? args.compareType : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["l7policyId"] = args ? args.l7policyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["listenerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L7rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L7rule resources.
 */
export interface L7ruleState {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The comparison type for the L7 rule - can either be STARTS_WITH, EQUAL_TO or REGEX
     */
    compareType?: pulumi.Input<string>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER. Changing this creates a new L7 Rule.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new L7 Rule.
     */
    l7policyId?: pulumi.Input<string>;
    /**
     * The ID of the Listener owning this resource.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The region in which to create the L7 Rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Rule.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The L7 Rule type - can either be HOST_NAME or PATH. Changing this creates a new
     * L7 Rule.
     */
    type?: pulumi.Input<string>;
    /**
     * The value to use for the comparison. For example, the file type to compare.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a L7rule resource.
 */
export interface L7ruleArgs {
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The comparison type for the L7 rule - can either be STARTS_WITH, EQUAL_TO or REGEX
     */
    compareType: pulumi.Input<string>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER. Changing this creates a new L7 Rule.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new L7 Rule.
     */
    l7policyId: pulumi.Input<string>;
    /**
     * The region in which to create the L7 Rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Rule.
     */
    region?: pulumi.Input<string>;
    /**
     * @deprecated tenant_id is deprecated
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The L7 Rule type - can either be HOST_NAME or PATH. Changing this creates a new
     * L7 Rule.
     */
    type: pulumi.Input<string>;
    /**
     * The value to use for the comparison. For example, the file type to compare.
     */
    value: pulumi.Input<string>;
}