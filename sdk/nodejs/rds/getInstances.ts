// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to list all available RDS instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const thisInstances = pulumi.output(huaweicloud.Rds.getInstances({
 *     name: "rds-instance",
 * }));
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Rds/getInstances:getInstances", {
        "datastoreType": args.datastoreType,
        "enterpriseProjectId": args.enterpriseProjectId,
        "name": args.name,
        "region": args.region,
        "subnetId": args.subnetId,
        "type": args.type,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Specifies the type of the database. Valid values are: MySQL, PostgreSQL, and SQLServer.
     */
    datastoreType?: string;
    /**
     * Specifies the enterprise project id.
     */
    enterpriseProjectId?: string;
    /**
     * Specifies the name of the instance.
     */
    name?: string;
    /**
     * The region in which to obtain the instances. If omitted, the provider-level region will
     * be used.
     */
    region?: string;
    /**
     * Specifies the network ID of a subnet.
     */
    subnetId?: string;
    /**
     * Specifies the type of the instance. Valid values are: Single, Ha, Replica, and Enterprise.
     */
    type?: string;
    /**
     * Specifies the VPC ID.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly datastoreType?: string;
    /**
     * Indicates the enterprise project id.
     */
    readonly enterpriseProjectId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * An array of available instances.
     */
    readonly instances: outputs.Rds.GetInstancesInstance[];
    /**
     * Indicates the node name.
     */
    readonly name?: string;
    /**
     * The region of the instance.
     */
    readonly region: string;
    /**
     * Indicates the network ID of a subnet.
     */
    readonly subnetId?: string;
    /**
     * Indicates the volume type.
     */
    readonly type?: string;
    /**
     * Indicates the VPC ID.
     */
    readonly vpcId?: string;
}

export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * Specifies the type of the database. Valid values are: MySQL, PostgreSQL, and SQLServer.
     */
    datastoreType?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the instances. If omitted, the provider-level region will
     * be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the network ID of a subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Specifies the type of the instance. Valid values are: Single, Ha, Replica, and Enterprise.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}