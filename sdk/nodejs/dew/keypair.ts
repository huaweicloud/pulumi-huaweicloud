// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a keypair resource within HuaweiCloud.
 *
 * By default, key pairs use the SSH-2 (RSA, 2048) algorithm for encryption and decryption.
 *
 * Keys imported support the following cryptographic algorithms:
 *
 *  * RSA-1024
 *  * RSA-2048
 *  * RSA-4096
 *
 * ## Example Usage
 * ### Create a new keypair which scope is Tenant-level and the private key is managed by HuaweiCloud
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const test = new huaweicloud.dew.Key("test", {keyAlias: "kms_test"});
 * const test_keypair = new huaweicloud.dew.Keypair("test-keypair", {
 *     scope: "account",
 *     encryptionType: "kms",
 *     kmsKeyName: test.keyAlias,
 * });
 * ```
 * ### Create a new keypair and export private key to current folder
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test_keypair = new huaweicloud.Dew.Keypair("test-keypair", {
 *     keyFile: "private_key.pem",
 * });
 * ```
 * ### Import an existing keypair
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test_keypair = new huaweicloud.Dew.Keypair("test-keypair", {
 *     publicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAlJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcAT... root@terra-dev",
 * });
 * ```
 *
 * ## Import
 *
 * Keypairs can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dew/keypair:Keypair my-keypair test-keypair
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`encryption_type`, and `kms_key_name`. It is generally recommended running `terraform plan` after importing a key pair. You can then decide if changes should be applied to the key pair, or the resource definition should be updated to align with the key pair. Also you can ignore changes as below. resource "huaweicloud_kps_keypair" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  encryption_type, kms_key_name,
 *
 *  ]
 *
 *  } }
 */
export class Keypair extends pulumi.CustomResource {
    /**
     * Get an existing Keypair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeypairState, opts?: pulumi.CustomResourceOptions): Keypair {
        return new Keypair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dew/keypair:Keypair';

    /**
     * Returns true if the given object is an instance of Keypair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Keypair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Keypair.__pulumiType;
    }

    /**
     * The key pair creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the description of key pair.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies encryption mode if manages the private key by HuaweiCloud.
     * The options are as follows:
     * - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
     * - **kms**: KMS encryption mode.
     */
    public readonly encryptionType!: pulumi.Output<string>;
    /**
     * Fingerprint information about an key pair.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Whether the private key is managed by HuaweiCloud.
     */
    public /*out*/ readonly isManaged!: pulumi.Output<boolean>;
    /**
     * Specifies the path of the created private key.
     * The private key file (**.pem**) is created only after the resource is created.
     * Changing this parameter will create a new resource.
     */
    public readonly keyFile!: pulumi.Output<string>;
    /**
     * Specifies the KMS key name to encrypt private keys.
     * It's mandatory when the `encryptionType` is `kms`. Changing this parameter will create a new resource.
     */
    public readonly kmsKeyName!: pulumi.Output<string | undefined>;
    /**
     * Specifies a unique name for the keypair. The name can contain a maximum of 64
     * characters, including letters, digits, underscores (_) and hyphens (-).
     * Changing this parameter will create a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the imported OpenSSH-formatted public key.
     * Changing this parameter will create a new resource.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the keypair resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the scope of key pair. The options are as follows:
     * - **account**: Tenant-level, available to all users under the same account.
     * - **user**: User-level, only available to that user.
     */
    public readonly scope!: pulumi.Output<string>;

    /**
     * Create a Keypair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeypairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeypairArgs | KeypairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeypairState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["isManaged"] = state ? state.isManaged : undefined;
            resourceInputs["keyFile"] = state ? state.keyFile : undefined;
            resourceInputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as KeypairArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["keyFile"] = args ? args.keyFile : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["isManaged"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Keypair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Keypair resources.
 */
export interface KeypairState {
    /**
     * The key pair creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the description of key pair.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies encryption mode if manages the private key by HuaweiCloud.
     * The options are as follows:
     * - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
     * - **kms**: KMS encryption mode.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * Fingerprint information about an key pair.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * Whether the private key is managed by HuaweiCloud.
     */
    isManaged?: pulumi.Input<boolean>;
    /**
     * Specifies the path of the created private key.
     * The private key file (**.pem**) is created only after the resource is created.
     * Changing this parameter will create a new resource.
     */
    keyFile?: pulumi.Input<string>;
    /**
     * Specifies the KMS key name to encrypt private keys.
     * It's mandatory when the `encryptionType` is `kms`. Changing this parameter will create a new resource.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Specifies a unique name for the keypair. The name can contain a maximum of 64
     * characters, including letters, digits, underscores (_) and hyphens (-).
     * Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the imported OpenSSH-formatted public key.
     * Changing this parameter will create a new resource.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the keypair resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the scope of key pair. The options are as follows:
     * - **account**: Tenant-level, available to all users under the same account.
     * - **user**: User-level, only available to that user.
     */
    scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Keypair resource.
 */
export interface KeypairArgs {
    /**
     * Specifies the description of key pair.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies encryption mode if manages the private key by HuaweiCloud.
     * The options are as follows:
     * - **default**: The default encryption mode. Applicable to sites where KMS is not deployed.
     * - **kms**: KMS encryption mode.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * Specifies the path of the created private key.
     * The private key file (**.pem**) is created only after the resource is created.
     * Changing this parameter will create a new resource.
     */
    keyFile?: pulumi.Input<string>;
    /**
     * Specifies the KMS key name to encrypt private keys.
     * It's mandatory when the `encryptionType` is `kms`. Changing this parameter will create a new resource.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Specifies a unique name for the keypair. The name can contain a maximum of 64
     * characters, including letters, digits, underscores (_) and hyphens (-).
     * Changing this parameter will create a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the imported OpenSSH-formatted public key.
     * Changing this parameter will create a new resource.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the keypair resource. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the scope of key pair. The options are as follows:
     * - **account**: Tenant-level, available to all users under the same account.
     * - **user**: User-level, only available to that user.
     */
    scope?: pulumi.Input<string>;
}
