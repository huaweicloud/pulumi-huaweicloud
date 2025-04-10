// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a WAF policy resource within HuaweiCloud.
//
// > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
// used. The policy resource can be used in Cloud Mode and Dedicated Mode.
//
// ## Import
//
// There are two ways to import WAF policy state. * Using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/policy:Policy test <id>
//
// ```
//
//   - Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/policy:Policy test <id>/<enterprise_project_id>
//
// ```
type Policy struct {
	pulumi.CustomResourceState

	// The protection switches. The options object structure is documented below.
	BindHosts PolicyBindHostArrayOutput `pulumi:"bindHosts"`
	// Specifies the deep inspection in basic web protection. Defaults to **false**.
	DeepInspection pulumi.BoolOutput `pulumi:"deepInspection"`
	// Specifies the enterprise project ID of WAF policy.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	// Specifies the detection mode in precise protection. Defaults to **false**.
	// + **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
	//   checks and blocks the request immediately.
	// + **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
	//   request immediately. Instead, it blocks the requests until other checks are finished.
	FullDetection pulumi.BoolPtrOutput `pulumi:"fullDetection"`
	// Specifies the header inspection in basic web protection. Defaults to **false**.
	HeaderInspection pulumi.BoolOutput `pulumi:"headerInspection"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
	//   false alarms have been reported, this value is recommended.
	// + `2`: Medium. This protection level meets web protection requirements in most scenarios.
	// + `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
	//   complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
	//   and Druid SQL injection attacks.
	Level pulumi.IntOutput `pulumi:"level"`
	// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
	// underscores (_), and hyphens (-) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the switch options of the protection item in the policy.
	// The options structure is documented below.
	Options PolicyOptionArrayOutput `pulumi:"options"`
	// Specifies the protective action after a rule is matched. Defaults to **log**.
	// Valid values are:
	// + **block**: WAF blocks and logs detected attacks.
	// + **log**: WAF logs detected attacks only.
	ProtectionMode pulumi.StringOutput `pulumi:"protectionMode"`
	// Specifies the region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the protective actions for each rule in anti-crawler protection.
	// Defaults to **log**. Valid values are:
	// + **block**: WAF blocks discovered attacks.
	// + **log**: WAF only logs discovered attacks.
	RobotAction pulumi.StringOutput `pulumi:"robotAction"`
	// Specifies the shiro decryption check in basic web protection.
	// Defaults to **false**.
	ShiroDecryptionCheck pulumi.BoolOutput `pulumi:"shiroDecryptionCheck"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("huaweicloud:Waf/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("huaweicloud:Waf/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The protection switches. The options object structure is documented below.
	BindHosts []PolicyBindHost `pulumi:"bindHosts"`
	// Specifies the deep inspection in basic web protection. Defaults to **false**.
	DeepInspection *bool `pulumi:"deepInspection"`
	// Specifies the enterprise project ID of WAF policy.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the detection mode in precise protection. Defaults to **false**.
	// + **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
	//   checks and blocks the request immediately.
	// + **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
	//   request immediately. Instead, it blocks the requests until other checks are finished.
	FullDetection *bool `pulumi:"fullDetection"`
	// Specifies the header inspection in basic web protection. Defaults to **false**.
	HeaderInspection *bool `pulumi:"headerInspection"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
	//   false alarms have been reported, this value is recommended.
	// + `2`: Medium. This protection level meets web protection requirements in most scenarios.
	// + `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
	//   complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
	//   and Druid SQL injection attacks.
	Level *int `pulumi:"level"`
	// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
	// underscores (_), and hyphens (-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the switch options of the protection item in the policy.
	// The options structure is documented below.
	Options []PolicyOption `pulumi:"options"`
	// Specifies the protective action after a rule is matched. Defaults to **log**.
	// Valid values are:
	// + **block**: WAF blocks and logs detected attacks.
	// + **log**: WAF logs detected attacks only.
	ProtectionMode *string `pulumi:"protectionMode"`
	// Specifies the region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region *string `pulumi:"region"`
	// Specifies the protective actions for each rule in anti-crawler protection.
	// Defaults to **log**. Valid values are:
	// + **block**: WAF blocks discovered attacks.
	// + **log**: WAF only logs discovered attacks.
	RobotAction *string `pulumi:"robotAction"`
	// Specifies the shiro decryption check in basic web protection.
	// Defaults to **false**.
	ShiroDecryptionCheck *bool `pulumi:"shiroDecryptionCheck"`
}

