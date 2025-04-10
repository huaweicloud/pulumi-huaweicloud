// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IoTDA device linkage rule within HuaweiCloud.
//
// > When accessing an IoTDA **standard** or **enterprise** edition instance, you need to specify the IoTDA service
// endpoint in `provider` block.
// You can login to the IoTDA console, choose the instance **Overview** and click **Access Details**
// to view the HTTPS application access address. An example of the access address might be
// **9bc34xxxxx.st1.iotda-app.ap-southeast-1.myhuaweicloud.com**, then you need to configure the
// `provider` block as follows:
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/IoTDA"
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Smn"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/IoTDA"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			spaceId := cfg.RequireObject("spaceId")
//			triggerDeviceId := cfg.RequireObject("triggerDeviceId")
//			actionDeviceId := cfg.RequireObject("actionDeviceId")
//			topic, err := Smn.NewTopic(ctx, "topic", nil)
//			if err != nil {
//				return err
//			}
//			_, err = IoTDA.NewDeviceLinkageRule(ctx, "test", &IoTDA.DeviceLinkageRuleArgs{
//				SpaceId: pulumi.Any(spaceId),
//				Triggers: iotda.DeviceLinkageRuleTriggerArray{
//					&iotda.DeviceLinkageRuleTriggerArgs{
//						Type: pulumi.String("SIMPLE_TIMER"),
//						SimpleTimerCondition: &iotda.DeviceLinkageRuleTriggerSimpleTimerConditionArgs{
//							StartTime:      pulumi.String("20220622T160000Z"),
//							RepeatInterval: pulumi.Int(2),
//							RepeatCount:    pulumi.Int(2),
//						},
//					},
//					&iotda.DeviceLinkageRuleTriggerArgs{
//						Type: pulumi.String("DEVICE_DATA"),
//						DeviceDataCondition: &iotda.DeviceLinkageRuleTriggerDeviceDataConditionArgs{
//							DeviceId:            pulumi.Any(triggerDeviceId),
//							Path:                pulumi.String("service_id/propertyName_1"),
//							Operator:            pulumi.String("="),
//							Value:               pulumi.String("5"),
//							TriggerStrategy:     pulumi.String("pulse"),
//							DataValidatiyPeriod: pulumi.Int(300),
//						},
//					},
//					&iotda.DeviceLinkageRuleTriggerArgs{
//						Type: pulumi.String("DAILY_TIMER"),
//						DailyTimerCondition: &iotda.DeviceLinkageRuleTriggerDailyTimerConditionArgs{
//							StartTime: pulumi.String("19:02"),
//						},
//					},
//				},
//				Actions: iotda.DeviceLinkageRuleActionArray{
//					&iotda.DeviceLinkageRuleActionArgs{
//						Type: pulumi.String("SMN_FORWARDING"),
//						SmnForwarding: &iotda.DeviceLinkageRuleActionSmnForwardingArgs{
//							Region:         topic.Region,
//							TopicName:      topic.Name,
//							TopicUrn:       topic.TopicUrn,
//							MessageTitle:   pulumi.String("message_title"),
//							MessageContent: pulumi.String("message_content"),
//						},
//					},
//					&iotda.DeviceLinkageRuleActionArgs{
//						Type: pulumi.String("DEVICE_CMD"),
//						DeviceCommand: &iotda.DeviceLinkageRuleActionDeviceCommandArgs{
//							DeviceId:    pulumi.Any(actionDeviceId),
//							ServiceId:   pulumi.String("service_id"),
//							CommandName: pulumi.String("cmd_name"),
//							CommandBody: pulumi.String("{\"cmd_parameter_1\":\"3\"}"),
//						},
//					},
//				},
//				EffectivePeriod: &iotda.DeviceLinkageRuleEffectivePeriodArgs{
//					StartTime:  pulumi.String("00:00"),
//					EndTime:    pulumi.String("23:59"),
//					DaysOfWeek: pulumi.String("1,2,3"),
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
// Device linkage rules can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:IoTDA/deviceLinkageRule:DeviceLinkageRule test 62b6cc5aa367f403fea86127
//
// ```
type DeviceLinkageRule struct {
	pulumi.CustomResourceState

	// Specifies the list of the actions, at most 10 actions.
	// The actions structure is documented below.
	Actions DeviceLinkageRuleActionArrayOutput `pulumi:"actions"`
	// Specifies the description of the alarm.\
	// The value can contain a maximum of `256` characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the effective period of the device linkage rule. Always effectives
	// by default. The effectivePeriod structure is documented below.
	EffectivePeriod DeviceLinkageRuleEffectivePeriodPtrOutput `pulumi:"effectivePeriod"`
	// Specifies whether to enable the device linkage rule. Defaults to **true**.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the alarm.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region to which the SMN belongs.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the resource space ID to which the device linkage rule belongs.
	// Changing this parameter will create a new resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// Specifies the logical relationship between multiple triggers.
	// The options are as follows:
	// + **and**: All of the triggers are met.
	// + **or**: Any of the triggers are met.
	TriggerLogic pulumi.StringPtrOutput `pulumi:"triggerLogic"`
	// Specifies the list of the triggers, at most 10 triggers.
	// The triggers structure is documented below.
	Triggers DeviceLinkageRuleTriggerArrayOutput `pulumi:"triggers"`
}

