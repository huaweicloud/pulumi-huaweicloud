// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the certificate of WAF within HuaweiCloud.
//
// > When multiple pieces of data are queried, the datasource will process the first piece of data and put it back.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			_, err := Waf.GetCertificate(ctx, &waf.GetCertificateArgs{
//				Name:                pulumi.StringRef("test-name"),
//				EnterpriseProjectId: pulumi.StringRef(enterpriseProjectId),
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
	err := ctx.Invoke("huaweicloud:Waf/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the certificate expiration status. The options are as follows:
	// + `0`: Not expired;
	// + `1`: Expired;
	// + `2`: Expired soon (The certificate will expire in one month.)
	ExpirationStatus *string `pulumi:"expirationStatus"`
	// Deprecated: Use 'expiration_status' instead.
	ExpireStatus *int `pulumi:"expireStatus"`
	// Specifies the name of certificate. The value is case-sensitive and supports fuzzy matching.
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain the WAF. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	// Indicates the time when the certificate uploaded, in RFC3339 format.
	CreatedAt           string `pulumi:"createdAt"`
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// Deprecated: Use 'expired_at' instead.
	Expiration       string `pulumi:"expiration"`
	ExpirationStatus string `pulumi:"expirationStatus"`
	// Deprecated: Use 'expiration_status' instead.
	ExpireStatus *int `pulumi:"expireStatus"`
	// Indicates the time when the certificate expires, in RFC3339 format.
	ExpiredAt string `pulumi:"expiredAt"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
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
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the certificate expiration status. The options are as follows:
	// + `0`: Not expired;
	// + `1`: Expired;
	// + `2`: Expired soon (The certificate will expire in one month.)
	ExpirationStatus pulumi.StringPtrInput `pulumi:"expirationStatus"`
	// Deprecated: Use 'expiration_status' instead.
	ExpireStatus pulumi.IntPtrInput `pulumi:"expireStatus"`
	// Specifies the name of certificate. The value is case-sensitive and supports fuzzy matching.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain the WAF. If omitted, the provider-level region
	// will be used.
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

// Indicates the time when the certificate uploaded, in RFC3339 format.
func (o LookupCertificateResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Deprecated: Use 'expired_at' instead.
func (o LookupCertificateResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Expiration }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ExpirationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpirationStatus }).(pulumi.StringOutput)
}

// Deprecated: Use 'expiration_status' instead.
func (o LookupCertificateResultOutput) ExpireStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *int { return v.ExpireStatus }).(pulumi.IntPtrOutput)
}

// Indicates the time when the certificate expires, in RFC3339 format.
func (o LookupCertificateResultOutput) ExpiredAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpiredAt }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
