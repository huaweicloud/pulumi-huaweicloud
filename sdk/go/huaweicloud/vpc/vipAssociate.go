// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this resource, one or more NICs (to which the ECS instance belongs) can be bound to the VIP.
//
// > A VIP can only have one resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vipId := cfg.RequireObject("vipId")
//			nicPortIds := cfg.Require("nicPortIds")
//			_, err := Vpc.NewVipAssociate(ctx, "vipAssociated", &Vpc.VipAssociateArgs{
//				VipId:   pulumi.Any(vipId),
//				PortIds: nicPortIds,
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
// Vip associate can be imported using the `vip_id` and port IDs separated by slashes (no limit on the number of port IDs), e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpc/vipAssociate:VipAssociate vip_associated vip_id/port1_id/port2_id
//
// ```
type VipAssociate struct {
	pulumi.CustomResourceState

	// The IP addresses of ports to attach the vip to.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// An array of one or more IDs of the ports to attach the vip to.
	PortIds pulumi.StringArrayOutput `pulumi:"portIds"`
	// The region in which to create the vip associate resource. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of vip to attach the ports to.
	VipId pulumi.StringOutput `pulumi:"vipId"`
	// The IP address in the subnet for this vip.
	VipIpAddress pulumi.StringOutput `pulumi:"vipIpAddress"`
	// The ID of the subnet this vip connects to.
	VipSubnetId pulumi.StringOutput `pulumi:"vipSubnetId"`
}

// NewVipAssociate registers a new resource with the given unique name, arguments, and options.
func NewVipAssociate(ctx *pulumi.Context,
	name string, args *VipAssociateArgs, opts ...pulumi.ResourceOption) (*VipAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortIds == nil {
		return nil, errors.New("invalid value for required argument 'PortIds'")
	}
	if args.VipId == nil {
		return nil, errors.New("invalid value for required argument 'VipId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VipAssociate
	err := ctx.RegisterResource("huaweicloud:Vpc/vipAssociate:VipAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVipAssociate gets an existing VipAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVipAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VipAssociateState, opts ...pulumi.ResourceOption) (*VipAssociate, error) {
	var resource VipAssociate
	err := ctx.ReadResource("huaweicloud:Vpc/vipAssociate:VipAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VipAssociate resources.
type vipAssociateState struct {
	// The IP addresses of ports to attach the vip to.
	IpAddresses []string `pulumi:"ipAddresses"`
	// An array of one or more IDs of the ports to attach the vip to.
	PortIds []string `pulumi:"portIds"`
	// The region in which to create the vip associate resource. If omitted, the
	// provider-level region will be used.
	Region *string `pulumi:"region"`
	// The ID of vip to attach the ports to.
	VipId *string `pulumi:"vipId"`
	// The IP address in the subnet for this vip.
	VipIpAddress *string `pulumi:"vipIpAddress"`
	// The ID of the subnet this vip connects to.
	VipSubnetId *string `pulumi:"vipSubnetId"`
}

type VipAssociateState struct {
	// The IP addresses of ports to attach the vip to.
	IpAddresses pulumi.StringArrayInput
	// An array of one or more IDs of the ports to attach the vip to.
	PortIds pulumi.StringArrayInput
	// The region in which to create the vip associate resource. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringPtrInput
	// The ID of vip to attach the ports to.
	VipId pulumi.StringPtrInput
	// The IP address in the subnet for this vip.
	VipIpAddress pulumi.StringPtrInput
	// The ID of the subnet this vip connects to.
	VipSubnetId pulumi.StringPtrInput
}

func (VipAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*vipAssociateState)(nil)).Elem()
}

type vipAssociateArgs struct {
	// An array of one or more IDs of the ports to attach the vip to.
	PortIds []string `pulumi:"portIds"`
	// The region in which to create the vip associate resource. If omitted, the
	// provider-level region will be used.
	Region *string `pulumi:"region"`
	// The ID of vip to attach the ports to.
	VipId string `pulumi:"vipId"`
}

// The set of arguments for constructing a VipAssociate resource.
type VipAssociateArgs struct {
	// An array of one or more IDs of the ports to attach the vip to.
	PortIds pulumi.StringArrayInput
	// The region in which to create the vip associate resource. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringPtrInput
	// The ID of vip to attach the ports to.
	VipId pulumi.StringInput
}

func (VipAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vipAssociateArgs)(nil)).Elem()
}

type VipAssociateInput interface {
	pulumi.Input

	ToVipAssociateOutput() VipAssociateOutput
	ToVipAssociateOutputWithContext(ctx context.Context) VipAssociateOutput
}

func (*VipAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**VipAssociate)(nil)).Elem()
}

func (i *VipAssociate) ToVipAssociateOutput() VipAssociateOutput {
	return i.ToVipAssociateOutputWithContext(context.Background())
}

