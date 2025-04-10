// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an MPC transcoding template group resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Mpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Mpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mpc.NewTranscodingTemplateGroup(ctx, "test", &Mpc.TranscodingTemplateGroupArgs{
//				Audio: &mpc.TranscodingTemplateGroupAudioArgs{
//					Bitrate:      pulumi.Int(0),
//					Channels:     pulumi.Int(2),
//					Codec:        pulumi.Int(2),
//					OutputPolicy: pulumi.String("transcode"),
//					SampleRate:   pulumi.Int(1),
//				},
//				DashSegmentDuration: pulumi.Int(5),
//				HlsSegmentDuration:  pulumi.Int(5),
//				LowBitrateHd:        pulumi.Bool(true),
//				OutputFormat:        pulumi.Int(1),
//				VideoCommon: &mpc.TranscodingTemplateGroupVideoCommonArgs{
//					BlackBarRemoval:       pulumi.Int(0),
//					Codec:                 pulumi.Int(2),
//					Fps:                   pulumi.Int(0),
//					Level:                 pulumi.Int(15),
//					MaxConsecutiveBframes: pulumi.Int(7),
//					MaxIframesInterval:    pulumi.Int(5),
//					OutputPolicy:          pulumi.String("transcode"),
//					Profile:               pulumi.Int(4),
//					Quality:               pulumi.Int(1),
//				},
//				Videos: mpc.TranscodingTemplateGroupVideoArray{
//					&mpc.TranscodingTemplateGroupVideoArgs{
//						Bitrate: pulumi.Int(0),
//						Height:  pulumi.Int(2160),
//						Width:   pulumi.Int(3840),
//					},
//				},
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
// MPC transcoding template groups can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup test 589e49809bb84447a759f6fa9aa19949
//
// ```
type TranscodingTemplateGroup struct {
	pulumi.CustomResourceState

	// Specifies the audio parameters. The object structure is documented below.
	Audio TranscodingTemplateGroupAudioPtrOutput `pulumi:"audio"`
	// Specifies the dash segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	DashSegmentDuration pulumi.IntPtrOutput `pulumi:"dashSegmentDuration"`
	// Specifies the HLS segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	HlsSegmentDuration pulumi.IntPtrOutput `pulumi:"hlsSegmentDuration"`
	// Specifies Whether to enable low bitrate HD. The default value is false.
	LowBitrateHd pulumi.BoolPtrOutput `pulumi:"lowBitrateHd"`
	// Specifies the name of a transcoding template group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the packaging type. Possible values are:
	// + **1**: HLS
	// + **2**: DASH
	// + **3**: HLS+DASH
	// + **4**: MP4
	// + **5**: MP3
	// + **6**: ADTS
	OutputFormat pulumi.IntOutput `pulumi:"outputFormat"`
	// The region in which to create the transcoding template group resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Indicates the IDs of templates in the template group
	TemplateIds pulumi.StringArrayOutput `pulumi:"templateIds"`
	// Specifies the video parameters.
	// The object structure is documented below.
	VideoCommon TranscodingTemplateGroupVideoCommonPtrOutput `pulumi:"videoCommon"`
	// Specifies the list of output video configurations.
	// The object structure is documented below.
	Videos TranscodingTemplateGroupVideoArrayOutput `pulumi:"videos"`
}

