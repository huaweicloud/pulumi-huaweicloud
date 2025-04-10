// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dli

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages DLI package resource within HuaweiCloud
//
// ## Example Usage
// ### Upload the specified python script as a resource package
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			groupName := cfg.RequireObject("groupName")
//			accessDomainName := cfg.RequireObject("accessDomainName")
//			_, err := Dli.NewPackage(ctx, "queue", &Dli.PackageArgs{
//				GroupName:  pulumi.Any(groupName),
//				ObjectPath: pulumi.String(fmt.Sprintf("https://%v/dli/packages/object_file.py", accessDomainName)),
//				Type:       pulumi.String("pyFile"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Package struct {
	pulumi.CustomResourceState

	// Time when a queue is created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the group name which the package belongs to.
	// Changing this parameter will delete the current package and upload a new package.
	GroupName pulumi.StringPtrOutput `pulumi:"groupName"`
	// Specifies whether to upload resource packages in asynchronous mode.
	// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
	IsAsync pulumi.BoolOutput `pulumi:"isAsync"`
	// The package name.
	ObjectName pulumi.StringOutput `pulumi:"objectName"`
	// Specifies the OBS storage path where the package is located.
	// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
	// Changing this parameter will delete the current package and upload a new package.
	ObjectPath pulumi.StringOutput `pulumi:"objectPath"`
	// Specifies the name of the package owner. The owner must be IAM user.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Specifies the region in which to upload packages.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will delete the current package and upload a new package.
	Region pulumi.StringOutput `pulumi:"region"`
	// Status of a package group to be uploaded.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the key/value pairs to associate with the package.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the package type.
	// + **jar**: `.jar` or jar related files.
	// + **pyFile**: `.py` or python related files.
	// + **file**: Other user files.
	// + **modelFile**: User AI model files.
	Type pulumi.StringOutput `pulumi:"type"`
	// The last time when the package configuration update has completed.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewPackage registers a new resource with the given unique name, arguments, and options.
