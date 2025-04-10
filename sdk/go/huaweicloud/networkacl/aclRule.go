// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkacl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **WARNING:** It has been deprecated, use `Vpc.NetworkAcl` instead.
//
// Manages a network ACL rule resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/NetworkAcl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := NetworkAcl.NewAclRule(ctx, "rule1", &NetworkAcl.AclRuleArgs{
//				Action:               pulumi.String("deny"),
//				DestinationIpAddress: pulumi.String("4.3.2.0/24"),
//				DestinationPort:      pulumi.String("555"),
//				Protocol:             pulumi.String("udp"),
//				SourceIpAddress:      pulumi.String("1.2.3.4"),
//				SourcePort:           pulumi.String("444"),
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
// network ACL rules can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:NetworkAcl/aclRule:AclRule rule_1 89a84b28-4cc2-4859-9885-c67e802a46a3
//
// ```
type AclRule struct {
	pulumi.CustomResourceState

	// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
	// *deny*.
	Action pulumi.StringOutput `pulumi:"action"`
	// Specifies the description for the network ACL rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the destination IP address to which the traffic is allowed.
	// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	DestinationIpAddress pulumi.StringPtrOutput `pulumi:"destinationIpAddress"`
	// Specifies the destination port number or port number range. The value ranges
	// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	DestinationPort pulumi.StringPtrOutput `pulumi:"destinationPort"`
	// Enabled status for the network ACL rule. Defaults to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the IP version, either 4 (default) or 6. This parameter is available after
	// the IPv6 function is enabled.
	IpVersion pulumi.IntPtrOutput `pulumi:"ipVersion"`
	// Specifies a unique name for the network ACL rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
	// *udp*, *icmp* and *any*.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to create the network ACL rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new network ACL rule resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the source IP address that the traffic is allowed from. The default
	// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	SourceIpAddress pulumi.StringPtrOutput `pulumi:"sourceIpAddress"`
	// Specifies the source port number or port number range. The value ranges from 1 to
	// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	SourcePort pulumi.StringPtrOutput `pulumi:"sourcePort"`
}

// NewAclRule registers a new resource with the given unique name, arguments, and options.
func NewAclRule(ctx *pulumi.Context,
	name string, args *AclRuleArgs, opts ...pulumi.ResourceOption) (*AclRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AclRule
	err := ctx.RegisterResource("huaweicloud:NetworkAcl/aclRule:AclRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclRule gets an existing AclRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclRuleState, opts ...pulumi.ResourceOption) (*AclRule, error) {
	var resource AclRule
	err := ctx.ReadResource("huaweicloud:NetworkAcl/aclRule:AclRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclRule resources.
type aclRuleState struct {
	// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
	// *deny*.
	Action *string `pulumi:"action"`
	// Specifies the description for the network ACL rule.
	Description *string `pulumi:"description"`
	// Specifies the destination IP address to which the traffic is allowed.
	// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// Specifies the destination port number or port number range. The value ranges
	// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the network ACL rule. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the IP version, either 4 (default) or 6. This parameter is available after
	// the IPv6 function is enabled.
	IpVersion *int `pulumi:"ipVersion"`
	// Specifies a unique name for the network ACL rule.
	Name *string `pulumi:"name"`
	// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
	// *udp*, *icmp* and *any*.
	Protocol *string `pulumi:"protocol"`
	// The region in which to create the network ACL rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new network ACL rule resource.
	Region *string `pulumi:"region"`
	// Specifies the source IP address that the traffic is allowed from. The default
	// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// Specifies the source port number or port number range. The value ranges from 1 to
	// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	SourcePort *string `pulumi:"sourcePort"`
}

type AclRuleState struct {
	// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
	// *deny*.
	Action pulumi.StringPtrInput
	// Specifies the description for the network ACL rule.
	Description pulumi.StringPtrInput
	// Specifies the destination IP address to which the traffic is allowed.
	// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	DestinationIpAddress pulumi.StringPtrInput
	// Specifies the destination port number or port number range. The value ranges
	// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the network ACL rule. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// Specifies the IP version, either 4 (default) or 6. This parameter is available after
	// the IPv6 function is enabled.
	IpVersion pulumi.IntPtrInput
	// Specifies a unique name for the network ACL rule.
	Name pulumi.StringPtrInput
	// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
	// *udp*, *icmp* and *any*.
	Protocol pulumi.StringPtrInput
	// The region in which to create the network ACL rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new network ACL rule resource.
	Region pulumi.StringPtrInput
	// Specifies the source IP address that the traffic is allowed from. The default
	// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	SourceIpAddress pulumi.StringPtrInput
	// Specifies the source port number or port number range. The value ranges from 1 to
	// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	SourcePort pulumi.StringPtrInput
}

func (AclRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleState)(nil)).Elem()
}

type aclRuleArgs struct {
	// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
	// *deny*.
	Action string `pulumi:"action"`
	// Specifies the description for the network ACL rule.
	Description *string `pulumi:"description"`
	// Specifies the destination IP address to which the traffic is allowed.
	// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// Specifies the destination port number or port number range. The value ranges
	// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the network ACL rule. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the IP version, either 4 (default) or 6. This parameter is available after
	// the IPv6 function is enabled.
	IpVersion *int `pulumi:"ipVersion"`
	// Specifies a unique name for the network ACL rule.
	Name *string `pulumi:"name"`
	// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
	// *udp*, *icmp* and *any*.
	Protocol string `pulumi:"protocol"`
	// The region in which to create the network ACL rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new network ACL rule resource.
	Region *string `pulumi:"region"`
	// Specifies the source IP address that the traffic is allowed from. The default
	// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// Specifies the source port number or port number range. The value ranges from 1 to
	// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	SourcePort *string `pulumi:"sourcePort"`
}

// The set of arguments for constructing a AclRule resource.
type AclRuleArgs struct {
	// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
	// *deny*.
	Action pulumi.StringInput
	// Specifies the description for the network ACL rule.
	Description pulumi.StringPtrInput
	// Specifies the destination IP address to which the traffic is allowed.
	// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	DestinationIpAddress pulumi.StringPtrInput
	// Specifies the destination port number or port number range. The value ranges
	// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the network ACL rule. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// Specifies the IP version, either 4 (default) or 6. This parameter is available after
	// the IPv6 function is enabled.
	IpVersion pulumi.IntPtrInput
	// Specifies a unique name for the network ACL rule.
	Name pulumi.StringPtrInput
	// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
	// *udp*, *icmp* and *any*.
	Protocol pulumi.StringInput
	// The region in which to create the network ACL rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new network ACL rule resource.
	Region pulumi.StringPtrInput
	// Specifies the source IP address that the traffic is allowed from. The default
	// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
	SourceIpAddress pulumi.StringPtrInput
	// Specifies the source port number or port number range. The value ranges from 1 to
	// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
	SourcePort pulumi.StringPtrInput
}

func (AclRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleArgs)(nil)).Elem()
}

