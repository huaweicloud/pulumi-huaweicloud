// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DMS kafka user resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const kafkaInstanceId = config.requireObject("kafkaInstanceId");
 * const userPassword = config.requireObject("userPassword");
 * const user = new huaweicloud.dms.KafkaUser("user", {
 *     instanceId: kafkaInstanceId,
 *     password: userPassword,
 *     description: "test_description",
 * });
 * ```
 *
 * ## Import
 *
 * DMS kafka users can be imported using the kafka instance ID and user name separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dms/kafkaUser:KafkaUser user c8057fe5-23a8-46ef-ad83-c0055b4e0c5c/user_1
 * ```
 */
export class KafkaUser extends pulumi.CustomResource {
    /**
     * Get an existing KafkaUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaUserState, opts?: pulumi.CustomResourceOptions): KafkaUser {
        return new KafkaUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dms/kafkaUser:KafkaUser';

    /**
     * Returns true if the given object is an instance of KafkaUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaUser.__pulumiType;
    }

    /**
     * Indicates the create time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Indicates whether the application is the default application.
     */
    public /*out*/ readonly defaultApp!: pulumi.Output<boolean>;
    /**
     * Specifies the description of the user.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the DMS kafka instance to which the user belongs.
     * Changing this creates a new resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the name of the user. Changing this creates a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the password of the user. The parameter must be 8 to 32 characters
     * long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
     * The value must be different from name.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The region in which to create the DMS kafka user resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Indicates the user role.
     */
    public /*out*/ readonly role!: pulumi.Output<string>;

    /**
     * Create a KafkaUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaUserArgs | KafkaUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaUserState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["defaultApp"] = state ? state.defaultApp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as KafkaUserArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["defaultApp"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KafkaUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaUser resources.
 */
export interface KafkaUserState {
    /**
     * Indicates the create time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Indicates whether the application is the default application.
     */
    defaultApp?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the DMS kafka instance to which the user belongs.
     * Changing this creates a new resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the user. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the user. The parameter must be 8 to 32 characters
     * long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
     * The value must be different from name.
     */
    password?: pulumi.Input<string>;
    /**
     * The region in which to create the DMS kafka user resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Indicates the user role.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KafkaUser resource.
 */
export interface KafkaUserArgs {
    /**
     * Specifies the description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the DMS kafka instance to which the user belongs.
     * Changing this creates a new resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the name of the user. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the user. The parameter must be 8 to 32 characters
     * long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
     * The value must be different from name.
     */
    password: pulumi.Input<string>;
    /**
     * The region in which to create the DMS kafka user resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
}
