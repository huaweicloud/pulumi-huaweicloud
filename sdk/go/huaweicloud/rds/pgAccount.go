// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages RDS PostgreSQL account resource within HuaweiCloud.
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
//			accountPassword := cfg.RequireObject("accountPassword")
//			_, err := Rds.NewPgAccount(ctx, "test", &Rds.PgAccountArgs{
//				InstanceId: pulumi.Any(instanceId),
//				Password:   pulumi.Any(accountPassword),
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
// The RDS PostgreSQL account can be imported using the `instance_id` and `name` separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Rds/pgAccount:PgAccount test <instance_id>/<name>
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`. It is generally recommended running `terraform plan` after importing the RDS PostgreSQL account. You can then decide if changes should be applied to the RDS PostgreSQL account, or the resource definition should be updated to align with the RDS PostgreSQL account. Also you can ignore changes as below. hcl resource "huaweicloud_rds_pg_account" "account_1" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	password
//
//	]
//
//	} }
type PgAccount struct {
	pulumi.CustomResourceState

	// Indicates the permission attributes of a user.
	// The attributes structure is documented below.
	Attributes PgAccountAttributeArrayOutput `pulumi:"attributes"`
	// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// schema: Deprecated
	Memberofs pulumi.StringArrayOutput `pulumi:"memberofs"`
	// Specifies the username of the DB account. The username contains 1 to 63
	// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
	// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
	// and **rdsDdm**.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the password of the DB account. The value must be 8 to 32 characters long
	// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
	// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
	Password pulumi.StringOutput `pulumi:"password"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPgAccount registers a new resource with the given unique name, arguments, and options.
func NewPgAccount(ctx *pulumi.Context,
	name string, args *PgAccountArgs, opts ...pulumi.ResourceOption) (*PgAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PgAccount
	err := ctx.RegisterResource("huaweicloud:Rds/pgAccount:PgAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgAccount gets an existing PgAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgAccountState, opts ...pulumi.ResourceOption) (*PgAccount, error) {
	var resource PgAccount
	err := ctx.ReadResource("huaweicloud:Rds/pgAccount:PgAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgAccount resources.
type pgAccountState struct {
	// Indicates the permission attributes of a user.
	// The attributes structure is documented below.
	Attributes []PgAccountAttribute `pulumi:"attributes"`
	// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
	Description *string `pulumi:"description"`
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId *string `pulumi:"instanceId"`
	// schema: Deprecated
	Memberofs []string `pulumi:"memberofs"`
	// Specifies the username of the DB account. The username contains 1 to 63
	// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
	// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
	// and **rdsDdm**.
	Name *string `pulumi:"name"`
	// Specifies the password of the DB account. The value must be 8 to 32 characters long
	// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
	// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
	Password *string `pulumi:"password"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

type PgAccountState struct {
	// Indicates the permission attributes of a user.
	// The attributes structure is documented below.
	Attributes PgAccountAttributeArrayInput
	// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
	Description pulumi.StringPtrInput
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringPtrInput
	// schema: Deprecated
	Memberofs pulumi.StringArrayInput
	// Specifies the username of the DB account. The username contains 1 to 63
	// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
	// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
	// and **rdsDdm**.
	Name pulumi.StringPtrInput
	// Specifies the password of the DB account. The value must be 8 to 32 characters long
	// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
	// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
	Password pulumi.StringPtrInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (PgAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountState)(nil)).Elem()
}

type pgAccountArgs struct {
	// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
	Description *string `pulumi:"description"`
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId string `pulumi:"instanceId"`
	// schema: Deprecated
	Memberofs []string `pulumi:"memberofs"`
	// Specifies the username of the DB account. The username contains 1 to 63
	// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
	// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
	// and **rdsDdm**.
	Name *string `pulumi:"name"`
	// Specifies the password of the DB account. The value must be 8 to 32 characters long
	// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
	// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
	Password string `pulumi:"password"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a PgAccount resource.
type PgAccountArgs struct {
	// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
	Description pulumi.StringPtrInput
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceId pulumi.StringInput
	// schema: Deprecated
	Memberofs pulumi.StringArrayInput
	// Specifies the username of the DB account. The username contains 1 to 63
	// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
	// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
	// and **rdsDdm**.
	Name pulumi.StringPtrInput
	// Specifies the password of the DB account. The value must be 8 to 32 characters long
	// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
	// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
	Password pulumi.StringInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (PgAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgAccountArgs)(nil)).Elem()
}

type PgAccountInput interface {
	pulumi.Input

	ToPgAccountOutput() PgAccountOutput
	ToPgAccountOutputWithContext(ctx context.Context) PgAccountOutput
}

func (*PgAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccount)(nil)).Elem()
}

func (i *PgAccount) ToPgAccountOutput() PgAccountOutput {
	return i.ToPgAccountOutputWithContext(context.Background())
}

func (i *PgAccount) ToPgAccountOutputWithContext(ctx context.Context) PgAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountOutput)
}

