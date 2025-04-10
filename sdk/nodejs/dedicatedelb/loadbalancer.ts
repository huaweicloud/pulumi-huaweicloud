// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Dedicated load balancer resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Loadbalancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const l4FlavorId = config.requireObject("l4FlavorId");
 * const l7FlavorId = config.requireObject("l7FlavorId");
 * const epsId = config.requireObject("epsId");
 * const basic = new huaweicloud.dedicatedelb.Loadbalancer("basic", {
 *     description: "basic example",
 *     crossVpcBackend: true,
 *     vpcId: vpcId,
 *     ipv4SubnetId: ipv4SubnetId,
 *     l4FlavorId: l4FlavorId,
 *     l7FlavorId: l7FlavorId,
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 *     enterpriseProjectId: epsId,
 * });
 * ```
 * ### Loadbalancer With Existing EIP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const ipv6NetworkId = config.requireObject("ipv6NetworkId");
 * const ipv6BandwidthId = config.requireObject("ipv6BandwidthId");
 * const l4FlavorId = config.requireObject("l4FlavorId");
 * const l7FlavorId = config.requireObject("l7FlavorId");
 * const epsId = config.requireObject("epsId");
 * const eipId = config.requireObject("eipId");
 * const basic = new huaweicloud.dedicatedelb.Loadbalancer("basic", {
 *     description: "basic example",
 *     crossVpcBackend: true,
 *     vpcId: vpcId,
 *     ipv6NetworkId: ipv6NetworkId,
 *     ipv6BandwidthId: ipv6BandwidthId,
 *     ipv4SubnetId: ipv4SubnetId,
 *     l4FlavorId: l4FlavorId,
 *     l7FlavorId: l7FlavorId,
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 *     enterpriseProjectId: epsId,
 *     ipv4EipId: eipId,
 * });
 * ```
 * ### Loadbalancer With EIP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const ipv6NetworkId = config.requireObject("ipv6NetworkId");
 * const ipv6BandwidthId = config.requireObject("ipv6BandwidthId");
 * const l4FlavorId = config.requireObject("l4FlavorId");
 * const l7FlavorId = config.requireObject("l7FlavorId");
 * const basic = new huaweicloud.dedicatedelb.Loadbalancer("basic", {
 *     description: "basic example",
 *     crossVpcBackend: true,
 *     vpcId: vpcId,
 *     ipv6NetworkId: ipv6NetworkId,
 *     ipv6BandwidthId: ipv6BandwidthId,
 *     ipv4SubnetId: ipv4SubnetId,
 *     l4FlavorId: l4FlavorId,
 *     l7FlavorId: l7FlavorId,
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 * });
 * ```
 * ### Loadbalancer With gateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.requireObject("vpcId");
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const ipv6NetworkId = config.requireObject("ipv6NetworkId");
 * const basic = new huaweicloud.dedicatedelb.Loadbalancer("basic", {
 *     description: "basic example",
 *     loadbalancerType: "gateway",
 *     vpcId: vpcId,
 *     ipv4SubnetId: ipv4SubnetId,
 *     ipv6NetworkId: ipv6NetworkId,
 *     availabilityZones: ["cn-north-4a"],
 * });
 * ```
 *
 * ## Import
 *
 * ELB load balancer can be imported using the ID, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedElb/loadbalancer:Loadbalancer loadbalancer_1 <id>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`ipv6_bandwidth_id`, `iptype`, `bandwidth_charge_mode`, `sharetype`,
 *
 * `bandwidth_size`, `bandwidth_id`, `force_delete` and `deletion_protection_enable`. It is generally recommended running `terraform plan` after importing a load balancer. You can then decide if changes should be applied to the load balancer, or the resource definition should be updated to align with the load balancer. Also you can ignore changes as below. hcl resource "huaweicloud_elb_loadbalancer" "loadbalancer_1" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size, bandwidth_id, force_delete,
 *
 *  deletion_protection_enable,
 *
 *  ]
 *
 *  } }
 */
