// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a VPC Peering Connection resource.
//
// > **NOTE:** For cross-tenant (requester's tenant differs from the accepter's tenant) VPC Peering Connections,
//
//	use the `Vpc.PeeringConnection` resource to manage the requester's side of the connection and
//	use the `Vpc.PeeringConnectionAccepter` resource to manage the accepter's side of the connection.
//	<br/>If you create a VPC peering connection with another VPC of your own, the connection is created without the need
//	for you to accept the connection.
//
// ## Example Usage
//
// ## Import
//
// VPC Peering resources can be imported using the `vpc peering id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpc/peeringConnection:PeeringConnection test_connection 22b76469-08e3-4937-8c1d-7aad34892be1
//
// ```
type PeeringConnection struct {
	pulumi.CustomResourceState

	// Specifies the description of the VPC peering connection.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the name of the VPC peering connection. The value can contain 1 to 64
	// characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the tenant ID of the accepter tenant. Changing this creates
	// a new VPC peering connection.
	PeerTenantId pulumi.StringOutput `pulumi:"peerTenantId"`
	// Specifies the VPC ID of the accepter tenant. Changing this creates a new
	// VPC peering connection.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`
	// The region in which to create the VPC peering connection. If omitted, the
	// provider-level region will be used. Changing this creates a new VPC peering connection resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or
	// ACTIVE.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
	// creates a new VPC peering connection.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewPeeringConnection(ctx *pulumi.Context,
	name string, args *PeeringConnectionArgs, opts ...pulumi.ResourceOption) (*PeeringConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerVpcId == nil {
		return nil, errors.New("invalid value for required argument 'PeerVpcId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PeeringConnection
	err := ctx.RegisterResource("huaweicloud:Vpc/peeringConnection:PeeringConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringConnection gets an existing PeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringConnectionState, opts ...pulumi.ResourceOption) (*PeeringConnection, error) {
	var resource PeeringConnection
	err := ctx.ReadResource("huaweicloud:Vpc/peeringConnection:PeeringConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringConnection resources.
type peeringConnectionState struct {
	// Specifies the description of the VPC peering connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the VPC peering connection. The value can contain 1 to 64
	// characters.
	Name *string `pulumi:"name"`
	// Specifies the tenant ID of the accepter tenant. Changing this creates
	// a new VPC peering connection.
	PeerTenantId *string `pulumi:"peerTenantId"`
	// Specifies the VPC ID of the accepter tenant. Changing this creates a new
	// VPC peering connection.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// The region in which to create the VPC peering connection. If omitted, the
	// provider-level region will be used. Changing this creates a new VPC peering connection resource.
	Region *string `pulumi:"region"`
	// The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or
	// ACTIVE.
	Status *string `pulumi:"status"`
	// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
	// creates a new VPC peering connection.
	VpcId *string `pulumi:"vpcId"`
}

type PeeringConnectionState struct {
	// Specifies the description of the VPC peering connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the VPC peering connection. The value can contain 1 to 64
	// characters.
	Name pulumi.StringPtrInput
	// Specifies the tenant ID of the accepter tenant. Changing this creates
	// a new VPC peering connection.
	PeerTenantId pulumi.StringPtrInput
	// Specifies the VPC ID of the accepter tenant. Changing this creates a new
	// VPC peering connection.
	PeerVpcId pulumi.StringPtrInput
	// The region in which to create the VPC peering connection. If omitted, the
	// provider-level region will be used. Changing this creates a new VPC peering connection resource.
	Region pulumi.StringPtrInput
	// The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or
	// ACTIVE.
	Status pulumi.StringPtrInput
	// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
	// creates a new VPC peering connection.
	VpcId pulumi.StringPtrInput
}

func (PeeringConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionState)(nil)).Elem()
}

type peeringConnectionArgs struct {
	// Specifies the description of the VPC peering connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the VPC peering connection. The value can contain 1 to 64
	// characters.
	Name *string `pulumi:"name"`
	// Specifies the tenant ID of the accepter tenant. Changing this creates
	// a new VPC peering connection.
	PeerTenantId *string `pulumi:"peerTenantId"`
	// Specifies the VPC ID of the accepter tenant. Changing this creates a new
	// VPC peering connection.
	PeerVpcId string `pulumi:"peerVpcId"`
	// The region in which to create the VPC peering connection. If omitted, the
	// provider-level region will be used. Changing this creates a new VPC peering connection resource.
	Region *string `pulumi:"region"`
	// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
	// creates a new VPC peering connection.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a PeeringConnection resource.
type PeeringConnectionArgs struct {
	// Specifies the description of the VPC peering connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the VPC peering connection. The value can contain 1 to 64
	// characters.
	Name pulumi.StringPtrInput
	// Specifies the tenant ID of the accepter tenant. Changing this creates
	// a new VPC peering connection.
	PeerTenantId pulumi.StringPtrInput
	// Specifies the VPC ID of the accepter tenant. Changing this creates a new
	// VPC peering connection.
	PeerVpcId pulumi.StringInput
	// The region in which to create the VPC peering connection. If omitted, the
	// provider-level region will be used. Changing this creates a new VPC peering connection resource.
	Region pulumi.StringPtrInput
	// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
	// creates a new VPC peering connection.
	VpcId pulumi.StringInput
}

func (PeeringConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionArgs)(nil)).Elem()
}

type PeeringConnectionInput interface {
	pulumi.Input

	ToPeeringConnectionOutput() PeeringConnectionOutput
	ToPeeringConnectionOutputWithContext(ctx context.Context) PeeringConnectionOutput
}

func (*PeeringConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringConnection)(nil)).Elem()
}

