// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package smn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionExtension struct {
	// Specifies the client ID. This field is the tenant ID field in
	// the WeLink subscription and is obtained by the tenant from WeLink. This field is mandatory when `protocol`
	// is set to **welink**. Changing this parameter will create a new resource.
	ClientId *string `pulumi:"clientId"`
	// Specifies the client secret. This field is the client secret
	// field obtained by the tenant from WeLink. This field is mandatory when `protocol` is set to **welink**.
	// Changing this parameter will create a new resource.
	ClientSecret *string `pulumi:"clientSecret"`
	// Specifies the keyword. When `protocol` is set to **feishu**,
	// either `keyword` or `signSecret` must be specified. When you use `keywords` to configure a security policy
	// for the Lark or DingTalk chatbot on SMN, the keywords must have one of the keywords configured on the Lark
	// or DingTalk client. Changing this parameter will create a new resource.
	Keyword *string `pulumi:"keyword"`
	// Specifies the key including signature. When `protocol` is set
	// to **feishu** or **dingding**, this field or `keyword` must be specified. The key configurations must be
	// the same as those on the Lark or DingTalk client. For example, if only key is configured on the Lark client,
	// enter the key field obtained from the Lark client. If only keyword is configured on the Lark client, skip this field.
	// Changing this parameter will create a new resource.
	SignSecret *string `pulumi:"signSecret"`
}

// SubscriptionExtensionInput is an input type that accepts SubscriptionExtensionArgs and SubscriptionExtensionOutput values.
// You can construct a concrete instance of `SubscriptionExtensionInput` via:
//
//	SubscriptionExtensionArgs{...}
type SubscriptionExtensionInput interface {
	pulumi.Input

	ToSubscriptionExtensionOutput() SubscriptionExtensionOutput
	ToSubscriptionExtensionOutputWithContext(context.Context) SubscriptionExtensionOutput
}

type SubscriptionExtensionArgs struct {
	// Specifies the client ID. This field is the tenant ID field in
	// the WeLink subscription and is obtained by the tenant from WeLink. This field is mandatory when `protocol`
	// is set to **welink**. Changing this parameter will create a new resource.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// Specifies the client secret. This field is the client secret
	// field obtained by the tenant from WeLink. This field is mandatory when `protocol` is set to **welink**.
	// Changing this parameter will create a new resource.
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
	// Specifies the keyword. When `protocol` is set to **feishu**,
	// either `keyword` or `signSecret` must be specified. When you use `keywords` to configure a security policy
	// for the Lark or DingTalk chatbot on SMN, the keywords must have one of the keywords configured on the Lark
	// or DingTalk client. Changing this parameter will create a new resource.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// Specifies the key including signature. When `protocol` is set
	// to **feishu** or **dingding**, this field or `keyword` must be specified. The key configurations must be
	// the same as those on the Lark or DingTalk client. For example, if only key is configured on the Lark client,
	// enter the key field obtained from the Lark client. If only keyword is configured on the Lark client, skip this field.
	// Changing this parameter will create a new resource.
	SignSecret pulumi.StringPtrInput `pulumi:"signSecret"`
}

func (SubscriptionExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionExtension)(nil)).Elem()
}

func (i SubscriptionExtensionArgs) ToSubscriptionExtensionOutput() SubscriptionExtensionOutput {
	return i.ToSubscriptionExtensionOutputWithContext(context.Background())
}

func (i SubscriptionExtensionArgs) ToSubscriptionExtensionOutputWithContext(ctx context.Context) SubscriptionExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionExtensionOutput)
}

func (i SubscriptionExtensionArgs) ToSubscriptionExtensionPtrOutput() SubscriptionExtensionPtrOutput {
	return i.ToSubscriptionExtensionPtrOutputWithContext(context.Background())
}

func (i SubscriptionExtensionArgs) ToSubscriptionExtensionPtrOutputWithContext(ctx context.Context) SubscriptionExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionExtensionOutput).ToSubscriptionExtensionPtrOutputWithContext(ctx)
}

