// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB active-standby pool resource within HuaweiCloud.
//
// ## Example Usage
// ### Create an active-standby Pool
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			_, err := DedicatedElb.NewActiveStandbyPool(ctx, "test", &DedicatedElb.ActiveStandbyPoolArgs{
//				Description:   pulumi.String("test description"),
//				Protocol:      pulumi.String("TCP"),
//				VpcId:         pulumi.Any(vpcId),
//				Type:          pulumi.String("instance"),
//				AnyPortEnable: pulumi.Bool(false),
//				Members: dedicatedelb.ActiveStandbyPoolMemberArray{
//					&dedicatedelb.ActiveStandbyPoolMemberArgs{
//						Address:      pulumi.String("192.168.0.1"),
//						Role:         pulumi.String("master"),
//						ProtocolPort: pulumi.Int(45),
//					},
//					&dedicatedelb.ActiveStandbyPoolMemberArgs{
//						Address:      pulumi.String("192.168.0.2"),
//						Role:         pulumi.String("slave"),
//						ProtocolPort: pulumi.Int(36),
//					},
//				},
//				Healthmonitor: &dedicatedelb.ActiveStandbyPoolHealthmonitorArgs{
//					Delay:          pulumi.Int(5),
//					ExpectedCodes:  pulumi.String("200"),
//					MaxRetries:     pulumi.Int(3),
//					MaxRetriesDown: pulumi.Int(3),
//					Timeout:        pulumi.Int(3),
//					Type:           pulumi.String("TCP"),
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
// ELB active-standby pool can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedElb/activeStandbyPool:ActiveStandbyPool test <id>
//
// ```
type ActiveStandbyPool struct {
	pulumi.CustomResourceState

	// Specifies whether to enable forward to same port for active-standby
	// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
	// frontend port. Value options:
	// + **false**: Disable forward to same port.
	// + **true**: Enable forward to same port.
	AnyPortEnable pulumi.BoolOutput `pulumi:"anyPortEnable"`
	// Specifies whether to enable delayed logout. This parameter can
	// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
	// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
	// + The pool member is removed from the pool.
	// + The health monitor status is abnormal.
	// + The pool member weight is changed to 0.
	ConnectionDrainEnabled pulumi.BoolOutput `pulumi:"connectionDrainEnabled"`
	// Specifies the timeout of the delayed logout in seconds. Value
	// ranges from `10` to `4,000`.
	ConnectionDrainTimeout pulumi.IntOutput `pulumi:"connectionDrainTimeout"`
	// The create time of the active-standby pool.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the description of the active-standby pool. Changing this
	// parameter will create a new resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the health check configured for the active-standby pool.
	// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
	Healthmonitor ActiveStandbyPoolHealthmonitorOutput `pulumi:"healthmonitor"`
	// Specifies the IP address version supported by active-standby pool.
	// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// Specifies the ID of the listener with which the active-standby pool is
	// associated. Changing this parameter will create a new resource.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Specifies the ID of the load balancer with which the active-standby
	// pool is associated. Changing this parameter will create a new resource.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the members in the active-standby pool.
	// The members structure is documented below. Changing this parameter will create a new resource.
	Members ActiveStandbyPoolMemberArrayOutput `pulumi:"members"`
	// Specifies the health check name. The length range of value is from `1` to `255`.
	// Changing this parameter will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the protocol used by the active-standby pool to receive requests.
	// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
	// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
	// + If the listener's protocol is **TCP**, the value must be **TCP**.
	// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Specifies the region in which to create the ELB active-standby pool resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
	// **HTTP**, and **HTTPS**.
	// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	Type pulumi.StringOutput `pulumi:"type"`
	// The update time of the active-standby pool.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Specifies the ID of the VPC where the active-standby pool works. Changing this
	// parameter will create a new resource.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewActiveStandbyPool registers a new resource with the given unique name, arguments, and options.
func NewActiveStandbyPool(ctx *pulumi.Context,
	name string, args *ActiveStandbyPoolArgs, opts ...pulumi.ResourceOption) (*ActiveStandbyPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Healthmonitor == nil {
		return nil, errors.New("invalid value for required argument 'Healthmonitor'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ActiveStandbyPool
	err := ctx.RegisterResource("huaweicloud:DedicatedElb/activeStandbyPool:ActiveStandbyPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveStandbyPool gets an existing ActiveStandbyPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveStandbyPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveStandbyPoolState, opts ...pulumi.ResourceOption) (*ActiveStandbyPool, error) {
	var resource ActiveStandbyPool
	err := ctx.ReadResource("huaweicloud:DedicatedElb/activeStandbyPool:ActiveStandbyPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveStandbyPool resources.
type activeStandbyPoolState struct {
	// Specifies whether to enable forward to same port for active-standby
	// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
	// frontend port. Value options:
	// + **false**: Disable forward to same port.
	// + **true**: Enable forward to same port.
	AnyPortEnable *bool `pulumi:"anyPortEnable"`
	// Specifies whether to enable delayed logout. This parameter can
	// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
	// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
	// + The pool member is removed from the pool.
	// + The health monitor status is abnormal.
	// + The pool member weight is changed to 0.
	ConnectionDrainEnabled *bool `pulumi:"connectionDrainEnabled"`
	// Specifies the timeout of the delayed logout in seconds. Value
	// ranges from `10` to `4,000`.
	ConnectionDrainTimeout *int `pulumi:"connectionDrainTimeout"`
	// The create time of the active-standby pool.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the description of the active-standby pool. Changing this
	// parameter will create a new resource.
	Description *string `pulumi:"description"`
	// Specifies the health check configured for the active-standby pool.
	// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
	Healthmonitor *ActiveStandbyPoolHealthmonitor `pulumi:"healthmonitor"`
	// Specifies the IP address version supported by active-standby pool.
	// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the ID of the listener with which the active-standby pool is
	// associated. Changing this parameter will create a new resource.
	ListenerId *string `pulumi:"listenerId"`
	// Specifies the ID of the load balancer with which the active-standby
	// pool is associated. Changing this parameter will create a new resource.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the members in the active-standby pool.
	// The members structure is documented below. Changing this parameter will create a new resource.
	Members []ActiveStandbyPoolMember `pulumi:"members"`
	// Specifies the health check name. The length range of value is from `1` to `255`.
	// Changing this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// Specifies the protocol used by the active-standby pool to receive requests.
	// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
	// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
	// + If the listener's protocol is **TCP**, the value must be **TCP**.
	// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
	Protocol *string `pulumi:"protocol"`
	// Specifies the region in which to create the ELB active-standby pool resource.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
	// **HTTP**, and **HTTPS**.
	// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	Type *string `pulumi:"type"`
	// The update time of the active-standby pool.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Specifies the ID of the VPC where the active-standby pool works. Changing this
	// parameter will create a new resource.
	VpcId *string `pulumi:"vpcId"`
}

type ActiveStandbyPoolState struct {
	// Specifies whether to enable forward to same port for active-standby
	// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
	// frontend port. Value options:
	// + **false**: Disable forward to same port.
	// + **true**: Enable forward to same port.
	AnyPortEnable pulumi.BoolPtrInput
	// Specifies whether to enable delayed logout. This parameter can
	// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
	// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
	// + The pool member is removed from the pool.
	// + The health monitor status is abnormal.
	// + The pool member weight is changed to 0.
	ConnectionDrainEnabled pulumi.BoolPtrInput
	// Specifies the timeout of the delayed logout in seconds. Value
	// ranges from `10` to `4,000`.
	ConnectionDrainTimeout pulumi.IntPtrInput
	// The create time of the active-standby pool.
	CreatedAt pulumi.StringPtrInput
	// Specifies the description of the active-standby pool. Changing this
	// parameter will create a new resource.
	Description pulumi.StringPtrInput
	// Specifies the health check configured for the active-standby pool.
	// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
	Healthmonitor ActiveStandbyPoolHealthmonitorPtrInput
	// Specifies the IP address version supported by active-standby pool.
	// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
	IpVersion pulumi.StringPtrInput
	// Specifies the ID of the listener with which the active-standby pool is
	// associated. Changing this parameter will create a new resource.
	ListenerId pulumi.StringPtrInput
	// Specifies the ID of the load balancer with which the active-standby
	// pool is associated. Changing this parameter will create a new resource.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the members in the active-standby pool.
	// The members structure is documented below. Changing this parameter will create a new resource.
	Members ActiveStandbyPoolMemberArrayInput
	// Specifies the health check name. The length range of value is from `1` to `255`.
	// Changing this parameter will create a new resource.
	Name pulumi.StringPtrInput
	// Specifies the protocol used by the active-standby pool to receive requests.
	// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
	// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
	// + If the listener's protocol is **TCP**, the value must be **TCP**.
	// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
	Protocol pulumi.StringPtrInput
	// Specifies the region in which to create the ELB active-standby pool resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput
	// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
	// **HTTP**, and **HTTPS**.
	// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	Type pulumi.StringPtrInput
	// The update time of the active-standby pool.
	UpdatedAt pulumi.StringPtrInput
	// Specifies the ID of the VPC where the active-standby pool works. Changing this
	// parameter will create a new resource.
	VpcId pulumi.StringPtrInput
}

func (ActiveStandbyPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeStandbyPoolState)(nil)).Elem()
}

type activeStandbyPoolArgs struct {
	// Specifies whether to enable forward to same port for active-standby
	// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
	// frontend port. Value options:
	// + **false**: Disable forward to same port.
	// + **true**: Enable forward to same port.
	AnyPortEnable *bool `pulumi:"anyPortEnable"`
	// Specifies whether to enable delayed logout. This parameter can
	// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
	// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
	// + The pool member is removed from the pool.
	// + The health monitor status is abnormal.
	// + The pool member weight is changed to 0.
	ConnectionDrainEnabled *bool `pulumi:"connectionDrainEnabled"`
	// Specifies the timeout of the delayed logout in seconds. Value
	// ranges from `10` to `4,000`.
	ConnectionDrainTimeout *int `pulumi:"connectionDrainTimeout"`
	// Specifies the description of the active-standby pool. Changing this
	// parameter will create a new resource.
	Description *string `pulumi:"description"`
	// Specifies the health check configured for the active-standby pool.
	// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
	Healthmonitor ActiveStandbyPoolHealthmonitor `pulumi:"healthmonitor"`
	// Specifies the IP address version supported by active-standby pool.
	// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the ID of the listener with which the active-standby pool is
	// associated. Changing this parameter will create a new resource.
	ListenerId *string `pulumi:"listenerId"`
	// Specifies the ID of the load balancer with which the active-standby
	// pool is associated. Changing this parameter will create a new resource.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the members in the active-standby pool.
	// The members structure is documented below. Changing this parameter will create a new resource.
	Members []ActiveStandbyPoolMember `pulumi:"members"`
	// Specifies the health check name. The length range of value is from `1` to `255`.
	// Changing this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// Specifies the protocol used by the active-standby pool to receive requests.
	// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
	// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
	// + If the listener's protocol is **TCP**, the value must be **TCP**.
	// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
	Protocol string `pulumi:"protocol"`
	// Specifies the region in which to create the ELB active-standby pool resource.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
	// **HTTP**, and **HTTPS**.
	// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	Type *string `pulumi:"type"`
	// Specifies the ID of the VPC where the active-standby pool works. Changing this
	// parameter will create a new resource.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ActiveStandbyPool resource.
type ActiveStandbyPoolArgs struct {
	// Specifies whether to enable forward to same port for active-standby
	// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
	// frontend port. Value options:
	// + **false**: Disable forward to same port.
	// + **true**: Enable forward to same port.
	AnyPortEnable pulumi.BoolPtrInput
	// Specifies whether to enable delayed logout. This parameter can
	// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
	// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
	// + The pool member is removed from the pool.
	// + The health monitor status is abnormal.
	// + The pool member weight is changed to 0.
	ConnectionDrainEnabled pulumi.BoolPtrInput
	// Specifies the timeout of the delayed logout in seconds. Value
	// ranges from `10` to `4,000`.
	ConnectionDrainTimeout pulumi.IntPtrInput
	// Specifies the description of the active-standby pool. Changing this
	// parameter will create a new resource.
	Description pulumi.StringPtrInput
	// Specifies the health check configured for the active-standby pool.
	// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
	Healthmonitor ActiveStandbyPoolHealthmonitorInput
	// Specifies the IP address version supported by active-standby pool.
	// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
	IpVersion pulumi.StringPtrInput
	// Specifies the ID of the listener with which the active-standby pool is
	// associated. Changing this parameter will create a new resource.
	ListenerId pulumi.StringPtrInput
	// Specifies the ID of the load balancer with which the active-standby
	// pool is associated. Changing this parameter will create a new resource.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the members in the active-standby pool.
	// The members structure is documented below. Changing this parameter will create a new resource.
	Members ActiveStandbyPoolMemberArrayInput
	// Specifies the health check name. The length range of value is from `1` to `255`.
	// Changing this parameter will create a new resource.
	Name pulumi.StringPtrInput
	// Specifies the protocol used by the active-standby pool to receive requests.
	// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
	// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
	// + If the listener's protocol is **TCP**, the value must be **TCP**.
	// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
	Protocol pulumi.StringInput
	// Specifies the region in which to create the ELB active-standby pool resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput
	// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
	// **HTTP**, and **HTTPS**.
	// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
	// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
	Type pulumi.StringPtrInput
	// Specifies the ID of the VPC where the active-standby pool works. Changing this
	// parameter will create a new resource.
	VpcId pulumi.StringPtrInput
}

func (ActiveStandbyPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeStandbyPoolArgs)(nil)).Elem()
}

type ActiveStandbyPoolInput interface {
	pulumi.Input

	ToActiveStandbyPoolOutput() ActiveStandbyPoolOutput
	ToActiveStandbyPoolOutputWithContext(ctx context.Context) ActiveStandbyPoolOutput
}

func (*ActiveStandbyPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveStandbyPool)(nil)).Elem()
}

func (i *ActiveStandbyPool) ToActiveStandbyPoolOutput() ActiveStandbyPoolOutput {
	return i.ToActiveStandbyPoolOutputWithContext(context.Background())
}

func (i *ActiveStandbyPool) ToActiveStandbyPoolOutputWithContext(ctx context.Context) ActiveStandbyPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveStandbyPoolOutput)
}

// ActiveStandbyPoolArrayInput is an input type that accepts ActiveStandbyPoolArray and ActiveStandbyPoolArrayOutput values.
// You can construct a concrete instance of `ActiveStandbyPoolArrayInput` via:
//
//	ActiveStandbyPoolArray{ ActiveStandbyPoolArgs{...} }
type ActiveStandbyPoolArrayInput interface {
	pulumi.Input

	ToActiveStandbyPoolArrayOutput() ActiveStandbyPoolArrayOutput
	ToActiveStandbyPoolArrayOutputWithContext(context.Context) ActiveStandbyPoolArrayOutput
}

type ActiveStandbyPoolArray []ActiveStandbyPoolInput

func (ActiveStandbyPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActiveStandbyPool)(nil)).Elem()
}