// NewDeviceLinkageRule registers a new resource with the given unique name, arguments, and options.
func NewDeviceLinkageRule(ctx *pulumi.Context,
	name string, args *DeviceLinkageRuleArgs, opts ...pulumi.ResourceOption) (*DeviceLinkageRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.SpaceId == nil {
		return nil, errors.New("invalid value for required argument 'SpaceId'")
	}
	if args.Triggers == nil {
		return nil, errors.New("invalid value for required argument 'Triggers'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DeviceLinkageRule
	err := ctx.RegisterResource("huaweicloud:IoTDA/deviceLinkageRule:DeviceLinkageRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceLinkageRule gets an existing DeviceLinkageRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceLinkageRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceLinkageRuleState, opts ...pulumi.ResourceOption) (*DeviceLinkageRule, error) {
	var resource DeviceLinkageRule
	err := ctx.ReadResource("huaweicloud:IoTDA/deviceLinkageRule:DeviceLinkageRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceLinkageRule resources.
type deviceLinkageRuleState struct {
	// Specifies the list of the actions, at most 10 actions.
	// The actions structure is documented below.
	Actions []DeviceLinkageRuleAction `pulumi:"actions"`
	// Specifies the description of the alarm.\
	// The value can contain a maximum of `256` characters.
	Description *string `pulumi:"description"`
	// Specifies the effective period of the device linkage rule. Always effectives
	// by default. The effectivePeriod structure is documented below.
	EffectivePeriod *DeviceLinkageRuleEffectivePeriod `pulumi:"effectivePeriod"`
	// Specifies whether to enable the device linkage rule. Defaults to **true**.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the alarm.
	Name *string `pulumi:"name"`
	// Specifies the region to which the SMN belongs.
	Region *string `pulumi:"region"`
	// Specifies the resource space ID to which the device linkage rule belongs.
	// Changing this parameter will create a new resource.
	SpaceId *string `pulumi:"spaceId"`
	// Specifies the logical relationship between multiple triggers.
	// The options are as follows:
	// + **and**: All of the triggers are met.
	// + **or**: Any of the triggers are met.
	TriggerLogic *string `pulumi:"triggerLogic"`
	// Specifies the list of the triggers, at most 10 triggers.
	// The triggers structure is documented below.
	Triggers []DeviceLinkageRuleTrigger `pulumi:"triggers"`
}

type DeviceLinkageRuleState struct {
	// Specifies the list of the actions, at most 10 actions.
	// The actions structure is documented below.
	Actions DeviceLinkageRuleActionArrayInput
	// Specifies the description of the alarm.\
	// The value can contain a maximum of `256` characters.
	Description pulumi.StringPtrInput
	// Specifies the effective period of the device linkage rule. Always effectives
	// by default. The effectivePeriod structure is documented below.
	EffectivePeriod DeviceLinkageRuleEffectivePeriodPtrInput
	// Specifies whether to enable the device linkage rule. Defaults to **true**.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the alarm.
	Name pulumi.StringPtrInput
	// Specifies the region to which the SMN belongs.
	Region pulumi.StringPtrInput
	// Specifies the resource space ID to which the device linkage rule belongs.
	// Changing this parameter will create a new resource.
	SpaceId pulumi.StringPtrInput
	// Specifies the logical relationship between multiple triggers.
	// The options are as follows:
	// + **and**: All of the triggers are met.
	// + **or**: Any of the triggers are met.
	TriggerLogic pulumi.StringPtrInput
	// Specifies the list of the triggers, at most 10 triggers.
	// The triggers structure is documented below.
	Triggers DeviceLinkageRuleTriggerArrayInput
}

func (DeviceLinkageRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceLinkageRuleState)(nil)).Elem()
}

type deviceLinkageRuleArgs struct {
	// Specifies the list of the actions, at most 10 actions.
	// The actions structure is documented below.
	Actions []DeviceLinkageRuleAction `pulumi:"actions"`
	// Specifies the description of the alarm.\
	// The value can contain a maximum of `256` characters.
	Description *string `pulumi:"description"`
	// Specifies the effective period of the device linkage rule. Always effectives
	// by default. The effectivePeriod structure is documented below.
	EffectivePeriod *DeviceLinkageRuleEffectivePeriod `pulumi:"effectivePeriod"`
	// Specifies whether to enable the device linkage rule. Defaults to **true**.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the alarm.
	Name *string `pulumi:"name"`
	// Specifies the region to which the SMN belongs.
	Region *string `pulumi:"region"`
	// Specifies the resource space ID to which the device linkage rule belongs.
	// Changing this parameter will create a new resource.
	SpaceId string `pulumi:"spaceId"`
	// Specifies the logical relationship between multiple triggers.
	// The options are as follows:
	// + **and**: All of the triggers are met.
	// + **or**: Any of the triggers are met.
	TriggerLogic *string `pulumi:"triggerLogic"`
	// Specifies the list of the triggers, at most 10 triggers.
	// The triggers structure is documented below.
	Triggers []DeviceLinkageRuleTrigger `pulumi:"triggers"`
}

// The set of arguments for constructing a DeviceLinkageRule resource.
type DeviceLinkageRuleArgs struct {
	// Specifies the list of the actions, at most 10 actions.
	// The actions structure is documented below.
	Actions DeviceLinkageRuleActionArrayInput
	// Specifies the description of the alarm.\
	// The value can contain a maximum of `256` characters.
	Description pulumi.StringPtrInput
	// Specifies the effective period of the device linkage rule. Always effectives
	// by default. The effectivePeriod structure is documented below.
	EffectivePeriod DeviceLinkageRuleEffectivePeriodPtrInput
	// Specifies whether to enable the device linkage rule. Defaults to **true**.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the alarm.
	Name pulumi.StringPtrInput
	// Specifies the region to which the SMN belongs.
	Region pulumi.StringPtrInput
	// Specifies the resource space ID to which the device linkage rule belongs.
	// Changing this parameter will create a new resource.
	SpaceId pulumi.StringInput
	// Specifies the logical relationship between multiple triggers.
	// The options are as follows:
	// + **and**: All of the triggers are met.
	// + **or**: Any of the triggers are met.
	TriggerLogic pulumi.StringPtrInput
	// Specifies the list of the triggers, at most 10 triggers.
	// The triggers structure is documented below.
	Triggers DeviceLinkageRuleTriggerArrayInput
}

func (DeviceLinkageRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceLinkageRuleArgs)(nil)).Elem()
}

type DeviceLinkageRuleInput interface {
	pulumi.Input

	ToDeviceLinkageRuleOutput() DeviceLinkageRuleOutput
	ToDeviceLinkageRuleOutputWithContext(ctx context.Context) DeviceLinkageRuleOutput
}

func (*DeviceLinkageRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceLinkageRule)(nil)).Elem()
}

