// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS PostgreSQL account roles resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			_, err := Rds.NewPgAccountRoles(ctx, "test", &Rds.PgAccountRolesArgs{
//				InstanceId: pulumi.Any(instanceId),
//				User:       pulumi.String("test_user"),
//				Roles: pulumi.StringArray{
//					pulumi.String("test111"),
//					pulumi.String("test222"),
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
// The RDS PostgreSQL account roles can be imported using the `instance_id` and `name` separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/pgAccountRoles:PgAccountRoles test <instance_id>/<name>
//
// ```
type PgAccountRoles struct {
	pulumi.CustomResourceState

	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the list of roles.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Specifies the username of the account.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewPgAccountRoles registers a new resource with the given unique name, arguments, and options.
func NewPgAccountRoles(ctx *pulumi.Context,
	name string, args *PgAccountRolesArgs, opts ...pulumi.ResourceOption) (*PgAccountRoles, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PgAccountRoles
	err := ctx.RegisterResource("huaweicloud:Rds/pgAccountRoles:PgAccountRoles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgAccountRoles gets an existing PgAccountRoles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgAccountRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgAccountRolesState, opts ...pulumi.ResourceOption) (*PgAccountRoles, error) {
	var resource PgAccountRoles
	err := ctx.ReadResource("huaweicloud:Rds/pgAccountRoles:PgAccountRoles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgAccountRoles resources.
type pgAccountRolesState struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the list of roles.
	Roles []string `pulumi:"roles"`
	// Specifies the username of the account.
	User *string `pulumi:"user"`
}

type PgAccountRolesState struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the list of roles.
	Roles pulumi.StringArrayInput
	// Specifies the username of the account.
	User pulumi.StringPtrInput
}

func (PgAccountRolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountRolesState)(nil)).Elem()
}

type pgAccountRolesArgs struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the list of roles.
	Roles []string `pulumi:"roles"`
	// Specifies the username of the account.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a PgAccountRoles resource.
type PgAccountRolesArgs struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the list of roles.
	Roles pulumi.StringArrayInput
	// Specifies the username of the account.
	User pulumi.StringInput
}

func (PgAccountRolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountRolesArgs)(nil)).Elem()
}

type PgAccountRolesInput interface {
	pulumi.Input

	ToPgAccountRolesOutput() PgAccountRolesOutput
	ToPgAccountRolesOutputWithContext(ctx context.Context) PgAccountRolesOutput
}

func (*PgAccountRoles) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccountRoles)(nil)).Elem()
}

func (i *PgAccountRoles) ToPgAccountRolesOutput() PgAccountRolesOutput {
	return i.ToPgAccountRolesOutputWithContext(context.Background())
}

func (i *PgAccountRoles) ToPgAccountRolesOutputWithContext(ctx context.Context) PgAccountRolesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountRolesOutput)
}

// PgAccountRolesArrayInput is an input type that accepts PgAccountRolesArray and PgAccountRolesArrayOutput values.
// You can construct a concrete instance of `PgAccountRolesArrayInput` via:
//
//	PgAccountRolesArray{ PgAccountRolesArgs{...} }
type PgAccountRolesArrayInput interface {
	pulumi.Input

	ToPgAccountRolesArrayOutput() PgAccountRolesArrayOutput
	ToPgAccountRolesArrayOutputWithContext(context.Context) PgAccountRolesArrayOutput
}

type PgAccountRolesArray []PgAccountRolesInput

func (PgAccountRolesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccountRoles)(nil)).Elem()
}

func (i PgAccountRolesArray) ToPgAccountRolesArrayOutput() PgAccountRolesArrayOutput {
	return i.ToPgAccountRolesArrayOutputWithContext(context.Background())
}

func (i PgAccountRolesArray) ToPgAccountRolesArrayOutputWithContext(ctx context.Context) PgAccountRolesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountRolesArrayOutput)
}

// PgAccountRolesMapInput is an input type that accepts PgAccountRolesMap and PgAccountRolesMapOutput values.
// You can construct a concrete instance of `PgAccountRolesMapInput` via:
//
//	PgAccountRolesMap{ "key": PgAccountRolesArgs{...} }
type PgAccountRolesMapInput interface {
	pulumi.Input

	ToPgAccountRolesMapOutput() PgAccountRolesMapOutput
	ToPgAccountRolesMapOutputWithContext(context.Context) PgAccountRolesMapOutput
}

type PgAccountRolesMap map[string]PgAccountRolesInput

func (PgAccountRolesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccountRoles)(nil)).Elem()
}

func (i PgAccountRolesMap) ToPgAccountRolesMapOutput() PgAccountRolesMapOutput {
	return i.ToPgAccountRolesMapOutputWithContext(context.Background())
}

func (i PgAccountRolesMap) ToPgAccountRolesMapOutputWithContext(ctx context.Context) PgAccountRolesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountRolesMapOutput)
}

type PgAccountRolesOutput struct{ *pulumi.OutputState }

func (PgAccountRolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccountRoles)(nil)).Elem()
}

func (o PgAccountRolesOutput) ToPgAccountRolesOutput() PgAccountRolesOutput {
	return o
}

func (o PgAccountRolesOutput) ToPgAccountRolesOutputWithContext(ctx context.Context) PgAccountRolesOutput {
	return o
}

// Specifies the ID of the RDS PostgreSQL instance.
func (o PgAccountRolesOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountRoles) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o PgAccountRolesOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountRoles) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the list of roles.
func (o PgAccountRolesOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PgAccountRoles) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// Specifies the username of the account.
func (o PgAccountRolesOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountRoles) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type PgAccountRolesArrayOutput struct{ *pulumi.OutputState }

func (PgAccountRolesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccountRoles)(nil)).Elem()
}

func (o PgAccountRolesArrayOutput) ToPgAccountRolesArrayOutput() PgAccountRolesArrayOutput {
	return o
}

func (o PgAccountRolesArrayOutput) ToPgAccountRolesArrayOutputWithContext(ctx context.Context) PgAccountRolesArrayOutput {
	return o
}

func (o PgAccountRolesArrayOutput) Index(i pulumi.IntInput) PgAccountRolesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgAccountRoles {
		return vs[0].([]*PgAccountRoles)[vs[1].(int)]
	}).(PgAccountRolesOutput)
}

type PgAccountRolesMapOutput struct{ *pulumi.OutputState }

func (PgAccountRolesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccountRoles)(nil)).Elem()
}

func (o PgAccountRolesMapOutput) ToPgAccountRolesMapOutput() PgAccountRolesMapOutput {
	return o
}

func (o PgAccountRolesMapOutput) ToPgAccountRolesMapOutputWithContext(ctx context.Context) PgAccountRolesMapOutput {
	return o
}

func (o PgAccountRolesMapOutput) MapIndex(k pulumi.StringInput) PgAccountRolesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgAccountRoles {
		return vs[0].(map[string]*PgAccountRoles)[vs[1].(string)]
	}).(PgAccountRolesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountRolesInput)(nil)).Elem(), &PgAccountRoles{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountRolesArrayInput)(nil)).Elem(), PgAccountRolesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountRolesMapInput)(nil)).Elem(), PgAccountRolesMap{})
	pulumi.RegisterOutputType(PgAccountRolesOutput{})
	pulumi.RegisterOutputType(PgAccountRolesArrayOutput{})
	pulumi.RegisterOutputType(PgAccountRolesMapOutput{})
}
