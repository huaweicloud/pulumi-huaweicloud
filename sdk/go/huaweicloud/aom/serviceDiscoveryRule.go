// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aom

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AOM service discovery rule resource within HuaweiCloud.
//
// ## Example Usage
// ### Basic example
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Aom"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Aom"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Aom.NewServiceDiscoveryRule(ctx, "discoveryRule", &Aom.ServiceDiscoveryRuleArgs{
//				DetectLogEnabled:     pulumi.Bool(true),
//				DiscoveryRuleEnabled: pulumi.Bool(true),
//				DiscoveryRules: aom.ServiceDiscoveryRuleDiscoveryRuleArray{
//					&aom.ServiceDiscoveryRuleDiscoveryRuleArgs{
//						CheckContents: pulumi.StringArray{
//							pulumi.String("java"),
//						},
//						CheckMode: pulumi.String("contain"),
//						CheckType: pulumi.String("cmdLine"),
//					},
//				},
//				IsDefaultRule: pulumi.Bool(false),
//				LogFileSuffixes: pulumi.StringArray{
//					pulumi.String("log"),
//				},
//				LogPathRules: aom.ServiceDiscoveryRuleLogPathRuleArray{
//					&aom.ServiceDiscoveryRuleLogPathRuleArgs{
//						Args: pulumi.StringArray{
//							pulumi.String("java"),
//						},
//						NameType: pulumi.String("cmdLineHash"),
//						Values: pulumi.StringArray{
//							pulumi.String("/tmp/log"),
//						},
//					},
//				},
//				NameRules: &aom.ServiceDiscoveryRuleNameRulesArgs{
//					ApplicationNameRules: aom.ServiceDiscoveryRuleNameRulesApplicationNameRuleArray{
//						&aom.ServiceDiscoveryRuleNameRulesApplicationNameRuleArgs{
//							Args: pulumi.StringArray{
//								pulumi.String("java"),
//							},
//							NameType: pulumi.String("str"),
//						},
//					},
//					ServiceNameRules: aom.ServiceDiscoveryRuleNameRulesServiceNameRuleArray{
//						&aom.ServiceDiscoveryRuleNameRulesServiceNameRuleArgs{
//							Args: pulumi.StringArray{
//								pulumi.String("java"),
//							},
//							NameType: pulumi.String("str"),
//						},
//					},
//				},
//				Priority:    pulumi.Int(9999),
//				ServiceType: pulumi.String("Java"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// AOM service discovery rules can be imported using the `name`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Aom/serviceDiscoveryRule:ServiceDiscoveryRule alarm_rule <name>
//
// ```
type ServiceDiscoveryRule struct {
	pulumi.CustomResourceState

	// The rule create time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the rule description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to enable log collection. The default value is true.
	DetectLogEnabled pulumi.BoolPtrOutput `pulumi:"detectLogEnabled"`
	// Specifies whether the rule is enabled. The default value is true.
	DiscoveryRuleEnabled pulumi.BoolPtrOutput `pulumi:"discoveryRuleEnabled"`
	// Specifies the discovery rule. If the array contains multiple conditions, only the
	// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
	// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
	// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
	// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
	// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
	// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
	DiscoveryRules ServiceDiscoveryRuleDiscoveryRuleArrayOutput `pulumi:"discoveryRules"`
	// Specifies whether the rule is the default one. The default value is false.
	IsDefaultRule pulumi.BoolPtrOutput `pulumi:"isDefaultRule"`
	// Specifies the log file suffix. This is a list of strings.
	// The values can be: **log**, **trace**, and **out**.
	LogFileSuffixes pulumi.StringArrayOutput `pulumi:"logFileSuffixes"`
	// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
	// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
	// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
	LogPathRules ServiceDiscoveryRuleLogPathRuleArrayOutput `pulumi:"logPathRules"`
	// Specifies the rule name, which contains `4` to `63` characters. It must start
	// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the naming rules for discovered services and applications.
	// The object structure is documented below.
	NameRules ServiceDiscoveryRuleNameRulesOutput `pulumi:"nameRules"`
	// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The region in which to create the service discovery rule resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The rule ID in uuid format.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Specifies the service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
}

// NewServiceDiscoveryRule registers a new resource with the given unique name, arguments, and options.
func NewServiceDiscoveryRule(ctx *pulumi.Context,
	name string, args *ServiceDiscoveryRuleArgs, opts ...pulumi.ResourceOption) (*ServiceDiscoveryRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiscoveryRules == nil {
		return nil, errors.New("invalid value for required argument 'DiscoveryRules'")
	}
	if args.LogFileSuffixes == nil {
		return nil, errors.New("invalid value for required argument 'LogFileSuffixes'")
	}
	if args.NameRules == nil {
		return nil, errors.New("invalid value for required argument 'NameRules'")
	}
	if args.ServiceType == nil {
		return nil, errors.New("invalid value for required argument 'ServiceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServiceDiscoveryRule
	err := ctx.RegisterResource("huaweicloud:Aom/serviceDiscoveryRule:ServiceDiscoveryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceDiscoveryRule gets an existing ServiceDiscoveryRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceDiscoveryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceDiscoveryRuleState, opts ...pulumi.ResourceOption) (*ServiceDiscoveryRule, error) {
	var resource ServiceDiscoveryRule
	err := ctx.ReadResource("huaweicloud:Aom/serviceDiscoveryRule:ServiceDiscoveryRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceDiscoveryRule resources.
type serviceDiscoveryRuleState struct {
	// The rule create time.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the rule description.
	Description *string `pulumi:"description"`
	// Specifies whether to enable log collection. The default value is true.
	DetectLogEnabled *bool `pulumi:"detectLogEnabled"`
	// Specifies whether the rule is enabled. The default value is true.
	DiscoveryRuleEnabled *bool `pulumi:"discoveryRuleEnabled"`
	// Specifies the discovery rule. If the array contains multiple conditions, only the
	// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
	// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
	// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
	// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
	// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
	// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
	DiscoveryRules []ServiceDiscoveryRuleDiscoveryRule `pulumi:"discoveryRules"`
	// Specifies whether the rule is the default one. The default value is false.
	IsDefaultRule *bool `pulumi:"isDefaultRule"`
	// Specifies the log file suffix. This is a list of strings.
	// The values can be: **log**, **trace**, and **out**.
	LogFileSuffixes []string `pulumi:"logFileSuffixes"`
	// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
	// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
	// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
	LogPathRules []ServiceDiscoveryRuleLogPathRule `pulumi:"logPathRules"`
	// Specifies the rule name, which contains `4` to `63` characters. It must start
	// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
	Name *string `pulumi:"name"`
	// Specifies the naming rules for discovered services and applications.
	// The object structure is documented below.
	NameRules *ServiceDiscoveryRuleNameRules `pulumi:"nameRules"`
	// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
	Priority *int `pulumi:"priority"`
	// The region in which to create the service discovery rule resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// The rule ID in uuid format.
	RuleId *string `pulumi:"ruleId"`
	// Specifies the service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
	ServiceType *string `pulumi:"serviceType"`
}

type ServiceDiscoveryRuleState struct {
	// The rule create time.
	CreatedAt pulumi.StringPtrInput
	// Specifies the rule description.
	Description pulumi.StringPtrInput
	// Specifies whether to enable log collection. The default value is true.
	DetectLogEnabled pulumi.BoolPtrInput
	// Specifies whether the rule is enabled. The default value is true.
	DiscoveryRuleEnabled pulumi.BoolPtrInput
	// Specifies the discovery rule. If the array contains multiple conditions, only the
	// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
	// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
	// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
	// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
	// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
	// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
	DiscoveryRules ServiceDiscoveryRuleDiscoveryRuleArrayInput
	// Specifies whether the rule is the default one. The default value is false.
	IsDefaultRule pulumi.BoolPtrInput
	// Specifies the log file suffix. This is a list of strings.
	// The values can be: **log**, **trace**, and **out**.
	LogFileSuffixes pulumi.StringArrayInput
	// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
	// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
	// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
	LogPathRules ServiceDiscoveryRuleLogPathRuleArrayInput
	// Specifies the rule name, which contains `4` to `63` characters. It must start
	// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
	Name pulumi.StringPtrInput
	// Specifies the naming rules for discovered services and applications.
	// The object structure is documented below.
	NameRules ServiceDiscoveryRuleNameRulesPtrInput
	// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
	Priority pulumi.IntPtrInput
	// The region in which to create the service discovery rule resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// The rule ID in uuid format.
	RuleId pulumi.StringPtrInput
	// Specifies the service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
	ServiceType pulumi.StringPtrInput
}

func (ServiceDiscoveryRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDiscoveryRuleState)(nil)).Elem()
}

type serviceDiscoveryRuleArgs struct {
	// Specifies the rule description.
	Description *string `pulumi:"description"`
	// Specifies whether to enable log collection. The default value is true.
	DetectLogEnabled *bool `pulumi:"detectLogEnabled"`
	// Specifies whether the rule is enabled. The default value is true.
	DiscoveryRuleEnabled *bool `pulumi:"discoveryRuleEnabled"`
	// Specifies the discovery rule. If the array contains multiple conditions, only the
	// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
	// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
	// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
	// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
	// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
	// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
	DiscoveryRules []ServiceDiscoveryRuleDiscoveryRule `pulumi:"discoveryRules"`
	// Specifies whether the rule is the default one. The default value is false.
	IsDefaultRule *bool `pulumi:"isDefaultRule"`
	// Specifies the log file suffix. This is a list of strings.
	// The values can be: **log**, **trace**, and **out**.
	LogFileSuffixes []string `pulumi:"logFileSuffixes"`
	// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
	// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
	// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
	LogPathRules []ServiceDiscoveryRuleLogPathRule `pulumi:"logPathRules"`
	// Specifies the rule name, which contains `4` to `63` characters. It must start
	// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
	Name *string `pulumi:"name"`
	// Specifies the naming rules for discovered services and applications.
	// The object structure is documented below.
	NameRules ServiceDiscoveryRuleNameRules `pulumi:"nameRules"`
	// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
	Priority *int `pulumi:"priority"`
	// The region in which to create the service discovery rule resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
	ServiceType string `pulumi:"serviceType"`
}

// The set of arguments for constructing a ServiceDiscoveryRule resource.
type ServiceDiscoveryRuleArgs struct {
	// Specifies the rule description.
	Description pulumi.StringPtrInput
	// Specifies whether to enable log collection. The default value is true.
	DetectLogEnabled pulumi.BoolPtrInput
	// Specifies whether the rule is enabled. The default value is true.
	DiscoveryRuleEnabled pulumi.BoolPtrInput
	// Specifies the discovery rule. If the array contains multiple conditions, only the
	// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
	// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
	// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
	// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
	// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
	// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
	DiscoveryRules ServiceDiscoveryRuleDiscoveryRuleArrayInput
	// Specifies whether the rule is the default one. The default value is false.
	IsDefaultRule pulumi.BoolPtrInput
	// Specifies the log file suffix. This is a list of strings.
	// The values can be: **log**, **trace**, and **out**.
	LogFileSuffixes pulumi.StringArrayInput
	// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
	// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
	// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
	LogPathRules ServiceDiscoveryRuleLogPathRuleArrayInput
	// Specifies the rule name, which contains `4` to `63` characters. It must start
	// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
	Name pulumi.StringPtrInput
	// Specifies the naming rules for discovered services and applications.
	// The object structure is documented below.
	NameRules ServiceDiscoveryRuleNameRulesInput
	// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
	Priority pulumi.IntPtrInput
	// The region in which to create the service discovery rule resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
	ServiceType pulumi.StringInput
}

func (ServiceDiscoveryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDiscoveryRuleArgs)(nil)).Elem()
}

type ServiceDiscoveryRuleInput interface {
	pulumi.Input

	ToServiceDiscoveryRuleOutput() ServiceDiscoveryRuleOutput
	ToServiceDiscoveryRuleOutputWithContext(ctx context.Context) ServiceDiscoveryRuleOutput
}

func (*ServiceDiscoveryRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDiscoveryRule)(nil)).Elem()
}

func (i *ServiceDiscoveryRule) ToServiceDiscoveryRuleOutput() ServiceDiscoveryRuleOutput {
	return i.ToServiceDiscoveryRuleOutputWithContext(context.Background())
}

func (i *ServiceDiscoveryRule) ToServiceDiscoveryRuleOutputWithContext(ctx context.Context) ServiceDiscoveryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDiscoveryRuleOutput)
}

