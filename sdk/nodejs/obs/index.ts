// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./bucket";
export * from "./bucketObject";
export * from "./bucketPolicy";
export * from "./getBucketObject";
export * from "./getBuckets";

// Import resources to register:
import { Bucket } from "./bucket";
import { BucketObject } from "./bucketObject";
import { BucketPolicy } from "./bucketPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "huaweicloud:Obs/bucket:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "huaweicloud:Obs/bucketObject:BucketObject":
                return new BucketObject(name, <any>undefined, { urn })
            case "huaweicloud:Obs/bucketPolicy:BucketPolicy":
                return new BucketPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("huaweicloud", "Obs/bucket", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Obs/bucketObject", _module)
pulumi.runtime.registerResourceModule("huaweicloud", "Obs/bucketPolicy", _module)
