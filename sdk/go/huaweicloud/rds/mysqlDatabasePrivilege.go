// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages RDS Mysql database privilege resource within HuaweiCloud.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			instanceId := cfg.RequireObject("instanceId")
//			dbName := cfg.RequireObject("dbName")
//			userName1 := cfg.RequireObject("userName1")
//			userName2 := cfg.RequireObject("userName2")
//			_, err := Rds.NewMysqlDatabasePrivilege(ctx, "test", &Rds.MysqlDatabasePrivilegeArgs{
//				InstanceId: pulumi.Any(instanceId),
//				DbName:     pulumi.Any(dbName),
//				Users: rds.MysqlDatabasePrivilegeUserArray{
//					&rds.MysqlDatabasePrivilegeUserArgs{
//						Name:     pulumi.Any(userName1),
//						Readonly: pulumi.Bool(true),
//					},
//					&rds.MysqlDatabasePrivilegeUserArgs{
//						Name:     pulumi.Any(userName2),
//						Readonly: pulumi.Bool(false),
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
// RDS database privilege can be imported using the `instance id` and `db_name`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/mysqlDatabasePrivilege:MysqlDatabasePrivilege test <instance_id>/<db_name>
//
// ```
type MysqlDatabasePrivilege struct {
	pulumi.CustomResourceState

	// Specifies the database name. Changing this creates a new resource.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Specifies the RDS instance ID. Changing this will create a new resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the account that associated with the database. Structure is documented below.
	Users MysqlDatabasePrivilegeUserArrayOutput `pulumi:"users"`
}

// NewMysqlDatabasePrivilege registers a new resource with the given unique name, arguments, and options.
func NewMysqlDatabasePrivilege(ctx *pulumi.Context,
	name string, args *MysqlDatabasePrivilegeArgs, opts ...pulumi.ResourceOption) (*MysqlDatabasePrivilege, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MysqlDatabasePrivilege
	err := ctx.RegisterResource("huaweicloud:Rds/mysqlDatabasePrivilege:MysqlDatabasePrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMysqlDatabasePrivilege gets an existing MysqlDatabasePrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMysqlDatabasePrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MysqlDatabasePrivilegeState, opts ...pulumi.ResourceOption) (*MysqlDatabasePrivilege, error) {
	var resource MysqlDatabasePrivilege
	err := ctx.ReadResource("huaweicloud:Rds/mysqlDatabasePrivilege:MysqlDatabasePrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MysqlDatabasePrivilege resources.
type mysqlDatabasePrivilegeState struct {
	// Specifies the database name. Changing this creates a new resource.
	DbName *string `pulumi:"dbName"`
	// Specifies the RDS instance ID. Changing this will create a new resource.
	InstanceId *string `pulumi:"instanceId"`
	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the account that associated with the database. Structure is documented below.
	Users []MysqlDatabasePrivilegeUser `pulumi:"users"`
}

type MysqlDatabasePrivilegeState struct {
	// Specifies the database name. Changing this creates a new resource.
	DbName pulumi.StringPtrInput
	// Specifies the RDS instance ID. Changing this will create a new resource.
	InstanceId pulumi.StringPtrInput
	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the account that associated with the database. Structure is documented below.
	Users MysqlDatabasePrivilegeUserArrayInput
}

func (MysqlDatabasePrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlDatabasePrivilegeState)(nil)).Elem()
}

type mysqlDatabasePrivilegeArgs struct {
	// Specifies the database name. Changing this creates a new resource.
	DbName string `pulumi:"dbName"`
	// Specifies the RDS instance ID. Changing this will create a new resource.
	InstanceId string `pulumi:"instanceId"`
	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the account that associated with the database. Structure is documented below.
	Users []MysqlDatabasePrivilegeUser `pulumi:"users"`
}

// The set of arguments for constructing a MysqlDatabasePrivilege resource.
type MysqlDatabasePrivilegeArgs struct {
	// Specifies the database name. Changing this creates a new resource.
	DbName pulumi.StringInput
	// Specifies the RDS instance ID. Changing this will create a new resource.
	InstanceId pulumi.StringInput
	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the account that associated with the database. Structure is documented below.
	Users MysqlDatabasePrivilegeUserArrayInput
}

func (MysqlDatabasePrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlDatabasePrivilegeArgs)(nil)).Elem()
}

type MysqlDatabasePrivilegeInput interface {
	pulumi.Input

	ToMysqlDatabasePrivilegeOutput() MysqlDatabasePrivilegeOutput
	ToMysqlDatabasePrivilegeOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeOutput
}

func (*MysqlDatabasePrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlDatabasePrivilege)(nil)).Elem()
}

func (i *MysqlDatabasePrivilege) ToMysqlDatabasePrivilegeOutput() MysqlDatabasePrivilegeOutput {
	return i.ToMysqlDatabasePrivilegeOutputWithContext(context.Background())
}

func (i *MysqlDatabasePrivilege) ToMysqlDatabasePrivilegeOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlDatabasePrivilegeOutput)
}

// MysqlDatabasePrivilegeArrayInput is an input type that accepts MysqlDatabasePrivilegeArray and MysqlDatabasePrivilegeArrayOutput values.
// You can construct a concrete instance of `MysqlDatabasePrivilegeArrayInput` via:
//
//	MysqlDatabasePrivilegeArray{ MysqlDatabasePrivilegeArgs{...} }
type MysqlDatabasePrivilegeArrayInput interface {
	pulumi.Input

	ToMysqlDatabasePrivilegeArrayOutput() MysqlDatabasePrivilegeArrayOutput
	ToMysqlDatabasePrivilegeArrayOutputWithContext(context.Context) MysqlDatabasePrivilegeArrayOutput
}

