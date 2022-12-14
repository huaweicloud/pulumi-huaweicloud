// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNAT rule resource within HuaweiCloud.
//
// ## Example Usage
// ### DNAT rule in Direct Connect scenario
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Nat"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Nat.NewDnatRule(ctx, "dnat2", &Nat.DnatRuleArgs{
//				NatGatewayId:        pulumi.Any(_var.Natgw_id),
//				FloatingIpId:        pulumi.Any(_var.Publicip_id),
//				PrivateIp:           pulumi.String("10.0.0.12"),
//				Protocol:            pulumi.String("tcp"),
//				InternalServicePort: pulumi.Int(80),
//				ExternalServicePort: pulumi.Int(8080),
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
// # DNAT rules can be imported using the following format
//
// ```sh
//
//	$ pulumi import huaweicloud:Nat/dnatRule:DnatRule dnat_1 f4f783a7-b908-4215-b018-724960e5df4a
//
// ```
type DnatRule struct {
	pulumi.CustomResourceState

	// Dnat rule creation time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the description of the dnat rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// Changing this creates a new dnat rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	ExternalServicePort pulumi.IntOutput `pulumi:"externalServicePort"`
	// The actual floating IP address.
	FloatingIpAddress pulumi.StringOutput `pulumi:"floatingIpAddress"`
	// Specifies the ID of the floating IP address. Changing this creates a
	// new dnat rule.
	FloatingIpId pulumi.StringOutput `pulumi:"floatingIpId"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	InternalServicePort pulumi.IntOutput `pulumi:"internalServicePort"`
	// ID of the nat gateway this dnat rule belongs to. Changing this creates
	// a new dnat rule.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// Specifies the port ID of network. This parameter is mandatory in VPC
	// scenario. Use Vpc.Port to get the port if just know a fixed IP
	// addresses on the port. Changing this creates a new dnat rule.
	PortId pulumi.StringPtrOutput `pulumi:"portId"`
	// Specifies the private IP address of a user. This parameter is mandatory in
	// Direct Connect scenario. Changing this creates a new dnat rule.
	PrivateIp pulumi.StringPtrOutput `pulumi:"privateIp"`
	// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
	// Changing this creates a new dnat rule.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to create the dnat rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new dnat rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// Dnat rule status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDnatRule registers a new resource with the given unique name, arguments, and options.
func NewDnatRule(ctx *pulumi.Context,
	name string, args *DnatRuleArgs, opts ...pulumi.ResourceOption) (*DnatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalServicePort == nil {
		return nil, errors.New("invalid value for required argument 'ExternalServicePort'")
	}
	if args.FloatingIpId == nil {
		return nil, errors.New("invalid value for required argument 'FloatingIpId'")
	}
	if args.InternalServicePort == nil {
		return nil, errors.New("invalid value for required argument 'InternalServicePort'")
	}
	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DnatRule
	err := ctx.RegisterResource("huaweicloud:Nat/dnatRule:DnatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnatRule gets an existing DnatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnatRuleState, opts ...pulumi.ResourceOption) (*DnatRule, error) {
	var resource DnatRule
	err := ctx.ReadResource("huaweicloud:Nat/dnatRule:DnatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnatRule resources.
type dnatRuleState struct {
	// Dnat rule creation time.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the description of the dnat rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// Changing this creates a new dnat rule.
	Description *string `pulumi:"description"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	ExternalServicePort *int `pulumi:"externalServicePort"`
	// The actual floating IP address.
	FloatingIpAddress *string `pulumi:"floatingIpAddress"`
	// Specifies the ID of the floating IP address. Changing this creates a
	// new dnat rule.
	FloatingIpId *string `pulumi:"floatingIpId"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	InternalServicePort *int `pulumi:"internalServicePort"`
	// ID of the nat gateway this dnat rule belongs to. Changing this creates
	// a new dnat rule.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Specifies the port ID of network. This parameter is mandatory in VPC
	// scenario. Use Vpc.Port to get the port if just know a fixed IP
	// addresses on the port. Changing this creates a new dnat rule.
	PortId *string `pulumi:"portId"`
	// Specifies the private IP address of a user. This parameter is mandatory in
	// Direct Connect scenario. Changing this creates a new dnat rule.
	PrivateIp *string `pulumi:"privateIp"`
	// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
	// Changing this creates a new dnat rule.
	Protocol *string `pulumi:"protocol"`
	// The region in which to create the dnat rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new dnat rule.
	Region *string `pulumi:"region"`
	// Dnat rule status.
	Status *string `pulumi:"status"`
}

type DnatRuleState struct {
	// Dnat rule creation time.
	CreatedAt pulumi.StringPtrInput
	// Specifies the description of the dnat rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// Changing this creates a new dnat rule.
	Description pulumi.StringPtrInput
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	ExternalServicePort pulumi.IntPtrInput
	// The actual floating IP address.
	FloatingIpAddress pulumi.StringPtrInput
	// Specifies the ID of the floating IP address. Changing this creates a
	// new dnat rule.
	FloatingIpId pulumi.StringPtrInput
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	InternalServicePort pulumi.IntPtrInput
	// ID of the nat gateway this dnat rule belongs to. Changing this creates
	// a new dnat rule.
	NatGatewayId pulumi.StringPtrInput
	// Specifies the port ID of network. This parameter is mandatory in VPC
	// scenario. Use Vpc.Port to get the port if just know a fixed IP
	// addresses on the port. Changing this creates a new dnat rule.
	PortId pulumi.StringPtrInput
	// Specifies the private IP address of a user. This parameter is mandatory in
	// Direct Connect scenario. Changing this creates a new dnat rule.
	PrivateIp pulumi.StringPtrInput
	// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
	// Changing this creates a new dnat rule.
	Protocol pulumi.StringPtrInput
	// The region in which to create the dnat rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new dnat rule.
	Region pulumi.StringPtrInput
	// Dnat rule status.
	Status pulumi.StringPtrInput
}

func (DnatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatRuleState)(nil)).Elem()
}

type dnatRuleArgs struct {
	// Specifies the description of the dnat rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// Changing this creates a new dnat rule.
	Description *string `pulumi:"description"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	ExternalServicePort int `pulumi:"externalServicePort"`
	// Specifies the ID of the floating IP address. Changing this creates a
	// new dnat rule.
	FloatingIpId string `pulumi:"floatingIpId"`
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	InternalServicePort int `pulumi:"internalServicePort"`
	// ID of the nat gateway this dnat rule belongs to. Changing this creates
	// a new dnat rule.
	NatGatewayId string `pulumi:"natGatewayId"`
	// Specifies the port ID of network. This parameter is mandatory in VPC
	// scenario. Use Vpc.Port to get the port if just know a fixed IP
	// addresses on the port. Changing this creates a new dnat rule.
	PortId *string `pulumi:"portId"`
	// Specifies the private IP address of a user. This parameter is mandatory in
	// Direct Connect scenario. Changing this creates a new dnat rule.
	PrivateIp *string `pulumi:"privateIp"`
	// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
	// Changing this creates a new dnat rule.
	Protocol string `pulumi:"protocol"`
	// The region in which to create the dnat rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new dnat rule.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DnatRule resource.
type DnatRuleArgs struct {
	// Specifies the description of the dnat rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// Changing this creates a new dnat rule.
	Description pulumi.StringPtrInput
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	ExternalServicePort pulumi.IntInput
	// Specifies the ID of the floating IP address. Changing this creates a
	// new dnat rule.
	FloatingIpId pulumi.StringInput
	// Specifies port used by ECSs or BMSs to provide services for
	// external systems. Changing this creates a new dnat rule.
	InternalServicePort pulumi.IntInput
	// ID of the nat gateway this dnat rule belongs to. Changing this creates
	// a new dnat rule.
	NatGatewayId pulumi.StringInput
	// Specifies the port ID of network. This parameter is mandatory in VPC
	// scenario. Use Vpc.Port to get the port if just know a fixed IP
	// addresses on the port. Changing this creates a new dnat rule.
	PortId pulumi.StringPtrInput
	// Specifies the private IP address of a user. This parameter is mandatory in
	// Direct Connect scenario. Changing this creates a new dnat rule.
	PrivateIp pulumi.StringPtrInput
	// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
	// Changing this creates a new dnat rule.
	Protocol pulumi.StringInput
	// The region in which to create the dnat rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new dnat rule.
	Region pulumi.StringPtrInput
}

func (DnatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatRuleArgs)(nil)).Elem()
}

