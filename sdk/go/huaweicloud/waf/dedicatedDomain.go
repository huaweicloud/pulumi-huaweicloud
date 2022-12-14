// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a dedicated mode domain resource within HuaweiCloud.
//
// > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
// used. The dedicated mode domain name resource can be used in Dedicated Mode and ELB Mode.
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
//			certificatedId := cfg.RequireObject("certificatedId")
//			vpcId := cfg.RequireObject("vpcId")
//			dedicatedEngineId := cfg.RequireObject("dedicatedEngineId")
//			_, err := Waf.NewDedicatedDomain(ctx, "domain1", &Waf.DedicatedDomainArgs{
//				Domain:        pulumi.String("www.example.com"),
//				CertificateId: pulumi.Any(huaweicloud_waf_certificate.Certificate_1.Id),
//				Servers: waf.DedicatedDomainServerArray{
//					&waf.DedicatedDomainServerArgs{
//						ClientProtocol: pulumi.String("HTTPS"),
//						ServerProtocol: pulumi.String("HTTP"),
//						Address:        pulumi.String("192.168.1.100"),
//						Port:           pulumi.Int(8080),
//						Type:           pulumi.String("ipv4"),
//						VpcId:          pulumi.Any(vpcId),
//					},
//				},
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
// Dedicated mode domain can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/dedicatedDomain:DedicatedDomain domain_1 69e9a86becb4424298cc6bdeacbf69d5
//
// ```
type DedicatedDomain struct {
	pulumi.CustomResourceState

	// Whether a domain name is connected to WAF. Valid values are:
	AccessStatus pulumi.IntOutput `pulumi:"accessStatus"`
	// The alarm page of domain. Valid values are:
	AlarmPage pulumi.StringMapOutput `pulumi:"alarmPage"`
	// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
	// is set to HTTPS.
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// The name of the certificate used by the domain name.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	Cipher          pulumi.StringOutput `pulumi:"cipher"`
	// The compliance certifications of the domain, values are:
	ComplianceCertification pulumi.BoolMapOutput `pulumi:"complianceCertification"`
	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to `true`.
	KeepPolicy pulumi.BoolPtrOutput `pulumi:"keepPolicy"`
	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The protection status of domain, `0`: suspended, `1`: enabled.
	// Default value is `1`.
	ProtectStatus pulumi.IntOutput `pulumi:"protectStatus"`
	// The protocol type of the client. The options are `HTTP` and `HTTPS`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Specifies whether a proxy is configured. Default value is `false`.
	Proxy pulumi.BoolPtrOutput `pulumi:"proxy"`
	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region pulumi.StringOutput `pulumi:"region"`
	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	Servers DedicatedDomainServerArrayOutput `pulumi:"servers"`
	// The TLS configuration of domain.
	Tls pulumi.StringOutput `pulumi:"tls"`
	// The traffic identifier of domain. Valid values are:
	TrafficIdentifier pulumi.StringMapOutput `pulumi:"trafficIdentifier"`
}

