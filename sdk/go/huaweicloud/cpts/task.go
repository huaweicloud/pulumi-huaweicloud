// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cpts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a pressure Test Task resource within HuaweiCloud CPTS.
// The task resource only supports serial execution mode.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cpts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testProject, err := Cpts.NewProject(ctx, "testProject", nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cpts.NewTask(ctx, "testTask", &Cpts.TaskArgs{
//				ProjectId: testProject.ID(),
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
// Tasks can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Cpts/task:Task test 1090
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`operation`. It is generally recommended running `terraform plan` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. resource "huaweicloud_cpts_task" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	operation,
//
//	]
//
//	} }
type Task struct {
	pulumi.CustomResourceState

	// Specifies benchmark concurrency of the task, the value range is 0 to
	// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
	// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
	BenchmarkConcurrency pulumi.IntPtrOutput `pulumi:"benchmarkConcurrency"`
	// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
	// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
	ClusterId pulumi.IntPtrOutput `pulumi:"clusterId"`
	// Specifies the name of the task, which can contain a maximum of 42 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether to enable the task or stop the task. The options are as follows:
	// + **enable**: Starting the pressure test task.
	// + **stop**: Stop the pressure test tasks.
	Operation pulumi.StringPtrOutput `pulumi:"operation"`
	// Specifies the CPTS project ID which the task belongs to.
	// Changing this parameter will create a new resource.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Specifies the region in which to create the task resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of the task. The options are as follows:
	// + **0**: The task is running.
	// + **2**: The task is idle.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Task
	err := ctx.RegisterResource("huaweicloud:Cpts/task:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("huaweicloud:Cpts/task:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
	// Specifies benchmark concurrency of the task, the value range is 0 to
	// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
	// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
	BenchmarkConcurrency *int `pulumi:"benchmarkConcurrency"`
	// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
	// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
	ClusterId *int `pulumi:"clusterId"`
	// Specifies the name of the task, which can contain a maximum of 42 characters.
	Name *string `pulumi:"name"`
	// Specifies whether to enable the task or stop the task. The options are as follows:
	// + **enable**: Starting the pressure test task.
	// + **stop**: Stop the pressure test tasks.
	Operation *string `pulumi:"operation"`
	// Specifies the CPTS project ID which the task belongs to.
	// Changing this parameter will create a new resource.
	ProjectId *int `pulumi:"projectId"`
	// Specifies the region in which to create the task resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// The status of the task. The options are as follows:
	// + **0**: The task is running.
	// + **2**: The task is idle.
	Status *int `pulumi:"status"`
}

type TaskState struct {
	// Specifies benchmark concurrency of the task, the value range is 0 to
	// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
	// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
	BenchmarkConcurrency pulumi.IntPtrInput
	// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
	// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
	ClusterId pulumi.IntPtrInput
	// Specifies the name of the task, which can contain a maximum of 42 characters.
	Name pulumi.StringPtrInput
	// Specifies whether to enable the task or stop the task. The options are as follows:
	// + **enable**: Starting the pressure test task.
	// + **stop**: Stop the pressure test tasks.
	Operation pulumi.StringPtrInput
	// Specifies the CPTS project ID which the task belongs to.
	// Changing this parameter will create a new resource.
	ProjectId pulumi.IntPtrInput
	// Specifies the region in which to create the task resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// The status of the task. The options are as follows:
	// + **0**: The task is running.
	// + **2**: The task is idle.
	Status pulumi.IntPtrInput
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// Specifies benchmark concurrency of the task, the value range is 0 to
	// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
	// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
	BenchmarkConcurrency *int `pulumi:"benchmarkConcurrency"`
	// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
	// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
	ClusterId *int `pulumi:"clusterId"`
	// Specifies the name of the task, which can contain a maximum of 42 characters.
	Name *string `pulumi:"name"`
	// Specifies whether to enable the task or stop the task. The options are as follows:
	// + **enable**: Starting the pressure test task.
	// + **stop**: Stop the pressure test tasks.
	Operation *string `pulumi:"operation"`
	// Specifies the CPTS project ID which the task belongs to.
	// Changing this parameter will create a new resource.
	ProjectId int `pulumi:"projectId"`
	// Specifies the region in which to create the task resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// Specifies benchmark concurrency of the task, the value range is 0 to
	// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
	// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
	BenchmarkConcurrency pulumi.IntPtrInput
	// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
	// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
	ClusterId pulumi.IntPtrInput
	// Specifies the name of the task, which can contain a maximum of 42 characters.
	Name pulumi.StringPtrInput
	// Specifies whether to enable the task or stop the task. The options are as follows:
	// + **enable**: Starting the pressure test task.
	// + **stop**: Stop the pressure test tasks.
	Operation pulumi.StringPtrInput
	// Specifies the CPTS project ID which the task belongs to.
	// Changing this parameter will create a new resource.
	ProjectId pulumi.IntInput
	// Specifies the region in which to create the task resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}

type TaskInput interface {
	pulumi.Input

	ToTaskOutput() TaskOutput
	ToTaskOutputWithContext(ctx context.Context) TaskOutput
}

func (*Task) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

// TaskArrayInput is an input type that accepts TaskArray and TaskArrayOutput values.
// You can construct a concrete instance of `TaskArrayInput` via:
//
//	TaskArray{ TaskArgs{...} }
type TaskArrayInput interface {
	pulumi.Input

	ToTaskArrayOutput() TaskArrayOutput
	ToTaskArrayOutputWithContext(context.Context) TaskArrayOutput
}

type TaskArray []TaskInput

func (TaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (i TaskArray) ToTaskArrayOutput() TaskArrayOutput {
	return i.ToTaskArrayOutputWithContext(context.Background())
}

func (i TaskArray) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskArrayOutput)
}

// TaskMapInput is an input type that accepts TaskMap and TaskMapOutput values.
// You can construct a concrete instance of `TaskMapInput` via:
//
//	TaskMap{ "key": TaskArgs{...} }
type TaskMapInput interface {
	pulumi.Input

	ToTaskMapOutput() TaskMapOutput
	ToTaskMapOutputWithContext(context.Context) TaskMapOutput
}

type TaskMap map[string]TaskInput

func (TaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (i TaskMap) ToTaskMapOutput() TaskMapOutput {
	return i.ToTaskMapOutputWithContext(context.Background())
}

func (i TaskMap) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskMapOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

// Specifies benchmark concurrency of the task, the value range is 0 to
// 2000000. The default value is `100`. Reference for the calculation of the number of concurrent users.
// `Number of concurrent users` = `benchmark concurrency` * `concurrency ratio`.
func (o TaskOutput) BenchmarkConcurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.IntPtrOutput { return v.BenchmarkConcurrency }).(pulumi.IntPtrOutput)
}