// ServiceDiscoveryRuleArrayInput is an input type that accepts ServiceDiscoveryRuleArray and ServiceDiscoveryRuleArrayOutput values.
// You can construct a concrete instance of `ServiceDiscoveryRuleArrayInput` via:
//
//	ServiceDiscoveryRuleArray{ ServiceDiscoveryRuleArgs{...} }
type ServiceDiscoveryRuleArrayInput interface {
	pulumi.Input

	ToServiceDiscoveryRuleArrayOutput() ServiceDiscoveryRuleArrayOutput
	ToServiceDiscoveryRuleArrayOutputWithContext(context.Context) ServiceDiscoveryRuleArrayOutput
}

type ServiceDiscoveryRuleArray []ServiceDiscoveryRuleInput

func (ServiceDiscoveryRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceDiscoveryRule)(nil)).Elem()
}

func (i ServiceDiscoveryRuleArray) ToServiceDiscoveryRuleArrayOutput() ServiceDiscoveryRuleArrayOutput {
	return i.ToServiceDiscoveryRuleArrayOutputWithContext(context.Background())
}

func (i ServiceDiscoveryRuleArray) ToServiceDiscoveryRuleArrayOutputWithContext(ctx context.Context) ServiceDiscoveryRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDiscoveryRuleArrayOutput)
}

