// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package modelarts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage storages mounted to the notebook resource within HuaweiCloud. A maximum of 10 storages can be mounted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/ModelArts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			notebookId := cfg.RequireObject("notebookId")
//			uriParallelObs := cfg.RequireObject("uriParallelObs")
//			_, err := ModelArts.NewNotebookMountStorage(ctx, "test", &ModelArts.NotebookMountStorageArgs{
//				NotebookId:          pulumi.Any(notebookId),
//				StoragePath:         pulumi.Any(uriParallelObs),
//				LocalMountDirectory: pulumi.String("/data/test/"),
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
// The mount storage can be imported by `id`, It is composed of the ID of notebook and mount ID, separated by a slash.
//
// For example,
//
// ```sh
//
//	$ pulumi import huaweicloud:ModelArts/notebookMountStorage:NotebookMountStorage test b11b407c-e604-4e8d-8bc4-92398320b847/4e206d3c-6831-4267-b93d-e236105cda38
//
// ```
type NotebookMountStorage struct {
	pulumi.CustomResourceState

	// Specifies the local mount directory.
	// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
	// Changing this parameter will create a new resource.
	LocalMountDirectory pulumi.StringOutput `pulumi:"localMountDirectory"`
	// The mount ID.
	MountId pulumi.StringOutput `pulumi:"mountId"`
	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookId pulumi.StringOutput `pulumi:"notebookId"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
	// `UNMOUNT_FAILED`, `UNMOUNTED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
	StoragePath pulumi.StringOutput `pulumi:"storagePath"`
	// The type of storage system.  The value is `OBSFS`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotebookMountStorage registers a new resource with the given unique name, arguments, and options.
func NewNotebookMountStorage(ctx *pulumi.Context,
	name string, args *NotebookMountStorageArgs, opts ...pulumi.ResourceOption) (*NotebookMountStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalMountDirectory == nil {
		return nil, errors.New("invalid value for required argument 'LocalMountDirectory'")
	}
	if args.NotebookId == nil {
		return nil, errors.New("invalid value for required argument 'NotebookId'")
	}
	if args.StoragePath == nil {
		return nil, errors.New("invalid value for required argument 'StoragePath'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NotebookMountStorage
	err := ctx.RegisterResource("huaweicloud:ModelArts/notebookMountStorage:NotebookMountStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookMountStorage gets an existing NotebookMountStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookMountStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookMountStorageState, opts ...pulumi.ResourceOption) (*NotebookMountStorage, error) {
	var resource NotebookMountStorage
	err := ctx.ReadResource("huaweicloud:ModelArts/notebookMountStorage:NotebookMountStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookMountStorage resources.
type notebookMountStorageState struct {
	// Specifies the local mount directory.
	// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
	// Changing this parameter will create a new resource.
	LocalMountDirectory *string `pulumi:"localMountDirectory"`
	// The mount ID.
	MountId *string `pulumi:"mountId"`
	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookId *string `pulumi:"notebookId"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
	// `UNMOUNT_FAILED`, `UNMOUNTED`.
	Status *string `pulumi:"status"`
	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
	StoragePath *string `pulumi:"storagePath"`
	// The type of storage system.  The value is `OBSFS`.
	Type *string `pulumi:"type"`
}

type NotebookMountStorageState struct {
	// Specifies the local mount directory.
	// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
	// Changing this parameter will create a new resource.
	LocalMountDirectory pulumi.StringPtrInput
	// The mount ID.
	MountId pulumi.StringPtrInput
	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookId pulumi.StringPtrInput
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
	// `UNMOUNT_FAILED`, `UNMOUNTED`.
	Status pulumi.StringPtrInput
	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
	StoragePath pulumi.StringPtrInput
	// The type of storage system.  The value is `OBSFS`.
	Type pulumi.StringPtrInput
}

func (NotebookMountStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookMountStorageState)(nil)).Elem()
}

type notebookMountStorageArgs struct {
	// Specifies the local mount directory.
	// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
	// Changing this parameter will create a new resource.
	LocalMountDirectory string `pulumi:"localMountDirectory"`
	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookId string `pulumi:"notebookId"`
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
	StoragePath string `pulumi:"storagePath"`
}

// The set of arguments for constructing a NotebookMountStorage resource.
type NotebookMountStorageArgs struct {
	// Specifies the local mount directory.
	// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
	// Changing this parameter will create a new resource.
	LocalMountDirectory pulumi.StringInput
	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookId pulumi.StringInput
	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
	StoragePath pulumi.StringInput
}

func (NotebookMountStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookMountStorageArgs)(nil)).Elem()
}

type NotebookMountStorageInput interface {
	pulumi.Input

	ToNotebookMountStorageOutput() NotebookMountStorageOutput
	ToNotebookMountStorageOutputWithContext(ctx context.Context) NotebookMountStorageOutput
}

func (*NotebookMountStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookMountStorage)(nil)).Elem()
}

func (i *NotebookMountStorage) ToNotebookMountStorageOutput() NotebookMountStorageOutput {
	return i.ToNotebookMountStorageOutputWithContext(context.Background())
}

func (i *NotebookMountStorage) ToNotebookMountStorageOutputWithContext(ctx context.Context) NotebookMountStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookMountStorageOutput)
}

// NotebookMountStorageArrayInput is an input type that accepts NotebookMountStorageArray and NotebookMountStorageArrayOutput values.
// You can construct a concrete instance of `NotebookMountStorageArrayInput` via:
//
//	NotebookMountStorageArray{ NotebookMountStorageArgs{...} }
type NotebookMountStorageArrayInput interface {
	pulumi.Input

	ToNotebookMountStorageArrayOutput() NotebookMountStorageArrayOutput
	ToNotebookMountStorageArrayOutputWithContext(context.Context) NotebookMountStorageArrayOutput
}

type NotebookMountStorageArray []NotebookMountStorageInput

func (NotebookMountStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotebookMountStorage)(nil)).Elem()
}

func (i NotebookMountStorageArray) ToNotebookMountStorageArrayOutput() NotebookMountStorageArrayOutput {
	return i.ToNotebookMountStorageArrayOutputWithContext(context.Background())
}

func (i NotebookMountStorageArray) ToNotebookMountStorageArrayOutputWithContext(ctx context.Context) NotebookMountStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookMountStorageArrayOutput)
}

// NotebookMountStorageMapInput is an input type that accepts NotebookMountStorageMap and NotebookMountStorageMapOutput values.
// You can construct a concrete instance of `NotebookMountStorageMapInput` via:
//
//	NotebookMountStorageMap{ "key": NotebookMountStorageArgs{...} }
type NotebookMountStorageMapInput interface {
	pulumi.Input

	ToNotebookMountStorageMapOutput() NotebookMountStorageMapOutput
	ToNotebookMountStorageMapOutputWithContext(context.Context) NotebookMountStorageMapOutput
}

type NotebookMountStorageMap map[string]NotebookMountStorageInput

func (NotebookMountStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotebookMountStorage)(nil)).Elem()
}

func (i NotebookMountStorageMap) ToNotebookMountStorageMapOutput() NotebookMountStorageMapOutput {
	return i.ToNotebookMountStorageMapOutputWithContext(context.Background())
}

func (i NotebookMountStorageMap) ToNotebookMountStorageMapOutputWithContext(ctx context.Context) NotebookMountStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookMountStorageMapOutput)
}

type NotebookMountStorageOutput struct{ *pulumi.OutputState }

func (NotebookMountStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookMountStorage)(nil)).Elem()
}

func (o NotebookMountStorageOutput) ToNotebookMountStorageOutput() NotebookMountStorageOutput {
	return o
}

func (o NotebookMountStorageOutput) ToNotebookMountStorageOutputWithContext(ctx context.Context) NotebookMountStorageOutput {
	return o
}

// Specifies the local mount directory.
// Only the subdirectory of `/data/` can be mounted. The format is : `/data/dir1/`.
// Changing this parameter will create a new resource.
func (o NotebookMountStorageOutput) LocalMountDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.LocalMountDirectory }).(pulumi.StringOutput)
}

// The mount ID.
func (o NotebookMountStorageOutput) MountId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.MountId }).(pulumi.StringOutput)
}

// Specifies ID of notebook which the storage be mounted to.
// Changing this parameter will create a new resource.
func (o NotebookMountStorageOutput) NotebookId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.NotebookId }).(pulumi.StringOutput)
}

// The region in which to create the resource. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o NotebookMountStorageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// mount status. Valid values include: `MOUNTING`, `MOUNT_FAILED`, `MOUNTED`, `UNMOUNTING`,
// `UNMOUNT_FAILED`, `UNMOUNTED`.
func (o NotebookMountStorageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the path of Parallel File System (PFS) or its folders in OBS.
// The format is : `obs://obs-bucket/folder/`. Changing this parameter will create a new resource.
func (o NotebookMountStorageOutput) StoragePath() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.StoragePath }).(pulumi.StringOutput)
}

// The type of storage system.  The value is `OBSFS`.
func (o NotebookMountStorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookMountStorage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type NotebookMountStorageArrayOutput struct{ *pulumi.OutputState }

func (NotebookMountStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotebookMountStorage)(nil)).Elem()
}

func (o NotebookMountStorageArrayOutput) ToNotebookMountStorageArrayOutput() NotebookMountStorageArrayOutput {
	return o
}

func (o NotebookMountStorageArrayOutput) ToNotebookMountStorageArrayOutputWithContext(ctx context.Context) NotebookMountStorageArrayOutput {
	return o
}

func (o NotebookMountStorageArrayOutput) Index(i pulumi.IntInput) NotebookMountStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotebookMountStorage {
		return vs[0].([]*NotebookMountStorage)[vs[1].(int)]
	}).(NotebookMountStorageOutput)
}

type NotebookMountStorageMapOutput struct{ *pulumi.OutputState }

func (NotebookMountStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotebookMountStorage)(nil)).Elem()
}

func (o NotebookMountStorageMapOutput) ToNotebookMountStorageMapOutput() NotebookMountStorageMapOutput {
	return o
}

func (o NotebookMountStorageMapOutput) ToNotebookMountStorageMapOutputWithContext(ctx context.Context) NotebookMountStorageMapOutput {
	return o
}

func (o NotebookMountStorageMapOutput) MapIndex(k pulumi.StringInput) NotebookMountStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotebookMountStorage {
		return vs[0].(map[string]*NotebookMountStorage)[vs[1].(string)]
	}).(NotebookMountStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookMountStorageInput)(nil)).Elem(), &NotebookMountStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookMountStorageArrayInput)(nil)).Elem(), NotebookMountStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookMountStorageMapInput)(nil)).Elem(), NotebookMountStorageMap{})
	pulumi.RegisterOutputType(NotebookMountStorageOutput{})
	pulumi.RegisterOutputType(NotebookMountStorageArrayOutput{})
	pulumi.RegisterOutputType(NotebookMountStorageMapOutput{})
}