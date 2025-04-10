// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a BMS instance resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceName = config.requireObject("instanceName");
 * const imageId = config.requireObject("imageId");
 * const flavorId = config.requireObject("flavorId");
 * const userId = config.requireObject("userId");
 * const keyPair = config.requireObject("keyPair");
 * const eipId = config.requireObject("eipId");
 * const enterpriseProjectId = config.requireObject("enterpriseProjectId");
 * const myaz = huaweicloud.getAvailabilityZones({});
 * const myvpc = huaweicloud.Vpc.getVpc({
 *     name: "vpc-default",
 * });
 * const mynet = huaweicloud.Vpc.getSubnet({
 *     name: "subnet-default",
 * });
 * const mysecgroup = huaweicloud.Vpc.getSecgroup({
 *     name: "default",
 * });
 * const test = new huaweicloud.bms.Instance("test", {
 *     imageId: imageId,
 *     flavorId: flavorId,
 *     userId: userId,
 *     securityGroups: [mysecgroup.then(mysecgroup => mysecgroup.id)],
 *     availabilityZone: myaz.then(myaz => myaz.names?[0]),
 *     vpcId: myvpc.then(myvpc => myvpc.id),
 *     eipId: huaweicloud_vpc_eip.myeip.id,
 *     chargingMode: "prePaid",
 *     periodUnit: "month",
 *     period: 1,
 *     keyPair: keyPair,
 *     enterpriseProjectId: enterpriseProjectId,
 *     systemDiskSize: 150,
 *     systemDiskType: "SSD",
 *     dataDisks: [{
 *         type: "SSD",
 *         size: 100,
 *     }],
 *     nics: [{
 *         subnetId: mynet.then(mynet => mynet.id),
 *         ipAddress: "192.168.0.123",
 *     }],
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 * });
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Bms/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Specifies the login password of the administrator for logging in to the
     * BMS using password authentication. Changing this creates a new instance. The password must meet the following
     * complexity requirements:
     * + Contains 8 to 26 characters.
     * + Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special
     * characters !@$%^-_=+[{}]:,./?
     * + Cannot contain the username or the username in reverse.
     */
    public readonly adminPass!: pulumi.Output<string | undefined>;
    /**
     * Specifies the IAM agency name which is created on IAM to provide
     * temporary credentials for BMS to access cloud services.
     */
    public readonly agencyName!: pulumi.Output<string>;
    /**
     * Specifies whether auto renew is enabled. Valid values are "true" and "
     * false", defaults to *false*.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * Specifies the availability zone in which to create the instance.
     * Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?BMS)
     * for the values. Changing this creates a new instance.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Bandwidth billing type. Available options are:
     * + `traffic`: billing mode is traffic.
     * + `bandwidth`: billing mode is bandwidth.
     */
    public readonly bandwidthChargeMode!: pulumi.Output<string | undefined>;
    /**
     * Bandwidth size. Changing this creates a new instance.
     */
    public readonly bandwidthSize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging mode of the instance. Valid value is *prePaid*.
     * Changing this creates a new instance.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more data disks to attach to the instance. The
     * dataDisks object structure is documented below. A maximum of 59 disks can be mounted. Changing this creates a new
     * instance.
     */
    public readonly dataDisks!: pulumi.Output<outputs.Bms.InstanceDataDisk[] | undefined>;
    /**
     * The description of the instance.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The ID of disks attached.
     */
    public /*out*/ readonly diskIds!: pulumi.Output<string[]>;
    /**
     * Elastic IP billing type. If the bandwidth billing mode is bandwidth,
     * both prePaid and postPaid are supported. If the bandwidth billing mode is traffic, only postPaid is supported.
     * Changing this creates a new instance. Available options are:
     * + `prePaid`: indicates the yearly/monthly billing mode.
     * + `postPaid`: indicates the pay-per-use billing mode.
     */
    public readonly eipChargeMode!: pulumi.Output<string | undefined>;
    /**
     * The ID of the EIP. Changing this creates a new instance.
     */
    public readonly eipId!: pulumi.Output<string | undefined>;
    /**
     * Specifies a unique id in UUID format of enterprise project.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies the flavor ID of the desired flavor for the instance. Changing
     * this creates a new instance.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The host ID of the instance.
     */
    public /*out*/ readonly hostId!: pulumi.Output<string>;
    /**
     * Specifies the image ID of the desired image for the instance. Changing this
     * creates a new instance.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The imageName of the instance.
     */
    public /*out*/ readonly imageName!: pulumi.Output<string>;
    /**
     * Elastic IP type. Changing this creates a new instance.
     * Available options are:
     * + `5Bgp`: dynamic BGP.
     * + `5Sbgp`: static BGP.
     */
    public readonly iptype!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of a key pair for logging in to the BMS using key pair
     * authentication. The key pair must already be created and associated with the tenant's account. The parameter is
     * required when using a Windows image to create a BMS. Changing this creates a new instance.
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * Specifies the user-defined metadata key-value pair.
     * + A metadata key contains of a maximum of 255 Unicode characters which can be letters, digits, hyphens (-),
     * underscores (_), colons (:), and point (.).
     * + A metadata value consists of a maximum of 255 Unicode characters.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies a unique name for the instance. The name consists of 1 to 63 characters,
     * including letters, digits, underscores (_), hyphens (-), and periods (.).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more networks to attach to the instance. The network
     * object structure is documented below.
     */
    public readonly nics!: pulumi.Output<outputs.Bms.InstanceNic[]>;
    /**
     * Specifies the charging period of the instance. If `periodUnit` is set to *month*
     * , the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value is 1. This parameter is mandatory
     * if `chargingMode` is set to *prePaid*. Changing this creates a new instance.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of the instance. Valid values are *
     * month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * instance.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The EIP address that is associated to the instance.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the instance. If omitted, the
     * provider-level region will be used. Changing this creates a new instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the instance. Changing this creates a new instance.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * Bandwidth sharing type. Changing this creates a new instance. Available
     * options are:
     * + `PER`: indicates dedicated bandwidth.
     * + `WHOLE`: indicates shared bandwidth.
     */
    public readonly sharetype!: pulumi.Output<string | undefined>;
    /**
     * The status of the instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the system disk size in GB. The value ranges from 40 to 1024.
     * The system disk size must be greater than or equal to the minimum system disk size of the image. Changing this creates
     * a new instance.
     */
    public readonly systemDiskSize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the system disk type of the instance. For details about
     * disk types, see
     * [Disk Types and Disk Performance](https://support.huaweicloud.com/intl/en-us/productdesc-evs/en-us_topic_0014580744.html)
     * . Changing this creates a new instance. Available options are:
     * + `SSD`: ultra-high I/O disk type.
     * + `GPSSD`: general purpose SSD disk type.
     * + `SAS`: high I/O disk type.
     */
    public readonly systemDiskType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the key/value pairs to associate with the instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the user data to be injected during the instance creation. Text
     * and text files can be injected. `userData` can come from a variety of sources: inline, read in from the *file*
     * function. The content of `userData` can be plaint text or encoded with base64. Changing this creates a new instance.
     */
    public readonly userData!: pulumi.Output<string>;
    /**
     * Specifies the user ID. You can obtain the user ID from My Credential on the
     * management console. Changing this creates a new instance.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Specifies id of vpc in which to create the instance. Changing this creates a
     * new instance.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["adminPass"] = state ? state.adminPass : undefined;
            resourceInputs["agencyName"] = state ? state.agencyName : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["bandwidthChargeMode"] = state ? state.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthSize"] = state ? state.bandwidthSize : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskIds"] = state ? state.diskIds : undefined;
            resourceInputs["eipChargeMode"] = state ? state.eipChargeMode : undefined;
            resourceInputs["eipId"] = state ? state.eipId : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["hostId"] = state ? state.hostId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["iptype"] = state ? state.iptype : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nics"] = state ? state.nics : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["sharetype"] = state ? state.sharetype : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["systemDiskType"] = state ? state.systemDiskType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.flavorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavorId'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.nics === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nics'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["adminPass"] = args ? args.adminPass : undefined;
            resourceInputs["agencyName"] = args ? args.agencyName : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["bandwidthChargeMode"] = args ? args.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthSize"] = args ? args.bandwidthSize : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["eipChargeMode"] = args ? args.eipChargeMode : undefined;
            resourceInputs["eipId"] = args ? args.eipId : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["iptype"] = args ? args.iptype : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nics"] = args ? args.nics : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["sharetype"] = args ? args.sharetype : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["systemDiskType"] = args ? args.systemDiskType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["diskIds"] = undefined /*out*/;
            resourceInputs["hostId"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Specifies the login password of the administrator for logging in to the
     * BMS using password authentication. Changing this creates a new instance. The password must meet the following
     * complexity requirements:
     * + Contains 8 to 26 characters.
     * + Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special
     * characters !@$%^-_=+[{}]:,./?
     * + Cannot contain the username or the username in reverse.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies the IAM agency name which is created on IAM to provide
     * temporary credentials for BMS to access cloud services.
     */
    agencyName?: pulumi.Input<string>;
    /**
     * Specifies whether auto renew is enabled. Valid values are "true" and "
     * false", defaults to *false*.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone in which to create the instance.
     * Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?BMS)
     * for the values. Changing this creates a new instance.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Bandwidth billing type. Available options are:
     * + `traffic`: billing mode is traffic.
     * + `bandwidth`: billing mode is bandwidth.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth size. Changing this creates a new instance.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the charging mode of the instance. Valid value is *prePaid*.
     * Changing this creates a new instance.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more data disks to attach to the instance. The
     * dataDisks object structure is documented below. A maximum of 59 disks can be mounted. Changing this creates a new
     * instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Bms.InstanceDataDisk>[]>;
    /**
     * The description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of disks attached.
     */
    diskIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Elastic IP billing type. If the bandwidth billing mode is bandwidth,
     * both prePaid and postPaid are supported. If the bandwidth billing mode is traffic, only postPaid is supported.
     * Changing this creates a new instance. Available options are:
     * + `prePaid`: indicates the yearly/monthly billing mode.
     * + `postPaid`: indicates the pay-per-use billing mode.
     */
    eipChargeMode?: pulumi.Input<string>;
    /**
     * The ID of the EIP. Changing this creates a new instance.
     */
    eipId?: pulumi.Input<string>;
    /**
     * Specifies a unique id in UUID format of enterprise project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the flavor ID of the desired flavor for the instance. Changing
     * this creates a new instance.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The host ID of the instance.
     */
    hostId?: pulumi.Input<string>;
    /**
     * Specifies the image ID of the desired image for the instance. Changing this
     * creates a new instance.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The imageName of the instance.
     */
    imageName?: pulumi.Input<string>;
    /**
     * Elastic IP type. Changing this creates a new instance.
     * Available options are:
     * + `5Bgp`: dynamic BGP.
     * + `5Sbgp`: static BGP.
     */
    iptype?: pulumi.Input<string>;
    /**
     * Specifies the name of a key pair for logging in to the BMS using key pair
     * authentication. The key pair must already be created and associated with the tenant's account. The parameter is
     * required when using a Windows image to create a BMS. Changing this creates a new instance.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Specifies the user-defined metadata key-value pair.
     * + A metadata key contains of a maximum of 255 Unicode characters which can be letters, digits, hyphens (-),
     * underscores (_), colons (:), and point (.).
     * + A metadata value consists of a maximum of 255 Unicode characters.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a unique name for the instance. The name consists of 1 to 63 characters,
     * including letters, digits, underscores (_), hyphens (-), and periods (.).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more networks to attach to the instance. The network
     * object structure is documented below.
     */
    nics?: pulumi.Input<pulumi.Input<inputs.Bms.InstanceNic>[]>;
    /**
     * Specifies the charging period of the instance. If `periodUnit` is set to *month*
     * , the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value is 1. This parameter is mandatory
     * if `chargingMode` is set to *prePaid*. Changing this creates a new instance.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the instance. Valid values are *
     * month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * instance.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The EIP address that is associated to the instance.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the instance. If omitted, the
     * provider-level region will be used. Changing this creates a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the instance. Changing this creates a new instance.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth sharing type. Changing this creates a new instance. Available
     * options are:
     * + `PER`: indicates dedicated bandwidth.
     * + `WHOLE`: indicates shared bandwidth.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * The status of the instance.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the system disk size in GB. The value ranges from 40 to 1024.
     * The system disk size must be greater than or equal to the minimum system disk size of the image. Changing this creates
     * a new instance.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * Specifies the system disk type of the instance. For details about
     * disk types, see
     * [Disk Types and Disk Performance](https://support.huaweicloud.com/intl/en-us/productdesc-evs/en-us_topic_0014580744.html)
     * . Changing this creates a new instance. Available options are:
     * + `SSD`: ultra-high I/O disk type.
     * + `GPSSD`: general purpose SSD disk type.
     * + `SAS`: high I/O disk type.
     */
    systemDiskType?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the user data to be injected during the instance creation. Text
     * and text files can be injected. `userData` can come from a variety of sources: inline, read in from the *file*
     * function. The content of `userData` can be plaint text or encoded with base64. Changing this creates a new instance.
     */
    userData?: pulumi.Input<string>;
    /**
     * Specifies the user ID. You can obtain the user ID from My Credential on the
     * management console. Changing this creates a new instance.
     */
    userId?: pulumi.Input<string>;
    /**
     * Specifies id of vpc in which to create the instance. Changing this creates a
     * new instance.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Specifies the login password of the administrator for logging in to the
     * BMS using password authentication. Changing this creates a new instance. The password must meet the following
     * complexity requirements:
     * + Contains 8 to 26 characters.
     * + Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special
     * characters !@$%^-_=+[{}]:,./?
     * + Cannot contain the username or the username in reverse.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies the IAM agency name which is created on IAM to provide
     * temporary credentials for BMS to access cloud services.
     */
    agencyName?: pulumi.Input<string>;
    /**
     * Specifies whether auto renew is enabled. Valid values are "true" and "
     * false", defaults to *false*.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone in which to create the instance.
     * Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?BMS)
     * for the values. Changing this creates a new instance.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Bandwidth billing type. Available options are:
     * + `traffic`: billing mode is traffic.
     * + `bandwidth`: billing mode is bandwidth.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth size. Changing this creates a new instance.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the charging mode of the instance. Valid value is *prePaid*.
     * Changing this creates a new instance.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more data disks to attach to the instance. The
     * dataDisks object structure is documented below. A maximum of 59 disks can be mounted. Changing this creates a new
     * instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Bms.InstanceDataDisk>[]>;
    /**
     * Elastic IP billing type. If the bandwidth billing mode is bandwidth,
     * both prePaid and postPaid are supported. If the bandwidth billing mode is traffic, only postPaid is supported.
     * Changing this creates a new instance. Available options are:
     * + `prePaid`: indicates the yearly/monthly billing mode.
     * + `postPaid`: indicates the pay-per-use billing mode.
     */
    eipChargeMode?: pulumi.Input<string>;
    /**
     * The ID of the EIP. Changing this creates a new instance.
     */
    eipId?: pulumi.Input<string>;
    /**
     * Specifies a unique id in UUID format of enterprise project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the flavor ID of the desired flavor for the instance. Changing
     * this creates a new instance.
     */
    flavorId: pulumi.Input<string>;
    /**
     * Specifies the image ID of the desired image for the instance. Changing this
     * creates a new instance.
     */
    imageId: pulumi.Input<string>;
    /**
     * Elastic IP type. Changing this creates a new instance.
     * Available options are:
     * + `5Bgp`: dynamic BGP.
     * + `5Sbgp`: static BGP.
     */
    iptype?: pulumi.Input<string>;
    /**
     * Specifies the name of a key pair for logging in to the BMS using key pair
     * authentication. The key pair must already be created and associated with the tenant's account. The parameter is
     * required when using a Windows image to create a BMS. Changing this creates a new instance.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Specifies the user-defined metadata key-value pair.
     * + A metadata key contains of a maximum of 255 Unicode characters which can be letters, digits, hyphens (-),
     * underscores (_), colons (:), and point (.).
     * + A metadata value consists of a maximum of 255 Unicode characters.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a unique name for the instance. The name consists of 1 to 63 characters,
     * including letters, digits, underscores (_), hyphens (-), and periods (.).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more networks to attach to the instance. The network
     * object structure is documented below.
     */
    nics: pulumi.Input<pulumi.Input<inputs.Bms.InstanceNic>[]>;
    /**
     * Specifies the charging period of the instance. If `periodUnit` is set to *month*
     * , the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value is 1. This parameter is mandatory
     * if `chargingMode` is set to *prePaid*. Changing this creates a new instance.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the instance. Valid values are *
     * month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
     * instance.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the instance. If omitted, the
     * provider-level region will be used. Changing this creates a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with
     * the instance. Changing this creates a new instance.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth sharing type. Changing this creates a new instance. Available
     * options are:
     * + `PER`: indicates dedicated bandwidth.
     * + `WHOLE`: indicates shared bandwidth.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * Specifies the system disk size in GB. The value ranges from 40 to 1024.
     * The system disk size must be greater than or equal to the minimum system disk size of the image. Changing this creates
     * a new instance.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * Specifies the system disk type of the instance. For details about
     * disk types, see
     * [Disk Types and Disk Performance](https://support.huaweicloud.com/intl/en-us/productdesc-evs/en-us_topic_0014580744.html)
     * . Changing this creates a new instance. Available options are:
     * + `SSD`: ultra-high I/O disk type.
     * + `GPSSD`: general purpose SSD disk type.
     * + `SAS`: high I/O disk type.
     */
    systemDiskType?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the user data to be injected during the instance creation. Text
     * and text files can be injected. `userData` can come from a variety of sources: inline, read in from the *file*
     * function. The content of `userData` can be plaint text or encoded with base64. Changing this creates a new instance.
     */
    userData?: pulumi.Input<string>;
    /**
     * Specifies the user ID. You can obtain the user ID from My Credential on the
     * management console. Changing this creates a new instance.
     */
    userId: pulumi.Input<string>;
    /**
     * Specifies id of vpc in which to create the instance. Changing this creates a
     * new instance.
     */
    vpcId: pulumi.Input<string>;
}