// SubscriptionExtensionPtrInput is an input type that accepts SubscriptionExtensionArgs, SubscriptionExtensionPtr and SubscriptionExtensionPtrOutput values.
// You can construct a concrete instance of `SubscriptionExtensionPtrInput` via:
//
//	        SubscriptionExtensionArgs{...}
//
//	or:
//
//	        nil
type SubscriptionExtensionPtrInput interface {
	pulumi.Input

	ToSubscriptionExtensionPtrOutput() SubscriptionExtensionPtrOutput
	ToSubscriptionExtensionPtrOutputWithContext(context.Context) SubscriptionExtensionPtrOutput
}

type subscriptionExtensionPtrType SubscriptionExtensionArgs

func SubscriptionExtensionPtr(v *SubscriptionExtensionArgs) SubscriptionExtensionPtrInput {
	return (*subscriptionExtensionPtrType)(v)
}

func (*subscriptionExtensionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionExtension)(nil)).Elem()
}

func (i *subscriptionExtensionPtrType) ToSubscriptionExtensionPtrOutput() SubscriptionExtensionPtrOutput {
	return i.ToSubscriptionExtensionPtrOutputWithContext(context.Background())
}

func (i *subscriptionExtensionPtrType) ToSubscriptionExtensionPtrOutputWithContext(ctx context.Context) SubscriptionExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionExtensionPtrOutput)
}

type SubscriptionExtensionOutput struct{ *pulumi.OutputState }

func (SubscriptionExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionExtension)(nil)).Elem()
}

func (o SubscriptionExtensionOutput) ToSubscriptionExtensionOutput() SubscriptionExtensionOutput {
	return o
}

func (o SubscriptionExtensionOutput) ToSubscriptionExtensionOutputWithContext(ctx context.Context) SubscriptionExtensionOutput {
	return o
}

func (o SubscriptionExtensionOutput) ToSubscriptionExtensionPtrOutput() SubscriptionExtensionPtrOutput {
	return o.ToSubscriptionExtensionPtrOutputWithContext(context.Background())
}

func (o SubscriptionExtensionOutput) ToSubscriptionExtensionPtrOutputWithContext(ctx context.Context) SubscriptionExtensionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionExtension) *SubscriptionExtension {
		return &v
	}).(SubscriptionExtensionPtrOutput)
}

// Specifies the client ID. This field is the tenant ID field in
// the WeLink subscription and is obtained by the tenant from WeLink. This field is mandatory when `protocol`
// is set to **welink**. Changing this parameter will create a new resource.
func (o SubscriptionExtensionOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionExtension) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// Specifies the client secret. This field is the client secret
// field obtained by the tenant from WeLink. This field is mandatory when `protocol` is set to **welink**.
// Changing this parameter will create a new resource.
func (o SubscriptionExtensionOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionExtension) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Specifies the keyword. When `protocol` is set to **feishu**,
// either `keyword` or `signSecret` must be specified. When you use `keywords` to configure a security policy
// for the Lark or DingTalk chatbot on SMN, the keywords must have one of the keywords configured on the Lark
// or DingTalk client. Changing this parameter will create a new resource.
func (o SubscriptionExtensionOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionExtension) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

// Specifies the key including signature. When `protocol` is set
// to **feishu** or **dingding**, this field or `keyword` must be specified. The key configurations must be
// the same as those on the Lark or DingTalk client. For example, if only key is configured on the Lark client,
// enter the key field obtained from the Lark client. If only keyword is configured on the Lark client, skip this field.
// Changing this parameter will create a new resource.
func (o SubscriptionExtensionOutput) SignSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionExtension) *string { return v.SignSecret }).(pulumi.StringPtrOutput)
}

type SubscriptionExtensionPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionExtensionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionExtension)(nil)).Elem()
}

func (o SubscriptionExtensionPtrOutput) ToSubscriptionExtensionPtrOutput() SubscriptionExtensionPtrOutput {
	return o
}

func (o SubscriptionExtensionPtrOutput) ToSubscriptionExtensionPtrOutputWithContext(ctx context.Context) SubscriptionExtensionPtrOutput {
	return o
}

func (o SubscriptionExtensionPtrOutput) Elem() SubscriptionExtensionOutput {
	return o.ApplyT(func(v *SubscriptionExtension) SubscriptionExtension {
		if v != nil {
			return *v
		}
		var ret SubscriptionExtension
		return ret
	}).(SubscriptionExtensionOutput)
}

