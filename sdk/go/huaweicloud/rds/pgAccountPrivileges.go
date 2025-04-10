// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS PostgreSQL account privileges resource within HuaweiCloud.
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
//			userName := cfg.RequireObject("userName")
//			_, err := Rds.NewPgAccountPrivileges(ctx, "test", &Rds.PgAccountPrivilegesArgs{
//				InstanceId: pulumi.Any(instanceId),
//				UserName:   pulumi.Any(userName),
//				RolePrivileges: pulumi.StringArray{
//					pulumi.String("CREATEROLE"),
//					pulumi.String("LOGIN"),
//					pulumi.String("REPLICATION"),
//				},
//				SystemRolePrivileges: pulumi.StringArray{
//					pulumi.String("pg_signal_backend"),
//					pulumi.String("root"),
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
// The RDS PostgreSQL privileges roles can be imported using the `instance_id` and `user_name` separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/pgAccountPrivileges:PgAccountPrivileges test <instance_id>/<user_name>
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`system_role_privileges`. It is generally recommended running `terraform plan` after importing the RDS PostgreSQL account privileges. You can then decide if changes should be applied to the RDS PostgreSQL account privileges, or the RDS PostgreSQL account privileges definition should be updated to align with the account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_pg_account_privileges" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	system_role_privileges,
//
//	]
//
//	} }
type PgAccountPrivileges struct {
	pulumi.CustomResourceState

	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the list of role privileges. Value options: **CREATEDB**,
	// **CREATEROLE**, **LOGIN**, **REPLICATION**.
	RolePrivileges pulumi.StringArrayOutput `pulumi:"rolePrivileges"`
	// Specifies the list of system role privileges. Value options:
	// **pg_monitor**, **pg_signal_backend**, **root**.
	SystemRolePrivileges pulumi.StringArrayOutput `pulumi:"systemRolePrivileges"`
	// Specifies the username of the account.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewPgAccountPrivileges registers a new resource with the given unique name, arguments, and options.
func NewPgAccountPrivileges(ctx *pulumi.Context,
	name string, args *PgAccountPrivilegesArgs, opts ...pulumi.ResourceOption) (*PgAccountPrivileges, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PgAccountPrivileges
	err := ctx.RegisterResource("huaweicloud:Rds/pgAccountPrivileges:PgAccountPrivileges", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgAccountPrivileges gets an existing PgAccountPrivileges resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgAccountPrivileges(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgAccountPrivilegesState, opts ...pulumi.ResourceOption) (*PgAccountPrivileges, error) {
	var resource PgAccountPrivileges
	err := ctx.ReadResource("huaweicloud:Rds/pgAccountPrivileges:PgAccountPrivileges", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgAccountPrivileges resources.
type pgAccountPrivilegesState struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the list of role privileges. Value options: **CREATEDB**,
	// **CREATEROLE**, **LOGIN**, **REPLICATION**.
	RolePrivileges []string `pulumi:"rolePrivileges"`
	// Specifies the list of system role privileges. Value options:
	// **pg_monitor**, **pg_signal_backend**, **root**.
	SystemRolePrivileges []string `pulumi:"systemRolePrivileges"`
	// Specifies the username of the account.
	UserName *string `pulumi:"userName"`
}

type PgAccountPrivilegesState struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the list of role privileges. Value options: **CREATEDB**,
	// **CREATEROLE**, **LOGIN**, **REPLICATION**.
	RolePrivileges pulumi.StringArrayInput
	// Specifies the list of system role privileges. Value options:
	// **pg_monitor**, **pg_signal_backend**, **root**.
	SystemRolePrivileges pulumi.StringArrayInput
	// Specifies the username of the account.
	UserName pulumi.StringPtrInput
}

func (PgAccountPrivilegesState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountPrivilegesState)(nil)).Elem()
}

type pgAccountPrivilegesArgs struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the list of role privileges. Value options: **CREATEDB**,
	// **CREATEROLE**, **LOGIN**, **REPLICATION**.
	RolePrivileges []string `pulumi:"rolePrivileges"`
	// Specifies the list of system role privileges. Value options:
	// **pg_monitor**, **pg_signal_backend**, **root**.
	SystemRolePrivileges []string `pulumi:"systemRolePrivileges"`
	// Specifies the username of the account.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a PgAccountPrivileges resource.
type PgAccountPrivilegesArgs struct {
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the list of role privileges. Value options: **CREATEDB**,
	// **CREATEROLE**, **LOGIN**, **REPLICATION**.
	RolePrivileges pulumi.StringArrayInput
	// Specifies the list of system role privileges. Value options:
	// **pg_monitor**, **pg_signal_backend**, **root**.
	SystemRolePrivileges pulumi.StringArrayInput
	// Specifies the username of the account.
	UserName pulumi.StringInput
}

func (PgAccountPrivilegesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountPrivilegesArgs)(nil)).Elem()
}

type PgAccountPrivilegesInput interface {
	pulumi.Input

	ToPgAccountPrivilegesOutput() PgAccountPrivilegesOutput
	ToPgAccountPrivilegesOutputWithContext(ctx context.Context) PgAccountPrivilegesOutput
}

func (*PgAccountPrivileges) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccountPrivileges)(nil)).Elem()
}

func (i *PgAccountPrivileges) ToPgAccountPrivilegesOutput() PgAccountPrivilegesOutput {
	return i.ToPgAccountPrivilegesOutputWithContext(context.Background())
}

func (i *PgAccountPrivileges) ToPgAccountPrivilegesOutputWithContext(ctx context.Context) PgAccountPrivilegesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountPrivilegesOutput)
}

