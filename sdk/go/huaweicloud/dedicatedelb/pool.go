// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB pool resource within HuaweiCloud.
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DedicatedElb.NewPool(ctx, "pool1", &DedicatedElb.PoolArgs{
//				LbMethod:   pulumi.String("ROUND_ROBIN"),
//				ListenerId: pulumi.String("{{ listener_id }}"),
//				Persistences: dedicatedelb.PoolPersistenceArray{
//					&dedicatedelb.PoolPersistenceArgs{
//						CookieName: pulumi.String("testCookie"),
//						Type:       pulumi.String("HTTP_COOKIE"),
//					},
//				},
//				Protocol: pulumi.String("HTTP"),
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
// ELB pool can be imported using the pool ID, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedElb/pool:Pool pool_1 5c20fdad-7288-11eb-b817-0255ac10158b
//
// ```
type Pool struct {
	pulumi.CustomResourceState

	// Human-readable description for the pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The load balancing algorithm to distribute traffic to the pool's members. Must be one
	// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
	LbMethod pulumi.StringOutput `pulumi:"lbMethod"`
	// The Listener on which the members of the pool will be associated with.
	// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The load balancer on which to provision this pool. Changing this
	// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Omit this field to prevent session persistence. Indicates whether
	// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
	Persistences PoolPersistenceArrayOutput `pulumi:"persistences"`
	// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to create the ELB pool resource. If omitted, the the
	// provider-level region will be used. Changing this creates a new pool.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbMethod == nil {
		return nil, errors.New("invalid value for required argument 'LbMethod'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("huaweicloud:DedicatedElb/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("huaweicloud:DedicatedElb/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// Human-readable description for the pool.
	Description *string `pulumi:"description"`
	// The load balancing algorithm to distribute traffic to the pool's members. Must be one
	// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
	LbMethod *string `pulumi:"lbMethod"`
	// The Listener on which the members of the pool will be associated with.
	// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	ListenerId *string `pulumi:"listenerId"`
	// The load balancer on which to provision this pool. Changing this
	// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name *string `pulumi:"name"`
	// Omit this field to prevent session persistence. Indicates whether
	// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
	Persistences []PoolPersistence `pulumi:"persistences"`
	// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
	Protocol *string `pulumi:"protocol"`
	// The region in which to create the ELB pool resource. If omitted, the the
	// provider-level region will be used. Changing this creates a new pool.
	Region *string `pulumi:"region"`
}

type PoolState struct {
	// Human-readable description for the pool.
	Description pulumi.StringPtrInput
	// The load balancing algorithm to distribute traffic to the pool's members. Must be one
	// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
	LbMethod pulumi.StringPtrInput
	// The Listener on which the members of the pool will be associated with.
	// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringPtrInput
	// The load balancer on which to provision this pool. Changing this
	// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the pool.
	Name pulumi.StringPtrInput
	// Omit this field to prevent session persistence. Indicates whether
	// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
	Persistences PoolPersistenceArrayInput
	// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
	Protocol pulumi.StringPtrInput
	// The region in which to create the ELB pool resource. If omitted, the the
	// provider-level region will be used. Changing this creates a new pool.
	Region pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Human-readable description for the pool.
	Description *string `pulumi:"description"`
	// The load balancing algorithm to distribute traffic to the pool's members. Must be one
	// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
	LbMethod string `pulumi:"lbMethod"`
	// The Listener on which the members of the pool will be associated with.
	// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	ListenerId *string `pulumi:"listenerId"`
	// The load balancer on which to provision this pool. Changing this
	// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name *string `pulumi:"name"`
	// Omit this field to prevent session persistence. Indicates whether
	// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
	Persistences []PoolPersistence `pulumi:"persistences"`
	// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
	Protocol string `pulumi:"protocol"`
	// The region in which to create the ELB pool resource. If omitted, the the
	// provider-level region will be used. Changing this creates a new pool.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Human-readable description for the pool.
	Description pulumi.StringPtrInput
	// The load balancing algorithm to distribute traffic to the pool's members. Must be one
	// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
	LbMethod pulumi.StringInput
	// The Listener on which the members of the pool will be associated with.
	// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringPtrInput
	// The load balancer on which to provision this pool. Changing this
	// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the pool.
	Name pulumi.StringPtrInput
	// Omit this field to prevent session persistence. Indicates whether
	// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
	Persistences PoolPersistenceArrayInput
	// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
	Protocol pulumi.StringInput
	// The region in which to create the ELB pool resource. If omitted, the the
	// provider-level region will be used. Changing this creates a new pool.
	Region pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//	PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//	PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// Human-readable description for the pool.
func (o PoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The load balancing algorithm to distribute traffic to the pool's members. Must be one
// of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
func (o PoolOutput) LbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LbMethod }).(pulumi.StringOutput)
}

// The Listener on which the members of the pool will be associated with.
// Changing this creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
func (o PoolOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The load balancer on which to provision this pool. Changing this
// creates a new pool. Note:  Exactly one of LoadbalancerID or ListenerID must be provided.
func (o PoolOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// Human-readable name for the pool.
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Omit this field to prevent session persistence. Indicates whether
// connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.
func (o PoolOutput) Persistences() PoolPersistenceArrayOutput {
	return o.ApplyT(func(v *Pool) PoolPersistenceArrayOutput { return v.Persistences }).(PoolPersistenceArrayOutput)
}

// The protocol - can either be TCP, UDP, HTTP, HTTPS or QUIC.
func (o PoolOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The region in which to create the ELB pool resource. If omitted, the the
// provider-level region will be used. Changing this creates a new pool.
func (o PoolOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].([]*Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].(map[string]*Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PoolInput)(nil)).Elem(), &Pool{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolArrayInput)(nil)).Elem(), PoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolMapInput)(nil)).Elem(), PoolMap{})
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}