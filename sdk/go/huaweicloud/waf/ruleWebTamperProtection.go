// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a WAF web tamper protection rule resource within HuaweiCloud.
//
// > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
// used. The web tamper protection rule resource can be used in Cloud Mode, Dedicated Mode and ELB Mode.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			policyId := cfg.RequireObject("policyId")
//			_, err := Waf.NewRuleWebTamperProtection(ctx, "rule1", &Waf.RuleWebTamperProtectionArgs{
//				PolicyId:            pulumi.Any(policyId),
//				EnterpriseProjectId: pulumi.Any(enterpriseProjectId),
//				Domain:              pulumi.String("www.your-domain.com"),
//				Path:                pulumi.String("/payment"),
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
// There are two ways to import WAF rule web tamper protection state. * Using `policy_id` and `rule_id`, separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/ruleWebTamperProtection:RuleWebTamperProtection test <policy_id>/<rule_id>
//
// ```
//
//   - Using `policy_id`, `rule_id` and `enterprise_project_id`, separated by slashes, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/ruleWebTamperProtection:RuleWebTamperProtection test <policy_id>/<rule_id>/<enterprise_project_id>
//
// ```
type RuleWebTamperProtection struct {
	pulumi.CustomResourceState

	// Specifies the domain name. Changing this creates a new rule.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Specifies the enterprise project ID of WAF tamper protection
	// rule. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	// Specifies the URL protected by the web tamper protection rule, excluding a
	// domain name. Changing this creates a new rule.
	Path pulumi.StringOutput `pulumi:"path"`
	// Specifies the WAF policy ID. Changing this creates a new rule.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The region in which to create the WAF web tamper protection rules resource. If
	// omitted, the provider-level region will be used. Changing this creates a new rule.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRuleWebTamperProtection registers a new resource with the given unique name, arguments, and options.
func NewRuleWebTamperProtection(ctx *pulumi.Context,
	name string, args *RuleWebTamperProtectionArgs, opts ...pulumi.ResourceOption) (*RuleWebTamperProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RuleWebTamperProtection
	err := ctx.RegisterResource("huaweicloud:Waf/ruleWebTamperProtection:RuleWebTamperProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleWebTamperProtection gets an existing RuleWebTamperProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleWebTamperProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleWebTamperProtectionState, opts ...pulumi.ResourceOption) (*RuleWebTamperProtection, error) {
	var resource RuleWebTamperProtection
	err := ctx.ReadResource("huaweicloud:Waf/ruleWebTamperProtection:RuleWebTamperProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleWebTamperProtection resources.
type ruleWebTamperProtectionState struct {
	// Specifies the domain name. Changing this creates a new rule.
	Domain *string `pulumi:"domain"`
	// Specifies the enterprise project ID of WAF tamper protection
	// rule. Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the URL protected by the web tamper protection rule, excluding a
	// domain name. Changing this creates a new rule.
	Path *string `pulumi:"path"`
	// Specifies the WAF policy ID. Changing this creates a new rule.
	PolicyId *string `pulumi:"policyId"`
	// The region in which to create the WAF web tamper protection rules resource. If
	// omitted, the provider-level region will be used. Changing this creates a new rule.
	Region *string `pulumi:"region"`
}

type RuleWebTamperProtectionState struct {
	// Specifies the domain name. Changing this creates a new rule.
	Domain pulumi.StringPtrInput
	// Specifies the enterprise project ID of WAF tamper protection
	// rule. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the URL protected by the web tamper protection rule, excluding a
	// domain name. Changing this creates a new rule.
	Path pulumi.StringPtrInput
	// Specifies the WAF policy ID. Changing this creates a new rule.
	PolicyId pulumi.StringPtrInput
	// The region in which to create the WAF web tamper protection rules resource. If
	// omitted, the provider-level region will be used. Changing this creates a new rule.
	Region pulumi.StringPtrInput
}

func (RuleWebTamperProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleWebTamperProtectionState)(nil)).Elem()
}

type ruleWebTamperProtectionArgs struct {
	// Specifies the domain name. Changing this creates a new rule.
	Domain string `pulumi:"domain"`
	// Specifies the enterprise project ID of WAF tamper protection
	// rule. Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the URL protected by the web tamper protection rule, excluding a
	// domain name. Changing this creates a new rule.
	Path string `pulumi:"path"`
	// Specifies the WAF policy ID. Changing this creates a new rule.
	PolicyId string `pulumi:"policyId"`
	// The region in which to create the WAF web tamper protection rules resource. If
	// omitted, the provider-level region will be used. Changing this creates a new rule.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RuleWebTamperProtection resource.
type RuleWebTamperProtectionArgs struct {
	// Specifies the domain name. Changing this creates a new rule.
	Domain pulumi.StringInput
	// Specifies the enterprise project ID of WAF tamper protection
	// rule. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the URL protected by the web tamper protection rule, excluding a
	// domain name. Changing this creates a new rule.
	Path pulumi.StringInput
	// Specifies the WAF policy ID. Changing this creates a new rule.
	PolicyId pulumi.StringInput
	// The region in which to create the WAF web tamper protection rules resource. If
	// omitted, the provider-level region will be used. Changing this creates a new rule.
	Region pulumi.StringPtrInput
}

func (RuleWebTamperProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleWebTamperProtectionArgs)(nil)).Elem()
}

type RuleWebTamperProtectionInput interface {
	pulumi.Input

	ToRuleWebTamperProtectionOutput() RuleWebTamperProtectionOutput
	ToRuleWebTamperProtectionOutputWithContext(ctx context.Context) RuleWebTamperProtectionOutput
}

func (*RuleWebTamperProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleWebTamperProtection)(nil)).Elem()
}

