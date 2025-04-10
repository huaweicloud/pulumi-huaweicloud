// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dli

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages DLI SQL job resource within HuaweiCloud
//
// ## Example Usage
// ### Create a Sql job
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dli"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			databaseName := cfg.RequireObject("databaseName")
//			queueName := cfg.RequireObject("queueName")
//			sql := cfg.RequireObject("sql")
//			_, err := Dli.NewSqlJob(ctx, "test", &Dli.SqlJobArgs{
//				Sql:          pulumi.Any(sql),
//				DatabaseName: pulumi.Any(databaseName),
//				QueueName:    pulumi.Any(queueName),
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
// DLI SQL job can be imported by `id`. For example, bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Dli/sqlJob:SqlJob example 7f803d70-c533-469f-8431-e378f3e97123
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`conf`, `rows` and `schema`. It is generally recommended running `terraform plan` after importing a resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the resource. Also you can ignore changes as below. hcl resource "huaweicloud_dli_sql_job" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	conf, rows, schema
//
//	]
//
//	} }
type SqlJob struct {
	pulumi.CustomResourceState

	// Specifies the configuration parameters for the SQL job. Changing this parameter
	// will create a new resource. Structure is documented below.
	Conf SqlJobConfPtrOutput `pulumi:"conf"`
	// Specifies the database where the SQL is executed. This argument does
	// not need to be configured during database creation. Changing this parameter will create a new resource.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// Job running duration (unit: millisecond).
	Duration pulumi.IntOutput `pulumi:"duration"`
	// Type of a job, Includes **DDL**, **DCL**, **IMPORT**, **EXPORT**, **QUERY**, **INSERT**,
	// **DATA_MIGRATION**, **UPDATE**, **DELETE**, **RESTART_QUEUE** and **SCALE_QUEUE**.
	JobType pulumi.StringOutput `pulumi:"jobType"`
	// User who submits a job.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Specifies queue which this job to be submitted belongs.
	// Changing this parameter will create a new resource.
	QueueName pulumi.StringOutput `pulumi:"queueName"`
	// Specifies the region in which to create the DLI table resource. If omitted,
	// the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// When the statement type is DDL, results of the DDL are displayed.
	Rows pulumi.StringArrayArrayOutput `pulumi:"rows"`
	// When the statement type is DDL, the column name and type of DDL are displayed.
	Schemas pulumi.StringMapArrayOutput `pulumi:"schemas"`
	// Specifies SQL statement that you want to execute.
	// Changing this parameter will create a new resource.
	Sql pulumi.StringOutput `pulumi:"sql"`
	// Time when a job is started, in RFC-3339 format. e.g. `2019-10-12T07:20:50.52Z`
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Status of a job, including **RUNNING**, **SCALING**, **LAUNCHING**, **FINISHED**, **FAILED**,
	// and **CANCELLED.**
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies label of a Job. Changing this parameter will create a new resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSqlJob registers a new resource with the given unique name, arguments, and options.
func NewSqlJob(ctx *pulumi.Context,
	name string, args *SqlJobArgs, opts ...pulumi.ResourceOption) (*SqlJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sql == nil {
		return nil, errors.New("invalid value for required argument 'Sql'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SqlJob
	err := ctx.RegisterResource("huaweicloud:Dli/sqlJob:SqlJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlJob gets an existing SqlJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlJobState, opts ...pulumi.ResourceOption) (*SqlJob, error) {
	var resource SqlJob
	err := ctx.ReadResource("huaweicloud:Dli/sqlJob:SqlJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlJob resources.
type sqlJobState struct {
	// Specifies the configuration parameters for the SQL job. Changing this parameter
	// will create a new resource. Structure is documented below.
	Conf *SqlJobConf `pulumi:"conf"`
	// Specifies the database where the SQL is executed. This argument does
	// not need to be configured during database creation. Changing this parameter will create a new resource.
	DatabaseName *string `pulumi:"databaseName"`
	// Job running duration (unit: millisecond).
	Duration *int `pulumi:"duration"`
	// Type of a job, Includes **DDL**, **DCL**, **IMPORT**, **EXPORT**, **QUERY**, **INSERT**,
	// **DATA_MIGRATION**, **UPDATE**, **DELETE**, **RESTART_QUEUE** and **SCALE_QUEUE**.
	JobType *string `pulumi:"jobType"`
	// User who submits a job.
	Owner *string `pulumi:"owner"`
	// Specifies queue which this job to be submitted belongs.
	// Changing this parameter will create a new resource.
	QueueName *string `pulumi:"queueName"`
	// Specifies the region in which to create the DLI table resource. If omitted,
	// the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// When the statement type is DDL, results of the DDL are displayed.
	Rows [][]string `pulumi:"rows"`
	// When the statement type is DDL, the column name and type of DDL are displayed.
	Schemas []map[string]string `pulumi:"schemas"`
	// Specifies SQL statement that you want to execute.
	// Changing this parameter will create a new resource.
	Sql *string `pulumi:"sql"`
	// Time when a job is started, in RFC-3339 format. e.g. `2019-10-12T07:20:50.52Z`
	StartTime *string `pulumi:"startTime"`
	// Status of a job, including **RUNNING**, **SCALING**, **LAUNCHING**, **FINISHED**, **FAILED**,
	// and **CANCELLED.**
	Status *string `pulumi:"status"`
	// Specifies label of a Job. Changing this parameter will create a new resource.
	Tags map[string]string `pulumi:"tags"`
}

type SqlJobState struct {
	// Specifies the configuration parameters for the SQL job. Changing this parameter
	// will create a new resource. Structure is documented below.
	Conf SqlJobConfPtrInput
	// Specifies the database where the SQL is executed. This argument does
	// not need to be configured during database creation. Changing this parameter will create a new resource.
	DatabaseName pulumi.StringPtrInput
	// Job running duration (unit: millisecond).
	Duration pulumi.IntPtrInput
	// Type of a job, Includes **DDL**, **DCL**, **IMPORT**, **EXPORT**, **QUERY**, **INSERT**,
	// **DATA_MIGRATION**, **UPDATE**, **DELETE**, **RESTART_QUEUE** and **SCALE_QUEUE**.
	JobType pulumi.StringPtrInput
	// User who submits a job.
	Owner pulumi.StringPtrInput
	// Specifies queue which this job to be submitted belongs.
	// Changing this parameter will create a new resource.
	QueueName pulumi.StringPtrInput
	// Specifies the region in which to create the DLI table resource. If omitted,
	// the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// When the statement type is DDL, results of the DDL are displayed.
	Rows pulumi.StringArrayArrayInput
	// When the statement type is DDL, the column name and type of DDL are displayed.
	Schemas pulumi.StringMapArrayInput
	// Specifies SQL statement that you want to execute.
	// Changing this parameter will create a new resource.
	Sql pulumi.StringPtrInput
	// Time when a job is started, in RFC-3339 format. e.g. `2019-10-12T07:20:50.52Z`
	StartTime pulumi.StringPtrInput
	// Status of a job, including **RUNNING**, **SCALING**, **LAUNCHING**, **FINISHED**, **FAILED**,
	// and **CANCELLED.**
	Status pulumi.StringPtrInput
	// Specifies label of a Job. Changing this parameter will create a new resource.
	Tags pulumi.StringMapInput
}

func (SqlJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlJobState)(nil)).Elem()
}

type sqlJobArgs struct {
	// Specifies the configuration parameters for the SQL job. Changing this parameter
	// will create a new resource. Structure is documented below.
	Conf *SqlJobConf `pulumi:"conf"`
	// Specifies the database where the SQL is executed. This argument does
	// not need to be configured during database creation. Changing this parameter will create a new resource.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies queue which this job to be submitted belongs.
	// Changing this parameter will create a new resource.
	QueueName *string `pulumi:"queueName"`
	// Specifies the region in which to create the DLI table resource. If omitted,
	// the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies SQL statement that you want to execute.
	// Changing this parameter will create a new resource.
	Sql string `pulumi:"sql"`
	// Specifies label of a Job. Changing this parameter will create a new resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlJob resource.
type SqlJobArgs struct {
	// Specifies the configuration parameters for the SQL job. Changing this parameter
	// will create a new resource. Structure is documented below.
	Conf SqlJobConfPtrInput
	// Specifies the database where the SQL is executed. This argument does
	// not need to be configured during database creation. Changing this parameter will create a new resource.
	DatabaseName pulumi.StringPtrInput
	// Specifies queue which this job to be submitted belongs.
	// Changing this parameter will create a new resource.
	QueueName pulumi.StringPtrInput
	// Specifies the region in which to create the DLI table resource. If omitted,
	// the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies SQL statement that you want to execute.
	// Changing this parameter will create a new resource.
	Sql pulumi.StringInput
	// Specifies label of a Job. Changing this parameter will create a new resource.
	Tags pulumi.StringMapInput
}

func (SqlJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlJobArgs)(nil)).Elem()
}

type SqlJobInput interface {
	pulumi.Input

	ToSqlJobOutput() SqlJobOutput
	ToSqlJobOutputWithContext(ctx context.Context) SqlJobOutput
}

func (*SqlJob) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlJob)(nil)).Elem()
}