// NewDedicatedDomain registers a new resource with the given unique name, arguments, and options.
func NewDedicatedDomain(ctx *pulumi.Context,
	name string, args *DedicatedDomainArgs, opts ...pulumi.ResourceOption) (*DedicatedDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Servers == nil {
		return nil, errors.New("invalid value for required argument 'Servers'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DedicatedDomain
	err := ctx.RegisterResource("huaweicloud:Waf/dedicatedDomain:DedicatedDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedDomain gets an existing DedicatedDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedDomainState, opts ...pulumi.ResourceOption) (*DedicatedDomain, error) {
	var resource DedicatedDomain
	err := ctx.ReadResource("huaweicloud:Waf/dedicatedDomain:DedicatedDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedDomain resources.
type dedicatedDomainState struct {
	// Whether a domain name is connected to WAF. Valid values are:
	AccessStatus *int `pulumi:"accessStatus"`
	// The alarm page of domain. Valid values are:
	AlarmPage map[string]string `pulumi:"alarmPage"`
	// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
	// is set to HTTPS.
	CertificateId *string `pulumi:"certificateId"`
	// The name of the certificate used by the domain name.
	CertificateName *string `pulumi:"certificateName"`
	Cipher          *string `pulumi:"cipher"`
	// The compliance certifications of the domain, values are:
	ComplianceCertification map[string]bool `pulumi:"complianceCertification"`
	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	Domain *string `pulumi:"domain"`
	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to `true`.
	KeepPolicy *bool `pulumi:"keepPolicy"`
	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyId *string `pulumi:"policyId"`
	// The protection status of domain, `0`: suspended, `1`: enabled.
	// Default value is `1`.
	ProtectStatus *int `pulumi:"protectStatus"`
	// The protocol type of the client. The options are `HTTP` and `HTTPS`.
	Protocol *string `pulumi:"protocol"`
	// Specifies whether a proxy is configured. Default value is `false`.
	Proxy *bool `pulumi:"proxy"`
	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region *string `pulumi:"region"`
	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	Servers []DedicatedDomainServer `pulumi:"servers"`
	// The TLS configuration of domain.
	Tls *string `pulumi:"tls"`
	// The traffic identifier of domain. Valid values are:
	TrafficIdentifier map[string]string `pulumi:"trafficIdentifier"`
}

type DedicatedDomainState struct {
	// Whether a domain name is connected to WAF. Valid values are:
	AccessStatus pulumi.IntPtrInput
	// The alarm page of domain. Valid values are:
	AlarmPage pulumi.StringMapInput
	// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
	// is set to HTTPS.
	CertificateId pulumi.StringPtrInput
	// The name of the certificate used by the domain name.
	CertificateName pulumi.StringPtrInput
	Cipher          pulumi.StringPtrInput
	// The compliance certifications of the domain, values are:
	ComplianceCertification pulumi.BoolMapInput
	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	Domain pulumi.StringPtrInput
	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to `true`.
	KeepPolicy pulumi.BoolPtrInput
	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyId pulumi.StringPtrInput
	// The protection status of domain, `0`: suspended, `1`: enabled.
	// Default value is `1`.
	ProtectStatus pulumi.IntPtrInput
	// The protocol type of the client. The options are `HTTP` and `HTTPS`.
	Protocol pulumi.StringPtrInput
	// Specifies whether a proxy is configured. Default value is `false`.
	Proxy pulumi.BoolPtrInput
	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region pulumi.StringPtrInput
	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	Servers DedicatedDomainServerArrayInput
	// The TLS configuration of domain.
	Tls pulumi.StringPtrInput
	// The traffic identifier of domain. Valid values are:
	TrafficIdentifier pulumi.StringMapInput
}

func (DedicatedDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedDomainState)(nil)).Elem()
}

type dedicatedDomainArgs struct {
	// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
	// is set to HTTPS.
	CertificateId *string `pulumi:"certificateId"`
	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	Domain string `pulumi:"domain"`
	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to `true`.
	KeepPolicy *bool `pulumi:"keepPolicy"`
	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyId *string `pulumi:"policyId"`
	// The protection status of domain, `0`: suspended, `1`: enabled.
	// Default value is `1`.
	ProtectStatus *int `pulumi:"protectStatus"`
	// Specifies whether a proxy is configured. Default value is `false`.
	Proxy *bool `pulumi:"proxy"`
	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region *string `pulumi:"region"`
	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	Servers []DedicatedDomainServer `pulumi:"servers"`
}

// The set of arguments for constructing a DedicatedDomain resource.
type DedicatedDomainArgs struct {
	// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
	// is set to HTTPS.
	CertificateId pulumi.StringPtrInput
	// Specifies the domain name to be protected. For example, www.example.com or
	// *.example.com. Changing this creates a new domain.
	Domain pulumi.StringInput
	// Specifies whether to retain the policy when deleting a domain name.
	// Defaults to `true`.
	KeepPolicy pulumi.BoolPtrInput
	// Specifies the policy ID associated with the domain. If not specified, a new policy
	// will be created automatically.
	PolicyId pulumi.StringPtrInput
	// The protection status of domain, `0`: suspended, `1`: enabled.
	// Default value is `1`.
	ProtectStatus pulumi.IntPtrInput
	// Specifies whether a proxy is configured. Default value is `false`.
	Proxy pulumi.BoolPtrInput
	// The region in which to create the dedicated mode domain resource. If omitted,
	// the provider-level region will be used. Changing this setting will push a new domain.
	Region pulumi.StringPtrInput
	// The server configuration list of the domain. A maximum of 80 can be configured.
	// The object structure is documented below.
	Servers DedicatedDomainServerArrayInput
}

func (DedicatedDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedDomainArgs)(nil)).Elem()
}

type DedicatedDomainInput interface {
	pulumi.Input

	ToDedicatedDomainOutput() DedicatedDomainOutput
	ToDedicatedDomainOutputWithContext(ctx context.Context) DedicatedDomainOutput
}

func (*DedicatedDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedDomain)(nil)).Elem()
}

