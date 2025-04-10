// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of IAM identity providers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const providerName = config.requireObject("providerName");
 * const test = huaweicloud.Iam.getProviders({
 *     name: providerName,
 * });
 * ```
 */
export function getProviders(args?: GetProvidersArgs, opts?: pulumi.InvokeOptions): Promise<GetProvidersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Iam/getProviders:getProviders", {
        "name": args.name,
        "ssoType": args.ssoType,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviders.
 */
export interface GetProvidersArgs {
    /**
     * Specifies the name of the identity provider.
     */
    name?: string;
    /**
     * Specifies the single sign-on type of the identity provider.
     */
    ssoType?: string;
    /**
     * Specifies the status of the identity provider. The value can be **true** or **false**
     */
    status?: string;
}

/**
 * A collection of values returned by getProviders.
 */
export interface GetProvidersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of identity providers.
     */
    readonly identityProviders: outputs.Iam.GetProvidersIdentityProvider[];
    readonly name?: string;
    /**
     * The single sign-on type of the identity provider.
     */
    readonly ssoType?: string;
    /**
     * The enabled status for the identity provider.
     */
    readonly status?: string;
}

export function getProvidersOutput(args?: GetProvidersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProvidersResult> {
    return pulumi.output(args).apply(a => getProviders(a, opts))
}

/**
 * A collection of arguments for invoking getProviders.
 */
export interface GetProvidersOutputArgs {
    /**
     * Specifies the name of the identity provider.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the single sign-on type of the identity provider.
     */
    ssoType?: pulumi.Input<string>;
    /**
     * Specifies the status of the identity provider. The value can be **true** or **false**
     */
    status?: pulumi.Input<string>;
}
