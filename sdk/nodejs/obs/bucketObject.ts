// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an OBS bucket object resource.
 *
 * ## Example Usage
 * ### Uploading to a bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const object = new huaweicloud.Obs.BucketObject("object", {
 *     bucket: "your_bucket_name",
 *     content: "some object content",
 *     contentType: "application/xml",
 *     key: "new_key_from_content",
 * });
 * ```
 * ### Uploading a file to a bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const examplebucket = new huaweicloud.obs.Bucket("examplebucket", {
 *     bucket: "examplebuckettftest",
 *     acl: "private",
 * });
 * const object = new huaweicloud.obs.BucketObject("object", {
 *     bucket: examplebucket.bucket,
 *     key: "new_key_from_file",
 *     source: "index.html",
 * });
 * ```
 * ### Server Side Encryption with OBS Default Master Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const examplebucketObject = new huaweicloud.Obs.BucketObject("examplebucket_object", {
 *     bucket: "your_bucket_name",
 *     encryption: true,
 *     key: "someobject",
 *     source: "index.html",
 * });
 * ```
 *
 * ## Import
 *
 * OBS bucket object can be imported using the bucket and key separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Obs/bucketObject:BucketObject object bucket/key
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`encryption`, `source`, `acl` and `kms_key_id`. It is generally recommended running `terraform plan` after importing an object. You can then decide if changes should be applied to the object, or the resource definition should be updated to align with the object. Also you can ignore changes as below. resource "huaweicloud_obs_bucket_object" "object" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  encryption, source, acl, kms_key_id,
 *
 *  ]
 *
 *  } }
 */
export class BucketObject extends pulumi.CustomResource {
    /**
     * Get an existing BucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketObjectState, opts?: pulumi.CustomResourceOptions): BucketObject {
        return new BucketObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Obs/bucketObject:BucketObject';

    /**
     * Returns true if the given object is an instance of BucketObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketObject.__pulumiType;
    }

    /**
     * The ACL policy to apply. Defaults to `private`.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The name of the bucket to put the file in.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The literal content being uploaded to the bucket.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * A standard MIME type describing the format of the object data, e.g.
     * application/octet-stream. All Valid MIME Types are valid for this input.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * Whether enable server-side encryption of the object in SSE-KMS mode.
     */
    public readonly encryption!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the unique identifier of the object content. It can be used to trigger updates.
     * The only meaningful value is `md5(file("pathToFile"))`.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the kms key. If omitted, the default master key will be used.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the OBS bucket object resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket object resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * the size of the object in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The path to the source file being uploaded to the bucket.
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * Specifies the storage class of the object. Defaults to `STANDARD`.
     */
    public readonly storageClass!: pulumi.Output<string>;
    /**
     * A unique version ID value for the object, if bucket versioning is enabled.
     */
    public /*out*/ readonly versionId!: pulumi.Output<string>;

    /**
     * Create a BucketObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketObjectArgs | BucketObjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketObjectState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["storageClass"] = state ? state.storageClass : undefined;
            resourceInputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as BucketObjectArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["storageClass"] = args ? args.storageClass : undefined;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketObject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketObject resources.
 */
export interface BucketObjectState {
    /**
     * The ACL policy to apply. Defaults to `private`.
     */
    acl?: pulumi.Input<string>;
    /**
     * The name of the bucket to put the file in.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The literal content being uploaded to the bucket.
     */
    content?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g.
     * application/octet-stream. All Valid MIME Types are valid for this input.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Whether enable server-side encryption of the object in SSE-KMS mode.
     */
    encryption?: pulumi.Input<boolean>;
    /**
     * Specifies the unique identifier of the object content. It can be used to trigger updates.
     * The only meaningful value is `md5(file("pathToFile"))`.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the kms key. If omitted, the default master key will be used.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The region in which to create the OBS bucket object resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket object resource.
     */
    region?: pulumi.Input<string>;
    /**
     * the size of the object in bytes.
     */
    size?: pulumi.Input<number>;
    /**
     * The path to the source file being uploaded to the bucket.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the storage class of the object. Defaults to `STANDARD`.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * A unique version ID value for the object, if bucket versioning is enabled.
     */
    versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketObject resource.
 */
export interface BucketObjectArgs {
    /**
     * The ACL policy to apply. Defaults to `private`.
     */
    acl?: pulumi.Input<string>;
    /**
     * The name of the bucket to put the file in.
     */
    bucket: pulumi.Input<string>;
    /**
     * The literal content being uploaded to the bucket.
     */
    content?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g.
     * application/octet-stream. All Valid MIME Types are valid for this input.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Whether enable server-side encryption of the object in SSE-KMS mode.
     */
    encryption?: pulumi.Input<boolean>;
    /**
     * Specifies the unique identifier of the object content. It can be used to trigger updates.
     * The only meaningful value is `md5(file("pathToFile"))`.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name of the object once it is in the bucket.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the kms key. If omitted, the default master key will be used.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The region in which to create the OBS bucket object resource. If omitted, the
     * provider-level region will be used. Changing this creates a new OBS bucket object resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The path to the source file being uploaded to the bucket.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the storage class of the object. Defaults to `STANDARD`.
     */
    storageClass?: pulumi.Input<string>;
}