func (i *PeeringConnection) ToPeeringConnectionOutput() PeeringConnectionOutput {
	return i.ToPeeringConnectionOutputWithContext(context.Background())
}

func (i *PeeringConnection) ToPeeringConnectionOutputWithContext(ctx context.Context) PeeringConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionOutput)
}

// PeeringConnectionArrayInput is an input type that accepts PeeringConnectionArray and PeeringConnectionArrayOutput values.
// You can construct a concrete instance of `PeeringConnectionArrayInput` via:
//
//	PeeringConnectionArray{ PeeringConnectionArgs{...} }
type PeeringConnectionArrayInput interface {
	pulumi.Input

	ToPeeringConnectionArrayOutput() PeeringConnectionArrayOutput
	ToPeeringConnectionArrayOutputWithContext(context.Context) PeeringConnectionArrayOutput
}

type PeeringConnectionArray []PeeringConnectionInput

func (PeeringConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeeringConnection)(nil)).Elem()
}

func (i PeeringConnectionArray) ToPeeringConnectionArrayOutput() PeeringConnectionArrayOutput {
	return i.ToPeeringConnectionArrayOutputWithContext(context.Background())
}

func (i PeeringConnectionArray) ToPeeringConnectionArrayOutputWithContext(ctx context.Context) PeeringConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionArrayOutput)
}

// PeeringConnectionMapInput is an input type that accepts PeeringConnectionMap and PeeringConnectionMapOutput values.
// You can construct a concrete instance of `PeeringConnectionMapInput` via:
//
//	PeeringConnectionMap{ "key": PeeringConnectionArgs{...} }
type PeeringConnectionMapInput interface {
	pulumi.Input

	ToPeeringConnectionMapOutput() PeeringConnectionMapOutput
	ToPeeringConnectionMapOutputWithContext(context.Context) PeeringConnectionMapOutput
}

type PeeringConnectionMap map[string]PeeringConnectionInput

func (PeeringConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeeringConnection)(nil)).Elem()
}

func (i PeeringConnectionMap) ToPeeringConnectionMapOutput() PeeringConnectionMapOutput {
	return i.ToPeeringConnectionMapOutputWithContext(context.Background())
}

func (i PeeringConnectionMap) ToPeeringConnectionMapOutputWithContext(ctx context.Context) PeeringConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionMapOutput)
}

type PeeringConnectionOutput struct{ *pulumi.OutputState }

func (PeeringConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringConnection)(nil)).Elem()
}

func (o PeeringConnectionOutput) ToPeeringConnectionOutput() PeeringConnectionOutput {
	return o
}

func (o PeeringConnectionOutput) ToPeeringConnectionOutputWithContext(ctx context.Context) PeeringConnectionOutput {
	return o
}

// Specifies the description of the VPC peering connection.
func (o PeeringConnectionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the name of the VPC peering connection. The value can contain 1 to 64
// characters.
func (o PeeringConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the tenant ID of the accepter tenant. Changing this creates
// a new VPC peering connection.
func (o PeeringConnectionOutput) PeerTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.PeerTenantId }).(pulumi.StringOutput)
}

// Specifies the VPC ID of the accepter tenant. Changing this creates a new
// VPC peering connection.
func (o PeeringConnectionOutput) PeerVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.PeerVpcId }).(pulumi.StringOutput)
}

// The region in which to create the VPC peering connection. If omitted, the
// provider-level region will be used. Changing this creates a new VPC peering connection resource.
func (o PeeringConnectionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or
// ACTIVE.
func (o PeeringConnectionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the ID of a VPC involved in a VPC peering connection. Changing this
// creates a new VPC peering connection.
func (o PeeringConnectionOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnection) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type PeeringConnectionArrayOutput struct{ *pulumi.OutputState }

func (PeeringConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeeringConnection)(nil)).Elem()
}

func (o PeeringConnectionArrayOutput) ToPeeringConnectionArrayOutput() PeeringConnectionArrayOutput {
	return o
}

func (o PeeringConnectionArrayOutput) ToPeeringConnectionArrayOutputWithContext(ctx context.Context) PeeringConnectionArrayOutput {
	return o
}

func (o PeeringConnectionArrayOutput) Index(i pulumi.IntInput) PeeringConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PeeringConnection {
		return vs[0].([]*PeeringConnection)[vs[1].(int)]
	}).(PeeringConnectionOutput)
}

type PeeringConnectionMapOutput struct{ *pulumi.OutputState }

func (PeeringConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeeringConnection)(nil)).Elem()
}

func (o PeeringConnectionMapOutput) ToPeeringConnectionMapOutput() PeeringConnectionMapOutput {
	return o
}

func (o PeeringConnectionMapOutput) ToPeeringConnectionMapOutputWithContext(ctx context.Context) PeeringConnectionMapOutput {
	return o
}

func (o PeeringConnectionMapOutput) MapIndex(k pulumi.StringInput) PeeringConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PeeringConnection {
		return vs[0].(map[string]*PeeringConnection)[vs[1].(string)]
	}).(PeeringConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionInput)(nil)).Elem(), &PeeringConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionArrayInput)(nil)).Elem(), PeeringConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionMapInput)(nil)).Elem(), PeeringConnectionMap{})
	pulumi.RegisterOutputType(PeeringConnectionOutput{})
	pulumi.RegisterOutputType(PeeringConnectionArrayOutput{})
	pulumi.RegisterOutputType(PeeringConnectionMapOutput{})
}
