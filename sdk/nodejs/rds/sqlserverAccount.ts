// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages RDS SQLServer account resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const test = new huaweicloud.rds.SqlserverAccount("test", {
 *     instanceId: instanceId,
 *     password: "Test@12345678",
 * });
 * ```
 *
 * ## Import
 *
 * The RDS sqlserver account can be imported using the `instance_id` and `name` separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Rds/sqlserverAccount:SqlserverAccount test <instance_id>/<name>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`. It is generally recommended running `terraform plan` after importing a RDS SQLServer account. You can then decide if changes should be applied to the RDS SQLServer account, or the resource definition should be updated to align with the RDS SQLServer account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_sqlserver_account" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  password,
 *
 *  ]
 *
 *  } }
 */
export class SqlserverAccount extends pulumi.CustomResource {
    /**
     * Get an existing SqlserverAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlserverAccountState, opts?: pulumi.CustomResourceOptions): SqlserverAccount {
        return new SqlserverAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Rds/sqlserverAccount:SqlserverAccount';

    /**
     * Returns true if the given object is an instance of SqlserverAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlserverAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlserverAccount.__pulumiType;
    }

    /**
     * Specifies the ID of the RDS SQLServer instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the username of the DB account. The username consists of 1 to 128
     * characters and must be different from system usernames. System users include **rdsadmin**, **rdsuser**, **rdsbackup**,
     * and **rdsmirror**.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the password of the DB account. It consists of 8 to 128 characters and
     * contains at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Indicates the DB user status. Its value can be any of the following:
     * + **unavailable**: The database user is unavailable.
     * + **available**: The database user is available.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a SqlserverAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlserverAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlserverAccountArgs | SqlserverAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SqlserverAccountState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as SqlserverAccountArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SqlserverAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SqlserverAccount resources.
 */
export interface SqlserverAccountState {
    /**
     * Specifies the ID of the RDS SQLServer instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the username of the DB account. The username consists of 1 to 128
     * characters and must be different from system usernames. System users include **rdsadmin**, **rdsuser**, **rdsbackup**,
     * and **rdsmirror**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the DB account. It consists of 8 to 128 characters and
     * contains at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Indicates the DB user status. Its value can be any of the following:
     * + **unavailable**: The database user is unavailable.
     * + **available**: The database user is available.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SqlserverAccount resource.
 */
export interface SqlserverAccountArgs {
    /**
     * Specifies the ID of the RDS SQLServer instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the username of the DB account. The username consists of 1 to 128
     * characters and must be different from system usernames. System users include **rdsadmin**, **rdsuser**, **rdsbackup**,
     * and **rdsmirror**.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the password of the DB account. It consists of 8 to 128 characters and
     * contains at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
     * characters.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}
