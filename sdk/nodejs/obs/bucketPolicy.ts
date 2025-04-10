// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a policy to an OBS bucket resource.
 *
 * > **NOTE:** When creating or updating the OBS bucket policy, the original policy will be overwritten.
 *
 * ## Example Usage
 * ### Policy with OBS format
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const bucket = new huaweicloud.obs.Bucket("bucket", {bucket: "my-test-bucket"});
 * const policy = new huaweicloud.obs.BucketPolicy("policy", {
 *     bucket: bucket.id,
 *     policy: `{
 *   "Statement": [
 *     {
 *       "Sid": "AddPerm",
 *       "Effect": "Allow",
 *       "Principal": {"ID": "*"},
 *       "Action": ["GetObject"],
 *       "Resource": "my-test-bucket/*"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 * ### Policy with S3 format
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const bucket = new huaweicloud.obs.Bucket("bucket", {bucket: "my-test-bucket"});
 * const s3Policy = new huaweicloud.obs.BucketPolicy("s3Policy", {
 *     bucket: bucket.id,
 *     policyFormat: "s3",
 *     policy: `{
 *   "Version": "2008-10-17",
 *   "Id": "MYBUCKETPOLICY",
 *   "Statement": [
 *     {
 *       "Sid": "IPAllow",
 *       "Effect": "Allow",
 *       "Principal": "*",
 *       "Action": "s3:*",
 *       "Resource": "arn:aws:s3:::my-test-bucket/*",
 *       "Condition": {
 *         "IpAddress": {"aws:SourceIp": "8.8.8.8/32"}
 *       }
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * OBS format bucket policy can be imported using the `<bucket>`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Obs/bucketPolicy:BucketPolicy policy <bucket-name>
 * ```
 *
 *  S3 format bucket policy can be imported using the `<bucket>` and "s3" by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Obs/bucketPolicy:BucketPolicy s3_policy <bucket-name>/s3
 * ```
 */
export class BucketPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketPolicyState, opts?: pulumi.CustomResourceOptions): BucketPolicy {
        return new BucketPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Obs/bucketPolicy:BucketPolicy';

    /**
     * Returns true if the given object is an instance of BucketPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketPolicy.__pulumiType;
    }

    /**
     * Specifies the name of the bucket to which to apply the policy.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Specifies the text of the bucket policy in JSON format. For more information about obs
     * format bucket policy,
     * see the [Developer Guide](https://support.huaweicloud.com/intl/en-us/perms-cfg-obs/obs_40_0004.html).
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Specifies the policy format, the supported values are *obs* and *s3*. Defaults
     * to *obs* .
     */
    public readonly policyFormat!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the OBS bucket policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket policy resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a BucketPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketPolicyArgs | BucketPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketPolicyState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["policyFormat"] = state ? state.policyFormat : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as BucketPolicyArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["policyFormat"] = args ? args.policyFormat : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketPolicy resources.
 */
export interface BucketPolicyState {
    /**
     * Specifies the name of the bucket to which to apply the policy.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Specifies the text of the bucket policy in JSON format. For more information about obs
     * format bucket policy,
     * see the [Developer Guide](https://support.huaweicloud.com/intl/en-us/perms-cfg-obs/obs_40_0004.html).
     */
    policy?: pulumi.Input<string>;
    /**
     * Specifies the policy format, the supported values are *obs* and *s3*. Defaults
     * to *obs* .
     */
    policyFormat?: pulumi.Input<string>;
    /**
     * The region in which to create the OBS bucket policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket policy resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketPolicy resource.
 */
export interface BucketPolicyArgs {
    /**
     * Specifies the name of the bucket to which to apply the policy.
     */
    bucket: pulumi.Input<string>;
    /**
     * Specifies the text of the bucket policy in JSON format. For more information about obs
     * format bucket policy,
     * see the [Developer Guide](https://support.huaweicloud.com/intl/en-us/perms-cfg-obs/obs_40_0004.html).
     */
    policy: pulumi.Input<string>;
    /**
     * Specifies the policy format, the supported values are *obs* and *s3*. Defaults
     * to *obs* .
     */
    policyFormat?: pulumi.Input<string>;
    /**
     * The region in which to create the OBS bucket policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket policy resource.
     */
    region?: pulumi.Input<string>;
}
