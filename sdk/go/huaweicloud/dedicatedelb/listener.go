// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB listener resource within HuaweiCloud.
//
// ## Import
//
// ELB listener can be imported using the listener ID, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedElb/listener:Listener listener_1 5c20fdad-7288-11eb-b817-0255ac10158b
//
// ```
type Listener struct {
	pulumi.CustomResourceState

	// Specifies the access policy for the listener. Valid options are *white* and
	// *black*.
	AccessPolicy pulumi.StringPtrOutput `pulumi:"accessPolicy"`
	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	AdvancedForwardingEnabled pulumi.BoolOutput `pulumi:"advancedForwardingEnabled"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is
	// valid when protocol is set to *HTTPS*.
	CaCertificate pulumi.StringPtrOutput `pulumi:"caCertificate"`
	// The ID of the default pool with which the listener is associated. Changing this
	// creates a new listener.
	DefaultPoolId pulumi.StringOutput `pulumi:"defaultPoolId"`
	// Human-readable description for the listener.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
	// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
	// *HTTPS*.
	ForwardEip pulumi.BoolPtrOutput `pulumi:"forwardEip"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *HTTPS*.
	Http2Enable pulumi.BoolPtrOutput `pulumi:"http2Enable"`
	// Specifies the idle timeout for the listener. Value range: 0 to 4000.
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`
	// Specifies the ip group id for the listener.
	IpGroup pulumi.StringPtrOutput `pulumi:"ipGroup"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Human-readable name for the listener.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
	// new listener.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntOutput `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the request timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	RequestTimeout pulumi.IntOutput `pulumi:"requestTimeout"`
	// Specifies the response timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	ResponseTimeout pulumi.IntOutput `pulumi:"responseTimeout"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *HTTPS*.
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// Lists the IDs of SNI certificates (server certificates with a domain name) used
	// by the listener. This parameter is valid when protocol is set to *HTTPS*.
	SniCertificates pulumi.StringArrayOutput `pulumi:"sniCertificates"`
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the TLS cipher policy for the listener. Valid options are:
	// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
	// This parameter is valid when protocol is set to *HTTPS*.
	TlsCiphersPolicy pulumi.StringOutput `pulumi:"tlsCiphersPolicy"`
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
	err := ctx.RegisterResource("huaweicloud:DedicatedElb/listener:Listener", name, args, &resource, opts...)
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
	err := ctx.ReadResource("huaweicloud:DedicatedElb/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// Specifies the access policy for the listener. Valid options are *white* and
	// *black*.
	AccessPolicy *string `pulumi:"accessPolicy"`
	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	AdvancedForwardingEnabled *bool `pulumi:"advancedForwardingEnabled"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is
	// valid when protocol is set to *HTTPS*.
	CaCertificate *string `pulumi:"caCertificate"`
	// The ID of the default pool with which the listener is associated. Changing this
	// creates a new listener.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// Human-readable description for the listener.
	Description *string `pulumi:"description"`
	// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
	// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
	// *HTTPS*.
	ForwardEip *bool `pulumi:"forwardEip"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *HTTPS*.
	Http2Enable *bool `pulumi:"http2Enable"`
	// Specifies the idle timeout for the listener. Value range: 0 to 4000.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the ip group id for the listener.
	IpGroup *string `pulumi:"ipGroup"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the listener.
	Name *string `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
	// new listener.
	Protocol *string `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort *int `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region *string `pulumi:"region"`
	// Specifies the request timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// Specifies the response timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	ResponseTimeout *int `pulumi:"responseTimeout"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *HTTPS*.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Lists the IDs of SNI certificates (server certificates with a domain name) used
	// by the listener. This parameter is valid when protocol is set to *HTTPS*.
	SniCertificates []string `pulumi:"sniCertificates"`
	// The key/value pairs to associate with the listener.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the TLS cipher policy for the listener. Valid options are:
	// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
	// This parameter is valid when protocol is set to *HTTPS*.
	TlsCiphersPolicy *string `pulumi:"tlsCiphersPolicy"`
}

type ListenerState struct {
	// Specifies the access policy for the listener. Valid options are *white* and
	// *black*.
	AccessPolicy pulumi.StringPtrInput
	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	AdvancedForwardingEnabled pulumi.BoolPtrInput
	// Specifies the ID of the CA certificate used by the listener. This parameter is
	// valid when protocol is set to *HTTPS*.
	CaCertificate pulumi.StringPtrInput
	// The ID of the default pool with which the listener is associated. Changing this
	// creates a new listener.
	DefaultPoolId pulumi.StringPtrInput
	// Human-readable description for the listener.
	Description pulumi.StringPtrInput
	// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
	// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
	// *HTTPS*.
	ForwardEip pulumi.BoolPtrInput
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *HTTPS*.
	Http2Enable pulumi.BoolPtrInput
	// Specifies the idle timeout for the listener. Value range: 0 to 4000.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the ip group id for the listener.
	IpGroup pulumi.StringPtrInput
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the listener.
	Name pulumi.StringPtrInput
	// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
	// new listener.
	Protocol pulumi.StringPtrInput
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntPtrInput
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringPtrInput
	// Specifies the request timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	RequestTimeout pulumi.IntPtrInput
	// Specifies the response timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	ResponseTimeout pulumi.IntPtrInput
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *HTTPS*.
	ServerCertificate pulumi.StringPtrInput
	// Lists the IDs of SNI certificates (server certificates with a domain name) used
	// by the listener. This parameter is valid when protocol is set to *HTTPS*.
	SniCertificates pulumi.StringArrayInput
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapInput
	// Specifies the TLS cipher policy for the listener. Valid options are:
	// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
	// This parameter is valid when protocol is set to *HTTPS*.
	TlsCiphersPolicy pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// Specifies the access policy for the listener. Valid options are *white* and
	// *black*.
	AccessPolicy *string `pulumi:"accessPolicy"`
	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	AdvancedForwardingEnabled *bool `pulumi:"advancedForwardingEnabled"`
	// Specifies the ID of the CA certificate used by the listener. This parameter is
	// valid when protocol is set to *HTTPS*.
	CaCertificate *string `pulumi:"caCertificate"`
	// The ID of the default pool with which the listener is associated. Changing this
	// creates a new listener.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// Human-readable description for the listener.
	Description *string `pulumi:"description"`
	// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
	// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
	// *HTTPS*.
	ForwardEip *bool `pulumi:"forwardEip"`
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *HTTPS*.
	Http2Enable *bool `pulumi:"http2Enable"`
	// Specifies the idle timeout for the listener. Value range: 0 to 4000.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the ip group id for the listener.
	IpGroup *string `pulumi:"ipGroup"`
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Human-readable name for the listener.
	Name *string `pulumi:"name"`
	// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
	// new listener.
	Protocol string `pulumi:"protocol"`
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort int `pulumi:"protocolPort"`
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region *string `pulumi:"region"`
	// Specifies the request timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// Specifies the response timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	ResponseTimeout *int `pulumi:"responseTimeout"`
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *HTTPS*.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Lists the IDs of SNI certificates (server certificates with a domain name) used
	// by the listener. This parameter is valid when protocol is set to *HTTPS*.
	SniCertificates []string `pulumi:"sniCertificates"`
	// The key/value pairs to associate with the listener.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the TLS cipher policy for the listener. Valid options are:
	// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
	// This parameter is valid when protocol is set to *HTTPS*.
	TlsCiphersPolicy *string `pulumi:"tlsCiphersPolicy"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// Specifies the access policy for the listener. Valid options are *white* and
	// *black*.
	AccessPolicy pulumi.StringPtrInput
	// Specifies whether to enable advanced forwarding.
	// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
	AdvancedForwardingEnabled pulumi.BoolPtrInput
	// Specifies the ID of the CA certificate used by the listener. This parameter is
	// valid when protocol is set to *HTTPS*.
	CaCertificate pulumi.StringPtrInput
	// The ID of the default pool with which the listener is associated. Changing this
	// creates a new listener.
	DefaultPoolId pulumi.StringPtrInput
	// Human-readable description for the listener.
	Description pulumi.StringPtrInput
	// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
	// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
	// *HTTPS*.
	ForwardEip pulumi.BoolPtrInput
	// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
	// only when the protocol is set to *HTTPS*.
	Http2Enable pulumi.BoolPtrInput
	// Specifies the idle timeout for the listener. Value range: 0 to 4000.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the ip group id for the listener.
	IpGroup pulumi.StringPtrInput
	// The load balancer on which to provision this listener. Changing this
	// creates a new listener.
	LoadbalancerId pulumi.StringInput
	// Human-readable name for the listener.
	Name pulumi.StringPtrInput
	// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
	// new listener.
	Protocol pulumi.StringInput
	// The port on which to listen for client traffic. Changing this creates a
	// new listener.
	ProtocolPort pulumi.IntInput
	// The region in which to create the listener resource. If omitted, the
	// provider-level region will be used. Changing this creates a new listener.
	Region pulumi.StringPtrInput
	// Specifies the request timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	RequestTimeout pulumi.IntPtrInput
	// Specifies the response timeout for the listener. Value range: 1 to 300. This
	// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
	ResponseTimeout pulumi.IntPtrInput
	// Specifies the ID of the server certificate used by the listener. This
	// parameter is mandatory when protocol is set to *HTTPS*.
	ServerCertificate pulumi.StringPtrInput
	// Lists the IDs of SNI certificates (server certificates with a domain name) used
	// by the listener. This parameter is valid when protocol is set to *HTTPS*.
	SniCertificates pulumi.StringArrayInput
	// The key/value pairs to associate with the listener.
	Tags pulumi.StringMapInput
	// Specifies the TLS cipher policy for the listener. Valid options are:
	// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
	// This parameter is valid when protocol is set to *HTTPS*.
	TlsCiphersPolicy pulumi.StringPtrInput
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

// Specifies the access policy for the listener. Valid options are *white* and
// *black*.
func (o ListenerOutput) AccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.AccessPolicy }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable advanced forwarding.
// If advanced forwarding is enabled, more flexible forwarding policies and rules are supported.
func (o ListenerOutput) AdvancedForwardingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolOutput { return v.AdvancedForwardingEnabled }).(pulumi.BoolOutput)
}

// Specifies the ID of the CA certificate used by the listener. This parameter is
// valid when protocol is set to *HTTPS*.
func (o ListenerOutput) CaCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.CaCertificate }).(pulumi.StringPtrOutput)
}

// The ID of the default pool with which the listener is associated. Changing this
// creates a new listener.
func (o ListenerOutput) DefaultPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.DefaultPoolId }).(pulumi.StringOutput)
}

// Human-readable description for the listener.
func (o ListenerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
// backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or
// *HTTPS*.
func (o ListenerOutput) ForwardEip() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.ForwardEip }).(pulumi.BoolPtrOutput)
}

// Specifies whether to use HTTP/2. The default value is false. This parameter is valid
// only when the protocol is set to *HTTPS*.
func (o ListenerOutput) Http2Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.Http2Enable }).(pulumi.BoolPtrOutput)
}

// Specifies the idle timeout for the listener. Value range: 0 to 4000.
func (o ListenerOutput) IdleTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.IdleTimeout }).(pulumi.IntOutput)
}

// Specifies the ip group id for the listener.
func (o ListenerOutput) IpGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.IpGroup }).(pulumi.StringPtrOutput)
}

// The load balancer on which to provision this listener. Changing this
// creates a new listener.
func (o ListenerOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// Human-readable name for the listener.
func (o ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
// new listener.
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

// Specifies the request timeout for the listener. Value range: 1 to 300. This
// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
func (o ListenerOutput) RequestTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.RequestTimeout }).(pulumi.IntOutput)
}

// Specifies the response timeout for the listener. Value range: 1 to 300. This
// parameter is valid when protocol is set to *HTTP* or *HTTPS*.
func (o ListenerOutput) ResponseTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.ResponseTimeout }).(pulumi.IntOutput)
}

// Specifies the ID of the server certificate used by the listener. This
// parameter is mandatory when protocol is set to *HTTPS*.
func (o ListenerOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

// Lists the IDs of SNI certificates (server certificates with a domain name) used
// by the listener. This parameter is valid when protocol is set to *HTTPS*.
func (o ListenerOutput) SniCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.SniCertificates }).(pulumi.StringArrayOutput)
}

// The key/value pairs to associate with the listener.
func (o ListenerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the TLS cipher policy for the listener. Valid options are:
// tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
// This parameter is valid when protocol is set to *HTTPS*.
func (o ListenerOutput) TlsCiphersPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.TlsCiphersPolicy }).(pulumi.StringOutput)
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