// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cce

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to scale the CCE node pool within HuaweiCloud.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.RequireObject("clusterId")
//			nodepoolId := cfg.RequireObject("nodepoolId")
//			_, err := Cce.NewNodePoolScale(ctx, "test", &Cce.NodePoolScaleArgs{
//				ClusterId:  pulumi.Any(clusterId),
//				NodepoolId: pulumi.Any(nodepoolId),
//				ScaleGroups: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//				DesiredNodeCount: pulumi.Int(2),
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
// > Deleting node pool scale is not supported, it will only be removed from the state.
type NodePoolScale struct {
	pulumi.CustomResourceState

	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrOutput `pulumi:"autoRenew"`
	// Specifies the charging mode of the nodes.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new cluster resource.
	ChargingMode pulumi.StringOutput `pulumi:"chargingMode"`
	// Specifies the cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Specifies the number of desired nodes.
	DesiredNodeCount pulumi.IntOutput       `pulumi:"desiredNodeCount"`
	EnableForceNew   pulumi.StringPtrOutput `pulumi:"enableForceNew"`
	// Specifies the node pool ID.
	NodepoolId pulumi.StringOutput `pulumi:"nodepoolId"`
	// Specifies the charging period of the nodes.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Specifies the charging period unit of the nodes.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// Specifies the region in which to create the node pool scale resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the scalable checking.
	// The value can be **instant** and **async**, defaults to **instant**.
	ScalableChecking pulumi.StringPtrOutput `pulumi:"scalableChecking"`
	// Specifies the IDs of scale groups to scale.
	// **default** indicates the default group.
	ScaleGroups pulumi.StringArrayOutput `pulumi:"scaleGroups"`
}

// NewNodePoolScale registers a new resource with the given unique name, arguments, and options.
func NewNodePoolScale(ctx *pulumi.Context,
	name string, args *NodePoolScaleArgs, opts ...pulumi.ResourceOption) (*NodePoolScale, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.DesiredNodeCount == nil {
		return nil, errors.New("invalid value for required argument 'DesiredNodeCount'")
	}
	if args.NodepoolId == nil {
		return nil, errors.New("invalid value for required argument 'NodepoolId'")
	}
	if args.ScaleGroups == nil {
		return nil, errors.New("invalid value for required argument 'ScaleGroups'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NodePoolScale
	err := ctx.RegisterResource("huaweicloud:Cce/nodePoolScale:NodePoolScale", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodePoolScale gets an existing NodePoolScale resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePoolScale(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodePoolScaleState, opts ...pulumi.ResourceOption) (*NodePoolScale, error) {
	var resource NodePoolScale
	err := ctx.ReadResource("huaweicloud:Cce/nodePoolScale:NodePoolScale", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodePoolScale resources.
type nodePoolScaleState struct {
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the charging mode of the nodes.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new cluster resource.
	ChargingMode *string `pulumi:"chargingMode"`
	// Specifies the cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Specifies the number of desired nodes.
	DesiredNodeCount *int    `pulumi:"desiredNodeCount"`
	EnableForceNew   *string `pulumi:"enableForceNew"`
	// Specifies the node pool ID.
	NodepoolId *string `pulumi:"nodepoolId"`
	// Specifies the charging period of the nodes.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the nodes.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	PeriodUnit *string `pulumi:"periodUnit"`
	// Specifies the region in which to create the node pool scale resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the scalable checking.
	// The value can be **instant** and **async**, defaults to **instant**.
	ScalableChecking *string `pulumi:"scalableChecking"`
	// Specifies the IDs of scale groups to scale.
	// **default** indicates the default group.
	ScaleGroups []string `pulumi:"scaleGroups"`
}

type NodePoolScaleState struct {
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrInput
	// Specifies the charging mode of the nodes.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new cluster resource.
	ChargingMode pulumi.StringPtrInput
	// Specifies the cluster ID.
	ClusterId pulumi.StringPtrInput
	// Specifies the number of desired nodes.
	DesiredNodeCount pulumi.IntPtrInput
	EnableForceNew   pulumi.StringPtrInput
	// Specifies the node pool ID.
	NodepoolId pulumi.StringPtrInput
	// Specifies the charging period of the nodes.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the nodes.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	PeriodUnit pulumi.StringPtrInput
	// Specifies the region in which to create the node pool scale resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the scalable checking.
	// The value can be **instant** and **async**, defaults to **instant**.
	ScalableChecking pulumi.StringPtrInput
	// Specifies the IDs of scale groups to scale.
	// **default** indicates the default group.
	ScaleGroups pulumi.StringArrayInput
}

func (NodePoolScaleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolScaleState)(nil)).Elem()
}

type nodePoolScaleArgs struct {
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the charging mode of the nodes.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new cluster resource.
	ChargingMode *string `pulumi:"chargingMode"`
	// Specifies the cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the number of desired nodes.
	DesiredNodeCount int     `pulumi:"desiredNodeCount"`
	EnableForceNew   *string `pulumi:"enableForceNew"`
	// Specifies the node pool ID.
	NodepoolId string `pulumi:"nodepoolId"`
	// Specifies the charging period of the nodes.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the nodes.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	PeriodUnit *string `pulumi:"periodUnit"`
	// Specifies the region in which to create the node pool scale resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the scalable checking.
	// The value can be **instant** and **async**, defaults to **instant**.
	ScalableChecking *string `pulumi:"scalableChecking"`
	// Specifies the IDs of scale groups to scale.
	// **default** indicates the default group.
	ScaleGroups []string `pulumi:"scaleGroups"`
}

// The set of arguments for constructing a NodePoolScale resource.
type NodePoolScaleArgs struct {
	// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
	AutoRenew pulumi.StringPtrInput
	// Specifies the charging mode of the nodes.
	// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
	// Changing this parameter will create a new cluster resource.
	ChargingMode pulumi.StringPtrInput
	// Specifies the cluster ID.
	ClusterId pulumi.StringInput
	// Specifies the number of desired nodes.
	DesiredNodeCount pulumi.IntInput
	EnableForceNew   pulumi.StringPtrInput
	// Specifies the node pool ID.
	NodepoolId pulumi.StringInput
	// Specifies the charging period of the nodes.
	// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
	// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the nodes.
	// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
	// Changing this parameter will create a new cluster resource.
	PeriodUnit pulumi.StringPtrInput
	// Specifies the region in which to create the node pool scale resource.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the scalable checking.
	// The value can be **instant** and **async**, defaults to **instant**.
	ScalableChecking pulumi.StringPtrInput
	// Specifies the IDs of scale groups to scale.
	// **default** indicates the default group.
	ScaleGroups pulumi.StringArrayInput
}

func (NodePoolScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolScaleArgs)(nil)).Elem()
}

type NodePoolScaleInput interface {
	pulumi.Input

	ToNodePoolScaleOutput() NodePoolScaleOutput
	ToNodePoolScaleOutputWithContext(ctx context.Context) NodePoolScaleOutput
}

func (*NodePoolScale) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePoolScale)(nil)).Elem()
}

