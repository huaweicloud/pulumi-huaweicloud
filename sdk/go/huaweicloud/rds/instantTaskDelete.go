// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS instant task delete resource within HuaweiCloud.
//
// > This resource is only a one-time action resource for operating the API.
// Deleting this resource will not clear the corresponding request record,
// but will only remove the resource information from the tfstate file.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			jobId := cfg.RequireObject("jobId")
//			_, err := Rds.NewInstantTaskDelete(ctx, "test", &Rds.InstantTaskDeleteArgs{
//				JobId: pulumi.Any(jobId),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type InstantTaskDelete struct {
	pulumi.CustomResourceState

	EnableForceNew pulumi.StringPtrOutput `pulumi:"enableForceNew"`
	// Specifies the task ID.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewInstantTaskDelete registers a new resource with the given unique name, arguments, and options.
func NewInstantTaskDelete(ctx *pulumi.Context,
	name string, args *InstantTaskDeleteArgs, opts ...pulumi.ResourceOption) (*InstantTaskDelete, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstantTaskDelete
	err := ctx.RegisterResource("huaweicloud:Rds/instantTaskDelete:InstantTaskDelete", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstantTaskDelete gets an existing InstantTaskDelete resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstantTaskDelete(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstantTaskDeleteState, opts ...pulumi.ResourceOption) (*InstantTaskDelete, error) {
	var resource InstantTaskDelete
	err := ctx.ReadResource("huaweicloud:Rds/instantTaskDelete:InstantTaskDelete", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstantTaskDelete resources.
type instantTaskDeleteState struct {
	EnableForceNew *string `pulumi:"enableForceNew"`
	// Specifies the task ID.
	JobId *string `pulumi:"jobId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

type InstantTaskDeleteState struct {
	EnableForceNew pulumi.StringPtrInput
	// Specifies the task ID.
	JobId pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (InstantTaskDeleteState) ElementType() reflect.Type {
	return reflect.TypeOf((*instantTaskDeleteState)(nil)).Elem()
}

type instantTaskDeleteArgs struct {
	EnableForceNew *string `pulumi:"enableForceNew"`
	// Specifies the task ID.
	JobId string `pulumi:"jobId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a InstantTaskDelete resource.
type InstantTaskDeleteArgs struct {
	EnableForceNew pulumi.StringPtrInput
	// Specifies the task ID.
	JobId pulumi.StringInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (InstantTaskDeleteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instantTaskDeleteArgs)(nil)).Elem()
}

type InstantTaskDeleteInput interface {
	pulumi.Input

	ToInstantTaskDeleteOutput() InstantTaskDeleteOutput
	ToInstantTaskDeleteOutputWithContext(ctx context.Context) InstantTaskDeleteOutput
}

func (*InstantTaskDelete) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantTaskDelete)(nil)).Elem()
}

func (i *InstantTaskDelete) ToInstantTaskDeleteOutput() InstantTaskDeleteOutput {
	return i.ToInstantTaskDeleteOutputWithContext(context.Background())
}

func (i *InstantTaskDelete) ToInstantTaskDeleteOutputWithContext(ctx context.Context) InstantTaskDeleteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantTaskDeleteOutput)
}

// InstantTaskDeleteArrayInput is an input type that accepts InstantTaskDeleteArray and InstantTaskDeleteArrayOutput values.
// You can construct a concrete instance of `InstantTaskDeleteArrayInput` via:
//
//	InstantTaskDeleteArray{ InstantTaskDeleteArgs{...} }
type InstantTaskDeleteArrayInput interface {
	pulumi.Input

	ToInstantTaskDeleteArrayOutput() InstantTaskDeleteArrayOutput
	ToInstantTaskDeleteArrayOutputWithContext(context.Context) InstantTaskDeleteArrayOutput
}

type InstantTaskDeleteArray []InstantTaskDeleteInput

func (InstantTaskDeleteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstantTaskDelete)(nil)).Elem()
}

func (i InstantTaskDeleteArray) ToInstantTaskDeleteArrayOutput() InstantTaskDeleteArrayOutput {
	return i.ToInstantTaskDeleteArrayOutputWithContext(context.Background())
}

