// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of IAM identity providers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			providerName := cfg.RequireObject("providerName")
//			_, err := Iam.GetProviders(ctx, &iam.GetProvidersArgs{
//				Name: pulumi.StringRef(providerName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProviders(ctx *pulumi.Context, args *GetProvidersArgs, opts ...pulumi.InvokeOption) (*GetProvidersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProvidersResult
	err := ctx.Invoke("huaweicloud:Iam/getProviders:getProviders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProviders.
type GetProvidersArgs struct {
	// Specifies the name of the identity provider.
	Name *string `pulumi:"name"`
	// Specifies the single sign-on type of the identity provider.
	SsoType *string `pulumi:"ssoType"`
	// Specifies the status of the identity provider. The value can be **true** or **false**
	Status *string `pulumi:"status"`
}

// A collection of values returned by getProviders.
type GetProvidersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of identity providers.
	IdentityProviders []GetProvidersIdentityProvider `pulumi:"identityProviders"`
	Name              *string                        `pulumi:"name"`
	// The single sign-on type of the identity provider.
	SsoType *string `pulumi:"ssoType"`
	// The enabled status for the identity provider.
	Status *string `pulumi:"status"`
}

func GetProvidersOutput(ctx *pulumi.Context, args GetProvidersOutputArgs, opts ...pulumi.InvokeOption) GetProvidersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProvidersResult, error) {
			args := v.(GetProvidersArgs)
			r, err := GetProviders(ctx, &args, opts...)
			var s GetProvidersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProvidersResultOutput)
}

// A collection of arguments for invoking getProviders.
type GetProvidersOutputArgs struct {
	// Specifies the name of the identity provider.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the single sign-on type of the identity provider.
	SsoType pulumi.StringPtrInput `pulumi:"ssoType"`
	// Specifies the status of the identity provider. The value can be **true** or **false**
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetProvidersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProvidersArgs)(nil)).Elem()
}

// A collection of values returned by getProviders.
type GetProvidersResultOutput struct{ *pulumi.OutputState }

func (GetProvidersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProvidersResult)(nil)).Elem()
}

func (o GetProvidersResultOutput) ToGetProvidersResultOutput() GetProvidersResultOutput {
	return o
}

func (o GetProvidersResultOutput) ToGetProvidersResultOutputWithContext(ctx context.Context) GetProvidersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProvidersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProvidersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of identity providers.
func (o GetProvidersResultOutput) IdentityProviders() GetProvidersIdentityProviderArrayOutput {
	return o.ApplyT(func(v GetProvidersResult) []GetProvidersIdentityProvider { return v.IdentityProviders }).(GetProvidersIdentityProviderArrayOutput)
}

func (o GetProvidersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProvidersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The single sign-on type of the identity provider.
func (o GetProvidersResultOutput) SsoType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProvidersResult) *string { return v.SsoType }).(pulumi.StringPtrOutput)
}

// The enabled status for the identity provider.
func (o GetProvidersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProvidersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProvidersResultOutput{})
}