func NewPackage(ctx *pulumi.Context,
	name string, args *PackageArgs, opts ...pulumi.ResourceOption) (*Package, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ObjectPath == nil {
		return nil, errors.New("invalid value for required argument 'ObjectPath'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Package
	err := ctx.RegisterResource("huaweicloud:Dli/package:Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPackage gets an existing Package resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PackageState, opts ...pulumi.ResourceOption) (*Package, error) {
	var resource Package
	err := ctx.ReadResource("huaweicloud:Dli/package:Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Package resources.
type packageState struct {
	// Time when a queue is created.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the group name which the package belongs to.
	// Changing this parameter will delete the current package and upload a new package.
	GroupName *string `pulumi:"groupName"`
	// Specifies whether to upload resource packages in asynchronous mode.
	// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
	IsAsync *bool `pulumi:"isAsync"`
	// The package name.
	ObjectName *string `pulumi:"objectName"`
	// Specifies the OBS storage path where the package is located.
	// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
	// Changing this parameter will delete the current package and upload a new package.
	ObjectPath *string `pulumi:"objectPath"`
	// Specifies the name of the package owner. The owner must be IAM user.
	Owner *string `pulumi:"owner"`
	// Specifies the region in which to upload packages.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will delete the current package and upload a new package.
	Region *string `pulumi:"region"`
	// Status of a package group to be uploaded.
	Status *string `pulumi:"status"`
	// Specifies the key/value pairs to associate with the package.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the package type.
	// + **jar**: `.jar` or jar related files.
	// + **pyFile**: `.py` or python related files.
	// + **file**: Other user files.
	// + **modelFile**: User AI model files.
	Type *string `pulumi:"type"`
	// The last time when the package configuration update has completed.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type PackageState struct {
	// Time when a queue is created.
	CreatedAt pulumi.StringPtrInput
	// Specifies the group name which the package belongs to.
	// Changing this parameter will delete the current package and upload a new package.
	GroupName pulumi.StringPtrInput
	// Specifies whether to upload resource packages in asynchronous mode.
	// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
	IsAsync pulumi.BoolPtrInput
	// The package name.
	ObjectName pulumi.StringPtrInput
	// Specifies the OBS storage path where the package is located.
	// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
	// Changing this parameter will delete the current package and upload a new package.
	ObjectPath pulumi.StringPtrInput
	// Specifies the name of the package owner. The owner must be IAM user.
	Owner pulumi.StringPtrInput
	// Specifies the region in which to upload packages.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will delete the current package and upload a new package.
	Region pulumi.StringPtrInput
	// Status of a package group to be uploaded.
	Status pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the package.
	Tags pulumi.StringMapInput
	// Specifies the package type.
	// + **jar**: `.jar` or jar related files.
	// + **pyFile**: `.py` or python related files.
	// + **file**: Other user files.
	// + **modelFile**: User AI model files.
	Type pulumi.StringPtrInput
	// The last time when the package configuration update has completed.
	UpdatedAt pulumi.StringPtrInput
}

func (PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*packageState)(nil)).Elem()
}

type packageArgs struct {
	// Specifies the group name which the package belongs to.
	// Changing this parameter will delete the current package and upload a new package.
	GroupName *string `pulumi:"groupName"`
	// Specifies whether to upload resource packages in asynchronous mode.
	// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
	IsAsync *bool `pulumi:"isAsync"`
	// Specifies the OBS storage path where the package is located.
	// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
	// Changing this parameter will delete the current package and upload a new package.
	ObjectPath string `pulumi:"objectPath"`
	// Specifies the name of the package owner. The owner must be IAM user.
	Owner *string `pulumi:"owner"`
	// Specifies the region in which to upload packages.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will delete the current package and upload a new package.
	Region *string `pulumi:"region"`
	// Specifies the key/value pairs to associate with the package.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the package type.
	// + **jar**: `.jar` or jar related files.
	// + **pyFile**: `.py` or python related files.
	// + **file**: Other user files.
	// + **modelFile**: User AI model files.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Package resource.
type PackageArgs struct {
	// Specifies the group name which the package belongs to.
	// Changing this parameter will delete the current package and upload a new package.
	GroupName pulumi.StringPtrInput
	// Specifies whether to upload resource packages in asynchronous mode.
	// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
	IsAsync pulumi.BoolPtrInput
	// Specifies the OBS storage path where the package is located.
	// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
	// Changing this parameter will delete the current package and upload a new package.
	ObjectPath pulumi.StringInput
	// Specifies the name of the package owner. The owner must be IAM user.
	Owner pulumi.StringPtrInput
	// Specifies the region in which to upload packages.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will delete the current package and upload a new package.
	Region pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the package.
	Tags pulumi.StringMapInput
	// Specifies the package type.
	// + **jar**: `.jar` or jar related files.
	// + **pyFile**: `.py` or python related files.
	// + **file**: Other user files.
	// + **modelFile**: User AI model files.
	Type pulumi.StringInput
}

func (PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packageArgs)(nil)).Elem()
}

type PackageInput interface {
	pulumi.Input

	ToPackageOutput() PackageOutput
	ToPackageOutputWithContext(ctx context.Context) PackageOutput
}

func (*Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (i *Package) ToPackageOutput() PackageOutput {
	return i.ToPackageOutputWithContext(context.Background())
}

func (i *Package) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageOutput)
}

// PackageArrayInput is an input type that accepts PackageArray and PackageArrayOutput values.
// You can construct a concrete instance of `PackageArrayInput` via:
//
//	PackageArray{ PackageArgs{...} }
type PackageArrayInput interface {
	pulumi.Input

	ToPackageArrayOutput() PackageArrayOutput
	ToPackageArrayOutputWithContext(context.Context) PackageArrayOutput
}

type PackageArray []PackageInput

func (PackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Package)(nil)).Elem()
}

func (i PackageArray) ToPackageArrayOutput() PackageArrayOutput {
	return i.ToPackageArrayOutputWithContext(context.Background())
}

func (i PackageArray) ToPackageArrayOutputWithContext(ctx context.Context) PackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageArrayOutput)
}