func (i *VipAssociate) ToVipAssociateOutputWithContext(ctx context.Context) VipAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipAssociateOutput)
}

// VipAssociateArrayInput is an input type that accepts VipAssociateArray and VipAssociateArrayOutput values.
// You can construct a concrete instance of `VipAssociateArrayInput` via:
//
//	VipAssociateArray{ VipAssociateArgs{...} }
type VipAssociateArrayInput interface {
	pulumi.Input

	ToVipAssociateArrayOutput() VipAssociateArrayOutput
	ToVipAssociateArrayOutputWithContext(context.Context) VipAssociateArrayOutput
}

type VipAssociateArray []VipAssociateInput

func (VipAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VipAssociate)(nil)).Elem()
}

func (i VipAssociateArray) ToVipAssociateArrayOutput() VipAssociateArrayOutput {
	return i.ToVipAssociateArrayOutputWithContext(context.Background())
}

func (i VipAssociateArray) ToVipAssociateArrayOutputWithContext(ctx context.Context) VipAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipAssociateArrayOutput)
}

// VipAssociateMapInput is an input type that accepts VipAssociateMap and VipAssociateMapOutput values.
// You can construct a concrete instance of `VipAssociateMapInput` via:
//
//	VipAssociateMap{ "key": VipAssociateArgs{...} }
type VipAssociateMapInput interface {
	pulumi.Input

	ToVipAssociateMapOutput() VipAssociateMapOutput
	ToVipAssociateMapOutputWithContext(context.Context) VipAssociateMapOutput
}

type VipAssociateMap map[string]VipAssociateInput

func (VipAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VipAssociate)(nil)).Elem()
}

func (i VipAssociateMap) ToVipAssociateMapOutput() VipAssociateMapOutput {
	return i.ToVipAssociateMapOutputWithContext(context.Background())
}

func (i VipAssociateMap) ToVipAssociateMapOutputWithContext(ctx context.Context) VipAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipAssociateMapOutput)
}

type VipAssociateOutput struct{ *pulumi.OutputState }

func (VipAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VipAssociate)(nil)).Elem()
}

func (o VipAssociateOutput) ToVipAssociateOutput() VipAssociateOutput {
	return o
}

func (o VipAssociateOutput) ToVipAssociateOutputWithContext(ctx context.Context) VipAssociateOutput {
	return o
}

// The IP addresses of ports to attach the vip to.
func (o VipAssociateOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// An array of one or more IDs of the ports to attach the vip to.
func (o VipAssociateOutput) PortIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringArrayOutput { return v.PortIds }).(pulumi.StringArrayOutput)
}

// The region in which to create the vip associate resource. If omitted, the
// provider-level region will be used.
func (o VipAssociateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of vip to attach the ports to.
func (o VipAssociateOutput) VipId() pulumi.StringOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringOutput { return v.VipId }).(pulumi.StringOutput)
}

// The IP address in the subnet for this vip.
func (o VipAssociateOutput) VipIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringOutput { return v.VipIpAddress }).(pulumi.StringOutput)
}

// The ID of the subnet this vip connects to.
func (o VipAssociateOutput) VipSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VipAssociate) pulumi.StringOutput { return v.VipSubnetId }).(pulumi.StringOutput)
}

type VipAssociateArrayOutput struct{ *pulumi.OutputState }

func (VipAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VipAssociate)(nil)).Elem()
}

func (o VipAssociateArrayOutput) ToVipAssociateArrayOutput() VipAssociateArrayOutput {
	return o
}

func (o VipAssociateArrayOutput) ToVipAssociateArrayOutputWithContext(ctx context.Context) VipAssociateArrayOutput {
	return o
}

func (o VipAssociateArrayOutput) Index(i pulumi.IntInput) VipAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VipAssociate {
		return vs[0].([]*VipAssociate)[vs[1].(int)]
	}).(VipAssociateOutput)
}

type VipAssociateMapOutput struct{ *pulumi.OutputState }

func (VipAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VipAssociate)(nil)).Elem()
}

func (o VipAssociateMapOutput) ToVipAssociateMapOutput() VipAssociateMapOutput {
	return o
}

func (o VipAssociateMapOutput) ToVipAssociateMapOutputWithContext(ctx context.Context) VipAssociateMapOutput {
	return o
}

func (o VipAssociateMapOutput) MapIndex(k pulumi.StringInput) VipAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VipAssociate {
		return vs[0].(map[string]*VipAssociate)[vs[1].(string)]
	}).(VipAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VipAssociateInput)(nil)).Elem(), &VipAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipAssociateArrayInput)(nil)).Elem(), VipAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipAssociateMapInput)(nil)).Elem(), VipAssociateMap{})
	pulumi.RegisterOutputType(VipAssociateOutput{})
	pulumi.RegisterOutputType(VipAssociateArrayOutput{})
	pulumi.RegisterOutputType(VipAssociateMapOutput{})
}
