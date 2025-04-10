// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicestage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource is used to manage a component under specified application within HuaweiCloud ServiceStage service.
//
// ## Example Usage
//
// ## Import
//
// Components can be imported using their `application_id` and `id`, separated by a slash (/), e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:ServiceStage/component:Component test dd7a1ce2-c48c-4f41-85bb-d0d09969eec9/9ab8ef79-d318-4de5-acf9-e1e1e25a0395
//
// ```
type Component struct {
	pulumi.CustomResourceState

	// Specifies the application ID to which the component belongs.
	// Changing this parameter will create a new component.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Specifies the component builder.
	// The object structure is documented below.
	Builder ComponentBuilderOutput `pulumi:"builder"`
	// Specifies the component framework.
	// + The framework of type **Webapp** is **Web**.
	// + The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
	//   **Dubbo**.
	// + The framework of type **Common** can be empty.
	Framework pulumi.StringPtrOutput `pulumi:"framework"`
	// Specifies the authorization name.
	// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
	// and the name must start with a letter and end with a letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region where the application and component are located.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the component runtime, such as **Docker**, **Java8**, etc.
	// Changing this parameter will create a new component.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// Specifies the repository source.
	// The object structure is documented below.
	Source ComponentSourceOutput `pulumi:"source"`
	// Specifies the type of repository source or storage.
	// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Component
	err := ctx.RegisterResource("huaweicloud:ServiceStage/component:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("huaweicloud:ServiceStage/component:Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
	// Specifies the application ID to which the component belongs.
	// Changing this parameter will create a new component.
	ApplicationId *string `pulumi:"applicationId"`
	// Specifies the component builder.
	// The object structure is documented below.
	Builder *ComponentBuilder `pulumi:"builder"`
	// Specifies the component framework.
	// + The framework of type **Webapp** is **Web**.
	// + The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
	//   **Dubbo**.
	// + The framework of type **Common** can be empty.
	Framework *string `pulumi:"framework"`
	// Specifies the authorization name.
	// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
	// and the name must start with a letter and end with a letter or digit.
	Name *string `pulumi:"name"`
	// Specifies the region where the application and component are located.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
	Region *string `pulumi:"region"`
	// Specifies the component runtime, such as **Docker**, **Java8**, etc.
	// Changing this parameter will create a new component.
	Runtime *string `pulumi:"runtime"`
	// Specifies the repository source.
	// The object structure is documented below.
	Source *ComponentSource `pulumi:"source"`
	// Specifies the type of repository source or storage.
	// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
	Type *string `pulumi:"type"`
}

type ComponentState struct {
	// Specifies the application ID to which the component belongs.
	// Changing this parameter will create a new component.
	ApplicationId pulumi.StringPtrInput
	// Specifies the component builder.
	// The object structure is documented below.
	Builder ComponentBuilderPtrInput
	// Specifies the component framework.
	// + The framework of type **Webapp** is **Web**.
	// + The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
	//   **Dubbo**.
	// + The framework of type **Common** can be empty.
	Framework pulumi.StringPtrInput
	// Specifies the authorization name.
	// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
	// and the name must start with a letter and end with a letter or digit.
	Name pulumi.StringPtrInput
	// Specifies the region where the application and component are located.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
	Region pulumi.StringPtrInput
	// Specifies the component runtime, such as **Docker**, **Java8**, etc.
	// Changing this parameter will create a new component.
	Runtime pulumi.StringPtrInput
	// Specifies the repository source.
	// The object structure is documented below.
	Source ComponentSourcePtrInput
	// Specifies the type of repository source or storage.
	// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
	Type pulumi.StringPtrInput
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	// Specifies the application ID to which the component belongs.
	// Changing this parameter will create a new component.
	ApplicationId string `pulumi:"applicationId"`
	// Specifies the component builder.
	// The object structure is documented below.
	Builder *ComponentBuilder `pulumi:"builder"`
	// Specifies the component framework.
	// + The framework of type **Webapp** is **Web**.
	// + The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
	//   **Dubbo**.
	// + The framework of type **Common** can be empty.
	Framework *string `pulumi:"framework"`
	// Specifies the authorization name.
	// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
	// and the name must start with a letter and end with a letter or digit.
	Name *string `pulumi:"name"`
	// Specifies the region where the application and component are located.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
	Region *string `pulumi:"region"`
	// Specifies the component runtime, such as **Docker**, **Java8**, etc.
	// Changing this parameter will create a new component.
	Runtime string `pulumi:"runtime"`
	// Specifies the repository source.
	// The object structure is documented below.
	Source *ComponentSource `pulumi:"source"`
	// Specifies the type of repository source or storage.
	// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	// Specifies the application ID to which the component belongs.
	// Changing this parameter will create a new component.
	ApplicationId pulumi.StringInput
	// Specifies the component builder.
	// The object structure is documented below.
	Builder ComponentBuilderPtrInput
	// Specifies the component framework.
	// + The framework of type **Webapp** is **Web**.
	// + The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
	//   **Dubbo**.
	// + The framework of type **Common** can be empty.
	Framework pulumi.StringPtrInput
	// Specifies the authorization name.
	// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
	// and the name must start with a letter and end with a letter or digit.
	Name pulumi.StringPtrInput
	// Specifies the region where the application and component are located.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
	Region pulumi.StringPtrInput
	// Specifies the component runtime, such as **Docker**, **Java8**, etc.
	// Changing this parameter will create a new component.
	Runtime pulumi.StringInput
	// Specifies the repository source.
	// The object structure is documented below.
	Source ComponentSourcePtrInput
	// Specifies the type of repository source or storage.
	// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
	Type pulumi.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

// ComponentArrayInput is an input type that accepts ComponentArray and ComponentArrayOutput values.
// You can construct a concrete instance of `ComponentArrayInput` via:
//
//	ComponentArray{ ComponentArgs{...} }
type ComponentArrayInput interface {
	pulumi.Input

	ToComponentArrayOutput() ComponentArrayOutput
	ToComponentArrayOutputWithContext(context.Context) ComponentArrayOutput
}

type ComponentArray []ComponentInput

func (ComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (i ComponentArray) ToComponentArrayOutput() ComponentArrayOutput {
	return i.ToComponentArrayOutputWithContext(context.Background())
}

func (i ComponentArray) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentArrayOutput)
}

// ComponentMapInput is an input type that accepts ComponentMap and ComponentMapOutput values.
// You can construct a concrete instance of `ComponentMapInput` via:
//
//	ComponentMap{ "key": ComponentArgs{...} }
type ComponentMapInput interface {
	pulumi.Input

	ToComponentMapOutput() ComponentMapOutput
	ToComponentMapOutputWithContext(context.Context) ComponentMapOutput
}

type ComponentMap map[string]ComponentInput

func (ComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (i ComponentMap) ToComponentMapOutput() ComponentMapOutput {
	return i.ToComponentMapOutputWithContext(context.Background())
}

func (i ComponentMap) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentMapOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

// Specifies the application ID to which the component belongs.
// Changing this parameter will create a new component.
func (o ComponentOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Specifies the component builder.
// The object structure is documented below.
func (o ComponentOutput) Builder() ComponentBuilderOutput {
	return o.ApplyT(func(v *Component) ComponentBuilderOutput { return v.Builder }).(ComponentBuilderOutput)
}

// Specifies the component framework.
//   - The framework of type **Webapp** is **Web**.
//   - The framework of type **MicroService** supports: **Java Classis**, **Go Classis**, **Mesher**, **Spring Cloud**,
//     **Dubbo**.
//   - The framework of type **Common** can be empty.
func (o ComponentOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.Framework }).(pulumi.StringPtrOutput)
}

// Specifies the authorization name.
// The name can contain of `2` to `64` characters, only letters, digits, underscores (_) and hyphens (-) are allowed,
// and the name must start with a letter and end with a letter or digit.
func (o ComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region where the application and component are located.
// If omitted, the provider-level region will be used. Changing this parameter will create a new component.
func (o ComponentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the component runtime, such as **Docker**, **Java8**, etc.
// Changing this parameter will create a new component.
func (o ComponentOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Runtime }).(pulumi.StringOutput)
}

// Specifies the repository source.
// The object structure is documented below.
func (o ComponentOutput) Source() ComponentSourceOutput {
	return o.ApplyT(func(v *Component) ComponentSourceOutput { return v.Source }).(ComponentSourceOutput)
}

// Specifies the type of repository source or storage.
// The valid values are **GitHub**, **GitLab**, **Gitee**, **Bitbucket** and **package**.
func (o ComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ComponentArrayOutput struct{ *pulumi.OutputState }

func (ComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (o ComponentArrayOutput) ToComponentArrayOutput() ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) Index(i pulumi.IntInput) ComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Component {
		return vs[0].([]*Component)[vs[1].(int)]
	}).(ComponentOutput)
}

type ComponentMapOutput struct{ *pulumi.OutputState }

func (ComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (o ComponentMapOutput) ToComponentMapOutput() ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) MapIndex(k pulumi.StringInput) ComponentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Component {
		return vs[0].(map[string]*Component)[vs[1].(string)]
	}).(ComponentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentArrayInput)(nil)).Elem(), ComponentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentMapInput)(nil)).Elem(), ComponentMap{})
	pulumi.RegisterOutputType(ComponentOutput{})
	pulumi.RegisterOutputType(ComponentArrayOutput{})
	pulumi.RegisterOutputType(ComponentMapOutput{})
}
