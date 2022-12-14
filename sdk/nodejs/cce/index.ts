// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./addon";
export * from "./cluster";
export * from "./getAddonTemplate";
export * from "./getCluster";
export * from "./getClusters";
export * from "./getNode";
export * from "./getNodePool";
export * from "./getNodes";
export * from "./namespace";
export * from "./node";
export * from "./nodeAttach";
export * from "./nodePool";
export * from "./pvc";

// Import resources to register:
import { Addon } from "./addon";
import { Cluster } from "./cluster";
import { Namespace } from "./namespace";
import { Node } from "./node";
import { NodeAttach } from "./nodeAttach";
import { NodePool } from "./nodePool";
import { Pvc } from "./pvc";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:Cce/addon:Addon":
                return new Addon(name, <any>undefined, { urn })
            case "huaweicloud:Cce/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "huaweicloud:Cce/namespace:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "huaweicloud:Cce/node:Node":
                return new Node(name, <any>undefined, { urn })
            case "huaweicloud:Cce/nodeAttach:NodeAttach":
                return new NodeAttach(name, <any>undefined, { urn })
            case "huaweicloud:Cce/nodePool:NodePool":
                return new NodePool(name, <any>undefined, { urn })
            case "huaweicloud:Cce/pvc:Pvc":
                return new Pvc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/addon", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/cluster", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/namespace", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/node", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/nodeAttach", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/nodePool", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Cce/pvc", _module)