func (i *RuleWebTamperProtection) ToRuleWebTamperProtectionOutput() RuleWebTamperProtectionOutput {
	return i.ToRuleWebTamperProtectionOutputWithContext(context.Background())
}

func (i *RuleWebTamperProtection) ToRuleWebTamperProtectionOutputWithContext(ctx context.Context) RuleWebTamperProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebTamperProtectionOutput)
}

// RuleWebTamperProtectionArrayInput is an input type that accepts RuleWebTamperProtectionArray and RuleWebTamperProtectionArrayOutput values.
// You can construct a concrete instance of `RuleWebTamperProtectionArrayInput` via:
//
//	RuleWebTamperProtectionArray{ RuleWebTamperProtectionArgs{...} }
type RuleWebTamperProtectionArrayInput interface {
	pulumi.Input

	ToRuleWebTamperProtectionArrayOutput() RuleWebTamperProtectionArrayOutput
	ToRuleWebTamperProtectionArrayOutputWithContext(context.Context) RuleWebTamperProtectionArrayOutput
}

type RuleWebTamperProtectionArray []RuleWebTamperProtectionInput

func (RuleWebTamperProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleWebTamperProtection)(nil)).Elem()
}

func (i RuleWebTamperProtectionArray) ToRuleWebTamperProtectionArrayOutput() RuleWebTamperProtectionArrayOutput {
	return i.ToRuleWebTamperProtectionArrayOutputWithContext(context.Background())
}

func (i RuleWebTamperProtectionArray) ToRuleWebTamperProtectionArrayOutputWithContext(ctx context.Context) RuleWebTamperProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebTamperProtectionArrayOutput)
}

// RuleWebTamperProtectionMapInput is an input type that accepts RuleWebTamperProtectionMap and RuleWebTamperProtectionMapOutput values.
// You can construct a concrete instance of `RuleWebTamperProtectionMapInput` via:
//
//	RuleWebTamperProtectionMap{ "key": RuleWebTamperProtectionArgs{...} }
type RuleWebTamperProtectionMapInput interface {
	pulumi.Input

	ToRuleWebTamperProtectionMapOutput() RuleWebTamperProtectionMapOutput
	ToRuleWebTamperProtectionMapOutputWithContext(context.Context) RuleWebTamperProtectionMapOutput
}

type RuleWebTamperProtectionMap map[string]RuleWebTamperProtectionInput

