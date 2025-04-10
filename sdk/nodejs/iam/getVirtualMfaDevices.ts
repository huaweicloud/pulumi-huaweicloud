// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of IAM virtual MFA devices within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test = pulumi.output(huaweicloud.Iam.getVirtualMfaDevices());
 * ```
 */
export function getVirtualMfaDevices(args?: GetVirtualMfaDevicesArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMfaDevicesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iam/getVirtualMfaDevices:getVirtualMfaDevices", {
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualMfaDevices.
 */
export interface GetVirtualMfaDevicesArgs {
    /**
     * Specifies the user ID to which the virtual MFA device belongs.
     */
    userId?: string;
}

/**
 * A collection of values returned by getVirtualMfaDevices.
 */
export interface GetVirtualMfaDevicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The user ID to which the virtual MFA device belongs.
     */
    readonly userId?: string;
    /**
     * The list of virtual MFA devices.
     */
    readonly virtualMfaDevices: outputs.Iam.GetVirtualMfaDevicesVirtualMfaDevice[];
}

export function getVirtualMfaDevicesOutput(args?: GetVirtualMfaDevicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualMfaDevicesResult> {
    return pulumi.output(args).apply(a => getVirtualMfaDevices(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualMfaDevices.
 */
export interface GetVirtualMfaDevicesOutputArgs {
    /**
     * Specifies the user ID to which the virtual MFA device belongs.
     */
    userId?: pulumi.Input<string>;
}
