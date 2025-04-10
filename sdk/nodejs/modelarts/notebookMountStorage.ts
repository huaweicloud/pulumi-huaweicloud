// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage storages mounted to the notebook resource within HuaweiCloud. A maximum of 10 storages can be mounted.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const notebookId = config.requireObject("notebookId");
 * const uriParallelObs = config.requireObject("uriParallelObs");
 * const test = new huaweicloud.modelarts.NotebookMountStorage("test", {
 *     notebookId: notebookId,
 *     storagePath: uriParallelObs,
 *     localMountDirectory: "/data/test/",
 * });
 * ```
 *
 * ## Import
 *
 * The mount storage can be imported by `id`, It is composed of the ID of notebook and mount ID, separated by a slash.
 *
 * For example, bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:ModelArts/notebookMountStorage:NotebookMountStorage test b11b407c-e604-4e8d-8bc4-92398320b847/4e206d3c-6831-4267-b93d-e236105cda38
 * ```
 */
export class NotebookMountStorage extends pulumi.CustomResource {
    /**
     * Get an existing NotebookMountStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotebookMountStorageState, opts?: pulumi.CustomResourceOptions): NotebookMountStorage {
        return new NotebookMountStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:ModelArts/notebookMountStorage:NotebookMountStorage';

    /**
     * Returns true if the given object is an instance of NotebookMountStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotebookMountStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotebookMountStorage.__pulumiType;
    }

    /**
     * Specifies the local mount directory.
     * Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
     * Changing this parameter will create a new resource.
     */
    public readonly localMountDirectory!: pulumi.Output<string>;
    /**
     * The mount ID.
     */
    public /*out*/ readonly mountId!: pulumi.Output<string>;
    /**
     * Specifies ID of notebook which the storage be mounted to.
     * Changing this parameter will create a new resource.
     */
    public readonly notebookId!: pulumi.Output<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
     * `UNMOUNT_FAILED`, `UNMOUNTED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the path of Parallel File System (PFS) or its folders in OBS.
     * The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
     */
    public readonly storagePath!: pulumi.Output<string>;
    /**
     * The type of storage system.  The value is `OBSFS`.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NotebookMountStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotebookMountStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotebookMountStorageArgs | NotebookMountStorageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotebookMountStorageState | undefined;
            resourceInputs["localMountDirectory"] = state ? state.localMountDirectory : undefined;
            resourceInputs["mountId"] = state ? state.mountId : undefined;
            resourceInputs["notebookId"] = state ? state.notebookId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storagePath"] = state ? state.storagePath : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NotebookMountStorageArgs | undefined;
            if ((!args || args.localMountDirectory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localMountDirectory'");
            }
            if ((!args || args.notebookId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notebookId'");
            }
            if ((!args || args.storagePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storagePath'");
            }
            resourceInputs["localMountDirectory"] = args ? args.localMountDirectory : undefined;
            resourceInputs["notebookId"] = args ? args.notebookId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["storagePath"] = args ? args.storagePath : undefined;
            resourceInputs["mountId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotebookMountStorage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotebookMountStorage resources.
 */
export interface NotebookMountStorageState {
    /**
     * Specifies the local mount directory.
     * Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
     * Changing this parameter will create a new resource.
     */
    localMountDirectory?: pulumi.Input<string>;
    /**
     * The mount ID.
     */
    mountId?: pulumi.Input<string>;
    /**
     * Specifies ID of notebook which the storage be mounted to.
     * Changing this parameter will create a new resource.
     */
    notebookId?: pulumi.Input<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
     * `UNMOUNT_FAILED`, `UNMOUNTED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the path of Parallel File System (PFS) or its folders in OBS.
     * The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
     */
    storagePath?: pulumi.Input<string>;
    /**
     * The type of storage system.  The value is `OBSFS`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotebookMountStorage resource.
 */
export interface NotebookMountStorageArgs {
    /**
     * Specifies the local mount directory.
     * Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
     * Changing this parameter will create a new resource.
     */
    localMountDirectory: pulumi.Input<string>;
    /**
     * Specifies ID of notebook which the storage be mounted to.
     * Changing this parameter will create a new resource.
     */
    notebookId: pulumi.Input<string>;
    /**
     * The region in which to create the resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the path of Parallel File System (PFS) or its folders in OBS.
     * The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
     */
    storagePath: pulumi.Input<string>;
}
