// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VOD transcoding template group resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vod"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vod"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vod.NewTranscodingTemplateGroup(ctx, "test", &Vod.TranscodingTemplateGroupArgs{
//				AudioCodec:         pulumi.String("AAC"),
//				Description:        pulumi.String("test group"),
//				HlsSegmentDuration: pulumi.Int(5),
//				LowBitrateHd:       pulumi.Bool(false),
//				QualityInfos: vod.TranscodingTemplateGroupQualityInfoArray{
//					&vod.TranscodingTemplateGroupQualityInfoArgs{
//						Audio: &vod.TranscodingTemplateGroupQualityInfoAudioArgs{
//							Bitrate:    pulumi.Int(0),
//							Channels:   pulumi.Int(2),
//							SampleRate: pulumi.Int(1),
//						},
//						OutputFormat: pulumi.String("HLS"),
//						Video: &vod.TranscodingTemplateGroupQualityInfoVideoArgs{
//							Bitrate:   pulumi.Int(1000),
//							FrameRate: pulumi.Int(1),
//							Height:    pulumi.Int(720),
//							Quality:   pulumi.String("HD"),
//							Width:     pulumi.Int(1280),
//						},
//					},
//				},
//				VideoCodec: pulumi.String("H264"),
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
// VOD transcoding template groups can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vod/transcodingTemplateGroup:TranscodingTemplateGroup test 589e49809bb84447a759f6fa9aa19949
//
// ```
type TranscodingTemplateGroup struct {
	pulumi.CustomResourceState

	// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
	// Defaults to: **AAC**.
	AudioCodec pulumi.StringOutput `pulumi:"audioCodec"`
	// Specifies whether to automatically encrypt. Before enabling, you need to configure
	// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
	// Defaults to: **false**.
	AutoEncrypt pulumi.BoolPtrOutput `pulumi:"autoEncrypt"`
	// Specifies the description of the template group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
	// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
	HlsSegmentDuration pulumi.IntOutput `pulumi:"hlsSegmentDuration"`
	// Specifies whether to use this group as default group. Defaults to: **false**.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// Specifies whether to enable low bitrate HD. Defaults to: **false**.
	LowBitrateHd pulumi.BoolPtrOutput `pulumi:"lowBitrateHd"`
	// Specifies the name of the template group. The value can be a string of `1` to `128`
	// characters that can consist of letters, digits and underscores (_).
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the quality info list of the template group.
	// The object structure is documented below.
	QualityInfos TranscodingTemplateGroupQualityInfoArrayOutput `pulumi:"qualityInfos"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Indicates the type of the template group.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the video codec. The value can be: **H264** and **H265**.
	// Defaults to: **H264**.
	VideoCodec pulumi.StringOutput `pulumi:"videoCodec"`
	// Specifies the list of the watermark template IDs.
	WatermarkTemplateIds pulumi.StringArrayOutput `pulumi:"watermarkTemplateIds"`
}

// NewTranscodingTemplateGroup registers a new resource with the given unique name, arguments, and options.
func NewTranscodingTemplateGroup(ctx *pulumi.Context,
	name string, args *TranscodingTemplateGroupArgs, opts ...pulumi.ResourceOption) (*TranscodingTemplateGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QualityInfos == nil {
		return nil, errors.New("invalid value for required argument 'QualityInfos'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TranscodingTemplateGroup
	err := ctx.RegisterResource("huaweicloud:Vod/transcodingTemplateGroup:TranscodingTemplateGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTranscodingTemplateGroup gets an existing TranscodingTemplateGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTranscodingTemplateGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TranscodingTemplateGroupState, opts ...pulumi.ResourceOption) (*TranscodingTemplateGroup, error) {
	var resource TranscodingTemplateGroup
	err := ctx.ReadResource("huaweicloud:Vod/transcodingTemplateGroup:TranscodingTemplateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TranscodingTemplateGroup resources.
type transcodingTemplateGroupState struct {
	// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
	// Defaults to: **AAC**.
	AudioCodec *string `pulumi:"audioCodec"`
	// Specifies whether to automatically encrypt. Before enabling, you need to configure
	// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
	// Defaults to: **false**.
	AutoEncrypt *bool `pulumi:"autoEncrypt"`
	// Specifies the description of the template group.
	Description *string `pulumi:"description"`
	// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
	// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
	HlsSegmentDuration *int `pulumi:"hlsSegmentDuration"`
	// Specifies whether to use this group as default group. Defaults to: **false**.
	IsDefault *bool `pulumi:"isDefault"`
	// Specifies whether to enable low bitrate HD. Defaults to: **false**.
	LowBitrateHd *bool `pulumi:"lowBitrateHd"`
	// Specifies the name of the template group. The value can be a string of `1` to `128`
	// characters that can consist of letters, digits and underscores (_).
	Name *string `pulumi:"name"`
	// Specifies the quality info list of the template group.
	// The object structure is documented below.
	QualityInfos []TranscodingTemplateGroupQualityInfo `pulumi:"qualityInfos"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Indicates the type of the template group.
	Type *string `pulumi:"type"`
	// Specifies the video codec. The value can be: **H264** and **H265**.
	// Defaults to: **H264**.
	VideoCodec *string `pulumi:"videoCodec"`
	// Specifies the list of the watermark template IDs.
	WatermarkTemplateIds []string `pulumi:"watermarkTemplateIds"`
}