func (i ActiveStandbyPoolArray) ToActiveStandbyPoolArrayOutput() ActiveStandbyPoolArrayOutput {
	return i.ToActiveStandbyPoolArrayOutputWithContext(context.Background())
}

func (i ActiveStandbyPoolArray) ToActiveStandbyPoolArrayOutputWithContext(ctx context.Context) ActiveStandbyPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveStandbyPoolArrayOutput)
}

// ActiveStandbyPoolMapInput is an input type that accepts ActiveStandbyPoolMap and ActiveStandbyPoolMapOutput values.
// You can construct a concrete instance of `ActiveStandbyPoolMapInput` via:
//
//	ActiveStandbyPoolMap{ "key": ActiveStandbyPoolArgs{...} }
type ActiveStandbyPoolMapInput interface {
	pulumi.Input

	ToActiveStandbyPoolMapOutput() ActiveStandbyPoolMapOutput
	ToActiveStandbyPoolMapOutputWithContext(context.Context) ActiveStandbyPoolMapOutput
}

type ActiveStandbyPoolMap map[string]ActiveStandbyPoolInput

func (ActiveStandbyPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActiveStandbyPool)(nil)).Elem()
}

func (i ActiveStandbyPoolMap) ToActiveStandbyPoolMapOutput() ActiveStandbyPoolMapOutput {
	return i.ToActiveStandbyPoolMapOutputWithContext(context.Background())
}