// Specifies the client ID. This field is the tenant ID field in
// the WeLink subscription and is obtained by the tenant from WeLink. This field is mandatory when `protocol`
// is set to **welink**. Changing this parameter will create a new resource.
func (o SubscriptionExtensionPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionExtension) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

// Specifies the client secret. This field is the client secret
// field obtained by the tenant from WeLink. This field is mandatory when `protocol` is set to **welink**.
// Changing this parameter will create a new resource.
func (o SubscriptionExtensionPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionExtension) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

// Specifies the keyword. When `protocol` is set to **feishu**,
// either `keyword` or `signSecret` must be specified. When you use `keywords` to configure a security policy
// for the Lark or DingTalk chatbot on SMN, the keywords must have one of the keywords configured on the Lark
// or DingTalk client. Changing this parameter will create a new resource.
func (o SubscriptionExtensionPtrOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionExtension) *string {
		if v == nil {
			return nil
		}
		return v.Keyword
	}).(pulumi.StringPtrOutput)
}

// Specifies the key including signature. When `protocol` is set
// to **feishu** or **dingding**, this field or `keyword` must be specified. The key configurations must be
// the same as those on the Lark or DingTalk client. For example, if only key is configured on the Lark client,
// enter the key field obtained from the Lark client. If only keyword is configured on the Lark client, skip this field.
// Changing this parameter will create a new resource.
func (o SubscriptionExtensionPtrOutput) SignSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionExtension) *string {
		if v == nil {
			return nil
		}
		return v.SignSecret
	}).(pulumi.StringPtrOutput)
}

type SubscriptionFilterPolicy struct {
	// The filter policy name.
	Name *string `pulumi:"name"`
	// The string array for exact match.
	StringEquals []string `pulumi:"stringEquals"`
}

// SubscriptionFilterPolicyInput is an input type that accepts SubscriptionFilterPolicyArgs and SubscriptionFilterPolicyOutput values.
// You can construct a concrete instance of `SubscriptionFilterPolicyInput` via:
//
//	SubscriptionFilterPolicyArgs{...}
type SubscriptionFilterPolicyInput interface {
	pulumi.Input

	ToSubscriptionFilterPolicyOutput() SubscriptionFilterPolicyOutput
	ToSubscriptionFilterPolicyOutputWithContext(context.Context) SubscriptionFilterPolicyOutput
}

type SubscriptionFilterPolicyArgs struct {
	// The filter policy name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The string array for exact match.
	StringEquals pulumi.StringArrayInput `pulumi:"stringEquals"`
}

func (SubscriptionFilterPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFilterPolicy)(nil)).Elem()
}

func (i SubscriptionFilterPolicyArgs) ToSubscriptionFilterPolicyOutput() SubscriptionFilterPolicyOutput {
	return i.ToSubscriptionFilterPolicyOutputWithContext(context.Background())
}

func (i SubscriptionFilterPolicyArgs) ToSubscriptionFilterPolicyOutputWithContext(ctx context.Context) SubscriptionFilterPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFilterPolicyOutput)
}

// SubscriptionFilterPolicyArrayInput is an input type that accepts SubscriptionFilterPolicyArray and SubscriptionFilterPolicyArrayOutput values.
// You can construct a concrete instance of `SubscriptionFilterPolicyArrayInput` via:
//
//	SubscriptionFilterPolicyArray{ SubscriptionFilterPolicyArgs{...} }
type SubscriptionFilterPolicyArrayInput interface {
	pulumi.Input

	ToSubscriptionFilterPolicyArrayOutput() SubscriptionFilterPolicyArrayOutput
	ToSubscriptionFilterPolicyArrayOutputWithContext(context.Context) SubscriptionFilterPolicyArrayOutput
}

type SubscriptionFilterPolicyArray []SubscriptionFilterPolicyInput

func (SubscriptionFilterPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionFilterPolicy)(nil)).Elem()
}

func (i SubscriptionFilterPolicyArray) ToSubscriptionFilterPolicyArrayOutput() SubscriptionFilterPolicyArrayOutput {
	return i.ToSubscriptionFilterPolicyArrayOutputWithContext(context.Background())
}