// NewTranscodingTemplateGroup registers a new resource with the given unique name, arguments, and options.
func NewTranscodingTemplateGroup(ctx *pulumi.Context,
	name string, args *TranscodingTemplateGroupArgs, opts ...pulumi.ResourceOption) (*TranscodingTemplateGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OutputFormat == nil {
		return nil, errors.New("invalid value for required argument 'OutputFormat'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TranscodingTemplateGroup
	err := ctx.RegisterResource("huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TranscodingTemplateGroup resources.
type transcodingTemplateGroupState struct {
	// Specifies the audio parameters. The object structure is documented below.
	Audio *TranscodingTemplateGroupAudio `pulumi:"audio"`
	// Specifies the dash segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	DashSegmentDuration *int `pulumi:"dashSegmentDuration"`
	// Specifies the HLS segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	HlsSegmentDuration *int `pulumi:"hlsSegmentDuration"`
	// Specifies Whether to enable low bitrate HD. The default value is false.
	LowBitrateHd *bool `pulumi:"lowBitrateHd"`
	// Specifies the name of a transcoding template group.
	Name *string `pulumi:"name"`
	// Specifies the packaging type. Possible values are:
	// + **1**: HLS
	// + **2**: DASH
	// + **3**: HLS+DASH
	// + **4**: MP4
	// + **5**: MP3
	// + **6**: ADTS
	OutputFormat *int `pulumi:"outputFormat"`
	// The region in which to create the transcoding template group resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Indicates the IDs of templates in the template group
	TemplateIds []string `pulumi:"templateIds"`
	// Specifies the video parameters.
	// The object structure is documented below.
	VideoCommon *TranscodingTemplateGroupVideoCommon `pulumi:"videoCommon"`
	// Specifies the list of output video configurations.
	// The object structure is documented below.
	Videos []TranscodingTemplateGroupVideo `pulumi:"videos"`
}

type TranscodingTemplateGroupState struct {
	// Specifies the audio parameters. The object structure is documented below.
	Audio TranscodingTemplateGroupAudioPtrInput
	// Specifies the dash segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	DashSegmentDuration pulumi.IntPtrInput
	// Specifies the HLS segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	HlsSegmentDuration pulumi.IntPtrInput
	// Specifies Whether to enable low bitrate HD. The default value is false.
	LowBitrateHd pulumi.BoolPtrInput
	// Specifies the name of a transcoding template group.
	Name pulumi.StringPtrInput
	// Specifies the packaging type. Possible values are:
	// + **1**: HLS
	// + **2**: DASH
	// + **3**: HLS+DASH
	// + **4**: MP4
	// + **5**: MP3
	// + **6**: ADTS
	OutputFormat pulumi.IntPtrInput
	// The region in which to create the transcoding template group resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Indicates the IDs of templates in the template group
	TemplateIds pulumi.StringArrayInput
	// Specifies the video parameters.
	// The object structure is documented below.
	VideoCommon TranscodingTemplateGroupVideoCommonPtrInput
	// Specifies the list of output video configurations.
	// The object structure is documented below.
	Videos TranscodingTemplateGroupVideoArrayInput
}

func (TranscodingTemplateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*transcodingTemplateGroupState)(nil)).Elem()
}

type transcodingTemplateGroupArgs struct {
	// Specifies the audio parameters. The object structure is documented below.
	Audio *TranscodingTemplateGroupAudio `pulumi:"audio"`
	// Specifies the dash segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	DashSegmentDuration *int `pulumi:"dashSegmentDuration"`
	// Specifies the HLS segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	HlsSegmentDuration *int `pulumi:"hlsSegmentDuration"`
	// Specifies Whether to enable low bitrate HD. The default value is false.
	LowBitrateHd *bool `pulumi:"lowBitrateHd"`
	// Specifies the name of a transcoding template group.
	Name *string `pulumi:"name"`
	// Specifies the packaging type. Possible values are:
	// + **1**: HLS
	// + **2**: DASH
	// + **3**: HLS+DASH
	// + **4**: MP4
	// + **5**: MP3
	// + **6**: ADTS
	OutputFormat int `pulumi:"outputFormat"`
	// The region in which to create the transcoding template group resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// Specifies the video parameters.
	// The object structure is documented below.
	VideoCommon *TranscodingTemplateGroupVideoCommon `pulumi:"videoCommon"`
	// Specifies the list of output video configurations.
	// The object structure is documented below.
	Videos []TranscodingTemplateGroupVideo `pulumi:"videos"`
}

// The set of arguments for constructing a TranscodingTemplateGroup resource.
type TranscodingTemplateGroupArgs struct {
	// Specifies the audio parameters. The object structure is documented below.
	Audio TranscodingTemplateGroupAudioPtrInput
	// Specifies the dash segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	DashSegmentDuration pulumi.IntPtrInput
	// Specifies the HLS segment duration, in second.\
	// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
	// The default value is `5`.
	HlsSegmentDuration pulumi.IntPtrInput
	// Specifies Whether to enable low bitrate HD. The default value is false.
	LowBitrateHd pulumi.BoolPtrInput
	// Specifies the name of a transcoding template group.
	Name pulumi.StringPtrInput
	// Specifies the packaging type. Possible values are:
	// + **1**: HLS
	// + **2**: DASH
	// + **3**: HLS+DASH
	// + **4**: MP4
	// + **5**: MP3
	// + **6**: ADTS
	OutputFormat pulumi.IntInput
	// The region in which to create the transcoding template group resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// Specifies the video parameters.
	// The object structure is documented below.
	VideoCommon TranscodingTemplateGroupVideoCommonPtrInput
	// Specifies the list of output video configurations.
	// The object structure is documented below.
	Videos TranscodingTemplateGroupVideoArrayInput
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

// Specifies the audio parameters. The object structure is documented below.
func (o TranscodingTemplateGroupOutput) Audio() TranscodingTemplateGroupAudioPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) TranscodingTemplateGroupAudioPtrOutput { return v.Audio }).(TranscodingTemplateGroupAudioPtrOutput)
}

