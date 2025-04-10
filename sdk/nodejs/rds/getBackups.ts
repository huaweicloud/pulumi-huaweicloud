// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of RDS backups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const test = huaweicloud.Rds.getBackups({
 *     instanceId: instanceId,
 * });
 * ```
 */
export function getBackups(args: GetBackupsArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Rds/getBackups:getBackups", {
        "backupId": args.backupId,
        "backupType": args.backupType,
        "beginTime": args.beginTime,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackups.
 */
export interface GetBackupsArgs {
    /**
     * Backup ID.
     */
    backupId?: string;
    /**
     * Backup type.  
     * The options are as follows:
     * + **auto**: Automated full backup.
     * + **manual**: Manual full backup.
     * + **fragment**: Differential full backup.
     * + **incremental**: Automated incremental backup.
     */
    backupType?: string;
    /**
     * Start time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    beginTime?: string;
    /**
     * End time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    endTime?: string;
    /**
     * Instance ID.
     */
    instanceId: string;
    /**
     * Backup name.
     */
    name?: string;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getBackups.
 */
export interface GetBackupsResult {
    readonly backupId?: string;
    readonly backupType?: string;
    /**
     * Backup list. For details, see Data structure of the Backup field.
     * The backups structure is documented below.
     */
    readonly backups: outputs.Rds.GetBackupsBackup[];
    /**
     * Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    readonly beginTime?: string;
    /**
     * Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * RDS instance ID.
     */
    readonly instanceId: string;
    /**
     * Database to be backed up for Microsoft SQL Server.
     */
    readonly name?: string;
    readonly region: string;
}

export function getBackupsOutput(args: GetBackupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupsResult> {
    return pulumi.output(args).apply(a => getBackups(a, opts))
}

/**
 * A collection of arguments for invoking getBackups.
 */
export interface GetBackupsOutputArgs {
    /**
     * Backup ID.
     */
    backupId?: pulumi.Input<string>;
    /**
     * Backup type.  
     * The options are as follows:
     * + **auto**: Automated full backup.
     * + **manual**: Manual full backup.
     * + **fragment**: Differential full backup.
     * + **incremental**: Automated incremental backup.
     */
    backupType?: pulumi.Input<string>;
    /**
     * Start time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    beginTime?: pulumi.Input<string>;
    /**
     * End time in the "yyyy-mm-ddThh:mm:ssZ" format.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Backup name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
}