func (i *SqlJob) ToSqlJobOutput() SqlJobOutput {
	return i.ToSqlJobOutputWithContext(context.Background())
}

func (i *SqlJob) ToSqlJobOutputWithContext(ctx context.Context) SqlJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlJobOutput)
}

// SqlJobArrayInput is an input type that accepts SqlJobArray and SqlJobArrayOutput values.
// You can construct a concrete instance of `SqlJobArrayInput` via:
//
//	SqlJobArray{ SqlJobArgs{...} }
type SqlJobArrayInput interface {
	pulumi.Input

	ToSqlJobArrayOutput() SqlJobArrayOutput
	ToSqlJobArrayOutputWithContext(context.Context) SqlJobArrayOutput
}

type SqlJobArray []SqlJobInput

func (SqlJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlJob)(nil)).Elem()
}

func (i SqlJobArray) ToSqlJobArrayOutput() SqlJobArrayOutput {
	return i.ToSqlJobArrayOutputWithContext(context.Background())
}

func (i SqlJobArray) ToSqlJobArrayOutputWithContext(ctx context.Context) SqlJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlJobArrayOutput)
}

// SqlJobMapInput is an input type that accepts SqlJobMap and SqlJobMapOutput values.
// You can construct a concrete instance of `SqlJobMapInput` via:
//
//	SqlJobMap{ "key": SqlJobArgs{...} }
type SqlJobMapInput interface {
	pulumi.Input

	ToSqlJobMapOutput() SqlJobMapOutput
	ToSqlJobMapOutputWithContext(context.Context) SqlJobMapOutput
}