type DnatRuleInput interface {
	pulumi.Input

	ToDnatRuleOutput() DnatRuleOutput
	ToDnatRuleOutputWithContext(ctx context.Context) DnatRuleOutput
}

func (*DnatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DnatRule)(nil)).Elem()
}

func (i *DnatRule) ToDnatRuleOutput() DnatRuleOutput {
	return i.ToDnatRuleOutputWithContext(context.Background())
}

func (i *DnatRule) ToDnatRuleOutputWithContext(ctx context.Context) DnatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatRuleOutput)
}

// DnatRuleArrayInput is an input type that accepts DnatRuleArray and DnatRuleArrayOutput values.
// You can construct a concrete instance of `DnatRuleArrayInput` via:
//
//	DnatRuleArray{ DnatRuleArgs{...} }
type DnatRuleArrayInput interface {
	pulumi.Input

	ToDnatRuleArrayOutput() DnatRuleArrayOutput
	ToDnatRuleArrayOutputWithContext(context.Context) DnatRuleArrayOutput
}

type DnatRuleArray []DnatRuleInput

func (DnatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnatRule)(nil)).Elem()
}

func (i DnatRuleArray) ToDnatRuleArrayOutput() DnatRuleArrayOutput {
	return i.ToDnatRuleArrayOutputWithContext(context.Background())
}

func (i DnatRuleArray) ToDnatRuleArrayOutputWithContext(ctx context.Context) DnatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatRuleArrayOutput)
}

