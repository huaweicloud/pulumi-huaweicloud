// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VPC network interface resource within HuaweiCloud.
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
//			name := cfg.RequireObject("name")
//			subnetId := cfg.RequireObject("subnetId")
//			_, err := Vpc.NewNetworkInterface(ctx, "test", &Vpc.NetworkInterfaceArgs{
//				SubnetId: pulumi.Any(subnetId),
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
// The network interface can be imported using `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpc/networkInterface:NetworkInterface test <id>
//
// ```
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Specifies an array of IP addresses that can be active on the
	// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
	// check switch is turned off.
	AllowedAddresses pulumi.StringArrayOutput `pulumi:"allowedAddresses"`
	// Indicates the availability zone to which the network interface belongs.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Indicates the ID of the device to which the network interface belongs.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Indicates the belonged device, which can be the DHCP server, router, load balancer, or Nova.
	DeviceOwner pulumi.StringOutput `pulumi:"deviceOwner"`
	// Specifies the DHCP lease time. The value format of value is "Xh",
	// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
	DhcpLeaseTime pulumi.StringOutput `pulumi:"dhcpLeaseTime"`
	// Indicates the default private network DNS name of the primary NIC.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Indicates whether to enable EFI.
	EnableEfi pulumi.BoolOutput `pulumi:"enableEfi"`
	// Specifies the network interface IPv4 address.
	// Changing this creates a new resource.
	FixedIpV4 pulumi.StringOutput `pulumi:"fixedIpV4"`
	// Indicates the ID of the instance to which the network interface belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicates the type of the instance to which the network interface belongs.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Indicates the Shared bandwidth ID bound to IPv6 network interface.
	Ipv6BandwidthId pulumi.StringOutput `pulumi:"ipv6BandwidthId"`
	// Indicates the network interface MAC address.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// Specifies the network interface name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the security option is enabled for the network interface.
	// If the option is not enabled, the security group and DHCP snooping do not take effect.
	PortSecurityEnabled pulumi.BoolOutput `pulumi:"portSecurityEnabled"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Indicates the network interface status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the ID of the subnet to which the network interface belongs.
	// Changing this creates a new resource.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Specifies the network interface tags in the format of key-value pairs.
	// This parameter can only be used in **cn-south-2** for now.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NetworkInterface
	err := ctx.RegisterResource("huaweicloud:Vpc/networkInterface:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("huaweicloud:Vpc/networkInterface:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// Specifies an array of IP addresses that can be active on the
	// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
	// check switch is turned off.
	AllowedAddresses []string `pulumi:"allowedAddresses"`
	// Indicates the availability zone to which the network interface belongs.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Indicates the ID of the device to which the network interface belongs.
	DeviceId *string `pulumi:"deviceId"`
	// Indicates the belonged device, which can be the DHCP server, router, load balancer, or Nova.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// Specifies the DHCP lease time. The value format of value is "Xh",
	// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
	DhcpLeaseTime *string `pulumi:"dhcpLeaseTime"`
	// Indicates the default private network DNS name of the primary NIC.
	DnsName *string `pulumi:"dnsName"`
	// Indicates whether to enable EFI.
	EnableEfi *bool `pulumi:"enableEfi"`
	// Specifies the network interface IPv4 address.
	// Changing this creates a new resource.
	FixedIpV4 *string `pulumi:"fixedIpV4"`
	// Indicates the ID of the instance to which the network interface belongs.
	InstanceId *string `pulumi:"instanceId"`
	// Indicates the type of the instance to which the network interface belongs.
	InstanceType *string `pulumi:"instanceType"`
	// Indicates the Shared bandwidth ID bound to IPv6 network interface.
	Ipv6BandwidthId *string `pulumi:"ipv6BandwidthId"`
	// Indicates the network interface MAC address.
	MacAddress *string `pulumi:"macAddress"`
	// Specifies the network interface name.
	Name *string `pulumi:"name"`
	// Indicates whether the security option is enabled for the network interface.
	// If the option is not enabled, the security group and DHCP snooping do not take effect.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies an array of one or more security group IDs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Indicates the network interface status.
	Status *string `pulumi:"status"`
	// Specifies the ID of the subnet to which the network interface belongs.
	// Changing this creates a new resource.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the network interface tags in the format of key-value pairs.
	// This parameter can only be used in **cn-south-2** for now.
	Tags map[string]string `pulumi:"tags"`
}

type NetworkInterfaceState struct {
	// Specifies an array of IP addresses that can be active on the
	// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
	// check switch is turned off.
	AllowedAddresses pulumi.StringArrayInput
	// Indicates the availability zone to which the network interface belongs.
	AvailabilityZone pulumi.StringPtrInput
	// Indicates the ID of the device to which the network interface belongs.
	DeviceId pulumi.StringPtrInput
	// Indicates the belonged device, which can be the DHCP server, router, load balancer, or Nova.
	DeviceOwner pulumi.StringPtrInput
	// Specifies the DHCP lease time. The value format of value is "Xh",
	// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
	DhcpLeaseTime pulumi.StringPtrInput
	// Indicates the default private network DNS name of the primary NIC.
	DnsName pulumi.StringPtrInput
	// Indicates whether to enable EFI.
	EnableEfi pulumi.BoolPtrInput
	// Specifies the network interface IPv4 address.
	// Changing this creates a new resource.
	FixedIpV4 pulumi.StringPtrInput
	// Indicates the ID of the instance to which the network interface belongs.
	InstanceId pulumi.StringPtrInput
	// Indicates the type of the instance to which the network interface belongs.
	InstanceType pulumi.StringPtrInput
	// Indicates the Shared bandwidth ID bound to IPv6 network interface.
	Ipv6BandwidthId pulumi.StringPtrInput
	// Indicates the network interface MAC address.
	MacAddress pulumi.StringPtrInput
	// Specifies the network interface name.
	Name pulumi.StringPtrInput
	// Indicates whether the security option is enabled for the network interface.
	// If the option is not enabled, the security group and DHCP snooping do not take effect.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayInput
	// Indicates the network interface status.
	Status pulumi.StringPtrInput
	// Specifies the ID of the subnet to which the network interface belongs.
	// Changing this creates a new resource.
	SubnetId pulumi.StringPtrInput
	// Specifies the network interface tags in the format of key-value pairs.
	// This parameter can only be used in **cn-south-2** for now.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Specifies an array of IP addresses that can be active on the
	// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
	// check switch is turned off.
	AllowedAddresses []string `pulumi:"allowedAddresses"`
	// Specifies the DHCP lease time. The value format of value is "Xh",
	// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
	DhcpLeaseTime *string `pulumi:"dhcpLeaseTime"`
	// Specifies the network interface IPv4 address.
	// Changing this creates a new resource.
	FixedIpV4 *string `pulumi:"fixedIpV4"`
	// Specifies the network interface name.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies an array of one or more security group IDs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Specifies the ID of the subnet to which the network interface belongs.
	// Changing this creates a new resource.
	SubnetId string `pulumi:"subnetId"`
	// Specifies the network interface tags in the format of key-value pairs.
	// This parameter can only be used in **cn-south-2** for now.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Specifies an array of IP addresses that can be active on the
	// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
	// check switch is turned off.
	AllowedAddresses pulumi.StringArrayInput
	// Specifies the DHCP lease time. The value format of value is "Xh",
	// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
	DhcpLeaseTime pulumi.StringPtrInput
	// Specifies the network interface IPv4 address.
	// Changing this creates a new resource.
	FixedIpV4 pulumi.StringPtrInput
	// Specifies the network interface name.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayInput
	// Specifies the ID of the subnet to which the network interface belongs.
	// Changing this creates a new resource.
	SubnetId pulumi.StringInput
	// Specifies the network interface tags in the format of key-value pairs.
	// This parameter can only be used in **cn-south-2** for now.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

// NetworkInterfaceArrayInput is an input type that accepts NetworkInterfaceArray and NetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceArrayInput` via:
//
//	NetworkInterfaceArray{ NetworkInterfaceArgs{...} }
type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

