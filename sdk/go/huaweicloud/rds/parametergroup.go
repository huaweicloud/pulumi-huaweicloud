// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a RDS ParameterGroup resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Rds.NewParametergroup(ctx, "pg1", &Rds.ParametergroupArgs{
//				Datastore: &rds.ParametergroupDatastoreArgs{
//					Type:    pulumi.String("mysql"),
//					Version: pulumi.String("5.6"),
//				},
//				Description: pulumi.String("description_1"),
//				Values: pulumi.StringMap{
//					"autocommit":      pulumi.String("OFF"),
//					"max_connections": pulumi.String("10"),
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
// Parameter groups can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/parametergroup:Parametergroup pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c
//
// ```
type Parametergroup struct {
	pulumi.CustomResourceState

	// Indicates the parameter configuration defined by users based on the default parameters
	// groups.
	ConfigurationParameters ParametergroupConfigurationParameterArrayOutput `pulumi:"configurationParameters"`
	// The creation time, in UTC format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Database object. The database object structure is documented below. Changing
	// this creates a new parameter group.
	Datastore ParametergroupDatastoreOutput `pulumi:"datastore"`
	// The parameter group description. It contains a maximum of 256 characters and cannot
	// contain the following special characters:>!<"&'= the value is left blank by default.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The parameter group name. It contains a maximum of 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the RDS parameter group. If omitted, the
	// provider-level region will be used. Changing this creates a new parameter group.
	Region pulumi.StringOutput `pulumi:"region"`
	// The last update time, in UTC format.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Parameter group values key/value pairs defined by users based on the default parameter
	// groups.
	Values pulumi.StringMapOutput `pulumi:"values"`
}

// NewParametergroup registers a new resource with the given unique name, arguments, and options.
func NewParametergroup(ctx *pulumi.Context,
	name string, args *ParametergroupArgs, opts ...pulumi.ResourceOption) (*Parametergroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Parametergroup
	err := ctx.RegisterResource("huaweicloud:Rds/parametergroup:Parametergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParametergroup gets an existing Parametergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParametergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParametergroupState, opts ...pulumi.ResourceOption) (*Parametergroup, error) {
	var resource Parametergroup
	err := ctx.ReadResource("huaweicloud:Rds/parametergroup:Parametergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parametergroup resources.
type parametergroupState struct {
	// Indicates the parameter configuration defined by users based on the default parameters
	// groups.
	ConfigurationParameters []ParametergroupConfigurationParameter `pulumi:"configurationParameters"`
	// The creation time, in UTC format.
	CreatedAt *string `pulumi:"createdAt"`
	// Database object. The database object structure is documented below. Changing
	// this creates a new parameter group.
	Datastore *ParametergroupDatastore `pulumi:"datastore"`
	// The parameter group description. It contains a maximum of 256 characters and cannot
	// contain the following special characters:>!<"&'= the value is left blank by default.
	Description *string `pulumi:"description"`
	// The parameter group name. It contains a maximum of 64 characters.
	Name *string `pulumi:"name"`
	// The region in which to create the RDS parameter group. If omitted, the
	// provider-level region will be used. Changing this creates a new parameter group.
	Region *string `pulumi:"region"`
	// The last update time, in UTC format.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Parameter group values key/value pairs defined by users based on the default parameter
	// groups.
	Values map[string]string `pulumi:"values"`
}

type ParametergroupState struct {
	// Indicates the parameter configuration defined by users based on the default parameters
	// groups.
	ConfigurationParameters ParametergroupConfigurationParameterArrayInput
	// The creation time, in UTC format.
	CreatedAt pulumi.StringPtrInput
	// Database object. The database object structure is documented below. Changing
	// this creates a new parameter group.
	Datastore ParametergroupDatastorePtrInput
	// The parameter group description. It contains a maximum of 256 characters and cannot
	// contain the following special characters:>!<"&'= the value is left blank by default.
	Description pulumi.StringPtrInput
	// The parameter group name. It contains a maximum of 64 characters.
	Name pulumi.StringPtrInput
	// The region in which to create the RDS parameter group. If omitted, the
	// provider-level region will be used. Changing this creates a new parameter group.
	Region pulumi.StringPtrInput
	// The last update time, in UTC format.
	UpdatedAt pulumi.StringPtrInput
	// Parameter group values key/value pairs defined by users based on the default parameter
	// groups.
	Values pulumi.StringMapInput
}

func (ParametergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parametergroupState)(nil)).Elem()
}

type parametergroupArgs struct {
	// Database object. The database object structure is documented below. Changing
	// this creates a new parameter group.
	Datastore ParametergroupDatastore `pulumi:"datastore"`
	// The parameter group description. It contains a maximum of 256 characters and cannot
	// contain the following special characters:>!<"&'= the value is left blank by default.
	Description *string `pulumi:"description"`
	// The parameter group name. It contains a maximum of 64 characters.
	Name *string `pulumi:"name"`
	// The region in which to create the RDS parameter group. If omitted, the
	// provider-level region will be used. Changing this creates a new parameter group.
	Region *string `pulumi:"region"`
	// Parameter group values key/value pairs defined by users based on the default parameter
	// groups.
	Values map[string]string `pulumi:"values"`
}

// The set of arguments for constructing a Parametergroup resource.
type ParametergroupArgs struct {
	// Database object. The database object structure is documented below. Changing
	// this creates a new parameter group.
	Datastore ParametergroupDatastoreInput
	// The parameter group description. It contains a maximum of 256 characters and cannot
	// contain the following special characters:>!<"&'= the value is left blank by default.
	Description pulumi.StringPtrInput
	// The parameter group name. It contains a maximum of 64 characters.
	Name pulumi.StringPtrInput
	// The region in which to create the RDS parameter group. If omitted, the
	// provider-level region will be used. Changing this creates a new parameter group.
	Region pulumi.StringPtrInput
	// Parameter group values key/value pairs defined by users based on the default parameter
	// groups.
	Values pulumi.StringMapInput
}

func (ParametergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parametergroupArgs)(nil)).Elem()
}

