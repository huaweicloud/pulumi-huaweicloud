// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb mysql dedicated resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDB.GetMysqlDedicatedResource(ctx, &gaussdb.GetMysqlDedicatedResourceArgs{
//				ResourceName: pulumi.StringRef("test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMysqlDedicatedResource(ctx *pulumi.Context, args *GetMysqlDedicatedResourceArgs, opts ...pulumi.InvokeOption) (*GetMysqlDedicatedResourceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMysqlDedicatedResourceResult
	err := ctx.Invoke("huaweicloud:GaussDB/getMysqlDedicatedResource:getMysqlDedicatedResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlDedicatedResource.
type GetMysqlDedicatedResourceArgs struct {
	// The region in which to obtain the dedicated resource. If omitted, the provider-level
	// region will be used.
	Region *string `pulumi:"region"`
	// Specifies the dedicated resource name.
	ResourceName *string `pulumi:"resourceName"`
}

// A collection of values returned by getMysqlDedicatedResource.
type GetMysqlDedicatedResourceResult struct {
	// Indicates the architecture of the dedicated resource.
	Architecture string `pulumi:"architecture"`
	// Indicates the availability zone of the dedicated resource.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates the ram size of the dedicated resource.
	Ram          int    `pulumi:"ram"`
	Region       string `pulumi:"region"`
	ResourceName string `pulumi:"resourceName"`
	// Indicates the status of the dedicated resource.
	Status string `pulumi:"status"`
	// Indicates the vcpus count of the dedicated resource.
	Vcpus int `pulumi:"vcpus"`
	// Indicates the volume size of the dedicated resource.
	Volume int `pulumi:"volume"`
}

func GetMysqlDedicatedResourceOutput(ctx *pulumi.Context, args GetMysqlDedicatedResourceOutputArgs, opts ...pulumi.InvokeOption) GetMysqlDedicatedResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMysqlDedicatedResourceResult, error) {
			args := v.(GetMysqlDedicatedResourceArgs)
			r, err := GetMysqlDedicatedResource(ctx, &args, opts...)
			var s GetMysqlDedicatedResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMysqlDedicatedResourceResultOutput)
}

// A collection of arguments for invoking getMysqlDedicatedResource.
type GetMysqlDedicatedResourceOutputArgs struct {
	// The region in which to obtain the dedicated resource. If omitted, the provider-level
	// region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the dedicated resource name.
	ResourceName pulumi.StringPtrInput `pulumi:"resourceName"`
}

func (GetMysqlDedicatedResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlDedicatedResourceArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlDedicatedResource.
type GetMysqlDedicatedResourceResultOutput struct{ *pulumi.OutputState }

func (GetMysqlDedicatedResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlDedicatedResourceResult)(nil)).Elem()
}

func (o GetMysqlDedicatedResourceResultOutput) ToGetMysqlDedicatedResourceResultOutput() GetMysqlDedicatedResourceResultOutput {
	return o
}

func (o GetMysqlDedicatedResourceResultOutput) ToGetMysqlDedicatedResourceResultOutputWithContext(ctx context.Context) GetMysqlDedicatedResourceResultOutput {
	return o
}

// Indicates the architecture of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// Indicates the availability zone of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMysqlDedicatedResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates the ram size of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) int { return v.Ram }).(pulumi.IntOutput)
}

func (o GetMysqlDedicatedResourceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetMysqlDedicatedResourceResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

// Indicates the status of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

// Indicates the vcpus count of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) Vcpus() pulumi.IntOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) int { return v.Vcpus }).(pulumi.IntOutput)
}

// Indicates the volume size of the dedicated resource.
func (o GetMysqlDedicatedResourceResultOutput) Volume() pulumi.IntOutput {
	return o.ApplyT(func(v GetMysqlDedicatedResourceResult) int { return v.Volume }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMysqlDedicatedResourceResultOutput{})
}