type TranscodingTemplateGroupState struct {
	// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
	// Defaults to: **AAC**.
	AudioCodec pulumi.StringPtrInput
	// Specifies whether to automatically encrypt. Before enabling, you need to configure
	// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
	// Defaults to: **false**.
	AutoEncrypt pulumi.BoolPtrInput
	// Specifies the description of the template group.
	Description pulumi.StringPtrInput
	// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
	// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
	HlsSegmentDuration pulumi.IntPtrInput
	// Specifies whether to use this group as default group. Defaults to: **false**.
	IsDefault pulumi.BoolPtrInput
	// Specifies whether to enable low bitrate HD. Defaults to: **false**.
	LowBitrateHd pulumi.BoolPtrInput
	// Specifies the name of the template group. The value can be a string of `1` to `128`
	// characters that can consist of letters, digits and underscores (_).
	Name pulumi.StringPtrInput
	// Specifies the quality info list of the template group.
	// The object structure is documented below.
	QualityInfos TranscodingTemplateGroupQualityInfoArrayInput
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Indicates the type of the template group.
	Type pulumi.StringPtrInput
	// Specifies the video codec. The value can be: **H264** and **H265**.
	// Defaults to: **H264**.
	VideoCodec pulumi.StringPtrInput
	// Specifies the list of the watermark template IDs.
	WatermarkTemplateIds pulumi.StringArrayInput
}

func (TranscodingTemplateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*transcodingTemplateGroupState)(nil)).Elem()
}

type transcodingTemplateGroupArgs struct {
	// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
	// Defaults to: **AAC**.
	AudioCodec *string `pulumi:"audioCodec"`
	// Specifies whether to automatically encrypt. Before enabling, you need to configure
	// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
	// Defaults to: **false**.
	AutoEncrypt *bool `pulumi:"autoEncrypt"`
	// Specifies the description of the template group.
	Description *string `pulumi:"description"`
	// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
	// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
	HlsSegmentDuration *int `pulumi:"hlsSegmentDuration"`
	// Specifies whether to use this group as default group. Defaults to: **false**.
	IsDefault *bool `pulumi:"isDefault"`
	// Specifies whether to enable low bitrate HD. Defaults to: **false**.
	LowBitrateHd *bool `pulumi:"lowBitrateHd"`
	// Specifies the name of the template group. The value can be a string of `1` to `128`
	// characters that can consist of letters, digits and underscores (_).
	Name *string `pulumi:"name"`
	// Specifies the quality info list of the template group.
	// The object structure is documented below.
	QualityInfos []TranscodingTemplateGroupQualityInfo `pulumi:"qualityInfos"`
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the video codec. The value can be: **H264** and **H265**.
	// Defaults to: **H264**.
	VideoCodec *string `pulumi:"videoCodec"`
	// Specifies the list of the watermark template IDs.
	WatermarkTemplateIds []string `pulumi:"watermarkTemplateIds"`
}

// The set of arguments for constructing a TranscodingTemplateGroup resource.
type TranscodingTemplateGroupArgs struct {
	// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
	// Defaults to: **AAC**.
	AudioCodec pulumi.StringPtrInput
	// Specifies whether to automatically encrypt. Before enabling, you need to configure
	// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
	// Defaults to: **false**.
	AutoEncrypt pulumi.BoolPtrInput
	// Specifies the description of the template group.
	Description pulumi.StringPtrInput
	// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
	// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
	HlsSegmentDuration pulumi.IntPtrInput
	// Specifies whether to use this group as default group. Defaults to: **false**.
	IsDefault pulumi.BoolPtrInput
	// Specifies whether to enable low bitrate HD. Defaults to: **false**.
	LowBitrateHd pulumi.BoolPtrInput
	// Specifies the name of the template group. The value can be a string of `1` to `128`
	// characters that can consist of letters, digits and underscores (_).
	Name pulumi.StringPtrInput
	// Specifies the quality info list of the template group.
	// The object structure is documented below.
	QualityInfos TranscodingTemplateGroupQualityInfoArrayInput
	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the video codec. The value can be: **H264** and **H265**.
	// Defaults to: **H264**.
	VideoCodec pulumi.StringPtrInput
	// Specifies the list of the watermark template IDs.
	WatermarkTemplateIds pulumi.StringArrayInput
}