export class Loadbalancer extends pulumi.CustomResource {
    /**
     * Get an existing Loadbalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerState, opts?: pulumi.CustomResourceOptions): Loadbalancer {
        return new Loadbalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedElb/loadbalancer:Loadbalancer';

    /**
     * Returns true if the given object is an instance of Loadbalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Loadbalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Loadbalancer.__pulumiType;
    }

    /**
     * @deprecated Deprecated
     */
    public readonly autoPay!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether auto-renew is enabled. Valid values are **true** and **false**.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    public readonly autoscalingEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the list of AZ names.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * The IDs of subnets on the downstream plane.
     * + If this parameter is not specified, select subnets as follows:
     * - If IPv6 is enabled for a load balancer, the ID of subnet specified in `ipv6NetworkId` will be used.
     * - If IPv4 is enabled for a load balancer, the ID of subnet specified in `ipv4SubnetId` will be used.
     * - If only public network is available for a load balancer, the ID of any subnet in the VPC where the load balancer
     * resides will be used. Subnets with more IP addresses are preferred.
     * + If there is more than one subnet, the first subnet in the list will be used, and the subnets must be in the VPC
     * where the load balancer resides.
     */
    public readonly backendSubnets!: pulumi.Output<string[]>;
    /**
     * Bandwidth billing type. Value options:
     * + **bandwidth**: Billed by bandwidth.
     * + **traffic**: Billed by traffic.
     */
    public readonly bandwidthChargeMode!: pulumi.Output<string>;
    /**
     * Bandwidth ID of the shared bandwidth. It is mandatory when `sharetype`
     * is **WHOLE**. Changing this parameter will create a new resource.
     */
    public readonly bandwidthId!: pulumi.Output<string>;
    /**
     * Bandwidth size. It is mandatory when `iptype` is set and `bandwidthId`
     * is empty. Changing this parameter will create a new resource.
     */
    public readonly bandwidthSize!: pulumi.Output<number>;
    /**
     * Indicates the billing mode. The value can be one of the following:
     * + **flavor**: Billed by the specifications you will select.
     * + **lcu**: Billed by LCU usage.
     */
    public /*out*/ readonly chargeMode!: pulumi.Output<string>;
    /**
     * Specifies the charging mode of the ELB load balancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Indicates the time when the load balancer was created, in RFC3339 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    public readonly crossVpcBackend!: pulumi.Output<boolean>;
    /**
     * Specifies whether to enable deletion protection
     * for the load balancer. Value options:
     * + **true**: Enable deletion protection.
     * + **false**: Disable deletion protection.
     */
    public readonly deletionProtectionEnable!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable description for the load balancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The enterprise project id of the load balancer.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies whether to forcibly delete the load balancer, remove the load balancer,
     * listeners, unbind associated pools. Defaults to **false**.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the load balancer is a dedicated load balancer.
     * The value can be one of the following:
     * + **false**: The load balancer is a shared load balancer.
     * + **true**: The load balancer is a dedicated load balancer.
     */
    public /*out*/ readonly guaranteed!: pulumi.Output<boolean>;
    /**
     * The flavor ID of the gateway load balancer.
     */
    public /*out*/ readonly gwFlavorId!: pulumi.Output<string>;
    /**
     * Elastic IP type. Changing this parameter will create a new resource.
     */
    public readonly iptype!: pulumi.Output<string>;
    /**
     * The ipv4 address of the load balancer.
     */
    public readonly ipv4Address!: pulumi.Output<string>;
    /**
     * The ipv4 eip address of the load balancer.
     */
    public /*out*/ readonly ipv4Eip!: pulumi.Output<string>;
    /**
     * The ID of the EIP. Changing this parameter will create a new resource.
     */
    public readonly ipv4EipId!: pulumi.Output<string>;
    /**
     * The ID of the port bound to the private IPv4 address of the load balancer.
     */
    public /*out*/ readonly ipv4PortId!: pulumi.Output<string>;
    /**
     * The **IPv4 subnet ID** of the subnet on which to allocate the load balancer
     * ipv4 address.
     */
    public readonly ipv4SubnetId!: pulumi.Output<string | undefined>;
    /**
     * The ipv6 address of the Load Balancer.
     */
    public readonly ipv6Address!: pulumi.Output<string>;
    /**
     * The ipv6 bandwidth id. Only support shared bandwidth.
     */
    public readonly ipv6BandwidthId!: pulumi.Output<string | undefined>;
    /**
     * The ipv6 eip address of the load balancer.
     */
    public /*out*/ readonly ipv6Eip!: pulumi.Output<string>;
    /**
     * The ipv6 eip id of the load balancer.
     */
    public /*out*/ readonly ipv6EipId!: pulumi.Output<string>;
    /**
     * The **ID** of the subnet on which to allocate the load balancer ipv6 address.
     */
    public readonly ipv6NetworkId!: pulumi.Output<string | undefined>;
    /**
     * The L4 flavor id of the load balancer.
     */
    public readonly l4FlavorId!: pulumi.Output<string>;
    /**
     * The L7 flavor id of the load balancer.
     */
    public readonly l7FlavorId!: pulumi.Output<string>;
    /**
     * Specifies the type of the load balancer. Value options:
     * + **gateway**: indicates a gateway load balancer.
     * + Keep empty(default) indicates other types of load balancers.
     */
    public readonly loadbalancerType!: pulumi.Output<string>;
    public readonly minL7FlavorId!: pulumi.Output<string>;
    /**
     * Human-readable name for the load balancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the charging period of the ELB load balancer.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `9`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of the ELB load balancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The reason for update protection. Only valid when `protectionStatus` is
     * **consoleProtection**.
     */
    public readonly protectionReason!: pulumi.Output<string | undefined>;
    /**
     * The protection status for update. Value options:
     * + **nonProtection**: No protection.
     * + **consoleProtection**: Console modification protection.
     */
    public readonly protectionStatus!: pulumi.Output<string>;
    /**
     * The region in which to create the load balancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new load balancer.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Bandwidth sharing type. Value options:
     * + **PER**: Dedicated bandwidth.
     * + **WHOLE**: Shared bandwidth.
     */
    public readonly sharetype!: pulumi.Output<string>;
    /**
     * The key/value pairs to associate with the load balancer.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Indicates the time when the load balancer was updated, in RFC3339 format.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The vpc on which to create the load balancer. Changing this creates a new
     * load balancer.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Specifies traffic distributing policies when the WAF is faulty.
     * Value options:
     * + **discard**: Traffic will not be distributed.
     * + **forward**: Traffic will be distributed to the default backend servers.
     */
    public readonly wafFailureAction!: pulumi.Output<string>;

