// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cce

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a CCE namespace resource within HuaweiCloud.
//
// ## Example Usage
// ### Basic
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
//			_, err := Cce.NewNamespace(ctx, "test", &Cce.NamespaceArgs{
//				ClusterId: pulumi.Any(clusterId),
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
// # CCE namespace can be imported using the cluster ID and namespace name separated by a slash, e.g.bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Cce/namespace:Namespace test bb6923e4-b16e-11eb-b0cd-0255ac101da1/test-namespace
//
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The server time when namespace was created.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Specifies the unique name of the namespace.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a prefix used by the server to generate a unique name.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current phase of the namespace.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("huaweicloud:Cce/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("huaweicloud:Cce/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	ClusterId *string `pulumi:"clusterId"`
	// The server time when namespace was created.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the unique name of the namespace.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Name *string `pulumi:"name"`
	// Specifies a prefix used by the server to generate a unique name.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Prefix *string `pulumi:"prefix"`
	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	Region *string `pulumi:"region"`
	// The current phase of the namespace.
	Status *string `pulumi:"status"`
}

type NamespaceState struct {
	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	Annotations pulumi.StringMapInput
	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	ClusterId pulumi.StringPtrInput
	// The server time when namespace was created.
	CreationTimestamp pulumi.StringPtrInput
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	Labels pulumi.StringMapInput
	// Specifies the unique name of the namespace.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Name pulumi.StringPtrInput
	// Specifies a prefix used by the server to generate a unique name.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Prefix pulumi.StringPtrInput
	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	Region pulumi.StringPtrInput
	// The current phase of the namespace.
	Status pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the unique name of the namespace.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Name *string `pulumi:"name"`
	// Specifies a prefix used by the server to generate a unique name.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Prefix *string `pulumi:"prefix"`
	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	Annotations pulumi.StringMapInput
	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	ClusterId pulumi.StringInput
	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	Labels pulumi.StringMapInput
	// Specifies the unique name of the namespace.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Name pulumi.StringPtrInput
	// Specifies a prefix used by the server to generate a unique name.\
	// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource.
	Prefix pulumi.StringPtrInput
	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	Region pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//	NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//	NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// Specifies an unstructured key value map for external parameters.
// Changing this will create a new namespace resource.
func (o NamespaceOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Specifies the cluster ID to which the CCE namespace belongs.
// Changing this will create a new namespace resource.
func (o NamespaceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The server time when namespace was created.
func (o NamespaceOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Specifies the map of string keys and values for labels.
// Changing this will create a new namespace resource.
func (o NamespaceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Specifies the unique name of the namespace.\
// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
// hyphens (-), and must start and end with lowercase letters and digits.
// Changing this will create a new namespace resource.
func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies a prefix used by the server to generate a unique name.\
// This parameter can contain a maximum of `63` characters, which may consist of lowercase letters, digits and
// hyphens (-), and must start and end with lowercase letters and digits.
// Changing this will create a new namespace resource.
func (o NamespaceOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

// Specifies the region in which to create the namespace resource.
// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
func (o NamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current phase of the namespace.
func (o NamespaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