func (i *DeviceLinkageRule) ToDeviceLinkageRuleOutput() DeviceLinkageRuleOutput {
	return i.ToDeviceLinkageRuleOutputWithContext(context.Background())
}

func (i *DeviceLinkageRule) ToDeviceLinkageRuleOutputWithContext(ctx context.Context) DeviceLinkageRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceLinkageRuleOutput)
}

// DeviceLinkageRuleArrayInput is an input type that accepts DeviceLinkageRuleArray and DeviceLinkageRuleArrayOutput values.
// You can construct a concrete instance of `DeviceLinkageRuleArrayInput` via:
//
//	DeviceLinkageRuleArray{ DeviceLinkageRuleArgs{...} }
type DeviceLinkageRuleArrayInput interface {
	pulumi.Input

	ToDeviceLinkageRuleArrayOutput() DeviceLinkageRuleArrayOutput
	ToDeviceLinkageRuleArrayOutputWithContext(context.Context) DeviceLinkageRuleArrayOutput
}

type DeviceLinkageRuleArray []DeviceLinkageRuleInput

func (DeviceLinkageRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceLinkageRule)(nil)).Elem()
}

func (i DeviceLinkageRuleArray) ToDeviceLinkageRuleArrayOutput() DeviceLinkageRuleArrayOutput {
	return i.ToDeviceLinkageRuleArrayOutputWithContext(context.Background())
}