func (TranscodingTemplateGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transcodingTemplateGroupArgs)(nil)).Elem()
}

type TranscodingTemplateGroupInput interface {
	pulumi.Input

	ToTranscodingTemplateGroupOutput() TranscodingTemplateGroupOutput
	ToTranscodingTemplateGroupOutputWithContext(ctx context.Context) TranscodingTemplateGroupOutput
}

func (*TranscodingTemplateGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TranscodingTemplateGroup)(nil)).Elem()
}

func (i *TranscodingTemplateGroup) ToTranscodingTemplateGroupOutput() TranscodingTemplateGroupOutput {
	return i.ToTranscodingTemplateGroupOutputWithContext(context.Background())
}

func (i *TranscodingTemplateGroup) ToTranscodingTemplateGroupOutputWithContext(ctx context.Context) TranscodingTemplateGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TranscodingTemplateGroupOutput)
}

// TranscodingTemplateGroupArrayInput is an input type that accepts TranscodingTemplateGroupArray and TranscodingTemplateGroupArrayOutput values.
// You can construct a concrete instance of `TranscodingTemplateGroupArrayInput` via:
//
//	TranscodingTemplateGroupArray{ TranscodingTemplateGroupArgs{...} }
type TranscodingTemplateGroupArrayInput interface {
	pulumi.Input

	ToTranscodingTemplateGroupArrayOutput() TranscodingTemplateGroupArrayOutput
	ToTranscodingTemplateGroupArrayOutputWithContext(context.Context) TranscodingTemplateGroupArrayOutput
}

type TranscodingTemplateGroupArray []TranscodingTemplateGroupInput

func (TranscodingTemplateGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TranscodingTemplateGroup)(nil)).Elem()
}

func (i TranscodingTemplateGroupArray) ToTranscodingTemplateGroupArrayOutput() TranscodingTemplateGroupArrayOutput {
	return i.ToTranscodingTemplateGroupArrayOutputWithContext(context.Background())
}

func (i TranscodingTemplateGroupArray) ToTranscodingTemplateGroupArrayOutputWithContext(ctx context.Context) TranscodingTemplateGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TranscodingTemplateGroupArrayOutput)
}

// TranscodingTemplateGroupMapInput is an input type that accepts TranscodingTemplateGroupMap and TranscodingTemplateGroupMapOutput values.
// You can construct a concrete instance of `TranscodingTemplateGroupMapInput` via:
//
//	TranscodingTemplateGroupMap{ "key": TranscodingTemplateGroupArgs{...} }
type TranscodingTemplateGroupMapInput interface {
	pulumi.Input

	ToTranscodingTemplateGroupMapOutput() TranscodingTemplateGroupMapOutput
	ToTranscodingTemplateGroupMapOutputWithContext(context.Context) TranscodingTemplateGroupMapOutput
}

type TranscodingTemplateGroupMap map[string]TranscodingTemplateGroupInput

func (TranscodingTemplateGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TranscodingTemplateGroup)(nil)).Elem()
}

func (i TranscodingTemplateGroupMap) ToTranscodingTemplateGroupMapOutput() TranscodingTemplateGroupMapOutput {
	return i.ToTranscodingTemplateGroupMapOutputWithContext(context.Background())
}

func (i TranscodingTemplateGroupMap) ToTranscodingTemplateGroupMapOutputWithContext(ctx context.Context) TranscodingTemplateGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TranscodingTemplateGroupMapOutput)
}

type TranscodingTemplateGroupOutput struct{ *pulumi.OutputState }

func (TranscodingTemplateGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TranscodingTemplateGroup)(nil)).Elem()
}

func (o TranscodingTemplateGroupOutput) ToTranscodingTemplateGroupOutput() TranscodingTemplateGroupOutput {
	return o
}

func (o TranscodingTemplateGroupOutput) ToTranscodingTemplateGroupOutputWithContext(ctx context.Context) TranscodingTemplateGroupOutput {
	return o
}

// Specifies the audio codec. The value can be: **AAC** and **HEAAC1**.
// Defaults to: **AAC**.
func (o TranscodingTemplateGroupOutput) AudioCodec() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.AudioCodec }).(pulumi.StringOutput)
}

