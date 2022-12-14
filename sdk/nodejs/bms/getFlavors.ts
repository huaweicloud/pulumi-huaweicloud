// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get available BMS flavors.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const demo = huaweicloud.Bms.getFlavors({
 *     availabilityZone: "cn-north-1a",
 *     vcpus: 48,
 * });
 * // Create BMS instance with the matched flavor
 * const instance = new huaweicloud.bms.Instance("instance", {flavorId: demo.then(demo => demo.flavors?[0]?.id)});
 * // Other properties...
 * ```
 */
export function getFlavors(args?: GetFlavorsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Bms/getFlavors:getFlavors", {
        "availabilityZone": args.availabilityZone,
        "cpuArch": args.cpuArch,
        "memory": args.memory,
        "region": args.region,
        "vcpus": args.vcpus,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsArgs {
    /**
     * Specifies the AZ name.
     */
    availabilityZone?: string;
    /**
     * Specifies the CPU architecture of the BMS flavor.
     * The value can be x8664 and aarch64, defaults to **x86_64**.
     */
    cpuArch?: string;
    /**
     * Specifies the memory size(GB) in the BMS flavor.
     */
    memory?: number;
    /**
     * The region in which to obtain the flavors.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
    /**
     * Specifies the number of vCPUs in the BMS flavor.
     */
    vcpus?: number;
}

/**
 * A collection of values returned by getFlavors.
 */
export interface GetFlavorsResult {
    readonly availabilityZone?: string;
    /**
     * The CPU architecture of the BMS flavor.
     */
    readonly cpuArch?: string;
    /**
     * Indicates the flavors information. Structure is documented below.
     */
    readonly flavors: outputs.Bms.GetFlavorsFlavor[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The memory size in GB.
     */
    readonly memory?: number;
    readonly region: string;
    /**
     * The number of vCPUs.
     */
    readonly vcpus?: number;
}

export function getFlavorsOutput(args?: GetFlavorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlavorsResult> {
    return pulumi.output(args).apply(a => getFlavors(a, opts))
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsOutputArgs {
    /**
     * Specifies the AZ name.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the CPU architecture of the BMS flavor.
     * The value can be x8664 and aarch64, defaults to **x86_64**.
     */
    cpuArch?: pulumi.Input<string>;
    /**
     * Specifies the memory size(GB) in the BMS flavor.
     */
    memory?: pulumi.Input<number>;
    /**
     * The region in which to obtain the flavors.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the number of vCPUs in the BMS flavor.
     */
    vcpus?: pulumi.Input<number>;
}
