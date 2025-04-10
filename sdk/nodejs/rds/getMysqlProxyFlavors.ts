// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of RDS MySQL proxy flavors.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const flavor = huaweicloud.Rds.getMysqlProxyFlavors({
 *     instanceId: instanceId,
 * });
 * ```
 */
export function getMysqlProxyFlavors(args: GetMysqlProxyFlavorsArgs, opts?: pulumi.InvokeOptions): Promise<GetMysqlProxyFlavorsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Rds/getMysqlProxyFlavors:getMysqlProxyFlavors", {
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getMysqlProxyFlavors.
 */
export interface GetMysqlProxyFlavorsArgs {
    /**
     * Specifies the ID of RDS MySQL instance.
     */
    instanceId: string;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getMysqlProxyFlavors.
 */
export interface GetMysqlProxyFlavorsResult {
    /**
     * Indicates the list of flavor groups.
     * The flavorGroups structure is documented below.
     */
    readonly flavorGroups: outputs.Rds.GetMysqlProxyFlavorsFlavorGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region: string;
}

export function getMysqlProxyFlavorsOutput(args: GetMysqlProxyFlavorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMysqlProxyFlavorsResult> {
    return pulumi.output(args).apply(a => getMysqlProxyFlavors(a, opts))
}

/**
 * A collection of arguments for invoking getMysqlProxyFlavors.
 */
export interface GetMysqlProxyFlavorsOutputArgs {
    /**
     * Specifies the ID of RDS MySQL instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
}