func (i ActiveStandbyPoolMap) ToActiveStandbyPoolMapOutputWithContext(ctx context.Context) ActiveStandbyPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveStandbyPoolMapOutput)
}

type ActiveStandbyPoolOutput struct{ *pulumi.OutputState }

func (ActiveStandbyPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveStandbyPool)(nil)).Elem()
}

func (o ActiveStandbyPoolOutput) ToActiveStandbyPoolOutput() ActiveStandbyPoolOutput {
	return o
}

func (o ActiveStandbyPoolOutput) ToActiveStandbyPoolOutputWithContext(ctx context.Context) ActiveStandbyPoolOutput {
	return o
}

// Specifies whether to enable forward to same port for active-standby
// pool. If this option is enabled, the listener routes the requests to the backend server over the same port as the
// frontend port. Value options:
// + **false**: Disable forward to same port.
// + **true**: Enable forward to same port.
func (o ActiveStandbyPoolOutput) AnyPortEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.BoolOutput { return v.AnyPortEnable }).(pulumi.BoolOutput)
}

// Specifies whether to enable delayed logout. This parameter can
// be set to **true** when the `protocol` is set to **TCP**, **UDP** or **QUIC**, and the value of `protocol` of the
// associated listener must be **TCP** or **UDP**. It will be triggered for the following scenes:
// + The pool member is removed from the pool.
// + The health monitor status is abnormal.
// + The pool member weight is changed to 0.
func (o ActiveStandbyPoolOutput) ConnectionDrainEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.BoolOutput { return v.ConnectionDrainEnabled }).(pulumi.BoolOutput)
}

