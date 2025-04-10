// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a SFS Turbo resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a STANDARD Shared File System (SFS) Turbo
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const secgroupId = config.requireObject("secgroupId");
 * const testAz = config.requireObject("testAz");
 * const test = new huaweicloud.sfs.Turbo("test", {
 *     size: 500,
 *     shareProto: "NFS",
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     securityGroupId: secgroupId,
 *     availabilityZone: testAz,
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 * });
 * ```
 * ### Create an HPC Shared File System (SFS) Turbo
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const secgroupId = config.requireObject("secgroupId");
 * const testAz = config.requireObject("testAz");
 * const test = new huaweicloud.sfs.Turbo("test", {
 *     size: 3686,
 *     shareProto: "NFS",
 *     shareType: "HPC",
 *     hpcBandwidth: "40M",
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     securityGroupId: secgroupId,
 *     availabilityZone: testAz,
 * });
 * ```
 * ### Create an HPC CACHE Shared File System (SFS) Turbo
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const secgroupId = config.requireObject("secgroupId");
 * const testAz = config.requireObject("testAz");
 * const test = new huaweicloud.sfs.Turbo("test", {
 *     size: 4096,
 *     shareProto: "NFS",
 *     shareType: "HPC_CACHE",
 *     hpcCacheBandwidth: "2G",
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     securityGroupId: secgroupId,
 *     availabilityZone: testAz,
 * });
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Sfs/turbo:Turbo test <id>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to payment attributes missing from the API response. The missing attributes include`charging_mode`, `period_unit`, `period`, `auto_renew`. It is generally recommended running `terraform plan` after importing an instance. You can ignore changes as below. hcl resource "huaweicloud_sfs_turbo" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  charging_mode, period_unit, period, auto_renew,
 *
 *  ]
 *
 *  } }
 */
