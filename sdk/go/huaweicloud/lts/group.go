// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a log group resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Lts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lts.NewGroup(ctx, "test", &Lts.GroupArgs{
//				GroupName: pulumi.String("log_group1"),
//				TtlInDays: pulumi.Int(30),
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
// The log group can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Lts/group:Group test <id>
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// The creation time of the log group.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the enterprise project ID to which the log group belongs.
	// Changing this parameter will create a new resource.
	// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
	// will be used.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Specifies the log group name. Changing this parameter will create a new resource.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Specifies the region in which to create the log group resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the key/value pairs to associate with the log group.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the log expiration time(days).\
	// The value is range from `1` to `365`.
	TtlInDays pulumi.IntOutput `pulumi:"ttlInDays"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.TtlInDays == nil {
		return nil, errors.New("invalid value for required argument 'TtlInDays'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("huaweicloud:Lts/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("huaweicloud:Lts/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The creation time of the log group.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the enterprise project ID to which the log group belongs.
	// Changing this parameter will create a new resource.
	// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
	// will be used.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the log group name. Changing this parameter will create a new resource.
	GroupName *string `pulumi:"groupName"`
	// Specifies the region in which to create the log group resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs to associate with the log group.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the log expiration time(days).\
	// The value is range from `1` to `365`.
	TtlInDays *int `pulumi:"ttlInDays"`
}

type GroupState struct {
	// The creation time of the log group.
	CreatedAt pulumi.StringPtrInput
	// Specifies the enterprise project ID to which the log group belongs.
	// Changing this parameter will create a new resource.
	// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
	// will be used.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the log group name. Changing this parameter will create a new resource.
	GroupName pulumi.StringPtrInput
	// Specifies the region in which to create the log group resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the log group.
	Tags pulumi.StringMapInput
	// Specifies the log expiration time(days).\
	// The value is range from `1` to `365`.
	TtlInDays pulumi.IntPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Specifies the enterprise project ID to which the log group belongs.
	// Changing this parameter will create a new resource.
	// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
	// will be used.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the log group name. Changing this parameter will create a new resource.
	GroupName string `pulumi:"groupName"`
	// Specifies the region in which to create the log group resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs to associate with the log group.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the log expiration time(days).\
	// The value is range from `1` to `365`.
	TtlInDays int `pulumi:"ttlInDays"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Specifies the enterprise project ID to which the log group belongs.
	// Changing this parameter will create a new resource.
	// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
	// will be used.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the log group name. Changing this parameter will create a new resource.
	GroupName pulumi.StringInput
	// Specifies the region in which to create the log group resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the log group.
	Tags pulumi.StringMapInput
	// Specifies the log expiration time(days).\
	// The value is range from `1` to `365`.
	TtlInDays pulumi.IntInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The creation time of the log group.
func (o GroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the enterprise project ID to which the log group belongs.
// Changing this parameter will create a new resource.
// This parameter is valid only when the enterprise project function is enabled, if omitted, default enterprise project
// will be used.
func (o GroupOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Specifies the log group name. Changing this parameter will create a new resource.
func (o GroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Specifies the region in which to create the log group resource. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o GroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the key/value pairs to associate with the log group.
func (o GroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Group) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the log expiration time(days).\
// The value is range from `1` to `365`.
func (o GroupOutput) TtlInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Group) pulumi.IntOutput { return v.TtlInDays }).(pulumi.IntOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