// Specifies the timeout of the delayed logout in seconds. Value
// ranges from `10` to `4,000`.
func (o ActiveStandbyPoolOutput) ConnectionDrainTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.IntOutput { return v.ConnectionDrainTimeout }).(pulumi.IntOutput)
}

// The create time of the active-standby pool.
func (o ActiveStandbyPoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the description of the active-standby pool. Changing this
// parameter will create a new resource.
func (o ActiveStandbyPoolOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the health check configured for the active-standby pool.
// The healthmonitor structure is documented below. Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) Healthmonitor() ActiveStandbyPoolHealthmonitorOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) ActiveStandbyPoolHealthmonitorOutput { return v.Healthmonitor }).(ActiveStandbyPoolHealthmonitorOutput)
}

// Specifies the IP address version supported by active-standby pool.
// The value can be **dualstack**, **v6**, or **v4**. Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// Specifies the ID of the listener with which the active-standby pool is
// associated. Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// Specifies the ID of the load balancer with which the active-standby
// pool is associated. Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// Specifies the members in the active-standby pool.
// The members structure is documented below. Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) Members() ActiveStandbyPoolMemberArrayOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) ActiveStandbyPoolMemberArrayOutput { return v.Members }).(ActiveStandbyPoolMemberArrayOutput)
}

// Specifies the health check name. The length range of value is from `1` to `255`.
// Changing this parameter will create a new resource.
func (o ActiveStandbyPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the protocol used by the active-standby pool to receive requests.
// Value options: **TCP**, **UDP**, **QUIC** or **TLS**.
// + If the listener's protocol is **UDP**, the value must be **UDP** or **QUIC**.
// + If the listener's protocol is **TCP**, the value must be **TCP**.
// + If the listener's protocol is **TLS**, the value must be **TLS** or **TCP**.
func (o ActiveStandbyPoolOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Specifies the region in which to create the ELB active-standby pool resource.
// If omitted, the provider-level region will be used.
func (o ActiveStandbyPoolOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the health check protocol. Value options: **TCP**, **UDP_CONNECT**,
// **HTTP**, and **HTTPS**.
// + If the protocol of the backend server is **QUIC**, the value can only be **UDP_CONNECT**.
// + If the protocol of the backend server is **UDP**, the value can only be **UDP_CONNECT**.
// + If the protocol of the backend server is **TCP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
// + If the protocol of the backend server is **HTTP**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
// + If the protocol of the backend server is **HTTPS**, the value can only be **TCP**, **HTTP**, or **HTTPS**.
func (o ActiveStandbyPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The update time of the active-standby pool.
func (o ActiveStandbyPoolOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Specifies the ID of the VPC where the active-standby pool works. Changing this
// parameter will create a new resource.
func (o ActiveStandbyPoolOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveStandbyPool) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ActiveStandbyPoolArrayOutput struct{ *pulumi.OutputState }

func (ActiveStandbyPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActiveStandbyPool)(nil)).Elem()
}

func (o ActiveStandbyPoolArrayOutput) ToActiveStandbyPoolArrayOutput() ActiveStandbyPoolArrayOutput {
	return o
}

func (o ActiveStandbyPoolArrayOutput) ToActiveStandbyPoolArrayOutputWithContext(ctx context.Context) ActiveStandbyPoolArrayOutput {
	return o
}

func (o ActiveStandbyPoolArrayOutput) Index(i pulumi.IntInput) ActiveStandbyPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActiveStandbyPool {
		return vs[0].([]*ActiveStandbyPool)[vs[1].(int)]
	}).(ActiveStandbyPoolOutput)
}

type ActiveStandbyPoolMapOutput struct{ *pulumi.OutputState }

func (ActiveStandbyPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActiveStandbyPool)(nil)).Elem()
}

func (o ActiveStandbyPoolMapOutput) ToActiveStandbyPoolMapOutput() ActiveStandbyPoolMapOutput {
	return o
}

func (o ActiveStandbyPoolMapOutput) ToActiveStandbyPoolMapOutputWithContext(ctx context.Context) ActiveStandbyPoolMapOutput {
	return o
}

func (o ActiveStandbyPoolMapOutput) MapIndex(k pulumi.StringInput) ActiveStandbyPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActiveStandbyPool {
		return vs[0].(map[string]*ActiveStandbyPool)[vs[1].(string)]
	}).(ActiveStandbyPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveStandbyPoolInput)(nil)).Elem(), &ActiveStandbyPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveStandbyPoolArrayInput)(nil)).Elem(), ActiveStandbyPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActiveStandbyPoolMapInput)(nil)).Elem(), ActiveStandbyPoolMap{})
	pulumi.RegisterOutputType(ActiveStandbyPoolOutput{})
	pulumi.RegisterOutputType(ActiveStandbyPoolArrayOutput{})
	pulumi.RegisterOutputType(ActiveStandbyPoolMapOutput{})
}