// NetworkInterfaceMapInput is an input type that accepts NetworkInterfaceMap and NetworkInterfaceMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceMapInput` via:
//
//	NetworkInterfaceMap{ "key": NetworkInterfaceArgs{...} }
type NetworkInterfaceMapInput interface {
	pulumi.Input

	ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput
	ToNetworkInterfaceMapOutputWithContext(context.Context) NetworkInterfaceMapOutput
}

type NetworkInterfaceMap map[string]NetworkInterfaceInput

func (NetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return i.ToNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceMapOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

// Specifies an array of IP addresses that can be active on the
// network interface. If the IP address is "1.1.1.1/0", it means that the source/destination address
// check switch is turned off.
func (o NetworkInterfaceOutput) AllowedAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.AllowedAddresses }).(pulumi.StringArrayOutput)
}

// Indicates the availability zone to which the network interface belongs.
func (o NetworkInterfaceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Indicates the ID of the device to which the network interface belongs.
func (o NetworkInterfaceOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// Indicates the belonged device, which can be the DHCP server, router, load balancer, or Nova.
func (o NetworkInterfaceOutput) DeviceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.DeviceOwner }).(pulumi.StringOutput)
}

// Specifies the DHCP lease time. The value format of value is "Xh",
// the value of "X" is "-1" or from "1" to "30000". If the value is "-1", the DHCP lease time is infinite.
func (o NetworkInterfaceOutput) DhcpLeaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.DhcpLeaseTime }).(pulumi.StringOutput)
}

