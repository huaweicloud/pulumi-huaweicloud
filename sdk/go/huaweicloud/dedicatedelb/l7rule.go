// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB L7 Rule resource within HuaweiCloud.
//
// ## Example Usage
// ### Create by value
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			l7policyId := cfg.RequireObject("l7policyId")
//			_, err := DedicatedElb.NewL7rule(ctx, "l7rule1", &DedicatedElb.L7ruleArgs{
//				L7policyId:  pulumi.Any(l7policyId),
//				Type:        pulumi.String("PATH"),
//				CompareType: pulumi.String("EQUAL_TO"),
//				Value:       pulumi.String("/api"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create by conditions and type is HOST_NAME
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			l7policyId := cfg.RequireObject("l7policyId")
//			_, err := DedicatedElb.NewL7rule(ctx, "l7rule1", &DedicatedElb.L7ruleArgs{
//				L7policyId:  pulumi.Any(l7policyId),
//				Type:        pulumi.String("HOST_NAME"),
//				CompareType: pulumi.String("EQUAL_TO"),
//				Conditions: dedicatedelb.L7ruleConditionArray{
//					&dedicatedelb.L7ruleConditionArgs{
//						Value: pulumi.String("test.com"),
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
// ### Create by conditions and type is HEADER
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			l7policyId := cfg.RequireObject("l7policyId")
//			_, err := DedicatedElb.NewL7rule(ctx, "l7rule1", &DedicatedElb.L7ruleArgs{
//				L7policyId:  pulumi.Any(l7policyId),
//				Type:        pulumi.String("HEADER"),
//				CompareType: pulumi.String("EQUAL_TO"),
//				Conditions: dedicatedelb.L7ruleConditionArray{
//					&dedicatedelb.L7ruleConditionArgs{
//						Key:   pulumi.String("testKey"),
//						Value: pulumi.String("testValue"),
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
// ### Create by conditions and type is SOURCE_IP
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			l7policyId := cfg.RequireObject("l7policyId")
//			_, err := DedicatedElb.NewL7rule(ctx, "l7rule1", &DedicatedElb.L7ruleArgs{
//				L7policyId:  pulumi.Any(l7policyId),
//				Type:        pulumi.String("SOURCE_IP"),
//				CompareType: pulumi.String("EQUAL_TO"),
//				Conditions: dedicatedelb.L7ruleConditionArray{
//					&dedicatedelb.L7ruleConditionArgs{
//						Value: pulumi.String("192.168.0.2/32"),
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
// ELB L7 rule can be imported using the `l7policy_id` and `id` separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:DedicatedElb/l7rule:L7rule rule_1 <l7policy_id>/<id>
//
// ```
type L7rule struct {
	pulumi.CustomResourceState

	// Specifies how requests are matched with the forwarding rule. Value options:
	// + **EQUAL_TO**: Exact match.
	// + **REGEX**: Regular expression match.
	// + **STARTS_WITH**: Prefix match.
	CompareType pulumi.StringOutput `pulumi:"compareType"`
	// Specifies the matching conditions of the forwarding rule. This parameter is available
	// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
	// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
	// must be the same, whereas each value must be unique.
	// The condition structure is documented below.
	Conditions L7ruleConditionArrayOutput `pulumi:"conditions"`
	// The create time of the L7 Rule.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
	L7policyId pulumi.StringOutput `pulumi:"l7policyId"`
	// The region in which to create the L7 Rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new L7 Rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the L7 Rule type. Value options:
	// + **HOST_NAME**: A domain name will be used for matching.
	// + **PATH**: A URL will be used for matching.
	// + **METHOD**: An HTTP request method will be used for matching.
	// + **HEADER**: The request header will be used for matching.
	// + **QUERY_STRING**: A query string will be used for matching.
	// + **SOURCE_IP**: The source IP address will be used for matching.
	// + **COOKIE**: The cookie will be used for matching.
	Type pulumi.StringOutput `pulumi:"type"`
	// The update time of the L7 Rule.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Specifies the value of the match item.
	// + If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
	//   letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
	//   If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	// + If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
	//   `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
	//   slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
	// + If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
	//   characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
	//   An asterisk can match zero or more characters, and a question mark can match 1 character.
	// + If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
	//   and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
	//   backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
	//   and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
	//   more characters, and a question mark can match 1 character.
	// + If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
	//   **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
	// + If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
	//   **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewL7rule registers a new resource with the given unique name, arguments, and options.
func NewL7rule(ctx *pulumi.Context,
	name string, args *L7ruleArgs, opts ...pulumi.ResourceOption) (*L7rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompareType == nil {
		return nil, errors.New("invalid value for required argument 'CompareType'")
	}
	if args.L7policyId == nil {
		return nil, errors.New("invalid value for required argument 'L7policyId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource L7rule
	err := ctx.RegisterResource("huaweicloud:DedicatedElb/l7rule:L7rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL7rule gets an existing L7rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL7rule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L7ruleState, opts ...pulumi.ResourceOption) (*L7rule, error) {
	var resource L7rule
	err := ctx.ReadResource("huaweicloud:DedicatedElb/l7rule:L7rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L7rule resources.
type l7ruleState struct {
	// Specifies how requests are matched with the forwarding rule. Value options:
	// + **EQUAL_TO**: Exact match.
	// + **REGEX**: Regular expression match.
	// + **STARTS_WITH**: Prefix match.
	CompareType *string `pulumi:"compareType"`
	// Specifies the matching conditions of the forwarding rule. This parameter is available
	// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
	// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
	// must be the same, whereas each value must be unique.
	// The condition structure is documented below.
	Conditions []L7ruleCondition `pulumi:"conditions"`
	// The create time of the L7 Rule.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
	L7policyId *string `pulumi:"l7policyId"`
	// The region in which to create the L7 Rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new L7 Rule.
	Region *string `pulumi:"region"`
	// Specifies the L7 Rule type. Value options:
	// + **HOST_NAME**: A domain name will be used for matching.
	// + **PATH**: A URL will be used for matching.
	// + **METHOD**: An HTTP request method will be used for matching.
	// + **HEADER**: The request header will be used for matching.
	// + **QUERY_STRING**: A query string will be used for matching.
	// + **SOURCE_IP**: The source IP address will be used for matching.
	// + **COOKIE**: The cookie will be used for matching.
	Type *string `pulumi:"type"`
	// The update time of the L7 Rule.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Specifies the value of the match item.
	// + If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
	//   letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
	//   If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	// + If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
	//   `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
	//   slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
	// + If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
	//   characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
	//   An asterisk can match zero or more characters, and a question mark can match 1 character.
	// + If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
	//   and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
	//   backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
	//   and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
	//   more characters, and a question mark can match 1 character.
	// + If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
	//   **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
	// + If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
	//   **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	Value *string `pulumi:"value"`
}

type L7ruleState struct {
	// Specifies how requests are matched with the forwarding rule. Value options:
	// + **EQUAL_TO**: Exact match.
	// + **REGEX**: Regular expression match.
	// + **STARTS_WITH**: Prefix match.
	CompareType pulumi.StringPtrInput
	// Specifies the matching conditions of the forwarding rule. This parameter is available
	// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
	// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
	// must be the same, whereas each value must be unique.
	// The condition structure is documented below.
	Conditions L7ruleConditionArrayInput
	// The create time of the L7 Rule.
	CreatedAt pulumi.StringPtrInput
	// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
	L7policyId pulumi.StringPtrInput
	// The region in which to create the L7 Rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new L7 Rule.
	Region pulumi.StringPtrInput
	// Specifies the L7 Rule type. Value options:
	// + **HOST_NAME**: A domain name will be used for matching.
	// + **PATH**: A URL will be used for matching.
	// + **METHOD**: An HTTP request method will be used for matching.
	// + **HEADER**: The request header will be used for matching.
	// + **QUERY_STRING**: A query string will be used for matching.
	// + **SOURCE_IP**: The source IP address will be used for matching.
	// + **COOKIE**: The cookie will be used for matching.
	Type pulumi.StringPtrInput
	// The update time of the L7 Rule.
	UpdatedAt pulumi.StringPtrInput
	// Specifies the value of the match item.
	// + If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
	//   letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
	//   If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	// + If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
	//   `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
	//   slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
	// + If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
	//   characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
	//   An asterisk can match zero or more characters, and a question mark can match 1 character.
	// + If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
	//   and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
	//   backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
	//   and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
	//   more characters, and a question mark can match 1 character.
	// + If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
	//   **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
	// + If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
	//   **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	Value pulumi.StringPtrInput
}

func (L7ruleState) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleState)(nil)).Elem()
}

type l7ruleArgs struct {
	// Specifies how requests are matched with the forwarding rule. Value options:
	// + **EQUAL_TO**: Exact match.
	// + **REGEX**: Regular expression match.
	// + **STARTS_WITH**: Prefix match.
	CompareType string `pulumi:"compareType"`
	// Specifies the matching conditions of the forwarding rule. This parameter is available
	// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
	// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
	// must be the same, whereas each value must be unique.
	// The condition structure is documented below.
	Conditions []L7ruleCondition `pulumi:"conditions"`
	// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
	L7policyId string `pulumi:"l7policyId"`
	// The region in which to create the L7 Rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new L7 Rule.
	Region *string `pulumi:"region"`
	// Specifies the L7 Rule type. Value options:
	// + **HOST_NAME**: A domain name will be used for matching.
	// + **PATH**: A URL will be used for matching.
	// + **METHOD**: An HTTP request method will be used for matching.
	// + **HEADER**: The request header will be used for matching.
	// + **QUERY_STRING**: A query string will be used for matching.
	// + **SOURCE_IP**: The source IP address will be used for matching.
	// + **COOKIE**: The cookie will be used for matching.
	Type string `pulumi:"type"`
	// Specifies the value of the match item.
	// + If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
	//   letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
	//   If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	// + If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
	//   `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
	//   slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
	// + If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
	//   characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
	//   An asterisk can match zero or more characters, and a question mark can match 1 character.
	// + If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
	//   and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
	//   backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
	//   and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
	//   more characters, and a question mark can match 1 character.
	// + If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
	//   **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
	// + If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
	//   **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a L7rule resource.
type L7ruleArgs struct {
	// Specifies how requests are matched with the forwarding rule. Value options:
	// + **EQUAL_TO**: Exact match.
	// + **REGEX**: Regular expression match.
	// + **STARTS_WITH**: Prefix match.
	CompareType pulumi.StringInput
	// Specifies the matching conditions of the forwarding rule. This parameter is available
	// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
	// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
	// must be the same, whereas each value must be unique.
	// The condition structure is documented below.
	Conditions L7ruleConditionArrayInput
	// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
	L7policyId pulumi.StringInput
	// The region in which to create the L7 Rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new L7 Rule.
	Region pulumi.StringPtrInput
	// Specifies the L7 Rule type. Value options:
	// + **HOST_NAME**: A domain name will be used for matching.
	// + **PATH**: A URL will be used for matching.
	// + **METHOD**: An HTTP request method will be used for matching.
	// + **HEADER**: The request header will be used for matching.
	// + **QUERY_STRING**: A query string will be used for matching.
	// + **SOURCE_IP**: The source IP address will be used for matching.
	// + **COOKIE**: The cookie will be used for matching.
	Type pulumi.StringInput
	// Specifies the value of the match item.
	// + If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
	//   letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
	//   If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
	// + If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
	//   `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
	//   slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
	// + If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
	//   characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
	//   An asterisk can match zero or more characters, and a question mark can match 1 character.
	// + If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
	//   and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
	//   backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
	//   and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
	//   more characters, and a question mark can match 1 character.
	// + If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
	//   **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
	// + If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
	//   **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
	Value pulumi.StringPtrInput
}

func (L7ruleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleArgs)(nil)).Elem()
}

type L7ruleInput interface {
	pulumi.Input

	ToL7ruleOutput() L7ruleOutput
	ToL7ruleOutputWithContext(ctx context.Context) L7ruleOutput
}

func (*L7rule) ElementType() reflect.Type {
	return reflect.TypeOf((**L7rule)(nil)).Elem()
}

func (i *L7rule) ToL7ruleOutput() L7ruleOutput {
	return i.ToL7ruleOutputWithContext(context.Background())
}

func (i *L7rule) ToL7ruleOutputWithContext(ctx context.Context) L7ruleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7ruleOutput)
}