// PgAccountPrivilegesArrayInput is an input type that accepts PgAccountPrivilegesArray and PgAccountPrivilegesArrayOutput values.
// You can construct a concrete instance of `PgAccountPrivilegesArrayInput` via:
//
//	PgAccountPrivilegesArray{ PgAccountPrivilegesArgs{...} }
type PgAccountPrivilegesArrayInput interface {
	pulumi.Input

	ToPgAccountPrivilegesArrayOutput() PgAccountPrivilegesArrayOutput
	ToPgAccountPrivilegesArrayOutputWithContext(context.Context) PgAccountPrivilegesArrayOutput
}

type PgAccountPrivilegesArray []PgAccountPrivilegesInput

func (PgAccountPrivilegesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccountPrivileges)(nil)).Elem()
}

func (i PgAccountPrivilegesArray) ToPgAccountPrivilegesArrayOutput() PgAccountPrivilegesArrayOutput {
	return i.ToPgAccountPrivilegesArrayOutputWithContext(context.Background())
}

func (i PgAccountPrivilegesArray) ToPgAccountPrivilegesArrayOutputWithContext(ctx context.Context) PgAccountPrivilegesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountPrivilegesArrayOutput)
}

// PgAccountPrivilegesMapInput is an input type that accepts PgAccountPrivilegesMap and PgAccountPrivilegesMapOutput values.
// You can construct a concrete instance of `PgAccountPrivilegesMapInput` via:
//
//	PgAccountPrivilegesMap{ "key": PgAccountPrivilegesArgs{...} }
type PgAccountPrivilegesMapInput interface {
	pulumi.Input

	ToPgAccountPrivilegesMapOutput() PgAccountPrivilegesMapOutput
	ToPgAccountPrivilegesMapOutputWithContext(context.Context) PgAccountPrivilegesMapOutput
}

type PgAccountPrivilegesMap map[string]PgAccountPrivilegesInput

func (PgAccountPrivilegesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccountPrivileges)(nil)).Elem()
}

func (i PgAccountPrivilegesMap) ToPgAccountPrivilegesMapOutput() PgAccountPrivilegesMapOutput {
	return i.ToPgAccountPrivilegesMapOutputWithContext(context.Background())
}

func (i PgAccountPrivilegesMap) ToPgAccountPrivilegesMapOutputWithContext(ctx context.Context) PgAccountPrivilegesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountPrivilegesMapOutput)
}

type PgAccountPrivilegesOutput struct{ *pulumi.OutputState }

func (PgAccountPrivilegesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccountPrivileges)(nil)).Elem()
}

func (o PgAccountPrivilegesOutput) ToPgAccountPrivilegesOutput() PgAccountPrivilegesOutput {
	return o
}

func (o PgAccountPrivilegesOutput) ToPgAccountPrivilegesOutputWithContext(ctx context.Context) PgAccountPrivilegesOutput {
	return o
}

// Specifies the ID of the RDS PostgreSQL instance.
func (o PgAccountPrivilegesOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountPrivileges) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o PgAccountPrivilegesOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountPrivileges) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the list of role privileges. Value options: **CREATEDB**,
// **CREATEROLE**, **LOGIN**, **REPLICATION**.
func (o PgAccountPrivilegesOutput) RolePrivileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PgAccountPrivileges) pulumi.StringArrayOutput { return v.RolePrivileges }).(pulumi.StringArrayOutput)
}

// Specifies the list of system role privileges. Value options:
// **pg_monitor**, **pg_signal_backend**, **root**.
func (o PgAccountPrivilegesOutput) SystemRolePrivileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PgAccountPrivileges) pulumi.StringArrayOutput { return v.SystemRolePrivileges }).(pulumi.StringArrayOutput)
}

// Specifies the username of the account.
func (o PgAccountPrivilegesOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccountPrivileges) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type PgAccountPrivilegesArrayOutput struct{ *pulumi.OutputState }

func (PgAccountPrivilegesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccountPrivileges)(nil)).Elem()
}

func (o PgAccountPrivilegesArrayOutput) ToPgAccountPrivilegesArrayOutput() PgAccountPrivilegesArrayOutput {
	return o
}

func (o PgAccountPrivilegesArrayOutput) ToPgAccountPrivilegesArrayOutputWithContext(ctx context.Context) PgAccountPrivilegesArrayOutput {
	return o
}

func (o PgAccountPrivilegesArrayOutput) Index(i pulumi.IntInput) PgAccountPrivilegesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgAccountPrivileges {
		return vs[0].([]*PgAccountPrivileges)[vs[1].(int)]
	}).(PgAccountPrivilegesOutput)
}

type PgAccountPrivilegesMapOutput struct{ *pulumi.OutputState }

func (PgAccountPrivilegesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccountPrivileges)(nil)).Elem()
}

func (o PgAccountPrivilegesMapOutput) ToPgAccountPrivilegesMapOutput() PgAccountPrivilegesMapOutput {
	return o
}

func (o PgAccountPrivilegesMapOutput) ToPgAccountPrivilegesMapOutputWithContext(ctx context.Context) PgAccountPrivilegesMapOutput {
	return o
}

func (o PgAccountPrivilegesMapOutput) MapIndex(k pulumi.StringInput) PgAccountPrivilegesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgAccountPrivileges {
		return vs[0].(map[string]*PgAccountPrivileges)[vs[1].(string)]
	}).(PgAccountPrivilegesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountPrivilegesInput)(nil)).Elem(), &PgAccountPrivileges{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountPrivilegesArrayInput)(nil)).Elem(), PgAccountPrivilegesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountPrivilegesMapInput)(nil)).Elem(), PgAccountPrivilegesMap{})
	pulumi.RegisterOutputType(PgAccountPrivilegesOutput{})
	pulumi.RegisterOutputType(PgAccountPrivilegesArrayOutput{})
	pulumi.RegisterOutputType(PgAccountPrivilegesMapOutput{})
}
