// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages RDS SQL audit resource within HuaweiCloud.
//
// > **NOTE:** Only MySQL and PostgreSQL engines are supported.
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
//			_, err := Rds.NewSqlAudit(ctx, "test", &Rds.SqlAuditArgs{
//				InstanceId: pulumi.Any(instanceId),
//				KeepDays:   pulumi.Int(5),
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
// The RDS SQL audit can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/sqlAudit:SqlAudit test <id>
//
// ```
type SqlAudit struct {
	pulumi.CustomResourceState

	// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
	// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
	// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
	// It is not supported for PostgreSQL.
	AuditTypes pulumi.StringArrayOutput `pulumi:"auditTypes"`
	// Specifies the ID of the RDS instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
	KeepDays pulumi.IntOutput `pulumi:"keepDays"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies whether the historical audit logs will be reserved for some time
	// when SQL audit is disabled. It is valid only when SQL audit is disabled.
	ReserveAuditlogs pulumi.BoolPtrOutput `pulumi:"reserveAuditlogs"`
}

// NewSqlAudit registers a new resource with the given unique name, arguments, and options.
func NewSqlAudit(ctx *pulumi.Context,
	name string, args *SqlAuditArgs, opts ...pulumi.ResourceOption) (*SqlAudit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.KeepDays == nil {
		return nil, errors.New("invalid value for required argument 'KeepDays'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SqlAudit
	err := ctx.RegisterResource("huaweicloud:Rds/sqlAudit:SqlAudit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlAudit gets an existing SqlAudit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlAudit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlAuditState, opts ...pulumi.ResourceOption) (*SqlAudit, error) {
	var resource SqlAudit
	err := ctx.ReadResource("huaweicloud:Rds/sqlAudit:SqlAudit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlAudit resources.
type sqlAuditState struct {
	// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
	// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
	// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
	// It is not supported for PostgreSQL.
	AuditTypes []string `pulumi:"auditTypes"`
	// Specifies the ID of the RDS instance.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
	KeepDays *int `pulumi:"keepDays"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies whether the historical audit logs will be reserved for some time
	// when SQL audit is disabled. It is valid only when SQL audit is disabled.
	ReserveAuditlogs *bool `pulumi:"reserveAuditlogs"`
}

type SqlAuditState struct {
	// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
	// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
	// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
	// It is not supported for PostgreSQL.
	AuditTypes pulumi.StringArrayInput
	// Specifies the ID of the RDS instance.
	InstanceId pulumi.StringPtrInput
	// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
	KeepDays pulumi.IntPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies whether the historical audit logs will be reserved for some time
	// when SQL audit is disabled. It is valid only when SQL audit is disabled.
	ReserveAuditlogs pulumi.BoolPtrInput
}

func (SqlAuditState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlAuditState)(nil)).Elem()
}