func (i InstantTaskDeleteArray) ToInstantTaskDeleteArrayOutputWithContext(ctx context.Context) InstantTaskDeleteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantTaskDeleteArrayOutput)
}

// InstantTaskDeleteMapInput is an input type that accepts InstantTaskDeleteMap and InstantTaskDeleteMapOutput values.
// You can construct a concrete instance of `InstantTaskDeleteMapInput` via:
//
//	InstantTaskDeleteMap{ "key": InstantTaskDeleteArgs{...} }
type InstantTaskDeleteMapInput interface {
	pulumi.Input

	ToInstantTaskDeleteMapOutput() InstantTaskDeleteMapOutput
	ToInstantTaskDeleteMapOutputWithContext(context.Context) InstantTaskDeleteMapOutput
}

type InstantTaskDeleteMap map[string]InstantTaskDeleteInput

func (InstantTaskDeleteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstantTaskDelete)(nil)).Elem()
}

func (i InstantTaskDeleteMap) ToInstantTaskDeleteMapOutput() InstantTaskDeleteMapOutput {
	return i.ToInstantTaskDeleteMapOutputWithContext(context.Background())
}

func (i InstantTaskDeleteMap) ToInstantTaskDeleteMapOutputWithContext(ctx context.Context) InstantTaskDeleteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantTaskDeleteMapOutput)
}

type InstantTaskDeleteOutput struct{ *pulumi.OutputState }

func (InstantTaskDeleteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantTaskDelete)(nil)).Elem()
}

func (o InstantTaskDeleteOutput) ToInstantTaskDeleteOutput() InstantTaskDeleteOutput {
	return o
}

func (o InstantTaskDeleteOutput) ToInstantTaskDeleteOutputWithContext(ctx context.Context) InstantTaskDeleteOutput {
	return o
}

func (o InstantTaskDeleteOutput) EnableForceNew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstantTaskDelete) pulumi.StringPtrOutput { return v.EnableForceNew }).(pulumi.StringPtrOutput)
}

// Specifies the task ID.
func (o InstantTaskDeleteOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstantTaskDelete) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o InstantTaskDeleteOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InstantTaskDelete) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type InstantTaskDeleteArrayOutput struct{ *pulumi.OutputState }

func (InstantTaskDeleteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstantTaskDelete)(nil)).Elem()
}

func (o InstantTaskDeleteArrayOutput) ToInstantTaskDeleteArrayOutput() InstantTaskDeleteArrayOutput {
	return o
}

func (o InstantTaskDeleteArrayOutput) ToInstantTaskDeleteArrayOutputWithContext(ctx context.Context) InstantTaskDeleteArrayOutput {
	return o
}

func (o InstantTaskDeleteArrayOutput) Index(i pulumi.IntInput) InstantTaskDeleteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstantTaskDelete {
		return vs[0].([]*InstantTaskDelete)[vs[1].(int)]
	}).(InstantTaskDeleteOutput)
}

type InstantTaskDeleteMapOutput struct{ *pulumi.OutputState }

func (InstantTaskDeleteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstantTaskDelete)(nil)).Elem()
}

func (o InstantTaskDeleteMapOutput) ToInstantTaskDeleteMapOutput() InstantTaskDeleteMapOutput {
	return o
}

func (o InstantTaskDeleteMapOutput) ToInstantTaskDeleteMapOutputWithContext(ctx context.Context) InstantTaskDeleteMapOutput {
	return o
}

func (o InstantTaskDeleteMapOutput) MapIndex(k pulumi.StringInput) InstantTaskDeleteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstantTaskDelete {
		return vs[0].(map[string]*InstantTaskDelete)[vs[1].(string)]
	}).(InstantTaskDeleteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstantTaskDeleteInput)(nil)).Elem(), &InstantTaskDelete{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstantTaskDeleteArrayInput)(nil)).Elem(), InstantTaskDeleteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstantTaskDeleteMapInput)(nil)).Elem(), InstantTaskDeleteMap{})
	pulumi.RegisterOutputType(InstantTaskDeleteOutput{})
	pulumi.RegisterOutputType(InstantTaskDeleteArrayOutput{})
	pulumi.RegisterOutputType(InstantTaskDeleteMapOutput{})
}
