// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DDS database role resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			roleName := cfg.RequireObject("roleName")
//			dbName := cfg.RequireObject("dbName")
//			ownedRoleName := cfg.RequireObject("ownedRoleName")
//			ownedRoleDbName := cfg.RequireObject("ownedRoleDbName")
//			_, err := Dds.NewDatabaseRole(ctx, "test", &Dds.DatabaseRoleArgs{
//				InstanceId: pulumi.Any(instanceId),
//				DbName:     pulumi.Any(dbName),
//				Roles: dds.DatabaseRoleRoleArray{
//					&dds.DatabaseRoleRoleArgs{
//						Name:   pulumi.Any(ownedRoleName),
//						DbName: pulumi.Any(ownedRoleDbName),
//					},
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
// DDS database roles can be imported using the `instance_id`, `db_name` and `name` separated by slashes (/), e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Dds/databaseRole:DatabaseRole test <instance_id>/<db_name>/<name>
//
// ```
type DatabaseRole struct {
	pulumi.CustomResourceState

	// Specifies the database name to which this owned role belongs.
	// Changing this parameter will create a new role.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// The list of database privileges owned by the current role, includes all privileges
	// inherited by owned roles. The inheritedPrivileges structure is documented below.
	InheritedPrivileges DatabaseRoleInheritedPrivilegeArrayOutput `pulumi:"inheritedPrivileges"`
	// Specifies the DDS instance ID to which the role belongs.
	// Changing this parameter will create a new role.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the name of role owned by the current role.
	// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
	// allowed. Changing this parameter will create a new role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of database privileges owned by the current role.
	// The privileges structure is documented below.
	Privileges DatabaseRolePrivilegeArrayOutput `pulumi:"privileges"`
	// Specifies the region where the DDS instance is located.
	// Changing this parameter will create a new role.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the list of roles owned by the current role.
	// The roles structure is documented below.
	// Changing this parameter will create a new role.
	Roles DatabaseRoleRoleArrayOutput `pulumi:"roles"`
}

