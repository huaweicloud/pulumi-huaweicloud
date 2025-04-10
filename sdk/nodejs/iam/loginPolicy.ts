// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = new huaweicloud.Iam.LoginPolicy("test", {
 *     accountValidityPeriod: 20,
 *     customInfoForLogin: "hello Terraform",
 *     lockoutDuration: 30,
 *     loginFailedTimes: 10,
 *     periodWithLoginFailures: 30,
 *     sessionTimeout: 120,
 *     showRecentLoginInfo: true,
 * });
 * ```
 *
 * ## Import
 *
 * Identity login policy can be imported using the account ID or domain ID, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iam/loginPolicy:LoginPolicy example <your account ID>
 * ```
 */
export class LoginPolicy extends pulumi.CustomResource {
    /**
     * Get an existing LoginPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoginPolicyState, opts?: pulumi.CustomResourceOptions): LoginPolicy {
        return new LoginPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iam/loginPolicy:LoginPolicy';

    /**
     * Returns true if the given object is an instance of LoginPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoginPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoginPolicy.__pulumiType;
    }

    /**
     * Specifies the validity period (days) to disable users
     * if they have not logged in within the period. The valid value is range from `0` to `240`.
     */
    public readonly accountValidityPeriod!: pulumi.Output<number>;
    /**
     * Specifies the custom information that will be displayed
     * upon successful login.
     */
    public readonly customInfoForLogin!: pulumi.Output<string | undefined>;
    /**
     * Specifies the duration (minutes) to lock users out.
     * The valid value is range from `15` to `1440`, defaults to `15`.
     */
    public readonly lockoutDuration!: pulumi.Output<number | undefined>;
    /**
     * Specifies the number of unsuccessful login attempts to lock users out.
     * The valid value is range from `3` to `10`, defaults to `5`.
     */
    public readonly loginFailedTimes!: pulumi.Output<number | undefined>;
    /**
     * Specifies the period (minutes) to count the number of unsuccessful
     * login attempts. The valid value is range from `15` to `60`, defaults to `15`.
     */
    public readonly periodWithLoginFailures!: pulumi.Output<number | undefined>;
    /**
     * Specifies the session timeout (minutes) that will apply if you or users created
     * using your account do not perform any operations within a specific period.
     * The valid value is range from `15` to `1,440`, defaults to `60`.
     */
    public readonly sessionTimeout!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether to display last login information upon successful login.
     * The value can be **true** or **false**.
     */
    public readonly showRecentLoginInfo!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LoginPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoginPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoginPolicyArgs | LoginPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoginPolicyState | undefined;
            resourceInputs["accountValidityPeriod"] = state ? state.accountValidityPeriod : undefined;
            resourceInputs["customInfoForLogin"] = state ? state.customInfoForLogin : undefined;
            resourceInputs["lockoutDuration"] = state ? state.lockoutDuration : undefined;
            resourceInputs["loginFailedTimes"] = state ? state.loginFailedTimes : undefined;
            resourceInputs["periodWithLoginFailures"] = state ? state.periodWithLoginFailures : undefined;
            resourceInputs["sessionTimeout"] = state ? state.sessionTimeout : undefined;
            resourceInputs["showRecentLoginInfo"] = state ? state.showRecentLoginInfo : undefined;
        } else {
            const args = argsOrState as LoginPolicyArgs | undefined;
            resourceInputs["accountValidityPeriod"] = args ? args.accountValidityPeriod : undefined;
            resourceInputs["customInfoForLogin"] = args ? args.customInfoForLogin : undefined;
            resourceInputs["lockoutDuration"] = args ? args.lockoutDuration : undefined;
            resourceInputs["loginFailedTimes"] = args ? args.loginFailedTimes : undefined;
            resourceInputs["periodWithLoginFailures"] = args ? args.periodWithLoginFailures : undefined;
            resourceInputs["sessionTimeout"] = args ? args.sessionTimeout : undefined;
            resourceInputs["showRecentLoginInfo"] = args ? args.showRecentLoginInfo : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoginPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoginPolicy resources.
 */
export interface LoginPolicyState {
    /**
     * Specifies the validity period (days) to disable users
     * if they have not logged in within the period. The valid value is range from `0` to `240`.
     */
    accountValidityPeriod?: pulumi.Input<number>;
    /**
     * Specifies the custom information that will be displayed
     * upon successful login.
     */
    customInfoForLogin?: pulumi.Input<string>;
    /**
     * Specifies the duration (minutes) to lock users out.
     * The valid value is range from `15` to `1440`, defaults to `15`.
     */
    lockoutDuration?: pulumi.Input<number>;
    /**
     * Specifies the number of unsuccessful login attempts to lock users out.
     * The valid value is range from `3` to `10`, defaults to `5`.
     */
    loginFailedTimes?: pulumi.Input<number>;
    /**
     * Specifies the period (minutes) to count the number of unsuccessful
     * login attempts. The valid value is range from `15` to `60`, defaults to `15`.
     */
    periodWithLoginFailures?: pulumi.Input<number>;
    /**
     * Specifies the session timeout (minutes) that will apply if you or users created
     * using your account do not perform any operations within a specific period.
     * The valid value is range from `15` to `1,440`, defaults to `60`.
     */
    sessionTimeout?: pulumi.Input<number>;
    /**
     * Specifies whether to display last login information upon successful login.
     * The value can be **true** or **false**.
     */
    showRecentLoginInfo?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LoginPolicy resource.
 */
export interface LoginPolicyArgs {
    /**
     * Specifies the validity period (days) to disable users
     * if they have not logged in within the period. The valid value is range from `0` to `240`.
     */
    accountValidityPeriod?: pulumi.Input<number>;
    /**
     * Specifies the custom information that will be displayed
     * upon successful login.
     */
    customInfoForLogin?: pulumi.Input<string>;
    /**
     * Specifies the duration (minutes) to lock users out.
     * The valid value is range from `15` to `1440`, defaults to `15`.
     */
    lockoutDuration?: pulumi.Input<number>;
    /**
     * Specifies the number of unsuccessful login attempts to lock users out.
     * The valid value is range from `3` to `10`, defaults to `5`.
     */
    loginFailedTimes?: pulumi.Input<number>;
    /**
     * Specifies the period (minutes) to count the number of unsuccessful
     * login attempts. The valid value is range from `15` to `60`, defaults to `15`.
     */
    periodWithLoginFailures?: pulumi.Input<number>;
    /**
     * Specifies the session timeout (minutes) that will apply if you or users created
     * using your account do not perform any operations within a specific period.
     * The valid value is range from `15` to `1,440`, defaults to `60`.
     */
    sessionTimeout?: pulumi.Input<number>;
    /**
     * Specifies whether to display last login information upon successful login.
     * The value can be **true** or **false**.
     */
    showRecentLoginInfo?: pulumi.Input<boolean>;
}