// ServiceDiscoveryRuleMapInput is an input type that accepts ServiceDiscoveryRuleMap and ServiceDiscoveryRuleMapOutput values.
// You can construct a concrete instance of `ServiceDiscoveryRuleMapInput` via:
//
//	ServiceDiscoveryRuleMap{ "key": ServiceDiscoveryRuleArgs{...} }
type ServiceDiscoveryRuleMapInput interface {
	pulumi.Input

	ToServiceDiscoveryRuleMapOutput() ServiceDiscoveryRuleMapOutput
	ToServiceDiscoveryRuleMapOutputWithContext(context.Context) ServiceDiscoveryRuleMapOutput
}

type ServiceDiscoveryRuleMap map[string]ServiceDiscoveryRuleInput

func (ServiceDiscoveryRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceDiscoveryRule)(nil)).Elem()
}

func (i ServiceDiscoveryRuleMap) ToServiceDiscoveryRuleMapOutput() ServiceDiscoveryRuleMapOutput {
	return i.ToServiceDiscoveryRuleMapOutputWithContext(context.Background())
}

func (i ServiceDiscoveryRuleMap) ToServiceDiscoveryRuleMapOutputWithContext(ctx context.Context) ServiceDiscoveryRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDiscoveryRuleMapOutput)
}