type PolicyState struct {
	// The protection switches. The options object structure is documented below.
	BindHosts PolicyBindHostArrayInput
	// Specifies the deep inspection in basic web protection. Defaults to **false**.
	DeepInspection pulumi.BoolPtrInput
	// Specifies the enterprise project ID of WAF policy.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the detection mode in precise protection. Defaults to **false**.
	// + **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
	//   checks and blocks the request immediately.
	// + **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
	//   request immediately. Instead, it blocks the requests until other checks are finished.
	FullDetection pulumi.BoolPtrInput
	// Specifies the header inspection in basic web protection. Defaults to **false**.
	HeaderInspection pulumi.BoolPtrInput
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
	//   false alarms have been reported, this value is recommended.
	// + `2`: Medium. This protection level meets web protection requirements in most scenarios.
	// + `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
	//   complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
	//   and Druid SQL injection attacks.
	Level pulumi.IntPtrInput
	// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
	// underscores (_), and hyphens (-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the switch options of the protection item in the policy.
	// The options structure is documented below.
	Options PolicyOptionArrayInput
	// Specifies the protective action after a rule is matched. Defaults to **log**.
	// Valid values are:
	// + **block**: WAF blocks and logs detected attacks.
	// + **log**: WAF logs detected attacks only.
	ProtectionMode pulumi.StringPtrInput
	// Specifies the region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringPtrInput
	// Specifies the protective actions for each rule in anti-crawler protection.
	// Defaults to **log**. Valid values are:
	// + **block**: WAF blocks discovered attacks.
	// + **log**: WAF only logs discovered attacks.
	RobotAction pulumi.StringPtrInput
	// Specifies the shiro decryption check in basic web protection.
	// Defaults to **false**.
	ShiroDecryptionCheck pulumi.BoolPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Specifies the deep inspection in basic web protection. Defaults to **false**.
	DeepInspection *bool `pulumi:"deepInspection"`
	// Specifies the enterprise project ID of WAF policy.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the detection mode in precise protection. Defaults to **false**.
	// + **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
	//   checks and blocks the request immediately.
	// + **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
	//   request immediately. Instead, it blocks the requests until other checks are finished.
	FullDetection *bool `pulumi:"fullDetection"`
	// Specifies the header inspection in basic web protection. Defaults to **false**.
	HeaderInspection *bool `pulumi:"headerInspection"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
	//   false alarms have been reported, this value is recommended.
	// + `2`: Medium. This protection level meets web protection requirements in most scenarios.
	// + `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
	//   complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
	//   and Druid SQL injection attacks.
	Level *int `pulumi:"level"`
	// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
	// underscores (_), and hyphens (-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the switch options of the protection item in the policy.
	// The options structure is documented below.
	Options []PolicyOption `pulumi:"options"`
	// Specifies the protective action after a rule is matched. Defaults to **log**.
	// Valid values are:
	// + **block**: WAF blocks and logs detected attacks.
	// + **log**: WAF logs detected attacks only.
	ProtectionMode *string `pulumi:"protectionMode"`
	// Specifies the region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region *string `pulumi:"region"`
	// Specifies the protective actions for each rule in anti-crawler protection.
	// Defaults to **log**. Valid values are:
	// + **block**: WAF blocks discovered attacks.
	// + **log**: WAF only logs discovered attacks.
	RobotAction *string `pulumi:"robotAction"`
	// Specifies the shiro decryption check in basic web protection.
	// Defaults to **false**.
	ShiroDecryptionCheck *bool `pulumi:"shiroDecryptionCheck"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies the deep inspection in basic web protection. Defaults to **false**.
	DeepInspection pulumi.BoolPtrInput
	// Specifies the enterprise project ID of WAF policy.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the detection mode in precise protection. Defaults to **false**.
	// + **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
	//   checks and blocks the request immediately.
	// + **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
	//   request immediately. Instead, it blocks the requests until other checks are finished.
	FullDetection pulumi.BoolPtrInput
	// Specifies the header inspection in basic web protection. Defaults to **false**.
	HeaderInspection pulumi.BoolPtrInput
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
	//   false alarms have been reported, this value is recommended.
	// + `2`: Medium. This protection level meets web protection requirements in most scenarios.
	// + `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
	//   complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
	//   and Druid SQL injection attacks.
	Level pulumi.IntPtrInput
	// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
	// underscores (_), and hyphens (-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the switch options of the protection item in the policy.
	// The options structure is documented below.
	Options PolicyOptionArrayInput
	// Specifies the protective action after a rule is matched. Defaults to **log**.
	// Valid values are:
	// + **block**: WAF blocks and logs detected attacks.
	// + **log**: WAF logs detected attacks only.
	ProtectionMode pulumi.StringPtrInput
	// Specifies the region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringPtrInput
	// Specifies the protective actions for each rule in anti-crawler protection.
	// Defaults to **log**. Valid values are:
	// + **block**: WAF blocks discovered attacks.
	// + **log**: WAF only logs discovered attacks.
	RobotAction pulumi.StringPtrInput
	// Specifies the shiro decryption check in basic web protection.
	// Defaults to **false**.
	ShiroDecryptionCheck pulumi.BoolPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// The protection switches. The options object structure is documented below.
func (o PolicyOutput) BindHosts() PolicyBindHostArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyBindHostArrayOutput { return v.BindHosts }).(PolicyBindHostArrayOutput)
}

