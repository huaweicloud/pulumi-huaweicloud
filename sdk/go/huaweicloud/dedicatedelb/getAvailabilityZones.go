// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of available AZs when create a load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DedicatedElb.GetAvailabilityZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("huaweicloud:DedicatedElb/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// Specifies the load balancer ID.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the public border group.
	PublicBorderGroup *string `pulumi:"publicBorderGroup"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	// Indicates the AZs that are available during load balancer creation. For example, in [az1,az2] and
	// [az2,az3] sets, you can select **az1** and **az2** or **az2** and **az3**, but cannot select **az1** and **az3**.
	// The availabilityZones structure is documented below.
	AvailabilityZones []GetAvailabilityZonesAvailabilityZone `pulumi:"availabilityZones"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Indicates the public border group, for example, **center**.
	PublicBorderGroup *string `pulumi:"publicBorderGroup"`
	Region            string  `pulumi:"region"`
}

func GetAvailabilityZonesOutput(ctx *pulumi.Context, args GetAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZonesResult, error) {
			args := v.(GetAvailabilityZonesArgs)
			r, err := GetAvailabilityZones(ctx, &args, opts...)
			var s GetAvailabilityZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesOutputArgs struct {
	// Specifies the load balancer ID.
	LoadbalancerId pulumi.StringPtrInput `pulumi:"loadbalancerId"`
	// Specifies the public border group.
	PublicBorderGroup pulumi.StringPtrInput `pulumi:"publicBorderGroup"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesResult)(nil)).Elem()
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutput() GetAvailabilityZonesResultOutput {
	return o
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutputWithContext(ctx context.Context) GetAvailabilityZonesResultOutput {
	return o
}

// Indicates the AZs that are available during load balancer creation. For example, in [az1,az2] and
// [az2,az3] sets, you can select **az1** and **az2** or **az2** and **az3**, but cannot select **az1** and **az3**.
// The availabilityZones structure is documented below.
func (o GetAvailabilityZonesResultOutput) AvailabilityZones() GetAvailabilityZonesAvailabilityZoneArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []GetAvailabilityZonesAvailabilityZone { return v.AvailabilityZones }).(GetAvailabilityZonesAvailabilityZoneArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAvailabilityZonesResultOutput) LoadbalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *string { return v.LoadbalancerId }).(pulumi.StringPtrOutput)
}

// Indicates the public border group, for example, **center**.
func (o GetAvailabilityZonesResultOutput) PublicBorderGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *string { return v.PublicBorderGroup }).(pulumi.StringPtrOutput)
}

func (o GetAvailabilityZonesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZonesResultOutput{})
}
