// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get available HuaweiCloud gaussdb opengauss instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const thisOpengaussInstances = pulumi.output(huaweicloud.GaussDBforOpenGauss.getOpengaussInstances({
 *     name: "gaussdb-instance",
 * }));
 * ```
 */
export function getOpengaussInstances(args?: GetOpengaussInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetOpengaussInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:GaussDBforOpenGauss/getOpengaussInstances:getOpengaussInstances", {
        "name": args.name,
        "region": args.region,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpengaussInstances.
 */
export interface GetOpengaussInstancesArgs {
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
 * A collection of values returned by getOpengaussInstances.
 */
export interface GetOpengaussInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * An array of available instances.
     */
    readonly instances: outputs.GaussDBforOpenGauss.GetOpengaussInstancesInstance[];
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
     * Indicates the VPC ID.
     */
    readonly vpcId?: string;
}

export function getOpengaussInstancesOutput(args?: GetOpengaussInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOpengaussInstancesResult> {
    return pulumi.output(args).apply(a => getOpengaussInstances(a, opts))
}

/**
 * A collection of arguments for invoking getOpengaussInstances.
 */
export interface GetOpengaussInstancesOutputArgs {
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