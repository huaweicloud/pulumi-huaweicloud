// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./desktop";
export * from "./service";
export * from "./user";

// Import resources to register:
import { Desktop } from "./desktop";
import { Service } from "./service";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:Workspace/desktop:Desktop":
                return new Desktop(name, <any>undefined, { urn })
            case "huaweicloud:Workspace/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "huaweicloud:Workspace/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "Workspace/desktop", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Workspace/service", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Workspace/user", _module)