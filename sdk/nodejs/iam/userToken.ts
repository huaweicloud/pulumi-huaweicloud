// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IAM user token resource within HuaweiCloud.
 *
 * ->**Note** The token can not be destroyed. It will be invalid after expiration time. If password or AK/SK is changed,
 * the token valid time will last less than 30 minutes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const accountName = config.requireObject("accountName");
 * const userName = config.requireObject("userName");
 * const password = config.requireObject("password");
 * const test = new huaweicloud.iam.UserToken("test", {
 *     accountName: accountName,
 *     userName: userName,
 *     password: password,
 * });
 * ```
 */
export class UserToken extends pulumi.CustomResource {
    /**
     * Get an existing UserToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserTokenState, opts?: pulumi.CustomResourceOptions): UserToken {
        return new UserToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iam/userToken:UserToken';

    /**
     * Returns true if the given object is an instance of UserToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserToken.__pulumiType;
    }

    /**
     * Specifies the account name to which the IAM user belongs.
     * Changing this will create a new token.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The Time when the token will expire. The value is a UTC time in the YYYY-MM-DDTHH:mm:ss.ssssssZ format.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * Specifies the IAM user password. Changing this will create a new token.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the project name. If it is blank, the token applies to global
     * services, otherwise the token applies to project-level services. Changing this will create a new token.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The token. Validity period is 24 hours.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Specifies the IAM user name. Changing this will create a new token.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a UserToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserTokenArgs | UserTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserTokenState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserTokenArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserToken resources.
 */
export interface UserTokenState {
    /**
     * Specifies the account name to which the IAM user belongs.
     * Changing this will create a new token.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The Time when the token will expire. The value is a UTC time in the YYYY-MM-DDTHH:mm:ss.ssssssZ format.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Specifies the IAM user password. Changing this will create a new token.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the project name. If it is blank, the token applies to global
     * services, otherwise the token applies to project-level services. Changing this will create a new token.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The token. Validity period is 24 hours.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies the IAM user name. Changing this will create a new token.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserToken resource.
 */
export interface UserTokenArgs {
    /**
     * Specifies the account name to which the IAM user belongs.
     * Changing this will create a new token.
     */
    accountName: pulumi.Input<string>;
    /**
     * Specifies the IAM user password. Changing this will create a new token.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the project name. If it is blank, the token applies to global
     * services, otherwise the token applies to project-level services. Changing this will create a new token.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Specifies the IAM user name. Changing this will create a new token.
     */
    userName: pulumi.Input<string>;
}
