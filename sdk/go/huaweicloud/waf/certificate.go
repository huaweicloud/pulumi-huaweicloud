// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a WAF certificate resource within HuaweiCloud.
//
// > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
// used. The certificate resource can be used in Cloud Mode and Dedicated Mode.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			_, err := Waf.NewCertificate(ctx, "test", &Waf.CertificateArgs{
//				EnterpriseProjectId: pulumi.Any(enterpriseProjectId),
//				Certificate: pulumi.String(fmt.Sprintf(`-----BEGIN CERTIFICATE-----
//
// MIIFmQl5dh2QUAeo39TIKtadgAgh4zHx09kSgayS9Wph9LEqq7MA+2042L3J9aOa
// DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUR+SosWwALt6PkP0J9iOIxA6RW8gVsLwq
// ...
// +HhDvD/VeOHytX3RAs2GeTOtxyAV5XpKY5r+PkyUqPJj04t3d0Fopi0gNtLpMF=
// -----END CERTIFICATE-----
// `)),
//
//	PrivateKey: pulumi.String(fmt.Sprintf(`-----BEGIN PRIVATE KEY-----
//
// MIIJwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAM
// ATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2VuY3J5cHQu
// ...
// he8Y4IWS6wY7bCkjCWDcRQJMEhg76fsO3txE+FiYruq9RUWhiF1myv4Q6W+CyBFC
// 1qoJFlcDyqSMo5iHq3HLjs
// -----END PRIVATE KEY-----
// `)),
//
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// There are two ways to import WAF certificate state. * Using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/certificate:Certificate test <id>
//
// ```
//
//   - Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/certificate:Certificate test <id>/<enterprise_project_id>
//
// ```
//
//	Note that the imported state is not identical to your resource definition, due to security reason. The missing attributes include `certificate`, and `private_key`. You can ignore changes as below. hcl resource "huaweicloud_waf_certificate" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	certificate, private_key
//
//	]
//
//	} }
type Certificate struct {
	pulumi.CustomResourceState

	// Specifies the certificate content.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Indicates the time when the certificate uploaded, in RFC3339 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	// schema: Deprecated; The certificate expiration time.
	//
	// Deprecated: Use 'expired_at' instead.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// Indicates the time when the certificate expires, in RFC3339 format.
	ExpiredAt pulumi.StringOutput `pulumi:"expiredAt"`
	// Specifies the certificate name. The maximum length is `256` characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the private key. This field does not support individual editing.
	// Changes to this field will only take effect when `certificate` changes.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// Specifies the region in which to create the WAF certificate. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("huaweicloud:Waf/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("huaweicloud:Waf/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Specifies the certificate content.
	Certificate *string `pulumi:"certificate"`
	// Indicates the time when the certificate uploaded, in RFC3339 format.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// schema: Deprecated; The certificate expiration time.
	//
	// Deprecated: Use 'expired_at' instead.
	Expiration *string `pulumi:"expiration"`
	// Indicates the time when the certificate expires, in RFC3339 format.
	ExpiredAt *string `pulumi:"expiredAt"`
	// Specifies the certificate name. The maximum length is `256` characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the private key. This field does not support individual editing.
	// Changes to this field will only take effect when `certificate` changes.
	PrivateKey *string `pulumi:"privateKey"`
	// Specifies the region in which to create the WAF certificate. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

type CertificateState struct {
	// Specifies the certificate content.
	Certificate pulumi.StringPtrInput
	// Indicates the time when the certificate uploaded, in RFC3339 format.
	CreatedAt pulumi.StringPtrInput
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// schema: Deprecated; The certificate expiration time.
	//
	// Deprecated: Use 'expired_at' instead.
	Expiration pulumi.StringPtrInput
	// Indicates the time when the certificate expires, in RFC3339 format.
	ExpiredAt pulumi.StringPtrInput
	// Specifies the certificate name. The maximum length is `256` characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the private key. This field does not support individual editing.
	// Changes to this field will only take effect when `certificate` changes.
	PrivateKey pulumi.StringPtrInput
	// Specifies the region in which to create the WAF certificate. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Specifies the certificate content.
	Certificate string `pulumi:"certificate"`
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the certificate name. The maximum length is `256` characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the private key. This field does not support individual editing.
	// Changes to this field will only take effect when `certificate` changes.
	PrivateKey string `pulumi:"privateKey"`
	// Specifies the region in which to create the WAF certificate. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Specifies the certificate content.
	Certificate pulumi.StringInput
	// Specifies the enterprise project ID of WAF certificate.
	// For enterprise users, if omitted, default enterprise project will be used.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the certificate name. The maximum length is `256` characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the private key. This field does not support individual editing.
	// Changes to this field will only take effect when `certificate` changes.
	PrivateKey pulumi.StringInput
	// Specifies the region in which to create the WAF certificate. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Specifies the certificate content.
func (o CertificateOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Indicates the time when the certificate uploaded, in RFC3339 format.
func (o CertificateOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the enterprise project ID of WAF certificate.
// For enterprise users, if omitted, default enterprise project will be used.
// Changing this parameter will create a new resource.
func (o CertificateOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// schema: Deprecated; The certificate expiration time.
//
// Deprecated: Use 'expired_at' instead.
func (o CertificateOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// Indicates the time when the certificate expires, in RFC3339 format.
func (o CertificateOutput) ExpiredAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpiredAt }).(pulumi.StringOutput)
}

// Specifies the certificate name. The maximum length is `256` characters. Only digits,
// letters, underscores(_), and hyphens(-) are allowed.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the private key. This field does not support individual editing.
// Changes to this field will only take effect when `certificate` changes.
func (o CertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// Specifies the region in which to create the WAF certificate. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o CertificateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
