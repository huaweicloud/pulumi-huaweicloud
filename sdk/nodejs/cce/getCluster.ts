// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about the cluster and obtains certificate for accessing cluster information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const clusterName = config.requireObject("clusterName");
 * const cluster = huaweicloud.Cce.getCluster({
 *     name: clusterName,
 *     status: "Available",
 * });
 * ```
 */
export function getCluster(args?: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Cce/getCluster:getCluster", {
        "clusterType": args.clusterType,
        "id": args.id,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * Specifies the type of the cluster. Possible values: **VirtualMachine**, **ARM64**.
     */
    clusterType?: string;
    /**
     * Specifies the ID of the cluster.
     */
    id?: string;
    /**
     * Specifies the name of the cluster.
     */
    name?: string;
    /**
     * Specifies the region in which to obtain the CCE cluster. If omitted, the provider-level
     * region will be used.
     */
    region?: string;
    /**
     * Specifies the status of the cluster.
     */
    status?: string;
    /**
     * Specifies the VPC ID to which the cluster belongs.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * Authentication mode of the cluster, possible values are x509 and rbac. Defaults to **rbac**.
     */
    readonly authenticationMode: string;
    /**
     * Charging mode of the cluster.
     */
    readonly billingMode: number;
    /**
     * The certificate clusters. Structure is documented below.
     */
    readonly certificateClusters: outputs.Cce.GetClusterCertificateCluster[];
    /**
     * The certificate users. Structure is documented below.
     */
    readonly certificateUsers: outputs.Cce.GetClusterCertificateUser[];
    readonly clusterType: string;
    /**
     * The version of cluster in string format.
     */
    readonly clusterVersion: string;
    /**
     * The container network segment.
     */
    readonly containerNetworkCidr: string;
    /**
     * The container network type: **overlay_l2** , **underlay_ipvlan**, **vpc-router** or **eni**.
     */
    readonly containerNetworkType: string;
    /**
     * Cluster description.
     */
    readonly description: string;
    /**
     * The access addresses of kube-apiserver in the cluster. Structure is documented below.
     */
    readonly endpoints: outputs.Cce.GetClusterEndpoint[];
    /**
     * ENI network segment. Specified when creating a CCE Turbo cluster.
     */
    readonly eniSubnetCidr: string;
    /**
     * The **IPv4 subnet ID** of the subnet where the ENI resides.
     * Specified when creating a CCE Turbo cluster.
     */
    readonly eniSubnetId: string;
    /**
     * The enterprise project ID of the CCE cluster.
     */
    readonly enterpriseProjectId: string;
    /**
     * The cluster specification in string format.
     */
    readonly flavorId: string;
    /**
     * The ID of the high speed network used to create bare metal nodes.
     */
    readonly highwaySubnetId: string;
    readonly id: string;
    /**
     * Raw Kubernetes config to be used by kubectl and other compatible tools.
     */
    readonly kubeConfigRaw: string;
    /**
     * Advanced configuration of master nodes. Structure is documented below.
     */
    readonly masters: outputs.Cce.GetClusterMaster[];
    /**
     * The user name.
     */
    readonly name: string;
    readonly region: string;
    /**
     * Security group ID of the cluster.
     */
    readonly securityGroupId: string;
    /**
     * The service network segment.
     */
    readonly serviceNetworkCidr: string;
    readonly status: string;
    /**
     * The ID of the subnet used to create the node.
     */
    readonly subnetId: string;
    readonly vpcId: string;
}

export function getClusterOutput(args?: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply(a => getCluster(a, opts))
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterOutputArgs {
    /**
     * Specifies the type of the cluster. Possible values: **VirtualMachine**, **ARM64**.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Specifies the ID of the cluster.
     */
    id?: pulumi.Input<string>;
    /**
     * Specifies the name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region in which to obtain the CCE cluster. If omitted, the provider-level
     * region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the status of the cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the VPC ID to which the cluster belongs.
     */
    vpcId?: pulumi.Input<string>;
}