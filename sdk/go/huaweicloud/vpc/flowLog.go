// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VPC flow log resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Lts"
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testGroup, err := Lts.NewGroup(ctx, "testGroup", &Lts.GroupArgs{
//				GroupName: pulumi.String("test_group"),
//				TtlInDays: pulumi.Int(7),
//			})
//			if err != nil {
//				return err
//			}
//			testStream, err := Lts.NewStream(ctx, "testStream", &Lts.StreamArgs{
//				GroupId:    testGroup.ID(),
//				StreamName: pulumi.String("test_stream"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.NewFlowLog(ctx, "testFlowlog", &Vpc.FlowLogArgs{
//				ResourceType: pulumi.String("network"),
//				ResourceId:   pulumi.Any(_var.Subnet_id),
//				TrafficType:  pulumi.String("all"),
//				LogGroupId:   testGroup.ID(),
//				LogStreamId:  testStream.ID(),
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
// VPC flow logs can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpc/flowLog:FlowLog flowlog1 41b9d73f-eb1c-4795-a100-59a99b062513
//
// ```
type FlowLog struct {
	pulumi.CustomResourceState

	// Specifies description about the VPC flow log.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether to enable the flow log function, the default value is *true*.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the LTS log group ID.
	// Changing this creates a new VPC flow log.
	LogGroupId pulumi.StringOutput `pulumi:"logGroupId"`
	// Specifies the LTS log stream ID.
	// Changing this creates a new VPC flow log.
	LogStreamId pulumi.StringOutput `pulumi:"logStreamId"`
	// Specifies the VPC flow log name. The value can contain no more than 64 characters,
	// including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the resource ID for which that the logs to be collected.
	// Changing this creates a new VPC flow log.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Specifies the resource type for which that the logs to be collected.
	// The value can be:
	// + *port*: Select this to record traffic information about a specific NIC.
	// + *network*: Select this to record traffic information about all NICs in a specific subnet.
	// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The status of the flow log. The value can be `ACTIVE`, `DOWN` or `ERROR`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the type of traffic to log. The value can be:
	// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
	// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
	// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
	TrafficType pulumi.StringPtrOutput `pulumi:"trafficType"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogGroupId == nil {
		return nil, errors.New("invalid value for required argument 'LogGroupId'")
	}
	if args.LogStreamId == nil {
		return nil, errors.New("invalid value for required argument 'LogStreamId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FlowLog
	err := ctx.RegisterResource("huaweicloud:Vpc/flowLog:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("huaweicloud:Vpc/flowLog:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// Specifies description about the VPC flow log.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description *string `pulumi:"description"`
	// Specifies whether to enable the flow log function, the default value is *true*.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the LTS log group ID.
	// Changing this creates a new VPC flow log.
	LogGroupId *string `pulumi:"logGroupId"`
	// Specifies the LTS log stream ID.
	// Changing this creates a new VPC flow log.
	LogStreamId *string `pulumi:"logStreamId"`
	// Specifies the VPC flow log name. The value can contain no more than 64 characters,
	// including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
	Region *string `pulumi:"region"`
	// Specifies the resource ID for which that the logs to be collected.
	// Changing this creates a new VPC flow log.
	ResourceId *string `pulumi:"resourceId"`
	// Specifies the resource type for which that the logs to be collected.
	// The value can be:
	// + *port*: Select this to record traffic information about a specific NIC.
	// + *network*: Select this to record traffic information about all NICs in a specific subnet.
	// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
	ResourceType *string `pulumi:"resourceType"`
	// The status of the flow log. The value can be `ACTIVE`, `DOWN` or `ERROR`.
	Status *string `pulumi:"status"`
	// Specifies the type of traffic to log. The value can be:
	// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
	// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
	// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
	TrafficType *string `pulumi:"trafficType"`
}

type FlowLogState struct {
	// Specifies description about the VPC flow log.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description pulumi.StringPtrInput
	// Specifies whether to enable the flow log function, the default value is *true*.
	Enabled pulumi.BoolPtrInput
	// Specifies the LTS log group ID.
	// Changing this creates a new VPC flow log.
	LogGroupId pulumi.StringPtrInput
	// Specifies the LTS log stream ID.
	// Changing this creates a new VPC flow log.
	LogStreamId pulumi.StringPtrInput
	// Specifies the VPC flow log name. The value can contain no more than 64 characters,
	// including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
	Region pulumi.StringPtrInput
	// Specifies the resource ID for which that the logs to be collected.
	// Changing this creates a new VPC flow log.
	ResourceId pulumi.StringPtrInput
	// Specifies the resource type for which that the logs to be collected.
	// The value can be:
	// + *port*: Select this to record traffic information about a specific NIC.
	// + *network*: Select this to record traffic information about all NICs in a specific subnet.
	// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
	ResourceType pulumi.StringPtrInput
	// The status of the flow log. The value can be `ACTIVE`, `DOWN` or `ERROR`.
	Status pulumi.StringPtrInput
	// Specifies the type of traffic to log. The value can be:
	// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
	// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
	// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
	TrafficType pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// Specifies description about the VPC flow log.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description *string `pulumi:"description"`
	// Specifies whether to enable the flow log function, the default value is *true*.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the LTS log group ID.
	// Changing this creates a new VPC flow log.
	LogGroupId string `pulumi:"logGroupId"`
	// Specifies the LTS log stream ID.
	// Changing this creates a new VPC flow log.
	LogStreamId string `pulumi:"logStreamId"`
	// Specifies the VPC flow log name. The value can contain no more than 64 characters,
	// including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
	Region *string `pulumi:"region"`
	// Specifies the resource ID for which that the logs to be collected.
	// Changing this creates a new VPC flow log.
	ResourceId string `pulumi:"resourceId"`
	// Specifies the resource type for which that the logs to be collected.
	// The value can be:
	// + *port*: Select this to record traffic information about a specific NIC.
	// + *network*: Select this to record traffic information about all NICs in a specific subnet.
	// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
	ResourceType string `pulumi:"resourceType"`
	// Specifies the type of traffic to log. The value can be:
	// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
	// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
	// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
	TrafficType *string `pulumi:"trafficType"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// Specifies description about the VPC flow log.
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description pulumi.StringPtrInput
	// Specifies whether to enable the flow log function, the default value is *true*.
	Enabled pulumi.BoolPtrInput
	// Specifies the LTS log group ID.
	// Changing this creates a new VPC flow log.
	LogGroupId pulumi.StringInput
	// Specifies the LTS log stream ID.
	// Changing this creates a new VPC flow log.
	LogStreamId pulumi.StringInput
	// Specifies the VPC flow log name. The value can contain no more than 64 characters,
	// including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
	Region pulumi.StringPtrInput
	// Specifies the resource ID for which that the logs to be collected.
	// Changing this creates a new VPC flow log.
	ResourceId pulumi.StringInput
	// Specifies the resource type for which that the logs to be collected.
	// The value can be:
	// + *port*: Select this to record traffic information about a specific NIC.
	// + *network*: Select this to record traffic information about all NICs in a specific subnet.
	// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
	ResourceType pulumi.StringInput
	// Specifies the type of traffic to log. The value can be:
	// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
	// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
	// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
	TrafficType pulumi.StringPtrInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (*FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

// FlowLogArrayInput is an input type that accepts FlowLogArray and FlowLogArrayOutput values.
// You can construct a concrete instance of `FlowLogArrayInput` via:
//
//	FlowLogArray{ FlowLogArgs{...} }
type FlowLogArrayInput interface {
	pulumi.Input

	ToFlowLogArrayOutput() FlowLogArrayOutput
	ToFlowLogArrayOutputWithContext(context.Context) FlowLogArrayOutput
}

type FlowLogArray []FlowLogInput

func (FlowLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (i FlowLogArray) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return i.ToFlowLogArrayOutputWithContext(context.Background())
}

func (i FlowLogArray) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogArrayOutput)
}

// FlowLogMapInput is an input type that accepts FlowLogMap and FlowLogMapOutput values.
// You can construct a concrete instance of `FlowLogMapInput` via:
//
//	FlowLogMap{ "key": FlowLogArgs{...} }
type FlowLogMapInput interface {
	pulumi.Input

	ToFlowLogMapOutput() FlowLogMapOutput
	ToFlowLogMapOutputWithContext(context.Context) FlowLogMapOutput
}

type FlowLogMap map[string]FlowLogInput

func (FlowLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (i FlowLogMap) ToFlowLogMapOutput() FlowLogMapOutput {
	return i.ToFlowLogMapOutputWithContext(context.Background())
}

func (i FlowLogMap) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogMapOutput)
}

type FlowLogOutput struct{ *pulumi.OutputState }

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

// Specifies description about the VPC flow log.
// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
func (o FlowLogOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies whether to enable the flow log function, the default value is *true*.
func (o FlowLogOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies the LTS log group ID.
// Changing this creates a new VPC flow log.
func (o FlowLogOutput) LogGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogGroupId }).(pulumi.StringOutput)
}

