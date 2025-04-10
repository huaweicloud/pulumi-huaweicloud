// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages RDS SQLServer database resource within HuaweiCloud.
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
//			_, err := Rds.NewSqlserverDatabase(ctx, "test", &Rds.SqlserverDatabaseArgs{
//				InstanceId: pulumi.Any(instanceId),
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
// The RDS sqlserver database can be imported using the `instance_id` and `name` separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/sqlserverDatabase:SqlserverDatabase test <instance_id>/<name>
//
// ```
type SqlserverDatabase struct {
	pulumi.CustomResourceState

	// Indicates the character set used by the database.
	CharacterSet pulumi.StringOutput `pulumi:"characterSet"`
	// Specifies the ID of the RDS SQLServer instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
	// **tempdb**, **resource**, and **rdsadmin**.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Indicates the database status. Its value can be any of the following:
	// + **Creating**: The database is being created.
	// + **Running**: The database is running.
	// + **Deleting**: The database is being deleted.
	// + **Not Exists**: The database does not exist.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewSqlserverDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlserverDatabase(ctx *pulumi.Context,
	name string, args *SqlserverDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlserverDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SqlserverDatabase
	err := ctx.RegisterResource("huaweicloud:Rds/sqlserverDatabase:SqlserverDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlserverDatabase gets an existing SqlserverDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlserverDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlserverDatabaseState, opts ...pulumi.ResourceOption) (*SqlserverDatabase, error) {
	var resource SqlserverDatabase
	err := ctx.ReadResource("huaweicloud:Rds/sqlserverDatabase:SqlserverDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlserverDatabase resources.
type sqlserverDatabaseState struct {
	// Indicates the character set used by the database.
	CharacterSet *string `pulumi:"characterSet"`
	// Specifies the ID of the RDS SQLServer instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
	// **tempdb**, **resource**, and **rdsadmin**.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Indicates the database status. Its value can be any of the following:
	// + **Creating**: The database is being created.
	// + **Running**: The database is running.
	// + **Deleting**: The database is being deleted.
	// + **Not Exists**: The database does not exist.
	State *string `pulumi:"state"`
}

type SqlserverDatabaseState struct {
	// Indicates the character set used by the database.
	CharacterSet pulumi.StringPtrInput
	// Specifies the ID of the RDS SQLServer instance.
	InstanceId pulumi.StringPtrInput
	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
	// **tempdb**, **resource**, and **rdsadmin**.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Indicates the database status. Its value can be any of the following:
	// + **Creating**: The database is being created.
	// + **Running**: The database is running.
	// + **Deleting**: The database is being deleted.
	// + **Not Exists**: The database does not exist.
	State pulumi.StringPtrInput
}

func (SqlserverDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlserverDatabaseState)(nil)).Elem()
}

type sqlserverDatabaseArgs struct {
	// Specifies the ID of the RDS SQLServer instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
	// **tempdb**, **resource**, and **rdsadmin**.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a SqlserverDatabase resource.
type SqlserverDatabaseArgs struct {
	// Specifies the ID of the RDS SQLServer instance.
	InstanceId pulumi.StringInput
	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
	// **tempdb**, **resource**, and **rdsadmin**.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (SqlserverDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlserverDatabaseArgs)(nil)).Elem()
}

type SqlserverDatabaseInput interface {
	pulumi.Input

	ToSqlserverDatabaseOutput() SqlserverDatabaseOutput
	ToSqlserverDatabaseOutputWithContext(ctx context.Context) SqlserverDatabaseOutput
}

func (*SqlserverDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlserverDatabase)(nil)).Elem()
}

func (i *SqlserverDatabase) ToSqlserverDatabaseOutput() SqlserverDatabaseOutput {
	return i.ToSqlserverDatabaseOutputWithContext(context.Background())
}

func (i *SqlserverDatabase) ToSqlserverDatabaseOutputWithContext(ctx context.Context) SqlserverDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlserverDatabaseOutput)
}

// SqlserverDatabaseArrayInput is an input type that accepts SqlserverDatabaseArray and SqlserverDatabaseArrayOutput values.
// You can construct a concrete instance of `SqlserverDatabaseArrayInput` via:
//
//	SqlserverDatabaseArray{ SqlserverDatabaseArgs{...} }
type SqlserverDatabaseArrayInput interface {
	pulumi.Input

	ToSqlserverDatabaseArrayOutput() SqlserverDatabaseArrayOutput
	ToSqlserverDatabaseArrayOutputWithContext(context.Context) SqlserverDatabaseArrayOutput
}

type SqlserverDatabaseArray []SqlserverDatabaseInput

func (SqlserverDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlserverDatabase)(nil)).Elem()
}

