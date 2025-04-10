// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := cfg.RequireObject("name")
//			userId := cfg.RequireObject("userId")
//			_, err := Iam.NewVirtualMfaDevice(ctx, "test", &Iam.VirtualMfaDeviceArgs{
//				UserId: pulumi.Any(userId),
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
// The virtual MFA device can be imported using the `user_id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Iam/virtualMfaDevice:VirtualMfaDevice test <user_id>
//
// ```
type VirtualMfaDevice struct {
	pulumi.CustomResourceState

	// The base32 seed, which a third-patry system can use to generate a `CAPTCHA` code.
	Base32StringSeed pulumi.StringOutput `pulumi:"base32StringSeed"`
	// Specifies the virtual MFA device name. Changing this will create a new virtual
	// MFA device.
	Name pulumi.StringOutput `pulumi:"name"`
	// A QR code PNG image that encodes `otpauth://totp/huawei:$domainName@$userName?secret=$Base32String`
	// where `$domainName` is IAM domain name, `$userName` is IAM user name, and `Base32String` is the seed in base32 format.
	QrCodePng pulumi.StringOutput `pulumi:"qrCodePng"`
	// Specifies the user ID which the virtual MFA device belongs to.
	// Changing this will create a new virtual MFA device.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewVirtualMfaDevice registers a new resource with the given unique name, arguments, and options.
func NewVirtualMfaDevice(ctx *pulumi.Context,
	name string, args *VirtualMfaDeviceArgs, opts ...pulumi.ResourceOption) (*VirtualMfaDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VirtualMfaDevice
	err := ctx.RegisterResource("huaweicloud:Iam/virtualMfaDevice:VirtualMfaDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMfaDevice gets an existing VirtualMfaDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMfaDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMfaDeviceState, opts ...pulumi.ResourceOption) (*VirtualMfaDevice, error) {
	var resource VirtualMfaDevice
	err := ctx.ReadResource("huaweicloud:Iam/virtualMfaDevice:VirtualMfaDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMfaDevice resources.
type virtualMfaDeviceState struct {
	// The base32 seed, which a third-patry system can use to generate a `CAPTCHA` code.
	Base32StringSeed *string `pulumi:"base32StringSeed"`
	// Specifies the virtual MFA device name. Changing this will create a new virtual
	// MFA device.
	Name *string `pulumi:"name"`
	// A QR code PNG image that encodes `otpauth://totp/huawei:$domainName@$userName?secret=$Base32String`
	// where `$domainName` is IAM domain name, `$userName` is IAM user name, and `Base32String` is the seed in base32 format.
	QrCodePng *string `pulumi:"qrCodePng"`
	// Specifies the user ID which the virtual MFA device belongs to.
	// Changing this will create a new virtual MFA device.
	UserId *string `pulumi:"userId"`
}

type VirtualMfaDeviceState struct {
	// The base32 seed, which a third-patry system can use to generate a `CAPTCHA` code.
	Base32StringSeed pulumi.StringPtrInput
	// Specifies the virtual MFA device name. Changing this will create a new virtual
	// MFA device.
	Name pulumi.StringPtrInput
	// A QR code PNG image that encodes `otpauth://totp/huawei:$domainName@$userName?secret=$Base32String`
	// where `$domainName` is IAM domain name, `$userName` is IAM user name, and `Base32String` is the seed in base32 format.
	QrCodePng pulumi.StringPtrInput
	// Specifies the user ID which the virtual MFA device belongs to.
	// Changing this will create a new virtual MFA device.
	UserId pulumi.StringPtrInput
}

func (VirtualMfaDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMfaDeviceState)(nil)).Elem()
}

type virtualMfaDeviceArgs struct {
	// Specifies the virtual MFA device name. Changing this will create a new virtual
	// MFA device.
	Name *string `pulumi:"name"`
	// Specifies the user ID which the virtual MFA device belongs to.
	// Changing this will create a new virtual MFA device.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a VirtualMfaDevice resource.
type VirtualMfaDeviceArgs struct {
	// Specifies the virtual MFA device name. Changing this will create a new virtual
	// MFA device.
	Name pulumi.StringPtrInput
	// Specifies the user ID which the virtual MFA device belongs to.
	// Changing this will create a new virtual MFA device.
	UserId pulumi.StringInput
}

func (VirtualMfaDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMfaDeviceArgs)(nil)).Elem()
}

type VirtualMfaDeviceInput interface {
	pulumi.Input

	ToVirtualMfaDeviceOutput() VirtualMfaDeviceOutput
	ToVirtualMfaDeviceOutputWithContext(ctx context.Context) VirtualMfaDeviceOutput
}

func (*VirtualMfaDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMfaDevice)(nil)).Elem()
}

func (i *VirtualMfaDevice) ToVirtualMfaDeviceOutput() VirtualMfaDeviceOutput {
	return i.ToVirtualMfaDeviceOutputWithContext(context.Background())
}

func (i *VirtualMfaDevice) ToVirtualMfaDeviceOutputWithContext(ctx context.Context) VirtualMfaDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMfaDeviceOutput)
}

// VirtualMfaDeviceArrayInput is an input type that accepts VirtualMfaDeviceArray and VirtualMfaDeviceArrayOutput values.
// You can construct a concrete instance of `VirtualMfaDeviceArrayInput` via:
//
//	VirtualMfaDeviceArray{ VirtualMfaDeviceArgs{...} }
type VirtualMfaDeviceArrayInput interface {
	pulumi.Input

	ToVirtualMfaDeviceArrayOutput() VirtualMfaDeviceArrayOutput
	ToVirtualMfaDeviceArrayOutputWithContext(context.Context) VirtualMfaDeviceArrayOutput
}

