// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query the IAM project list within HuaweiCloud.
 *
 * > **NOTE:** You *must* have IAM read privileges to use this data source.
 *
 * ## Example Usage
 * ### Obtain project information by name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.Iam.getProjects({
 *     name: "cn-north-4_demo",
 * }));
 * ```
 * ### Obtain special project information by name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.Iam.getProjects({
 *     name: "MOS", // The project for OBS Billing
 * }));
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iam/getProjects:getProjects", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * Specifies the IAM project name to query.
     */
    name?: string;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IAM project name.
     */
    readonly name?: string;
    /**
     * The details of the query projects. The structure is documented below.
     */
    readonly projects: outputs.Iam.GetProjectsProject[];
}

export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply(a => getProjects(a, opts))
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    /**
     * Specifies the IAM project name to query.
     */
    name?: pulumi.Input<string>;
}
