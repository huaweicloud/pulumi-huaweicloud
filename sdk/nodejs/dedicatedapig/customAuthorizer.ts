// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an APIG custom authorizer resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const authorizerName = config.requireObject("authorizerName");
 * const functionUrn = config.requireObject("functionUrn");
 * const test = new huaweicloud.dedicatedapig.CustomAuthorizer("test", {
 *     instanceId: instanceId,
 *     functionUrn: functionUrn,
 *     type: "FRONTEND",
 *     cacheAge: 60,
 *     identities: [{
 *         name: "user_name",
 *         location: "QUERY",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Custom Authorizers of the APIG can be imported using their `name` and related dedicated instance IDs, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer test <instance_id>/<name>
 * ```
 */
export class CustomAuthorizer extends pulumi.CustomResource {
    /**
     * Get an existing CustomAuthorizer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomAuthorizerState, opts?: pulumi.CustomResourceOptions): CustomAuthorizer {
        return new CustomAuthorizer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer';

    /**
     * Returns true if the given object is an instance of CustomAuthorizer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomAuthorizer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomAuthorizer.__pulumiType;
    }

    /**
     * Specifies the maximum cache age.
     */
    public readonly cacheAge!: pulumi.Output<number | undefined>;
    /**
     * The creation time of the custom authorizer.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the uniform function URN of the function graph resource.
     */
    public readonly functionUrn!: pulumi.Output<string>;
    /**
     * Specifies an array of one or more parameter identities of the custom authorizer.
     * The object structure is documented below.
     */
    public readonly identities!: pulumi.Output<outputs.DedicatedApig.CustomAuthorizerIdentity[] | undefined>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the
     * custom authorizer belongs to.
     * Changing this will create a new custom authorizer resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies whether to send the body.
     */
    public readonly isBodySend!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the parameter to be verified.
     * The parameter includes front-end and back-end parameters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the custom authorizer resource.
     * If omitted, the provider-level region will be used.
     * Changing this will create a new custom authorizer resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the custom authoriz type.
     * The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
     * Changing this will create a new custom authorizer resource.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Specifies the user data, which can contain a maximum of `2,048` characters.
     * The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
     */
    public readonly userData!: pulumi.Output<string | undefined>;

    /**
     * Create a CustomAuthorizer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomAuthorizerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomAuthorizerArgs | CustomAuthorizerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomAuthorizerState | undefined;
            resourceInputs["cacheAge"] = state ? state.cacheAge : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["functionUrn"] = state ? state.functionUrn : undefined;
            resourceInputs["identities"] = state ? state.identities : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isBodySend"] = state ? state.isBodySend : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
        } else {
            const args = argsOrState as CustomAuthorizerArgs | undefined;
            if ((!args || args.functionUrn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionUrn'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["cacheAge"] = args ? args.cacheAge : undefined;
            resourceInputs["functionUrn"] = args ? args.functionUrn : undefined;
            resourceInputs["identities"] = args ? args.identities : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isBodySend"] = args ? args.isBodySend : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomAuthorizer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomAuthorizer resources.
 */
export interface CustomAuthorizerState {
    /**
     * Specifies the maximum cache age.
     */
    cacheAge?: pulumi.Input<number>;
    /**
     * The creation time of the custom authorizer.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the uniform function URN of the function graph resource.
     */
    functionUrn?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more parameter identities of the custom authorizer.
     * The object structure is documented below.
     */
    identities?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.CustomAuthorizerIdentity>[]>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the
     * custom authorizer belongs to.
     * Changing this will create a new custom authorizer resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies whether to send the body.
     */
    isBodySend?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the parameter to be verified.
     * The parameter includes front-end and back-end parameters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the custom authorizer resource.
     * If omitted, the provider-level region will be used.
     * Changing this will create a new custom authorizer resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the custom authoriz type.
     * The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
     * Changing this will create a new custom authorizer resource.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the user data, which can contain a maximum of `2,048` characters.
     * The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
     */
    userData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomAuthorizer resource.
 */
export interface CustomAuthorizerArgs {
    /**
     * Specifies the maximum cache age.
     */
    cacheAge?: pulumi.Input<number>;
    /**
     * Specifies the uniform function URN of the function graph resource.
     */
    functionUrn: pulumi.Input<string>;
    /**
     * Specifies an array of one or more parameter identities of the custom authorizer.
     * The object structure is documented below.
     */
    identities?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.CustomAuthorizerIdentity>[]>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the
     * custom authorizer belongs to.
     * Changing this will create a new custom authorizer resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies whether to send the body.
     */
    isBodySend?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the parameter to be verified.
     * The parameter includes front-end and back-end parameters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the custom authorizer resource.
     * If omitted, the provider-level region will be used.
     * Changing this will create a new custom authorizer resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the custom authoriz type.
     * The valid values are **FRONTEND** and **BACKEND**. Defaults to **FRONTEND**.
     * Changing this will create a new custom authorizer resource.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the user data, which can contain a maximum of `2,048` characters.
     * The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
     */
    userData?: pulumi.Input<string>;
}