// Indicates the default private network DNS name of the primary NIC.
func (o NetworkInterfaceOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// Indicates whether to enable EFI.
func (o NetworkInterfaceOutput) EnableEfi() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolOutput { return v.EnableEfi }).(pulumi.BoolOutput)
}

// Specifies the network interface IPv4 address.
// Changing this creates a new resource.
func (o NetworkInterfaceOutput) FixedIpV4() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.FixedIpV4 }).(pulumi.StringOutput)
}

// Indicates the ID of the instance to which the network interface belongs.
func (o NetworkInterfaceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicates the type of the instance to which the network interface belongs.
func (o NetworkInterfaceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// Indicates the Shared bandwidth ID bound to IPv6 network interface.
func (o NetworkInterfaceOutput) Ipv6BandwidthId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Ipv6BandwidthId }).(pulumi.StringOutput)
}

// Indicates the network interface MAC address.
func (o NetworkInterfaceOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// Specifies the network interface name.
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the security option is enabled for the network interface.
// If the option is not enabled, the security group and DHCP snooping do not take effect.
func (o NetworkInterfaceOutput) PortSecurityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolOutput { return v.PortSecurityEnabled }).(pulumi.BoolOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used.
// Changing this creates a new resource.
func (o NetworkInterfaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies an array of one or more security group IDs.
func (o NetworkInterfaceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Indicates the network interface status.
func (o NetworkInterfaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the ID of the subnet to which the network interface belongs.
// Changing this creates a new resource.
func (o NetworkInterfaceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Specifies the network interface tags in the format of key-value pairs.
// This parameter can only be used in **cn-south-2** for now.
func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkInterface {
		return vs[0].([]*NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkInterface {
		return vs[0].(map[string]*NetworkInterface)[vs[1].(string)]
	}).(NetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceInput)(nil)).Elem(), &NetworkInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceArrayInput)(nil)).Elem(), NetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceMapInput)(nil)).Elem(), NetworkInterfaceMap{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceMapOutput{})
}