func (i SubscriptionFilterPolicyArray) ToSubscriptionFilterPolicyArrayOutputWithContext(ctx context.Context) SubscriptionFilterPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFilterPolicyArrayOutput)
}

type SubscriptionFilterPolicyOutput struct{ *pulumi.OutputState }

func (SubscriptionFilterPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFilterPolicy)(nil)).Elem()
}

func (o SubscriptionFilterPolicyOutput) ToSubscriptionFilterPolicyOutput() SubscriptionFilterPolicyOutput {
	return o
}

func (o SubscriptionFilterPolicyOutput) ToSubscriptionFilterPolicyOutputWithContext(ctx context.Context) SubscriptionFilterPolicyOutput {
	return o
}

// The filter policy name.
func (o SubscriptionFilterPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionFilterPolicy) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The string array for exact match.
func (o SubscriptionFilterPolicyOutput) StringEquals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SubscriptionFilterPolicy) []string { return v.StringEquals }).(pulumi.StringArrayOutput)
}

type SubscriptionFilterPolicyArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionFilterPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionFilterPolicy)(nil)).Elem()
}

func (o SubscriptionFilterPolicyArrayOutput) ToSubscriptionFilterPolicyArrayOutput() SubscriptionFilterPolicyArrayOutput {
	return o
}

func (o SubscriptionFilterPolicyArrayOutput) ToSubscriptionFilterPolicyArrayOutputWithContext(ctx context.Context) SubscriptionFilterPolicyArrayOutput {
	return o
}

func (o SubscriptionFilterPolicyArrayOutput) Index(i pulumi.IntInput) SubscriptionFilterPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionFilterPolicy {
		return vs[0].([]SubscriptionFilterPolicy)[vs[1].(int)]
	}).(SubscriptionFilterPolicyOutput)
}

type GetTopicsTopic struct {
	// Specifies the topic display name.
	DisplayName string `pulumi:"displayName"`
	// Specifies the enterprise project ID of the SMN topic.
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// The topic ID. The value is the topic URN.
	Id string `pulumi:"id"`
	// Specifies the name of the topic.
	Name string `pulumi:"name"`
	// Message pushing policy.
	// + **0**: indicates that the message sending fails and the message is cached in the queue.
	// + **1**: indicates that the failed message is discarded.
	PushPolicy int `pulumi:"pushPolicy"`
	// The tags of the SMN topic, key/value pair format.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the topic URN.
	TopicUrn string `pulumi:"topicUrn"`
}

// GetTopicsTopicInput is an input type that accepts GetTopicsTopicArgs and GetTopicsTopicOutput values.
// You can construct a concrete instance of `GetTopicsTopicInput` via:
//
//	GetTopicsTopicArgs{...}
type GetTopicsTopicInput interface {
	pulumi.Input

	ToGetTopicsTopicOutput() GetTopicsTopicOutput
	ToGetTopicsTopicOutputWithContext(context.Context) GetTopicsTopicOutput
}

type GetTopicsTopicArgs struct {
	// Specifies the topic display name.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Specifies the enterprise project ID of the SMN topic.
	EnterpriseProjectId pulumi.StringInput `pulumi:"enterpriseProjectId"`
	// The topic ID. The value is the topic URN.
	Id pulumi.StringInput `pulumi:"id"`
	// Specifies the name of the topic.
	Name pulumi.StringInput `pulumi:"name"`
	// Message pushing policy.
	// + **0**: indicates that the message sending fails and the message is cached in the queue.
	// + **1**: indicates that the failed message is discarded.
	PushPolicy pulumi.IntInput `pulumi:"pushPolicy"`
	// The tags of the SMN topic, key/value pair format.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Specifies the topic URN.
	TopicUrn pulumi.StringInput `pulumi:"topicUrn"`
}

func (GetTopicsTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicsTopic)(nil)).Elem()
}

func (i GetTopicsTopicArgs) ToGetTopicsTopicOutput() GetTopicsTopicOutput {
	return i.ToGetTopicsTopicOutputWithContext(context.Background())
}

func (i GetTopicsTopicArgs) ToGetTopicsTopicOutputWithContext(ctx context.Context) GetTopicsTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTopicsTopicOutput)
}

