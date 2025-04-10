// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a IEC server resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Server Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const iecServerName = config.requireObject("iecServerName");
 * const iecIamgeId = config.requireObject("iecIamgeId");
 * const iecFlavorId = config.requireObject("iecFlavorId");
 * const iecSiteId = config.requireObject("iecSiteId");
 * const iecSiteOperator = config.requireObject("iecSiteOperator");
 * const iecVpcId = config.requireObject("iecVpcId");
 * const iecSubnetId = config.requireObject("iecSubnetId");
 * const iecSecgroupId = config.requireObject("iecSecgroupId");
 * const iecServerPassword = config.requireObject("iecServerPassword");
 * const serverTest = new huaweicloud.iec.Server("serverTest", {
 *     imageId: iecIamgeId,
 *     flavorId: iecFlavorId,
 *     vpcId: iecVpcId,
 *     subnetIds: [iecSubnetId],
 *     securityGroups: [iecSecgroupId],
 *     adminPass: iecServerPassword,
 *     bindEip: true,
 *     systemDiskType: "SAS",
 *     systemDiskSize: 40,
 *     coverageSites: [{
 *         siteId: iecSiteId,
 *         operator: iecSiteOperator,
 *     }],
 * });
 * ```
 * ### Server Instance With Multiple Data Disks
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const iecServerName = config.requireObject("iecServerName");
 * const iecIamgeId = config.requireObject("iecIamgeId");
 * const iecFlavorId = config.requireObject("iecFlavorId");
 * const iecSiteId = config.requireObject("iecSiteId");
 * const iecSiteOperator = config.requireObject("iecSiteOperator");
 * const iecVpcId = config.requireObject("iecVpcId");
 * const iecSubnetId = config.requireObject("iecSubnetId");
 * const iecSecgroupId = config.requireObject("iecSecgroupId");
 * const iecServerPassword = config.requireObject("iecServerPassword");
 * const serverTest = new huaweicloud.iec.Server("serverTest", {
 *     imageId: iecIamgeId,
 *     flavorId: iecFlavorId,
 *     vpcId: iecVpcId,
 *     subnetIds: [iecSubnetId],
 *     securityGroups: [iecSecgroupId],
 *     adminPass: iecServerPassword,
 *     bindEip: true,
 *     systemDiskType: "SAS",
 *     systemDiskSize: 40,
 *     dataDisks: [
 *         {
 *             type: "SAS",
 *             size: 20,
 *         },
 *         {
 *             type: "SAS",
 *             size: 40,
 *         },
 *     ],
 *     coverageSites: [{
 *         siteId: iecSiteId,
 *         operator: iecSiteOperator,
 *     }],
 * });
 * ```
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iec/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * Specifies the administrative password to assign to the IEC server. This
     * parameter can contain a maximum of 26 characters, which may consist of letters, digits and Special characters(~!?,.:
     * ;-_'"(){}[]/<>@#$%^&*+|\\=) and space. This parameter and `keyPair` are alternative. Changing this changes the root
     * password on the existing server.
     */
    public readonly adminPass!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the IEC server is bound to EIP. Changing this parameter
     * creates a new IEC server resource.
     */
    public readonly bindEip!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the coverage level of IEC sites. Valid value is *SITE*.
     * Changing this parameter creates a new IEC server resource.
     */
    public readonly coverageLevel!: pulumi.Output<string | undefined>;
    /**
     * Specifies the policy of IEC sites. Valid values are *centralize*
     * and *discrete*, *centralize* is default. Changing this parameter creates a new IEC server resource.
     */
    public readonly coveragePolicy!: pulumi.Output<string | undefined>;
    /**
     * Specifies an array of site ID and operator for the IEC server. The
     * object structure is documented below. Changing this parameter creates a new IEC server resource.
     */
    public readonly coverageSites!: pulumi.Output<outputs.Iec.ServerCoverageSite[]>;
    /**
     * Specifies the array of data disks to attach to the IEC server. Up to two
     * data disks can be specified. The object structure is documented below. Changing this parameter creates a new IEC
     * server resource.
     */
    public readonly dataDisks!: pulumi.Output<outputs.Iec.ServerDataDisk[] | undefined>;
    /**
     * The ID of the edgecloud service.
     */
    public /*out*/ readonly edgecloudId!: pulumi.Output<string>;
    /**
     * The Name of the edgecloud service.
     */
    public /*out*/ readonly edgecloudName!: pulumi.Output<string>;
    /**
     * Specifies the flavor ID of the desired flavor for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The flavor name of the IEC server.
     */
    public /*out*/ readonly flavorName!: pulumi.Output<string>;
    /**
     * Specifies the image ID of the desired image for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The image name of the IEC server.
     */
    public /*out*/ readonly imageName!: pulumi.Output<string>;
    /**
     * Specifies the name of a key pair to put on the IEC server. The key pair must
     * already be created and associated with the tenant's account. This parameter and `adminPass` are alternative. Changing
     * this parameter creates a new IEC server resource.
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * Specifies the IEC server name. This parameter can contain a maximum of 64
     * characters, which may consist of letters, digits, dot(.), underscores (_), and hyphens (-).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of one or more networks to attach to the IEC server. The object structure is documented below.
     */
    public /*out*/ readonly nics!: pulumi.Output<outputs.Iec.ServerNic[]>;
    /**
     * The ID of origin server.
     */
    public /*out*/ readonly originServerId!: pulumi.Output<string>;
    /**
     * The EIP address that is associted to the IEC server.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the IEC server. Changing this parameter creates a new IEC server resource.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The status of IEC server.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more subnet ID of Network for the IEC server
     * binding. Changing this parameter creates a new IEC server resource.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * The system disk volume ID.
     */
    public /*out*/ readonly systemDiskId!: pulumi.Output<string>;
    /**
     * Specifies the size of system disk for the IEC server binding. The
     * value range is 40 to 100 in GB. Changing this parameter creates a new IEC server resource.
     */
    public readonly systemDiskSize!: pulumi.Output<number>;
    /**
     * Specifies the type of system disk for the IEC server binding. Valid
     * value is *SAS*(high I/O disk type). Changing this parameter creates a new IEC server resource.
     */
    public readonly systemDiskType!: pulumi.Output<string>;
    /**
     * Specifies the user data (information after encoding) configured during IEC
     * server creation. The value can come from a variety of sources: inline, read in from the *file* function. Changing this
     * parameter creates a new IEC server resource.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * An array of one or more disks to attach to the IEC server. The object structure is documented
     * below.
     */
    public /*out*/ readonly volumeAttacheds!: pulumi.Output<outputs.Iec.ServerVolumeAttached[]>;
    /**
     * Specifies the ID of vpc for the IEC server. VPC mode only *CUSTOMER* can be
     * used to create IEC server. Changing this parameter creates a new IEC server resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            resourceInputs["adminPass"] = state ? state.adminPass : undefined;
            resourceInputs["bindEip"] = state ? state.bindEip : undefined;
            resourceInputs["coverageLevel"] = state ? state.coverageLevel : undefined;
            resourceInputs["coveragePolicy"] = state ? state.coveragePolicy : undefined;
            resourceInputs["coverageSites"] = state ? state.coverageSites : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["edgecloudId"] = state ? state.edgecloudId : undefined;
            resourceInputs["edgecloudName"] = state ? state.edgecloudName : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["flavorName"] = state ? state.flavorName : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nics"] = state ? state.nics : undefined;
            resourceInputs["originServerId"] = state ? state.originServerId : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["systemDiskId"] = state ? state.systemDiskId : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["systemDiskType"] = state ? state.systemDiskType : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["volumeAttacheds"] = state ? state.volumeAttacheds : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if ((!args || args.coverageSites === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coverageSites'");
            }
            if ((!args || args.flavorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavorId'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.securityGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroups'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.systemDiskSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'systemDiskSize'");
            }
            if ((!args || args.systemDiskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'systemDiskType'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["adminPass"] = args ? args.adminPass : undefined;
            resourceInputs["bindEip"] = args ? args.bindEip : undefined;
            resourceInputs["coverageLevel"] = args ? args.coverageLevel : undefined;
            resourceInputs["coveragePolicy"] = args ? args.coveragePolicy : undefined;
            resourceInputs["coverageSites"] = args ? args.coverageSites : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["systemDiskType"] = args ? args.systemDiskType : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["edgecloudId"] = undefined /*out*/;
            resourceInputs["edgecloudName"] = undefined /*out*/;
            resourceInputs["flavorName"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["nics"] = undefined /*out*/;
            resourceInputs["originServerId"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemDiskId"] = undefined /*out*/;
            resourceInputs["volumeAttacheds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * Specifies the administrative password to assign to the IEC server. This
     * parameter can contain a maximum of 26 characters, which may consist of letters, digits and Special characters(~!?,.:
     * ;-_'"(){}[]/<>@#$%^&*+|\\=) and space. This parameter and `keyPair` are alternative. Changing this changes the root
     * password on the existing server.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies whether the IEC server is bound to EIP. Changing this parameter
     * creates a new IEC server resource.
     */
    bindEip?: pulumi.Input<boolean>;
    /**
     * Specifies the coverage level of IEC sites. Valid value is *SITE*.
     * Changing this parameter creates a new IEC server resource.
     */
    coverageLevel?: pulumi.Input<string>;
    /**
     * Specifies the policy of IEC sites. Valid values are *centralize*
     * and *discrete*, *centralize* is default. Changing this parameter creates a new IEC server resource.
     */
    coveragePolicy?: pulumi.Input<string>;
    /**
     * Specifies an array of site ID and operator for the IEC server. The
     * object structure is documented below. Changing this parameter creates a new IEC server resource.
     */
    coverageSites?: pulumi.Input<pulumi.Input<inputs.Iec.ServerCoverageSite>[]>;
    /**
     * Specifies the array of data disks to attach to the IEC server. Up to two
     * data disks can be specified. The object structure is documented below. Changing this parameter creates a new IEC
     * server resource.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Iec.ServerDataDisk>[]>;
    /**
     * The ID of the edgecloud service.
     */
    edgecloudId?: pulumi.Input<string>;
    /**
     * The Name of the edgecloud service.
     */
    edgecloudName?: pulumi.Input<string>;
    /**
     * Specifies the flavor ID of the desired flavor for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The flavor name of the IEC server.
     */
    flavorName?: pulumi.Input<string>;
    /**
     * Specifies the image ID of the desired image for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The image name of the IEC server.
     */
    imageName?: pulumi.Input<string>;
    /**
     * Specifies the name of a key pair to put on the IEC server. The key pair must
     * already be created and associated with the tenant's account. This parameter and `adminPass` are alternative. Changing
     * this parameter creates a new IEC server resource.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Specifies the IEC server name. This parameter can contain a maximum of 64
     * characters, which may consist of letters, digits, dot(.), underscores (_), and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the IEC server. The object structure is documented below.
     */
    nics?: pulumi.Input<pulumi.Input<inputs.Iec.ServerNic>[]>;
    /**
     * The ID of origin server.
     */
    originServerId?: pulumi.Input<string>;
    /**
     * The EIP address that is associted to the IEC server.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the IEC server. Changing this parameter creates a new IEC server resource.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of IEC server.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more subnet ID of Network for the IEC server
     * binding. Changing this parameter creates a new IEC server resource.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The system disk volume ID.
     */
    systemDiskId?: pulumi.Input<string>;
    /**
     * Specifies the size of system disk for the IEC server binding. The
     * value range is 40 to 100 in GB. Changing this parameter creates a new IEC server resource.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * Specifies the type of system disk for the IEC server binding. Valid
     * value is *SAS*(high I/O disk type). Changing this parameter creates a new IEC server resource.
     */
    systemDiskType?: pulumi.Input<string>;
    /**
     * Specifies the user data (information after encoding) configured during IEC
     * server creation. The value can come from a variety of sources: inline, read in from the *file* function. Changing this
     * parameter creates a new IEC server resource.
     */
    userData?: pulumi.Input<string>;
    /**
     * An array of one or more disks to attach to the IEC server. The object structure is documented
     * below.
     */
    volumeAttacheds?: pulumi.Input<pulumi.Input<inputs.Iec.ServerVolumeAttached>[]>;
    /**
     * Specifies the ID of vpc for the IEC server. VPC mode only *CUSTOMER* can be
     * used to create IEC server. Changing this parameter creates a new IEC server resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * Specifies the administrative password to assign to the IEC server. This
     * parameter can contain a maximum of 26 characters, which may consist of letters, digits and Special characters(~!?,.:
     * ;-_'"(){}[]/<>@#$%^&*+|\\=) and space. This parameter and `keyPair` are alternative. Changing this changes the root
     * password on the existing server.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies whether the IEC server is bound to EIP. Changing this parameter
     * creates a new IEC server resource.
     */
    bindEip?: pulumi.Input<boolean>;
    /**
     * Specifies the coverage level of IEC sites. Valid value is *SITE*.
     * Changing this parameter creates a new IEC server resource.
     */
    coverageLevel?: pulumi.Input<string>;
    /**
     * Specifies the policy of IEC sites. Valid values are *centralize*
     * and *discrete*, *centralize* is default. Changing this parameter creates a new IEC server resource.
     */
    coveragePolicy?: pulumi.Input<string>;
    /**
     * Specifies an array of site ID and operator for the IEC server. The
     * object structure is documented below. Changing this parameter creates a new IEC server resource.
     */
    coverageSites: pulumi.Input<pulumi.Input<inputs.Iec.ServerCoverageSite>[]>;
    /**
     * Specifies the array of data disks to attach to the IEC server. Up to two
     * data disks can be specified. The object structure is documented below. Changing this parameter creates a new IEC
     * server resource.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Iec.ServerDataDisk>[]>;
    /**
     * Specifies the flavor ID of the desired flavor for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    flavorId: pulumi.Input<string>;
    /**
     * Specifies the image ID of the desired image for the IEC server. Changing
     * this parameter creates a new IEC server resource.
     */
    imageId: pulumi.Input<string>;
    /**
     * Specifies the name of a key pair to put on the IEC server. The key pair must
     * already be created and associated with the tenant's account. This parameter and `adminPass` are alternative. Changing
     * this parameter creates a new IEC server resource.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Specifies the IEC server name. This parameter can contain a maximum of 64
     * characters, which may consist of letters, digits, dot(.), underscores (_), and hyphens (-).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the IEC server. Changing this parameter creates a new IEC server resource.
     */
    securityGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an array of one or more subnet ID of Network for the IEC server
     * binding. Changing this parameter creates a new IEC server resource.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the size of system disk for the IEC server binding. The
     * value range is 40 to 100 in GB. Changing this parameter creates a new IEC server resource.
     */
    systemDiskSize: pulumi.Input<number>;
    /**
     * Specifies the type of system disk for the IEC server binding. Valid
     * value is *SAS*(high I/O disk type). Changing this parameter creates a new IEC server resource.
     */
    systemDiskType: pulumi.Input<string>;
    /**
     * Specifies the user data (information after encoding) configured during IEC
     * server creation. The value can come from a variety of sources: inline, read in from the *file* function. Changing this
     * parameter creates a new IEC server resource.
     */
    userData?: pulumi.Input<string>;
    /**
     * Specifies the ID of vpc for the IEC server. VPC mode only *CUSTOMER* can be
     * used to create IEC server. Changing this parameter creates a new IEC server resource.
     */
    vpcId: pulumi.Input<string>;
}
