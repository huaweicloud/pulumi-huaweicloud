// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dli

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages DLI Queue resource within HuaweiCloud
//
// ## Example Usage
// ### Create an exclusive mode queue
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
//			elasticResourcePoolName := cfg.RequireObject("elasticResourcePoolName")
//			queueName := cfg.RequireObject("queueName")
//			_, err := Dli.NewQueue(ctx, "test", &Dli.QueueArgs{
//				ElasticResourcePoolName: pulumi.Any(elasticResourcePoolName),
//				ResourceMode:            pulumi.Int(1),
//				CuCount:                 pulumi.Int(16),
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//					"key": pulumi.String("value"),
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
// ### Import a SQL type queue bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Dli/queue:Queue test <name>
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response. The missing attributes include`tags`. It is generally recommended running `terraform plan` after importing a DLI queue. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the resource. Also you can ignore changes as below. hcl resource "huaweicloud_dli_queue" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	tags
//
//	]
//
//	} }
type Queue struct {
	pulumi.CustomResourceState

	// Time when a queue is created.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
	// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
	CuCount pulumi.IntOutput `pulumi:"cuCount"`
	// Description of a queue. Changing this parameter will create a new
	// resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the name of the elastic resource pool to which the queue
	// belongs.
	ElasticResourcePoolName pulumi.StringOutput `pulumi:"elasticResourcePoolName"`
	// Enterprise project ID. The value 0 indicates the default
	// enterprise project. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Indicates the queue feature. Changing this parameter will create a new
	// resource. The options are as follows:
	// + basic: basic type (default value)
	// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
	Feature pulumi.StringPtrOutput `pulumi:"feature"`
	// Deprecated: management_subnet_cidr is Deprecated
	ManagementSubnetCidr pulumi.StringPtrOutput `pulumi:"managementSubnetCidr"`
	// Name of a queue. Name of a newly created resource queue. The name can contain
	// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
	// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// CPU architecture of queue compute resources. Changing this parameter will
	// create a new resource. The options are as follows:
	// + x8664 : default value
	// + aarch64
	Platform pulumi.StringPtrOutput `pulumi:"platform"`
	// Indicates the queue type. Changing this parameter will create a new
	// resource. The options are as follows:
	// + sql
	// + general
	QueueType pulumi.StringPtrOutput `pulumi:"queueType"`
	// Specifies the region in which to create the dli queue resource. If omitted,
	// the provider-level region will be used. Changing this will create a new VPC channel resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the queue resource mode.\
	// The valid value is as follows:
	// + 1: indicates the exclusive resource mode.
	ResourceMode pulumi.IntOutput `pulumi:"resourceMode"`
	// Specifies the list of scaling policies of the queue associated with
	// an elastic resource pool.
	// This parameter is only available if `resourceMode` is set to `1`.
	// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
	// The scalingPolicies structure is documented below.
	ScalingPolicies QueueScalingPolicyArrayOutput `pulumi:"scalingPolicies"`
	// Specifies spark driver configuration of the queue.
	// This parameter is only available if `queueType` is set to `sql`.
	// The sparkDriver structure is documented below.
	SparkDriver QueueSparkDriverPtrOutput `pulumi:"sparkDriver"`
	// Deprecated: subnet_cidr is Deprecated
	SubnetCidr pulumi.StringPtrOutput `pulumi:"subnetCidr"`
	// Label of a queue. Changing this parameter will create a new resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The CIDR block of the queue.
	VpcCidr pulumi.StringOutput `pulumi:"vpcCidr"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CuCount == nil {
		return nil, errors.New("invalid value for required argument 'CuCount'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Queue
	err := ctx.RegisterResource("huaweicloud:Dli/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("huaweicloud:Dli/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// Time when a queue is created.
	CreateTime *int `pulumi:"createTime"`
	// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
	// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
	CuCount *int `pulumi:"cuCount"`
	// Description of a queue. Changing this parameter will create a new
	// resource.
	Description *string `pulumi:"description"`
	// Specifies the name of the elastic resource pool to which the queue
	// belongs.
	ElasticResourcePoolName *string `pulumi:"elasticResourcePoolName"`
	// Enterprise project ID. The value 0 indicates the default
	// enterprise project. Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Indicates the queue feature. Changing this parameter will create a new
	// resource. The options are as follows:
	// + basic: basic type (default value)
	// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
	Feature *string `pulumi:"feature"`
	// Deprecated: management_subnet_cidr is Deprecated
	ManagementSubnetCidr *string `pulumi:"managementSubnetCidr"`
	// Name of a queue. Name of a newly created resource queue. The name can contain
	// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
	// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// CPU architecture of queue compute resources. Changing this parameter will
	// create a new resource. The options are as follows:
	// + x8664 : default value
	// + aarch64
	Platform *string `pulumi:"platform"`
	// Indicates the queue type. Changing this parameter will create a new
	// resource. The options are as follows:
	// + sql
	// + general
	QueueType *string `pulumi:"queueType"`
	// Specifies the region in which to create the dli queue resource. If omitted,
	// the provider-level region will be used. Changing this will create a new VPC channel resource.
	Region *string `pulumi:"region"`
	// Specifies the queue resource mode.\
	// The valid value is as follows:
	// + 1: indicates the exclusive resource mode.
	ResourceMode *int `pulumi:"resourceMode"`
	// Specifies the list of scaling policies of the queue associated with
	// an elastic resource pool.
	// This parameter is only available if `resourceMode` is set to `1`.
	// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
	// The scalingPolicies structure is documented below.
	ScalingPolicies []QueueScalingPolicy `pulumi:"scalingPolicies"`
	// Specifies spark driver configuration of the queue.
	// This parameter is only available if `queueType` is set to `sql`.
	// The sparkDriver structure is documented below.
	SparkDriver *QueueSparkDriver `pulumi:"sparkDriver"`
	// Deprecated: subnet_cidr is Deprecated
	SubnetCidr *string `pulumi:"subnetCidr"`
	// Label of a queue. Changing this parameter will create a new resource.
	Tags map[string]string `pulumi:"tags"`
	// The CIDR block of the queue.
	VpcCidr *string `pulumi:"vpcCidr"`
}

type QueueState struct {
	// Time when a queue is created.
	CreateTime pulumi.IntPtrInput
	// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
	// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
	CuCount pulumi.IntPtrInput
	// Description of a queue. Changing this parameter will create a new
	// resource.
	Description pulumi.StringPtrInput
	// Specifies the name of the elastic resource pool to which the queue
	// belongs.
	ElasticResourcePoolName pulumi.StringPtrInput
	// Enterprise project ID. The value 0 indicates the default
	// enterprise project. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Indicates the queue feature. Changing this parameter will create a new
	// resource. The options are as follows:
	// + basic: basic type (default value)
	// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
	Feature pulumi.StringPtrInput
	// Deprecated: management_subnet_cidr is Deprecated
	ManagementSubnetCidr pulumi.StringPtrInput
	// Name of a queue. Name of a newly created resource queue. The name can contain
	// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
	// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
	Name pulumi.StringPtrInput
	// CPU architecture of queue compute resources. Changing this parameter will
	// create a new resource. The options are as follows:
	// + x8664 : default value
	// + aarch64
	Platform pulumi.StringPtrInput
	// Indicates the queue type. Changing this parameter will create a new
	// resource. The options are as follows:
	// + sql
	// + general
	QueueType pulumi.StringPtrInput
	// Specifies the region in which to create the dli queue resource. If omitted,
	// the provider-level region will be used. Changing this will create a new VPC channel resource.
	Region pulumi.StringPtrInput
	// Specifies the queue resource mode.\
	// The valid value is as follows:
	// + 1: indicates the exclusive resource mode.
	ResourceMode pulumi.IntPtrInput
	// Specifies the list of scaling policies of the queue associated with
	// an elastic resource pool.
	// This parameter is only available if `resourceMode` is set to `1`.
	// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
	// The scalingPolicies structure is documented below.
	ScalingPolicies QueueScalingPolicyArrayInput
	// Specifies spark driver configuration of the queue.
	// This parameter is only available if `queueType` is set to `sql`.
	// The sparkDriver structure is documented below.
	SparkDriver QueueSparkDriverPtrInput
	// Deprecated: subnet_cidr is Deprecated
	SubnetCidr pulumi.StringPtrInput
	// Label of a queue. Changing this parameter will create a new resource.
	Tags pulumi.StringMapInput
	// The CIDR block of the queue.
	VpcCidr pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
	// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
	CuCount int `pulumi:"cuCount"`
	// Description of a queue. Changing this parameter will create a new
	// resource.
	Description *string `pulumi:"description"`
	// Specifies the name of the elastic resource pool to which the queue
	// belongs.
	ElasticResourcePoolName *string `pulumi:"elasticResourcePoolName"`
	// Enterprise project ID. The value 0 indicates the default
	// enterprise project. Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Indicates the queue feature. Changing this parameter will create a new
	// resource. The options are as follows:
	// + basic: basic type (default value)
	// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
	Feature *string `pulumi:"feature"`
	// Deprecated: management_subnet_cidr is Deprecated
	ManagementSubnetCidr *string `pulumi:"managementSubnetCidr"`
	// Name of a queue. Name of a newly created resource queue. The name can contain
	// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
	// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// CPU architecture of queue compute resources. Changing this parameter will
	// create a new resource. The options are as follows:
	// + x8664 : default value
	// + aarch64
	Platform *string `pulumi:"platform"`
	// Indicates the queue type. Changing this parameter will create a new
	// resource. The options are as follows:
	// + sql
	// + general
	QueueType *string `pulumi:"queueType"`
	// Specifies the region in which to create the dli queue resource. If omitted,
	// the provider-level region will be used. Changing this will create a new VPC channel resource.
	Region *string `pulumi:"region"`
	// Specifies the queue resource mode.\
	// The valid value is as follows:
	// + 1: indicates the exclusive resource mode.
	ResourceMode *int `pulumi:"resourceMode"`
	// Specifies the list of scaling policies of the queue associated with
	// an elastic resource pool.
	// This parameter is only available if `resourceMode` is set to `1`.
	// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
	// The scalingPolicies structure is documented below.
	ScalingPolicies []QueueScalingPolicy `pulumi:"scalingPolicies"`
	// Specifies spark driver configuration of the queue.
	// This parameter is only available if `queueType` is set to `sql`.
	// The sparkDriver structure is documented below.
	SparkDriver *QueueSparkDriver `pulumi:"sparkDriver"`
	// Deprecated: subnet_cidr is Deprecated
	SubnetCidr *string `pulumi:"subnetCidr"`
	// Label of a queue. Changing this parameter will create a new resource.
	Tags map[string]string `pulumi:"tags"`
	// The CIDR block of the queue.
	VpcCidr *string `pulumi:"vpcCidr"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
	// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
	CuCount pulumi.IntInput
	// Description of a queue. Changing this parameter will create a new
	// resource.
	Description pulumi.StringPtrInput
	// Specifies the name of the elastic resource pool to which the queue
	// belongs.
	ElasticResourcePoolName pulumi.StringPtrInput
	// Enterprise project ID. The value 0 indicates the default
	// enterprise project. Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Indicates the queue feature. Changing this parameter will create a new
	// resource. The options are as follows:
	// + basic: basic type (default value)
	// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
	Feature pulumi.StringPtrInput
	// Deprecated: management_subnet_cidr is Deprecated
	ManagementSubnetCidr pulumi.StringPtrInput
	// Name of a queue. Name of a newly created resource queue. The name can contain
	// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
	// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
	Name pulumi.StringPtrInput
	// CPU architecture of queue compute resources. Changing this parameter will
	// create a new resource. The options are as follows:
	// + x8664 : default value
	// + aarch64
	Platform pulumi.StringPtrInput
	// Indicates the queue type. Changing this parameter will create a new
	// resource. The options are as follows:
	// + sql
	// + general
	QueueType pulumi.StringPtrInput
	// Specifies the region in which to create the dli queue resource. If omitted,
	// the provider-level region will be used. Changing this will create a new VPC channel resource.
	Region pulumi.StringPtrInput
	// Specifies the queue resource mode.\
	// The valid value is as follows:
	// + 1: indicates the exclusive resource mode.
	ResourceMode pulumi.IntPtrInput
	// Specifies the list of scaling policies of the queue associated with
	// an elastic resource pool.
	// This parameter is only available if `resourceMode` is set to `1`.
	// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
	// The scalingPolicies structure is documented below.
	ScalingPolicies QueueScalingPolicyArrayInput
	// Specifies spark driver configuration of the queue.
	// This parameter is only available if `queueType` is set to `sql`.
	// The sparkDriver structure is documented below.
	SparkDriver QueueSparkDriverPtrInput
	// Deprecated: subnet_cidr is Deprecated
	SubnetCidr pulumi.StringPtrInput
	// Label of a queue. Changing this parameter will create a new resource.
	Tags pulumi.StringMapInput
	// The CIDR block of the queue.
	VpcCidr pulumi.StringPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//	QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//	QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

// Time when a queue is created.
func (o QueueOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// Minimum number of CUs that are bound to a queue. Initial value can be `16`,
// `64`, or `256`. When scaleOut or scale_in, the number must be a multiple of `16`.
func (o QueueOutput) CuCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntOutput { return v.CuCount }).(pulumi.IntOutput)
}

