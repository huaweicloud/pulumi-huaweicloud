// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of feature configurations of ELB of a tenant.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.DedicatedElb.getFeatureConfigurations());
 * ```
 */
export function getFeatureConfigurations(args?: GetFeatureConfigurationsArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureConfigurationsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:DedicatedElb/getFeatureConfigurations:getFeatureConfigurations", {
        "feature": args.feature,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureConfigurations.
 */
export interface GetFeatureConfigurationsArgs {
    /**
     * Specifies the feature name.
     */
    feature?: string;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getFeatureConfigurations.
 */
export interface GetFeatureConfigurationsResult {
    /**
     * Indicates the feature configuration list.
     */
    readonly configs: outputs.DedicatedElb.GetFeatureConfigurationsConfig[];
    /**
     * Indicates the feature name.
     */
    readonly feature?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region: string;
}

export function getFeatureConfigurationsOutput(args?: GetFeatureConfigurationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureConfigurationsResult> {
    return pulumi.output(args).apply(a => getFeatureConfigurations(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureConfigurations.
 */
export interface GetFeatureConfigurationsOutputArgs {
    /**
     * Specifies the feature name.
     */
    feature?: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
}