// DnatRuleMapInput is an input type that accepts DnatRuleMap and DnatRuleMapOutput values.
// You can construct a concrete instance of `DnatRuleMapInput` via:
//
//	DnatRuleMap{ "key": DnatRuleArgs{...} }
type DnatRuleMapInput interface {
	pulumi.Input

	ToDnatRuleMapOutput() DnatRuleMapOutput
	ToDnatRuleMapOutputWithContext(context.Context) DnatRuleMapOutput
}

type DnatRuleMap map[string]DnatRuleInput

func (DnatRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnatRule)(nil)).Elem()
}

func (i DnatRuleMap) ToDnatRuleMapOutput() DnatRuleMapOutput {
	return i.ToDnatRuleMapOutputWithContext(context.Background())
}

func (i DnatRuleMap) ToDnatRuleMapOutputWithContext(ctx context.Context) DnatRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatRuleMapOutput)
}

type DnatRuleOutput struct{ *pulumi.OutputState }

func (DnatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnatRule)(nil)).Elem()
}

func (o DnatRuleOutput) ToDnatRuleOutput() DnatRuleOutput {
	return o
}

func (o DnatRuleOutput) ToDnatRuleOutputWithContext(ctx context.Context) DnatRuleOutput {
	return o
}

// Dnat rule creation time.
func (o DnatRuleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the description of the dnat rule.
// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
// Changing this creates a new dnat rule.
func (o DnatRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies port used by ECSs or BMSs to provide services for
// external systems. Changing this creates a new dnat rule.
func (o DnatRuleOutput) ExternalServicePort() pulumi.IntOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.IntOutput { return v.ExternalServicePort }).(pulumi.IntOutput)
}

// The actual floating IP address.
func (o DnatRuleOutput) FloatingIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.FloatingIpAddress }).(pulumi.StringOutput)
}

// Specifies the ID of the floating IP address. Changing this creates a
// new dnat rule.
func (o DnatRuleOutput) FloatingIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.FloatingIpId }).(pulumi.StringOutput)
}

// Specifies port used by ECSs or BMSs to provide services for
// external systems. Changing this creates a new dnat rule.
func (o DnatRuleOutput) InternalServicePort() pulumi.IntOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.IntOutput { return v.InternalServicePort }).(pulumi.IntOutput)
}

// ID of the nat gateway this dnat rule belongs to. Changing this creates
// a new dnat rule.
func (o DnatRuleOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.NatGatewayId }).(pulumi.StringOutput)
}

// Specifies the port ID of network. This parameter is mandatory in VPC
// scenario. Use Vpc.Port to get the port if just know a fixed IP
// addresses on the port. Changing this creates a new dnat rule.
func (o DnatRuleOutput) PortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringPtrOutput { return v.PortId }).(pulumi.StringPtrOutput)
}

// Specifies the private IP address of a user. This parameter is mandatory in
// Direct Connect scenario. Changing this creates a new dnat rule.
func (o DnatRuleOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringPtrOutput { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

// Specifies the protocol type. Currently, TCP, UDP, and ANY are supported.
// Changing this creates a new dnat rule.
func (o DnatRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The region in which to create the dnat rule resource. If omitted, the
// provider-level region will be used. Changing this creates a new dnat rule.
func (o DnatRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Dnat rule status.
func (o DnatRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type DnatRuleArrayOutput struct{ *pulumi.OutputState }

func (DnatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnatRule)(nil)).Elem()
}

func (o DnatRuleArrayOutput) ToDnatRuleArrayOutput() DnatRuleArrayOutput {
	return o
}

func (o DnatRuleArrayOutput) ToDnatRuleArrayOutputWithContext(ctx context.Context) DnatRuleArrayOutput {
	return o
}

func (o DnatRuleArrayOutput) Index(i pulumi.IntInput) DnatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnatRule {
		return vs[0].([]*DnatRule)[vs[1].(int)]
	}).(DnatRuleOutput)
}

type DnatRuleMapOutput struct{ *pulumi.OutputState }

func (DnatRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnatRule)(nil)).Elem()
}

func (o DnatRuleMapOutput) ToDnatRuleMapOutput() DnatRuleMapOutput {
	return o
}

func (o DnatRuleMapOutput) ToDnatRuleMapOutputWithContext(ctx context.Context) DnatRuleMapOutput {
	return o
}

func (o DnatRuleMapOutput) MapIndex(k pulumi.StringInput) DnatRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnatRule {
		return vs[0].(map[string]*DnatRule)[vs[1].(string)]
	}).(DnatRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnatRuleInput)(nil)).Elem(), &DnatRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnatRuleArrayInput)(nil)).Elem(), DnatRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnatRuleMapInput)(nil)).Elem(), DnatRuleMap{})
	pulumi.RegisterOutputType(DnatRuleOutput{})
	pulumi.RegisterOutputType(DnatRuleArrayOutput{})
	pulumi.RegisterOutputType(DnatRuleMapOutput{})
}
