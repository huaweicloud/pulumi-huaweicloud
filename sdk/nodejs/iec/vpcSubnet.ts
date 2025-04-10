// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a VPC subnet resource within HuaweiCloud IEC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const sitesTest = huaweicloud.Iec.getSites({});
 * const vpcTest = new huaweicloud.iec.Vpc("vpcTest", {
 *     cidr: "192.168.0.0/16",
 *     mode: "CUSTOMER",
 * });
 * const subnetTest = new huaweicloud.iec.VpcSubnet("subnetTest", {
 *     cidr: "192.168.128.0/18",
 *     vpcId: vpcTest.id,
 *     siteId: sitesTest.then(sitesTest => sitesTest.sites?[0]?.id),
 *     gatewayIp: "192.168.128.1",
 * });
 * ```
 *
 * ## Import
 *
 * IEC vpc subnet can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Iec/vpcSubnet:VpcSubnet subnet_demo 51be9f2b-5a3b-406a-9271-36f0c929fbcc
 * ```
 */
export class VpcSubnet extends pulumi.CustomResource {
    /**
     * Get an existing VpcSubnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcSubnetState, opts?: pulumi.CustomResourceOptions): VpcSubnet {
        return new VpcSubnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Iec/vpcSubnet:VpcSubnet';

    /**
     * Returns true if the given object is an instance of VpcSubnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcSubnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcSubnet.__pulumiType;
    }

    /**
     * Specifies the network segment on which the subnet resides. The value must be in
     * CIDR format and within the CIDR block of the iec vpc. Changing this parameter creates a new subnet resource.
     */
    public readonly cidr!: pulumi.Output<string>;
    /**
     * Specifies the status of subnet DHCP is enabled or not.
     * Valid values are **true** and **false**, defaults to **true**.
     */
    public readonly dhcpEnable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the DNS server address list of a subnet. These DNS server address must be
     * valid IP addresses.
     */
    public readonly dnsLists!: pulumi.Output<string[]>;
    /**
     * Specifies the gateway of the subnet. The value must be a valid IP address
     * and in the subnet segment. Changing this parameter creates a new subnet resource.
     */
    public readonly gatewayIp!: pulumi.Output<string>;
    /**
     * Specifies the name of the iec vpc subnet. The value is a string of 1 to 64 characters that
     * can contain letters, digits, underscores(_), and hyphens(-).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the iec vpc subnet resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the iec site. Changing this parameter creates a new
     * subnet resource.
     */
    public readonly siteId!: pulumi.Output<string>;
    /**
     * The located information of the iec site. It contains area, province and city.
     */
    public /*out*/ readonly siteInfo!: pulumi.Output<string>;
    /**
     * The status of the subnet.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the ID of the iec **CUSTOMER**
     * vpc to which the subnet belongs. Changing this parameter creates a new subnet resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcSubnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcSubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcSubnetArgs | VpcSubnetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcSubnetState | undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["dhcpEnable"] = state ? state.dhcpEnable : undefined;
            resourceInputs["dnsLists"] = state ? state.dnsLists : undefined;
            resourceInputs["gatewayIp"] = state ? state.gatewayIp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["siteInfo"] = state ? state.siteInfo : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcSubnetArgs | undefined;
            if ((!args || args.cidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidr'");
            }
            if ((!args || args.gatewayIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayIp'");
            }
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["dhcpEnable"] = args ? args.dhcpEnable : undefined;
            resourceInputs["dnsLists"] = args ? args.dnsLists : undefined;
            resourceInputs["gatewayIp"] = args ? args.gatewayIp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["siteInfo"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcSubnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcSubnet resources.
 */
export interface VpcSubnetState {
    /**
     * Specifies the network segment on which the subnet resides. The value must be in
     * CIDR format and within the CIDR block of the iec vpc. Changing this parameter creates a new subnet resource.
     */
    cidr?: pulumi.Input<string>;
    /**
     * Specifies the status of subnet DHCP is enabled or not.
     * Valid values are **true** and **false**, defaults to **true**.
     */
    dhcpEnable?: pulumi.Input<boolean>;
    /**
     * Specifies the DNS server address list of a subnet. These DNS server address must be
     * valid IP addresses.
     */
    dnsLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the gateway of the subnet. The value must be a valid IP address
     * and in the subnet segment. Changing this parameter creates a new subnet resource.
     */
    gatewayIp?: pulumi.Input<string>;
    /**
     * Specifies the name of the iec vpc subnet. The value is a string of 1 to 64 characters that
     * can contain letters, digits, underscores(_), and hyphens(-).
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the iec vpc subnet resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of the iec site. Changing this parameter creates a new
     * subnet resource.
     */
    siteId?: pulumi.Input<string>;
    /**
     * The located information of the iec site. It contains area, province and city.
     */
    siteInfo?: pulumi.Input<string>;
    /**
     * The status of the subnet.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the ID of the iec **CUSTOMER**
     * vpc to which the subnet belongs. Changing this parameter creates a new subnet resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcSubnet resource.
 */
export interface VpcSubnetArgs {
    /**
     * Specifies the network segment on which the subnet resides. The value must be in
     * CIDR format and within the CIDR block of the iec vpc. Changing this parameter creates a new subnet resource.
     */
    cidr: pulumi.Input<string>;
    /**
     * Specifies the status of subnet DHCP is enabled or not.
     * Valid values are **true** and **false**, defaults to **true**.
     */
    dhcpEnable?: pulumi.Input<boolean>;
    /**
     * Specifies the DNS server address list of a subnet. These DNS server address must be
     * valid IP addresses.
     */
    dnsLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the gateway of the subnet. The value must be a valid IP address
     * and in the subnet segment. Changing this parameter creates a new subnet resource.
     */
    gatewayIp: pulumi.Input<string>;
    /**
     * Specifies the name of the iec vpc subnet. The value is a string of 1 to 64 characters that
     * can contain letters, digits, underscores(_), and hyphens(-).
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the iec vpc subnet resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the ID of the iec site. Changing this parameter creates a new
     * subnet resource.
     */
    siteId: pulumi.Input<string>;
    /**
     * Specifies the ID of the iec **CUSTOMER**
     * vpc to which the subnet belongs. Changing this parameter creates a new subnet resource.
     */
    vpcId: pulumi.Input<string>;
}
