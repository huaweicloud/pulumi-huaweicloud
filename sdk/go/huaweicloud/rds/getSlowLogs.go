// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of RDS slow logs.
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
//			startTime := cfg.RequireObject("startTime")
//			endTime := cfg.RequireObject("endTime")
//			_, err := Rds.GetSlowLogs(ctx, &rds.GetSlowLogsArgs{
//				InstanceId: instanceId,
//				StartTime:  startTime,
//				EndTime:    endTime,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSlowLogs(ctx *pulumi.Context, args *GetSlowLogsArgs, opts ...pulumi.InvokeOption) (*GetSlowLogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowLogsResult
	err := ctx.Invoke("huaweicloud:Rds/getSlowLogs:getSlowLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogs.
type GetSlowLogsArgs struct {
	// Specifies the name of the database.
	Database *string `pulumi:"database"`
	// Specifies the end time in the **yyyy-mm-ddThh:mm:ssZ** format.
	EndTime string `pulumi:"endTime"`
	// Specifies the ID of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies the start time in the **yyyy-mm-ddThh:mm:ssZ** format.
	StartTime string `pulumi:"startTime"`
	// Specifies the statement type. Value options: **INSERT**, **UPDATE**, **SELECT**,
	// **DELETE**, **CREATE**.
	Type *string `pulumi:"type"`
	// Specifies the name of the account.
	Users *string `pulumi:"users"`
}

// A collection of values returned by getSlowLogs.
type GetSlowLogsResult struct {
	// Indicates the name of the database.
	Database *string `pulumi:"database"`
	EndTime  string  `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	Region     string `pulumi:"region"`
	// Indicates the list of the slow logs.
	SlowLogs []GetSlowLogsSlowLog `pulumi:"slowLogs"`
	// Indicates the start time in the **yyyy-mm-ddThh:mm:ssZ** format.
	StartTime string `pulumi:"startTime"`
	// Indicates the statement type.
	Type *string `pulumi:"type"`
	// Indicates the name of the account.
	Users *string `pulumi:"users"`
}

func GetSlowLogsOutput(ctx *pulumi.Context, args GetSlowLogsOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogsResult, error) {
			args := v.(GetSlowLogsArgs)
			r, err := GetSlowLogs(ctx, &args, opts...)
			var s GetSlowLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogsResultOutput)
}

// A collection of arguments for invoking getSlowLogs.
type GetSlowLogsOutputArgs struct {
	// Specifies the name of the database.
	Database pulumi.StringPtrInput `pulumi:"database"`
	// Specifies the end time in the **yyyy-mm-ddThh:mm:ssZ** format.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Specifies the ID of the RDS instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Specifies the region in which to query the resource.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the start time in the **yyyy-mm-ddThh:mm:ssZ** format.
	StartTime pulumi.StringInput `pulumi:"startTime"`
	// Specifies the statement type. Value options: **INSERT**, **UPDATE**, **SELECT**,
	// **DELETE**, **CREATE**.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Specifies the name of the account.
	Users pulumi.StringPtrInput `pulumi:"users"`
}

func (GetSlowLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogs.
type GetSlowLogsResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogsResult)(nil)).Elem()
}

func (o GetSlowLogsResultOutput) ToGetSlowLogsResultOutput() GetSlowLogsResultOutput {
	return o
}

func (o GetSlowLogsResultOutput) ToGetSlowLogsResultOutputWithContext(ctx context.Context) GetSlowLogsResultOutput {
	return o
}

// Indicates the name of the database.
func (o GetSlowLogsResultOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSlowLogsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the list of the slow logs.
func (o GetSlowLogsResultOutput) SlowLogs() GetSlowLogsSlowLogArrayOutput {
	return o.ApplyT(func(v GetSlowLogsResult) []GetSlowLogsSlowLog { return v.SlowLogs }).(GetSlowLogsSlowLogArrayOutput)
}

// Indicates the start time in the **yyyy-mm-ddThh:mm:ssZ** format.
func (o GetSlowLogsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Indicates the statement type.
func (o GetSlowLogsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Indicates the name of the account.
func (o GetSlowLogsResultOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogsResult) *string { return v.Users }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogsResultOutput{})
}