type SqlJobMap map[string]SqlJobInput

func (SqlJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlJob)(nil)).Elem()
}

func (i SqlJobMap) ToSqlJobMapOutput() SqlJobMapOutput {
	return i.ToSqlJobMapOutputWithContext(context.Background())
}

func (i SqlJobMap) ToSqlJobMapOutputWithContext(ctx context.Context) SqlJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlJobMapOutput)
}

type SqlJobOutput struct{ *pulumi.OutputState }

func (SqlJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlJob)(nil)).Elem()
}

func (o SqlJobOutput) ToSqlJobOutput() SqlJobOutput {
	return o
}

func (o SqlJobOutput) ToSqlJobOutputWithContext(ctx context.Context) SqlJobOutput {
	return o
}

// Specifies the configuration parameters for the SQL job. Changing this parameter
// will create a new resource. Structure is documented below.
func (o SqlJobOutput) Conf() SqlJobConfPtrOutput {
	return o.ApplyT(func(v *SqlJob) SqlJobConfPtrOutput { return v.Conf }).(SqlJobConfPtrOutput)
}

// Specifies the database where the SQL is executed. This argument does
// not need to be configured during database creation. Changing this parameter will create a new resource.
func (o SqlJobOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// Job running duration (unit: millisecond).
func (o SqlJobOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

// Type of a job, Includes **DDL**, **DCL**, **IMPORT**, **EXPORT**, **QUERY**, **INSERT**,
// **DATA_MIGRATION**, **UPDATE**, **DELETE**, **RESTART_QUEUE** and **SCALE_QUEUE**.
func (o SqlJobOutput) JobType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.JobType }).(pulumi.StringOutput)
}

