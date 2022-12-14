// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getTopics";
export * from "./subscription";
export * from "./topic";

// Import resources to register:
import { Subscription } from "./subscription";
import { Topic } from "./topic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:Smn/subscription:Subscription":
                return new Subscription(name, <any>undefined, { urn })
            case "huaweicloud:Smn/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "Smn/subscription", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Smn/topic", _module)
