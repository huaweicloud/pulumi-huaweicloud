// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages DLI SQL database resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a database
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const databaseName = config.requireObject("databaseName");
 * const test = new huaweicloud.dli.Database("test", {});
 * ```
 *
 * ## Import
 *
 * DLI SQL databases can be imported by their `name`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dli/database:Database test terraform_test
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`tags`. It is generally recommended running `terraform plan` after importing a resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the resource. Also you can ignore changes as below. hcl resource "huaweicloud_dataarts_factory_script" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  tags,
 *
 *  ]
 *
 *  } }
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dli/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Specifies the description of a queue.
     * Changing this parameter will create a new database resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the enterprise project ID.
     * The value 0 indicates the default enterprise project. Changing this parameter will create a new database resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies the database name.  
     * The name consists of `1` to `128` characters, starting with a letter or digit.
     * Only letters, digits and underscores (_) are allowed and the name cannot be all digits.
     * Changing this parameter will create a new database resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the SQL database owner.
     * The owner must be IAM user.
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the DLI database resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new database resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the database.  
     * Changing this parameter will create a new resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * Specifies the description of a queue.
     * Changing this parameter will create a new database resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID.
     * The value 0 indicates the default enterprise project. Changing this parameter will create a new database resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the database name.  
     * The name consists of `1` to `128` characters, starting with a letter or digit.
     * Only letters, digits and underscores (_) are allowed and the name cannot be all digits.
     * Changing this parameter will create a new database resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the name of the SQL database owner.
     * The owner must be IAM user.
     */
    owner?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the DLI database resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new database resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the database.  
     * Changing this parameter will create a new resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Specifies the description of a queue.
     * Changing this parameter will create a new database resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID.
     * The value 0 indicates the default enterprise project. Changing this parameter will create a new database resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the database name.  
     * The name consists of `1` to `128` characters, starting with a letter or digit.
     * Only letters, digits and underscores (_) are allowed and the name cannot be all digits.
     * Changing this parameter will create a new database resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the name of the SQL database owner.
     * The owner must be IAM user.
     */
    owner?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the DLI database resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new database resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the database.  
     * Changing this parameter will create a new resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