func (i DeviceLinkageRuleArray) ToDeviceLinkageRuleArrayOutputWithContext(ctx context.Context) DeviceLinkageRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceLinkageRuleArrayOutput)
}

// DeviceLinkageRuleMapInput is an input type that accepts DeviceLinkageRuleMap and DeviceLinkageRuleMapOutput values.
// You can construct a concrete instance of `DeviceLinkageRuleMapInput` via:
//
//	DeviceLinkageRuleMap{ "key": DeviceLinkageRuleArgs{...} }
type DeviceLinkageRuleMapInput interface {
	pulumi.Input

	ToDeviceLinkageRuleMapOutput() DeviceLinkageRuleMapOutput
	ToDeviceLinkageRuleMapOutputWithContext(context.Context) DeviceLinkageRuleMapOutput
}

type DeviceLinkageRuleMap map[string]DeviceLinkageRuleInput

func (DeviceLinkageRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceLinkageRule)(nil)).Elem()
}

func (i DeviceLinkageRuleMap) ToDeviceLinkageRuleMapOutput() DeviceLinkageRuleMapOutput {
	return i.ToDeviceLinkageRuleMapOutputWithContext(context.Background())
}

func (i DeviceLinkageRuleMap) ToDeviceLinkageRuleMapOutputWithContext(ctx context.Context) DeviceLinkageRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceLinkageRuleMapOutput)
}

type DeviceLinkageRuleOutput struct{ *pulumi.OutputState }

func (DeviceLinkageRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceLinkageRule)(nil)).Elem()
}

func (o DeviceLinkageRuleOutput) ToDeviceLinkageRuleOutput() DeviceLinkageRuleOutput {
	return o
}

func (o DeviceLinkageRuleOutput) ToDeviceLinkageRuleOutputWithContext(ctx context.Context) DeviceLinkageRuleOutput {
	return o
}