// L7ruleArrayInput is an input type that accepts L7ruleArray and L7ruleArrayOutput values.
// You can construct a concrete instance of `L7ruleArrayInput` via:
//
//	L7ruleArray{ L7ruleArgs{...} }
type L7ruleArrayInput interface {
	pulumi.Input

	ToL7ruleArrayOutput() L7ruleArrayOutput
	ToL7ruleArrayOutputWithContext(context.Context) L7ruleArrayOutput
}

type L7ruleArray []L7ruleInput

func (L7ruleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7rule)(nil)).Elem()
}

func (i L7ruleArray) ToL7ruleArrayOutput() L7ruleArrayOutput {
	return i.ToL7ruleArrayOutputWithContext(context.Background())
}

func (i L7ruleArray) ToL7ruleArrayOutputWithContext(ctx context.Context) L7ruleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7ruleArrayOutput)
}

// L7ruleMapInput is an input type that accepts L7ruleMap and L7ruleMapOutput values.
// You can construct a concrete instance of `L7ruleMapInput` via:
//
//	L7ruleMap{ "key": L7ruleArgs{...} }
type L7ruleMapInput interface {
	pulumi.Input

	ToL7ruleMapOutput() L7ruleMapOutput
	ToL7ruleMapOutputWithContext(context.Context) L7ruleMapOutput
}