// Description of a queue. Changing this parameter will create a new
// resource.
func (o QueueOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the name of the elastic resource pool to which the queue
// belongs.
func (o QueueOutput) ElasticResourcePoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.ElasticResourcePoolName }).(pulumi.StringOutput)
}

// Enterprise project ID. The value 0 indicates the default
// enterprise project. Changing this parameter will create a new resource.
func (o QueueOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Indicates the queue feature. Changing this parameter will create a new
// resource. The options are as follows:
// + basic: basic type (default value)
// + ai: AI-enhanced (Only the SQL x8664 dedicated queue supports this option.)
func (o QueueOutput) Feature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.Feature }).(pulumi.StringPtrOutput)
}

// Deprecated: management_subnet_cidr is Deprecated
func (o QueueOutput) ManagementSubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.ManagementSubnetCidr }).(pulumi.StringPtrOutput)
}

// Name of a queue. Name of a newly created resource queue. The name can contain
// only digits, lowercase letters, and underscores (\_), but cannot contain only digits or start with an underscore (_).
// Length range: `1` to `128` characters. Changing this parameter will create a new resource.
func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// CPU architecture of queue compute resources. Changing this parameter will
// create a new resource. The options are as follows:
// + x8664 : default value
// + aarch64
func (o QueueOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.Platform }).(pulumi.StringPtrOutput)
}