type AclRuleInput interface {
	pulumi.Input

	ToAclRuleOutput() AclRuleOutput
	ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput
}

func (*AclRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (i *AclRule) ToAclRuleOutput() AclRuleOutput {
	return i.ToAclRuleOutputWithContext(context.Background())
}

func (i *AclRule) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleOutput)
}

// AclRuleArrayInput is an input type that accepts AclRuleArray and AclRuleArrayOutput values.
// You can construct a concrete instance of `AclRuleArrayInput` via:
//
//	AclRuleArray{ AclRuleArgs{...} }
type AclRuleArrayInput interface {
	pulumi.Input

	ToAclRuleArrayOutput() AclRuleArrayOutput
	ToAclRuleArrayOutputWithContext(context.Context) AclRuleArrayOutput
}

type AclRuleArray []AclRuleInput

func (AclRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (i AclRuleArray) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return i.ToAclRuleArrayOutputWithContext(context.Background())
}

func (i AclRuleArray) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleArrayOutput)
}

// AclRuleMapInput is an input type that accepts AclRuleMap and AclRuleMapOutput values.
// You can construct a concrete instance of `AclRuleMapInput` via:
//
//	AclRuleMap{ "key": AclRuleArgs{...} }
type AclRuleMapInput interface {
	pulumi.Input

	ToAclRuleMapOutput() AclRuleMapOutput
	ToAclRuleMapOutputWithContext(context.Context) AclRuleMapOutput
}

type AclRuleMap map[string]AclRuleInput

func (AclRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (i AclRuleMap) ToAclRuleMapOutput() AclRuleMapOutput {
	return i.ToAclRuleMapOutputWithContext(context.Background())
}

func (i AclRuleMap) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleMapOutput)
}

type AclRuleOutput struct{ *pulumi.OutputState }

func (AclRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (o AclRuleOutput) ToAclRuleOutput() AclRuleOutput {
	return o
}

func (o AclRuleOutput) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return o
}

// Specifies the action in the network ACL rule. Currently, the value can be *allow* or
// *deny*.
func (o AclRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Specifies the description for the network ACL rule.
func (o AclRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the destination IP address to which the traffic is allowed.
// The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
func (o AclRuleOutput) DestinationIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.DestinationIpAddress }).(pulumi.StringPtrOutput)
}

// Specifies the destination port number or port number range. The value ranges
// from `1` to `65,535`. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
func (o AclRuleOutput) DestinationPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.DestinationPort }).(pulumi.StringPtrOutput)
}

// Enabled status for the network ACL rule. Defaults to true.
func (o AclRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies the IP version, either 4 (default) or 6. This parameter is available after
// the IPv6 function is enabled.
func (o AclRuleOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.IntPtrOutput { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// Specifies a unique name for the network ACL rule.
func (o AclRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the protocol supported by the network ACL rule. Valid values are: *tcp*,
// *udp*, *icmp* and *any*.
func (o AclRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The region in which to create the network ACL rule resource. If omitted, the
// provider-level region will be used. Changing this creates a new network ACL rule resource.
func (o AclRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the source IP address that the traffic is allowed from. The default
// value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).
func (o AclRuleOutput) SourceIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.SourceIpAddress }).(pulumi.StringPtrOutput)
}

// Specifies the source port number or port number range. The value ranges from 1 to
// 65535. For a port number range, enter two port numbers connected by a colon(:). For example, 1:100.
func (o AclRuleOutput) SourcePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.SourcePort }).(pulumi.StringPtrOutput)
}

type AclRuleArrayOutput struct{ *pulumi.OutputState }

func (AclRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) Index(i pulumi.IntInput) AclRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].([]*AclRule)[vs[1].(int)]
	}).(AclRuleOutput)
}

type AclRuleMapOutput struct{ *pulumi.OutputState }

func (AclRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (o AclRuleMapOutput) ToAclRuleMapOutput() AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) MapIndex(k pulumi.StringInput) AclRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].(map[string]*AclRule)[vs[1].(string)]
	}).(AclRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleInput)(nil)).Elem(), &AclRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleArrayInput)(nil)).Elem(), AclRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleMapInput)(nil)).Elem(), AclRuleMap{})
	pulumi.RegisterOutputType(AclRuleOutput{})
	pulumi.RegisterOutputType(AclRuleArrayOutput{})
	pulumi.RegisterOutputType(AclRuleMapOutput{})
}
