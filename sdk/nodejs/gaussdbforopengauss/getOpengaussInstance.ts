// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get available HuaweiCloud gaussdb opengauss instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const thisOpengaussInstance = pulumi.output(huaweicloud.GaussDBforOpenGauss.getOpengaussInstance({
 *     name: "gaussdb-instance",
 * }));
 * ```
 */
export function getOpengaussInstance(args?: GetOpengaussInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetOpengaussInstanceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:GaussDBforOpenGauss/getOpengaussInstance:getOpengaussInstance", {
        "name": args.name,
        "region": args.region,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpengaussInstance.
 */
export interface GetOpengaussInstanceArgs {
    /**
     * Specifies the name of the instance.
     */
    name?: string;
    /**
     * The region in which to obtain the instance. If omitted, the provider-level region will
     * be used.
     */
    region?: string;
    /**
     * Specifies the network ID of a subnet.
     */
    subnetId?: string;
    /**
     * Specifies the VPC ID.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getOpengaussInstance.
 */
export interface GetOpengaussInstanceResult {
    /**
     * Indicates the availability zone where the node resides.
     */
    readonly availabilityZone: string;
    /**
     * Indicates the advanced backup policy. Structure is documented below.
     */
    readonly backupStrategies: outputs.GaussDBforOpenGauss.GetOpengaussInstanceBackupStrategy[];
    /**
     * Indicates the count of coordinator node.
     */
    readonly coordinatorNum: number;
    /**
     * Indicates the database information. Structure is documented below.
     */
    readonly datastores: outputs.GaussDBforOpenGauss.GetOpengaussInstanceDatastore[];
    /**
     * Indicates the default username.
     */
    readonly dbUserName: string;
    /**
     * Indicates the enterprise project id.
     */
    readonly enterpriseProjectId: string;
    /**
     * Indicates the instance specifications.
     */
    readonly flavor: string;
    /**
     * Indicates the instance ha information. Structure is documented below.
     */
    readonly has: outputs.GaussDBforOpenGauss.GetOpengaussInstanceHa[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Indicates the maintenance window.
     */
    readonly maintenanceWindow: string;
    /**
     * Indicates the node name.
     */
    readonly name: string;
    /**
     * Indicates the instance nodes information. Structure is documented below.
     */
    readonly nodes: outputs.GaussDBforOpenGauss.GetOpengaussInstanceNode[];
    /**
     * Indicates the database port.
     */
    readonly port: number;
    /**
     * Indicates the list of private IP address of the nodes.
     */
    readonly privateIps: string[];
    /**
     * Indicates the public IP address of the DB instance.
     */
    readonly publicIps: string[];
    readonly region: string;
    /**
     * Indicates the replica num.
     */
    readonly replicaNum: number;
    /**
     * Indicates the security group ID.
     */
    readonly securityGroupId: string;
    /**
     * Indicates the sharding num.
     */
    readonly shardingNum: number;
    /**
     * Indicates the node status.
     */
    readonly status: string;
    readonly subnetId: string;
    /**
     * Indicates the switch strategy.
     */
    readonly switchStrategy: string;
    /**
     * Indicates the default username.
     */
    readonly timeZone: string;
    /**
     * Indicates the volume type. Value options: **ULTRAHIGH**, **ESSD**.
     */
    readonly type: string;
    /**
     * Indicates the volume information. Structure is documented below.
     */
    readonly volumes: outputs.GaussDBforOpenGauss.GetOpengaussInstanceVolume[];
    readonly vpcId: string;
}

export function getOpengaussInstanceOutput(args?: GetOpengaussInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOpengaussInstanceResult> {
    return pulumi.output(args).apply(a => getOpengaussInstance(a, opts))
}

/**
 * A collection of arguments for invoking getOpengaussInstance.
 */
export interface GetOpengaussInstanceOutputArgs {
    /**
     * Specifies the name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the instance. If omitted, the provider-level region will
     * be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the network ID of a subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Specifies the VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}
