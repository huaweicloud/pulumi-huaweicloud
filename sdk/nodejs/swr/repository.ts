// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a SWR repository resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const organizationName = config.requireObject("organizationName");
 * const test = new huaweicloud.swr.Repository("test", {
 *     organization: organizationName,
 *     description: "Test repository",
 *     category: "linux",
 * });
 * ```
 *
 * ## Import
 *
 * Repository can be imported using the organization name and repository name separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Swr/repository:Repository test org-name/repo-name
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Swr/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Specifies the category of the repository.
     * The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    /**
     * Specifies the description of the repository.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Intra-cluster image address for docker pull.
     */
    public /*out*/ readonly internalPath!: pulumi.Output<string>;
    /**
     * Specifies whether the repository is public. Default is false.
     */
    public readonly isPublic!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the repository. Changing this creates a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of image tags in a repository.
     */
    public /*out*/ readonly numImages!: pulumi.Output<number>;
    /**
     * Specifies the name of the organization (namespace) the repository belongs.
     * Changing this creates a new resource.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Image address for docker pull.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Numeric ID of the repository
     */
    public /*out*/ readonly repositoryId!: pulumi.Output<number>;
    /**
     * Repository size.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["internalPath"] = state ? state.internalPath : undefined;
            resourceInputs["isPublic"] = state ? state.isPublic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numImages"] = state ? state.numImages : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isPublic"] = args ? args.isPublic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["internalPath"] = undefined /*out*/;
            resourceInputs["numImages"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["repositoryId"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Specifies the category of the repository.
     * The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
     */
    category?: pulumi.Input<string>;
    /**
     * Specifies the description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * Intra-cluster image address for docker pull.
     */
    internalPath?: pulumi.Input<string>;
    /**
     * Specifies whether the repository is public. Default is false.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the repository. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of image tags in a repository.
     */
    numImages?: pulumi.Input<number>;
    /**
     * Specifies the name of the organization (namespace) the repository belongs.
     * Changing this creates a new resource.
     */
    organization?: pulumi.Input<string>;
    /**
     * Image address for docker pull.
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Numeric ID of the repository
     */
    repositoryId?: pulumi.Input<number>;
    /**
     * Repository size.
     */
    size?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Specifies the category of the repository.
     * The value can be `appServer`, `linux`, `frameworkApp`, `database`, `lang`, `other`, `windows`, `arm`.
     */
    category?: pulumi.Input<string>;
    /**
     * Specifies the description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether the repository is public. Default is false.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the repository. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the name of the organization (namespace) the repository belongs.
     * Changing this creates a new resource.
     */
    organization: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
}