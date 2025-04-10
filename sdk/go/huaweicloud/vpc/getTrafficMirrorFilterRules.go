// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the traffic mirror filter rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.GetTrafficMirrorFilterRules(ctx, &vpc.GetTrafficMirrorFilterRulesArgs{
//				Protocol: pulumi.StringRef("udp"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTrafficMirrorFilterRules(ctx *pulumi.Context, args *GetTrafficMirrorFilterRulesArgs, opts ...pulumi.InvokeOption) (*GetTrafficMirrorFilterRulesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTrafficMirrorFilterRulesResult
	err := ctx.Invoke("huaweicloud:Vpc/getTrafficMirrorFilterRules:getTrafficMirrorFilterRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrafficMirrorFilterRules.
type GetTrafficMirrorFilterRulesArgs struct {
	// The policy of in the traffic mirror filter rule.
	// Valid values are **accept** or **reject**.
	Action *string `pulumi:"action"`
	// The destination IP address of the traffic mirror filter rule.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination port number range of the traffic mirror filter rule.
	// The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// The direction of the traffic mirror filter rule.
	// Valid values are **ingress** or **egress**.
	Direction *string `pulumi:"direction"`
	// The priority number of the traffic mirror filter rule.
	// Valid value ranges from `1` to `65,535`.
	Priority *string `pulumi:"priority"`
	// The protocol of the traffic mirror filter rule.
	// Valid value are **tcp**, **udp**, **icmp**, **icmpv6**, **all**.
	Protocol *string `pulumi:"protocol"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// The source IP address of the traffic mirror filter rule.
	SourceCidrBlock *string `pulumi:"sourceCidrBlock"`
	// The source port number range of the traffic mirror filter rule.
	// The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The traffic mirror filter ID used as the query filter.
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
	// The traffic mirror filter rule ID used as the query filter.
	TrafficMirrorFilterRuleId *string `pulumi:"trafficMirrorFilterRuleId"`
}

// A collection of values returned by getTrafficMirrorFilterRules.
type GetTrafficMirrorFilterRulesResult struct {
	// Whether to accept or reject traffic.
	Action *string `pulumi:"action"`
	// Destination CIDR block of the mirrored traffic.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Source port range.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Traffic direction.
	Direction *string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Mirror filter rule priority.
	Priority *string `pulumi:"priority"`
	// Protocol of the mirrored traffic.
	Protocol *string `pulumi:"protocol"`
	Region   string  `pulumi:"region"`
	// Source CIDR block of the mirrored traffic.
	SourceCidrBlock *string `pulumi:"sourceCidrBlock"`
	// Source port range.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// Traffic mirror filter ID.
	TrafficMirrorFilterId     *string `pulumi:"trafficMirrorFilterId"`
	TrafficMirrorFilterRuleId *string `pulumi:"trafficMirrorFilterRuleId"`
	// List of traffic mirror filter rules.
	TrafficMirrorFilterRules []GetTrafficMirrorFilterRulesTrafficMirrorFilterRule `pulumi:"trafficMirrorFilterRules"`
}

func GetTrafficMirrorFilterRulesOutput(ctx *pulumi.Context, args GetTrafficMirrorFilterRulesOutputArgs, opts ...pulumi.InvokeOption) GetTrafficMirrorFilterRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrafficMirrorFilterRulesResult, error) {
			args := v.(GetTrafficMirrorFilterRulesArgs)
			r, err := GetTrafficMirrorFilterRules(ctx, &args, opts...)
			var s GetTrafficMirrorFilterRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTrafficMirrorFilterRulesResultOutput)
}

// A collection of arguments for invoking getTrafficMirrorFilterRules.
type GetTrafficMirrorFilterRulesOutputArgs struct {
	// The policy of in the traffic mirror filter rule.
	// Valid values are **accept** or **reject**.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// The destination IP address of the traffic mirror filter rule.
	DestinationCidrBlock pulumi.StringPtrInput `pulumi:"destinationCidrBlock"`
	// The destination port number range of the traffic mirror filter rule.
	// The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
	DestinationPortRange pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	// The direction of the traffic mirror filter rule.
	// Valid values are **ingress** or **egress**.
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	// The priority number of the traffic mirror filter rule.
	// Valid value ranges from `1` to `65,535`.
	Priority pulumi.StringPtrInput `pulumi:"priority"`
	// The protocol of the traffic mirror filter rule.
	// Valid value are **tcp**, **udp**, **icmp**, **icmpv6**, **all**.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The source IP address of the traffic mirror filter rule.
	SourceCidrBlock pulumi.StringPtrInput `pulumi:"sourceCidrBlock"`
	// The source port number range of the traffic mirror filter rule.
	// The value ranges from `1` to `65,535`, enter two port numbers connected by a hyphen (-). For example, **80-200**.
	SourcePortRange pulumi.StringPtrInput `pulumi:"sourcePortRange"`
	// The traffic mirror filter ID used as the query filter.
	TrafficMirrorFilterId pulumi.StringPtrInput `pulumi:"trafficMirrorFilterId"`
	// The traffic mirror filter rule ID used as the query filter.
	TrafficMirrorFilterRuleId pulumi.StringPtrInput `pulumi:"trafficMirrorFilterRuleId"`
}

func (GetTrafficMirrorFilterRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficMirrorFilterRulesArgs)(nil)).Elem()
}

// A collection of values returned by getTrafficMirrorFilterRules.
type GetTrafficMirrorFilterRulesResultOutput struct{ *pulumi.OutputState }

func (GetTrafficMirrorFilterRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficMirrorFilterRulesResult)(nil)).Elem()
}

func (o GetTrafficMirrorFilterRulesResultOutput) ToGetTrafficMirrorFilterRulesResultOutput() GetTrafficMirrorFilterRulesResultOutput {
	return o
}

func (o GetTrafficMirrorFilterRulesResultOutput) ToGetTrafficMirrorFilterRulesResultOutputWithContext(ctx context.Context) GetTrafficMirrorFilterRulesResultOutput {
	return o
}

// Whether to accept or reject traffic.
func (o GetTrafficMirrorFilterRulesResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// Destination CIDR block of the mirrored traffic.
func (o GetTrafficMirrorFilterRulesResultOutput) DestinationCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.DestinationCidrBlock }).(pulumi.StringPtrOutput)
}

// Source port range.
func (o GetTrafficMirrorFilterRulesResultOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

// Traffic direction.
func (o GetTrafficMirrorFilterRulesResultOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTrafficMirrorFilterRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Mirror filter rule priority.
func (o GetTrafficMirrorFilterRulesResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// Protocol of the mirrored traffic.
func (o GetTrafficMirrorFilterRulesResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o GetTrafficMirrorFilterRulesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) string { return v.Region }).(pulumi.StringOutput)
}

// Source CIDR block of the mirrored traffic.
func (o GetTrafficMirrorFilterRulesResultOutput) SourceCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.SourceCidrBlock }).(pulumi.StringPtrOutput)
}

// Source port range.
func (o GetTrafficMirrorFilterRulesResultOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

// Traffic mirror filter ID.
func (o GetTrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.TrafficMirrorFilterId }).(pulumi.StringPtrOutput)
}

func (o GetTrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) *string { return v.TrafficMirrorFilterRuleId }).(pulumi.StringPtrOutput)
}

// List of traffic mirror filter rules.
func (o GetTrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterRules() GetTrafficMirrorFilterRulesTrafficMirrorFilterRuleArrayOutput {
	return o.ApplyT(func(v GetTrafficMirrorFilterRulesResult) []GetTrafficMirrorFilterRulesTrafficMirrorFilterRule {
		return v.TrafficMirrorFilterRules
	}).(GetTrafficMirrorFilterRulesTrafficMirrorFilterRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrafficMirrorFilterRulesResultOutput{})
}