// Specifies the dash segment duration, in second.\
// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
// The default value is `5`.
func (o TranscodingTemplateGroupOutput) DashSegmentDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.IntPtrOutput { return v.DashSegmentDuration }).(pulumi.IntPtrOutput)
}

// Specifies the HLS segment duration, in second.\
// The valid value is range from `2` to `10`, and it is used only when `outputFormat` is set to `1` or `3`.
// The default value is `5`.
func (o TranscodingTemplateGroupOutput) HlsSegmentDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.IntPtrOutput { return v.HlsSegmentDuration }).(pulumi.IntPtrOutput)
}

// Specifies Whether to enable low bitrate HD. The default value is false.
func (o TranscodingTemplateGroupOutput) LowBitrateHd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.BoolPtrOutput { return v.LowBitrateHd }).(pulumi.BoolPtrOutput)
}

// Specifies the name of a transcoding template group.
func (o TranscodingTemplateGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the packaging type. Possible values are:
// + **1**: HLS
// + **2**: DASH
// + **3**: HLS+DASH
// + **4**: MP4
// + **5**: MP3
// + **6**: ADTS
func (o TranscodingTemplateGroupOutput) OutputFormat() pulumi.IntOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.IntOutput { return v.OutputFormat }).(pulumi.IntOutput)
}

// The region in which to create the transcoding template group resource. If omitted,
// the provider-level region will be used. Changing this creates a new resource.
func (o TranscodingTemplateGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Indicates the IDs of templates in the template group
func (o TranscodingTemplateGroupOutput) TemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) pulumi.StringArrayOutput { return v.TemplateIds }).(pulumi.StringArrayOutput)
}

// Specifies the video parameters.
// The object structure is documented below.
func (o TranscodingTemplateGroupOutput) VideoCommon() TranscodingTemplateGroupVideoCommonPtrOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) TranscodingTemplateGroupVideoCommonPtrOutput { return v.VideoCommon }).(TranscodingTemplateGroupVideoCommonPtrOutput)
}

// Specifies the list of output video configurations.
// The object structure is documented below.
func (o TranscodingTemplateGroupOutput) Videos() TranscodingTemplateGroupVideoArrayOutput {
	return o.ApplyT(func(v *TranscodingTemplateGroup) TranscodingTemplateGroupVideoArrayOutput { return v.Videos }).(TranscodingTemplateGroupVideoArrayOutput)
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
