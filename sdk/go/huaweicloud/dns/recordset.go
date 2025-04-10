// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS record set resource within HuaweiCloud.
//
// ## Example Usage
// ### Record Set with Public Zone
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			publicZoneId := cfg.RequireObject("publicZoneId")
//			publicRecordsetName := cfg.RequireObject("publicRecordsetName")
//			description := cfg.RequireObject("description")
//			_, err := Dns.NewRecordset(ctx, "test", &Dns.RecordsetArgs{
//				ZoneId:      pulumi.Any(publicZoneId),
//				Type:        pulumi.String("A"),
//				Description: pulumi.Any(description),
//				Status:      pulumi.String("ENABLE"),
//				Ttl:         pulumi.Int(300),
//				Records: pulumi.StringArray{
//					pulumi.String("10.1.0.0"),
//				},
//				LineId: pulumi.String("Dianxin_Shanxi"),
//				Weight: pulumi.Int(3),
//				Tags: pulumi.StringMap{
//					"key1": pulumi.String("value1"),
//					"key2": pulumi.String("value2"),
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
// ### Record Set with Private Zone
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			privateZoneId := cfg.RequireObject("privateZoneId")
//			privateRecordsetName := cfg.RequireObject("privateRecordsetName")
//			description := cfg.RequireObject("description")
//			_, err := Dns.NewRecordset(ctx, "test", &Dns.RecordsetArgs{
//				ZoneId:      pulumi.Any(privateZoneId),
//				Description: pulumi.Any(description),
//				Ttl:         pulumi.Int(3000),
//				Type:        pulumi.String("A"),
//				Records: pulumi.StringArray{
//					pulumi.String("10.0.0.1"),
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
// The DNS recordset can be imported using `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Dns/recordset:Recordset test <id>
//
// ```
type Recordset struct {
	pulumi.CustomResourceState

	// Specifies the description of the record set.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the resolution line ID.\
	// Changing this parameter will create a new resource.
	LineId pulumi.StringOutput `pulumi:"lineId"`
	// Specifies the name of the record set.\
	// The name suffixed with a zone name, which is a complete host name ended with a dot.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the list of the records of the record set.\
	// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the status of the record set, defaults to **ENABLE**.\
	// The valid values are as follows:
	// + **ENABLE**
	// + **DISABLE**
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Specifies the key/value pairs to associate with the DNS record set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the time to live (TTL) of the record set (in seconds).\
	// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// Specifies the type of the record set.
	// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
	// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the weight of the record set.
	// Only public zone support. The valid value is range from `1` to `1,000`.
	Weight pulumi.IntOutput `pulumi:"weight"`
	// Specifies the ID of the zone to which the record set belongs.\
	// Changing this parameter will create a new resource.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
	// The name of the zone to which the record set belongs.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
	// The type of the zone to which the record set belongs.
	// + **public**
	// + **private**
	ZoneType pulumi.StringOutput `pulumi:"zoneType"`
}