// Specifies the LTS log stream ID.
// Changing this creates a new VPC flow log.
func (o FlowLogOutput) LogStreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogStreamId }).(pulumi.StringOutput)
}

// Specifies the VPC flow log name. The value can contain no more than 64 characters,
// including letters, digits, underscores (_), hyphens (-), and periods (.).
func (o FlowLogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.
func (o FlowLogOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the resource ID for which that the logs to be collected.
// Changing this creates a new VPC flow log.
func (o FlowLogOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Specifies the resource type for which that the logs to be collected.
// The value can be:
// + *port*: Select this to record traffic information about a specific NIC.
// + *network*: Select this to record traffic information about all NICs in a specific subnet.
// + *vpc*: Select this to record traffic information about all NICs in a specific VPC.
func (o FlowLogOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The status of the flow log. The value can be `ACTIVE`, `DOWN` or `ERROR`.
func (o FlowLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the type of traffic to log. The value can be:
// + *all*: Specifies that both accepted and rejected traffic of the specified resource will be logged.
// + *accept*: Specifies that only accepted inbound and outbound traffic of the specified resource will be logged.
// + *reject*: Specifies that only rejected inbound and outbound traffic of the specified resource will be logged.
func (o FlowLogOutput) TrafficType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.TrafficType }).(pulumi.StringPtrOutput)
}

type FlowLogArrayOutput struct{ *pulumi.OutputState }

func (FlowLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) Index(i pulumi.IntInput) FlowLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].([]*FlowLog)[vs[1].(int)]
	}).(FlowLogOutput)
}

type FlowLogMapOutput struct{ *pulumi.OutputState }

func (FlowLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (o FlowLogMapOutput) ToFlowLogMapOutput() FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) MapIndex(k pulumi.StringInput) FlowLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].(map[string]*FlowLog)[vs[1].(string)]
	}).(FlowLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogArrayInput)(nil)).Elem(), FlowLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogMapInput)(nil)).Elem(), FlowLogMap{})
	pulumi.RegisterOutputType(FlowLogOutput{})
	pulumi.RegisterOutputType(FlowLogArrayOutput{})
	pulumi.RegisterOutputType(FlowLogMapOutput{})
}