type MysqlDatabasePrivilegeArray []MysqlDatabasePrivilegeInput

func (MysqlDatabasePrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MysqlDatabasePrivilege)(nil)).Elem()
}

func (i MysqlDatabasePrivilegeArray) ToMysqlDatabasePrivilegeArrayOutput() MysqlDatabasePrivilegeArrayOutput {
	return i.ToMysqlDatabasePrivilegeArrayOutputWithContext(context.Background())
}

func (i MysqlDatabasePrivilegeArray) ToMysqlDatabasePrivilegeArrayOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlDatabasePrivilegeArrayOutput)
}

// MysqlDatabasePrivilegeMapInput is an input type that accepts MysqlDatabasePrivilegeMap and MysqlDatabasePrivilegeMapOutput values.
// You can construct a concrete instance of `MysqlDatabasePrivilegeMapInput` via:
//
//	MysqlDatabasePrivilegeMap{ "key": MysqlDatabasePrivilegeArgs{...} }
type MysqlDatabasePrivilegeMapInput interface {
	pulumi.Input

	ToMysqlDatabasePrivilegeMapOutput() MysqlDatabasePrivilegeMapOutput
	ToMysqlDatabasePrivilegeMapOutputWithContext(context.Context) MysqlDatabasePrivilegeMapOutput
}

type MysqlDatabasePrivilegeMap map[string]MysqlDatabasePrivilegeInput

func (MysqlDatabasePrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MysqlDatabasePrivilege)(nil)).Elem()
}

func (i MysqlDatabasePrivilegeMap) ToMysqlDatabasePrivilegeMapOutput() MysqlDatabasePrivilegeMapOutput {
	return i.ToMysqlDatabasePrivilegeMapOutputWithContext(context.Background())
}

func (i MysqlDatabasePrivilegeMap) ToMysqlDatabasePrivilegeMapOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlDatabasePrivilegeMapOutput)
}

type MysqlDatabasePrivilegeOutput struct{ *pulumi.OutputState }

func (MysqlDatabasePrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlDatabasePrivilege)(nil)).Elem()
}

func (o MysqlDatabasePrivilegeOutput) ToMysqlDatabasePrivilegeOutput() MysqlDatabasePrivilegeOutput {
	return o
}

func (o MysqlDatabasePrivilegeOutput) ToMysqlDatabasePrivilegeOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeOutput {
	return o
}

// Specifies the database name. Changing this creates a new resource.
func (o MysqlDatabasePrivilegeOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *MysqlDatabasePrivilege) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// Specifies the RDS instance ID. Changing this will create a new resource.
func (o MysqlDatabasePrivilegeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MysqlDatabasePrivilege) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The region in which to create the RDS database privilege resource. If omitted,
// the provider-level region will be used. Changing this creates a new resource.
func (o MysqlDatabasePrivilegeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MysqlDatabasePrivilege) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the account that associated with the database. Structure is documented below.
func (o MysqlDatabasePrivilegeOutput) Users() MysqlDatabasePrivilegeUserArrayOutput {
	return o.ApplyT(func(v *MysqlDatabasePrivilege) MysqlDatabasePrivilegeUserArrayOutput { return v.Users }).(MysqlDatabasePrivilegeUserArrayOutput)
}

type MysqlDatabasePrivilegeArrayOutput struct{ *pulumi.OutputState }

func (MysqlDatabasePrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MysqlDatabasePrivilege)(nil)).Elem()
}

func (o MysqlDatabasePrivilegeArrayOutput) ToMysqlDatabasePrivilegeArrayOutput() MysqlDatabasePrivilegeArrayOutput {
	return o
}

func (o MysqlDatabasePrivilegeArrayOutput) ToMysqlDatabasePrivilegeArrayOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeArrayOutput {
	return o
}

func (o MysqlDatabasePrivilegeArrayOutput) Index(i pulumi.IntInput) MysqlDatabasePrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MysqlDatabasePrivilege {
		return vs[0].([]*MysqlDatabasePrivilege)[vs[1].(int)]
	}).(MysqlDatabasePrivilegeOutput)
}

type MysqlDatabasePrivilegeMapOutput struct{ *pulumi.OutputState }

func (MysqlDatabasePrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MysqlDatabasePrivilege)(nil)).Elem()
}

func (o MysqlDatabasePrivilegeMapOutput) ToMysqlDatabasePrivilegeMapOutput() MysqlDatabasePrivilegeMapOutput {
	return o
}

func (o MysqlDatabasePrivilegeMapOutput) ToMysqlDatabasePrivilegeMapOutputWithContext(ctx context.Context) MysqlDatabasePrivilegeMapOutput {
	return o
}

func (o MysqlDatabasePrivilegeMapOutput) MapIndex(k pulumi.StringInput) MysqlDatabasePrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MysqlDatabasePrivilege {
		return vs[0].(map[string]*MysqlDatabasePrivilege)[vs[1].(string)]
	}).(MysqlDatabasePrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlDatabasePrivilegeInput)(nil)).Elem(), &MysqlDatabasePrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlDatabasePrivilegeArrayInput)(nil)).Elem(), MysqlDatabasePrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlDatabasePrivilegeMapInput)(nil)).Elem(), MysqlDatabasePrivilegeMap{})
	pulumi.RegisterOutputType(MysqlDatabasePrivilegeOutput{})
	pulumi.RegisterOutputType(MysqlDatabasePrivilegeArrayOutput{})
	pulumi.RegisterOutputType(MysqlDatabasePrivilegeMapOutput{})
}