type L7ruleMap map[string]L7ruleInput

func (L7ruleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7rule)(nil)).Elem()
}

func (i L7ruleMap) ToL7ruleMapOutput() L7ruleMapOutput {
	return i.ToL7ruleMapOutputWithContext(context.Background())
}

func (i L7ruleMap) ToL7ruleMapOutputWithContext(ctx context.Context) L7ruleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7ruleMapOutput)
}

type L7ruleOutput struct{ *pulumi.OutputState }

func (L7ruleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L7rule)(nil)).Elem()
}

func (o L7ruleOutput) ToL7ruleOutput() L7ruleOutput {
	return o
}

func (o L7ruleOutput) ToL7ruleOutputWithContext(ctx context.Context) L7ruleOutput {
	return o
}

// Specifies how requests are matched with the forwarding rule. Value options:
// + **EQUAL_TO**: Exact match.
// + **REGEX**: Regular expression match.
// + **STARTS_WITH**: Prefix match.
func (o L7ruleOutput) CompareType() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.CompareType }).(pulumi.StringOutput)
}

// Specifies the matching conditions of the forwarding rule. This parameter is available
// only when `enhanceL7policyEnable` of the listener is set to **true**. If it is specified, parameter `value` will
// not take effect, and the value will contain all conditions configured for the forwarding rule. The keys in the list
// must be the same, whereas each value must be unique.
// The condition structure is documented below.
func (o L7ruleOutput) Conditions() L7ruleConditionArrayOutput {
	return o.ApplyT(func(v *L7rule) L7ruleConditionArrayOutput { return v.Conditions }).(L7ruleConditionArrayOutput)
}

