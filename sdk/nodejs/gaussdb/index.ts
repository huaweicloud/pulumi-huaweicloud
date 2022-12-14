// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getMysqlConfiguration";
export * from "./getMysqlDedicatedResource";
export * from "./getMysqlFlavors";
export * from "./getMysqlInstance";
export * from "./getMysqlInstances";
export * from "./mysqlInstance";
export * from "./mysqlProxy";

// Import resources to register:
import { MysqlInstance } from "./mysqlInstance";
import { MysqlProxy } from "./mysqlProxy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:GaussDB/mysqlInstance:MysqlInstance":
                return new MysqlInstance(name, <any>undefined, { urn })
            case "huaweicloud:GaussDB/mysqlProxy:MysqlProxy":
                return new MysqlProxy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "GaussDB/mysqlInstance", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "GaussDB/mysqlProxy", _module)