// NewDatabaseRole registers a new resource with the given unique name, arguments, and options.
func NewDatabaseRole(ctx *pulumi.Context,
	name string, args *DatabaseRoleArgs, opts ...pulumi.ResourceOption) (*DatabaseRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DatabaseRole
	err := ctx.RegisterResource("huaweicloud:Dds/databaseRole:DatabaseRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseRole gets an existing DatabaseRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseRoleState, opts ...pulumi.ResourceOption) (*DatabaseRole, error) {
	var resource DatabaseRole
	err := ctx.ReadResource("huaweicloud:Dds/databaseRole:DatabaseRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseRole resources.
type databaseRoleState struct {
	// Specifies the database name to which this owned role belongs.
	// Changing this parameter will create a new role.
	DbName *string `pulumi:"dbName"`
	// The list of database privileges owned by the current role, includes all privileges
	// inherited by owned roles. The inheritedPrivileges structure is documented below.
	InheritedPrivileges []DatabaseRoleInheritedPrivilege `pulumi:"inheritedPrivileges"`
	// Specifies the DDS instance ID to which the role belongs.
	// Changing this parameter will create a new role.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the name of role owned by the current role.
	// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
	// allowed. Changing this parameter will create a new role.
	Name *string `pulumi:"name"`
	// The list of database privileges owned by the current role.
	// The privileges structure is documented below.
	Privileges []DatabaseRolePrivilege `pulumi:"privileges"`
	// Specifies the region where the DDS instance is located.
	// Changing this parameter will create a new role.
	Region *string `pulumi:"region"`
	// Specifies the list of roles owned by the current role.
	// The roles structure is documented below.
	// Changing this parameter will create a new role.
	Roles []DatabaseRoleRole `pulumi:"roles"`
}

type DatabaseRoleState struct {
	// Specifies the database name to which this owned role belongs.
	// Changing this parameter will create a new role.
	DbName pulumi.StringPtrInput
	// The list of database privileges owned by the current role, includes all privileges
	// inherited by owned roles. The inheritedPrivileges structure is documented below.
	InheritedPrivileges DatabaseRoleInheritedPrivilegeArrayInput
	// Specifies the DDS instance ID to which the role belongs.
	// Changing this parameter will create a new role.
	InstanceId pulumi.StringPtrInput
	// Specifies the name of role owned by the current role.
	// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
	// allowed. Changing this parameter will create a new role.
	Name pulumi.StringPtrInput
	// The list of database privileges owned by the current role.
	// The privileges structure is documented below.
	Privileges DatabaseRolePrivilegeArrayInput
	// Specifies the region where the DDS instance is located.
	// Changing this parameter will create a new role.
	Region pulumi.StringPtrInput
	// Specifies the list of roles owned by the current role.
	// The roles structure is documented below.
	// Changing this parameter will create a new role.
	Roles DatabaseRoleRoleArrayInput
}

func (DatabaseRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseRoleState)(nil)).Elem()
}

type databaseRoleArgs struct {
	// Specifies the database name to which this owned role belongs.
	// Changing this parameter will create a new role.
	DbName string `pulumi:"dbName"`
	// Specifies the DDS instance ID to which the role belongs.
	// Changing this parameter will create a new role.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the name of role owned by the current role.
	// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
	// allowed. Changing this parameter will create a new role.
	Name *string `pulumi:"name"`
	// Specifies the region where the DDS instance is located.
	// Changing this parameter will create a new role.
	Region *string `pulumi:"region"`
	// Specifies the list of roles owned by the current role.
	// The roles structure is documented below.
	// Changing this parameter will create a new role.
	Roles []DatabaseRoleRole `pulumi:"roles"`
}

// The set of arguments for constructing a DatabaseRole resource.
type DatabaseRoleArgs struct {
	// Specifies the database name to which this owned role belongs.
	// Changing this parameter will create a new role.
	DbName pulumi.StringInput
	// Specifies the DDS instance ID to which the role belongs.
	// Changing this parameter will create a new role.
	InstanceId pulumi.StringInput
	// Specifies the name of role owned by the current role.
	// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
	// allowed. Changing this parameter will create a new role.
	Name pulumi.StringPtrInput
	// Specifies the region where the DDS instance is located.
	// Changing this parameter will create a new role.
	Region pulumi.StringPtrInput
	// Specifies the list of roles owned by the current role.
	// The roles structure is documented below.
	// Changing this parameter will create a new role.
	Roles DatabaseRoleRoleArrayInput
}

func (DatabaseRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseRoleArgs)(nil)).Elem()
}

type DatabaseRoleInput interface {
	pulumi.Input

	ToDatabaseRoleOutput() DatabaseRoleOutput
	ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput
}

func (*DatabaseRole) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseRole)(nil)).Elem()
}

func (i *DatabaseRole) ToDatabaseRoleOutput() DatabaseRoleOutput {
	return i.ToDatabaseRoleOutputWithContext(context.Background())
}

func (i *DatabaseRole) ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleOutput)
}

// DatabaseRoleArrayInput is an input type that accepts DatabaseRoleArray and DatabaseRoleArrayOutput values.
// You can construct a concrete instance of `DatabaseRoleArrayInput` via:
//
//	DatabaseRoleArray{ DatabaseRoleArgs{...} }
type DatabaseRoleArrayInput interface {
	pulumi.Input

	ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput
	ToDatabaseRoleArrayOutputWithContext(context.Context) DatabaseRoleArrayOutput
}

type DatabaseRoleArray []DatabaseRoleInput

func (DatabaseRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseRole)(nil)).Elem()
}

func (i DatabaseRoleArray) ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput {
	return i.ToDatabaseRoleArrayOutputWithContext(context.Background())
}

func (i DatabaseRoleArray) ToDatabaseRoleArrayOutputWithContext(ctx context.Context) DatabaseRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleArrayOutput)
}

// DatabaseRoleMapInput is an input type that accepts DatabaseRoleMap and DatabaseRoleMapOutput values.
// You can construct a concrete instance of `DatabaseRoleMapInput` via:
//
//	DatabaseRoleMap{ "key": DatabaseRoleArgs{...} }
type DatabaseRoleMapInput interface {
	pulumi.Input

	ToDatabaseRoleMapOutput() DatabaseRoleMapOutput
	ToDatabaseRoleMapOutputWithContext(context.Context) DatabaseRoleMapOutput
}

