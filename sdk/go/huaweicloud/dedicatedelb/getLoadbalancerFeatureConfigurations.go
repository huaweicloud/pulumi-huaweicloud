// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of feature configurations of a load balancer.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			loadbalancerId := cfg.RequireObject("loadbalancerId")
//			_, err := DedicatedElb.GetLoadbalancerFeatureConfigurations(ctx, &dedicatedelb.GetLoadbalancerFeatureConfigurationsArgs{
//				LoadbalancerId: loadbalancerId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLoadbalancerFeatureConfigurations(ctx *pulumi.Context, args *GetLoadbalancerFeatureConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetLoadbalancerFeatureConfigurationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLoadbalancerFeatureConfigurationsResult
	err := ctx.Invoke("huaweicloud:DedicatedElb/getLoadbalancerFeatureConfigurations:getLoadbalancerFeatureConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadbalancerFeatureConfigurations.
type GetLoadbalancerFeatureConfigurationsArgs struct {
	// Specifies the load balancer ID.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getLoadbalancerFeatureConfigurations.
type GetLoadbalancerFeatureConfigurationsResult struct {
	// Specifies the load balancer feature information list.
	Features []GetLoadbalancerFeatureConfigurationsFeature `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	LoadbalancerId string `pulumi:"loadbalancerId"`
	Region         string `pulumi:"region"`
}

func GetLoadbalancerFeatureConfigurationsOutput(ctx *pulumi.Context, args GetLoadbalancerFeatureConfigurationsOutputArgs, opts ...pulumi.InvokeOption) GetLoadbalancerFeatureConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLoadbalancerFeatureConfigurationsResult, error) {
			args := v.(GetLoadbalancerFeatureConfigurationsArgs)
			r, err := GetLoadbalancerFeatureConfigurations(ctx, &args, opts...)
			var s GetLoadbalancerFeatureConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLoadbalancerFeatureConfigurationsResultOutput)
}

// A collection of arguments for invoking getLoadbalancerFeatureConfigurations.
type GetLoadbalancerFeatureConfigurationsOutputArgs struct {
	// Specifies the load balancer ID.
	LoadbalancerId pulumi.StringInput `pulumi:"loadbalancerId"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetLoadbalancerFeatureConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadbalancerFeatureConfigurationsArgs)(nil)).Elem()
}

// A collection of values returned by getLoadbalancerFeatureConfigurations.
type GetLoadbalancerFeatureConfigurationsResultOutput struct{ *pulumi.OutputState }

func (GetLoadbalancerFeatureConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadbalancerFeatureConfigurationsResult)(nil)).Elem()
}

func (o GetLoadbalancerFeatureConfigurationsResultOutput) ToGetLoadbalancerFeatureConfigurationsResultOutput() GetLoadbalancerFeatureConfigurationsResultOutput {
	return o
}

func (o GetLoadbalancerFeatureConfigurationsResultOutput) ToGetLoadbalancerFeatureConfigurationsResultOutputWithContext(ctx context.Context) GetLoadbalancerFeatureConfigurationsResultOutput {
	return o
}

// Specifies the load balancer feature information list.
func (o GetLoadbalancerFeatureConfigurationsResultOutput) Features() GetLoadbalancerFeatureConfigurationsFeatureArrayOutput {
	return o.ApplyT(func(v GetLoadbalancerFeatureConfigurationsResult) []GetLoadbalancerFeatureConfigurationsFeature {
		return v.Features
	}).(GetLoadbalancerFeatureConfigurationsFeatureArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadbalancerFeatureConfigurationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadbalancerFeatureConfigurationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLoadbalancerFeatureConfigurationsResultOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadbalancerFeatureConfigurationsResult) string { return v.LoadbalancerId }).(pulumi.StringOutput)
}

func (o GetLoadbalancerFeatureConfigurationsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadbalancerFeatureConfigurationsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadbalancerFeatureConfigurationsResultOutput{})
}
