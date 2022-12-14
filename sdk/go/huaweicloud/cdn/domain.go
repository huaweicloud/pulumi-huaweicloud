// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CDN domain management.
//
// ## Example Usage
// ### Create a cdn domain
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			domainName := cfg.RequireObject("domainName")
//			originServer := cfg.RequireObject("originServer")
//			_, err := Cdn.NewDomain(ctx, "domain1", &Cdn.DomainArgs{
//				Type: pulumi.String("web"),
//				Sources: cdn.DomainSourceArray{
//					&cdn.DomainSourceArgs{
//						Origin:     pulumi.Any(originServer),
//						OriginType: pulumi.String("ipaddr"),
//						Active:     pulumi.Int(1),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("val"),
//					"foo": pulumi.String("bar"),
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
// ### Create a cdn domain with cache rules
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			domainName := cfg.RequireObject("domainName")
//			originServer := cfg.RequireObject("originServer")
//			_, err := Cdn.NewDomain(ctx, "domain1", &Cdn.DomainArgs{
//				Type: pulumi.String("web"),
//				Sources: cdn.DomainSourceArray{
//					&cdn.DomainSourceArgs{
//						Origin:     pulumi.Any(originServer),
//						OriginType: pulumi.String("ipaddr"),
//						Active:     pulumi.Int(1),
//					},
//				},
//				CacheSettings: &cdn.DomainCacheSettingsArgs{
//					Rules: cdn.DomainCacheSettingsRuleArray{
//						&cdn.DomainCacheSettingsRuleArgs{
//							RuleType: pulumi.Int(0),
//							Ttl:      pulumi.Int(180),
//							TtlType:  pulumi.Int(4),
//							Priority: pulumi.Int(2),
//						},
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
// ### Create a cdn domain with configs
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			domainName := cfg.RequireObject("domainName")
//			originServer := cfg.RequireObject("originServer")
//			_, err := Cdn.NewDomain(ctx, "domain1", &Cdn.DomainArgs{
//				Type: pulumi.String("web"),
//				Sources: cdn.DomainSourceArray{
//					&cdn.DomainSourceArgs{
//						Origin:     pulumi.Any(originServer),
//						OriginType: pulumi.String("ipaddr"),
//						Active:     pulumi.Int(1),
//					},
//				},
//				Configs: &cdn.DomainConfigsArgs{
//					OriginProtocol: pulumi.String("http"),
//					HttpsSettings: &cdn.DomainConfigsHttpsSettingsArgs{
//						CertificateName: pulumi.String("terraform-test"),
//						CertificateBody: readFileOrPanic("your_directory/chain.cer"),
//						Http2Enabled:    pulumi.Bool(true),
//						HttpsEnabled:    pulumi.Bool(true),
//						PrivateKey:      readFileOrPanic("your_directory/server_private.key"),
//					},
//					CacheUrlParameterFilter: &cdn.DomainConfigsCacheUrlParameterFilterArgs{
//						Type: pulumi.String("ignore_url_params"),
//					},
//					RetrievalRequestHeaders: cdn.DomainConfigsRetrievalRequestHeaderArray{
//						&cdn.DomainConfigsRetrievalRequestHeaderArgs{
//							Name:   pulumi.String("test-name"),
//							Value:  pulumi.String("test-val"),
//							Action: pulumi.String("set"),
//						},
//					},
//					HttpResponseHeaders: cdn.DomainConfigsHttpResponseHeaderArray{
//						&cdn.DomainConfigsHttpResponseHeaderArgs{
//							Name:   pulumi.String("test-name"),
//							Value:  pulumi.String("test-val"),
//							Action: pulumi.String("set"),
//						},
//					},
//					UrlSigning: &cdn.DomainConfigsUrlSigningArgs{
//						Enabled: pulumi.Bool(false),
//					},
//					Compress: &cdn.DomainConfigsCompressArgs{
//						Enabled: pulumi.Bool(false),
//					},
//					ForceRedirect: &cdn.DomainConfigsForceRedirectArgs{
//						Enabled: pulumi.Bool(true),
//						Type:    pulumi.String("http"),
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
// Domains can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Cdn/domain:Domain domain_1 fe2462fac09a4a42a76ecc4a1ef542f1
//
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Specifies the cache configuration. The object structure
	// is documented below.
	CacheSettings DomainCacheSettingsOutput `pulumi:"cacheSettings"`
	// The CNAME of the acceleration domain name.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Specifies the domain configuration items. The object structure is
	// documented below.
	Configs DomainConfigsOutput `pulumi:"configs"`
	// The status of the acceleration domain name. The available values are
	// 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'
	DomainStatus pulumi.StringOutput `pulumi:"domainStatus"`
	// The enterprise project id. Changing this parameter will create
	// a new resource.
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	// Specifies the request or response header.
	Name pulumi.StringOutput `pulumi:"name"`
	// The area covered by the acceleration service. Valid values are
	// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
	ServiceArea pulumi.StringOutput `pulumi:"serviceArea"`
	// An array of one or more objects specifies the domain name of the origin server.
	// The sources object structure is documented below.
	Sources DomainSourceArrayOutput `pulumi:"sources"`
	// Specifies the key/value pairs to associate with the domain.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the operation type for caching URL parameters. Posiible values are:
	// **full_url**: cache all parameters
	// **ignore_url_params**: ignore all parameters
	// **del_args**: ignore specific URL parameters
	// **reserve_args**: reserve specified URL parameters
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("huaweicloud:Cdn/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("huaweicloud:Cdn/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Specifies the cache configuration. The object structure
	// is documented below.
	CacheSettings *DomainCacheSettings `pulumi:"cacheSettings"`
	// The CNAME of the acceleration domain name.
	Cname *string `pulumi:"cname"`
	// Specifies the domain configuration items. The object structure is
	// documented below.
	Configs *DomainConfigs `pulumi:"configs"`
	// The status of the acceleration domain name. The available values are
	// 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'
	DomainStatus *string `pulumi:"domainStatus"`
	// The enterprise project id. Changing this parameter will create
	// a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the request or response header.
	Name *string `pulumi:"name"`
	// The area covered by the acceleration service. Valid values are
	// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
	ServiceArea *string `pulumi:"serviceArea"`
	// An array of one or more objects specifies the domain name of the origin server.
	// The sources object structure is documented below.
	Sources []DomainSource `pulumi:"sources"`
	// Specifies the key/value pairs to associate with the domain.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the operation type for caching URL parameters. Posiible values are:
	// **full_url**: cache all parameters
	// **ignore_url_params**: ignore all parameters
	// **del_args**: ignore specific URL parameters
	// **reserve_args**: reserve specified URL parameters
	Type *string `pulumi:"type"`
}

