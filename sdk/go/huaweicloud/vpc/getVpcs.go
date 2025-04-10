// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of VPC.
func GetVpcs(ctx *pulumi.Context, args *GetVpcsArgs, opts ...pulumi.InvokeOption) (*GetVpcsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetVpcsResult
	err := ctx.Invoke("huaweicloud:Vpc/getVpcs:getVpcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcs.
type GetVpcsArgs struct {
	// Specifies the cidr block of the desired VPC.
	Cidr *string `pulumi:"cidr"`
	// Specifies the enterprise project ID which the desired VPC belongs to.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the id of the desired VPC.
	Id *string `pulumi:"id"`
	// Specifies the name of the desired VPC. The value is a string of no more than 64 characters
	// and can contain digits, letters, underscores (_) and hyphens (-).
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain the VPC. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
	// Specifies the current status of the desired VPC. The value can be CREATING, OK or ERROR.
	Status *string `pulumi:"status"`
	// Specifies the included key/value pairs which associated with the desired VPC.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcs.
type GetVpcsResult struct {
	// The cidr block of the VPC.
	Cidr *string `pulumi:"cidr"`
	// The the enterprise project ID of the VPC.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The ID of the VPC.
	Id string `pulumi:"id"`
	// The name of the VPC.
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
	// The current status of the VPC.
	Status *string `pulumi:"status"`
	// The key/value pairs which associated with the VPC.
	Tags map[string]string `pulumi:"tags"`
	// The list of all VPCs found. Structure is documented below.
	Vpcs []GetVpcsVpc `pulumi:"vpcs"`
}

func GetVpcsOutput(ctx *pulumi.Context, args GetVpcsOutputArgs, opts ...pulumi.InvokeOption) GetVpcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcsResult, error) {
			args := v.(GetVpcsArgs)
			r, err := GetVpcs(ctx, &args, opts...)
			var s GetVpcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcsResultOutput)
}

// A collection of arguments for invoking getVpcs.
type GetVpcsOutputArgs struct {
	// Specifies the cidr block of the desired VPC.
	Cidr pulumi.StringPtrInput `pulumi:"cidr"`
	// Specifies the enterprise project ID which the desired VPC belongs to.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the id of the desired VPC.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies the name of the desired VPC. The value is a string of no more than 64 characters
	// and can contain digits, letters, underscores (_) and hyphens (-).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain the VPC. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the current status of the desired VPC. The value can be CREATING, OK or ERROR.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Specifies the included key/value pairs which associated with the desired VPC.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetVpcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcs.
type GetVpcsResultOutput struct{ *pulumi.OutputState }

func (GetVpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsResult)(nil)).Elem()
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutput() GetVpcsResultOutput {
	return o
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutputWithContext(ctx context.Context) GetVpcsResultOutput {
	return o
}

// The cidr block of the VPC.
func (o GetVpcsResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

// The the enterprise project ID of the VPC.
func (o GetVpcsResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The ID of the VPC.
func (o GetVpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the VPC.
func (o GetVpcsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetVpcsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.Region }).(pulumi.StringOutput)
}

// The current status of the VPC.
func (o GetVpcsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The key/value pairs which associated with the VPC.
func (o GetVpcsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetVpcsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The list of all VPCs found. Structure is documented below.
func (o GetVpcsResultOutput) Vpcs() GetVpcsVpcArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []GetVpcsVpc { return v.Vpcs }).(GetVpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcsResultOutput{})
}