func (i SqlserverDatabaseArray) ToSqlserverDatabaseArrayOutput() SqlserverDatabaseArrayOutput {
	return i.ToSqlserverDatabaseArrayOutputWithContext(context.Background())
}

func (i SqlserverDatabaseArray) ToSqlserverDatabaseArrayOutputWithContext(ctx context.Context) SqlserverDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlserverDatabaseArrayOutput)
}

// SqlserverDatabaseMapInput is an input type that accepts SqlserverDatabaseMap and SqlserverDatabaseMapOutput values.
// You can construct a concrete instance of `SqlserverDatabaseMapInput` via:
//
//	SqlserverDatabaseMap{ "key": SqlserverDatabaseArgs{...} }
type SqlserverDatabaseMapInput interface {
	pulumi.Input

	ToSqlserverDatabaseMapOutput() SqlserverDatabaseMapOutput
	ToSqlserverDatabaseMapOutputWithContext(context.Context) SqlserverDatabaseMapOutput
}

type SqlserverDatabaseMap map[string]SqlserverDatabaseInput

func (SqlserverDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlserverDatabase)(nil)).Elem()
}

func (i SqlserverDatabaseMap) ToSqlserverDatabaseMapOutput() SqlserverDatabaseMapOutput {
	return i.ToSqlserverDatabaseMapOutputWithContext(context.Background())
}

func (i SqlserverDatabaseMap) ToSqlserverDatabaseMapOutputWithContext(ctx context.Context) SqlserverDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlserverDatabaseMapOutput)
}

type SqlserverDatabaseOutput struct{ *pulumi.OutputState }

func (SqlserverDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlserverDatabase)(nil)).Elem()
}

func (o SqlserverDatabaseOutput) ToSqlserverDatabaseOutput() SqlserverDatabaseOutput {
	return o
}

func (o SqlserverDatabaseOutput) ToSqlserverDatabaseOutputWithContext(ctx context.Context) SqlserverDatabaseOutput {
	return o
}

// Indicates the character set used by the database.
func (o SqlserverDatabaseOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlserverDatabase) pulumi.StringOutput { return v.CharacterSet }).(pulumi.StringOutput)
}

// Specifies the ID of the RDS SQLServer instance.
func (o SqlserverDatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlserverDatabase) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the database name. The database name can contain 1 to 64 characters,
// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
// SQL Server system database name. RDS for SQL Server system databases include **master**, **msdb**, **model**,
// **tempdb**, **resource**, and **rdsadmin**.
func (o SqlserverDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlserverDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o SqlserverDatabaseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlserverDatabase) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Indicates the database status. Its value can be any of the following:
// + **Creating**: The database is being created.
// + **Running**: The database is running.
// + **Deleting**: The database is being deleted.
// + **Not Exists**: The database does not exist.
func (o SqlserverDatabaseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlserverDatabase) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type SqlserverDatabaseArrayOutput struct{ *pulumi.OutputState }

func (SqlserverDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlserverDatabase)(nil)).Elem()
}

func (o SqlserverDatabaseArrayOutput) ToSqlserverDatabaseArrayOutput() SqlserverDatabaseArrayOutput {
	return o
}

func (o SqlserverDatabaseArrayOutput) ToSqlserverDatabaseArrayOutputWithContext(ctx context.Context) SqlserverDatabaseArrayOutput {
	return o
}

func (o SqlserverDatabaseArrayOutput) Index(i pulumi.IntInput) SqlserverDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlserverDatabase {
		return vs[0].([]*SqlserverDatabase)[vs[1].(int)]
	}).(SqlserverDatabaseOutput)
}

type SqlserverDatabaseMapOutput struct{ *pulumi.OutputState }

func (SqlserverDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlserverDatabase)(nil)).Elem()
}

func (o SqlserverDatabaseMapOutput) ToSqlserverDatabaseMapOutput() SqlserverDatabaseMapOutput {
	return o
}

func (o SqlserverDatabaseMapOutput) ToSqlserverDatabaseMapOutputWithContext(ctx context.Context) SqlserverDatabaseMapOutput {
	return o
}

func (o SqlserverDatabaseMapOutput) MapIndex(k pulumi.StringInput) SqlserverDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlserverDatabase {
		return vs[0].(map[string]*SqlserverDatabase)[vs[1].(string)]
	}).(SqlserverDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlserverDatabaseInput)(nil)).Elem(), &SqlserverDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlserverDatabaseArrayInput)(nil)).Elem(), SqlserverDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlserverDatabaseMapInput)(nil)).Elem(), SqlserverDatabaseMap{})
	pulumi.RegisterOutputType(SqlserverDatabaseOutput{})
	pulumi.RegisterOutputType(SqlserverDatabaseArrayOutput{})
	pulumi.RegisterOutputType(SqlserverDatabaseMapOutput{})
}