// NewRecordset registers a new resource with the given unique name, arguments, and options.
func NewRecordset(ctx *pulumi.Context,
	name string, args *RecordsetArgs, opts ...pulumi.ResourceOption) (*Recordset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Records == nil {
		return nil, errors.New("invalid value for required argument 'Records'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Recordset
	err := ctx.RegisterResource("huaweicloud:Dns/recordset:Recordset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordset gets an existing Recordset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordsetState, opts ...pulumi.ResourceOption) (*Recordset, error) {
	var resource Recordset
	err := ctx.ReadResource("huaweicloud:Dns/recordset:Recordset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Recordset resources.
type recordsetState struct {
	// Specifies the description of the record set.
	Description *string `pulumi:"description"`
	// Specifies the resolution line ID.\
	// Changing this parameter will create a new resource.
	LineId *string `pulumi:"lineId"`
	// Specifies the name of the record set.\
	// The name suffixed with a zone name, which is a complete host name ended with a dot.
	Name *string `pulumi:"name"`
	// Specifies the list of the records of the record set.\
	// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
	Records []string `pulumi:"records"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the status of the record set, defaults to **ENABLE**.\
	// The valid values are as follows:
	// + **ENABLE**
	// + **DISABLE**
	Status *string `pulumi:"status"`
	// Specifies the key/value pairs to associate with the DNS record set.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the time to live (TTL) of the record set (in seconds).\
	// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
	Ttl *int `pulumi:"ttl"`
	// Specifies the type of the record set.
	// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
	// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
	Type *string `pulumi:"type"`
	// Specifies the weight of the record set.
	// Only public zone support. The valid value is range from `1` to `1,000`.
	Weight *int `pulumi:"weight"`
	// Specifies the ID of the zone to which the record set belongs.\
	// Changing this parameter will create a new resource.
	ZoneId *string `pulumi:"zoneId"`
	// The name of the zone to which the record set belongs.
	ZoneName *string `pulumi:"zoneName"`
	// The type of the zone to which the record set belongs.
	// + **public**
	// + **private**
	ZoneType *string `pulumi:"zoneType"`
}

type RecordsetState struct {
	// Specifies the description of the record set.
	Description pulumi.StringPtrInput
	// Specifies the resolution line ID.\
	// Changing this parameter will create a new resource.
	LineId pulumi.StringPtrInput
	// Specifies the name of the record set.\
	// The name suffixed with a zone name, which is a complete host name ended with a dot.
	Name pulumi.StringPtrInput
	// Specifies the list of the records of the record set.\
	// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
	Records pulumi.StringArrayInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the status of the record set, defaults to **ENABLE**.\
	// The valid values are as follows:
	// + **ENABLE**
	// + **DISABLE**
	Status pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the DNS record set.
	Tags pulumi.StringMapInput
	// Specifies the time to live (TTL) of the record set (in seconds).\
	// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
	Ttl pulumi.IntPtrInput
	// Specifies the type of the record set.
	// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
	// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
	Type pulumi.StringPtrInput
	// Specifies the weight of the record set.
	// Only public zone support. The valid value is range from `1` to `1,000`.
	Weight pulumi.IntPtrInput
	// Specifies the ID of the zone to which the record set belongs.\
	// Changing this parameter will create a new resource.
	ZoneId pulumi.StringPtrInput
	// The name of the zone to which the record set belongs.
	ZoneName pulumi.StringPtrInput
	// The type of the zone to which the record set belongs.
	// + **public**
	// + **private**
	ZoneType pulumi.StringPtrInput
}

func (RecordsetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordsetState)(nil)).Elem()
}

type recordsetArgs struct {
	// Specifies the description of the record set.
	Description *string `pulumi:"description"`
	// Specifies the resolution line ID.\
	// Changing this parameter will create a new resource.
	LineId *string `pulumi:"lineId"`
	// Specifies the name of the record set.\
	// The name suffixed with a zone name, which is a complete host name ended with a dot.
	Name *string `pulumi:"name"`
	// Specifies the list of the records of the record set.\
	// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
	Records []string `pulumi:"records"`
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the status of the record set, defaults to **ENABLE**.\
	// The valid values are as follows:
	// + **ENABLE**
	// + **DISABLE**
	Status *string `pulumi:"status"`
	// Specifies the key/value pairs to associate with the DNS record set.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the time to live (TTL) of the record set (in seconds).\
	// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
	Ttl *int `pulumi:"ttl"`
	// Specifies the type of the record set.
	// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
	// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
	Type string `pulumi:"type"`
	// Specifies the weight of the record set.
	// Only public zone support. The valid value is range from `1` to `1,000`.
	Weight *int `pulumi:"weight"`
	// Specifies the ID of the zone to which the record set belongs.\
	// Changing this parameter will create a new resource.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Recordset resource.
type RecordsetArgs struct {
	// Specifies the description of the record set.
	Description pulumi.StringPtrInput
	// Specifies the resolution line ID.\
	// Changing this parameter will create a new resource.
	LineId pulumi.StringPtrInput
	// Specifies the name of the record set.\
	// The name suffixed with a zone name, which is a complete host name ended with a dot.
	Name pulumi.StringPtrInput
	// Specifies the list of the records of the record set.\
	// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
	Records pulumi.StringArrayInput
	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the status of the record set, defaults to **ENABLE**.\
	// The valid values are as follows:
	// + **ENABLE**
	// + **DISABLE**
	Status pulumi.StringPtrInput
	// Specifies the key/value pairs to associate with the DNS record set.
	Tags pulumi.StringMapInput
	// Specifies the time to live (TTL) of the record set (in seconds).\
	// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
	Ttl pulumi.IntPtrInput
	// Specifies the type of the record set.
	// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
	// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
	Type pulumi.StringInput
	// Specifies the weight of the record set.
	// Only public zone support. The valid value is range from `1` to `1,000`.
	Weight pulumi.IntPtrInput
	// Specifies the ID of the zone to which the record set belongs.\
	// Changing this parameter will create a new resource.
	ZoneId pulumi.StringInput
}

func (RecordsetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordsetArgs)(nil)).Elem()
}

type RecordsetInput interface {
	pulumi.Input

	ToRecordsetOutput() RecordsetOutput
	ToRecordsetOutputWithContext(ctx context.Context) RecordsetOutput
}

func (*Recordset) ElementType() reflect.Type {
	return reflect.TypeOf((**Recordset)(nil)).Elem()
}

func (i *Recordset) ToRecordsetOutput() RecordsetOutput {
	return i.ToRecordsetOutputWithContext(context.Background())
}

func (i *Recordset) ToRecordsetOutputWithContext(ctx context.Context) RecordsetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordsetOutput)
}

// RecordsetArrayInput is an input type that accepts RecordsetArray and RecordsetArrayOutput values.
// You can construct a concrete instance of `RecordsetArrayInput` via:
//
//	RecordsetArray{ RecordsetArgs{...} }
type RecordsetArrayInput interface {
	pulumi.Input

	ToRecordsetArrayOutput() RecordsetArrayOutput
	ToRecordsetArrayOutputWithContext(context.Context) RecordsetArrayOutput
}

type RecordsetArray []RecordsetInput

func (RecordsetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Recordset)(nil)).Elem()
}

func (i RecordsetArray) ToRecordsetArrayOutput() RecordsetArrayOutput {
	return i.ToRecordsetArrayOutputWithContext(context.Background())
}

func (i RecordsetArray) ToRecordsetArrayOutputWithContext(ctx context.Context) RecordsetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordsetArrayOutput)
}

// RecordsetMapInput is an input type that accepts RecordsetMap and RecordsetMapOutput values.
// You can construct a concrete instance of `RecordsetMapInput` via:
//
//	RecordsetMap{ "key": RecordsetArgs{...} }
type RecordsetMapInput interface {
	pulumi.Input

	ToRecordsetMapOutput() RecordsetMapOutput
	ToRecordsetMapOutputWithContext(context.Context) RecordsetMapOutput
}

type RecordsetMap map[string]RecordsetInput

func (RecordsetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Recordset)(nil)).Elem()
}

func (i RecordsetMap) ToRecordsetMapOutput() RecordsetMapOutput {
	return i.ToRecordsetMapOutputWithContext(context.Background())
}

func (i RecordsetMap) ToRecordsetMapOutputWithContext(ctx context.Context) RecordsetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordsetMapOutput)
}

type RecordsetOutput struct{ *pulumi.OutputState }

func (RecordsetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Recordset)(nil)).Elem()
}

func (o RecordsetOutput) ToRecordsetOutput() RecordsetOutput {
	return o
}

func (o RecordsetOutput) ToRecordsetOutputWithContext(ctx context.Context) RecordsetOutput {
	return o
}

// Specifies the description of the record set.
func (o RecordsetOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the resolution line ID.\
// Changing this parameter will create a new resource.
func (o RecordsetOutput) LineId() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.LineId }).(pulumi.StringOutput)
}

// Specifies the name of the record set.\
// The name suffixed with a zone name, which is a complete host name ended with a dot.
func (o RecordsetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the list of the records of the record set.\
// The value depends on the `type` parameter, you can refer to this [document](https://support.huaweicloud.com/intl/en-us/usermanual-dns/dns_usermanual_0601.html#dns_usermanual_0601__table936244914119).
func (o RecordsetOutput) Records() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringArrayOutput { return v.Records }).(pulumi.StringArrayOutput)
}

// Specifies the region in which to create the resource.
// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
func (o RecordsetOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the status of the record set, defaults to **ENABLE**.\
// The valid values are as follows:
// + **ENABLE**
// + **DISABLE**
func (o RecordsetOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Specifies the key/value pairs to associate with the DNS record set.
func (o RecordsetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time to live (TTL) of the record set (in seconds).\
// The valid value is range from `1` to `2,147,483,647`. The default value is `300`.
func (o RecordsetOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Recordset) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// Specifies the type of the record set.
// + For the public record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV** and **CAA**.
// + For the private record set, the valid values are **A**, **AAAA**, **MX**, **CNAME**, **TXT** and **SRV**.
func (o RecordsetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the weight of the record set.
// Only public zone support. The valid value is range from `1` to `1,000`.
func (o RecordsetOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *Recordset) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

// Specifies the ID of the zone to which the record set belongs.\
// Changing this parameter will create a new resource.
func (o RecordsetOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

// The name of the zone to which the record set belongs.
func (o RecordsetOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.ZoneName }).(pulumi.StringOutput)
}

// The type of the zone to which the record set belongs.
// + **public**
// + **private**
func (o RecordsetOutput) ZoneType() pulumi.StringOutput {
	return o.ApplyT(func(v *Recordset) pulumi.StringOutput { return v.ZoneType }).(pulumi.StringOutput)
}

type RecordsetArrayOutput struct{ *pulumi.OutputState }

func (RecordsetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Recordset)(nil)).Elem()
}

func (o RecordsetArrayOutput) ToRecordsetArrayOutput() RecordsetArrayOutput {
	return o
}

func (o RecordsetArrayOutput) ToRecordsetArrayOutputWithContext(ctx context.Context) RecordsetArrayOutput {
	return o
}

func (o RecordsetArrayOutput) Index(i pulumi.IntInput) RecordsetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Recordset {
		return vs[0].([]*Recordset)[vs[1].(int)]
	}).(RecordsetOutput)
}

type RecordsetMapOutput struct{ *pulumi.OutputState }

func (RecordsetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Recordset)(nil)).Elem()
}

func (o RecordsetMapOutput) ToRecordsetMapOutput() RecordsetMapOutput {
	return o
}

func (o RecordsetMapOutput) ToRecordsetMapOutputWithContext(ctx context.Context) RecordsetMapOutput {
	return o
}

func (o RecordsetMapOutput) MapIndex(k pulumi.StringInput) RecordsetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Recordset {
		return vs[0].(map[string]*Recordset)[vs[1].(string)]
	}).(RecordsetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordsetInput)(nil)).Elem(), &Recordset{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordsetArrayInput)(nil)).Elem(), RecordsetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordsetMapInput)(nil)).Elem(), RecordsetMap{})
	pulumi.RegisterOutputType(RecordsetOutput{})
	pulumi.RegisterOutputType(RecordsetArrayOutput{})
	pulumi.RegisterOutputType(RecordsetMapOutput{})
}
