// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DMS kafka topic resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const kafkaInstanceId = config.requireObject("kafkaInstanceId");
 * const topic = new huaweicloud.dms.KafkaTopic("topic", {
 *     instanceId: kafkaInstanceId,
 *     partitions: 20,
 * });
 * ```
 *
 * ## Import
 *
 * DMS kafka topics can be imported using the kafka instance ID and topic name separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dms/kafkaTopic:KafkaTopic topic c8057fe5-23a8-46ef-ad83-c0055b4e0c5c/topic_1
 * ```
 */
export class KafkaTopic extends pulumi.CustomResource {
    /**
     * Get an existing KafkaTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaTopicState, opts?: pulumi.CustomResourceOptions): KafkaTopic {
        return new KafkaTopic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dms/kafkaTopic:KafkaTopic';

    /**
     * Returns true if the given object is an instance of KafkaTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaTopic.__pulumiType;
    }

    /**
     * Specifies the aging time in hours. The value ranges from 1 to 168 and defaults to 72.
     */
    public readonly agingTime!: pulumi.Output<number>;
    /**
     * Specifies the ID of the DMS kafka instance to which the topic belongs.
     * Changing this creates a new resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the name of the topic. The name starts with a letter, consists of 4 to
     * 64 characters, and supports only letters, digits, hyphens (-) and underscores (_). Changing this creates a new
     * resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the partition number. The value ranges from 1 to 100.
     */
    public readonly partitions!: pulumi.Output<number>;
    /**
     * The region in which to create the DMS kafka topic resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the replica number. The value ranges from 1 to 3 and defaults to 3.
     * Changing this creates a new resource.
     */
    public readonly replicas!: pulumi.Output<number>;
    /**
     * Whether or not to enable synchronous flushing.
     */
    public readonly syncFlushing!: pulumi.Output<boolean>;
    /**
     * Whether or not to enable synchronous replication.
     */
    public readonly syncReplication!: pulumi.Output<boolean>;

    /**
     * Create a KafkaTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaTopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaTopicArgs | KafkaTopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaTopicState | undefined;
            resourceInputs["agingTime"] = state ? state.agingTime : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partitions"] = state ? state.partitions : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicas"] = state ? state.replicas : undefined;
            resourceInputs["syncFlushing"] = state ? state.syncFlushing : undefined;
            resourceInputs["syncReplication"] = state ? state.syncReplication : undefined;
        } else {
            const args = argsOrState as KafkaTopicArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.partitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitions'");
            }
            resourceInputs["agingTime"] = args ? args.agingTime : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partitions"] = args ? args.partitions : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicas"] = args ? args.replicas : undefined;
            resourceInputs["syncFlushing"] = args ? args.syncFlushing : undefined;
            resourceInputs["syncReplication"] = args ? args.syncReplication : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KafkaTopic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaTopic resources.
 */
export interface KafkaTopicState {
    /**
     * Specifies the aging time in hours. The value ranges from 1 to 168 and defaults to 72.
     */
    agingTime?: pulumi.Input<number>;
    /**
     * Specifies the ID of the DMS kafka instance to which the topic belongs.
     * Changing this creates a new resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the topic. The name starts with a letter, consists of 4 to
     * 64 characters, and supports only letters, digits, hyphens (-) and underscores (_). Changing this creates a new
     * resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the partition number. The value ranges from 1 to 100.
     */
    partitions?: pulumi.Input<number>;
    /**
     * The region in which to create the DMS kafka topic resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the replica number. The value ranges from 1 to 3 and defaults to 3.
     * Changing this creates a new resource.
     */
    replicas?: pulumi.Input<number>;
    /**
     * Whether or not to enable synchronous flushing.
     */
    syncFlushing?: pulumi.Input<boolean>;
    /**
     * Whether or not to enable synchronous replication.
     */
    syncReplication?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a KafkaTopic resource.
 */
export interface KafkaTopicArgs {
    /**
     * Specifies the aging time in hours. The value ranges from 1 to 168 and defaults to 72.
     */
    agingTime?: pulumi.Input<number>;
    /**
     * Specifies the ID of the DMS kafka instance to which the topic belongs.
     * Changing this creates a new resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the name of the topic. The name starts with a letter, consists of 4 to
     * 64 characters, and supports only letters, digits, hyphens (-) and underscores (_). Changing this creates a new
     * resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the partition number. The value ranges from 1 to 100.
     */
    partitions: pulumi.Input<number>;
    /**
     * The region in which to create the DMS kafka topic resource. If omitted, the
     * provider-level region will be used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the replica number. The value ranges from 1 to 3 and defaults to 3.
     * Changing this creates a new resource.
     */
    replicas?: pulumi.Input<number>;
    /**
     * Whether or not to enable synchronous flushing.
     */
    syncFlushing?: pulumi.Input<boolean>;
    /**
     * Whether or not to enable synchronous replication.
     */
    syncReplication?: pulumi.Input<boolean>;
}