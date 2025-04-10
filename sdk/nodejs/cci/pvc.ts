// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a CCI Persistent Volume Claim resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Import an EVS volume
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const volumeId = config.requireObject("volumeId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const test = new huaweicloud.cci.Pvc("test", {
 *     namespace: namespace,
 *     volumeType: "ssd",
 *     volumeId: volumeId,
 * });
 * ```
 * ### Import an OBS bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const obsBucketName = config.requireObject("obsBucketName");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const test = new huaweicloud.cci.Pvc("test", {
 *     namespace: namespace,
 *     volumeType: "obs",
 *     volumeId: obsBucketName,
 * });
 * ```
 * ### Import an SFS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const sfsId = config.requireObject("sfsId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const exportLocation = config.requireObject("exportLocation");
 * const test = new huaweicloud.cci.Pvc("test", {
 *     namespace: namespace,
 *     volumeType: "nfs-rw",
 *     volumeId: sfsId,
 *     deviceMountPath: exportLocation,
 * });
 * ```
 * ### Import an SFS Turbo
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const sfsTurboId = config.requireObject("sfsTurboId");
 * const namespace = config.requireObject("namespace");
 * const pvcName = config.requireObject("pvcName");
 * const exportLocation = config.requireObject("exportLocation");
 * const test = new huaweicloud.cci.Pvc("test", {
 *     namespace: namespace,
 *     volumeType: "efs-standard",
 *     volumeId: sfsTurboId,
 *     deviceMountPath: exportLocation,
 * });
 * ```
 *
 * ## Import
 *
 * PVCs can be imported using the `namespace`, `volume_type` and `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cci/pvc:Pvc test <namespace>/<volume_type>/<id>
 * ```
 */
export class Pvc extends pulumi.CustomResource {
    /**
     * Get an existing Pvc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PvcState, opts?: pulumi.CustomResourceOptions): Pvc {
        return new Pvc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cci/pvc:Pvc';

    /**
     * Returns true if the given object is an instance of Pvc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pvc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pvc.__pulumiType;
    }

    /**
     * The access mode the volume should have.
     */
    public /*out*/ readonly accessModes!: pulumi.Output<string[]>;
    /**
     * The server time when PVC was created.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Specifies the share path of the SFS storage bound to the CCI
     * Namespace. Required if `volumeType` is **nfs-rw**, **efs-standard** or **efs-performance**.
     * Changing this will create a new PVC resource.
     */
    public readonly deviceMountPath!: pulumi.Output<string>;
    /**
     * Whether the PVC is available.
     */
    public /*out*/ readonly enable!: pulumi.Output<boolean>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens, and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the namespace to logically divide your cloud container instances
     * into different group. Changing this will create a new PVC resource.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the PVC resource. If omitted, the
     * provider-level region will be used. Changing this will create a new PVC resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current phase of the PVC.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the ID of the storage bound to the CCI Namespace. Changing this
     * will create a new PVC resource.
     */
    public readonly volumeId!: pulumi.Output<string>;
    /**
     * Specifies the type of the storage bound to the CCI Namespace. The valid
     * values are **sas**, **ssd**, **sata**, **obs**, **nfs-rw**, **efs-standard** and **efs-performance**,
     * Default to **sas**. Changing this will create a new PVC resource.
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a Pvc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PvcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PvcArgs | PvcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PvcState | undefined;
            resourceInputs["accessModes"] = state ? state.accessModes : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["deviceMountPath"] = state ? state.deviceMountPath : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as PvcArgs | undefined;
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["deviceMountPath"] = args ? args.deviceMountPath : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["accessModes"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["enable"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pvc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pvc resources.
 */
export interface PvcState {
    /**
     * The access mode the volume should have.
     */
    accessModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The server time when PVC was created.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * Specifies the share path of the SFS storage bound to the CCI
     * Namespace. Required if `volumeType` is **nfs-rw**, **efs-standard** or **efs-performance**.
     * Changing this will create a new PVC resource.
     */
    deviceMountPath?: pulumi.Input<string>;
    /**
     * Whether the PVC is available.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens, and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the namespace to logically divide your cloud container instances
     * into different group. Changing this will create a new PVC resource.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the PVC resource. If omitted, the
     * provider-level region will be used. Changing this will create a new PVC resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The current phase of the PVC.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the ID of the storage bound to the CCI Namespace. Changing this
     * will create a new PVC resource.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * Specifies the type of the storage bound to the CCI Namespace. The valid
     * values are **sas**, **ssd**, **sata**, **obs**, **nfs-rw**, **efs-standard** and **efs-performance**,
     * Default to **sas**. Changing this will create a new PVC resource.
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pvc resource.
 */
export interface PvcArgs {
    /**
     * Specifies the share path of the SFS storage bound to the CCI
     * Namespace. Required if `volumeType` is **nfs-rw**, **efs-standard** or **efs-performance**.
     * Changing this will create a new PVC resource.
     */
    deviceMountPath?: pulumi.Input<string>;
    /**
     * Specifies the unique name of the PVC resource. This parameter can contain a
     * maximum of 63 characters, which may consist of lowercase letters, digits and hyphens, and must start and end with
     * lowercase letters and digits. Changing this will create a new PVC resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the namespace to logically divide your cloud container instances
     * into different group. Changing this will create a new PVC resource.
     */
    namespace: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the PVC resource. If omitted, the
     * provider-level region will be used. Changing this will create a new PVC resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of the storage bound to the CCI Namespace. Changing this
     * will create a new PVC resource.
     */
    volumeId: pulumi.Input<string>;
    /**
     * Specifies the type of the storage bound to the CCI Namespace. The valid
     * values are **sas**, **ssd**, **sata**, **obs**, **nfs-rw**, **efs-standard** and **efs-performance**,
     * Default to **sas**. Changing this will create a new PVC resource.
     */
    volumeType?: pulumi.Input<string>;
}
