// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an access rule resource of Scalable File Resource (SFS).
 *
 * ## Example Usage
 * ### Usage in VPC authorization scenarios
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const shareName = config.requireObject("shareName");
 * const vpcId = config.requireObject("vpcId");
 * const share_file = new huaweicloud.sfs.FileSystem("share-file", {
 *     size: 100,
 *     shareProto: "NFS",
 * });
 * const rule1 = new huaweicloud.sfs.AccessRule("rule1", {
 *     sfsId: share_file.id,
 *     accessTo: vpcId,
 * });
 * ```
 *
 * ## Import
 *
 * SFS access rule can be imported by specifying the SFS ID and access rule ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Sfs/accessRule:AccessRule huaweicloud_sfs_access_rule <sfs_id>/<rule_id>
 * ```
 */
export class AccessRule extends pulumi.CustomResource {
    /**
     * Get an existing AccessRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessRuleState, opts?: pulumi.CustomResourceOptions): AccessRule {
        return new AccessRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Sfs/accessRule:AccessRule';

    /**
     * Returns true if the given object is an instance of AccessRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessRule.__pulumiType;
    }

    /**
     * Specifies the access level of the shared file system. Possible values
     * are *ro* (read-only)
     * and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
     */
    public readonly accessLevel!: pulumi.Output<string | undefined>;
    /**
     * Specifies the value that defines the access rule. The value contains 1 to
     * 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
     * + Set the VPC ID in VPC authorization scenarios.
     * + Set this parameter in IP address authorization scenario:
     * - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
     * For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
     */
    public readonly accessTo!: pulumi.Output<string>;
    /**
     * Specifies the type of the share access rule. The default value is *cert*.
     * Changing this will create a new access rule.
     */
    public readonly accessType!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the sfs access rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new access rule resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the UUID of the shared file system. Changing this will create a new
     * access rule.
     */
    public readonly sfsId!: pulumi.Output<string>;
    /**
     * The status of the share access rule.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AccessRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessRuleArgs | AccessRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessRuleState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["accessTo"] = state ? state.accessTo : undefined;
            resourceInputs["accessType"] = state ? state.accessType : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sfsId"] = state ? state.sfsId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccessRuleArgs | undefined;
            if ((!args || args.accessTo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessTo'");
            }
            if ((!args || args.sfsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sfsId'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["accessTo"] = args ? args.accessTo : undefined;
            resourceInputs["accessType"] = args ? args.accessType : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sfsId"] = args ? args.sfsId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessRule resources.
 */
export interface AccessRuleState {
    /**
     * Specifies the access level of the shared file system. Possible values
     * are *ro* (read-only)
     * and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Specifies the value that defines the access rule. The value contains 1 to
     * 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
     * + Set the VPC ID in VPC authorization scenarios.
     * + Set this parameter in IP address authorization scenario:
     * - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
     * For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
     */
    accessTo?: pulumi.Input<string>;
    /**
     * Specifies the type of the share access rule. The default value is *cert*.
     * Changing this will create a new access rule.
     */
    accessType?: pulumi.Input<string>;
    /**
     * The region in which to create the sfs access rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new access rule resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the UUID of the shared file system. Changing this will create a new
     * access rule.
     */
    sfsId?: pulumi.Input<string>;
    /**
     * The status of the share access rule.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessRule resource.
 */
export interface AccessRuleArgs {
    /**
     * Specifies the access level of the shared file system. Possible values
     * are *ro* (read-only)
     * and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Specifies the value that defines the access rule. The value contains 1 to
     * 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
     * + Set the VPC ID in VPC authorization scenarios.
     * + Set this parameter in IP address authorization scenario:
     * - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
     * For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
     */
    accessTo: pulumi.Input<string>;
    /**
     * Specifies the type of the share access rule. The default value is *cert*.
     * Changing this will create a new access rule.
     */
    accessType?: pulumi.Input<string>;
    /**
     * The region in which to create the sfs access rule resource. If omitted, the
     * provider-level region will be used. Changing this creates a new access rule resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the UUID of the shared file system. Changing this will create a new
     * access rule.
     */
    sfsId: pulumi.Input<string>;
}