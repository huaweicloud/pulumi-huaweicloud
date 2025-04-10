// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB listener resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Elb.NewListener(ctx, "listener1", &Elb.ListenerArgs{
//				LoadbalancerId: pulumi.String("d9415786-5f1a-428b-b35f-2f1523e146d2"),
//				Protocol:       pulumi.String("HTTP"),
//				ProtocolPort:   pulumi.Int(8080),
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("value"),
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
// ELB listener can be imported using the listener ID, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Elb/listener:Listener listener_1 5c20fdad-7288-11eb-b817-0255ac10158b
//
// ```
type Listener struct {
	pulumi.CustomResourceState

	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The maximum number of connections allowed for the listener. The value ranges from
	// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
	// number of connections.
	ConnectionLimit pulumi.IntOutput `pulumi:"connectionLimit"`
	// The ID of the default pool with which the listener is associated.
	// Changing this creates a new listener.
	DefaultPoolId pulumi.StringOutput `pulumi:"defaultPoolId"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
	DefaultTlsContainerRef pulumi.StringPtrOutput `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the listener.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *TERMINATED_HTTPS*.
	Http2Enable pulumi.BoolPtrOutput `pulumi:"http2Enable"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Human-readable name for the listener. Does not have to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
	// creates a new listener.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntOutput `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringOutput `pulumi:"region"`
	// Lists the IDs of SNI certificates (server certificates with a domain name)
	// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
	SniContainerRefs pulumi.StringArrayOutput `pulumi:"sniContainerRefs"`
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProtocolPort == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolPort'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Listener
	err := ctx.RegisterResource("huaweicloud:Elb/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("huaweicloud:Elb/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The maximum number of connections allowed for the listener. The value ranges from
	// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
	// number of connections.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// The ID of the default pool with which the listener is associated.
	// Changing this creates a new listener.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
	DefaultTlsContainerRef *string `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the listener.
	Description *string `pulumi:"description"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *TERMINATED_HTTPS*.
	Http2Enable *bool `pulumi:"http2Enable"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the listener. Does not have to be unique.
	Name *string `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
	// creates a new listener.
	Protocol *string `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort *int `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region *string `pulumi:"region"`
	// Lists the IDs of SNI certificates (server certificates with a domain name)
	// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
	SniContainerRefs []string `pulumi:"sniContainerRefs"`
	// The key/value pairs to associate with the listener.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
}

type ListenerState struct {
	AdminStateUp pulumi.BoolPtrInput
	// The maximum number of connections allowed for the listener. The value ranges from
	// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
	// number of connections.
	ConnectionLimit pulumi.IntPtrInput
	// The ID of the default pool with which the listener is associated.
	// Changing this creates a new listener.
	DefaultPoolId pulumi.StringPtrInput
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
	DefaultTlsContainerRef pulumi.StringPtrInput
	// Human-readable description for the listener.
	Description pulumi.StringPtrInput
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *TERMINATED_HTTPS*.
	Http2Enable pulumi.BoolPtrInput
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the listener. Does not have to be unique.
	Name pulumi.StringPtrInput
	// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
	// creates a new listener.
	Protocol pulumi.StringPtrInput
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntPtrInput
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringPtrInput
	// Lists the IDs of SNI certificates (server certificates with a domain name)
	// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
	SniContainerRefs pulumi.StringArrayInput
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The maximum number of connections allowed for the listener. The value ranges from
	// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
	// number of connections.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// The ID of the default pool with which the listener is associated.
	// Changing this creates a new listener.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
	DefaultTlsContainerRef *string `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the listener.
	Description *string `pulumi:"description"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *TERMINATED_HTTPS*.
	Http2Enable *bool `pulumi:"http2Enable"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Human-readable name for the listener. Does not have to be unique.
	Name *string `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
	// creates a new listener.
	Protocol string `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort int `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region *string `pulumi:"region"`
	// Lists the IDs of SNI certificates (server certificates with a domain name)
	// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
	SniContainerRefs []string `pulumi:"sniContainerRefs"`
	// The key/value pairs to associate with the listener.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	AdminStateUp pulumi.BoolPtrInput
	// The maximum number of connections allowed for the listener. The value ranges from
	// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
	// number of connections.
	ConnectionLimit pulumi.IntPtrInput
	// The ID of the default pool with which the listener is associated.
	// Changing this creates a new listener.
	DefaultPoolId pulumi.StringPtrInput
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
	DefaultTlsContainerRef pulumi.StringPtrInput
	// Human-readable description for the listener.
	Description pulumi.StringPtrInput
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *TERMINATED_HTTPS*.
	Http2Enable pulumi.BoolPtrInput
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringInput
	// Human-readable name for the listener. Does not have to be unique.
	Name pulumi.StringPtrInput
	// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
	// creates a new listener.
	Protocol pulumi.StringInput
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntInput
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringPtrInput
	// Lists the IDs of SNI certificates (server certificates with a domain name)
	// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
	SniContainerRefs pulumi.StringArrayInput
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//	ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//	ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

func (o ListenerOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The maximum number of connections allowed for the listener. The value ranges from
// -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum
// number of connections.
func (o ListenerOutput) ConnectionLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.ConnectionLimit }).(pulumi.IntOutput)
}

// The ID of the default pool with which the listener is associated.
// Changing this creates a new listener.
func (o ListenerOutput) DefaultPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.DefaultPoolId }).(pulumi.StringOutput)
}

// Specifies the ID of the server certificate used by the listener. This
// parameter is mandatory when protocol is set to *TERMINATED_HTTPS*.
func (o ListenerOutput) DefaultTlsContainerRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.DefaultTlsContainerRef }).(pulumi.StringPtrOutput)
}

// Human-readable description for the listener.
func (o ListenerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
// only when the protocol is set to *TERMINATED_HTTPS*.
func (o ListenerOutput) Http2Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.Http2Enable }).(pulumi.BoolPtrOutput)
}

// The load balancer on which to provision this listener. Changing this
// creates a new listener.
func (o ListenerOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// Human-readable name for the listener. Does not have to be unique.
func (o ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this
// creates a new listener.
func (o ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The port on which to listen for client traffic. Changing this creates a
// new listener.
func (o ListenerOutput) ProtocolPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.ProtocolPort }).(pulumi.IntOutput)
}

// The region in which to create the listener resource. If omitted, the
// provider-level region will be used. Changing this creates a new listener.
func (o ListenerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Lists the IDs of SNI certificates (server certificates with a domain name)
// used by the listener. This parameter is valid when protocol is set to *TERMINATED_HTTPS*.
func (o ListenerOutput) SniContainerRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.SniContainerRefs }).(pulumi.StringArrayOutput)
}

// The key/value pairs to associate with the listener.
func (o ListenerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: tenant_id is deprecated
func (o ListenerOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].([]*Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].(map[string]*Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerArrayInput)(nil)).Elem(), ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerMapInput)(nil)).Elem(), ListenerMap{})
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