func (RuleWebTamperProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleWebTamperProtection)(nil)).Elem()
}

func (i RuleWebTamperProtectionMap) ToRuleWebTamperProtectionMapOutput() RuleWebTamperProtectionMapOutput {
	return i.ToRuleWebTamperProtectionMapOutputWithContext(context.Background())
}

func (i RuleWebTamperProtectionMap) ToRuleWebTamperProtectionMapOutputWithContext(ctx context.Context) RuleWebTamperProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebTamperProtectionMapOutput)
}

type RuleWebTamperProtectionOutput struct{ *pulumi.OutputState }

func (RuleWebTamperProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleWebTamperProtection)(nil)).Elem()
}

func (o RuleWebTamperProtectionOutput) ToRuleWebTamperProtectionOutput() RuleWebTamperProtectionOutput {
	return o
}

func (o RuleWebTamperProtectionOutput) ToRuleWebTamperProtectionOutputWithContext(ctx context.Context) RuleWebTamperProtectionOutput {
	return o
}

// Specifies the domain name. Changing this creates a new rule.
func (o RuleWebTamperProtectionOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleWebTamperProtection) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Specifies the enterprise project ID of WAF tamper protection
// rule. Changing this parameter will create a new resource.
func (o RuleWebTamperProtectionOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleWebTamperProtection) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// Specifies the URL protected by the web tamper protection rule, excluding a
// domain name. Changing this creates a new rule.
func (o RuleWebTamperProtectionOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleWebTamperProtection) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Specifies the WAF policy ID. Changing this creates a new rule.
func (o RuleWebTamperProtectionOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleWebTamperProtection) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The region in which to create the WAF web tamper protection rules resource. If
// omitted, the provider-level region will be used. Changing this creates a new rule.
func (o RuleWebTamperProtectionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleWebTamperProtection) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RuleWebTamperProtectionArrayOutput struct{ *pulumi.OutputState }

func (RuleWebTamperProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleWebTamperProtection)(nil)).Elem()
}

func (o RuleWebTamperProtectionArrayOutput) ToRuleWebTamperProtectionArrayOutput() RuleWebTamperProtectionArrayOutput {
	return o
}

func (o RuleWebTamperProtectionArrayOutput) ToRuleWebTamperProtectionArrayOutputWithContext(ctx context.Context) RuleWebTamperProtectionArrayOutput {
	return o
}

func (o RuleWebTamperProtectionArrayOutput) Index(i pulumi.IntInput) RuleWebTamperProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleWebTamperProtection {
		return vs[0].([]*RuleWebTamperProtection)[vs[1].(int)]
	}).(RuleWebTamperProtectionOutput)
}

type RuleWebTamperProtectionMapOutput struct{ *pulumi.OutputState }

func (RuleWebTamperProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleWebTamperProtection)(nil)).Elem()
}

func (o RuleWebTamperProtectionMapOutput) ToRuleWebTamperProtectionMapOutput() RuleWebTamperProtectionMapOutput {
	return o
}

func (o RuleWebTamperProtectionMapOutput) ToRuleWebTamperProtectionMapOutputWithContext(ctx context.Context) RuleWebTamperProtectionMapOutput {
	return o
}

func (o RuleWebTamperProtectionMapOutput) MapIndex(k pulumi.StringInput) RuleWebTamperProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleWebTamperProtection {
		return vs[0].(map[string]*RuleWebTamperProtection)[vs[1].(string)]
	}).(RuleWebTamperProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleWebTamperProtectionInput)(nil)).Elem(), &RuleWebTamperProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleWebTamperProtectionArrayInput)(nil)).Elem(), RuleWebTamperProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleWebTamperProtectionMapInput)(nil)).Elem(), RuleWebTamperProtectionMap{})
	pulumi.RegisterOutputType(RuleWebTamperProtectionOutput{})
	pulumi.RegisterOutputType(RuleWebTamperProtectionArrayOutput{})
	pulumi.RegisterOutputType(RuleWebTamperProtectionMapOutput{})
}