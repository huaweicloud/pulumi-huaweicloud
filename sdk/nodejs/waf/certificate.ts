// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a WAF certificate resource within HuaweiCloud.
 *
 * > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
 * used. The certificate resource can be used in Cloud Mode and Dedicated Mode.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const enterpriseProjectId = config.requireObject("enterpriseProjectId");
 * const test = new huaweicloud.waf.Certificate("test", {
 *     enterpriseProjectId: enterpriseProjectId,
 *     certificate: `-----BEGIN CERTIFICATE-----
 * MIIFmQl5dh2QUAeo39TIKtadgAgh4zHx09kSgayS9Wph9LEqq7MA+2042L3J9aOa
 * DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUR+SosWwALt6PkP0J9iOIxA6RW8gVsLwq
 * ...
 * +HhDvD/VeOHytX3RAs2GeTOtxyAV5XpKY5r+PkyUqPJj04t3d0Fopi0gNtLpMF=
 * -----END CERTIFICATE-----
 * `,
 *     privateKey: `-----BEGIN PRIVATE KEY-----
 * MIIJwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAM
 * ATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2VuY3J5cHQu
 * ...
 * he8Y4IWS6wY7bCkjCWDcRQJMEhg76fsO3txE+FiYruq9RUWhiF1myv4Q6W+CyBFC
 * 1qoJFlcDyqSMo5iHq3HLjs
 * -----END PRIVATE KEY-----
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * There are two ways to import WAF certificate state. * Using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/certificate:Certificate test <id>
 * ```
 *
 *  * Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Waf/certificate:Certificate test <id>/<enterprise_project_id>
 * ```
 *
 *  Note that the imported state is not identical to your resource definition, due to security reason. The missing attributes include `certificate`, and `private_key`. You can ignore changes as below. hcl resource "huaweicloud_waf_certificate" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  certificate, private_key
 *
 *  ]
 *
 *  } }
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Waf/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Specifies the certificate content.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Indicates the time when the certificate uploaded, in RFC3339 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the enterprise project ID of WAF certificate.
     * For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    /**
     * schema: Deprecated; The certificate expiration time.
     *
     * @deprecated Use 'expired_at' instead. 
     */
    public /*out*/ readonly expiration!: pulumi.Output<string>;
    /**
     * Indicates the time when the certificate expires, in RFC3339 format.
     */
    public /*out*/ readonly expiredAt!: pulumi.Output<string>;
    /**
     * Specifies the certificate name. The maximum length is `256` characters. Only digits,
     * letters, underscores(_), and hyphens(-) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the private key. This field does not support individual editing.
     * Changes to this field will only take effect when `certificate` changes.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the WAF certificate. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["expiration"] = state ? state.expiration : undefined;
            resourceInputs["expiredAt"] = state ? state.expiredAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["expiration"] = undefined /*out*/;
            resourceInputs["expiredAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Specifies the certificate content.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Indicates the time when the certificate uploaded, in RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID of WAF certificate.
     * For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * schema: Deprecated; The certificate expiration time.
     *
     * @deprecated Use 'expired_at' instead. 
     */
    expiration?: pulumi.Input<string>;
    /**
     * Indicates the time when the certificate expires, in RFC3339 format.
     */
    expiredAt?: pulumi.Input<string>;
    /**
     * Specifies the certificate name. The maximum length is `256` characters. Only digits,
     * letters, underscores(_), and hyphens(-) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the private key. This field does not support individual editing.
     * Changes to this field will only take effect when `certificate` changes.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the WAF certificate. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Specifies the certificate content.
     */
    certificate: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID of WAF certificate.
     * For enterprise users, if omitted, default enterprise project will be used.
     * Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the certificate name. The maximum length is `256` characters. Only digits,
     * letters, underscores(_), and hyphens(-) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the private key. This field does not support individual editing.
     * Changes to this field will only take effect when `certificate` changes.
     */
    privateKey: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the WAF certificate. If omitted, the
     * provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}