type sqlAuditArgs struct {
	// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
	// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
	// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
	// It is not supported for PostgreSQL.
	AuditTypes []string `pulumi:"auditTypes"`
	// Specifies the ID of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
	KeepDays int `pulumi:"keepDays"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies whether the historical audit logs will be reserved for some time
	// when SQL audit is disabled. It is valid only when SQL audit is disabled.
	ReserveAuditlogs *bool `pulumi:"reserveAuditlogs"`
}

// The set of arguments for constructing a SqlAudit resource.
type SqlAuditArgs struct {
	// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
	// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
	// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
	// It is not supported for PostgreSQL.
	AuditTypes pulumi.StringArrayInput
	// Specifies the ID of the RDS instance.
	InstanceId pulumi.StringInput
	// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
	KeepDays pulumi.IntInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies whether the historical audit logs will be reserved for some time
	// when SQL audit is disabled. It is valid only when SQL audit is disabled.
	ReserveAuditlogs pulumi.BoolPtrInput
}

func (SqlAuditArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlAuditArgs)(nil)).Elem()
}

type SqlAuditInput interface {
	pulumi.Input

	ToSqlAuditOutput() SqlAuditOutput
	ToSqlAuditOutputWithContext(ctx context.Context) SqlAuditOutput
}

func (*SqlAudit) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlAudit)(nil)).Elem()
}

func (i *SqlAudit) ToSqlAuditOutput() SqlAuditOutput {
	return i.ToSqlAuditOutputWithContext(context.Background())
}

func (i *SqlAudit) ToSqlAuditOutputWithContext(ctx context.Context) SqlAuditOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlAuditOutput)
}

// SqlAuditArrayInput is an input type that accepts SqlAuditArray and SqlAuditArrayOutput values.
// You can construct a concrete instance of `SqlAuditArrayInput` via:
//
//	SqlAuditArray{ SqlAuditArgs{...} }
type SqlAuditArrayInput interface {
	pulumi.Input

	ToSqlAuditArrayOutput() SqlAuditArrayOutput
	ToSqlAuditArrayOutputWithContext(context.Context) SqlAuditArrayOutput
}

type SqlAuditArray []SqlAuditInput

func (SqlAuditArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlAudit)(nil)).Elem()
}

func (i SqlAuditArray) ToSqlAuditArrayOutput() SqlAuditArrayOutput {
	return i.ToSqlAuditArrayOutputWithContext(context.Background())
}

func (i SqlAuditArray) ToSqlAuditArrayOutputWithContext(ctx context.Context) SqlAuditArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlAuditArrayOutput)
}

// SqlAuditMapInput is an input type that accepts SqlAuditMap and SqlAuditMapOutput values.
// You can construct a concrete instance of `SqlAuditMapInput` via:
//
//	SqlAuditMap{ "key": SqlAuditArgs{...} }
type SqlAuditMapInput interface {
	pulumi.Input

	ToSqlAuditMapOutput() SqlAuditMapOutput
	ToSqlAuditMapOutputWithContext(context.Context) SqlAuditMapOutput
}

type SqlAuditMap map[string]SqlAuditInput

func (SqlAuditMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlAudit)(nil)).Elem()
}

func (i SqlAuditMap) ToSqlAuditMapOutput() SqlAuditMapOutput {
	return i.ToSqlAuditMapOutputWithContext(context.Background())
}

func (i SqlAuditMap) ToSqlAuditMapOutputWithContext(ctx context.Context) SqlAuditMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlAuditMapOutput)
}

type SqlAuditOutput struct{ *pulumi.OutputState }

func (SqlAuditOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlAudit)(nil)).Elem()
}

func (o SqlAuditOutput) ToSqlAuditOutput() SqlAuditOutput {
	return o
}

func (o SqlAuditOutput) ToSqlAuditOutputWithContext(ctx context.Context) SqlAuditOutput {
	return o
}

// Specifies the list of audit types. Value options: **CREATE_USER**, **DROP_USER**,
// **RENAME_USER**, **GRANT**, **REVOKE**, **CREATE**, **ALTER**, **DROP**, **RENAME**, **TRUNCATE**, **INSERT**,
// **DELETE**, **UPDATE**, **REPLACE**, **SELECT**, **BEGIN/COMMIT/ROLLBACK**, **PREPARED_STATEMENT**.
// It is not supported for PostgreSQL.
func (o SqlAuditOutput) AuditTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SqlAudit) pulumi.StringArrayOutput { return v.AuditTypes }).(pulumi.StringArrayOutput)
}

// Specifies the ID of the RDS instance.
func (o SqlAuditOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlAudit) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the number of days for storing audit logs. Value ranges from `1` to `732`.
func (o SqlAuditOutput) KeepDays() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlAudit) pulumi.IntOutput { return v.KeepDays }).(pulumi.IntOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o SqlAuditOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlAudit) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies whether the historical audit logs will be reserved for some time
// when SQL audit is disabled. It is valid only when SQL audit is disabled.
func (o SqlAuditOutput) ReserveAuditlogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlAudit) pulumi.BoolPtrOutput { return v.ReserveAuditlogs }).(pulumi.BoolPtrOutput)
}

type SqlAuditArrayOutput struct{ *pulumi.OutputState }

func (SqlAuditArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlAudit)(nil)).Elem()
}

func (o SqlAuditArrayOutput) ToSqlAuditArrayOutput() SqlAuditArrayOutput {
	return o
}

func (o SqlAuditArrayOutput) ToSqlAuditArrayOutputWithContext(ctx context.Context) SqlAuditArrayOutput {
	return o
}

func (o SqlAuditArrayOutput) Index(i pulumi.IntInput) SqlAuditOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlAudit {
		return vs[0].([]*SqlAudit)[vs[1].(int)]
	}).(SqlAuditOutput)
}

type SqlAuditMapOutput struct{ *pulumi.OutputState }

func (SqlAuditMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlAudit)(nil)).Elem()
}

func (o SqlAuditMapOutput) ToSqlAuditMapOutput() SqlAuditMapOutput {
	return o
}

func (o SqlAuditMapOutput) ToSqlAuditMapOutputWithContext(ctx context.Context) SqlAuditMapOutput {
	return o
}

func (o SqlAuditMapOutput) MapIndex(k pulumi.StringInput) SqlAuditOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlAudit {
		return vs[0].(map[string]*SqlAudit)[vs[1].(string)]
	}).(SqlAuditOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAuditInput)(nil)).Elem(), &SqlAudit{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAuditArrayInput)(nil)).Elem(), SqlAuditArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAuditMapInput)(nil)).Elem(), SqlAuditMap{})
	pulumi.RegisterOutputType(SqlAuditOutput{})
	pulumi.RegisterOutputType(SqlAuditArrayOutput{})
	pulumi.RegisterOutputType(SqlAuditMapOutput{})
}