// PackageMapInput is an input type that accepts PackageMap and PackageMapOutput values.
// You can construct a concrete instance of `PackageMapInput` via:
//
//	PackageMap{ "key": PackageArgs{...} }
type PackageMapInput interface {
	pulumi.Input

	ToPackageMapOutput() PackageMapOutput
	ToPackageMapOutputWithContext(context.Context) PackageMapOutput
}

type PackageMap map[string]PackageInput

func (PackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Package)(nil)).Elem()
}

func (i PackageMap) ToPackageMapOutput() PackageMapOutput {
	return i.ToPackageMapOutputWithContext(context.Background())
}

func (i PackageMap) ToPackageMapOutputWithContext(ctx context.Context) PackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageMapOutput)
}

type PackageOutput struct{ *pulumi.OutputState }

func (PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (o PackageOutput) ToPackageOutput() PackageOutput {
	return o
}

func (o PackageOutput) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return o
}

// Time when a queue is created.
func (o PackageOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the group name which the package belongs to.
// Changing this parameter will delete the current package and upload a new package.
func (o PackageOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Package) pulumi.StringPtrOutput { return v.GroupName }).(pulumi.StringPtrOutput)
}

// Specifies whether to upload resource packages in asynchronous mode.
// The default value is **false**. Changing this parameter will delete the current package and upload a new package.
func (o PackageOutput) IsAsync() pulumi.BoolOutput {
	return o.ApplyT(func(v *Package) pulumi.BoolOutput { return v.IsAsync }).(pulumi.BoolOutput)
}

// The package name.
func (o PackageOutput) ObjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.ObjectName }).(pulumi.StringOutput)
}

// Specifies the OBS storage path where the package is located.
// For example, `https://{bucket_name}.obs.{region}.myhuaweicloud.com/dli/packages/object_file.py`.
// Changing this parameter will delete the current package and upload a new package.
func (o PackageOutput) ObjectPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.ObjectPath }).(pulumi.StringOutput)
}

// Specifies the name of the package owner. The owner must be IAM user.
func (o PackageOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Specifies the region in which to upload packages.
// If omitted, the provider-level region will be used.
// Changing this parameter will delete the current package and upload a new package.
func (o PackageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Status of a package group to be uploaded.
func (o PackageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the key/value pairs to associate with the package.
func (o PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the package type.
// + **jar**: `.jar` or jar related files.
// + **pyFile**: `.py` or python related files.
// + **file**: Other user files.
// + **modelFile**: User AI model files.
func (o PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last time when the package configuration update has completed.
func (o PackageOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type PackageArrayOutput struct{ *pulumi.OutputState }

func (PackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Package)(nil)).Elem()
}

func (o PackageArrayOutput) ToPackageArrayOutput() PackageArrayOutput {
	return o
}

func (o PackageArrayOutput) ToPackageArrayOutputWithContext(ctx context.Context) PackageArrayOutput {
	return o
}

func (o PackageArrayOutput) Index(i pulumi.IntInput) PackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Package {
		return vs[0].([]*Package)[vs[1].(int)]
	}).(PackageOutput)
}

type PackageMapOutput struct{ *pulumi.OutputState }

func (PackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Package)(nil)).Elem()
}

func (o PackageMapOutput) ToPackageMapOutput() PackageMapOutput {
	return o
}

func (o PackageMapOutput) ToPackageMapOutputWithContext(ctx context.Context) PackageMapOutput {
	return o
}

func (o PackageMapOutput) MapIndex(k pulumi.StringInput) PackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Package {
		return vs[0].(map[string]*Package)[vs[1].(string)]
	}).(PackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PackageInput)(nil)).Elem(), &Package{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageArrayInput)(nil)).Elem(), PackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageMapInput)(nil)).Elem(), PackageMap{})
	pulumi.RegisterOutputType(PackageOutput{})
	pulumi.RegisterOutputType(PackageArrayOutput{})
	pulumi.RegisterOutputType(PackageMapOutput{})
}