// The create time of the L7 Rule.
func (o L7ruleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the ID of the L7 Policy. Changing this creates a new L7 Rule.
func (o L7ruleOutput) L7policyId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.L7policyId }).(pulumi.StringOutput)
}

// The region in which to create the L7 Rule resource. If omitted, the
// provider-level region will be used. Changing this creates a new L7 Rule.
func (o L7ruleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the L7 Rule type. Value options:
// + **HOST_NAME**: A domain name will be used for matching.
// + **PATH**: A URL will be used for matching.
// + **METHOD**: An HTTP request method will be used for matching.
// + **HEADER**: The request header will be used for matching.
// + **QUERY_STRING**: A query string will be used for matching.
// + **SOURCE_IP**: The source IP address will be used for matching.
// + **COOKIE**: The cookie will be used for matching.
func (o L7ruleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The update time of the L7 Rule.
func (o L7ruleOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Specifies the value of the match item.
//   - If `type` is set to **HOST_NAME**, it indicates the domain name, which can contain 1 to 128 characters, including
//     letters, digits, hyphens (-), periods (.), and asterisks (), and must start with a letter, digit, or asterisk ().
//     If you want to use a wildcard domain name, enter an asterisk (*) as the leftmost label of the domain name.
//   - If `type` is set to **PATH**, it indicates the request path, which can contain 1 to 128 characters. If
//     `compareType` is set to **STARTS_WITH** or **EQUAL_TO** for the forwarding rule, the value must start with a
//     slash (/) and can contain only letters, digits, and special characters _~';@^-%#&$.*+?,=!:|/()[]{}.
//   - If `type` is set to **HEADER**, it indicates the value of the HTTP header parameter. The value can contain 1 to 128
//     characters. Asterisks (*) and question marks (?)are allowed, but spaces and double quotation marks are not allowed.
//     An asterisk can match zero or more characters, and a question mark can match 1 character.
//   - If `type` is set to **QUERY_STRING**, it indicates the value of the query parameter. The value is case-sensitive
//     and can contain 1 to 128 characters. Spaces, square brackets ([]), curly brackets ({}), angle brackets (<>),
//     backslashes (), double quotation marks (""), pound signs (#), ampersands (&), vertical bars (|), percent signs (%),
//     and tildes (~) are not supported. Asterisks (*)and question marks (?) are allowed. An asterisk can match zero or
//     more characters, and a question mark can match 1 character.
//   - If `type` is set to **METHOD**, it indicates the HTTP method. The value can be **GET**, **PUT**, **POST**,
//     **DELETE**, **PATCH**, **HEAD**, or **OPTIONS**.
//   - If `type` is set to **SOURCE_IP**, it indicates the source IP address of the request. The value is an **IPv4** or
//     **IPv6** CIDR block, for example, 192.168.0.2/32 or 2049::49/64.
func (o L7ruleOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *L7rule) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type L7ruleArrayOutput struct{ *pulumi.OutputState }

func (L7ruleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7rule)(nil)).Elem()
}

func (o L7ruleArrayOutput) ToL7ruleArrayOutput() L7ruleArrayOutput {
	return o
}

func (o L7ruleArrayOutput) ToL7ruleArrayOutputWithContext(ctx context.Context) L7ruleArrayOutput {
	return o
}

func (o L7ruleArrayOutput) Index(i pulumi.IntInput) L7ruleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L7rule {
		return vs[0].([]*L7rule)[vs[1].(int)]
	}).(L7ruleOutput)
}

type L7ruleMapOutput struct{ *pulumi.OutputState }

func (L7ruleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7rule)(nil)).Elem()
}

func (o L7ruleMapOutput) ToL7ruleMapOutput() L7ruleMapOutput {
	return o
}

func (o L7ruleMapOutput) ToL7ruleMapOutputWithContext(ctx context.Context) L7ruleMapOutput {
	return o
}

func (o L7ruleMapOutput) MapIndex(k pulumi.StringInput) L7ruleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L7rule {
		return vs[0].(map[string]*L7rule)[vs[1].(string)]
	}).(L7ruleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L7ruleInput)(nil)).Elem(), &L7rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7ruleArrayInput)(nil)).Elem(), L7ruleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7ruleMapInput)(nil)).Elem(), L7ruleMap{})
	pulumi.RegisterOutputType(L7ruleOutput{})
	pulumi.RegisterOutputType(L7ruleArrayOutput{})
	pulumi.RegisterOutputType(L7ruleMapOutput{})
}
