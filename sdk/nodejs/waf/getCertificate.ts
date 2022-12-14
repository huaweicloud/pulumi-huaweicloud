// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get the certificate in the WAF, including the one pushed from SCM.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const certificate1 = huaweicloud.Waf.getCertificate({
 *     name: "certificate name",
 * });
 * const domain1 = new huaweicloud.waf.Domain("domain1", {
 *     domain: "www.domainname.com",
 *     certificateId: certificate1.then(certificate1 => certificate1.id),
 *     certificateName: certificate1.then(certificate1 => certificate1.name),
 *     keepPolicy: false,
 *     proxy: false,
 *     servers: [{
 *         clientProtocol: "HTTPS",
 *         serverProtocol: "HTTP",
 *         address: "192.168.10.1",
 *         port: 8080,
 *     }],
 * });
 * ```
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Waf/getCertificate:getCertificate", {
        "expireStatus": args.expireStatus,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * The expire status of certificate. Defaults is `0`. The value can be:
     * + `0`: not expire
     * + `1`: has expired
     * + `2`: wil expired soon
     */
    expireStatus?: number;
    /**
     * The name of certificate. The value is case sensitive and supports fuzzy matching.
     */
    name: string;
    /**
     * The region in which to obtain the WAF. If omitted, the provider-level region will be
     * used.
     */
    region?: string;
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    /**
     * Indicates the time when the certificate expires.
     */
    readonly expiration: string;
    readonly expireStatus?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly region: string;
}

export function getCertificateOutput(args: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply(a => getCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateOutputArgs {
    /**
     * The expire status of certificate. Defaults is `0`. The value can be:
     * + `0`: not expire
     * + `1`: has expired
     * + `2`: wil expired soon
     */
    expireStatus?: pulumi.Input<number>;
    /**
     * The name of certificate. The value is case sensitive and supports fuzzy matching.
     */
    name: pulumi.Input<string>;
    /**
     * The region in which to obtain the WAF. If omitted, the provider-level region will be
     * used.
     */
    region?: pulumi.Input<string>;
}
