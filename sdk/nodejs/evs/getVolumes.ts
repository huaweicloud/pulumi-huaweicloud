// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query the detailed information list of the EVS disks within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const targetServer = config.requireObject("targetServer");
 * const test = huaweicloud.Evs.getVolumes({
 *     serverId: targetServer,
 * });
 * ```
 */
export function getVolumes(args?: GetVolumesArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Evs/getVolumes:getVolumes", {
        "availabilityZone": args.availabilityZone,
        "enterpriseProjectId": args.enterpriseProjectId,
        "region": args.region,
        "serverId": args.serverId,
        "shareable": args.shareable,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolumes.
 */
export interface GetVolumesArgs {
    /**
     * Specifies the availability zone for the disks.
     */
    availabilityZone?: string;
    /**
     * Specifies the enterprise project ID for filtering.
     */
    enterpriseProjectId?: string;
    /**
     * Specifies the region in which to query the disk list.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
    /**
     * Specifies the server ID to which the disks are attached.
     */
    serverId?: string;
    /**
     * Specifies whether the disk is shareable.
     */
    shareable?: boolean;
    /**
     * Specifies the disk status. The valid values are as following:
     * + **FREEZED**
     * + **BIND_ERROR**
     * + **BINDING**
     * + **PENDING_DELETE**
     * + **PENDING_CREATE**
     * + **NOTIFYING**
     * + **NOTIFY_DELETE**
     * + **PENDING_UPDATE**
     * + **DOWN**
     * + **ACTIVE**
     * + **ELB**
     * + **ERROR**
     * + **VPN**
     */
    status?: string;
    /**
     * Specifies the included key/value pairs which associated with the desired disk.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVolumes.
 */
export interface GetVolumesResult {
    /**
     * The availability zone of the disk.
     */
    readonly availabilityZone?: string;
    /**
     * The ID of the enterprise project associated with the disk.
     */
    readonly enterpriseProjectId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region?: string;
    /**
     * The ID of the server to which the disk is attached.
     */
    readonly serverId?: string;
    /**
     * Whether the disk is shareable.
     */
    readonly shareable?: boolean;
    /**
     * The disk status.
     */
    readonly status?: string;
    /**
     * The disk tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The detailed information of the disks. Structure is documented below.
     */
    readonly volumes: outputs.Evs.GetVolumesVolume[];
}

export function getVolumesOutput(args?: GetVolumesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumesResult> {
    return pulumi.output(args).apply(a => getVolumes(a, opts))
}

/**
 * A collection of arguments for invoking getVolumes.
 */
export interface GetVolumesOutputArgs {
    /**
     * Specifies the availability zone for the disks.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the enterprise project ID for filtering.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the disk list.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the server ID to which the disks are attached.
     */
    serverId?: pulumi.Input<string>;
    /**
     * Specifies whether the disk is shareable.
     */
    shareable?: pulumi.Input<boolean>;
    /**
     * Specifies the disk status. The valid values are as following:
     * + **FREEZED**
     * + **BIND_ERROR**
     * + **BINDING**
     * + **PENDING_DELETE**
     * + **PENDING_CREATE**
     * + **NOTIFYING**
     * + **NOTIFY_DELETE**
     * + **PENDING_UPDATE**
     * + **DOWN**
     * + **ACTIVE**
     * + **ELB**
     * + **ERROR**
     * + **VPN**
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the included key/value pairs which associated with the desired disk.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}