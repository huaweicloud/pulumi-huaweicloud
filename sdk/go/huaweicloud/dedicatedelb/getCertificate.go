// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the certificate in Dedicated Load Balance (Dedicated ELB).
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
//			certificateName := cfg.RequireObject("certificateName")
//			_, err := DedicatedElb.GetCertificate(ctx, &dedicatedelb.GetCertificateArgs{
//				Name: certificateName,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("huaweicloud:DedicatedElb/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The name of certificate. The value is case sensitive and does not supports fuzzy matching.
	Name string `pulumi:"name"`
	// The region in which to obtain the Dedicated ELB certificate. If omitted, the
	// provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	// Human-readable description for the Certificate.
	Description string `pulumi:"description"`
	// The domain of the Certificate. This parameter is valid only when `type` is "server".
	Domain string `pulumi:"domain"`
	// Indicates the time when the certificate expires.
	Expiration string `pulumi:"expiration"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
	// Specifies the certificate type. The value can be one of the following:
	// + `server`: indicates the server certificate.
	// + `client`: indicates the CA certificate.
	Type string `pulumi:"type"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateOutputArgs struct {
	// The name of certificate. The value is case sensitive and does not supports fuzzy matching.
	Name pulumi.StringInput `pulumi:"name"`
	// The region in which to obtain the Dedicated ELB certificate. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getCertificate.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

// Human-readable description for the Certificate.
func (o LookupCertificateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Description }).(pulumi.StringOutput)
}

// The domain of the Certificate. This parameter is valid only when `type` is "server".
func (o LookupCertificateResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Domain }).(pulumi.StringOutput)
}

// Indicates the time when the certificate expires.
func (o LookupCertificateResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Region }).(pulumi.StringOutput)
}

// Specifies the certificate type. The value can be one of the following:
// + `server`: indicates the server certificate.
// + `client`: indicates the CA certificate.
func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
