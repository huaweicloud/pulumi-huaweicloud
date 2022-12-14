// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an APIG dedicated instance resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceName = config.requireObject("instanceName");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const securityGroupId = config.requireObject("securityGroupId");
 * const eipId = config.requireObject("eipId");
 * const enterpriseProjectId = config.requireObject("enterpriseProjectId");
 * const testAvailabilityZones = huaweicloud.getAvailabilityZones({});
 * const testInstance = new huaweicloud.dedicatedapig.Instance("testInstance", {
 *     edition: "BASIC",
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     securityGroupId: securityGroupId,
 *     enterpriseProjectId: enterpriseProjectId,
 *     maintainBegin: "06:00:00",
 *     description: "Created by script",
 *     bandwidthSize: 3,
 *     eipId: eipId,
 *     availableZones: [
 *         testAvailabilityZones.then(testAvailabilityZones => testAvailabilityZones.names?[0]),
 *         testAvailabilityZones.then(testAvailabilityZones => testAvailabilityZones.names?[1]),
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * APIG Dedicated Instances can be imported by their `id`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/instance:Instance test de379eed30aa4d31a84f426ea3c7ef4e
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
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/instance:Instance';

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
     * Specifies an array of available zone names for the APIG dedicated
     * instance. Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?APIG) for list elements.
     * Changing this will create a new APIG dedicated instance resource.
     */
    public readonly availableZones!: pulumi.Output<string[]>;
    /**
     * Specifies the egress bandwidth size of the APIG dedicated instance. The range of
     * valid value is from 1 to 2000.
     */
    public readonly bandwidthSize!: pulumi.Output<number | undefined>;
    /**
     * Time when the APIG instance is created, in RFC-3339 format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specifies the description about the APIG dedicated instance. The description
     * contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the edition of the APIG dedicated instance. The supported editions
     * are as follows: BASIC, PROFESSIONAL, ENTERPRISE, PLATINUM. Changing this will create a new APIG dedicated instance
     * resource.
     */
    public readonly edition!: pulumi.Output<string>;
    /**
     * The egress (nat) public ip address.
     */
    public /*out*/ readonly egressAddress!: pulumi.Output<string>;
    /**
     * Specifies the eip ID associated with the APIG dedicated instance.
     */
    public readonly eipId!: pulumi.Output<string | undefined>;
    /**
     * Specifies an enterprise project ID. This parameter is required
     * for enterprise users. Changing this will create a new APIG dedicated instance resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The ingress eip address.
     */
    public /*out*/ readonly ingressAddress!: pulumi.Output<string>;
    /**
     * Specifies a start time of the maintenance time window in the format 'xx:00:00'.
     * The value of xx can be 02, 06, 10, 14, 18 or 22.
     */
    public readonly maintainBegin!: pulumi.Output<string>;
    /**
     * End time of the maintenance time window, 4-hour difference between the start time and end time.
     */
    public /*out*/ readonly maintainEnd!: pulumi.Output<string>;
    /**
     * Specifies the name of the API dedicated instance. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the APIG dedicated instance resource.
     * If omitted, the provider-level region will be used. Changing this will create a new APIG dedicated instance resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies an ID of the security group to which the APIG dedicated instance
     * belongs to.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Status of the APIG dedicated instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies an ID of the VPC subnet used to create the APIG dedicated
     * instance. Changing this will create a new APIG dedicated instance resource.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The supported features of the APIG dedicated instance.
     */
    public /*out*/ readonly supportedFeatures!: pulumi.Output<string[]>;
    /**
     * Specifies an ID of the VPC used to create the APIG dedicated instance.
     * Changing this will create a new APIG dedicated instance resource.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The ingress private ip address of vpc.
     */
    public /*out*/ readonly vpcIngressAddress!: pulumi.Output<string>;

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
            resourceInputs["availableZones"] = state ? state.availableZones : undefined;
            resourceInputs["bandwidthSize"] = state ? state.bandwidthSize : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edition"] = state ? state.edition : undefined;
            resourceInputs["egressAddress"] = state ? state.egressAddress : undefined;
            resourceInputs["eipId"] = state ? state.eipId : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["ingressAddress"] = state ? state.ingressAddress : undefined;
            resourceInputs["maintainBegin"] = state ? state.maintainBegin : undefined;
            resourceInputs["maintainEnd"] = state ? state.maintainEnd : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["supportedFeatures"] = state ? state.supportedFeatures : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcIngressAddress"] = state ? state.vpcIngressAddress : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.availableZones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availableZones'");
            }
            if ((!args || args.edition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edition'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availableZones"] = args ? args.availableZones : undefined;
            resourceInputs["bandwidthSize"] = args ? args.bandwidthSize : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edition"] = args ? args.edition : undefined;
            resourceInputs["eipId"] = args ? args.eipId : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["maintainBegin"] = args ? args.maintainBegin : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["egressAddress"] = undefined /*out*/;
            resourceInputs["ingressAddress"] = undefined /*out*/;
            resourceInputs["maintainEnd"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["supportedFeatures"] = undefined /*out*/;
            resourceInputs["vpcIngressAddress"] = undefined /*out*/;
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
     * Specifies an array of available zone names for the APIG dedicated
     * instance. Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?APIG) for list elements.
     * Changing this will create a new APIG dedicated instance resource.
     */
    availableZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the egress bandwidth size of the APIG dedicated instance. The range of
     * valid value is from 1 to 2000.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Time when the APIG instance is created, in RFC-3339 format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Specifies the description about the APIG dedicated instance. The description
     * contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the edition of the APIG dedicated instance. The supported editions
     * are as follows: BASIC, PROFESSIONAL, ENTERPRISE, PLATINUM. Changing this will create a new APIG dedicated instance
     * resource.
     */
    edition?: pulumi.Input<string>;
    /**
     * The egress (nat) public ip address.
     */
    egressAddress?: pulumi.Input<string>;
    /**
     * Specifies the eip ID associated with the APIG dedicated instance.
     */
    eipId?: pulumi.Input<string>;
    /**
     * Specifies an enterprise project ID. This parameter is required
     * for enterprise users. Changing this will create a new APIG dedicated instance resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The ingress eip address.
     */
    ingressAddress?: pulumi.Input<string>;
    /**
     * Specifies a start time of the maintenance time window in the format 'xx:00:00'.
     * The value of xx can be 02, 06, 10, 14, 18 or 22.
     */
    maintainBegin?: pulumi.Input<string>;
    /**
     * End time of the maintenance time window, 4-hour difference between the start time and end time.
     */
    maintainEnd?: pulumi.Input<string>;
    /**
     * Specifies the name of the API dedicated instance. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the APIG dedicated instance resource.
     * If omitted, the provider-level region will be used. Changing this will create a new APIG dedicated instance resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an ID of the security group to which the APIG dedicated instance
     * belongs to.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Status of the APIG dedicated instance.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies an ID of the VPC subnet used to create the APIG dedicated
     * instance. Changing this will create a new APIG dedicated instance resource.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The supported features of the APIG dedicated instance.
     */
    supportedFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an ID of the VPC used to create the APIG dedicated instance.
     * Changing this will create a new APIG dedicated instance resource.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ingress private ip address of vpc.
     */
    vpcIngressAddress?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Specifies an array of available zone names for the APIG dedicated
     * instance. Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?APIG) for list elements.
     * Changing this will create a new APIG dedicated instance resource.
     */
    availableZones: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the egress bandwidth size of the APIG dedicated instance. The range of
     * valid value is from 1 to 2000.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the description about the APIG dedicated instance. The description
     * contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the edition of the APIG dedicated instance. The supported editions
     * are as follows: BASIC, PROFESSIONAL, ENTERPRISE, PLATINUM. Changing this will create a new APIG dedicated instance
     * resource.
     */
    edition: pulumi.Input<string>;
    /**
     * Specifies the eip ID associated with the APIG dedicated instance.
     */
    eipId?: pulumi.Input<string>;
    /**
     * Specifies an enterprise project ID. This parameter is required
     * for enterprise users. Changing this will create a new APIG dedicated instance resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies a start time of the maintenance time window in the format 'xx:00:00'.
     * The value of xx can be 02, 06, 10, 14, 18 or 22.
     */
    maintainBegin?: pulumi.Input<string>;
    /**
     * Specifies the name of the API dedicated instance. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the APIG dedicated instance resource.
     * If omitted, the provider-level region will be used. Changing this will create a new APIG dedicated instance resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies an ID of the security group to which the APIG dedicated instance
     * belongs to.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Specifies an ID of the VPC subnet used to create the APIG dedicated
     * instance. Changing this will create a new APIG dedicated instance resource.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Specifies an ID of the VPC used to create the APIG dedicated instance.
     * Changing this will create a new APIG dedicated instance resource.
     */
    vpcId: pulumi.Input<string>;
}