type ParametergroupInput interface {
	pulumi.Input

	ToParametergroupOutput() ParametergroupOutput
	ToParametergroupOutputWithContext(ctx context.Context) ParametergroupOutput
}

func (*Parametergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Parametergroup)(nil)).Elem()
}

func (i *Parametergroup) ToParametergroupOutput() ParametergroupOutput {
	return i.ToParametergroupOutputWithContext(context.Background())
}

func (i *Parametergroup) ToParametergroupOutputWithContext(ctx context.Context) ParametergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametergroupOutput)
}

// ParametergroupArrayInput is an input type that accepts ParametergroupArray and ParametergroupArrayOutput values.
// You can construct a concrete instance of `ParametergroupArrayInput` via:
//
//	ParametergroupArray{ ParametergroupArgs{...} }
type ParametergroupArrayInput interface {
	pulumi.Input

	ToParametergroupArrayOutput() ParametergroupArrayOutput
	ToParametergroupArrayOutputWithContext(context.Context) ParametergroupArrayOutput
}

type ParametergroupArray []ParametergroupInput

func (ParametergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parametergroup)(nil)).Elem()
}

func (i ParametergroupArray) ToParametergroupArrayOutput() ParametergroupArrayOutput {
	return i.ToParametergroupArrayOutputWithContext(context.Background())
}

func (i ParametergroupArray) ToParametergroupArrayOutputWithContext(ctx context.Context) ParametergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametergroupArrayOutput)
}

// ParametergroupMapInput is an input type that accepts ParametergroupMap and ParametergroupMapOutput values.
// You can construct a concrete instance of `ParametergroupMapInput` via:
//
//	ParametergroupMap{ "key": ParametergroupArgs{...} }
type ParametergroupMapInput interface {
	pulumi.Input

	ToParametergroupMapOutput() ParametergroupMapOutput
	ToParametergroupMapOutputWithContext(context.Context) ParametergroupMapOutput
}

