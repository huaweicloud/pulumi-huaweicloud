// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Dedicated Load Balancer resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Loadbalancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const basic = new huaweicloud.DedicatedElb.Loadbalancer("basic", {
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 *     crossVpcBackend: true,
 *     description: "basic example",
 *     enterpriseProjectId: "{{ eps_id }}",
 *     ipv4SubnetId: "{{ ipv4_subnet_id }}",
 *     l4FlavorId: "{{ l4_flavor_id }}",
 *     l7FlavorId: "{{ l7_flavor_id }}",
 *     vpcId: "{{ vpc_id }}",
 * });
 * ```
 * ### Loadbalancer With Existing EIP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const basic = new huaweicloud.DedicatedElb.Loadbalancer("basic", {
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 *     crossVpcBackend: true,
 *     description: "basic example",
 *     enterpriseProjectId: "{{ eps_id }}",
 *     ipv4EipId: "{{ eip_id }}",
 *     ipv4SubnetId: "{{ ipv4_subnet_id }}",
 *     ipv6BandwidthId: "{{ ipv6_bandwidth_id }}",
 *     ipv6NetworkId: "{{ ipv6_network_id }}",
 *     l4FlavorId: "{{ l4_flavor_id }}",
 *     l7FlavorId: "{{ l7_flavor_id }}",
 *     vpcId: "{{ vpc_id }}",
 * });
 * ```
 * ### Loadbalancer With EIP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const basic = new huaweicloud.DedicatedElb.Loadbalancer("basic", {
 *     availabilityZones: [
 *         "cn-north-4a",
 *         "cn-north-4b",
 *     ],
 *     bandwidthChargeMode: "traffic",
 *     bandwidthSize: 10,
 *     crossVpcBackend: true,
 *     description: "basic example",
 *     enterpriseProjectId: "{{ eps_id }}",
 *     iptype: "5_bgp",
 *     ipv4SubnetId: "{{ ipv4_subnet_id }}",
 *     ipv6BandwidthId: "{{ ipv6_bandwidth_id }}",
 *     ipv6NetworkId: "{{ ipv6_network_id }}",
 *     l4FlavorId: "{{ l4_flavor_id }}",
 *     l7FlavorId: "{{ l7_flavor_id }}",
 *     sharetype: "PER",
 *     vpcId: "{{ vpc_id }}",
 * });
 * ```
 *
 * ## Import
 *
 * ELB loadbalancer can be imported using the loadbalancer ID, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedElb/loadbalancer:Loadbalancer loadbalancer_1 5c20fdad-7288-11eb-b817-0255ac10158b
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`ipv6_bandwidth_id`, `iptype`, `bandwidth_charge_mode`, `sharetype` and `bandwidth_size`. It is generally recommended running `terraform plan` after importing a loadbalancer. You can then decide if changes should be applied to the loadbalancer, or the resource definition should be updated to align with the loadbalancer. Also you can ignore changes as below. resource "huaweicloud_elb_loadbalancer" "loadbalancer_1" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size,
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
     * Specifies whether auto renew is enabled. Valid values are **true** and **false**.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether autoscaling is enabled. Valid values are **true** and
     * **false**.
     */
    public readonly autoscalingEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the list of AZ names. Changing this parameter will create a
     * new resource.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * Bandwidth billing type. Changing this parameter will create a
     * new resource.
     */
    public readonly bandwidthChargeMode!: pulumi.Output<string>;
    /**
     * Bandwidth size. Changing this parameter will create a new resource.
     */
    public readonly bandwidthSize!: pulumi.Output<number>;
    /**
     * Specifies the charging mode of the ELB loadbalancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new resource.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    public readonly crossVpcBackend!: pulumi.Output<boolean>;
    /**
     * Human-readable description for the loadbalancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Elastic IP type. Changing this parameter will create a new resource.
     */
    public readonly iptype!: pulumi.Output<string>;
    /**
     * The ipv4 address of the load balancer.
     */
    public readonly ipv4Address!: pulumi.Output<string>;
    /**
     * The ipv4 eip address of the Load Balancer.
     */
    public /*out*/ readonly ipv4Eip!: pulumi.Output<string>;
    /**
     * The ID of the EIP. Changing this parameter will create a new resource.
     */
    public readonly ipv4EipId!: pulumi.Output<string>;
    /**
     * The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
     * ipv4 address.
     */
    public readonly ipv4SubnetId!: pulumi.Output<string | undefined>;
    /**
     * The ipv6 address of the Load Balancer.
     */
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    /**
     * The ipv6 bandwidth id. Only support shared bandwidth.
     */
    public readonly ipv6BandwidthId!: pulumi.Output<string | undefined>;
    /**
     * The ipv6 eip address of the Load Balancer.
     */
    public /*out*/ readonly ipv6Eip!: pulumi.Output<string>;
    /**
     * The ipv6 eip id of the Load Balancer.
     */
    public /*out*/ readonly ipv6EipId!: pulumi.Output<string>;
    /**
     * The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
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
     * Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
     * This parameter cannot be left blank if there are HTTP or HTTPS listeners.
     */
    public readonly minL7FlavorId!: pulumi.Output<string>;
    /**
     * Human-readable name for the loadbalancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the charging period of the ELB loadbalancer.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 3.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of the ELB loadbalancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Bandwidth sharing type. Changing this parameter will create a new resource.
     */
    public readonly sharetype!: pulumi.Output<string>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The vpc on which to create the loadbalancer. Changing this creates a new
     * loadbalancer.
     */
    public readonly vpcId!: pulumi.Output<string>;

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
            resourceInputs["bandwidthChargeMode"] = state ? state.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthSize"] = state ? state.bandwidthSize : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["crossVpcBackend"] = state ? state.crossVpcBackend : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["iptype"] = state ? state.iptype : undefined;
            resourceInputs["ipv4Address"] = state ? state.ipv4Address : undefined;
            resourceInputs["ipv4Eip"] = state ? state.ipv4Eip : undefined;
            resourceInputs["ipv4EipId"] = state ? state.ipv4EipId : undefined;
            resourceInputs["ipv4SubnetId"] = state ? state.ipv4SubnetId : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["ipv6BandwidthId"] = state ? state.ipv6BandwidthId : undefined;
            resourceInputs["ipv6Eip"] = state ? state.ipv6Eip : undefined;
            resourceInputs["ipv6EipId"] = state ? state.ipv6EipId : undefined;
            resourceInputs["ipv6NetworkId"] = state ? state.ipv6NetworkId : undefined;
            resourceInputs["l4FlavorId"] = state ? state.l4FlavorId : undefined;
            resourceInputs["l7FlavorId"] = state ? state.l7FlavorId : undefined;
            resourceInputs["minL7FlavorId"] = state ? state.minL7FlavorId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sharetype"] = state ? state.sharetype : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as LoadbalancerArgs | undefined;
            if ((!args || args.availabilityZones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZones'");
            }
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoscalingEnabled"] = args ? args.autoscalingEnabled : undefined;
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["bandwidthChargeMode"] = args ? args.bandwidthChargeMode : undefined;
            resourceInputs["bandwidthSize"] = args ? args.bandwidthSize : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["crossVpcBackend"] = args ? args.crossVpcBackend : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["iptype"] = args ? args.iptype : undefined;
            resourceInputs["ipv4Address"] = args ? args.ipv4Address : undefined;
            resourceInputs["ipv4EipId"] = args ? args.ipv4EipId : undefined;
            resourceInputs["ipv4SubnetId"] = args ? args.ipv4SubnetId : undefined;
            resourceInputs["ipv6BandwidthId"] = args ? args.ipv6BandwidthId : undefined;
            resourceInputs["ipv6NetworkId"] = args ? args.ipv6NetworkId : undefined;
            resourceInputs["l4FlavorId"] = args ? args.l4FlavorId : undefined;
            resourceInputs["l7FlavorId"] = args ? args.l7FlavorId : undefined;
            resourceInputs["minL7FlavorId"] = args ? args.minL7FlavorId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sharetype"] = args ? args.sharetype : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["ipv4Eip"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["ipv6Eip"] = undefined /*out*/;
            resourceInputs["ipv6EipId"] = undefined /*out*/;
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
     * Specifies whether auto renew is enabled. Valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies whether autoscaling is enabled. Valid values are **true** and
     * **false**.
     */
    autoscalingEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the list of AZ names. Changing this parameter will create a
     * new resource.
     */
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth billing type. Changing this parameter will create a
     * new resource.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth size. Changing this parameter will create a new resource.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the charging mode of the ELB loadbalancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new resource.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    crossVpcBackend?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Elastic IP type. Changing this parameter will create a new resource.
     */
    iptype?: pulumi.Input<string>;
    /**
     * The ipv4 address of the load balancer.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * The ipv4 eip address of the Load Balancer.
     */
    ipv4Eip?: pulumi.Input<string>;
    /**
     * The ID of the EIP. Changing this parameter will create a new resource.
     */
    ipv4EipId?: pulumi.Input<string>;
    /**
     * The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
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
     * The ipv6 eip address of the Load Balancer.
     */
    ipv6Eip?: pulumi.Input<string>;
    /**
     * The ipv6 eip id of the Load Balancer.
     */
    ipv6EipId?: pulumi.Input<string>;
    /**
     * The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
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
     * Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
     * This parameter cannot be left blank if there are HTTP or HTTPS listeners.
     */
    minL7FlavorId?: pulumi.Input<string>;
    /**
     * Human-readable name for the loadbalancer.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the ELB loadbalancer.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 3.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the ELB loadbalancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    region?: pulumi.Input<string>;
    /**
     * Bandwidth sharing type. Changing this parameter will create a new resource.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The vpc on which to create the loadbalancer. Changing this creates a new
     * loadbalancer.
     */
    vpcId?: pulumi.Input<string>;
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
     * Specifies whether auto renew is enabled. Valid values are **true** and **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies whether autoscaling is enabled. Valid values are **true** and
     * **false**.
     */
    autoscalingEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the list of AZ names. Changing this parameter will create a
     * new resource.
     */
    availabilityZones: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Bandwidth billing type. Changing this parameter will create a
     * new resource.
     */
    bandwidthChargeMode?: pulumi.Input<string>;
    /**
     * Bandwidth size. Changing this parameter will create a new resource.
     */
    bandwidthSize?: pulumi.Input<number>;
    /**
     * Specifies the charging mode of the ELB loadbalancer.
     * Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
     * Changing this parameter will create a new resource.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Enable this if you want to associate the IP addresses of backend servers with
     * your load balancer. Can only be true when updating.
     */
    crossVpcBackend?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project id of the loadbalancer. Changing this
     * creates a new loadbalancer.
     */
    enterpriseProjectId?: pulumi.Input<string>;
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
     * The **IPv4 subnet ID** of the subnet on which to allocate the loadbalancer's
     * ipv4 address.
     */
    ipv4SubnetId?: pulumi.Input<string>;
    /**
     * The ipv6 bandwidth id. Only support shared bandwidth.
     */
    ipv6BandwidthId?: pulumi.Input<string>;
    /**
     * The **ID** of the subnet on which to allocate the loadbalancer's ipv6 address.
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
     * Specifies the ID of the minimum Layer-7 flavor for elastic scaling.
     * This parameter cannot be left blank if there are HTTP or HTTPS listeners.
     */
    minL7FlavorId?: pulumi.Input<string>;
    /**
     * Human-readable name for the loadbalancer.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the charging period of the ELB loadbalancer.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 3.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of the ELB loadbalancer.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The region in which to create the loadbalancer resource. If omitted, the
     * provider-level region will be used. Changing this creates a new loadbalancer.
     */
    region?: pulumi.Input<string>;
    /**
     * Bandwidth sharing type. Changing this parameter will create a new resource.
     */
    sharetype?: pulumi.Input<string>;
    /**
     * The key/value pairs to associate with the loadbalancer.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The vpc on which to create the loadbalancer. Changing this creates a new
     * loadbalancer.
     */
    vpcId?: pulumi.Input<string>;
}