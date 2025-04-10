// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * GaussDB OpenGauss instance management within HuaweiCould.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * OpenGaussDB instance can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:GaussDBforOpenGauss/opengaussInstance:OpengaussInstance test <id>
 * ```
 *
 *  Note that the imported state may not be identical to your resource definition, due to the attribute missing from the API response. The missing attributes include`password`, `ha.0.mode`, `ha.0.instance_mode`, `configuration_id`, `disk_encryption_id`, `enable_force_switch`, `enable_single_float_ip`, `parameters`, `period_unit`, `period` and `auto_renew`. It is generally recommended running `terraform plan` after importing a GaussDB OpenGauss instance. You can then decide if changes should be applied to the GaussDB OpenGauss instance, or the resource definition should be updated to align with the GaussDB OpenGauss instance. Also you can ignore changes as below. hcl resource "huaweicloud_gaussdb_opengauss_instance" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  password, configuration_id, disk_encryption_id, enable_force_switch, enable_single_float_ip, parameters, period_unit,
 *
 *  period, auto_renew,
 *
 *  ]
 *
 *  } }
 */
export class OpengaussInstance extends pulumi.CustomResource {
    /**
     * Get an existing OpengaussInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpengaussInstanceState, opts?: pulumi.CustomResourceOptions): OpengaussInstance {
        return new OpengaussInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:GaussDBforOpenGauss/opengaussInstance:OpengaussInstance';

    /**
     * Returns true if the given object is an instance of OpengaussInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpengaussInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpengaussInstance.__pulumiType;
    }

    /**
     * Specifies the advanced features.
     * The advanceFeatures structure is documented below.
     */
    public readonly advanceFeatures!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceAdvanceFeature[]>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are **true** and **false**. Defaults to **false**.
     */
    public readonly autoRenew!: pulumi.Output<string | undefined>;
    /**
     * Specifies the availability zone information, can be three same or
     * different az like **cn-north-4a,cn-north-4a,cn-north-4a**. Changing this parameter will create a new resource.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Specifies the advanced backup policy.
     * The backupStrategy structure is documented below.
     */
    public readonly backupStrategy!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceBackupStrategy>;
    /**
     * Indicates whether the host load is balanced due to a primary/standby switchover.
     */
    public /*out*/ readonly balanceStatus!: pulumi.Output<boolean>;
    /**
     * Specifies the charging mode of opengauss instance.
     * The valid values are as follows:
     * + **prePaid**: the yearly/monthly billing mode.
     * + **postPaid**: the pay-per-use billing mode.
     */
    public readonly chargingMode!: pulumi.Output<string>;
    /**
     * Specifies the parameter template ID.
     * Changing this parameter will create a new resource.
     */
    public readonly configurationId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the coordinator number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     * The value must not be greater than twice value of `shardingNum`.
     */
    public readonly coordinatorNum!: pulumi.Output<number | undefined>;
    /**
     * Specifies the datastore information.
     * The datastore structure is documented below.
     * Changing this parameter will create a new resource.
     */
    public readonly datastore!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceDatastore>;
    /**
     * Indicates the default username.
     */
    public /*out*/ readonly dbUserName!: pulumi.Output<string>;
    /**
     * Specifies the key ID for disk encryption.
     * Changing this parameter will create a new resource.
     */
    public readonly diskEncryptionId!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to forcibly promote a standby node to primary.
     * Defaults to **false**. Changing this parameter will create a new resource.
     */
    public readonly enableForceSwitch!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable single floating IP address policy,
     * which is only suitable for primary/standby instances. Value options:
     * + **true**: This function is enabled. Only one floating IP address is bound to the primary node of a DB instance. If a
     * primary/standby fail over occurs, the floating IP address does not change.
     * + **false (default value)**: The function is disabled. Each node is bound to a floating IP address. If a primary/standby
     * fail over occurs, the floating IP addresses change.
     */
    public readonly enableSingleFloatIp!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates the connection endpoints list of the DB instance. Example: [127.0.0.1:8000].
     */
    public /*out*/ readonly endpoints!: pulumi.Output<string[]>;
    /**
     * Specifies the enterprise project ID.
     */
    public readonly enterpriseProjectId!: pulumi.Output<string>;
    /**
     * Indicates whether error log collection is enabled. The value can be:
     * + **ON**: enabled
     * + **OFF**: disabled
     */
    public /*out*/ readonly errorLogSwitchStatus!: pulumi.Output<string>;
    /**
     * Specifies the instance specifications.
     */
    public readonly flavor!: pulumi.Output<string>;
    /**
     * Specifies whether to import the instance with the given configuration instead of
     * creation. If specified, try to import the instance instead of creation if the instance already existed.
     */
    public readonly forceImport!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the HA information.
     * The object structure is documented below.
     * Changing this parameter will create a new resource.
     */
    public readonly ha!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceHa>;
    /**
     * Indicates the maintenance window.
     */
    public /*out*/ readonly maintenanceWindow!: pulumi.Output<string>;
    /**
     * Specifies the port for MySQL compatibility. Value range: **0** or
     * **1024** to **39989**.
     * + The following ports are used by the system and cannot be used: **2378**, **2379**, **2380**, **2400**, **4999**,
     * **5000**, **5001**, **5100**, **5500**, **5999**, **6000**, **6001**, **6009**, **6010**, **6500**, **8015**, **8097**,
     * **8098**, **8181**, **9090**, **9100**, **9180**, **9187**, **9200**, **12016**, **12017**, **20049**, **20050**,
     * **21731**, **21732**, **32122**, **32123**, **32124**, **32125**, **32126**, **39001**,
     * **[Database port, Database port + 10]**.
     * + If the value is **0**, the MySQL compatibility port is disabled.
     */
    public readonly mysqlCompatibilityPort!: pulumi.Output<string>;
    /**
     * Specifies the name of the advance feature.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates the instance nodes information. Structure is documented below.
     */
    public /*out*/ readonly nodes!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceNode[]>;
    /**
     * Specifies an array of one or more parameters to be set to the instance after launched.
     * The parameters structure is documented below.
     */
    public readonly parameters!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceParameter[]>;
    /**
     * Specifies the database password. The value must be `8` to `32` characters in length,
     * including uppercase and lowercase letters, digits, and special characters, such as **~!@#%^*-_=+?**. You are advised
     * to enter a strong password to improve security, preventing security risks such as brute force cracking.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the charging period of opengauss instance.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 5.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Specifies the charging period unit of opengauss instance.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * Specifies the port information. Defaults to `8,000`.
     * The valid values are as follows:
     * + `2,378` to `2,380`
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Indicates the private IP address of the DB instance.
     */
    public /*out*/ readonly privateIps!: pulumi.Output<string[]>;
    /**
     * Indicates the public IP address of the DB instance.
     */
    public /*out*/ readonly publicIps!: pulumi.Output<string[]>;
    /**
     * Specifies the region in which to create the instance.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The replica number. The valid values are `2` and `3`, defaults to `3`.
     * Double replicas are only available for specific users and supports only instance versions are v1.3.0 or later.
     * Changing this parameter will create a new resource.
     */
    public readonly replicaNum!: pulumi.Output<number | undefined>;
    /**
     * Specifies the security group ID to which the instance belongs.
     * If the `port` parameter is specified, please ensure that the TCP ports in the inbound rule of security group
     * includes the `100` ports starting with the database port.
     * (For example, if the database port is `8,000`, the TCP port must include the range from `8,000` to `8,100`.)
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Specifies the sharding number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     */
    public readonly shardingNum!: pulumi.Output<number | undefined>;
    /**
     * Indicates the node status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the network ID of VPC subnet to which the instance belongs.
     * Changing this parameter will create a new resource.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Indicates the switch strategy.
     */
    public /*out*/ readonly switchStrategy!: pulumi.Output<string>;
    /**
     * Specifies the key/value pairs to associate with the GaussDB OpenGauss instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the time zone. Defaults to **UTC+08:00**.
     * Changing this parameter will create a new resource.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;
    /**
     * Specifies the volume type. Only **ULTRAHIGH** is supported now.
     * Changing this parameter will create a new resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Specifies the volume storage information.
     * The object structure is documented below.
     */
    public readonly volume!: pulumi.Output<outputs.GaussDBforOpenGauss.OpengaussInstanceVolume>;
    /**
     * Specifies the VPC ID to which the subnet belongs.
     * Changing this parameter will create a new resource.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a OpengaussInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpengaussInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpengaussInstanceArgs | OpengaussInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OpengaussInstanceState | undefined;
            resourceInputs["advanceFeatures"] = state ? state.advanceFeatures : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["backupStrategy"] = state ? state.backupStrategy : undefined;
            resourceInputs["balanceStatus"] = state ? state.balanceStatus : undefined;
            resourceInputs["chargingMode"] = state ? state.chargingMode : undefined;
            resourceInputs["configurationId"] = state ? state.configurationId : undefined;
            resourceInputs["coordinatorNum"] = state ? state.coordinatorNum : undefined;
            resourceInputs["datastore"] = state ? state.datastore : undefined;
            resourceInputs["dbUserName"] = state ? state.dbUserName : undefined;
            resourceInputs["diskEncryptionId"] = state ? state.diskEncryptionId : undefined;
            resourceInputs["enableForceSwitch"] = state ? state.enableForceSwitch : undefined;
            resourceInputs["enableSingleFloatIp"] = state ? state.enableSingleFloatIp : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["enterpriseProjectId"] = state ? state.enterpriseProjectId : undefined;
            resourceInputs["errorLogSwitchStatus"] = state ? state.errorLogSwitchStatus : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["forceImport"] = state ? state.forceImport : undefined;
            resourceInputs["ha"] = state ? state.ha : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["mysqlCompatibilityPort"] = state ? state.mysqlCompatibilityPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["publicIps"] = state ? state.publicIps : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicaNum"] = state ? state.replicaNum : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["shardingNum"] = state ? state.shardingNum : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["switchStrategy"] = state ? state.switchStrategy : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["volume"] = state ? state.volume : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as OpengaussInstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.flavor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavor'");
            }
            if ((!args || args.ha === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ha'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.volume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volume'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["advanceFeatures"] = args ? args.advanceFeatures : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["backupStrategy"] = args ? args.backupStrategy : undefined;
            resourceInputs["chargingMode"] = args ? args.chargingMode : undefined;
            resourceInputs["configurationId"] = args ? args.configurationId : undefined;
            resourceInputs["coordinatorNum"] = args ? args.coordinatorNum : undefined;
            resourceInputs["datastore"] = args ? args.datastore : undefined;
            resourceInputs["diskEncryptionId"] = args ? args.diskEncryptionId : undefined;
            resourceInputs["enableForceSwitch"] = args ? args.enableForceSwitch : undefined;
            resourceInputs["enableSingleFloatIp"] = args ? args.enableSingleFloatIp : undefined;
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["forceImport"] = args ? args.forceImport : undefined;
            resourceInputs["ha"] = args ? args.ha : undefined;
            resourceInputs["mysqlCompatibilityPort"] = args ? args.mysqlCompatibilityPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicaNum"] = args ? args.replicaNum : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["shardingNum"] = args ? args.shardingNum : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["volume"] = args ? args.volume : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["balanceStatus"] = undefined /*out*/;
            resourceInputs["dbUserName"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["errorLogSwitchStatus"] = undefined /*out*/;
            resourceInputs["maintenanceWindow"] = undefined /*out*/;
            resourceInputs["nodes"] = undefined /*out*/;
            resourceInputs["privateIps"] = undefined /*out*/;
            resourceInputs["publicIps"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["switchStrategy"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OpengaussInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpengaussInstance resources.
 */
export interface OpengaussInstanceState {
    /**
     * Specifies the advanced features.
     * The advanceFeatures structure is documented below.
     */
    advanceFeatures?: pulumi.Input<pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceAdvanceFeature>[]>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are **true** and **false**. Defaults to **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone information, can be three same or
     * different az like **cn-north-4a,cn-north-4a,cn-north-4a**. Changing this parameter will create a new resource.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the advanced backup policy.
     * The backupStrategy structure is documented below.
     */
    backupStrategy?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceBackupStrategy>;
    /**
     * Indicates whether the host load is balanced due to a primary/standby switchover.
     */
    balanceStatus?: pulumi.Input<boolean>;
    /**
     * Specifies the charging mode of opengauss instance.
     * The valid values are as follows:
     * + **prePaid**: the yearly/monthly billing mode.
     * + **postPaid**: the pay-per-use billing mode.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies the parameter template ID.
     * Changing this parameter will create a new resource.
     */
    configurationId?: pulumi.Input<string>;
    /**
     * Specifies the coordinator number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     * The value must not be greater than twice value of `shardingNum`.
     */
    coordinatorNum?: pulumi.Input<number>;
    /**
     * Specifies the datastore information.
     * The datastore structure is documented below.
     * Changing this parameter will create a new resource.
     */
    datastore?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceDatastore>;
    /**
     * Indicates the default username.
     */
    dbUserName?: pulumi.Input<string>;
    /**
     * Specifies the key ID for disk encryption.
     * Changing this parameter will create a new resource.
     */
    diskEncryptionId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly promote a standby node to primary.
     * Defaults to **false**. Changing this parameter will create a new resource.
     */
    enableForceSwitch?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable single floating IP address policy,
     * which is only suitable for primary/standby instances. Value options:
     * + **true**: This function is enabled. Only one floating IP address is bound to the primary node of a DB instance. If a
     * primary/standby fail over occurs, the floating IP address does not change.
     * + **false (default value)**: The function is disabled. Each node is bound to a floating IP address. If a primary/standby
     * fail over occurs, the floating IP addresses change.
     */
    enableSingleFloatIp?: pulumi.Input<boolean>;
    /**
     * Indicates the connection endpoints list of the DB instance. Example: [127.0.0.1:8000].
     */
    endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Indicates whether error log collection is enabled. The value can be:
     * + **ON**: enabled
     * + **OFF**: disabled
     */
    errorLogSwitchStatus?: pulumi.Input<string>;
    /**
     * Specifies the instance specifications.
     */
    flavor?: pulumi.Input<string>;
    /**
     * Specifies whether to import the instance with the given configuration instead of
     * creation. If specified, try to import the instance instead of creation if the instance already existed.
     */
    forceImport?: pulumi.Input<boolean>;
    /**
     * Specifies the HA information.
     * The object structure is documented below.
     * Changing this parameter will create a new resource.
     */
    ha?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceHa>;
    /**
     * Indicates the maintenance window.
     */
    maintenanceWindow?: pulumi.Input<string>;
    /**
     * Specifies the port for MySQL compatibility. Value range: **0** or
     * **1024** to **39989**.
     * + The following ports are used by the system and cannot be used: **2378**, **2379**, **2380**, **2400**, **4999**,
     * **5000**, **5001**, **5100**, **5500**, **5999**, **6000**, **6001**, **6009**, **6010**, **6500**, **8015**, **8097**,
     * **8098**, **8181**, **9090**, **9100**, **9180**, **9187**, **9200**, **12016**, **12017**, **20049**, **20050**,
     * **21731**, **21732**, **32122**, **32123**, **32124**, **32125**, **32126**, **39001**,
     * **[Database port, Database port + 10]**.
     * + If the value is **0**, the MySQL compatibility port is disabled.
     */
    mysqlCompatibilityPort?: pulumi.Input<string>;
    /**
     * Specifies the name of the advance feature.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates the instance nodes information. Structure is documented below.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceNode>[]>;
    /**
     * Specifies an array of one or more parameters to be set to the instance after launched.
     * The parameters structure is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceParameter>[]>;
    /**
     * Specifies the database password. The value must be `8` to `32` characters in length,
     * including uppercase and lowercase letters, digits, and special characters, such as **~!@#%^*-_=+?**. You are advised
     * to enter a strong password to improve security, preventing security risks such as brute force cracking.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the charging period of opengauss instance.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 5.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of opengauss instance.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * Specifies the port information. Defaults to `8,000`.
     * The valid values are as follows:
     * + `2,378` to `2,380`
     */
    port?: pulumi.Input<string>;
    /**
     * Indicates the private IP address of the DB instance.
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates the public IP address of the DB instance.
     */
    publicIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region in which to create the instance.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The replica number. The valid values are `2` and `3`, defaults to `3`.
     * Double replicas are only available for specific users and supports only instance versions are v1.3.0 or later.
     * Changing this parameter will create a new resource.
     */
    replicaNum?: pulumi.Input<number>;
    /**
     * Specifies the security group ID to which the instance belongs.
     * If the `port` parameter is specified, please ensure that the TCP ports in the inbound rule of security group
     * includes the `100` ports starting with the database port.
     * (For example, if the database port is `8,000`, the TCP port must include the range from `8,000` to `8,100`.)
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Specifies the sharding number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     */
    shardingNum?: pulumi.Input<number>;
    /**
     * Indicates the node status.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the network ID of VPC subnet to which the instance belongs.
     * Changing this parameter will create a new resource.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Indicates the switch strategy.
     */
    switchStrategy?: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the GaussDB OpenGauss instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the time zone. Defaults to **UTC+08:00**.
     * Changing this parameter will create a new resource.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Specifies the volume type. Only **ULTRAHIGH** is supported now.
     * Changing this parameter will create a new resource.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the volume storage information.
     * The object structure is documented below.
     */
    volume?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceVolume>;
    /**
     * Specifies the VPC ID to which the subnet belongs.
     * Changing this parameter will create a new resource.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpengaussInstance resource.
 */
export interface OpengaussInstanceArgs {
    /**
     * Specifies the advanced features.
     * The advanceFeatures structure is documented below.
     */
    advanceFeatures?: pulumi.Input<pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceAdvanceFeature>[]>;
    /**
     * Specifies whether auto renew is enabled.
     * Valid values are **true** and **false**. Defaults to **false**.
     */
    autoRenew?: pulumi.Input<string>;
    /**
     * Specifies the availability zone information, can be three same or
     * different az like **cn-north-4a,cn-north-4a,cn-north-4a**. Changing this parameter will create a new resource.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specifies the advanced backup policy.
     * The backupStrategy structure is documented below.
     */
    backupStrategy?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceBackupStrategy>;
    /**
     * Specifies the charging mode of opengauss instance.
     * The valid values are as follows:
     * + **prePaid**: the yearly/monthly billing mode.
     * + **postPaid**: the pay-per-use billing mode.
     */
    chargingMode?: pulumi.Input<string>;
    /**
     * Specifies the parameter template ID.
     * Changing this parameter will create a new resource.
     */
    configurationId?: pulumi.Input<string>;
    /**
     * Specifies the coordinator number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     * The value must not be greater than twice value of `shardingNum`.
     */
    coordinatorNum?: pulumi.Input<number>;
    /**
     * Specifies the datastore information.
     * The datastore structure is documented below.
     * Changing this parameter will create a new resource.
     */
    datastore?: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceDatastore>;
    /**
     * Specifies the key ID for disk encryption.
     * Changing this parameter will create a new resource.
     */
    diskEncryptionId?: pulumi.Input<string>;
    /**
     * Specifies whether to forcibly promote a standby node to primary.
     * Defaults to **false**. Changing this parameter will create a new resource.
     */
    enableForceSwitch?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable single floating IP address policy,
     * which is only suitable for primary/standby instances. Value options:
     * + **true**: This function is enabled. Only one floating IP address is bound to the primary node of a DB instance. If a
     * primary/standby fail over occurs, the floating IP address does not change.
     * + **false (default value)**: The function is disabled. Each node is bound to a floating IP address. If a primary/standby
     * fail over occurs, the floating IP addresses change.
     */
    enableSingleFloatIp?: pulumi.Input<boolean>;
    /**
     * Specifies the enterprise project ID.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the instance specifications.
     */
    flavor: pulumi.Input<string>;
    /**
     * Specifies whether to import the instance with the given configuration instead of
     * creation. If specified, try to import the instance instead of creation if the instance already existed.
     */
    forceImport?: pulumi.Input<boolean>;
    /**
     * Specifies the HA information.
     * The object structure is documented below.
     * Changing this parameter will create a new resource.
     */
    ha: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceHa>;
    /**
     * Specifies the port for MySQL compatibility. Value range: **0** or
     * **1024** to **39989**.
     * + The following ports are used by the system and cannot be used: **2378**, **2379**, **2380**, **2400**, **4999**,
     * **5000**, **5001**, **5100**, **5500**, **5999**, **6000**, **6001**, **6009**, **6010**, **6500**, **8015**, **8097**,
     * **8098**, **8181**, **9090**, **9100**, **9180**, **9187**, **9200**, **12016**, **12017**, **20049**, **20050**,
     * **21731**, **21732**, **32122**, **32123**, **32124**, **32125**, **32126**, **39001**,
     * **[Database port, Database port + 10]**.
     * + If the value is **0**, the MySQL compatibility port is disabled.
     */
    mysqlCompatibilityPort?: pulumi.Input<string>;
    /**
     * Specifies the name of the advance feature.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more parameters to be set to the instance after launched.
     * The parameters structure is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceParameter>[]>;
    /**
     * Specifies the database password. The value must be `8` to `32` characters in length,
     * including uppercase and lowercase letters, digits, and special characters, such as **~!@#%^*-_=+?**. You are advised
     * to enter a strong password to improve security, preventing security risks such as brute force cracking.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the charging period of opengauss instance.
     * If `periodUnit` is set to **month**, the value ranges from 1 to 9.
     * If `periodUnit` is set to **year**, the value ranges from 1 to 5.
     * This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    period?: pulumi.Input<number>;
    /**
     * Specifies the charging period unit of opengauss instance.
     * Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
     * Changing this parameter will create a new resource.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * Specifies the port information. Defaults to `8,000`.
     * The valid values are as follows:
     * + `2,378` to `2,380`
     */
    port?: pulumi.Input<string>;
    /**
     * Specifies the region in which to create the instance.
     * If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The replica number. The valid values are `2` and `3`, defaults to `3`.
     * Double replicas are only available for specific users and supports only instance versions are v1.3.0 or later.
     * Changing this parameter will create a new resource.
     */
    replicaNum?: pulumi.Input<number>;
    /**
     * Specifies the security group ID to which the instance belongs.
     * If the `port` parameter is specified, please ensure that the TCP ports in the inbound rule of security group
     * includes the `100` ports starting with the database port.
     * (For example, if the database port is `8,000`, the TCP port must include the range from `8,000` to `8,100`.)
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Specifies the sharding number.  
     * The valid value is range form `1` to `9`. The default value is `3`.
     */
    shardingNum?: pulumi.Input<number>;
    /**
     * Specifies the network ID of VPC subnet to which the instance belongs.
     * Changing this parameter will create a new resource.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Specifies the key/value pairs to associate with the GaussDB OpenGauss instance.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the time zone. Defaults to **UTC+08:00**.
     * Changing this parameter will create a new resource.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * Specifies the volume storage information.
     * The object structure is documented below.
     */
    volume: pulumi.Input<inputs.GaussDBforOpenGauss.OpengaussInstanceVolume>;
    /**
     * Specifies the VPC ID to which the subnet belongs.
     * Changing this parameter will create a new resource.
     */
    vpcId: pulumi.Input<string>;
}
