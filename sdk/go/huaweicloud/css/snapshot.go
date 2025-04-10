// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSS cluster snapshot management
//
// ## Example Usage
// ### Create a snapshot
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Css"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Css.NewSnapshot(ctx, "snapshot", &Css.SnapshotArgs{
//				Description: pulumi.String("a snapshot created by manual"),
//				ClusterId:   pulumi.Any(_var.Css_cluster_id),
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
// # This resource can be imported by specifying the CSS cluster ID and snapshot ID separated by a slash, e.g.bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Css/snapshot:Snapshot snapshot_1 <cluster_id>/<snapshot_id>
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Indicates the snapshot creation mode, the value should be "manual" or "automated".
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// Specifies ID of the CSS cluster where index data is to be backed up.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Indicates the CSS cluster name.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specifies the description of a snapshot. The value contains 0 to 256
	// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the index to be backed up. Multiple index names are
	// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
	// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
	// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
	// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
	Index pulumi.StringOutput `pulumi:"index"`
	// Specifies the snapshot name. The snapshot name must start with a letter and
	// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
	// this parameter will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the snapshot status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("huaweicloud:Css/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("huaweicloud:Css/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Indicates the snapshot creation mode, the value should be "manual" or "automated".
	BackupType *string `pulumi:"backupType"`
	// Specifies ID of the CSS cluster where index data is to be backed up.
	// Changing this parameter will create a new resource.
	ClusterId *string `pulumi:"clusterId"`
	// Indicates the CSS cluster name.
	ClusterName *string `pulumi:"clusterName"`
	// Specifies the description of a snapshot. The value contains 0 to 256
	// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
	Description *string `pulumi:"description"`
	// Specifies the name of the index to be backed up. Multiple index names are
	// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
	// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
	// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
	// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
	Index *string `pulumi:"index"`
	// Specifies the snapshot name. The snapshot name must start with a letter and
	// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
	// this parameter will create a new resource.
	Name *string `pulumi:"name"`
	// Indicates the snapshot status.
	Status *string `pulumi:"status"`
}

type SnapshotState struct {
	// Indicates the snapshot creation mode, the value should be "manual" or "automated".
	BackupType pulumi.StringPtrInput
	// Specifies ID of the CSS cluster where index data is to be backed up.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringPtrInput
	// Indicates the CSS cluster name.
	ClusterName pulumi.StringPtrInput
	// Specifies the description of a snapshot. The value contains 0 to 256
	// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
	Description pulumi.StringPtrInput
	// Specifies the name of the index to be backed up. Multiple index names are
	// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
	// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
	// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
	// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
	Index pulumi.StringPtrInput
	// Specifies the snapshot name. The snapshot name must start with a letter and
	// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
	// this parameter will create a new resource.
	Name pulumi.StringPtrInput
	// Indicates the snapshot status.
	Status pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Specifies ID of the CSS cluster where index data is to be backed up.
	// Changing this parameter will create a new resource.
	ClusterId string `pulumi:"clusterId"`
	// Specifies the description of a snapshot. The value contains 0 to 256
	// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
	Description *string `pulumi:"description"`
	// Specifies the name of the index to be backed up. Multiple index names are
	// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
	// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
	// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
	// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
	Index *string `pulumi:"index"`
	// Specifies the snapshot name. The snapshot name must start with a letter and
	// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
	// this parameter will create a new resource.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Specifies ID of the CSS cluster where index data is to be backed up.
	// Changing this parameter will create a new resource.
	ClusterId pulumi.StringInput
	// Specifies the description of a snapshot. The value contains 0 to 256
	// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
	Description pulumi.StringPtrInput
	// Specifies the name of the index to be backed up. Multiple index names are
	// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
	// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
	// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
	// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
	Index pulumi.StringPtrInput
	// Specifies the snapshot name. The snapshot name must start with a letter and
	// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
	// this parameter will create a new resource.
	Name pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// Indicates the snapshot creation mode, the value should be "manual" or "automated".
func (o SnapshotOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// Specifies ID of the CSS cluster where index data is to be backed up.
// Changing this parameter will create a new resource.
func (o SnapshotOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Indicates the CSS cluster name.
func (o SnapshotOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Specifies the description of a snapshot. The value contains 0 to 256
// characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the name of the index to be backed up. Multiple index names are
// separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (*) to back up data of
// certain indices. For example, if you enter 2020-06*, then data of indices with the name prefix of 2020-06 will be
// backed up. The value contains 0 to 1024 characters. Uppercase letters, spaces, and certain special characters (
// including "\\<|>/?) are not allowed. Changing this parameter will create a new resource.
func (o SnapshotOutput) Index() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Index }).(pulumi.StringOutput)
}

// Specifies the snapshot name. The snapshot name must start with a letter and
// contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing
// this parameter will create a new resource.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates the snapshot status.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