// Specifies whether to automatically encrypt. Before enabling, you need to configure
// the HLS encryption key URL. When `autoEncrypt` is **true**, the `outputFormat` must be **HLS**.
// Defaults to: **false**.
func (o TranscodingTemplateGroupOutput) AutoEncrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.BoolPtrOutput { return v.AutoEncrypt }).(pulumi.BoolPtrOutput)
}

// Specifies the description of the template group.
func (o TranscodingTemplateGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the HLS segment duration. The value can be: `2`, `3`, `5`
// and `10`. Defaults to: `5`. This parameter is used only when `outputFormat` is set to **HLS** or **DASH_HLS**.
func (o TranscodingTemplateGroupOutput) HlsSegmentDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.IntOutput { return v.HlsSegmentDuration }).(pulumi.IntOutput)
}

// Specifies whether to use this group as default group. Defaults to: **false**.
func (o TranscodingTemplateGroupOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// Specifies whether to enable low bitrate HD. Defaults to: **false**.
func (o TranscodingTemplateGroupOutput) LowBitrateHd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.BoolPtrOutput { return v.LowBitrateHd }).(pulumi.BoolPtrOutput)
}

// Specifies the name of the template group. The value can be a string of `1` to `128`
// characters that can consist of letters, digits and underscores (_).
func (o TranscodingTemplateGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the quality info list of the template group.
// The object structure is documented below.
func (o TranscodingTemplateGroupOutput) QualityInfos() TranscodingTemplateGroupQualityInfoArrayOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) TranscodingTemplateGroupQualityInfoArrayOutput {
		return v.QualityInfos
	}).(TranscodingTemplateGroupQualityInfoArrayOutput)
}

// Specifies the region in which to create the resource. If omitted, the
// provider-level region will be used. Changing this creates a new resource.
func (o TranscodingTemplateGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Indicates the type of the template group.
func (o TranscodingTemplateGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the video codec. The value can be: **H264** and **H265**.
// Defaults to: **H264**.
func (o TranscodingTemplateGroupOutput) VideoCodec() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.VideoCodec }).(pulumi.StringOutput)
}

// Specifies the list of the watermark template IDs.
func (o TranscodingTemplateGroupOutput) WatermarkTemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringArrayOutput { return v.WatermarkTemplateIds }).(pulumi.StringArrayOutput)
}

type TranscodingTemplateGroupArrayOutput struct{ *pulumi.OutputState }

func (TranscodingTemplateGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TranscodingTemplateGroup)(nil)).Elem()
}

func (o TranscodingTemplateGroupArrayOutput) ToTranscodingTemplateGroupArrayOutput() TranscodingTemplateGroupArrayOutput {
	return o
}

func (o TranscodingTemplateGroupArrayOutput) ToTranscodingTemplateGroupArrayOutputWithContext(ctx context.Context) TranscodingTemplateGroupArrayOutput {
	return o
}

func (o TranscodingTemplateGroupArrayOutput) Index(i pulumi.IntInput) TranscodingTemplateGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TranscodingTemplateGroup {
		return vs[0].([]*TranscodingTemplateGroup)[vs[1].(int)]
	}).(TranscodingTemplateGroupOutput)
}

type TranscodingTemplateGroupMapOutput struct{ *pulumi.OutputState }

func (TranscodingTemplateGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TranscodingTemplateGroup)(nil)).Elem()
}

func (o TranscodingTemplateGroupMapOutput) ToTranscodingTemplateGroupMapOutput() TranscodingTemplateGroupMapOutput {
	return o
}

func (o TranscodingTemplateGroupMapOutput) ToTranscodingTemplateGroupMapOutputWithContext(ctx context.Context) TranscodingTemplateGroupMapOutput {
	return o
}

func (o TranscodingTemplateGroupMapOutput) MapIndex(k pulumi.StringInput) TranscodingTemplateGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TranscodingTemplateGroup {
		return vs[0].(map[string]*TranscodingTemplateGroup)[vs[1].(string)]
	}).(TranscodingTemplateGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TranscodingTemplateGroupInput)(nil)).Elem(), &TranscodingTemplateGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TranscodingTemplateGroupArrayInput)(nil)).Elem(), TranscodingTemplateGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TranscodingTemplateGroupMapInput)(nil)).Elem(), TranscodingTemplateGroupMap{})
	pulumi.RegisterOutputType(TranscodingTemplateGroupOutput{})
	pulumi.RegisterOutputType(TranscodingTemplateGroupArrayOutput{})
	pulumi.RegisterOutputType(TranscodingTemplateGroupMapOutput{})
}