// GetTopicsTopicArrayInput is an input type that accepts GetTopicsTopicArray and GetTopicsTopicArrayOutput values.
// You can construct a concrete instance of `GetTopicsTopicArrayInput` via:
//
//	GetTopicsTopicArray{ GetTopicsTopicArgs{...} }
type GetTopicsTopicArrayInput interface {
	pulumi.Input

	ToGetTopicsTopicArrayOutput() GetTopicsTopicArrayOutput
	ToGetTopicsTopicArrayOutputWithContext(context.Context) GetTopicsTopicArrayOutput
}

type GetTopicsTopicArray []GetTopicsTopicInput

func (GetTopicsTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTopicsTopic)(nil)).Elem()
}

func (i GetTopicsTopicArray) ToGetTopicsTopicArrayOutput() GetTopicsTopicArrayOutput {
	return i.ToGetTopicsTopicArrayOutputWithContext(context.Background())
}

func (i GetTopicsTopicArray) ToGetTopicsTopicArrayOutputWithContext(ctx context.Context) GetTopicsTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTopicsTopicArrayOutput)
}

type GetTopicsTopicOutput struct{ *pulumi.OutputState }

func (GetTopicsTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicsTopic)(nil)).Elem()
}

func (o GetTopicsTopicOutput) ToGetTopicsTopicOutput() GetTopicsTopicOutput {
	return o
}

func (o GetTopicsTopicOutput) ToGetTopicsTopicOutputWithContext(ctx context.Context) GetTopicsTopicOutput {
	return o
}

// Specifies the topic display name.
func (o GetTopicsTopicOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsTopic) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies the enterprise project ID of the SMN topic.
func (o GetTopicsTopicOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsTopic) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The topic ID. The value is the topic URN.
func (o GetTopicsTopicOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsTopic) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the name of the topic.
func (o GetTopicsTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsTopic) string { return v.Name }).(pulumi.StringOutput)
}

// Message pushing policy.
// + **0**: indicates that the message sending fails and the message is cached in the queue.
// + **1**: indicates that the failed message is discarded.
func (o GetTopicsTopicOutput) PushPolicy() pulumi.IntOutput {
	return o.ApplyT(func(v GetTopicsTopic) int { return v.PushPolicy }).(pulumi.IntOutput)
}

// The tags of the SMN topic, key/value pair format.
func (o GetTopicsTopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetTopicsTopic) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the topic URN.
func (o GetTopicsTopicOutput) TopicUrn() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsTopic) string { return v.TopicUrn }).(pulumi.StringOutput)
}

type GetTopicsTopicArrayOutput struct{ *pulumi.OutputState }

func (GetTopicsTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTopicsTopic)(nil)).Elem()
}

func (o GetTopicsTopicArrayOutput) ToGetTopicsTopicArrayOutput() GetTopicsTopicArrayOutput {
	return o
}

func (o GetTopicsTopicArrayOutput) ToGetTopicsTopicArrayOutputWithContext(ctx context.Context) GetTopicsTopicArrayOutput {
	return o
}

func (o GetTopicsTopicArrayOutput) Index(i pulumi.IntInput) GetTopicsTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTopicsTopic {
		return vs[0].([]GetTopicsTopic)[vs[1].(int)]
	}).(GetTopicsTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionExtensionInput)(nil)).Elem(), SubscriptionExtensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionExtensionPtrInput)(nil)).Elem(), SubscriptionExtensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionFilterPolicyInput)(nil)).Elem(), SubscriptionFilterPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionFilterPolicyArrayInput)(nil)).Elem(), SubscriptionFilterPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTopicsTopicInput)(nil)).Elem(), GetTopicsTopicArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTopicsTopicArrayInput)(nil)).Elem(), GetTopicsTopicArray{})
	pulumi.RegisterOutputType(SubscriptionExtensionOutput{})
	pulumi.RegisterOutputType(SubscriptionExtensionPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionFilterPolicyOutput{})
	pulumi.RegisterOutputType(SubscriptionFilterPolicyArrayOutput{})
	pulumi.RegisterOutputType(GetTopicsTopicOutput{})
	pulumi.RegisterOutputType(GetTopicsTopicArrayOutput{})
}