type ParametergroupMap map[string]ParametergroupInput

func (ParametergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parametergroup)(nil)).Elem()
}

func (i ParametergroupMap) ToParametergroupMapOutput() ParametergroupMapOutput {
	return i.ToParametergroupMapOutputWithContext(context.Background())
}

func (i ParametergroupMap) ToParametergroupMapOutputWithContext(ctx context.Context) ParametergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametergroupMapOutput)
}

type ParametergroupOutput struct{ *pulumi.OutputState }

func (ParametergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parametergroup)(nil)).Elem()
}

func (o ParametergroupOutput) ToParametergroupOutput() ParametergroupOutput {
	return o
}

func (o ParametergroupOutput) ToParametergroupOutputWithContext(ctx context.Context) ParametergroupOutput {
	return o
}

// Indicates the parameter configuration defined by users based on the default parameters
// groups.
func (o ParametergroupOutput) ConfigurationParameters() ParametergroupConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *Parametergroup) ParametergroupConfigurationParameterArrayOutput {
		return v.ConfigurationParameters
	}).(ParametergroupConfigurationParameterArrayOutput)
}

// The creation time, in UTC format.
func (o ParametergroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Database object. The database object structure is documented below. Changing
// this creates a new parameter group.
func (o ParametergroupOutput) Datastore() ParametergroupDatastoreOutput {
	return o.ApplyT(func(v *Parametergroup) ParametergroupDatastoreOutput { return v.Datastore }).(ParametergroupDatastoreOutput)
}

// The parameter group description. It contains a maximum of 256 characters and cannot
// contain the following special characters:>!<"&'= the value is left blank by default.
func (o ParametergroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The parameter group name. It contains a maximum of 64 characters.
func (o ParametergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the RDS parameter group. If omitted, the
// provider-level region will be used. Changing this creates a new parameter group.
func (o ParametergroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The last update time, in UTC format.
func (o ParametergroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Parameter group values key/value pairs defined by users based on the default parameter
// groups.
func (o ParametergroupOutput) Values() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Parametergroup) pulumi.StringMapOutput { return v.Values }).(pulumi.StringMapOutput)
}

type ParametergroupArrayOutput struct{ *pulumi.OutputState }

func (ParametergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parametergroup)(nil)).Elem()
}

func (o ParametergroupArrayOutput) ToParametergroupArrayOutput() ParametergroupArrayOutput {
	return o
}

func (o ParametergroupArrayOutput) ToParametergroupArrayOutputWithContext(ctx context.Context) ParametergroupArrayOutput {
	return o
}

func (o ParametergroupArrayOutput) Index(i pulumi.IntInput) ParametergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parametergroup {
		return vs[0].([]*Parametergroup)[vs[1].(int)]
	}).(ParametergroupOutput)
}

type ParametergroupMapOutput struct{ *pulumi.OutputState }

func (ParametergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parametergroup)(nil)).Elem()
}

func (o ParametergroupMapOutput) ToParametergroupMapOutput() ParametergroupMapOutput {
	return o
}

func (o ParametergroupMapOutput) ToParametergroupMapOutputWithContext(ctx context.Context) ParametergroupMapOutput {
	return o
}

func (o ParametergroupMapOutput) MapIndex(k pulumi.StringInput) ParametergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parametergroup {
		return vs[0].(map[string]*Parametergroup)[vs[1].(string)]
	}).(ParametergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParametergroupInput)(nil)).Elem(), &Parametergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametergroupArrayInput)(nil)).Elem(), ParametergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametergroupMapInput)(nil)).Elem(), ParametergroupMap{})
	pulumi.RegisterOutputType(ParametergroupOutput{})
	pulumi.RegisterOutputType(ParametergroupArrayOutput{})
	pulumi.RegisterOutputType(ParametergroupMapOutput{})
}
