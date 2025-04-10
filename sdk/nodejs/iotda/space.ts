// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IoTDA resource space within HuaweiCloud.
 *
 * A resource space is the basic unit of service management and provides independent device management and platform
 * configuration capabilities at the service layer. Resources (such as products and devices) must be created on
 * a resource space.
 *
 * > The **basic** edition instance does not support updating the resource.
 *
 * > When accessing an IoTDA **standard** or **enterprise** edition instance, you need to specify
 *   the IoTDA service endpoint in `provider` block.
 *   You can login to the IoTDA console, choose the instance **Overview** and click **Access Details**
 *   to view the HTTPS application access address. An example of the access address might be
 *   *9bc34xxxxx.st1.iotda-app.ap-southeast-1.myhuaweicloud.com*, then you need to configure the
 *   `provider` block as follows:
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const name = config.requireObject("name");
 * const test = new huaweicloud.iotda.Space("test", {});
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:IoTDA/space:Space test <id>
 * ```
 */
export class Space extends pulumi.CustomResource {
    /**
     * Get an existing Space resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpaceState, opts?: pulumi.CustomResourceOptions): Space {
        return new Space(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:IoTDA/space:Space';

    /**
     * Returns true if the given object is an instance of Space.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Space {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Space.__pulumiType;
    }

    /**
     * Whether it is the default resource space. The IoT platform automatically creates and assigns
     * a default resource space (undeletable) to your account.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * Specifies the space name. The name contains a maximum of `64` characters.
     * Only letters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region in which to create the IoTDA resource space resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Space resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SpaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpaceArgs | SpaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpaceState | undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as SpaceArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["isDefault"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Space.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Space resources.
 */
export interface SpaceState {
    /**
     * Whether it is the default resource space. The IoT platform automatically creates and assigns
     * a default resource space (undeletable) to your account.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Specifies the space name. The name contains a maximum of `64` characters.
     * Only letters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA resource space resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Space resource.
 */
export interface SpaceArgs {
    /**
     * Specifies the space name. The name contains a maximum of `64` characters.
     * Only letters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the IoTDA resource space resource.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
}
