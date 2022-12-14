// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the available ELB Flavors.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const flavors = huaweicloud.DedicatedElb.getFlavors({
 *     type: "L7",
 *     maxConnections: 200000,
 *     cps: 2000,
 *     bandwidth: 50,
 * });
 * const lb = new huaweicloud.dedicatedelb.Loadbalancer("lb", {l7FlavorId: flavors.then(flavors => flavors.ids?[0])});
 * // Other properties...
 * ```
 */
export function getFlavors(args?: GetFlavorsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:DedicatedElb/getFlavors:getFlavors", {
        "bandwidth": args.bandwidth,
        "cps": args.cps,
        "maxConnections": args.maxConnections,
        "qps": args.qps,
        "region": args.region,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsArgs {
    /**
     * Specifies the bandwidth size(Mbit/s) in the flavor.
     */
    bandwidth?: number;
    /**
     * Specifies the cps in the flavor.
     */
    cps?: number;
    /**
     * Specifies the maximum connections in the flavor.
     */
    maxConnections?: number;
    /**
     * Specifies the qps in the L7 flavor.
     */
    qps?: number;
    /**
     * The region in which to obtain the flavors. If omitted, the provider-level region will be
     * used.
     */
    region?: string;
    /**
     * Specifies the flavor type. Valid values are L4 and L7.
     */
    type?: string;
}

/**
 * A collection of values returned by getFlavors.
 */
export interface GetFlavorsResult {
    /**
     * Bandwidth size(Mbit/s) of the flavor.
     */
    readonly bandwidth?: number;
    /**
     * Cps of the flavor.
     */
    readonly cps?: number;
    /**
     * A list of flavors. Each element contains the following attributes:
     */
    readonly flavors: outputs.DedicatedElb.GetFlavorsFlavor[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of flavor IDs.
     */
    readonly ids: string[];
    /**
     * Maximum connections of the flavor.
     */
    readonly maxConnections?: number;
    /**
     * Qps of the L7 flavor.
     */
    readonly qps?: number;
    readonly region: string;
    /**
     * Type of the flavor.
     */
    readonly type?: string;
}

export function getFlavorsOutput(args?: GetFlavorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlavorsResult> {
    return pulumi.output(args).apply(a => getFlavors(a, opts))
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsOutputArgs {
    /**
     * Specifies the bandwidth size(Mbit/s) in the flavor.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Specifies the cps in the flavor.
     */
    cps?: pulumi.Input<number>;
    /**
     * Specifies the maximum connections in the flavor.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Specifies the qps in the L7 flavor.
     */
    qps?: pulumi.Input<number>;
    /**
     * The region in which to obtain the flavors. If omitted, the provider-level region will be
     * used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the flavor type. Valid values are L4 and L7.
     */
    type?: pulumi.Input<string>;
}
