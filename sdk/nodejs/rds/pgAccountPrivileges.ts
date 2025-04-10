// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an RDS PostgreSQL account privileges resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const userName = config.requireObject("userName");
 * const test = new huaweicloud.rds.PgAccountPrivileges("test", {
 *     instanceId: instanceId,
 *     userName: userName,
 *     rolePrivileges: [
 *         "CREATEROLE",
 *         "LOGIN",
 *         "REPLICATION",
 *     ],
 *     systemRolePrivileges: [
 *         "pg_signal_backend",
 *         "root",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * The RDS PostgreSQL privileges roles can be imported using the `instance_id` and `user_name` separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Rds/pgAccountPrivileges:PgAccountPrivileges test <instance_id>/<user_name>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`system_role_privileges`. It is generally recommended running `terraform plan` after importing the RDS PostgreSQL account privileges. You can then decide if changes should be applied to the RDS PostgreSQL account privileges, or the RDS PostgreSQL account privileges definition should be updated to align with the account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_pg_account_privileges" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  system_role_privileges,
 *
 *  ]
 *
 *  } }
 */
export class PgAccountPrivileges extends pulumi.CustomResource {
    /**
     * Get an existing PgAccountPrivileges resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PgAccountPrivilegesState, opts?: pulumi.CustomResourceOptions): PgAccountPrivileges {
        return new PgAccountPrivileges(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Rds/pgAccountPrivileges:PgAccountPrivileges';

    /**
     * Returns true if the given object is an instance of PgAccountPrivileges.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PgAccountPrivileges {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PgAccountPrivileges.__pulumiType;
    }

    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the list of role privileges. Value options: **CREATEDB**,
     * **CREATEROLE**, **LOGIN**, **REPLICATION**.
     */
    public readonly rolePrivileges!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the list of system role privileges. Value options:
     * **pg_monitor**, **pg_signal_backend**, **root**.
     */
    public readonly systemRolePrivileges!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the username of the account.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a PgAccountPrivileges resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PgAccountPrivilegesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PgAccountPrivilegesArgs | PgAccountPrivilegesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PgAccountPrivilegesState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rolePrivileges"] = state ? state.rolePrivileges : undefined;
            resourceInputs["systemRolePrivileges"] = state ? state.systemRolePrivileges : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as PgAccountPrivilegesArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rolePrivileges"] = args ? args.rolePrivileges : undefined;
            resourceInputs["systemRolePrivileges"] = args ? args.systemRolePrivileges : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PgAccountPrivileges.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PgAccountPrivileges resources.
 */
export interface PgAccountPrivilegesState {
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the list of role privileges. Value options: **CREATEDB**,
     * **CREATEROLE**, **LOGIN**, **REPLICATION**.
     */
    rolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of system role privileges. Value options:
     * **pg_monitor**, **pg_signal_backend**, **root**.
     */
    systemRolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the username of the account.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PgAccountPrivileges resource.
 */
export interface PgAccountPrivilegesArgs {
    /**
     * Specifies the ID of the RDS PostgreSQL instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the list of role privileges. Value options: **CREATEDB**,
     * **CREATEROLE**, **LOGIN**, **REPLICATION**.
     */
    rolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of system role privileges. Value options:
     * **pg_monitor**, **pg_signal_backend**, **root**.
     */
    systemRolePrivileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the username of the account.
     */
    userName: pulumi.Input<string>;
}
