// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EIP resource within HuaweiCloud.
//
// ## Example Usage
// ### EIP with Dedicated Bandwidth
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.NewEip(ctx, "eip1", &Vpc.EipArgs{
//				Bandwidth: &vpc.EipBandwidthArgs{
//					ChargeMode: pulumi.String("traffic"),
//					Name:       pulumi.String("test"),
//					ShareType:  pulumi.String("PER"),
//					Size:       pulumi.Int(10),
//				},
//				Publicip: &vpc.EipPublicipArgs{
//					Type: pulumi.String("5_bgp"),
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
// ### EIP with Shared Bandwidth
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bandwidth1, err := Vpc.NewBandwidth(ctx, "bandwidth1", &Vpc.BandwidthArgs{
//				Size: pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.NewEip(ctx, "eip1", &Vpc.EipArgs{
//				Publicip: &vpc.EipPublicipArgs{
//					Type: pulumi.String("5_bgp"),
//				},
//				Bandwidth: &vpc.EipBandwidthArgs{
//					ShareType: pulumi.String("WHOLE"),
//					Id:        bandwidth1.ID(),
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
// EIPs can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpc/eip:Eip eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
//
// ```
type Eip struct {
	pulumi.CustomResourceState

	// The IPv4 address of the EIP.
	Address pulumi.StringOutput    `pulumi:"address"`
	AutoPay pulumi.StringPtrOutput `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled.
	// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
	AutoRenew pulumi.StringPtrOutput `pulumi:"autoRenew"`
	// The bandwidth object.
	Bandwidth EipBandwidthOutput `pulumi:"bandwidth"`
	// Specifies the charging mode of the elastic IP. Valid values are
	// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
	ChargingMode pulumi.StringOutput `pulumi:"chargingMode"`
	// The enterprise project id of the elastic IP. Changing this
	// creates a new eip.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// The IPv6 address of the EIP.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
	// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the charging period of the elastic IP. If `periodUnit` is set to
	// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
	// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Specifies the charging period unit of the elastic IP. Valid values are
	// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
	// eip.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// The port ID which the EIP associated with.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The private IP address bound to the EIP.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The elastic IP address object.
	Publicip EipPublicipOutput `pulumi:"publicip"`
	// The region in which to create the EIP resource. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of EIP.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the key/value pairs to associate with the elastic IP.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEip registers a new resource with the given unique name, arguments, and options.
func NewEip(ctx *pulumi.Context,
	name string, args *EipArgs, opts ...pulumi.ResourceOption) (*Eip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Publicip == nil {
		return nil, errors.New("invalid value for required argument 'Publicip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Eip
	err := ctx.RegisterResource("huaweicloud:Vpc/eip:Eip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEip gets an existing Eip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipState, opts ...pulumi.ResourceOption) (*Eip, error) {
	var resource Eip
	err := ctx.ReadResource("huaweicloud:Vpc/eip:Eip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eip resources.
type eipState struct {
	// The IPv4 address of the EIP.
	Address *string `pulumi:"address"`
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled.
	// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
	AutoRenew *string `pulumi:"autoRenew"`
	// The bandwidth object.
	Bandwidth *EipBandwidth `pulumi:"bandwidth"`
	// Specifies the charging mode of the elastic IP. Valid values are
	// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
	ChargingMode *string `pulumi:"chargingMode"`
	// The enterprise project id of the elastic IP. Changing this
	// creates a new eip.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The IPv6 address of the EIP.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
	// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
	Name *string `pulumi:"name"`
	// Specifies the charging period of the elastic IP. If `periodUnit` is set to
	// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
	// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the elastic IP. Valid values are
	// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
	// eip.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The port ID which the EIP associated with.
	PortId *string `pulumi:"portId"`
	// The private IP address bound to the EIP.
	PrivateIp *string `pulumi:"privateIp"`
	// The elastic IP address object.
	Publicip *EipPublicip `pulumi:"publicip"`
	// The region in which to create the EIP resource. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// The status of EIP.
	Status *string `pulumi:"status"`
	// Specifies the key/value pairs to associate with the elastic IP.
	Tags map[string]string `pulumi:"tags"`
}

type EipState struct {
	// The IPv4 address of the EIP.
	Address pulumi.StringPtrInput
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto renew is enabled.
	// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
	AutoRenew pulumi.StringPtrInput
	// The bandwidth object.
	Bandwidth EipBandwidthPtrInput
	// Specifies the charging mode of the elastic IP. Valid values are
	// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
	ChargingMode pulumi.StringPtrInput
	// The enterprise project id of the elastic IP. Changing this
	// creates a new eip.
	EnterpriseProjectId pulumi.StringPtrInput
	// The IPv6 address of the EIP.
	Ipv6Address pulumi.StringPtrInput
	// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
	// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
	Name pulumi.StringPtrInput
	// Specifies the charging period of the elastic IP. If `periodUnit` is set to
	// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
	// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the elastic IP. Valid values are
	// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
	// eip.
	PeriodUnit pulumi.StringPtrInput
	// The port ID which the EIP associated with.
	PortId pulumi.StringPtrInput
	// The private IP address bound to the EIP.
	PrivateIp pulumi.StringPtrInput
	// The elastic IP address object.
	Publicip EipPublicipPtrInput
	// The region in which to create the EIP resource. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// The status of EIP.
	Status pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the elastic IP.
	Tags pulumi.StringMapInput
}

func (EipState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto renew is enabled.
	// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
	AutoRenew *string `pulumi:"autoRenew"`
	// The bandwidth object.
	Bandwidth EipBandwidth `pulumi:"bandwidth"`
	// Specifies the charging mode of the elastic IP. Valid values are
	// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
	ChargingMode *string `pulumi:"chargingMode"`
	// The enterprise project id of the elastic IP. Changing this
	// creates a new eip.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
	// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
	Name *string `pulumi:"name"`
	// Specifies the charging period of the elastic IP. If `periodUnit` is set to
	// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
	// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the elastic IP. Valid values are
	// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
	// eip.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The elastic IP address object.
	Publicip EipPublicip `pulumi:"publicip"`
	// The region in which to create the EIP resource. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs to associate with the elastic IP.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Eip resource.
type EipArgs struct {
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto renew is enabled.
	// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
	AutoRenew pulumi.StringPtrInput
	// The bandwidth object.
	Bandwidth EipBandwidthInput
	// Specifies the charging mode of the elastic IP. Valid values are
	// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
	ChargingMode pulumi.StringPtrInput
	// The enterprise project id of the elastic IP. Changing this
	// creates a new eip.
	EnterpriseProjectId pulumi.StringPtrInput
	// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
	// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
	Name pulumi.StringPtrInput
	// Specifies the charging period of the elastic IP. If `periodUnit` is set to
	// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
	// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the elastic IP. Valid values are
	// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
	// eip.
	PeriodUnit pulumi.StringPtrInput
	// The elastic IP address object.
	Publicip EipPublicipInput
	// The region in which to create the EIP resource. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the elastic IP.
	Tags pulumi.StringMapInput
}

func (EipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EipInput interface {
	pulumi.Input

	ToEipOutput() EipOutput
	ToEipOutputWithContext(ctx context.Context) EipOutput
}

func (*Eip) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (i *Eip) ToEipOutput() EipOutput {
	return i.ToEipOutputWithContext(context.Background())
}

func (i *Eip) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipOutput)
}

// EipArrayInput is an input type that accepts EipArray and EipArrayOutput values.
// You can construct a concrete instance of `EipArrayInput` via:
//
//	EipArray{ EipArgs{...} }
type EipArrayInput interface {
	pulumi.Input

	ToEipArrayOutput() EipArrayOutput
	ToEipArrayOutputWithContext(context.Context) EipArrayOutput
}

type EipArray []EipInput

func (EipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (i EipArray) ToEipArrayOutput() EipArrayOutput {
	return i.ToEipArrayOutputWithContext(context.Background())
}

func (i EipArray) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipArrayOutput)
}

// EipMapInput is an input type that accepts EipMap and EipMapOutput values.
// You can construct a concrete instance of `EipMapInput` via:
//
//	EipMap{ "key": EipArgs{...} }
type EipMapInput interface {
	pulumi.Input

	ToEipMapOutput() EipMapOutput
	ToEipMapOutputWithContext(context.Context) EipMapOutput
}

type EipMap map[string]EipInput

func (EipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (i EipMap) ToEipMapOutput() EipMapOutput {
	return i.ToEipMapOutputWithContext(context.Background())
}

func (i EipMap) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipMapOutput)
}

type EipOutput struct{ *pulumi.OutputState }

func (EipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (o EipOutput) ToEipOutput() EipOutput {
	return o
}

func (o EipOutput) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return o
}

// The IPv4 address of the EIP.
func (o EipOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

func (o EipOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

// Specifies whether auto renew is enabled.
// Valid values are *true* and *false*. Defaults to *false*. Changing this creates a new resource.
func (o EipOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

// The bandwidth object.
func (o EipOutput) Bandwidth() EipBandwidthOutput {
	return o.ApplyT(func(v *Eip) EipBandwidthOutput { return v.Bandwidth }).(EipBandwidthOutput)
}

// Specifies the charging mode of the elastic IP. Valid values are
// *prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.
func (o EipOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

// The enterprise project id of the elastic IP. Changing this
// creates a new eip.
func (o EipOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The IPv6 address of the EIP.
func (o EipOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
// underscores (_), and hyphens (-). This parameter is mandatory when `shareType` is set to **PER**.
func (o EipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the charging period of the elastic IP. If `periodUnit` is set to
// *month*, the value ranges from 1 to 9. If `periodUnit` is set to *year*, the value ranges from 1 to 3. This parameter
// is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new resource.
func (o EipOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Specifies the charging period unit of the elastic IP. Valid values are
// *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*. Changing this creates a new
// eip.
func (o EipOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// The port ID which the EIP associated with.
func (o EipOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The private IP address bound to the EIP.
func (o EipOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// The elastic IP address object.
func (o EipOutput) Publicip() EipPublicipOutput {
	return o.ApplyT(func(v *Eip) EipPublicipOutput { return v.Publicip }).(EipPublicipOutput)
}

// The region in which to create the EIP resource. If omitted, the provider-level
// region will be used. Changing this creates a new resource.
func (o EipOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The status of EIP.
func (o EipOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the key/value pairs to associate with the elastic IP.
func (o EipOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type EipArrayOutput struct{ *pulumi.OutputState }

func (EipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (o EipArrayOutput) ToEipArrayOutput() EipArrayOutput {
	return o
}

func (o EipArrayOutput) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return o
}

func (o EipArrayOutput) Index(i pulumi.IntInput) EipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].([]*Eip)[vs[1].(int)]
	}).(EipOutput)
}

type EipMapOutput struct{ *pulumi.OutputState }

func (EipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (o EipMapOutput) ToEipMapOutput() EipMapOutput {
	return o
}

func (o EipMapOutput) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return o
}

func (o EipMapOutput) MapIndex(k pulumi.StringInput) EipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].(map[string]*Eip)[vs[1].(string)]
	}).(EipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipInput)(nil)).Elem(), &Eip{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipArrayInput)(nil)).Elem(), EipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipMapInput)(nil)).Elem(), EipMap{})
	pulumi.RegisterOutputType(EipOutput{})
	pulumi.RegisterOutputType(EipArrayOutput{})
	pulumi.RegisterOutputType(EipMapOutput{})
}