    /**
     * Create a Loadbalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadbalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadbalancerArgs | LoadbalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerState | undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoscalingEnabled"] = state ? state.autoscalingEnabled : undefined;
            resourceInputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            resourceInputs["backendSubnets"] = state ? state.backendSubnets : undefined;
            resourceInputs["bandwidthChargeMode"] = state ? state.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthId"] = state ? state.bandwidthId : undefined;
            resourceInputs["bandwidthSize"] = state ? state.bandwidthSize : undefined;
            resourceInputs["chargeMode"] = state ? state.chargeMode : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["crossVpcBackend"] = state ? state.crossVpcBackend : undefined;
            resourceInputs["deletionProtectionEnable"] = state ? state.deletionProtectionEnable : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["guaranteed"] = state ? state.guaranteed : undefined;
            resourceInputs["gwFlavorId"] = state ? state.gwFlavorId : undefined;
            resourceInputs["iptype"] = state ? state.iptype : undefined;
            resourceInputs["ipv4Address"] = state ? state.ipv4Address : undefined;
            resourceInputs["ipv4Eip"] = state ? state.ipv4Eip : undefined;
            resourceInputs["ipv4EipId"] = state ? state.ipv4EipId : undefined;
            resourceInputs["ipv4PortId"] = state ? state.ipv4PortId : undefined;
            resourceInputs["ipv4SubnetId"] = state ? state.ipv4SubnetId : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["ipv6BandwidthId"] = state ? state.ipv6BandwidthId : undefined;
            resourceInputs["ipv6Eip"] = state ? state.ipv6Eip : undefined;
            resourceInputs["ipv6EipId"] = state ? state.ipv6EipId : undefined;
            resourceInputs["ipv6NetworkId"] = state ? state.ipv6NetworkId : undefined;
            resourceInputs["l4FlavorId"] = state ? state.l4FlavorId : undefined;
            resourceInputs["l7FlavorId"] = state ? state.l7FlavorId : undefined;
            resourceInputs["loadbalancerType"] = state ? state.loadbalancerType : undefined;
            resourceInputs["minL7FlavorId"] = state ? state.minL7FlavorId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["protectionReason"] = state ? state.protectionReason : undefined;
            resourceInputs["protectionStatus"] = state ? state.protectionStatus : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sharetype"] = state ? state.sharetype : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["wafFailureAction"] = state ? state.wafFailureAction : undefined;
        } else {
            const args = argsOrState as LoadbalancerArgs | undefined;
            if ((!args || args.availabilityZones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZones'");
            }
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoscalingEnabled"] = args ? args.autoscalingEnabled : undefined;
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["backendSubnets"] = args ? args.backendSubnets : undefined;
            resourceInputs["bandwidthChargeMode"] = args ? args.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthId"] = args ? args.bandwidthId : undefined;
            resourceInputs["bandwidthSize"] = args ? args.bandwidthSize : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["crossVpcBackend"] = args ? args.crossVpcBackend : undefined;
            resourceInputs["deletionProtectionEnable"] = args ? args.deletionProtectionEnable : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["iptype"] = args ? args.iptype : undefined;
            resourceInputs["ipv4Address"] = args ? args.ipv4Address : undefined;
            resourceInputs["ipv4EipId"] = args ? args.ipv4EipId : undefined;
            resourceInputs["ipv4SubnetId"] = args ? args.ipv4SubnetId : undefined;
            resourceInputs["ipv6Address"] = args ? args.ipv6Address : undefined;
            resourceInputs["ipv6BandwidthId"] = args ? args.ipv6BandwidthId : undefined;
            resourceInputs["ipv6NetworkId"] = args ? args.ipv6NetworkId : undefined;
            resourceInputs["l4FlavorId"] = args ? args.l4FlavorId : undefined;
            resourceInputs["l7FlavorId"] = args ? args.l7FlavorId : undefined;
            resourceInputs["loadbalancerType"] = args ? args.loadbalancerType : undefined;
            resourceInputs["minL7FlavorId"] = args ? args.minL7FlavorId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["protectionReason"] = args ? args.protectionReason : undefined;
            resourceInputs["protectionStatus"] = args ? args.protectionStatus : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sharetype"] = args ? args.sharetype : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["wafFailureAction"] = args ? args.wafFailureAction : undefined;
            resourceInputs["chargeMode"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["guaranteed"] = undefined /*out*/;
            resourceInputs["gwFlavorId"] = undefined /*out*/;
            resourceInputs["ipv4Eip"] = undefined /*out*/;
            resourceInputs["ipv4PortId"] = undefined /*out*/;
            resourceInputs["ipv6Eip"] = undefined /*out*/;
            resourceInputs["ipv6EipId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Loadbalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Loadbalancer resources.
 */
export interface LoadbalancerState {
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    /**
     * Specifies whether auto-renew is enabled. Valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    autoscalingEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the list of AZ names.
     */
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of subnets on the downstream plane.
     * + If this parameter is not specified, select subnets as follows:
     * - If IPv6 is enabled for a load balancer, the ID of subnet specified in `ipv6NetworkId` will be used.
     * - If IPv4 is enabled for a load balancer, the ID of subnet specified in `ipv4SubnetId` will be used.
     * - If only public network is available for a load balancer, the ID of any subnet in the VPC where the load balancer
     * resides will be used. Subnets with more IP addresses are preferred.
     * + If there is more than one subnet, the first subnet in the list will be used, and the subnets must be in the VPC
     * where the load balancer resides.
     */
    backendSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth billing type. Value options:
     * + **bandwidth**: Billed by bandwidth.
     * + **traffic**: Billed by traffic.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth ID of the shared bandwidth. It is mandatory when `sharetype`
     * is **WHOLE**. Changing this parameter will create a new resource.
     */
    bandwidthId?: pulumi.Input<string>;
    /**
     * Bandwidth size. It is mandatory when `iptype` is set and `bandwidthId`
     * is empty. Changing this parameter will create a new resource.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Indicates the billing mode. The value can be one of the following:
     * + **flavor**: Billed by the specifications you will select.
     * + **lcu**: Billed by LCU usage.
     */
    chargeMode?: pulumi.Input<string>;
    /**
     * Specifies the charging mode of the ELB load balancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Indicates the time when the load balancer was created, in RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    crossVpcBackend?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable deletion protection
     * for the load balancer. Value options:
     * + **true**: Enable deletion protection.
     * + **false**: Disable deletion protection.
     */
    deletionProtectionEnable?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the load balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the load balancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the load balancer, remove the load balancer,
     * listeners, unbind associated pools. Defaults to **false**.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Indicates whether the load balancer is a dedicated load balancer.
     * The value can be one of the following:
     * + **false**: The load balancer is a shared load balancer.
     * + **true**: The load balancer is a dedicated load balancer.
     */
    guaranteed?: pulumi.Input<boolean>;
    /**
     * The flavor ID of the gateway load balancer.
     */
    gwFlavorId?: pulumi.Input<string>;
    /**
     * Elastic IP type. Changing this parameter will create a new resource.
     */
    iptype?: pulumi.Input<string>;
    /**
     * The ipv4 address of the load balancer.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * The ipv4 eip address of the load balancer.
     */
    ipv4Eip?: pulumi.Input<string>;
    /**
     * The ID of the EIP. Changing this parameter will create a new resource.
     */
    ipv4EipId?: pulumi.Input<string>;
    /**
     * The ID of the port bound to the private IPv4 address of the load balancer.
     */
    ipv4PortId?: pulumi.Input<string>;
    /**
     * The **IPv4 subnet ID** of the subnet on which to allocate the load balancer
     * ipv4 address.
     */
    ipv4SubnetId?: pulumi.Input<string>;
    /**
     * The ipv6 address of the Load Balancer.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * The ipv6 bandwidth id. Only support shared bandwidth.
     */
    ipv6BandwidthId?: pulumi.Input<string>;
    /**
     * The ipv6 eip address of the load balancer.
     */
    ipv6Eip?: pulumi.Input<string>;
    /**
     * The ipv6 eip id of the load balancer.
     */
    ipv6EipId?: pulumi.Input<string>;
    /**
     * The **ID** of the subnet on which to allocate the load balancer ipv6 address.
     */
    ipv6NetworkId?: pulumi.Input<string>;
    /**
     * The L4 flavor id of the load balancer.
     */
    l4FlavorId?: pulumi.Input<string>;
    /**
     * The L7 flavor id of the load balancer.
     */
    l7FlavorId?: pulumi.Input<string>;
    /**
     * Specifies the type of the load balancer. Value options:
     * + **gateway**: indicates a gateway load balancer.
     * + Keep empty(default) indicates other types of load balancers.
     */
    loadbalancerType?: pulumi.Input<string>;
    minL7FlavorId?: pulumi.Input<string>;
    /**
     * Human-readable name for the load balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the ELB load balancer.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `9`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the ELB load balancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The reason for update protection. Only valid when `protectionStatus` is
     * **consoleProtection**.
     */
    protectionReason?: pulumi.Input<string>;
    /**
     * The protection status for update. Value options:
     * + **nonProtection**: No protection.
     * + **consoleProtection**: Console modification protection.
     */
    protectionStatus?: pulumi.Input<string>;
    /**
     * The region in which to create the load balancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new load balancer.
     */
    region?: pulumi.Input<string>;
    /**
     * Bandwidth sharing type. Value options:
     * + **PER**: Dedicated bandwidth.
     * + **WHOLE**: Shared bandwidth.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the load balancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the time when the load balancer was updated, in RFC3339 format.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The vpc on which to create the load balancer. Changing this creates a new
     * load balancer.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Specifies traffic distributing policies when the WAF is faulty.
     * Value options:
     * + **discard**: Traffic will not be distributed.
     * + **forward**: Traffic will be distributed to the default backend servers.
     */
    wafFailureAction?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Loadbalancer resource.
 */
export interface LoadbalancerArgs {
    /**
     * @deprecated Deprecated
     */
    autoPay?: pulumi.Input<string>;
    /**
     * Specifies whether auto-renew is enabled. Valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    autoscalingEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the list of AZ names.
     */
    availabilityZones: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of subnets on the downstream plane.
     * + If this parameter is not specified, select subnets as follows:
     * - If IPv6 is enabled for a load balancer, the ID of subnet specified in `ipv6NetworkId` will be used.
     * - If IPv4 is enabled for a load balancer, the ID of subnet specified in `ipv4SubnetId` will be used.
     * - If only public network is available for a load balancer, the ID of any subnet in the VPC where the load balancer
     * resides will be used. Subnets with more IP addresses are preferred.
     * + If there is more than one subnet, the first subnet in the list will be used, and the subnets must be in the VPC
     * where the load balancer resides.
     */
    backendSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth billing type. Value options:
     * + **bandwidth**: Billed by bandwidth.
     * + **traffic**: Billed by traffic.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth ID of the shared bandwidth. It is mandatory when `sharetype`
     * is **WHOLE**. Changing this parameter will create a new resource.
     */
    bandwidthId?: pulumi.Input<string>;
    /**
     * Bandwidth size. It is mandatory when `iptype` is set and `bandwidthId`
     * is empty. Changing this parameter will create a new resource.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the charging mode of the ELB load balancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    crossVpcBackend?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable deletion protection
     * for the load balancer. Value options:
     * + **true**: Enable deletion protection.
     * + **false**: Disable deletion protection.
     */
    deletionProtectionEnable?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the load balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the load balancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the load balancer, remove the load balancer,
     * listeners, unbind associated pools. Defaults to **false**.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Elastic IP type. Changing this parameter will create a new resource.
     */
    iptype?: pulumi.Input<string>;
    /**
     * The ipv4 address of the load balancer.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * The ID of the EIP. Changing this parameter will create a new resource.
     */
    ipv4EipId?: pulumi.Input<string>;
    /**
     * The **IPv4 subnet ID** of the subnet on which to allocate the load balancer
     * ipv4 address.
     */
    ipv4SubnetId?: pulumi.Input<string>;
    /**
     * The ipv6 address of the Load Balancer.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * The ipv6 bandwidth id. Only support shared bandwidth.
     */
    ipv6BandwidthId?: pulumi.Input<string>;
    /**
     * The **ID** of the subnet on which to allocate the load balancer ipv6 address.
     */
    ipv6NetworkId?: pulumi.Input<string>;
    /**
     * The L4 flavor id of the load balancer.
     */
    l4FlavorId?: pulumi.Input<string>;
    /**
     * The L7 flavor id of the load balancer.
     */
    l7FlavorId?: pulumi.Input<string>;
    /**
     * Specifies the type of the load balancer. Value options:
     * + **gateway**: indicates a gateway load balancer.
     * + Keep empty(default) indicates other types of load balancers.
     */
    loadbalancerType?: pulumi.Input<string>;
    minL7FlavorId?: pulumi.Input<string>;
    /**
     * Human-readable name for the load balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the ELB load balancer.
     * If `periodUnit` is set to **month**, the value ranges from `1` to `9`.
     * If `periodUnit` is set to **year**, the value ranges from `1` to `3`.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the ELB load balancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The reason for update protection. Only valid when `protectionStatus` is
     * **consoleProtection**.
     */
    protectionReason?: pulumi.Input<string>;
    /**
     * The protection status for update. Value options:
     * + **nonProtection**: No protection.
     * + **consoleProtection**: Console modification protection.
     */
    protectionStatus?: pulumi.Input<string>;
    /**
     * The region in which to create the load balancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new load balancer.
     */
    region?: pulumi.Input<string>;
    /**
     * Bandwidth sharing type. Value options:
     * + **PER**: Dedicated bandwidth.
     * + **WHOLE**: Shared bandwidth.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the load balancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The vpc on which to create the load balancer. Changing this creates a new
     * load balancer.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Specifies traffic distributing policies when the WAF is faulty.
     * Value options:
     * + **discard**: Traffic will not be distributed.
     * + **forward**: Traffic will be distributed to the default backend servers.
     */
    wafFailureAction?: pulumi.Input<string>;
}