// Indicates the queue type. Changing this parameter will create a new
// resource. The options are as follows:
// + sql
// + general
func (o QueueOutput) QueueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.QueueType }).(pulumi.StringPtrOutput)
}

// Specifies the region in which to create the dli queue resource. If omitted,
// the provider-level region will be used. Changing this will create a new VPC channel resource.
func (o QueueOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the queue resource mode.\
// The valid value is as follows:
// + 1: indicates the exclusive resource mode.
func (o QueueOutput) ResourceMode() pulumi.IntOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntOutput { return v.ResourceMode }).(pulumi.IntOutput)
}

// Specifies the list of scaling policies of the queue associated with
// an elastic resource pool.
// This parameter is only available if `resourceMode` is set to `1`.
// If you want to use this parameter, you must ensure that there is a scaling policy with a time period from `00:00` to `24:00`.
// The scalingPolicies structure is documented below.
func (o QueueOutput) ScalingPolicies() QueueScalingPolicyArrayOutput {
	return o.ApplyT(func(v *Queue) QueueScalingPolicyArrayOutput { return v.ScalingPolicies }).(QueueScalingPolicyArrayOutput)
}

// Specifies spark driver configuration of the queue.
// This parameter is only available if `queueType` is set to `sql`.
// The sparkDriver structure is documented below.
func (o QueueOutput) SparkDriver() QueueSparkDriverPtrOutput {
	return o.ApplyT(func(v *Queue) QueueSparkDriverPtrOutput { return v.SparkDriver }).(QueueSparkDriverPtrOutput)
}

// Deprecated: subnet_cidr is Deprecated
func (o QueueOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.SubnetCidr }).(pulumi.StringPtrOutput)
}

// Label of a queue. Changing this parameter will create a new resource.
func (o QueueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The CIDR block of the queue.
func (o QueueOutput) VpcCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.VpcCidr }).(pulumi.StringOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].([]*Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].(map[string]*Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueArrayInput)(nil)).Elem(), QueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueMapInput)(nil)).Elem(), QueueMap{})
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
