// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package obs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OBS bucket object resource.
//
// ## Example Usage
// ### Uploading to a bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Obs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Obs.NewBucketObject(ctx, "object", &Obs.BucketObjectArgs{
//				Bucket:      pulumi.String("your_bucket_name"),
//				Content:     pulumi.String("some object content"),
//				ContentType: pulumi.String("application/xml"),
//				Key:         pulumi.String("new_key_from_content"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Uploading a file to a bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Obs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplebucket, err := Obs.NewBucket(ctx, "examplebucket", &Obs.BucketArgs{
//				Bucket: pulumi.String("examplebuckettftest"),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Obs.NewBucketObject(ctx, "object", &Obs.BucketObjectArgs{
//				Bucket: examplebucket.Bucket,
//				Key:    pulumi.String("new_key_from_file"),
//				Source: pulumi.String("index.html"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Server Side Encryption with OBS Default Master Key
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Obs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Obs.NewBucketObject(ctx, "examplebucketObject", &Obs.BucketObjectArgs{
//				Bucket:     pulumi.String("your_bucket_name"),
//				Encryption: pulumi.Bool(true),
//				Key:        pulumi.String("someobject"),
//				Source:     pulumi.String("index.html"),
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
// OBS bucket object can be imported using the bucket and key separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Obs/bucketObject:BucketObject object bucket/key
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`encryption`, `source`, `acl` and `kms_key_id`. It is generally recommended running `terraform plan` after importing an object. You can then decide if changes should be applied to the object, or the resource definition should be updated to align with the object. Also you can ignore changes as below. resource "huaweicloud_obs_bucket_object" "object" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	encryption, source, acl, kms_key_id,
//
//	]
//
//	} }
type BucketObject struct {
	pulumi.CustomResourceState

	// The ACL policy to apply. Defaults to `private`.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption pulumi.BoolPtrOutput `pulumi:"encryption"`
	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is `md5(file("pathToFile"))`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the kms key. If omitted, the default master key will be used.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// the size of the object in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// Specifies the storage class of the object. Defaults to `STANDARD`.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BucketObject
	err := ctx.RegisterResource("huaweicloud:Obs/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("huaweicloud:Obs/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The ACL policy to apply. Defaults to `private`.
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket *string `pulumi:"bucket"`
	// The literal content being uploaded to the bucket.
	Content *string `pulumi:"content"`
	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption *bool `pulumi:"encryption"`
	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is `md5(file("pathToFile"))`.
	Etag *string `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// The ID of the kms key. If omitted, the default master key will be used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region *string `pulumi:"region"`
	// the size of the object in bytes.
	Size *int `pulumi:"size"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
	// Specifies the storage class of the object. Defaults to `STANDARD`.
	StorageClass *string `pulumi:"storageClass"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId *string `pulumi:"versionId"`
}

type BucketObjectState struct {
	// The ACL policy to apply. Defaults to `private`.
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringPtrInput
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption pulumi.BoolPtrInput
	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is `md5(file("pathToFile"))`.
	Etag pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// The ID of the kms key. If omitted, the default master key will be used.
	KmsKeyId pulumi.StringPtrInput
	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region pulumi.StringPtrInput
	// the size of the object in bytes.
	Size pulumi.IntPtrInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
	// Specifies the storage class of the object. Defaults to `STANDARD`.
	StorageClass pulumi.StringPtrInput
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringPtrInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The ACL policy to apply. Defaults to `private`.
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket string `pulumi:"bucket"`
	// The literal content being uploaded to the bucket.
	Content *string `pulumi:"content"`
	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption *bool `pulumi:"encryption"`
	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is `md5(file("pathToFile"))`.
	Etag *string `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key string `pulumi:"key"`
	// The ID of the kms key. If omitted, the default master key will be used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region *string `pulumi:"region"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
	// Specifies the storage class of the object. Defaults to `STANDARD`.
	StorageClass *string `pulumi:"storageClass"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The ACL policy to apply. Defaults to `private`.
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringInput
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption pulumi.BoolPtrInput
	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is `md5(file("pathToFile"))`.
	Etag pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringInput
	// The ID of the kms key. If omitted, the default master key will be used.
	KmsKeyId pulumi.StringPtrInput
	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region pulumi.StringPtrInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
	// Specifies the storage class of the object. Defaults to `STANDARD`.
	StorageClass pulumi.StringPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (*BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (i *BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i *BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

// BucketObjectArrayInput is an input type that accepts BucketObjectArray and BucketObjectArrayOutput values.
// You can construct a concrete instance of `BucketObjectArrayInput` via:
//
//	BucketObjectArray{ BucketObjectArgs{...} }
type BucketObjectArrayInput interface {
	pulumi.Input

	ToBucketObjectArrayOutput() BucketObjectArrayOutput
	ToBucketObjectArrayOutputWithContext(context.Context) BucketObjectArrayOutput
}

type BucketObjectArray []BucketObjectInput

func (BucketObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (i BucketObjectArray) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return i.ToBucketObjectArrayOutputWithContext(context.Background())
}

func (i BucketObjectArray) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectArrayOutput)
}

// BucketObjectMapInput is an input type that accepts BucketObjectMap and BucketObjectMapOutput values.
// You can construct a concrete instance of `BucketObjectMapInput` via:
//
//	BucketObjectMap{ "key": BucketObjectArgs{...} }
type BucketObjectMapInput interface {
	pulumi.Input

	ToBucketObjectMapOutput() BucketObjectMapOutput
	ToBucketObjectMapOutputWithContext(context.Context) BucketObjectMapOutput
}

type BucketObjectMap map[string]BucketObjectInput

func (BucketObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (i BucketObjectMap) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return i.ToBucketObjectMapOutputWithContext(context.Background())
}

func (i BucketObjectMap) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectMapOutput)
}

type BucketObjectOutput struct{ *pulumi.OutputState }

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

// The ACL policy to apply. Defaults to `private`.
func (o BucketObjectOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// The name of the bucket to put the file in.
func (o BucketObjectOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The literal content being uploaded to the bucket.
func (o BucketObjectOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// A standard MIME type describing the format of the object data, e.g.
// application/octet-stream. All Valid MIME Types are valid for this input.
func (o BucketObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// Whether enable server-side encryption of the object in SSE-KMS mode.
func (o BucketObjectOutput) Encryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.BoolPtrOutput { return v.Encryption }).(pulumi.BoolPtrOutput)
}

// Specifies the unique identifier of the object content. It can be used to trigger updates.
// The only meaningful value is `md5(file("pathToFile"))`.
func (o BucketObjectOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the object once it is in the bucket.
func (o BucketObjectOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The ID of the kms key. If omitted, the default master key will be used.
func (o BucketObjectOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The region in which to create the OBS bucket object resource. If omitted, the
// provider-level region will be used. Changing this creates a new OBS bucket object resource.
func (o BucketObjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// the size of the object in bytes.
func (o BucketObjectOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The path to the source file being uploaded to the bucket.
func (o BucketObjectOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// Specifies the storage class of the object. Defaults to `STANDARD`.
func (o BucketObjectOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.StorageClass }).(pulumi.StringOutput)
}

// A unique version ID value for the object, if bucket versioning is enabled.
func (o BucketObjectOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

type BucketObjectArrayOutput struct{ *pulumi.OutputState }

func (BucketObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) Index(i pulumi.IntInput) BucketObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].([]*BucketObject)[vs[1].(int)]
	}).(BucketObjectOutput)
}

type BucketObjectMapOutput struct{ *pulumi.OutputState }

func (BucketObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) MapIndex(k pulumi.StringInput) BucketObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].(map[string]*BucketObject)[vs[1].(string)]
	}).(BucketObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectInput)(nil)).Elem(), &BucketObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectArrayInput)(nil)).Elem(), BucketObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectMapInput)(nil)).Elem(), BucketObjectMap{})
	pulumi.RegisterOutputType(BucketObjectOutput{})
	pulumi.RegisterOutputType(BucketObjectArrayOutput{})
	pulumi.RegisterOutputType(BucketObjectMapOutput{})
}