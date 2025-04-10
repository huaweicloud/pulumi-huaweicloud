// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:ProjectMan/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Whether the project is archived.
     */
    public /*out*/ readonly archive!: pulumi.Output<number>;
    /**
     * The description about the project.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The enterprise project ID of the project.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * The project name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project code.
     */
    public /*out*/ readonly projectCode!: pulumi.Output<string>;
    /**
     * The number id of project.
     */
    public /*out*/ readonly projectNumId!: pulumi.Output<number>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The source of project.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The template id which used to create project.
     */
    public readonly templateId!: pulumi.Output<number>;
    /**
     * The type of project.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["archive"] = state ? state.archive : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectCode"] = state ? state.projectCode : undefined;
            resourceInputs["projectNumId"] = state ? state.projectNumId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["archive"] = undefined /*out*/;
            resourceInputs["projectCode"] = undefined /*out*/;
            resourceInputs["projectNumId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Whether the project is archived.
     */
    archive?: pulumi.Input<number>;
    /**
     * The description about the project.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project ID of the project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The project name.
     */
    name?: pulumi.Input<string>;
    /**
     * The project code.
     */
    projectCode?: pulumi.Input<string>;
    /**
     * The number id of project.
     */
    projectNumId?: pulumi.Input<number>;
    region?: pulumi.Input<string>;
    /**
     * The source of project.
     */
    source?: pulumi.Input<string>;
    /**
     * The template id which used to create project.
     */
    templateId?: pulumi.Input<number>;
    /**
     * The type of project.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The description about the project.
     */
    description?: pulumi.Input<string>;
    /**
     * The enterprise project ID of the project.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * The project name.
     */
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The source of project.
     */
    source?: pulumi.Input<string>;
    /**
     * The template id which used to create project.
     */
    templateId?: pulumi.Input<number>;
    /**
     * The type of project.
     */
    type: pulumi.Input<string>;
}