// PgAccountArrayInput is an input type that accepts PgAccountArray and PgAccountArrayOutput values.
// You can construct a concrete instance of `PgAccountArrayInput` via:
//
//	PgAccountArray{ PgAccountArgs{...} }
type PgAccountArrayInput interface {
	pulumi.Input

	ToPgAccountArrayOutput() PgAccountArrayOutput
	ToPgAccountArrayOutputWithContext(context.Context) PgAccountArrayOutput
}

type PgAccountArray []PgAccountInput

func (PgAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccount)(nil)).Elem()
}

func (i PgAccountArray) ToPgAccountArrayOutput() PgAccountArrayOutput {
	return i.ToPgAccountArrayOutputWithContext(context.Background())
}

func (i PgAccountArray) ToPgAccountArrayOutputWithContext(ctx context.Context) PgAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountArrayOutput)
}

// PgAccountMapInput is an input type that accepts PgAccountMap and PgAccountMapOutput values.
// You can construct a concrete instance of `PgAccountMapInput` via:
//
//	PgAccountMap{ "key": PgAccountArgs{...} }
type PgAccountMapInput interface {
	pulumi.Input

	ToPgAccountMapOutput() PgAccountMapOutput
	ToPgAccountMapOutputWithContext(context.Context) PgAccountMapOutput
}

type PgAccountMap map[string]PgAccountInput

func (PgAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccount)(nil)).Elem()
}

func (i PgAccountMap) ToPgAccountMapOutput() PgAccountMapOutput {
	return i.ToPgAccountMapOutputWithContext(context.Background())
}

func (i PgAccountMap) ToPgAccountMapOutputWithContext(ctx context.Context) PgAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgAccountMapOutput)
}

type PgAccountOutput struct{ *pulumi.OutputState }

func (PgAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgAccount)(nil)).Elem()
}

func (o PgAccountOutput) ToPgAccountOutput() PgAccountOutput {
	return o
}

func (o PgAccountOutput) ToPgAccountOutputWithContext(ctx context.Context) PgAccountOutput {
	return o
}

// Indicates the permission attributes of a user.
// The attributes structure is documented below.
func (o PgAccountOutput) Attributes() PgAccountAttributeArrayOutput {
	return o.ApplyT(func(v *PgAccount) PgAccountAttributeArrayOutput { return v.Attributes }).(PgAccountAttributeArrayOutput)
}

// Specifies the remarks of the DB account. The parameter must be 1 to 512 characters.
func (o PgAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the ID of the RDS PostgreSQL instance.
func (o PgAccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// schema: Deprecated
func (o PgAccountOutput) Memberofs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringArrayOutput { return v.Memberofs }).(pulumi.StringArrayOutput)
}

// Specifies the username of the DB account. The username contains 1 to 63
// characters, including letters, digits, and underscores (_). It cannot start with pg or a digit and must be different
// from system usernames. System users include **rdsAdmin**, **rdsMetric**, **rdsBackup**, **rdsRepl**, **rdsProxy**,
// and **rdsDdm**.
func (o PgAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the password of the DB account. The value must be 8 to 32 characters long
// and contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special
// characters (~!@#%^*-_=+?,). The value cannot contain the username or the username spelled backwards.
func (o PgAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o PgAccountOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PgAccount) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type PgAccountArrayOutput struct{ *pulumi.OutputState }

func (PgAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgAccount)(nil)).Elem()
}

func (o PgAccountArrayOutput) ToPgAccountArrayOutput() PgAccountArrayOutput {
	return o
}

func (o PgAccountArrayOutput) ToPgAccountArrayOutputWithContext(ctx context.Context) PgAccountArrayOutput {
	return o
}

func (o PgAccountArrayOutput) Index(i pulumi.IntInput) PgAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgAccount {
		return vs[0].([]*PgAccount)[vs[1].(int)]
	}).(PgAccountOutput)
}

type PgAccountMapOutput struct{ *pulumi.OutputState }

func (PgAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgAccount)(nil)).Elem()
}

func (o PgAccountMapOutput) ToPgAccountMapOutput() PgAccountMapOutput {
	return o
}

func (o PgAccountMapOutput) ToPgAccountMapOutputWithContext(ctx context.Context) PgAccountMapOutput {
	return o
}

func (o PgAccountMapOutput) MapIndex(k pulumi.StringInput) PgAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgAccount {
		return vs[0].(map[string]*PgAccount)[vs[1].(string)]
	}).(PgAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountInput)(nil)).Elem(), &PgAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountArrayInput)(nil)).Elem(), PgAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgAccountMapInput)(nil)).Elem(), PgAccountMap{})
	pulumi.RegisterOutputType(PgAccountOutput{})
	pulumi.RegisterOutputType(PgAccountArrayOutput{})
	pulumi.RegisterOutputType(PgAccountMapOutput{})
}