type DatabaseRoleMap map[string]DatabaseRoleInput

func (DatabaseRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseRole)(nil)).Elem()
}

func (i DatabaseRoleMap) ToDatabaseRoleMapOutput() DatabaseRoleMapOutput {
	return i.ToDatabaseRoleMapOutputWithContext(context.Background())
}

func (i DatabaseRoleMap) ToDatabaseRoleMapOutputWithContext(ctx context.Context) DatabaseRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleMapOutput)
}

type DatabaseRoleOutput struct{ *pulumi.OutputState }

func (DatabaseRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleOutput) ToDatabaseRoleOutput() DatabaseRoleOutput {
	return o
}

func (o DatabaseRoleOutput) ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput {
	return o
}

// Specifies the database name to which this owned role belongs.
// Changing this parameter will create a new role.
func (o DatabaseRoleOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// The list of database privileges owned by the current role, includes all privileges
// inherited by owned roles. The inheritedPrivileges structure is documented below.
func (o DatabaseRoleOutput) InheritedPrivileges() DatabaseRoleInheritedPrivilegeArrayOutput {
	return o.ApplyT(func(v *DatabaseRole) DatabaseRoleInheritedPrivilegeArrayOutput { return v.InheritedPrivileges }).(DatabaseRoleInheritedPrivilegeArrayOutput)
}

// Specifies the DDS instance ID to which the role belongs.
// Changing this parameter will create a new role.
func (o DatabaseRoleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the name of role owned by the current role.
// The name can contain `1` to `64` characters, only letters, digits, underscores (_), hyphens (-) and dots (.) are
// allowed. Changing this parameter will create a new role.
func (o DatabaseRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of database privileges owned by the current role.
// The privileges structure is documented below.
func (o DatabaseRoleOutput) Privileges() DatabaseRolePrivilegeArrayOutput {
	return o.ApplyT(func(v *DatabaseRole) DatabaseRolePrivilegeArrayOutput { return v.Privileges }).(DatabaseRolePrivilegeArrayOutput)
}

// Specifies the region where the DDS instance is located.
// Changing this parameter will create a new role.
func (o DatabaseRoleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the list of roles owned by the current role.
// The roles structure is documented below.
// Changing this parameter will create a new role.
func (o DatabaseRoleOutput) Roles() DatabaseRoleRoleArrayOutput {
	return o.ApplyT(func(v *DatabaseRole) DatabaseRoleRoleArrayOutput { return v.Roles }).(DatabaseRoleRoleArrayOutput)
}

type DatabaseRoleArrayOutput struct{ *pulumi.OutputState }

func (DatabaseRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleArrayOutput) ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput {
	return o
}

func (o DatabaseRoleArrayOutput) ToDatabaseRoleArrayOutputWithContext(ctx context.Context) DatabaseRoleArrayOutput {
	return o
}

func (o DatabaseRoleArrayOutput) Index(i pulumi.IntInput) DatabaseRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseRole {
		return vs[0].([]*DatabaseRole)[vs[1].(int)]
	}).(DatabaseRoleOutput)
}

type DatabaseRoleMapOutput struct{ *pulumi.OutputState }

func (DatabaseRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleMapOutput) ToDatabaseRoleMapOutput() DatabaseRoleMapOutput {
	return o
}

func (o DatabaseRoleMapOutput) ToDatabaseRoleMapOutputWithContext(ctx context.Context) DatabaseRoleMapOutput {
	return o
}

func (o DatabaseRoleMapOutput) MapIndex(k pulumi.StringInput) DatabaseRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseRole {
		return vs[0].(map[string]*DatabaseRole)[vs[1].(string)]
	}).(DatabaseRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleInput)(nil)).Elem(), &DatabaseRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleArrayInput)(nil)).Elem(), DatabaseRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleMapInput)(nil)).Elem(), DatabaseRoleMap{})
	pulumi.RegisterOutputType(DatabaseRoleOutput{})
	pulumi.RegisterOutputType(DatabaseRoleArrayOutput{})
	pulumi.RegisterOutputType(DatabaseRoleMapOutput{})
}
