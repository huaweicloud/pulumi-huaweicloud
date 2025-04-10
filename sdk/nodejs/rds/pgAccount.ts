// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages RDS PostgreSQL account resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const accountPassword = config.requireObject("accountPassword");
 * const test = new huaweicloud.rds.PgAccount("test", {
 *     instanceId: instanceId,
 *     password: accountPassword,
 * });
 * ```
 *
 * ## Import
 *
 * The RDS PostgreSQL account can be imported using the `instance_id` and `name` separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Rds/pgAccount:PgAccount test <instance_id>/<name>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`. It is generally recommended running `terraform plan` after importing the RDS PostgreSQL account. You can then decide if changes should be applied to the RDS PostgreSQL account, or the resource definition should be updated to align with the RDS PostgreSQL account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_pg_account" "account_1" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  password
 *
 *  ]
 *
 *  } }
 */
export class PgAccount extends pulumi.CustomResource {
    /**
     * Get an existing PgAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PgAccountState, opts?: pulumi.CustomResourceOptions): PgAccount {
        return new PgAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Rds/pgAccount:PgAccount';

    /**
     * Returns true if the given object is an instance of PgAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PgAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PgAccount.__pulumiType;
    }

    /**
     * Indicates the permission attributes of a user.
     * The attributes structure is documented below.
     */
    public /*out*/ readonly attributes!: pulumi.Output<outputs.Rds.PgAccountAttribute[]>;
    /**
     * Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * schema: Deprecated
     */
    public readonly memberofs!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the username of the DB account. The username contains 1 to 63
     * characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
     * from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
     * and **rdsDdm**.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the password of the DB account. The value must be 8 to 32 characters long
     * and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a PgAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PgAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PgAccountArgs | PgAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PgAccountState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["memberofs"] = state ? state.memberofs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as PgAccountArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["memberofs"] = args ? args.memberofs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["attributes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PgAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PgAccount resources.
 */
export interface PgAccountState {
    /**
     * Indicates the permission attributes of a user.
     * The attributes structure is documented below.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.Rds.PgAccountAttribute>[]>;
    /**
     * Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    memberofs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the username of the DB account. The username contains 1 to 63
     * characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
     * from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
     * and **rdsDdm**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the DB account. The value must be 8 to 32 characters long
     * and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PgAccount resource.
 */
export interface PgAccountArgs {
    /**
     * Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * schema: Deprecated
     */
    memberofs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the username of the DB account. The username contains 1 to 63
     * characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
     * from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
     * and **rdsDdm**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the DB account. The value must be 8 to 32 characters long
     * and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}