export class Turbo extends pulumi.CustomResource {
    /**
     * Get an existing Turbo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TurboState, opts?: pulumi.CustomResourceOptions): Turbo {
        return new Turbo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Sfs/turbo:Turbo';

    /**
     * Returns true if the given object is an instance of Turbo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Turbo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Turbo.__pulumiType;
    }

    /**
     * Specifies whether auto renew is enabled.  
     * The valid values are **true** and **false**.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * Specifies the availability zone where the file system is located.
     * Changing this will create a new resource.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The available capacity of the SFS Turbo file system in the unit of GB.
     */
    public /*out*/ readonly availableCapacity!: pulumi.Output<string>;
    /**
     * Specifies the backup ID.
     */
    public readonly backupId!: pulumi.Output<string>;
    /**
     * Specifies the charging mode of the SFS Turbo.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new cluster resource.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Specifies the ID of a KMS key to encrypt the file system. Changing this
     * will create a new resource.
     */
    public readonly cryptKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the VM flavor used for creating a dedicated file system.
     */
    public readonly dedicatedFlavor!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the dedicated distributed storage used
     * when creating a dedicated file system.
     */
    public readonly dedicatedStorageId!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the file system is enhanced or not. Changing this will
     * create a new resource.
     */
    public readonly enhanced!: pulumi.Output<boolean>;
    /**
     * The enterprise project id of the file system. Changing this
     * will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The mount point of the SFS Turbo file system.
     */
    public /*out*/ readonly exportLocation!: pulumi.Output<string>;
    /**
     * Specifies the HPC bandwidth. Changing this will create a new resource.
     * This parameter is valid and required when `shareType` is set to **HPC**.
     * Valid values are: **20M**, **40M**, **125M**, **250M**, **500M** and **1000M**.
     */
    public readonly hpcBandwidth!: pulumi.Output<string>;
    /**
     * Specifies the HPC cache bandwidth(GB/s).
     * This parameter is valid and required when `shareType` is set to **HPC_CACHE**.
     * Valid values are: **2G**, **4G**, **8G**, **16G**, **24G**, **32G** and **48G**.
     */
    public readonly hpcCacheBandwidth!: pulumi.Output<string>;
    /**
     * Specifies the name of an SFS Turbo file system. The value contains `4` to `64`
     * characters and must start with a letter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the charging period of the SFS Turbo.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `11`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of the SFS Turbo.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the SFS Turbo resource. If omitted, the
     * provider-level region will be used. Changing this creates a new SFS Turbo resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the security group ID.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Specifies the protocol for sharing file systems. The valid value is NFS.
     * Changing this will create a new resource.
     */
    public readonly shareProto!: pulumi.Output<string | undefined>;
    /**
     * Specifies the file system type. Changing this will create a new resource.
     * Valid values are **STANDARD**, **PERFORMANCE**, **HPC** and **HPC_CACHE**.
     * Defaults to **STANDARD**.
     */
    public readonly shareType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the capacity of a sharing file system, in GB.
     * + If `shareType` is set to **STANDARD** or **PERFORMANCE**, the value ranges from `500` to `32,768`, and ranges from
     * `10,240` to `327,680` for an enhanced file system.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The status of the SFS Turbo file system.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the network ID of the subnet. Changing this will create a new
     * resource.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the SFS Turbo.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The version ID of the SFS Turbo file system.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Specifies the VPC ID. Changing this will create a new resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Turbo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TurboArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TurboArgs | TurboState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TurboState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["availableCapacity"] = state ? state.availableCapacity : undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["cryptKeyId"] = state ? state.cryptKeyId : undefined;
            resourceInputs["dedicatedFlavor"] = state ? state.dedicatedFlavor : undefined;
            resourceInputs["dedicatedStorageId"] = state ? state.dedicatedStorageId : undefined;
            resourceInputs["enhanced"] = state ? state.enhanced : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["exportLocation"] = state ? state.exportLocation : undefined;
            resourceInputs["hpcBandwidth"] = state ? state.hpcBandwidth : undefined;
            resourceInputs["hpcCacheBandwidth"] = state ? state.hpcCacheBandwidth : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["shareProto"] = state ? state.shareProto : undefined;
            resourceInputs["shareType"] = state ? state.shareType : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as TurboArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["cryptKeyId"] = args ? args.cryptKeyId : undefined;
            resourceInputs["dedicatedFlavor"] = args ? args.dedicatedFlavor : undefined;
            resourceInputs["dedicatedStorageId"] = args ? args.dedicatedStorageId : undefined;
            resourceInputs["enhanced"] = args ? args.enhanced : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["hpcBandwidth"] = args ? args.hpcBandwidth : undefined;
            resourceInputs["hpcCacheBandwidth"] = args ? args.hpcCacheBandwidth : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["shareProto"] = args ? args.shareProto : undefined;
            resourceInputs["shareType"] = args ? args.shareType : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["availableCapacity"] = undefined /*out*/;
            resourceInputs["exportLocation"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Turbo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Turbo resources.
 */
export interface TurboState {
    /**
     * Specifies whether auto renew is enabled.  
     * The valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone where the file system is located.
     * Changing this will create a new resource.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The available capacity of the SFS Turbo file system in the unit of GB.
     */
    availableCapacity?: pulumi.Input<string>;
    /**
     * Specifies the backup ID.
     */
    backupId?: pulumi.Input<string>;
    /**
     * Specifies the charging mode of the SFS Turbo.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new cluster resource.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies the ID of a KMS key to encrypt the file system. Changing this
     * will create a new resource.
     */
    cryptKeyId?: pulumi.Input<string>;
    /**
     * Specifies the VM flavor used for creating a dedicated file system.
     */
    dedicatedFlavor?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated distributed storage used
     * when creating a dedicated file system.
     */
    dedicatedStorageId?: pulumi.Input<string>;
    /**
     * Specifies whether the file system is enhanced or not. Changing this will
     * create a new resource.
     */
    enhanced?: pulumi.Input<boolean>;
    /**
     * The enterprise project id of the file system. Changing this
     * will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The mount point of the SFS Turbo file system.
     */
    exportLocation?: pulumi.Input<string>;
    /**
     * Specifies the HPC bandwidth. Changing this will create a new resource.
     * This parameter is valid and required when `shareType` is set to **HPC**.
     * Valid values are: **20M**, **40M**, **125M**, **250M**, **500M** and **1000M**.
     */
    hpcBandwidth?: pulumi.Input<string>;
    /**
     * Specifies the HPC cache bandwidth(GB/s).
     * This parameter is valid and required when `shareType` is set to **HPC_CACHE**.
     * Valid values are: **2G**, **4G**, **8G**, **16G**, **24G**, **32G** and **48G**.
     */
    hpcCacheBandwidth?: pulumi.Input<string>;
    /**
     * Specifies the name of an SFS Turbo file system. The value contains `4` to `64`
     * characters and must start with a letter.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the SFS Turbo.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `11`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the SFS Turbo.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The region in which to create the SFS Turbo resource. If omitted, the
     * provider-level region will be used. Changing this creates a new SFS Turbo resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the security group ID.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Specifies the protocol for sharing file systems. The valid value is NFS.
     * Changing this will create a new resource.
     */
    shareProto?: pulumi.Input<string>;
    /**
     * Specifies the file system type. Changing this will create a new resource.
     * Valid values are **STANDARD**, **PERFORMANCE**, **HPC** and **HPC_CACHE**.
     * Defaults to **STANDARD**.
     */
    shareType?: pulumi.Input<string>;
    /**
     * Specifies the capacity of a sharing file system, in GB.
     * + If `shareType` is set to **STANDARD** or **PERFORMANCE**, the value ranges from `500` to `32,768`, and ranges from
     * `10,240` to `327,680` for an enhanced file system.
     */
    size?: pulumi.Input<number>;
    /**
     * The status of the SFS Turbo file system.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the network ID of the subnet. Changing this will create a new
     * resource.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the SFS Turbo.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The version ID of the SFS Turbo file system.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies the VPC ID. Changing this will create a new resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Turbo resource.
 */
export interface TurboArgs {
    /**
     * Specifies whether auto renew is enabled.  
     * The valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone where the file system is located.
     * Changing this will create a new resource.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specifies the backup ID.
     */
    backupId?: pulumi.Input<string>;
    /**
     * Specifies the charging mode of the SFS Turbo.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new cluster resource.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies the ID of a KMS key to encrypt the file system. Changing this
     * will create a new resource.
     */
    cryptKeyId?: pulumi.Input<string>;
    /**
     * Specifies the VM flavor used for creating a dedicated file system.
     */
    dedicatedFlavor?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated distributed storage used
     * when creating a dedicated file system.
     */
    dedicatedStorageId?: pulumi.Input<string>;
    /**
     * Specifies whether the file system is enhanced or not. Changing this will
     * create a new resource.
     */
    enhanced?: pulumi.Input<boolean>;
    /**
     * The enterprise project id of the file system. Changing this
     * will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the HPC bandwidth. Changing this will create a new resource.
     * This parameter is valid and required when `shareType` is set to **HPC**.
     * Valid values are: **20M**, **40M**, **125M**, **250M**, **500M** and **1000M**.
     */
    hpcBandwidth?: pulumi.Input<string>;
    /**
     * Specifies the HPC cache bandwidth(GB/s).
     * This parameter is valid and required when `shareType` is set to **HPC_CACHE**.
     * Valid values are: **2G**, **4G**, **8G**, **16G**, **24G**, **32G** and **48G**.
     */
    hpcCacheBandwidth?: pulumi.Input<string>;
    /**
     * Specifies the name of an SFS Turbo file system. The value contains `4` to `64`
     * characters and must start with a letter.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the SFS Turbo.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `11`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the SFS Turbo.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new cluster resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The region in which to create the SFS Turbo resource. If omitted, the
     * provider-level region will be used. Changing this creates a new SFS Turbo resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the security group ID.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Specifies the protocol for sharing file systems. The valid value is NFS.
     * Changing this will create a new resource.
     */
    shareProto?: pulumi.Input<string>;
    /**
     * Specifies the file system type. Changing this will create a new resource.
     * Valid values are **STANDARD**, **PERFORMANCE**, **HPC** and **HPC_CACHE**.
     * Defaults to **STANDARD**.
     */
    shareType?: pulumi.Input<string>;
    /**
     * Specifies the capacity of a sharing file system, in GB.
     * + If `shareType` is set to **STANDARD** or **PERFORMANCE**, the value ranges from `500` to `32,768`, and ranges from
     * `10,240` to `327,680` for an enhanced file system.
     */
    size: pulumi.Input<number>;
    /**
     * Specifies the network ID of the subnet. Changing this will create a new
     * resource.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the SFS Turbo.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the VPC ID. Changing this will create a new resource.
     */
    vpcId: pulumi.Input<string>;
}