// Specifies the list of the actions, at most 10 actions.
// The actions structure is documented below.
func (o DeviceLinkageRuleOutput) Actions() DeviceLinkageRuleActionArrayOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) DeviceLinkageRuleActionArrayOutput { return v.Actions }).(DeviceLinkageRuleActionArrayOutput)
}

// Specifies the description of the alarm.\
// The value can contain a maximum of `256` characters.
func (o DeviceLinkageRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the effective period of the device linkage rule. Always effectives
// by default. The effectivePeriod structure is documented below.
func (o DeviceLinkageRuleOutput) EffectivePeriod() DeviceLinkageRuleEffectivePeriodPtrOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) DeviceLinkageRuleEffectivePeriodPtrOutput { return v.EffectivePeriod }).(DeviceLinkageRuleEffectivePeriodPtrOutput)
}

// Specifies whether to enable the device linkage rule. Defaults to **true**.
func (o DeviceLinkageRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies the name of the alarm.
func (o DeviceLinkageRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region to which the SMN belongs.
func (o DeviceLinkageRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the resource space ID to which the device linkage rule belongs.
// Changing this parameter will create a new resource.
func (o DeviceLinkageRuleOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// Specifies the logical relationship between multiple triggers.
// The options are as follows:
// + **and**: All of the triggers are met.
// + **or**: Any of the triggers are met.
func (o DeviceLinkageRuleOutput) TriggerLogic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) pulumi.StringPtrOutput { return v.TriggerLogic }).(pulumi.StringPtrOutput)
}

// Specifies the list of the triggers, at most 10 triggers.
// The triggers structure is documented below.
func (o DeviceLinkageRuleOutput) Triggers() DeviceLinkageRuleTriggerArrayOutput {
	return o.ApplyT(func(v *DeviceLinkageRule) DeviceLinkageRuleTriggerArrayOutput { return v.Triggers }).(DeviceLinkageRuleTriggerArrayOutput)
}

type DeviceLinkageRuleArrayOutput struct{ *pulumi.OutputState }

func (DeviceLinkageRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceLinkageRule)(nil)).Elem()
}

func (o DeviceLinkageRuleArrayOutput) ToDeviceLinkageRuleArrayOutput() DeviceLinkageRuleArrayOutput {
	return o
}

func (o DeviceLinkageRuleArrayOutput) ToDeviceLinkageRuleArrayOutputWithContext(ctx context.Context) DeviceLinkageRuleArrayOutput {
	return o
}

func (o DeviceLinkageRuleArrayOutput) Index(i pulumi.IntInput) DeviceLinkageRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceLinkageRule {
		return vs[0].([]*DeviceLinkageRule)[vs[1].(int)]
	}).(DeviceLinkageRuleOutput)
}

type DeviceLinkageRuleMapOutput struct{ *pulumi.OutputState }

func (DeviceLinkageRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceLinkageRule)(nil)).Elem()
}

func (o DeviceLinkageRuleMapOutput) ToDeviceLinkageRuleMapOutput() DeviceLinkageRuleMapOutput {
	return o
}

func (o DeviceLinkageRuleMapOutput) ToDeviceLinkageRuleMapOutputWithContext(ctx context.Context) DeviceLinkageRuleMapOutput {
	return o
}

func (o DeviceLinkageRuleMapOutput) MapIndex(k pulumi.StringInput) DeviceLinkageRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceLinkageRule {
		return vs[0].(map[string]*DeviceLinkageRule)[vs[1].(string)]
	}).(DeviceLinkageRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceLinkageRuleInput)(nil)).Elem(), &DeviceLinkageRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceLinkageRuleArrayInput)(nil)).Elem(), DeviceLinkageRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceLinkageRuleMapInput)(nil)).Elem(), DeviceLinkageRuleMap{})
	pulumi.RegisterOutputType(DeviceLinkageRuleOutput{})
	pulumi.RegisterOutputType(DeviceLinkageRuleArrayOutput{})
	pulumi.RegisterOutputType(DeviceLinkageRuleMapOutput{})
}
