// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an AS group resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Basic Autoscaling Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const configurationId = config.requireObject("configurationId");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const myAsGroup = new huaweicloud.as.Group("myAsGroup", {
 *     scalingGroupName: "my_as_group",
 *     scalingConfigurationId: configurationId,
 *     desireInstanceNumber: 2,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 10,
 *     vpcId: vpcId,
 *     deletePublicip: true,
 *     deleteInstances: "yes",
 *     networks: [{
 *         id: subnetId,
 *     }],
 * });
 * ```
 * ### Autoscaling Group with tags
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const configurationId = config.requireObject("configurationId");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const myAsGroupTags = new huaweicloud.as.Group("myAsGroupTags", {
 *     scalingGroupName: "my_as_group_tags",
 *     scalingConfigurationId: configurationId,
 *     desireInstanceNumber: 2,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 10,
 *     vpcId: vpcId,
 *     deletePublicip: true,
 *     deleteInstances: "yes",
 *     networks: [{
 *         id: subnetId,
 *     }],
 *     tags: {
 *         foo: "bar",
 *         key: "value",
 *     },
 * });
 * ```
 * ### Autoscaling Group Only Remove Members When Scaling Down
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const configurationId = config.requireObject("configurationId");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const myAsGroupOnlyRemoveMembers = new huaweicloud.as.Group("myAsGroupOnlyRemoveMembers", {
 *     scalingGroupName: "my_as_group_only_remove_members",
 *     scalingConfigurationId: configurationId,
 *     desireInstanceNumber: 2,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 10,
 *     vpcId: vpcId,
 *     deletePublicip: true,
 *     deleteInstances: "no",
 *     networks: [{
 *         id: subnetId,
 *     }],
 * });
 * ```
 * ### Autoscaling Group With Elastic Load Balancer Listener
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const configurationId = config.requireObject("configurationId");
 * const vpcId = config.requireObject("vpcId");
 * const subnetId = config.requireObject("subnetId");
 * const ipv4SubnetId = config.requireObject("ipv4SubnetId");
 * const loadbalancer1 = new huaweicloud.elb.Loadbalancer("loadbalancer1", {vipSubnetId: ipv4SubnetId});
 * const listener1 = new huaweicloud.elb.Listener("listener1", {
 *     protocol: "HTTP",
 *     protocolPort: 8080,
 *     loadbalancerId: loadbalancer1.id,
 * });
 * const pool1 = new huaweicloud.elb.Pool("pool1", {
 *     protocol: "HTTP",
 *     lbMethod: "ROUND_ROBIN",
 *     listenerId: listener1.id,
 * });
 * const myAsGroupWithEnhancedLb = new huaweicloud.as.Group("myAsGroupWithEnhancedLb", {
 *     scalingGroupName: "my_as_group_with_enhanced_lb",
 *     scalingConfigurationId: configurationId,
 *     desireInstanceNumber: 2,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 10,
 *     vpcId: vpcId,
 *     networks: [{
 *         id: subnetId,
 *     }],
 *     lbaasListeners: [{
 *         poolId: pool1.id,
 *         protocolPort: listener1.protocolPort,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * AS groups can be imported by their `id`. For example,
 *
 * ```sh
 *  $ pulumi import huaweicloud:As/group:Group my_as_group 9ec5bea6-a728-4082-8109-5a7dc5c7af74
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:As/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * Specifies the IAM agency name. If you change the agency,
     * the new agency will be available for ECSs scaled out after the change.
     */
    public readonly agencyName!: pulumi.Output<string>;
    /**
     * Specifies the availability zones in which to create the instances in the
     * autoscaling group.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * schema: Deprecated; use availability_zones instead
     */
    public readonly availableZones!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the cooling duration (in seconds). The value ranges from 0 to 86400,
     * and is 300 by default.
     */
    public readonly coolDownTime!: pulumi.Output<number | undefined>;
    /**
     * The number of current instances in the AS group.
     */
    public /*out*/ readonly currentInstanceNumber!: pulumi.Output<number>;
    /**
     * Specifies whether to delete the instances in the AS group when deleting
     * the AS group. The options are `yes` and `no`.
     */
    public readonly deleteInstances!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to delete the elastic IP address bound to the instances of
     * AS group when deleting the instances. The options are `true` and `false`.
     */
    public readonly deletePublicip!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the description of the AS group.
     * The value can contain 0 to 256 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies the expected number of instances. The default value is the
     * minimum number of instances. The value ranges from the minimum number of instances to the maximum number of instances.
     */
    public readonly desireInstanceNumber!: pulumi.Output<number>;
    /**
     * Specifies whether to enable the AS Group. The options are `true` and `false`.
     * The default value is `true`.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the enterprise project id of the AS group.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Specifies whether to forcibly delete the AS group, remove the ECS instances and
     * release them. The default value is `false`.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the health check grace period for instances.
     * The unit is second and the value ranges from 0 to 86400. The default value is 600.
     */
    public readonly healthPeriodicAuditGracePeriod!: pulumi.Output<number>;
    /**
     * Specifies the health check method for instances in the AS group.
     * The health check methods include `ELB_AUDIT` and `NOVA_AUDIT`. If load balancing is configured, the default value of
     * this parameter is `ELB_AUDIT`. Otherwise, the default value is `NOVA_AUDIT`.
     */
    public readonly healthPeriodicAuditMethod!: pulumi.Output<string | undefined>;
    /**
     * Specifies the health check period for instances. The unit is minute
     * and value includes 0, 1, 5 (default), 15, 60, and 180. If the value is set to 0, health check is performed every 10 seconds.
     */
    public readonly healthPeriodicAuditTime!: pulumi.Output<number | undefined>;
    /**
     * Specifies the instance removal policy. The policy has four
     * options: `OLD_CONFIG_OLD_INSTANCE` (default), `OLD_CONFIG_NEW_INSTANCE`, `OLD_INSTANCE`, and `NEW_INSTANCE`.
     */
    public readonly instanceTerminatePolicy!: pulumi.Output<string | undefined>;
    /**
     * The instances IDs of the AS group.
     */
    public /*out*/ readonly instances!: pulumi.Output<string[]>;
    /**
     * The system supports the binding of up to six ELB listeners, the IDs of which are separated using a comma.
     *
     * @deprecated use lbaas_listeners instead
     */
    public readonly lbListenerId!: pulumi.Output<string | undefined>;
    /**
     * Specifies an array of one or more enhanced load balancer. The system supports
     * the binding of up to six load balancers. The object structure is documented below.
     */
    public readonly lbaasListeners!: pulumi.Output<outputs.As.GroupLbaasListener[]>;
    /**
     * Specifies the maximum number of instances. The default value is 0.
     */
    public readonly maxInstanceNumber!: pulumi.Output<number | undefined>;
    /**
     * Specifies the minimum number of instances. The default value is 0.
     */
    public readonly minInstanceNumber!: pulumi.Output<number | undefined>;
    /**
     * Specifies the priority policy used to select target AZs when adjusting
     * the number of instances in an AS group. The value can be `EQUILIBRIUM_DISTRIBUTE` and `PICK_FIRST`.
     */
    public readonly multiAzScalingPolicy!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more network IDs. The system supports up to five networks.
     * The object structure is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.As.GroupNetwork[]>;
    /**
     * schema: Deprecated; The notification mode has been canceled
     */
    public readonly notifications!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the region in which to create the AS group.
     * If omitted, the provider-level region will be used. Changing this creates a new group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the configuration ID which defines configurations
     * of instances in the AS group.
     */
    public readonly scalingConfigurationId!: pulumi.Output<string>;
    /**
     * Specifies the name of the scaling group. The name can contain
     * letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.
     */
    public readonly scalingGroupName!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with the group.
     * The object structure is documented below.
     */
    public readonly securityGroups!: pulumi.Output<outputs.As.GroupSecurityGroup[]>;
    /**
     * The status of the AS group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the AS group.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the VPC ID. Changing this creates a new group.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["agencyName"] = state ? state.agencyName : undefined;
            resourceInputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            resourceInputs["availableZones"] = state ? state.availableZones : undefined;
            resourceInputs["coolDownTime"] = state ? state.coolDownTime : undefined;
            resourceInputs["currentInstanceNumber"] = state ? state.currentInstanceNumber : undefined;
            resourceInputs["deleteInstances"] = state ? state.deleteInstances : undefined;
            resourceInputs["deletePublicip"] = state ? state.deletePublicip : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desireInstanceNumber"] = state ? state.desireInstanceNumber : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["healthPeriodicAuditGracePeriod"] = state ? state.healthPeriodicAuditGracePeriod : undefined;
            resourceInputs["healthPeriodicAuditMethod"] = state ? state.healthPeriodicAuditMethod : undefined;
            resourceInputs["healthPeriodicAuditTime"] = state ? state.healthPeriodicAuditTime : undefined;
            resourceInputs["instanceTerminatePolicy"] = state ? state.instanceTerminatePolicy : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
            resourceInputs["lbListenerId"] = state ? state.lbListenerId : undefined;
            resourceInputs["lbaasListeners"] = state ? state.lbaasListeners : undefined;
            resourceInputs["maxInstanceNumber"] = state ? state.maxInstanceNumber : undefined;
            resourceInputs["minInstanceNumber"] = state ? state.minInstanceNumber : undefined;
            resourceInputs["multiAzScalingPolicy"] = state ? state.multiAzScalingPolicy : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["notifications"] = state ? state.notifications : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scalingConfigurationId"] = state ? state.scalingConfigurationId : undefined;
            resourceInputs["scalingGroupName"] = state ? state.scalingGroupName : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            if ((!args || args.scalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupName'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["agencyName"] = args ? args.agencyName : undefined;
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["availableZones"] = args ? args.availableZones : undefined;
            resourceInputs["coolDownTime"] = args ? args.coolDownTime : undefined;
            resourceInputs["deleteInstances"] = args ? args.deleteInstances : undefined;
            resourceInputs["deletePublicip"] = args ? args.deletePublicip : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desireInstanceNumber"] = args ? args.desireInstanceNumber : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["healthPeriodicAuditGracePeriod"] = args ? args.healthPeriodicAuditGracePeriod : undefined;
            resourceInputs["healthPeriodicAuditMethod"] = args ? args.healthPeriodicAuditMethod : undefined;
            resourceInputs["healthPeriodicAuditTime"] = args ? args.healthPeriodicAuditTime : undefined;
            resourceInputs["instanceTerminatePolicy"] = args ? args.instanceTerminatePolicy : undefined;
            resourceInputs["lbListenerId"] = args ? args.lbListenerId : undefined;
            resourceInputs["lbaasListeners"] = args ? args.lbaasListeners : undefined;
            resourceInputs["maxInstanceNumber"] = args ? args.maxInstanceNumber : undefined;
            resourceInputs["minInstanceNumber"] = args ? args.minInstanceNumber : undefined;
            resourceInputs["multiAzScalingPolicy"] = args ? args.multiAzScalingPolicy : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["notifications"] = args ? args.notifications : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scalingConfigurationId"] = args ? args.scalingConfigurationId : undefined;
            resourceInputs["scalingGroupName"] = args ? args.scalingGroupName : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["currentInstanceNumber"] = undefined /*out*/;
            resourceInputs["instances"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * Specifies the IAM agency name. If you change the agency,
     * the new agency will be available for ECSs scaled out after the change.
     */
    agencyName?: pulumi.Input<string>;
    /**
     * Specifies the availability zones in which to create the instances in the
     * autoscaling group.
     */
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * schema: Deprecated; use availability_zones instead
     */
    availableZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the cooling duration (in seconds). The value ranges from 0 to 86400,
     * and is 300 by default.
     */
    coolDownTime?: pulumi.Input<number>;
    /**
     * The number of current instances in the AS group.
     */
    currentInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies whether to delete the instances in the AS group when deleting
     * the AS group. The options are `yes` and `no`.
     */
    deleteInstances?: pulumi.Input<string>;
    /**
     * Specifies whether to delete the elastic IP address bound to the instances of
     * AS group when deleting the instances. The options are `true` and `false`.
     */
    deletePublicip?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the AS group.
     * The value can contain 0 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the expected number of instances. The default value is the
     * minimum number of instances. The value ranges from the minimum number of instances to the maximum number of instances.
     */
    desireInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies whether to enable the AS Group. The options are `true` and `false`.
     * The default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Specifies the enterprise project id of the AS group.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the AS group, remove the ECS instances and
     * release them. The default value is `false`.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Specifies the health check grace period for instances.
     * The unit is second and the value ranges from 0 to 86400. The default value is 600.
     */
    healthPeriodicAuditGracePeriod?: pulumi.Input<number>;
    /**
     * Specifies the health check method for instances in the AS group.
     * The health check methods include `ELB_AUDIT` and `NOVA_AUDIT`. If load balancing is configured, the default value of
     * this parameter is `ELB_AUDIT`. Otherwise, the default value is `NOVA_AUDIT`.
     */
    healthPeriodicAuditMethod?: pulumi.Input<string>;
    /**
     * Specifies the health check period for instances. The unit is minute
     * and value includes 0, 1, 5 (default), 15, 60, and 180. If the value is set to 0, health check is performed every 10 seconds.
     */
    healthPeriodicAuditTime?: pulumi.Input<number>;
    /**
     * Specifies the instance removal policy. The policy has four
     * options: `OLD_CONFIG_OLD_INSTANCE` (default), `OLD_CONFIG_NEW_INSTANCE`, `OLD_INSTANCE`, and `NEW_INSTANCE`.
     */
    instanceTerminatePolicy?: pulumi.Input<string>;
    /**
     * The instances IDs of the AS group.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The system supports the binding of up to six ELB listeners, the IDs of which are separated using a comma.
     *
     * @deprecated use lbaas_listeners instead
     */
    lbListenerId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more enhanced load balancer. The system supports
     * the binding of up to six load balancers. The object structure is documented below.
     */
    lbaasListeners?: pulumi.Input<pulumi.Input<inputs.As.GroupLbaasListener>[]>;
    /**
     * Specifies the maximum number of instances. The default value is 0.
     */
    maxInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies the minimum number of instances. The default value is 0.
     */
    minInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies the priority policy used to select target AZs when adjusting
     * the number of instances in an AS group. The value can be `EQUILIBRIUM_DISTRIBUTE` and `PICK_FIRST`.
     */
    multiAzScalingPolicy?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more network IDs. The system supports up to five networks.
     * The object structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.As.GroupNetwork>[]>;
    /**
     * schema: Deprecated; The notification mode has been canceled
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region in which to create the AS group.
     * If omitted, the provider-level region will be used. Changing this creates a new group.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the configuration ID which defines configurations
     * of instances in the AS group.
     */
    scalingConfigurationId?: pulumi.Input<string>;
    /**
     * Specifies the name of the scaling group. The name can contain
     * letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.
     */
    scalingGroupName?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with the group.
     * The object structure is documented below.
     */
    securityGroups?: pulumi.Input<pulumi.Input<inputs.As.GroupSecurityGroup>[]>;
    /**
     * The status of the AS group.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the AS group.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the VPC ID. Changing this creates a new group.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Specifies the IAM agency name. If you change the agency,
     * the new agency will be available for ECSs scaled out after the change.
     */
    agencyName?: pulumi.Input<string>;
    /**
     * Specifies the availability zones in which to create the instances in the
     * autoscaling group.
     */
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * schema: Deprecated; use availability_zones instead
     */
    availableZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the cooling duration (in seconds). The value ranges from 0 to 86400,
     * and is 300 by default.
     */
    coolDownTime?: pulumi.Input<number>;
    /**
     * Specifies whether to delete the instances in the AS group when deleting
     * the AS group. The options are `yes` and `no`.
     */
    deleteInstances?: pulumi.Input<string>;
    /**
     * Specifies whether to delete the elastic IP address bound to the instances of
     * AS group when deleting the instances. The options are `true` and `false`.
     */
    deletePublicip?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the AS group.
     * The value can contain 0 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the expected number of instances. The default value is the
     * minimum number of instances. The value ranges from the minimum number of instances to the maximum number of instances.
     */
    desireInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies whether to enable the AS Group. The options are `true` and `false`.
     * The default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Specifies the enterprise project id of the AS group.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly delete the AS group, remove the ECS instances and
     * release them. The default value is `false`.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * Specifies the health check grace period for instances.
     * The unit is second and the value ranges from 0 to 86400. The default value is 600.
     */
    healthPeriodicAuditGracePeriod?: pulumi.Input<number>;
    /**
     * Specifies the health check method for instances in the AS group.
     * The health check methods include `ELB_AUDIT` and `NOVA_AUDIT`. If load balancing is configured, the default value of
     * this parameter is `ELB_AUDIT`. Otherwise, the default value is `NOVA_AUDIT`.
     */
    healthPeriodicAuditMethod?: pulumi.Input<string>;
    /**
     * Specifies the health check period for instances. The unit is minute
     * and value includes 0, 1, 5 (default), 15, 60, and 180. If the value is set to 0, health check is performed every 10 seconds.
     */
    healthPeriodicAuditTime?: pulumi.Input<number>;
    /**
     * Specifies the instance removal policy. The policy has four
     * options: `OLD_CONFIG_OLD_INSTANCE` (default), `OLD_CONFIG_NEW_INSTANCE`, `OLD_INSTANCE`, and `NEW_INSTANCE`.
     */
    instanceTerminatePolicy?: pulumi.Input<string>;
    /**
     * The system supports the binding of up to six ELB listeners, the IDs of which are separated using a comma.
     *
     * @deprecated use lbaas_listeners instead
     */
    lbListenerId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more enhanced load balancer. The system supports
     * the binding of up to six load balancers. The object structure is documented below.
     */
    lbaasListeners?: pulumi.Input<pulumi.Input<inputs.As.GroupLbaasListener>[]>;
    /**
     * Specifies the maximum number of instances. The default value is 0.
     */
    maxInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies the minimum number of instances. The default value is 0.
     */
    minInstanceNumber?: pulumi.Input<number>;
    /**
     * Specifies the priority policy used to select target AZs when adjusting
     * the number of instances in an AS group. The value can be `EQUILIBRIUM_DISTRIBUTE` and `PICK_FIRST`.
     */
    multiAzScalingPolicy?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more network IDs. The system supports up to five networks.
     * The object structure is documented below.
     */
    networks: pulumi.Input<pulumi.Input<inputs.As.GroupNetwork>[]>;
    /**
     * schema: Deprecated; The notification mode has been canceled
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region in which to create the AS group.
     * If omitted, the provider-level region will be used. Changing this creates a new group.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the configuration ID which defines configurations
     * of instances in the AS group.
     */
    scalingConfigurationId?: pulumi.Input<string>;
    /**
     * Specifies the name of the scaling group. The name can contain
     * letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.
     */
    scalingGroupName: pulumi.Input<string>;
    /**
     * Specifies an array of one or more security group IDs to associate with the group.
     * The object structure is documented below.
     */
    securityGroups?: pulumi.Input<pulumi.Input<inputs.As.GroupSecurityGroup>[]>;
    /**
     * Specifies the key/value pairs to associate with the AS group.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the VPC ID. Changing this creates a new group.
     */
    vpcId: pulumi.Input<string>;
}