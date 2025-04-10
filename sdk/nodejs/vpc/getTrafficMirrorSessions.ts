// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the traffic mirror sessions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const test1 = pulumi.output(huaweicloud.Vpc.getTrafficMirrorSessions({
 *     name: "mirror-session-a6b5",
 * }));
 * ```
 */
export function getTrafficMirrorSessions(args?: GetTrafficMirrorSessionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficMirrorSessionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Vpc/getTrafficMirrorSessions:getTrafficMirrorSessions", {
        "enabled": args.enabled,
        "name": args.name,
        "packetLength": args.packetLength,
        "priority": args.priority,
        "region": args.region,
        "trafficMirrorFilterId": args.trafficMirrorFilterId,
        "trafficMirrorSessionId": args.trafficMirrorSessionId,
        "trafficMirrorTargetId": args.trafficMirrorTargetId,
        "trafficMirrorTargetType": args.trafficMirrorTargetType,
        "type": args.type,
        "virtualNetworkId": args.virtualNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsArgs {
    /**
     * Specifies whether the mirror session is enabled. Defaults to **true**.
     */
    enabled?: string;
    /**
     * Specifies the traffic mirror session name used to query.
     */
    name?: string;
    /**
     * Specifies the maximum transmission unit (MTU).
     * The value range is **1-1460**, defaults to **96**.
     */
    packetLength?: string;
    /**
     * Specifies the mirror session priority. The value range is **1-32766**.
     * A smaller value indicates a higher priority.
     */
    priority?: string;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
    /**
     * Specifies the traffic mirror filter ID used in the session.
     */
    trafficMirrorFilterId?: string;
    /**
     * Specifies the traffic mirror session ID used to query.
     */
    trafficMirrorSessionId?: string;
    /**
     * Specifies the traffic mirror target ID.
     */
    trafficMirrorTargetId?: string;
    /**
     * Specifies the mirror target type. The value can be:
     * + **eni**: elastic network interface;
     * + **elb**: private network load balancer;
     */
    trafficMirrorTargetType?: string;
    /**
     * Specifies the mirror source type. The value can be **eni**(elastic network interface).
     */
    type?: string;
    /**
     * Specifies the VNI, which is used to distinguish mirrored traffic of different
     * sessions. The value range is **0-16777215**, defaults to **1**.
     */
    virtualNetworkId?: string;
}

/**
 * A collection of values returned by getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsResult {
    /**
     * Whether to enable a mirror session.
     */
    readonly enabled?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Traffic mirror session name.
     */
    readonly name?: string;
    /**
     * Maximum transmission unit (MTU).
     */
    readonly packetLength?: string;
    /**
     * Mirror session priority.
     */
    readonly priority?: string;
    readonly region: string;
    /**
     * Traffic mirror filter ID.
     */
    readonly trafficMirrorFilterId?: string;
    readonly trafficMirrorSessionId?: string;
    /**
     * List of traffic mirror sessions.
     */
    readonly trafficMirrorSessions: outputs.Vpc.GetTrafficMirrorSessionsTrafficMirrorSession[];
    /**
     * Mirror target ID.
     */
    readonly trafficMirrorTargetId?: string;
    /**
     * Mirror target type.
     */
    readonly trafficMirrorTargetType?: string;
    /**
     * Supported mirror source type.
     */
    readonly type?: string;
    /**
     * VNI, which is used to distinguish mirrored traffic of different sessions.
     */
    readonly virtualNetworkId?: string;
}

export function getTrafficMirrorSessionsOutput(args?: GetTrafficMirrorSessionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficMirrorSessionsResult> {
    return pulumi.output(args).apply(a => getTrafficMirrorSessions(a, opts))
}

/**
 * A collection of arguments for invoking getTrafficMirrorSessions.
 */
export interface GetTrafficMirrorSessionsOutputArgs {
    /**
     * Specifies whether the mirror session is enabled. Defaults to **true**.
     */
    enabled?: pulumi.Input<string>;
    /**
     * Specifies the traffic mirror session name used to query.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the maximum transmission unit (MTU).
     * The value range is **1-1460**, defaults to **96**.
     */
    packetLength?: pulumi.Input<string>;
    /**
     * Specifies the mirror session priority. The value range is **1-32766**.
     * A smaller value indicates a higher priority.
     */
    priority?: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the resource.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the traffic mirror filter ID used in the session.
     */
    trafficMirrorFilterId?: pulumi.Input<string>;
    /**
     * Specifies the traffic mirror session ID used to query.
     */
    trafficMirrorSessionId?: pulumi.Input<string>;
    /**
     * Specifies the traffic mirror target ID.
     */
    trafficMirrorTargetId?: pulumi.Input<string>;
    /**
     * Specifies the mirror target type. The value can be:
     * + **eni**: elastic network interface;
     * + **elb**: private network load balancer;
     */
    trafficMirrorTargetType?: pulumi.Input<string>;
    /**
     * Specifies the mirror source type. The value can be **eni**(elastic network interface).
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the VNI, which is used to distinguish mirrored traffic of different
     * sessions. The value range is **0-16777215**, defaults to **1**.
     */
    virtualNetworkId?: pulumi.Input<string>;
}
