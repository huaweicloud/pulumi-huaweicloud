// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an APIG application resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const appName = config.requireObject("appName");
 * const appCode = config.requireObject("appCode");
 * const test = new huaweicloud.dedicatedapig.Application("test", {
 *     instanceId: instanceId,
 *     description: "Created by script",
 *     appCodes: [appCode],
 * });
 * ```
 *
 * ## Import
 *
 * APIG Applications can be imported using their `id` and ID of the APIG dedicated instance to which the application belongs, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/application:Application test <instance id>/<id>
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Specifies an array of one or more application codes which the APIG application belongs
     * to. Up to five application codes can be created. The code consists of 64 to 180 characters, starting with a letter,
     * digit, plus sign (+) or slash (/). Only letters, digits and following special special characters are allowed: !@#$%+-_
     * /=
     */
    public readonly appCodes!: pulumi.Output<string[] | undefined>;
    /**
     * App key.
     */
    public /*out*/ readonly appKey!: pulumi.Output<string>;
    /**
     * App secret.
     */
    public /*out*/ readonly appSecret!: pulumi.Output<string>;
    /**
     * Specifies the description about the APIG application. The description contain a
     * maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
     * Unicode format.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the APIG
     * application belongs to. Changing this will create a new APIG application resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the name of the API application. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits and underscores (_) are allowed. Chinese characters must be
     * in UTF-8 or Unicode format.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the APIG application resource. If
     * omitted, the provider-level region will be used. Changing this will create a new APIG application resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Registration time, in RFC-3339 format.
     */
    public /*out*/ readonly registraionTime!: pulumi.Output<string>;
    /**
     * Specifies the secret action to be done for the APIG application. The valid action
     * is *RESET*.
     */
    public readonly secretAction!: pulumi.Output<string | undefined>;
    /**
     * Time when the API group was last modified, in RFC-3339 format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["appCodes"] = state ? state.appCodes : undefined;
            resourceInputs["appKey"] = state ? state.appKey : undefined;
            resourceInputs["appSecret"] = state ? state.appSecret : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registraionTime"] = state ? state.registraionTime : undefined;
            resourceInputs["secretAction"] = state ? state.secretAction : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["appCodes"] = args ? args.appCodes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretAction"] = args ? args.secretAction : undefined;
            resourceInputs["appKey"] = undefined /*out*/;
            resourceInputs["appSecret"] = undefined /*out*/;
            resourceInputs["registraionTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * Specifies an array of one or more application codes which the APIG application belongs
     * to. Up to five application codes can be created. The code consists of 64 to 180 characters, starting with a letter,
     * digit, plus sign (+) or slash (/). Only letters, digits and following special special characters are allowed: !@#$%+-_
     * /=
     */
    appCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * App key.
     */
    appKey?: pulumi.Input<string>;
    /**
     * App secret.
     */
    appSecret?: pulumi.Input<string>;
    /**
     * Specifies the description about the APIG application. The description contain a
     * maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
     * Unicode format.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the APIG
     * application belongs to. Changing this will create a new APIG application resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the API application. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits and underscores (_) are allowed. Chinese characters must be
     * in UTF-8 or Unicode format.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the APIG application resource. If
     * omitted, the provider-level region will be used. Changing this will create a new APIG application resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Registration time, in RFC-3339 format.
     */
    registraionTime?: pulumi.Input<string>;
    /**
     * Specifies the secret action to be done for the APIG application. The valid action
     * is *RESET*.
     */
    secretAction?: pulumi.Input<string>;
    /**
     * Time when the API group was last modified, in RFC-3339 format.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * Specifies an array of one or more application codes which the APIG application belongs
     * to. Up to five application codes can be created. The code consists of 64 to 180 characters, starting with a letter,
     * digit, plus sign (+) or slash (/). Only letters, digits and following special special characters are allowed: !@#$%+-_
     * /=
     */
    appCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the description about the APIG application. The description contain a
     * maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
     * Unicode format.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the APIG
     * application belongs to. Changing this will create a new APIG application resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the name of the API application. The API group name consists of 3 to 64
     * characters, starting with a letter. Only letters, digits and underscores (_) are allowed. Chinese characters must be
     * in UTF-8 or Unicode format.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the APIG application resource. If
     * omitted, the provider-level region will be used. Changing this will create a new APIG application resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the secret action to be done for the APIG application. The valid action
     * is *RESET*.
     */
    secretAction?: pulumi.Input<string>;
}
