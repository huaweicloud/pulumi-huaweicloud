// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of RDS MySQL accounts.
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Rds.GetMysqlAccounts(ctx, &rds.GetMysqlAccountsArgs{
//				InstanceId: _var.Instance_id,
//				Name:       pulumi.StringRef("test"),
//				Host:       pulumi.StringRef("10.10.10.10"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMysqlAccounts(ctx *pulumi.Context, args *GetMysqlAccountsArgs, opts ...pulumi.InvokeOption) (*GetMysqlAccountsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMysqlAccountsResult
	err := ctx.Invoke("huaweicloud:Rds/getMysqlAccounts:getMysqlAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlAccounts.
type GetMysqlAccountsArgs struct {
	// Specifies the IP address that is allowed to access your DB instance.
	Host *string `pulumi:"host"`
	// Specifies the ID of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the username of the DB account.
	Name *string `pulumi:"name"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getMysqlAccounts.
type GetMysqlAccountsResult struct {
	Host *string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Indicates the username of the DB account.
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
	// Indicates the list of users.
	// The users structure is documented below.
	Users []GetMysqlAccountsUser `pulumi:"users"`
}

func GetMysqlAccountsOutput(ctx *pulumi.Context, args GetMysqlAccountsOutputArgs, opts ...pulumi.InvokeOption) GetMysqlAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMysqlAccountsResult, error) {
			args := v.(GetMysqlAccountsArgs)
			r, err := GetMysqlAccounts(ctx, &args, opts...)
			var s GetMysqlAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMysqlAccountsResultOutput)
}

// A collection of arguments for invoking getMysqlAccounts.
type GetMysqlAccountsOutputArgs struct {
	// Specifies the IP address that is allowed to access your DB instance.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// Specifies the ID of the RDS instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Specifies the username of the DB account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetMysqlAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlAccounts.
type GetMysqlAccountsResultOutput struct{ *pulumi.OutputState }

func (GetMysqlAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlAccountsResult)(nil)).Elem()
}

func (o GetMysqlAccountsResultOutput) ToGetMysqlAccountsResultOutput() GetMysqlAccountsResultOutput {
	return o
}

func (o GetMysqlAccountsResultOutput) ToGetMysqlAccountsResultOutputWithContext(ctx context.Context) GetMysqlAccountsResultOutput {
	return o
}

func (o GetMysqlAccountsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMysqlAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMysqlAccountsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicates the username of the DB account.
func (o GetMysqlAccountsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetMysqlAccountsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the list of users.
// The users structure is documented below.
func (o GetMysqlAccountsResultOutput) Users() GetMysqlAccountsUserArrayOutput {
	return o.ApplyT(func(v GetMysqlAccountsResult) []GetMysqlAccountsUser { return v.Users }).(GetMysqlAccountsUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMysqlAccountsResultOutput{})
}