func (i *DedicatedDomain) ToDedicatedDomainOutput() DedicatedDomainOutput {
	return i.ToDedicatedDomainOutputWithContext(context.Background())
}

func (i *DedicatedDomain) ToDedicatedDomainOutputWithContext(ctx context.Context) DedicatedDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedDomainOutput)
}

// DedicatedDomainArrayInput is an input type that accepts DedicatedDomainArray and DedicatedDomainArrayOutput values.
// You can construct a concrete instance of `DedicatedDomainArrayInput` via:
//
//	DedicatedDomainArray{ DedicatedDomainArgs{...} }
type DedicatedDomainArrayInput interface {
	pulumi.Input

	ToDedicatedDomainArrayOutput() DedicatedDomainArrayOutput
	ToDedicatedDomainArrayOutputWithContext(context.Context) DedicatedDomainArrayOutput
}

type DedicatedDomainArray []DedicatedDomainInput

func (DedicatedDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedDomain)(nil)).Elem()
}

func (i DedicatedDomainArray) ToDedicatedDomainArrayOutput() DedicatedDomainArrayOutput {
	return i.ToDedicatedDomainArrayOutputWithContext(context.Background())
}

func (i DedicatedDomainArray) ToDedicatedDomainArrayOutputWithContext(ctx context.Context) DedicatedDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedDomainArrayOutput)
}

// DedicatedDomainMapInput is an input type that accepts DedicatedDomainMap and DedicatedDomainMapOutput values.
// You can construct a concrete instance of `DedicatedDomainMapInput` via:
//
//	DedicatedDomainMap{ "key": DedicatedDomainArgs{...} }
type DedicatedDomainMapInput interface {
	pulumi.Input

	ToDedicatedDomainMapOutput() DedicatedDomainMapOutput
	ToDedicatedDomainMapOutputWithContext(context.Context) DedicatedDomainMapOutput
}

type DedicatedDomainMap map[string]DedicatedDomainInput

func (DedicatedDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedDomain)(nil)).Elem()
}

func (i DedicatedDomainMap) ToDedicatedDomainMapOutput() DedicatedDomainMapOutput {
	return i.ToDedicatedDomainMapOutputWithContext(context.Background())
}

func (i DedicatedDomainMap) ToDedicatedDomainMapOutputWithContext(ctx context.Context) DedicatedDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedDomainMapOutput)
}

type DedicatedDomainOutput struct{ *pulumi.OutputState }

func (DedicatedDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedDomain)(nil)).Elem()
}

func (o DedicatedDomainOutput) ToDedicatedDomainOutput() DedicatedDomainOutput {
	return o
}

func (o DedicatedDomainOutput) ToDedicatedDomainOutputWithContext(ctx context.Context) DedicatedDomainOutput {
	return o
}

// Whether a domain name is connected to WAF. Valid values are:
func (o DedicatedDomainOutput) AccessStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.IntOutput { return v.AccessStatus }).(pulumi.IntOutput)
}

// The alarm page of domain. Valid values are:
func (o DedicatedDomainOutput) AlarmPage() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringMapOutput { return v.AlarmPage }).(pulumi.StringMapOutput)
}

// Specifies the certificate ID. This parameter is mandatory when `clientProtocol`
// is set to HTTPS.
func (o DedicatedDomainOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// The name of the certificate used by the domain name.
func (o DedicatedDomainOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.CertificateName }).(pulumi.StringOutput)
}

