// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdbforopengauss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available GaussDB OpenGauss instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforOpenGauss"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforOpenGauss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDBforOpenGauss.GetOpengaussInstances(ctx, &gaussdbforopengauss.GetOpengaussInstancesArgs{
//				Name: pulumi.StringRef("gaussdb-instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOpengaussInstances(ctx *pulumi.Context, args *GetOpengaussInstancesArgs, opts ...pulumi.InvokeOption) (*GetOpengaussInstancesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOpengaussInstancesResult
	err := ctx.Invoke("huaweicloud:GaussDBforOpenGauss/getOpengaussInstances:getOpengaussInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOpengaussInstances.
type GetOpengaussInstancesArgs struct {
	// Specifies the name of the instance.
	Name *string `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region *string `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getOpengaussInstances.
type GetOpengaussInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An array of available instances.
	Instances []GetOpengaussInstancesInstance `pulumi:"instances"`
	// Indicates the node name.
	Name *string `pulumi:"name"`
	// The region of the instance.
	Region string `pulumi:"region"`
	// Indicates the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Indicates the VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

func GetOpengaussInstancesOutput(ctx *pulumi.Context, args GetOpengaussInstancesOutputArgs, opts ...pulumi.InvokeOption) GetOpengaussInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOpengaussInstancesResult, error) {
			args := v.(GetOpengaussInstancesArgs)
			r, err := GetOpengaussInstances(ctx, &args, opts...)
			var s GetOpengaussInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOpengaussInstancesResultOutput)
}

// A collection of arguments for invoking getOpengaussInstances.
type GetOpengaussInstancesOutputArgs struct {
	// Specifies the name of the instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetOpengaussInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpengaussInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getOpengaussInstances.
type GetOpengaussInstancesResultOutput struct{ *pulumi.OutputState }

func (GetOpengaussInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpengaussInstancesResult)(nil)).Elem()
}

func (o GetOpengaussInstancesResultOutput) ToGetOpengaussInstancesResultOutput() GetOpengaussInstancesResultOutput {
	return o
}

func (o GetOpengaussInstancesResultOutput) ToGetOpengaussInstancesResultOutputWithContext(ctx context.Context) GetOpengaussInstancesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOpengaussInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// An array of available instances.
func (o GetOpengaussInstancesResultOutput) Instances() GetOpengaussInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) []GetOpengaussInstancesInstance { return v.Instances }).(GetOpengaussInstancesInstanceArrayOutput)
}

// Indicates the node name.
func (o GetOpengaussInstancesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The region of the instance.
func (o GetOpengaussInstancesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the network ID of a subnet.
func (o GetOpengaussInstancesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Indicates the VPC ID.
func (o GetOpengaussInstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpengaussInstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOpengaussInstancesResultOutput{})
}