type ServiceDiscoveryRuleOutput struct{ *pulumi.OutputState }

func (ServiceDiscoveryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDiscoveryRule)(nil)).Elem()
}

func (o ServiceDiscoveryRuleOutput) ToServiceDiscoveryRuleOutput() ServiceDiscoveryRuleOutput {
	return o
}

func (o ServiceDiscoveryRuleOutput) ToServiceDiscoveryRuleOutputWithContext(ctx context.Context) ServiceDiscoveryRuleOutput {
	return o
}

// The rule create time.
func (o ServiceDiscoveryRuleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the rule description.
func (o ServiceDiscoveryRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable log collection. The default value is true.
func (o ServiceDiscoveryRuleOutput) DetectLogEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.BoolPtrOutput { return v.DetectLogEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether the rule is enabled. The default value is true.
func (o ServiceDiscoveryRuleOutput) DiscoveryRuleEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.BoolPtrOutput { return v.DiscoveryRuleEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the discovery rule. If the array contains multiple conditions, only the
// processes that meet all the conditions will be matched. If the value of `checkType` is **cmdLine**, set the value of
// `checkMode` to **contain**. `checkContent` is in the format of ["xxx"], indicating that the process must contain
// the xxx parameter. If the value of `checkType` is **env**, set the value of `checkMode` to **contain**.
// `checkContent` is in the format of ["k1","v1"], indicating that the process must contain the environment variable
// whose name is k1 and value is v1. If the value of `checkType` is **scope**, set the value of `checkMode`
// to **equals**. `checkContent` is in the format of ["hostId1","hostId2"], indicating that the rule takes effect only
// on specified nodes. If no nodes are specified, the rule applies to all nodes of the project.
func (o ServiceDiscoveryRuleOutput) DiscoveryRules() ServiceDiscoveryRuleDiscoveryRuleArrayOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) ServiceDiscoveryRuleDiscoveryRuleArrayOutput { return v.DiscoveryRules }).(ServiceDiscoveryRuleDiscoveryRuleArrayOutput)
}

// Specifies whether the rule is the default one. The default value is false.
func (o ServiceDiscoveryRuleOutput) IsDefaultRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.BoolPtrOutput { return v.IsDefaultRule }).(pulumi.BoolPtrOutput)
}

