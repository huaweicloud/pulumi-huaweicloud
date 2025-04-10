// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of RDS storage types.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.require("instanceId");
 *
 * const test = pulumi.output(huaweicloud.Rds.getStorageTypes({
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 * }));
 * ```
 */
export function getStorageTypes(args: GetStorageTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageTypesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Rds/getStorageTypes:getStorageTypes", {
        "dbType": args.dbType,
        "dbVersion": args.dbVersion,
        "instanceMode": args.instanceMode,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageTypes.
 */
export interface GetStorageTypesArgs {
    /**
     * DB engine. The valid values are **MySQL**, **PostgreSQL**, **SQLServer**, **MariaDB**.
     */
    dbType: string;
    /**
     * DB version number.
     */
    dbVersion: string;
    /**
     * HA mode. The valid values are **single**, **ha**, **replica**.
     */
    instanceMode?: string;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
}

/**
 * A collection of values returned by getStorageTypes.
 */
export interface GetStorageTypesResult {
    readonly dbType: string;
    readonly dbVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceMode?: string;
    readonly region: string;
    /**
     * Storage type list. For details, see Data structure of the storageType field.
     * The storageType structure is documented below.
     */
    readonly storageTypes: outputs.Rds.GetStorageTypesStorageType[];
}

export function getStorageTypesOutput(args: GetStorageTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStorageTypesResult> {
    return pulumi.output(args).apply(a => getStorageTypes(a, opts))
}

/**
 * A collection of arguments for invoking getStorageTypes.
 */
export interface GetStorageTypesOutputArgs {
    /**
     * DB engine. The valid values are **MySQL**, **PostgreSQL**, **SQLServer**, **MariaDB**.
     */
    dbType: pulumi.Input<string>;
    /**
     * DB version number.
     */
    dbVersion: pulumi.Input<string>;
    /**
     * HA mode. The valid values are **single**, **ha**, **replica**.
     */
    instanceMode?: pulumi.Input<string>;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
}