type DomainState struct {
	// Specifies the cache configuration. The object structure
	// is documented below.
	CacheSettings DomainCacheSettingsPtrInput
	// The CNAME of the acceleration domain name.
	Cname pulumi.StringPtrInput
	// Specifies the domain configuration items. The object structure is
	// documented below.
	Configs DomainConfigsPtrInput
	// The status of the acceleration domain name. The available values are
	// 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'
	DomainStatus pulumi.StringPtrInput
	// The enterprise project id. Changing this parameter will create
	// a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the request or response header.
	Name pulumi.StringPtrInput
	// The area covered by the acceleration service. Valid values are
	// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
	ServiceArea pulumi.StringPtrInput
	// An array of one or more objects specifies the domain name of the origin server.
	// The sources object structure is documented below.
	Sources DomainSourceArrayInput
	// Specifies the key/value pairs to associate with the domain.
	Tags pulumi.StringMapInput
	// Specifies the operation type for caching URL parameters. Posiible values are:
	// **full_url**: cache all parameters
	// **ignore_url_params**: ignore all parameters
	// **del_args**: ignore specific URL parameters
	// **reserve_args**: reserve specified URL parameters
	Type pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Specifies the cache configuration. The object structure
	// is documented below.
	CacheSettings *DomainCacheSettings `pulumi:"cacheSettings"`
	// Specifies the domain configuration items. The object structure is
	// documented below.
	Configs *DomainConfigs `pulumi:"configs"`
	// The enterprise project id. Changing this parameter will create
	// a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the request or response header.
	Name *string `pulumi:"name"`
	// The area covered by the acceleration service. Valid values are
	// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
	ServiceArea *string `pulumi:"serviceArea"`
	// An array of one or more objects specifies the domain name of the origin server.
	// The sources object structure is documented below.
	Sources []DomainSource `pulumi:"sources"`
	// Specifies the key/value pairs to associate with the domain.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the operation type for caching URL parameters. Posiible values are:
	// **full_url**: cache all parameters
	// **ignore_url_params**: ignore all parameters
	// **del_args**: ignore specific URL parameters
	// **reserve_args**: reserve specified URL parameters
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Specifies the cache configuration. The object structure
	// is documented below.
	CacheSettings DomainCacheSettingsPtrInput
	// Specifies the domain configuration items. The object structure is
	// documented below.
	Configs DomainConfigsPtrInput
	// The enterprise project id. Changing this parameter will create
	// a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the request or response header.
	Name pulumi.StringPtrInput
	// The area covered by the acceleration service. Valid values are
	// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
	ServiceArea pulumi.StringPtrInput
	// An array of one or more objects specifies the domain name of the origin server.
	// The sources object structure is documented below.
	Sources DomainSourceArrayInput
	// Specifies the key/value pairs to associate with the domain.
	Tags pulumi.StringMapInput
	// Specifies the operation type for caching URL parameters. Posiible values are:
	// **full_url**: cache all parameters
	// **ignore_url_params**: ignore all parameters
	// **del_args**: ignore specific URL parameters
	// **reserve_args**: reserve specified URL parameters
	Type pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//	DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//	DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Specifies the cache configuration. The object structure
// is documented below.
func (o DomainOutput) CacheSettings() DomainCacheSettingsOutput {
	return o.ApplyT(func(v *Domain) DomainCacheSettingsOutput { return v.CacheSettings }).(DomainCacheSettingsOutput)
}

// The CNAME of the acceleration domain name.
func (o DomainOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

// Specifies the domain configuration items. The object structure is
// documented below.
func (o DomainOutput) Configs() DomainConfigsOutput {
	return o.ApplyT(func(v *Domain) DomainConfigsOutput { return v.Configs }).(DomainConfigsOutput)
}

// The status of the acceleration domain name. The available values are
// 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'
func (o DomainOutput) DomainStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainStatus }).(pulumi.StringOutput)
}

// The enterprise project id. Changing this parameter will create
// a new resource.
func (o DomainOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// Specifies the request or response header.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The area covered by the acceleration service. Valid values are
// `mainlandChina`, `outsideMainlandChina`, and `global`. Changing this parameter will create a new resource.
func (o DomainOutput) ServiceArea() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ServiceArea }).(pulumi.StringOutput)
}

// An array of one or more objects specifies the domain name of the origin server.
// The sources object structure is documented below.
func (o DomainOutput) Sources() DomainSourceArrayOutput {
	return o.ApplyT(func(v *Domain) DomainSourceArrayOutput { return v.Sources }).(DomainSourceArrayOutput)
}

// Specifies the key/value pairs to associate with the domain.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the operation type for caching URL parameters. Posiible values are:
// **full_url**: cache all parameters
// **ignore_url_params**: ignore all parameters
// **del_args**: ignore specific URL parameters
// **reserve_args**: reserve specified URL parameters
func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
