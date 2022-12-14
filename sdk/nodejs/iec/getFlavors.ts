// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the available of HuaweiCloud IEC flavors.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const flavorName = config.get("flavorName") || "c6.large.2";
 * const iecFlavorTest = huaweicloud.Iec.getFlavors({
 *     name: flavorName,
 * });
 * ```
 */
export function getFlavors(args?: GetFlavorsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iec/getFlavors:getFlavors", {
        "area": args.area,
        "city": args.city,
        "name": args.name,
        "operator": args.operator,
        "province": args.province,
        "region": args.region,
        "siteIds": args.siteIds,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsArgs {
    /**
     * Specifies the province of the iec instance located.
     */
    area?: string;
    /**
     * Specifies the province of the iec instance located.
     */
    city?: string;
    /**
     * Specifies the flavor name, which can be queried with a regular expression.
     */
    name?: string;
    /**
     * Specifies the operator supported of the iec instance.
     */
    operator?: string;
    /**
     * Specifies the province of the iec instance located.
     */
    province?: string;
    /**
     * The region in which to obtain the flavors. If omitted, the provider-level region will be
     * used.
     */
    region?: string;
    /**
     * Specifies the list of edge service site.
     */
    siteIds?: string;
}

/**
 * A collection of values returned by getFlavors.
 */
export interface GetFlavorsResult {
    readonly area?: string;
    readonly city?: string;
    /**
     * An array of one or more flavors. The flavors object structure is documented below.
     */
    readonly flavors: outputs.Iec.GetFlavorsFlavor[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the iec flavor.
     */
    readonly name?: string;
    readonly operator?: string;
    readonly province?: string;
    readonly region: string;
    readonly siteIds?: string;
}

export function getFlavorsOutput(args?: GetFlavorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlavorsResult> {
    return pulumi.output(args).apply(a => getFlavors(a, opts))
}

/**
 * A collection of arguments for invoking getFlavors.
 */
export interface GetFlavorsOutputArgs {
    /**
     * Specifies the province of the iec instance located.
     */
    area?: pulumi.Input<string>;
    /**
     * Specifies the province of the iec instance located.
     */
    city?: pulumi.Input<string>;
    /**
     * Specifies the flavor name, which can be queried with a regular expression.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the operator supported of the iec instance.
     */
    operator?: pulumi.Input<string>;
    /**
     * Specifies the province of the iec instance located.
     */
    province?: pulumi.Input<string>;
    /**
     * The region in which to obtain the flavors. If omitted, the provider-level region will be
     * used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the list of edge service site.
     */
    siteIds?: pulumi.Input<string>;
}