// User who submits a job.
func (o SqlJobOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Specifies queue which this job to be submitted belongs.
// Changing this parameter will create a new resource.
func (o SqlJobOutput) QueueName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.QueueName }).(pulumi.StringOutput)
}

// Specifies the region in which to create the DLI table resource. If omitted,
// the provider-level region will be used. Changing this parameter will create a new resource.
func (o SqlJobOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// When the statement type is DDL, results of the DDL are displayed.
func (o SqlJobOutput) Rows() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringArrayArrayOutput { return v.Rows }).(pulumi.StringArrayArrayOutput)
}

// When the statement type is DDL, the column name and type of DDL are displayed.
func (o SqlJobOutput) Schemas() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringMapArrayOutput { return v.Schemas }).(pulumi.StringMapArrayOutput)
}

// Specifies SQL statement that you want to execute.
// Changing this parameter will create a new resource.
func (o SqlJobOutput) Sql() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.Sql }).(pulumi.StringOutput)
}

// Time when a job is started, in RFC-3339 format. e.g. `2019-10-12T07:20:50.52Z`
func (o SqlJobOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Status of a job, including **RUNNING**, **SCALING**, **LAUNCHING**, **FINISHED**, **FAILED**,
// and **CANCELLED.**
func (o SqlJobOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies label of a Job. Changing this parameter will create a new resource.
func (o SqlJobOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlJob) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type SqlJobArrayOutput struct{ *pulumi.OutputState }

func (SqlJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlJob)(nil)).Elem()
}

func (o SqlJobArrayOutput) ToSqlJobArrayOutput() SqlJobArrayOutput {
	return o
}

func (o SqlJobArrayOutput) ToSqlJobArrayOutputWithContext(ctx context.Context) SqlJobArrayOutput {
	return o
}

func (o SqlJobArrayOutput) Index(i pulumi.IntInput) SqlJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlJob {
		return vs[0].([]*SqlJob)[vs[1].(int)]
	}).(SqlJobOutput)
}

type SqlJobMapOutput struct{ *pulumi.OutputState }

func (SqlJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlJob)(nil)).Elem()
}

func (o SqlJobMapOutput) ToSqlJobMapOutput() SqlJobMapOutput {
	return o
}

func (o SqlJobMapOutput) ToSqlJobMapOutputWithContext(ctx context.Context) SqlJobMapOutput {
	return o
}

func (o SqlJobMapOutput) MapIndex(k pulumi.StringInput) SqlJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlJob {
		return vs[0].(map[string]*SqlJob)[vs[1].(string)]
	}).(SqlJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlJobInput)(nil)).Elem(), &SqlJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlJobArrayInput)(nil)).Elem(), SqlJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlJobMapInput)(nil)).Elem(), SqlJobMap{})
	pulumi.RegisterOutputType(SqlJobOutput{})
	pulumi.RegisterOutputType(SqlJobArrayOutput{})
	pulumi.RegisterOutputType(SqlJobMapOutput{})
}