// Specifies the log file suffix. This is a list of strings.
// The values can be: **log**, **trace**, and **out**.
func (o ServiceDiscoveryRuleOutput) LogFileSuffixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringArrayOutput { return v.LogFileSuffixes }).(pulumi.StringArrayOutput)
}

// Specifies the log path configuration rule. If cmdLineHash is a fixed string,
// logs in the specified log path or log file are collected. Otherwise, only the files whose names end with
// .log or .trace are collected. If the value of `nameType` is **cmdLineHash**, args is in the format of ["00001"] and
// value is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when the startup command is 00001.
func (o ServiceDiscoveryRuleOutput) LogPathRules() ServiceDiscoveryRuleLogPathRuleArrayOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) ServiceDiscoveryRuleLogPathRuleArrayOutput { return v.LogPathRules }).(ServiceDiscoveryRuleLogPathRuleArrayOutput)
}

// Specifies the rule name, which contains `4` to `63` characters. It must start
// with a lowercase letter but cannot end with a hyphen (-). Only digits, lowercase letters, and hyphens are allowed.
func (o ServiceDiscoveryRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the naming rules for discovered services and applications.
// The object structure is documented below.
func (o ServiceDiscoveryRuleOutput) NameRules() ServiceDiscoveryRuleNameRulesOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) ServiceDiscoveryRuleNameRulesOutput { return v.NameRules }).(ServiceDiscoveryRuleNameRulesOutput)
}

// Specifies the rule priority. Value range: 1 to 9999. The default value is 9999.
func (o ServiceDiscoveryRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The region in which to create the service discovery rule resource. If omitted,
// the provider-level region will be used. Changing this creates a new resource.
func (o ServiceDiscoveryRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The rule ID in uuid format.
func (o ServiceDiscoveryRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Specifies the service type, which is used only for rule classification and UI display.
// You can enter any field. For example, enter Java or Python by technology stack, or enter collector or database by function.
func (o ServiceDiscoveryRuleOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDiscoveryRule) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

type ServiceDiscoveryRuleArrayOutput struct{ *pulumi.OutputState }

func (ServiceDiscoveryRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceDiscoveryRule)(nil)).Elem()
}

func (o ServiceDiscoveryRuleArrayOutput) ToServiceDiscoveryRuleArrayOutput() ServiceDiscoveryRuleArrayOutput {
	return o
}

func (o ServiceDiscoveryRuleArrayOutput) ToServiceDiscoveryRuleArrayOutputWithContext(ctx context.Context) ServiceDiscoveryRuleArrayOutput {
	return o
}

func (o ServiceDiscoveryRuleArrayOutput) Index(i pulumi.IntInput) ServiceDiscoveryRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceDiscoveryRule {
		return vs[0].([]*ServiceDiscoveryRule)[vs[1].(int)]
	}).(ServiceDiscoveryRuleOutput)
}

type ServiceDiscoveryRuleMapOutput struct{ *pulumi.OutputState }

func (ServiceDiscoveryRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceDiscoveryRule)(nil)).Elem()
}

func (o ServiceDiscoveryRuleMapOutput) ToServiceDiscoveryRuleMapOutput() ServiceDiscoveryRuleMapOutput {
	return o
}

func (o ServiceDiscoveryRuleMapOutput) ToServiceDiscoveryRuleMapOutputWithContext(ctx context.Context) ServiceDiscoveryRuleMapOutput {
	return o
}

func (o ServiceDiscoveryRuleMapOutput) MapIndex(k pulumi.StringInput) ServiceDiscoveryRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceDiscoveryRule {
		return vs[0].(map[string]*ServiceDiscoveryRule)[vs[1].(string)]
	}).(ServiceDiscoveryRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDiscoveryRuleInput)(nil)).Elem(), &ServiceDiscoveryRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDiscoveryRuleArrayInput)(nil)).Elem(), ServiceDiscoveryRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDiscoveryRuleMapInput)(nil)).Elem(), ServiceDiscoveryRuleMap{})
	pulumi.RegisterOutputType(ServiceDiscoveryRuleOutput{})
	pulumi.RegisterOutputType(ServiceDiscoveryRuleArrayOutput{})
	pulumi.RegisterOutputType(ServiceDiscoveryRuleMapOutput{})
}
