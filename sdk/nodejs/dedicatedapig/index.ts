// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./api";
export * from "./apiPublishment";
export * from "./application";
export * from "./customAuthorizer";
export * from "./environment";
export * from "./getEnvironments";
export * from "./group";
export * from "./instance";
export * from "./response";
export * from "./throttlingPolicy";
export * from "./throttlingPolicyAssociate";
export * from "./vpcChannel";

// Import resources to register:
import { Api } from "./api";
import { ApiPublishment } from "./apiPublishment";
import { Application } from "./application";
import { CustomAuthorizer } from "./customAuthorizer";
import { Environment } from "./environment";
import { Group } from "./group";
import { Instance } from "./instance";
import { Response } from "./response";
import { ThrottlingPolicy } from "./throttlingPolicy";
import { ThrottlingPolicyAssociate } from "./throttlingPolicyAssociate";
import { VpcChannel } from "./vpcChannel";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:DedicatedApig/api:Api":
                return new Api(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/apiPublishment:ApiPublishment":
                return new ApiPublishment(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/customAuthorizer:CustomAuthorizer":
                return new CustomAuthorizer(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/environment:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/response:Response":
                return new Response(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/throttlingPolicy:ThrottlingPolicy":
                return new ThrottlingPolicy(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/throttlingPolicyAssociate:ThrottlingPolicyAssociate":
                return new ThrottlingPolicyAssociate(name, <any>undefined, { urn })
            case "huaweicloud:DedicatedApig/vpcChannel:VpcChannel":
                return new VpcChannel(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/api", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/apiPublishment", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/application", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/customAuthorizer", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/environment", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/group", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/instance", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/response", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/throttlingPolicy", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/throttlingPolicyAssociate", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "DedicatedApig/vpcChannel", _module)
