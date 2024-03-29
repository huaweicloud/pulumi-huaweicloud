// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the plaintext and the ciphertext of an available HuaweiCloud KMS DEK (data encryption key).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const key1 = new huaweicloud.dew.Key("key1", {
 *     keyAlias: "key_1",
 *     pendingDays: "7",
 *     keyDescription: "first test key",
 * });
 * const kmsDatakey1 = huaweicloud.Dew.getDataKeyOutput({
 *     keyId: key1.id,
 *     datakeyLength: "512",
 * });
 * ```
 */
export function getDataKey(args: GetDataKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetDataKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Dew/getDataKey:getDataKey", {
        "datakeyLength": args.datakeyLength,
        "encryptionContext": args.encryptionContext,
        "keyId": args.keyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataKey.
 */
export interface GetDataKeyArgs {
    /**
     * Number of bits in the length of a DEK (data encryption keys). The maximum number
     * is 512. Changing this gets the new data encryption key.
     */
    datakeyLength: string;
    /**
     * The value of this parameter must be a series of
     * "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive
     * information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
     */
    encryptionContext?: string;
    /**
     * The globally unique identifier for the key. Changing this gets the new data encryption
     * key.
     */
    keyId: string;
    /**
     * The region in which to obtain the keys. If omitted, the provider-level region will be
     * used.
     */
    region?: string;
}

/**
 * A collection of values returned by getDataKey.
 */
export interface GetDataKeyResult {
    /**
     * The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
     */
    readonly cipherText: string;
    readonly datakeyLength: string;
    readonly encryptionContext?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyId: string;
    /**
     * The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
     */
    readonly plainText: string;
    readonly region: string;
}

export function getDataKeyOutput(args: GetDataKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataKeyResult> {
    return pulumi.output(args).apply(a => getDataKey(a, opts))
}

/**
 * A collection of arguments for invoking getDataKey.
 */
export interface GetDataKeyOutputArgs {
    /**
     * Number of bits in the length of a DEK (data encryption keys). The maximum number
     * is 512. Changing this gets the new data encryption key.
     */
    datakeyLength: pulumi.Input<string>;
    /**
     * The value of this parameter must be a series of
     * "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive
     * information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
     */
    encryptionContext?: pulumi.Input<string>;
    /**
     * The globally unique identifier for the key. Changing this gets the new data encryption
     * key.
     */
    keyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the keys. If omitted, the provider-level region will be
     * used.
     */
    region?: pulumi.Input<string>;
}