func (o DedicatedDomainOutput) Cipher() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.Cipher }).(pulumi.StringOutput)
}

// The compliance certifications of the domain, values are:
func (o DedicatedDomainOutput) ComplianceCertification() pulumi.BoolMapOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.BoolMapOutput { return v.ComplianceCertification }).(pulumi.BoolMapOutput)
}

// Specifies the domain name to be protected. For example, www.example.com or
// *.example.com. Changing this creates a new domain.
func (o DedicatedDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Specifies whether to retain the policy when deleting a domain name.
// Defaults to `true`.
func (o DedicatedDomainOutput) KeepPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.BoolPtrOutput { return v.KeepPolicy }).(pulumi.BoolPtrOutput)
}

// Specifies the policy ID associated with the domain. If not specified, a new policy
// will be created automatically.
func (o DedicatedDomainOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The protection status of domain, `0`: suspended, `1`: enabled.
// Default value is `1`.
func (o DedicatedDomainOutput) ProtectStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.IntOutput { return v.ProtectStatus }).(pulumi.IntOutput)
}

// The protocol type of the client. The options are `HTTP` and `HTTPS`.
func (o DedicatedDomainOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Specifies whether a proxy is configured. Default value is `false`.
func (o DedicatedDomainOutput) Proxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.BoolPtrOutput { return v.Proxy }).(pulumi.BoolPtrOutput)
}

// The region in which to create the dedicated mode domain resource. If omitted,
// the provider-level region will be used. Changing this setting will push a new domain.
func (o DedicatedDomainOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The server configuration list of the domain. A maximum of 80 can be configured.
// The object structure is documented below.
func (o DedicatedDomainOutput) Servers() DedicatedDomainServerArrayOutput {
	return o.ApplyT(func(v *DedicatedDomain) DedicatedDomainServerArrayOutput { return v.Servers }).(DedicatedDomainServerArrayOutput)
}

// The TLS configuration of domain.
func (o DedicatedDomainOutput) Tls() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringOutput { return v.Tls }).(pulumi.StringOutput)
}

// The traffic identifier of domain. Valid values are:
func (o DedicatedDomainOutput) TrafficIdentifier() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedDomain) pulumi.StringMapOutput { return v.TrafficIdentifier }).(pulumi.StringMapOutput)
}

type DedicatedDomainArrayOutput struct{ *pulumi.OutputState }

func (DedicatedDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedDomain)(nil)).Elem()
}

func (o DedicatedDomainArrayOutput) ToDedicatedDomainArrayOutput() DedicatedDomainArrayOutput {
	return o
}

func (o DedicatedDomainArrayOutput) ToDedicatedDomainArrayOutputWithContext(ctx context.Context) DedicatedDomainArrayOutput {
	return o
}

func (o DedicatedDomainArrayOutput) Index(i pulumi.IntInput) DedicatedDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedDomain {
		return vs[0].([]*DedicatedDomain)[vs[1].(int)]
	}).(DedicatedDomainOutput)
}

type DedicatedDomainMapOutput struct{ *pulumi.OutputState }

func (DedicatedDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedDomain)(nil)).Elem()
}

func (o DedicatedDomainMapOutput) ToDedicatedDomainMapOutput() DedicatedDomainMapOutput {
	return o
}

func (o DedicatedDomainMapOutput) ToDedicatedDomainMapOutputWithContext(ctx context.Context) DedicatedDomainMapOutput {
	return o
}

func (o DedicatedDomainMapOutput) MapIndex(k pulumi.StringInput) DedicatedDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedDomain {
		return vs[0].(map[string]*DedicatedDomain)[vs[1].(string)]
	}).(DedicatedDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedDomainInput)(nil)).Elem(), &DedicatedDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedDomainArrayInput)(nil)).Elem(), DedicatedDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedDomainMapInput)(nil)).Elem(), DedicatedDomainMap{})
	pulumi.RegisterOutputType(DedicatedDomainOutput{})
	pulumi.RegisterOutputType(DedicatedDomainArrayOutput{})
	pulumi.RegisterOutputType(DedicatedDomainMapOutput{})
}
