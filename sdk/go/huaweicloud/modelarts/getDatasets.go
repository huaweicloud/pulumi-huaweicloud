// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package modelarts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of ModelArts datasets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/ModelArts"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/ModelArts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ModelArts.GetDatasets(ctx, &modelarts.GetDatasetsArgs{
//				Name: pulumi.StringRef("your_dataset_name"),
//				Type: pulumi.IntRef(1),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDatasets(ctx *pulumi.Context, args *GetDatasetsArgs, opts ...pulumi.InvokeOption) (*GetDatasetsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatasetsResult
	err := ctx.Invoke("huaweicloud:ModelArts/getDatasets:getDatasets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatasets.
type GetDatasetsArgs struct {
	// Specifies the name of datasets.
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain datasets. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
	// Specifies the type of datasets. The options are:
	// + **0**: Image classification, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **1**: Object detection, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **3**: Image segmentation, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **100**: Text classification, supported formats: `.txt`, `.csv`.
	// + **200**: Sound classification, Supported formats: `.wav`.
	// + **400**: Table type, supported formats: Carbon type.
	// + **600**: Video, supported formats: `.mp4`
	// + **900**: Free format.
	Type *int `pulumi:"type"`
}

// A collection of values returned by getDatasets.
type GetDatasetsResult struct {
	// Indicates a list of all datasets found. Structure is documented below.
	Datasets []GetDatasetsDataset `pulumi:"datasets"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of label.
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
	// The field type. Valid values include: `String`, `Short`, `Int`, `Long`, `Double`, `Float`, `Byte`,
	// `Date`, `Timestamp`, `Bool`.
	Type *int `pulumi:"type"`
}

func GetDatasetsOutput(ctx *pulumi.Context, args GetDatasetsOutputArgs, opts ...pulumi.InvokeOption) GetDatasetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatasetsResult, error) {
			args := v.(GetDatasetsArgs)
			r, err := GetDatasets(ctx, &args, opts...)
			var s GetDatasetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatasetsResultOutput)
}

// A collection of arguments for invoking getDatasets.
type GetDatasetsOutputArgs struct {
	// Specifies the name of datasets.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain datasets. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the type of datasets. The options are:
	// + **0**: Image classification, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **1**: Object detection, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **3**: Image segmentation, supported formats: `.jpg`, `.png`, `.jpeg`, `.bmp`.
	// + **100**: Text classification, supported formats: `.txt`, `.csv`.
	// + **200**: Sound classification, Supported formats: `.wav`.
	// + **400**: Table type, supported formats: Carbon type.
	// + **600**: Video, supported formats: `.mp4`
	// + **900**: Free format.
	Type pulumi.IntPtrInput `pulumi:"type"`
}

func (GetDatasetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatasetsArgs)(nil)).Elem()
}

// A collection of values returned by getDatasets.
type GetDatasetsResultOutput struct{ *pulumi.OutputState }

func (GetDatasetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatasetsResult)(nil)).Elem()
}

func (o GetDatasetsResultOutput) ToGetDatasetsResultOutput() GetDatasetsResultOutput {
	return o
}

func (o GetDatasetsResultOutput) ToGetDatasetsResultOutputWithContext(ctx context.Context) GetDatasetsResultOutput {
	return o
}

// Indicates a list of all datasets found. Structure is documented below.
func (o GetDatasetsResultOutput) Datasets() GetDatasetsDatasetArrayOutput {
	return o.ApplyT(func(v GetDatasetsResult) []GetDatasetsDataset { return v.Datasets }).(GetDatasetsDatasetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatasetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatasetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of label.
func (o GetDatasetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatasetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetDatasetsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatasetsResult) string { return v.Region }).(pulumi.StringOutput)
}

// The field type. Valid values include: `String`, `Short`, `Int`, `Long`, `Double`, `Float`, `Byte`,
// `Date`, `Timestamp`, `Bool`.
func (o GetDatasetsResultOutput) Type() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDatasetsResult) *int { return v.Type }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatasetsResultOutput{})
}