// Specifies a cluster ID of the CPTS resource group. If the number of concurrent users
// is less than 1000, you can use a shared resource group for testing and do not have to create a resource group.
func (o TaskOutput) ClusterId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.IntPtrOutput { return v.ClusterId }).(pulumi.IntPtrOutput)
}

// Specifies the name of the task, which can contain a maximum of 42 characters.
func (o TaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether to enable the task or stop the task. The options are as follows:
// + **enable**: Starting the pressure test task.
// + **stop**: Stop the pressure test tasks.
func (o TaskOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.StringPtrOutput { return v.Operation }).(pulumi.StringPtrOutput)
}

// Specifies the CPTS project ID which the task belongs to.
// Changing this parameter will create a new resource.
func (o TaskOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Specifies the region in which to create the task resource. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o TaskOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The status of the task. The options are as follows:
// + **0**: The task is running.
// + **2**: The task is idle.
func (o TaskOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type TaskArrayOutput struct{ *pulumi.OutputState }

func (TaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (o TaskArrayOutput) ToTaskArrayOutput() TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) Index(i pulumi.IntInput) TaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Task {
		return vs[0].([]*Task)[vs[1].(int)]
	}).(TaskOutput)
}

type TaskMapOutput struct{ *pulumi.OutputState }

func (TaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (o TaskMapOutput) ToTaskMapOutput() TaskMapOutput {
	return o
}

func (o TaskMapOutput) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return o
}

func (o TaskMapOutput) MapIndex(k pulumi.StringInput) TaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Task {
		return vs[0].(map[string]*Task)[vs[1].(string)]
	}).(TaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskInput)(nil)).Elem(), &Task{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskArrayInput)(nil)).Elem(), TaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskMapInput)(nil)).Elem(), TaskMap{})
	pulumi.RegisterOutputType(TaskOutput{})
	pulumi.RegisterOutputType(TaskArrayOutput{})
	pulumi.RegisterOutputType(TaskMapOutput{})
}