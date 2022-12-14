// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the list of ELB listeners.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			protocol := cfg.RequireObject("protocol")
//			_, err := Elb.GetListeners(ctx, &elb.GetListenersArgs{
//				Protocol: pulumi.StringRef(protocol),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetListeners(ctx *pulumi.Context, args *GetListenersArgs, opts ...pulumi.InvokeOption) (*GetListenersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetListenersResult
	err := ctx.Invoke("huaweicloud:Elb/getListeners:getListeners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListeners.
type GetListenersArgs struct {
	// The listener name.
	Name *string `pulumi:"name"`
	// The listener protocol.\
	// The valid values are **TCP**, **UDP**, **HTTP** and **TERMINATED_HTTPS**.
	Protocol *string `pulumi:"protocol"`
	// The front-end listening port of the listener.\
	// The valid value is range from `1` to `65535`.
	ProtocolPort *string `pulumi:"protocolPort"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getListeners.
type GetListenersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Listener list.
	// The object structure is documented below.
	Listeners []GetListenersListener `pulumi:"listeners"`
	// The listener name.
	Name *string `pulumi:"name"`
	// The listener protocol.
	Protocol *string `pulumi:"protocol"`
	// The front-end listening port of the listener.
	ProtocolPort *string `pulumi:"protocolPort"`
	Region       string  `pulumi:"region"`
}

func GetListenersOutput(ctx *pulumi.Context, args GetListenersOutputArgs, opts ...pulumi.InvokeOption) GetListenersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListenersResult, error) {
			args := v.(GetListenersArgs)
			r, err := GetListeners(ctx, &args, opts...)
			var s GetListenersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListenersResultOutput)
}

// A collection of arguments for invoking getListeners.
type GetListenersOutputArgs struct {
	// The listener name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The listener protocol.\
	// The valid values are **TCP**, **UDP**, **HTTP** and **TERMINATED_HTTPS**.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// The front-end listening port of the listener.\
	// The valid value is range from `1` to `65535`.
	ProtocolPort pulumi.StringPtrInput `pulumi:"protocolPort"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetListenersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenersArgs)(nil)).Elem()
}

// A collection of values returned by getListeners.
type GetListenersResultOutput struct{ *pulumi.OutputState }

func (GetListenersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenersResult)(nil)).Elem()
}

func (o GetListenersResultOutput) ToGetListenersResultOutput() GetListenersResultOutput {
	return o
}

func (o GetListenersResultOutput) ToGetListenersResultOutputWithContext(ctx context.Context) GetListenersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetListenersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenersResult) string { return v.Id }).(pulumi.StringOutput)
}

// Listener list.
// The object structure is documented below.
func (o GetListenersResultOutput) Listeners() GetListenersListenerArrayOutput {
	return o.ApplyT(func(v GetListenersResult) []GetListenersListener { return v.Listeners }).(GetListenersListenerArrayOutput)
}

// The listener name.
func (o GetListenersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The listener protocol.
func (o GetListenersResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenersResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The front-end listening port of the listener.
func (o GetListenersResultOutput) ProtocolPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenersResult) *string { return v.ProtocolPort }).(pulumi.StringPtrOutput)
}

func (o GetListenersResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenersResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListenersResultOutput{})
}
