// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages DIS Stream resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a stream that type is BLOB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const stream = new huaweicloud.Dis.Stream("stream", {
 *     partitionCount: 1,
 *     streamName: "terraform_test_dis_stream",
 * });
 * ```
 * ### Create a stream that type is JSON
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const stream = new huaweicloud.Dis.Stream("stream", {
 *     dataSchema: "{\"type\":\"record\",\"name\":\"RecordName\",\"fields\":[{\"name\":\"id\",\"type\":\"string\",\"doc\":\"Type inferred from '\\\"2017/10/11 11:11:11\\\"'\"},{\"name\":\"info\",\"type\":{\"type\":\"array\",\"items\":{\"type\":\"record\",\"name\":\"info\",\"fields\":[{\"name\":\"date\",\"type\":\"string\",\"doc\":\"Type inferred from '\\\"2018/10/11 11:11:11\\\"'\"}]}},\"doc\":\"Type inferred from '[{\\\"date\\\":\\\"2018/10/11 11:11:11\\\"}]'\"}]}",
 *     dataType: "JSON",
 *     partitionCount: 1,
 *     streamName: "terraform_test_dis_stream",
 * });
 * ```
 *
 * ## Import
 *
 * Dis stream can be imported by `stream_name`. For example, bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:Dis/stream:Stream example _abc123
 * ```
 */
export class Stream extends pulumi.CustomResource {
    /**
     * Get an existing Stream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamState, opts?: pulumi.CustomResourceOptions): Stream {
        return new Stream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Dis/stream:Stream';

    /**
     * Returns true if the given object is an instance of Stream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stream.__pulumiType;
    }

    /**
     * Maximum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    public readonly autoScaleMaxPartitionCount!: pulumi.Output<number>;
    /**
     * Minimum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    public readonly autoScaleMinPartitionCount!: pulumi.Output<number>;
    /**
     * Data compression type. The value is one of snappy, gzip and zip.
     * Changing this parameter will create a new resource.
     */
    public readonly compressionFormat!: pulumi.Output<string>;
    /**
     * Timestamp at which the DIS stream was created.
     */
    public /*out*/ readonly created!: pulumi.Output<number>;
    /**
     * Field separator for CSV file. Changing this parameter will create a new
     * resource.
     */
    public readonly csvDelimiter!: pulumi.Output<string>;
    /**
     * User's JSON, CSV format data schema, described with Avro schema. Changing
     * this parameter will create a new resource.
     */
    public readonly dataSchema!: pulumi.Output<string>;
    /**
     * Data type of the data putting into the stream. The value is one of **BLOB**,
     * **JSON** and **CSV**. Changing this parameter will create a new resource.
     */
    public readonly dataType!: pulumi.Output<string>;
    /**
     * Specifies the enterprise project id of the dis stream, Value 0
     * indicates the default enterprise project. Changing this parameter will create a new resource.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Number of the expect partitions. NOTE: Each stream can be scaled up and down a
     * total of five times within one hour. After the stream is successfully scaled up or down, it cannot be scaled up or
     * down again within the next one hour.
     */
    public readonly partitionCount!: pulumi.Output<number>;
    /**
     * The information of stream partitions. Structure is documented below.
     */
    public /*out*/ readonly partitions!: pulumi.Output<outputs.Dis.StreamPartition[]>;
    /**
     * Total number of readable partitions (including partitions in ACTIVE state only).
     */
    public /*out*/ readonly readablePartitionCount!: pulumi.Output<number>;
    /**
     * The region in which to create the DIS stream resource. If omitted, the
     * provider-level region will be used. Changing this creates a new DIS Stream resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The number of hours for which data from the stream will be retained in DIS.
     * Value range: `24` to `72`. Unit: **hour**. Default:`24`. Changing this parameter will create a new resource.
     */
    public readonly retentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The status of the partition.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Indicates a stream ID in UUID format.
     */
    public /*out*/ readonly streamId!: pulumi.Output<string>;
    /**
     * Name of the DIS stream to be created.
     * Changing this parameter will create a new resource.
     */
    public readonly streamName!: pulumi.Output<string>;
    /**
     * Stream Type. The value is COMMON(means 1M bandwidth) or ADVANCED(means 5M
     * bandwidth). Changing this parameter will create a new resource.
     */
    public readonly streamType!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the stream.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Total number of writable partitions (including partitions in ACTIVE and DELETED states).
     */
    public /*out*/ readonly writablePartitionCount!: pulumi.Output<number>;

    /**
     * Create a Stream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamArgs | StreamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamState | undefined;
            resourceInputs["autoScaleMaxPartitionCount"] = state ? state.autoScaleMaxPartitionCount : undefined;
            resourceInputs["autoScaleMinPartitionCount"] = state ? state.autoScaleMinPartitionCount : undefined;
            resourceInputs["compressionFormat"] = state ? state.compressionFormat : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["csvDelimiter"] = state ? state.csvDelimiter : undefined;
            resourceInputs["dataSchema"] = state ? state.dataSchema : undefined;
            resourceInputs["dataType"] = state ? state.dataType : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["partitionCount"] = state ? state.partitionCount : undefined;
            resourceInputs["partitions"] = state ? state.partitions : undefined;
            resourceInputs["readablePartitionCount"] = state ? state.readablePartitionCount : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["streamId"] = state ? state.streamId : undefined;
            resourceInputs["streamName"] = state ? state.streamName : undefined;
            resourceInputs["streamType"] = state ? state.streamType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["writablePartitionCount"] = state ? state.writablePartitionCount : undefined;
        } else {
            const args = argsOrState as StreamArgs | undefined;
            if ((!args || args.partitionCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionCount'");
            }
            if ((!args || args.streamName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamName'");
            }
            resourceInputs["autoScaleMaxPartitionCount"] = args ? args.autoScaleMaxPartitionCount : undefined;
            resourceInputs["autoScaleMinPartitionCount"] = args ? args.autoScaleMinPartitionCount : undefined;
            resourceInputs["compressionFormat"] = args ? args.compressionFormat : undefined;
            resourceInputs["csvDelimiter"] = args ? args.csvDelimiter : undefined;
            resourceInputs["dataSchema"] = args ? args.dataSchema : undefined;
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["partitionCount"] = args ? args.partitionCount : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["streamName"] = args ? args.streamName : undefined;
            resourceInputs["streamType"] = args ? args.streamType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["partitions"] = undefined /*out*/;
            resourceInputs["readablePartitionCount"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["streamId"] = undefined /*out*/;
            resourceInputs["writablePartitionCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stream resources.
 */
export interface StreamState {
    /**
     * Maximum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    autoScaleMaxPartitionCount?: pulumi.Input<number>;
    /**
     * Minimum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    autoScaleMinPartitionCount?: pulumi.Input<number>;
    /**
     * Data compression type. The value is one of snappy, gzip and zip.
     * Changing this parameter will create a new resource.
     */
    compressionFormat?: pulumi.Input<string>;
    /**
     * Timestamp at which the DIS stream was created.
     */
    created?: pulumi.Input<number>;
    /**
     * Field separator for CSV file. Changing this parameter will create a new
     * resource.
     */
    csvDelimiter?: pulumi.Input<string>;
    /**
     * User's JSON, CSV format data schema, described with Avro schema. Changing
     * this parameter will create a new resource.
     */
    dataSchema?: pulumi.Input<string>;
    /**
     * Data type of the data putting into the stream. The value is one of **BLOB**,
     * **JSON** and **CSV**. Changing this parameter will create a new resource.
     */
    dataType?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id of the dis stream, Value 0
     * indicates the default enterprise project. Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Number of the expect partitions. NOTE: Each stream can be scaled up and down a
     * total of five times within one hour. After the stream is successfully scaled up or down, it cannot be scaled up or
     * down again within the next one hour.
     */
    partitionCount?: pulumi.Input<number>;
    /**
     * The information of stream partitions. Structure is documented below.
     */
    partitions?: pulumi.Input<pulumi.Input<inputs.Dis.StreamPartition>[]>;
    /**
     * Total number of readable partitions (including partitions in ACTIVE state only).
     */
    readablePartitionCount?: pulumi.Input<number>;
    /**
     * The region in which to create the DIS stream resource. If omitted, the
     * provider-level region will be used. Changing this creates a new DIS Stream resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The number of hours for which data from the stream will be retained in DIS.
     * Value range: `24` to `72`. Unit: **hour**. Default:`24`. Changing this parameter will create a new resource.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * The status of the partition.
     */
    status?: pulumi.Input<string>;
    /**
     * Indicates a stream ID in UUID format.
     */
    streamId?: pulumi.Input<string>;
    /**
     * Name of the DIS stream to be created.
     * Changing this parameter will create a new resource.
     */
    streamName?: pulumi.Input<string>;
    /**
     * Stream Type. The value is COMMON(means 1M bandwidth) or ADVANCED(means 5M
     * bandwidth). Changing this parameter will create a new resource.
     */
    streamType?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the stream.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Total number of writable partitions (including partitions in ACTIVE and DELETED states).
     */
    writablePartitionCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Stream resource.
 */
export interface StreamArgs {
    /**
     * Maximum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    autoScaleMaxPartitionCount?: pulumi.Input<number>;
    /**
     * Minimum number of partition for automatic scaling.
     * Changing this parameter will create a new resource.
     */
    autoScaleMinPartitionCount?: pulumi.Input<number>;
    /**
     * Data compression type. The value is one of snappy, gzip and zip.
     * Changing this parameter will create a new resource.
     */
    compressionFormat?: pulumi.Input<string>;
    /**
     * Field separator for CSV file. Changing this parameter will create a new
     * resource.
     */
    csvDelimiter?: pulumi.Input<string>;
    /**
     * User's JSON, CSV format data schema, described with Avro schema. Changing
     * this parameter will create a new resource.
     */
    dataSchema?: pulumi.Input<string>;
    /**
     * Data type of the data putting into the stream. The value is one of **BLOB**,
     * **JSON** and **CSV**. Changing this parameter will create a new resource.
     */
    dataType?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project id of the dis stream, Value 0
     * indicates the default enterprise project. Changing this parameter will create a new resource.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Number of the expect partitions. NOTE: Each stream can be scaled up and down a
     * total of five times within one hour. After the stream is successfully scaled up or down, it cannot be scaled up or
     * down again within the next one hour.
     */
    partitionCount: pulumi.Input<number>;
    /**
     * The region in which to create the DIS stream resource. If omitted, the
     * provider-level region will be used. Changing this creates a new DIS Stream resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The number of hours for which data from the stream will be retained in DIS.
     * Value range: `24` to `72`. Unit: **hour**. Default:`24`. Changing this parameter will create a new resource.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * Name of the DIS stream to be created.
     * Changing this parameter will create a new resource.
     */
    streamName: pulumi.Input<string>;
    /**
     * Stream Type. The value is COMMON(means 1M bandwidth) or ADVANCED(means 5M
     * bandwidth). Changing this parameter will create a new resource.
     */
    streamType?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the stream.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