func (i *NodePoolScale) ToNodePoolScaleOutput() NodePoolScaleOutput {
	return i.ToNodePoolScaleOutputWithContext(context.Background())
}

func (i *NodePoolScale) ToNodePoolScaleOutputWithContext(ctx context.Context) NodePoolScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolScaleOutput)
}

// NodePoolScaleArrayInput is an input type that accepts NodePoolScaleArray and NodePoolScaleArrayOutput values.
// You can construct a concrete instance of `NodePoolScaleArrayInput` via:
//
//	NodePoolScaleArray{ NodePoolScaleArgs{...} }
type NodePoolScaleArrayInput interface {
	pulumi.Input

	ToNodePoolScaleArrayOutput() NodePoolScaleArrayOutput
	ToNodePoolScaleArrayOutputWithContext(context.Context) NodePoolScaleArrayOutput
}

type NodePoolScaleArray []NodePoolScaleInput

func (NodePoolScaleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePoolScale)(nil)).Elem()
}

func (i NodePoolScaleArray) ToNodePoolScaleArrayOutput() NodePoolScaleArrayOutput {
	return i.ToNodePoolScaleArrayOutputWithContext(context.Background())
}

func (i NodePoolScaleArray) ToNodePoolScaleArrayOutputWithContext(ctx context.Context) NodePoolScaleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolScaleArrayOutput)
}

// NodePoolScaleMapInput is an input type that accepts NodePoolScaleMap and NodePoolScaleMapOutput values.
// You can construct a concrete instance of `NodePoolScaleMapInput` via:
//
//	NodePoolScaleMap{ "key": NodePoolScaleArgs{...} }
type NodePoolScaleMapInput interface {
	pulumi.Input

	ToNodePoolScaleMapOutput() NodePoolScaleMapOutput
	ToNodePoolScaleMapOutputWithContext(context.Context) NodePoolScaleMapOutput
}

type NodePoolScaleMap map[string]NodePoolScaleInput

func (NodePoolScaleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePoolScale)(nil)).Elem()
}

func (i NodePoolScaleMap) ToNodePoolScaleMapOutput() NodePoolScaleMapOutput {
	return i.ToNodePoolScaleMapOutputWithContext(context.Background())
}