type VirtualMfaDeviceArray []VirtualMfaDeviceInput

func (VirtualMfaDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMfaDevice)(nil)).Elem()
}

func (i VirtualMfaDeviceArray) ToVirtualMfaDeviceArrayOutput() VirtualMfaDeviceArrayOutput {
	return i.ToVirtualMfaDeviceArrayOutputWithContext(context.Background())
}

func (i VirtualMfaDeviceArray) ToVirtualMfaDeviceArrayOutputWithContext(ctx context.Context) VirtualMfaDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMfaDeviceArrayOutput)
}

// VirtualMfaDeviceMapInput is an input type that accepts VirtualMfaDeviceMap and VirtualMfaDeviceMapOutput values.
// You can construct a concrete instance of `VirtualMfaDeviceMapInput` via:
//
//	VirtualMfaDeviceMap{ "key": VirtualMfaDeviceArgs{...} }
type VirtualMfaDeviceMapInput interface {
	pulumi.Input

	ToVirtualMfaDeviceMapOutput() VirtualMfaDeviceMapOutput
	ToVirtualMfaDeviceMapOutputWithContext(context.Context) VirtualMfaDeviceMapOutput
}

type VirtualMfaDeviceMap map[string]VirtualMfaDeviceInput

func (VirtualMfaDeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMfaDevice)(nil)).Elem()
}

func (i VirtualMfaDeviceMap) ToVirtualMfaDeviceMapOutput() VirtualMfaDeviceMapOutput {
	return i.ToVirtualMfaDeviceMapOutputWithContext(context.Background())
}

func (i VirtualMfaDeviceMap) ToVirtualMfaDeviceMapOutputWithContext(ctx context.Context) VirtualMfaDeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMfaDeviceMapOutput)
}

type VirtualMfaDeviceOutput struct{ *pulumi.OutputState }

func (VirtualMfaDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMfaDevice)(nil)).Elem()
}

func (o VirtualMfaDeviceOutput) ToVirtualMfaDeviceOutput() VirtualMfaDeviceOutput {
	return o
}

func (o VirtualMfaDeviceOutput) ToVirtualMfaDeviceOutputWithContext(ctx context.Context) VirtualMfaDeviceOutput {
	return o
}

// The base32 seed, which a third-patry system can use to generate a `CAPTCHA` code.
func (o VirtualMfaDeviceOutput) Base32StringSeed() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMfaDevice) pulumi.StringOutput { return v.Base32StringSeed }).(pulumi.StringOutput)
}

// Specifies the virtual MFA device name. Changing this will create a new virtual
// MFA device.
func (o VirtualMfaDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMfaDevice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A QR code PNG image that encodes `otpauth://totp/huawei:$domainName@$userName?secret=$Base32String`
// where `$domainName` is IAM domain name, `$userName` is IAM user name, and `Base32String` is the seed in base32 format.
func (o VirtualMfaDeviceOutput) QrCodePng() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMfaDevice) pulumi.StringOutput { return v.QrCodePng }).(pulumi.StringOutput)
}

// Specifies the user ID which the virtual MFA device belongs to.
// Changing this will create a new virtual MFA device.
func (o VirtualMfaDeviceOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMfaDevice) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type VirtualMfaDeviceArrayOutput struct{ *pulumi.OutputState }

func (VirtualMfaDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMfaDevice)(nil)).Elem()
}

func (o VirtualMfaDeviceArrayOutput) ToVirtualMfaDeviceArrayOutput() VirtualMfaDeviceArrayOutput {
	return o
}

func (o VirtualMfaDeviceArrayOutput) ToVirtualMfaDeviceArrayOutputWithContext(ctx context.Context) VirtualMfaDeviceArrayOutput {
	return o
}

func (o VirtualMfaDeviceArrayOutput) Index(i pulumi.IntInput) VirtualMfaDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualMfaDevice {
		return vs[0].([]*VirtualMfaDevice)[vs[1].(int)]
	}).(VirtualMfaDeviceOutput)
}

type VirtualMfaDeviceMapOutput struct{ *pulumi.OutputState }

func (VirtualMfaDeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMfaDevice)(nil)).Elem()
}

func (o VirtualMfaDeviceMapOutput) ToVirtualMfaDeviceMapOutput() VirtualMfaDeviceMapOutput {
	return o
}

func (o VirtualMfaDeviceMapOutput) ToVirtualMfaDeviceMapOutputWithContext(ctx context.Context) VirtualMfaDeviceMapOutput {
	return o
}

func (o VirtualMfaDeviceMapOutput) MapIndex(k pulumi.StringInput) VirtualMfaDeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualMfaDevice {
		return vs[0].(map[string]*VirtualMfaDevice)[vs[1].(string)]
	}).(VirtualMfaDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMfaDeviceInput)(nil)).Elem(), &VirtualMfaDevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMfaDeviceArrayInput)(nil)).Elem(), VirtualMfaDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMfaDeviceMapInput)(nil)).Elem(), VirtualMfaDeviceMap{})
	pulumi.RegisterOutputType(VirtualMfaDeviceOutput{})
	pulumi.RegisterOutputType(VirtualMfaDeviceArrayOutput{})
	pulumi.RegisterOutputType(VirtualMfaDeviceMapOutput{})
}
