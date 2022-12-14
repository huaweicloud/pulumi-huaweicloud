// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb mysql configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDB.GetMysqlConfiguration(ctx, &gaussdb.GetMysqlConfigurationArgs{
//				Name: pulumi.StringRef("Default-GaussDB-for-MySQL 8.0"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMysqlConfiguration(ctx *pulumi.Context, args *GetMysqlConfigurationArgs, opts ...pulumi.InvokeOption) (*GetMysqlConfigurationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMysqlConfigurationResult
	err := ctx.Invoke("huaweicloud:GaussDB/getMysqlConfiguration:getMysqlConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlConfiguration.
type GetMysqlConfigurationArgs struct {
	// Specifies the name of the parameter template.
	Name *string `pulumi:"name"`
	// The region in which to obtain the configurations. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getMysqlConfiguration.
type GetMysqlConfigurationResult struct {
	// Indicates the datastore name of the configuration.
	DatastoreName string `pulumi:"datastoreName"`
	// Indicates the datastore version of the configuration.
	DatastoreVersion string `pulumi:"datastoreVersion"`
	// Indicates the description of the configuration.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
}

func GetMysqlConfigurationOutput(ctx *pulumi.Context, args GetMysqlConfigurationOutputArgs, opts ...pulumi.InvokeOption) GetMysqlConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMysqlConfigurationResult, error) {
			args := v.(GetMysqlConfigurationArgs)
			r, err := GetMysqlConfiguration(ctx, &args, opts...)
			var s GetMysqlConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMysqlConfigurationResultOutput)
}

// A collection of arguments for invoking getMysqlConfiguration.
type GetMysqlConfigurationOutputArgs struct {
	// Specifies the name of the parameter template.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the configurations. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetMysqlConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlConfiguration.
type GetMysqlConfigurationResultOutput struct{ *pulumi.OutputState }

func (GetMysqlConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlConfigurationResult)(nil)).Elem()
}

func (o GetMysqlConfigurationResultOutput) ToGetMysqlConfigurationResultOutput() GetMysqlConfigurationResultOutput {
	return o
}

func (o GetMysqlConfigurationResultOutput) ToGetMysqlConfigurationResultOutputWithContext(ctx context.Context) GetMysqlConfigurationResultOutput {
	return o
}

// Indicates the datastore name of the configuration.
func (o GetMysqlConfigurationResultOutput) DatastoreName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.DatastoreName }).(pulumi.StringOutput)
}

// Indicates the datastore version of the configuration.
func (o GetMysqlConfigurationResultOutput) DatastoreVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.DatastoreVersion }).(pulumi.StringOutput)
}

// Indicates the description of the configuration.
func (o GetMysqlConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMysqlConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMysqlConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetMysqlConfigurationResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlConfigurationResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMysqlConfigurationResultOutput{})
}