// Specifies the deep inspection in basic web protection. Defaults to **false**.
func (o PolicyOutput) DeepInspection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.DeepInspection }).(pulumi.BoolOutput)
}

// Specifies the enterprise project ID of WAF policy.
// For enterprise users, if omitted, default enterprise project will be used.
// Changing this parameter will create a new resource.
func (o PolicyOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// Specifies the detection mode in precise protection. Defaults to **false**.
//   - **false**: Instant detection. When a request hits the blocking conditions in precise protection, WAF terminates
//     checks and blocks the request immediately.
//   - **true**: Full detection. If a request hits the blocking conditions in precise protection, WAF does not block the
//     request immediately. Instead, it blocks the requests until other checks are finished.
func (o PolicyOutput) FullDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolPtrOutput { return v.FullDetection }).(pulumi.BoolPtrOutput)
}

// Specifies the header inspection in basic web protection. Defaults to **false**.
func (o PolicyOutput) HeaderInspection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.HeaderInspection }).(pulumi.BoolOutput)
}

// Specifies the protection level. Defaults to `2`. Valid values are:
//   - `1`: Low. At this protection level, WAF blocks only requests with obvious attack features. If a large number of
//     false alarms have been reported, this value is recommended.
//   - `2`: Medium. This protection level meets web protection requirements in most scenarios.
//   - `3`: High. At this protection level, WAF provides the finest granular protection and can intercept attacks with
//     complex bypass features, such as Jolokia cyberattacks, common gateway interface (CGI) vulnerability detection,
//     and Druid SQL injection attacks.
func (o PolicyOutput) Level() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy) pulumi.IntOutput { return v.Level }).(pulumi.IntOutput)
}

// Specifies the policy name. The maximum length is `256` characters. Only digits, letters,
// underscores (_), and hyphens (-) are allowed.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the switch options of the protection item in the policy.
// The options structure is documented below.
func (o PolicyOutput) Options() PolicyOptionArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyOptionArrayOutput { return v.Options }).(PolicyOptionArrayOutput)
}

// Specifies the protective action after a rule is matched. Defaults to **log**.
// Valid values are:
// + **block**: WAF blocks and logs detected attacks.
// + **log**: WAF logs detected attacks only.
func (o PolicyOutput) ProtectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProtectionMode }).(pulumi.StringOutput)
}

// Specifies the region in which to create the WAF policy resource. If omitted, the
// provider-level region will be used. Changing this setting will push a new certificate.
func (o PolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the protective actions for each rule in anti-crawler protection.
// Defaults to **log**. Valid values are:
// + **block**: WAF blocks discovered attacks.
// + **log**: WAF only logs discovered attacks.
func (o PolicyOutput) RobotAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.RobotAction }).(pulumi.StringOutput)
}

// Specifies the shiro decryption check in basic web protection.
// Defaults to **false**.
func (o PolicyOutput) ShiroDecryptionCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.ShiroDecryptionCheck }).(pulumi.BoolOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