func (i NodePoolScaleMap) ToNodePoolScaleMapOutputWithContext(ctx context.Context) NodePoolScaleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolScaleMapOutput)
}

type NodePoolScaleOutput struct{ *pulumi.OutputState }

func (NodePoolScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePoolScale)(nil)).Elem()
}

func (o NodePoolScaleOutput) ToNodePoolScaleOutput() NodePoolScaleOutput {
	return o
}

func (o NodePoolScaleOutput) ToNodePoolScaleOutputWithContext(ctx context.Context) NodePoolScaleOutput {
	return o
}

// Specifies whether auto renew is enabled. Valid values are **true** and **false**.
func (o NodePoolScaleOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

// Specifies the charging mode of the nodes.
// Valid values are **prePaid** and **postPaid**, defaults to **postPaid**.
// Changing this parameter will create a new cluster resource.
func (o NodePoolScaleOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

// Specifies the cluster ID.
func (o NodePoolScaleOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Specifies the number of desired nodes.
func (o NodePoolScaleOutput) DesiredNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.IntOutput { return v.DesiredNodeCount }).(pulumi.IntOutput)
}

func (o NodePoolScaleOutput) EnableForceNew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringPtrOutput { return v.EnableForceNew }).(pulumi.StringPtrOutput)
}

// Specifies the node pool ID.
func (o NodePoolScaleOutput) NodepoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringOutput { return v.NodepoolId }).(pulumi.StringOutput)
}

// Specifies the charging period of the nodes.
// If `periodUnit` is set to **month**, the value ranges from 1 to 9.
// If `periodUnit` is set to **year**, the value ranges from 1 to 3.
// This parameter is mandatory if `chargingMode` is set to **prePaid**.
// Changing this parameter will create a new cluster resource.
func (o NodePoolScaleOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Specifies the charging period unit of the nodes.
// Valid values are **month** and **year**. This parameter is mandatory if `chargingMode` is set to **prePaid**.
// Changing this parameter will create a new cluster resource.
func (o NodePoolScaleOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// Specifies the region in which to create the node pool scale resource.
// If omitted, the provider-level region will be used. Changing this will create a new resource.
func (o NodePoolScaleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the scalable checking.
// The value can be **instant** and **async**, defaults to **instant**.
func (o NodePoolScaleOutput) ScalableChecking() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringPtrOutput { return v.ScalableChecking }).(pulumi.StringPtrOutput)
}

// Specifies the IDs of scale groups to scale.
// **default** indicates the default group.
func (o NodePoolScaleOutput) ScaleGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodePoolScale) pulumi.StringArrayOutput { return v.ScaleGroups }).(pulumi.StringArrayOutput)
}

type NodePoolScaleArrayOutput struct{ *pulumi.OutputState }

func (NodePoolScaleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePoolScale)(nil)).Elem()
}

func (o NodePoolScaleArrayOutput) ToNodePoolScaleArrayOutput() NodePoolScaleArrayOutput {
	return o
}

func (o NodePoolScaleArrayOutput) ToNodePoolScaleArrayOutputWithContext(ctx context.Context) NodePoolScaleArrayOutput {
	return o
}

func (o NodePoolScaleArrayOutput) Index(i pulumi.IntInput) NodePoolScaleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodePoolScale {
		return vs[0].([]*NodePoolScale)[vs[1].(int)]
	}).(NodePoolScaleOutput)
}

type NodePoolScaleMapOutput struct{ *pulumi.OutputState }

func (NodePoolScaleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePoolScale)(nil)).Elem()
}

func (o NodePoolScaleMapOutput) ToNodePoolScaleMapOutput() NodePoolScaleMapOutput {
	return o
}

func (o NodePoolScaleMapOutput) ToNodePoolScaleMapOutputWithContext(ctx context.Context) NodePoolScaleMapOutput {
	return o
}

func (o NodePoolScaleMapOutput) MapIndex(k pulumi.StringInput) NodePoolScaleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodePoolScale {
		return vs[0].(map[string]*NodePoolScale)[vs[1].(string)]
	}).(NodePoolScaleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolScaleInput)(nil)).Elem(), &NodePoolScale{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolScaleArrayInput)(nil)).Elem(), NodePoolScaleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolScaleMapInput)(nil)).Elem(), NodePoolScaleMap{})
	pulumi.RegisterOutputType(NodePoolScaleOutput{})
	pulumi.RegisterOutputType(NodePoolScaleArrayOutput{})
	pulumi.RegisterOutputType(NodePoolScaleMapOutput{})
}
