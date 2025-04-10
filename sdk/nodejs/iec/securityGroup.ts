// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a IEC security group resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const iecSecgroupName = config.requireObject("iecSecgroupName");
 * const secgroupTest = new huaweicloud.iec.SecurityGroup("secgroupTest", {});
 * ```
 *
 * ## Import
 *
 * IEC Security Groups can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iec/securityGroup:SecurityGroup secgroup_test 2a02d1d3-437c-11eb-b721-fa163e8ac569
 * ```
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iec/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * Specifies the description of the iec security group. description must be
     * 0 to 64 characters in length, and does not contain angle brackets (<) and (>). Changing this parameter will creates a
     * new iec security group resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name for the security group. This parameter can contain a maximum
     * of 64 characters, which may consist of letters, digits, dot (.), underscores (_), and hyphens (-). The iec security
     * group allowed to have the same name. Changing this parameter will creates a new iec security group resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An Array of one or more security group rules. The securityGroupRules object structure is
     * documented below.
     */
    public /*out*/ readonly securityGroupRules!: pulumi.Output<outputs.Iec.SecurityGroupSecurityGroupRule[]>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["securityGroupRules"] = state ? state.securityGroupRules : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityGroupRules"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * Specifies the description of the iec security group. description must be
     * 0 to 64 characters in length, and does not contain angle brackets (<) and (>). Changing this parameter will creates a
     * new iec security group resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the name for the security group. This parameter can contain a maximum
     * of 64 characters, which may consist of letters, digits, dot (.), underscores (_), and hyphens (-). The iec security
     * group allowed to have the same name. Changing this parameter will creates a new iec security group resource.
     */
    name?: pulumi.Input<string>;
    /**
     * An Array of one or more security group rules. The securityGroupRules object structure is
     * documented below.
     */
    securityGroupRules?: pulumi.Input<pulumi.Input<inputs.Iec.SecurityGroupSecurityGroupRule>[]>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * Specifies the description of the iec security group. description must be
     * 0 to 64 characters in length, and does not contain angle brackets (<) and (>). Changing this parameter will creates a
     * new iec security group resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the name for the security group. This parameter can contain a maximum
     * of 64 characters, which may consist of letters, digits, dot (.), underscores (_), and hyphens (-). The iec security
     * group allowed to have the same name. Changing this parameter will creates a new iec security group resource.
     */
    name?: pulumi.Input<string>;
}
