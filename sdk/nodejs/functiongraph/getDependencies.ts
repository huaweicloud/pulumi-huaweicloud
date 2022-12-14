// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to filter dependent packages of FGS from HuaweiCloud.
 *
 * ## Example Usage
 * ### Obtain all public dependent packages
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.FunctionGraph.getDependencies());
 * ```
 * ### Obtain specific public dependent package by name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.FunctionGraph.getDependencies({
 *     name: "obssdk-3.0.2",
 *     type: "public",
 * }));
 * ```
 * ### Obtain all public Python2.7 dependent packages
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.FunctionGraph.getDependencies({
 *     runtime: "Python2.7",
 *     type: "public",
 * }));
 * ```
 */
export function getDependencies(args?: GetDependenciesArgs, opts?: pulumi.InvokeOptions): Promise<GetDependenciesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:FunctionGraph/getDependencies:getDependencies", {
        "name": args.name,
        "region": args.region,
        "runtime": args.runtime,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getDependencies.
 */
export interface GetDependenciesArgs {
    /**
     * Specifies the dependent package runtime to match.
     */
    name?: string;
    /**
     * Specifies the region in which to obtain the dependent packages. If omitted, the
     * provider-level region will be used.
     */
    region?: string;
    /**
     * Specifies the dependent package runtime to match. Valid values: **Java8**,
     * **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**, **Python3.6**, **Go1.8**,
     * **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and **PHP7.3**.
     */
    runtime?: string;
    /**
     * Specifies the dependent package type to match. Valid values: **public** and **private**.
     */
    type?: string;
}

/**
 * A collection of values returned by getDependencies.
 */
export interface GetDependenciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Dependent package name.
     */
    readonly name?: string;
    /**
     * All dependent packages that match.
     */
    readonly packages: outputs.FunctionGraph.GetDependenciesPackage[];
    readonly region: string;
    /**
     * Dependent package runtime.
     */
    readonly runtime?: string;
    readonly type?: string;
}

export function getDependenciesOutput(args?: GetDependenciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDependenciesResult> {
    return pulumi.output(args).apply(a => getDependencies(a, opts))
}

/**
 * A collection of arguments for invoking getDependencies.
 */
export interface GetDependenciesOutputArgs {
    /**
     * Specifies the dependent package runtime to match.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to obtain the dependent packages. If omitted, the
     * provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the dependent package runtime to match. Valid values: **Java8**,
     * **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**, **Python3.6**, **Go1.8**,
     * **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and **PHP7.3**.
     */
    runtime?: pulumi.Input<string>;
    /**
     * Specifies the dependent package type to match. Valid values: **public** and **private**.
     */
    type?: pulumi.Input<string>;
}
