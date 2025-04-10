// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of RDS PostgreSQL accounts.
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
//			_, err := Rds.GetPgAccounts(ctx, &rds.GetPgAccountsArgs{
//				InstanceId: instanceId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPgAccounts(ctx *pulumi.Context, args *GetPgAccountsArgs, opts ...pulumi.InvokeOption) (*GetPgAccountsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPgAccountsResult
	err := ctx.Invoke("huaweicloud:Rds/getPgAccounts:getPgAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPgAccounts.
type GetPgAccountsArgs struct {
	// Specifies the PostgreSQL instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies the username of the DB account.
	UserName *string `pulumi:"userName"`
}

// A collection of values returned by getPgAccounts.
type GetPgAccountsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	Region     string  `pulumi:"region"`
	UserName   *string `pulumi:"userName"`
	// Indicates the user list.
	// The users structure is documented below.
	Users []GetPgAccountsUser `pulumi:"users"`
}

func GetPgAccountsOutput(ctx *pulumi.Context, args GetPgAccountsOutputArgs, opts ...pulumi.InvokeOption) GetPgAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPgAccountsResult, error) {
			args := v.(GetPgAccountsArgs)
			r, err := GetPgAccounts(ctx, &args, opts...)
			var s GetPgAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPgAccountsResultOutput)
}

// A collection of arguments for invoking getPgAccounts.
type GetPgAccountsOutputArgs struct {
	// Specifies the PostgreSQL instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the username of the DB account.
	UserName pulumi.StringPtrInput `pulumi:"userName"`
}

func (GetPgAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPgAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getPgAccounts.
type GetPgAccountsResultOutput struct{ *pulumi.OutputState }

func (GetPgAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPgAccountsResult)(nil)).Elem()
}

func (o GetPgAccountsResultOutput) ToGetPgAccountsResultOutput() GetPgAccountsResultOutput {
	return o
}

func (o GetPgAccountsResultOutput) ToGetPgAccountsResultOutputWithContext(ctx context.Context) GetPgAccountsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPgAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPgAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPgAccountsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPgAccountsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetPgAccountsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetPgAccountsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetPgAccountsResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPgAccountsResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

// Indicates the user list.
// The users structure is documented below.
func (o GetPgAccountsResultOutput) Users() GetPgAccountsUserArrayOutput {
	return o.ApplyT(func(v GetPgAccountsResult) []GetPgAccountsUser { return v.Users }).(GetPgAccountsUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPgAccountsResultOutput{})
}
