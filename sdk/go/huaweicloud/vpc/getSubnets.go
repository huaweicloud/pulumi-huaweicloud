// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of VPC subnet.
//
// ## **Attributes Reference**
//
// The following attributes are exported:
//
// * `id` - Indicates a data source ID.
// * `subnets` - Indicates a list of all subnets found. Structure is documented below.
//
// The `subnets` block supports:
//
// * `id` - Indicates the ID of the subnet.
// * `name` - Indicates the name of the subnet.
// * `description` - Indicates the description of the subnet.
// * `cidr` - Indicates the cidr block of the subnet.
// * `status` - Indicates the current status of the subnet.
// * `vpcId` - Indicates the Id of the VPC that the subnet belongs to.
// * `gatewayIp` - Indicates the subnet gateway address of the subnet.
// * `primaryDns` - Indicates the IP address of DNS server 1 on the subnet.
// * `secondaryDns` - Indicates the IP address of DNS server 2 on the subnet.
// * `availabilityZone` - Indicates the availability zone (AZ) to which the subnet belongs to.
// * `dhcpEnable` - Indicates whether the DHCP is enabled.
// * `dnsList` - Indicates The IP address list of DNS servers on the subnet.
// * `ipv4SubnetId` - Indicates the ID of the IPv4 subnet (Native OpenStack API).
// * `ipv6Enable` - Indicates whether the IPv6 is enabled.
// * `ipv6SubnetId` - Indicates the ID of the IPv6 subnet (Native OpenStack API).
// * `ipv6Cidr` - Indicates the IPv6 subnet CIDR block.
// * `ipv6Gateway` - Indicates the IPv6 subnet gateway.
// * `tags` - Indicates the key/value pairs which associated with the subnet.
func GetSubnets(ctx *pulumi.Context, args *GetSubnetsArgs, opts ...pulumi.InvokeOption) (*GetSubnetsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSubnetsResult
	err := ctx.Invoke("huaweicloud:Vpc/getSubnets:getSubnets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsArgs struct {
	// Specifies the availability zone (AZ) to which the desired subnet belongs to.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the network segment of desired subnet. The value must be in CIDR format.
	Cidr *string `pulumi:"cidr"`
	// Specifies the subnet gateway address of desired subnet.
	GatewayIp *string `pulumi:"gatewayIp"`
	// - Specifies the id of the desired subnet.
	Id *string `pulumi:"id"`
	// Specifies the name of the desired subnet.
	Name *string `pulumi:"name"`
	// Specifies the IP address of DNS server 1 on the desired subnet.
	PrimaryDns *string `pulumi:"primaryDns"`
	// Specifies the region in which to obtain the subnet. If omitted, the provider-level
	// region will be used.
	Region *string `pulumi:"region"`
	// Specifies the IP address of DNS server 2 on the desired subnet.
	SecondaryDns *string `pulumi:"secondaryDns"`
	// Specifies the current status of the desired subnet.
	// the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status *string `pulumi:"status"`
	// Specifies the included key/value pairs which associated with the desired subnet.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the id of the VPC that the desired subnet belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSubnets.
type GetSubnetsResult struct {
	AvailabilityZone *string            `pulumi:"availabilityZone"`
	Cidr             *string            `pulumi:"cidr"`
	GatewayIp        *string            `pulumi:"gatewayIp"`
	Id               string             `pulumi:"id"`
	Name             *string            `pulumi:"name"`
	PrimaryDns       *string            `pulumi:"primaryDns"`
	Region           string             `pulumi:"region"`
	SecondaryDns     *string            `pulumi:"secondaryDns"`
	Status           *string            `pulumi:"status"`
	Subnets          []GetSubnetsSubnet `pulumi:"subnets"`
	Tags             map[string]string  `pulumi:"tags"`
	VpcId            *string            `pulumi:"vpcId"`
}

func GetSubnetsOutput(ctx *pulumi.Context, args GetSubnetsOutputArgs, opts ...pulumi.InvokeOption) GetSubnetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetsResult, error) {
			args := v.(GetSubnetsArgs)
			r, err := GetSubnets(ctx, &args, opts...)
			var s GetSubnetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetsResultOutput)
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsOutputArgs struct {
	// Specifies the availability zone (AZ) to which the desired subnet belongs to.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Specifies the network segment of desired subnet. The value must be in CIDR format.
	Cidr pulumi.StringPtrInput `pulumi:"cidr"`
	// Specifies the subnet gateway address of desired subnet.
	GatewayIp pulumi.StringPtrInput `pulumi:"gatewayIp"`
	// - Specifies the id of the desired subnet.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies the name of the desired subnet.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the IP address of DNS server 1 on the desired subnet.
	PrimaryDns pulumi.StringPtrInput `pulumi:"primaryDns"`
	// Specifies the region in which to obtain the subnet. If omitted, the provider-level
	// region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the IP address of DNS server 2 on the desired subnet.
	SecondaryDns pulumi.StringPtrInput `pulumi:"secondaryDns"`
	// Specifies the current status of the desired subnet.
	// the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Specifies the included key/value pairs which associated with the desired subnet.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Specifies the id of the VPC that the desired subnet belongs to.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetSubnetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsArgs)(nil)).Elem()
}

// A collection of values returned by getSubnets.
type GetSubnetsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsResult)(nil)).Elem()
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutput() GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutputWithContext(ctx context.Context) GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) GatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.GatewayIp }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSubnetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) PrimaryDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.PrimaryDns }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetSubnetsResultOutput) SecondaryDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.SecondaryDns }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Subnets() GetSubnetsSubnetArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []GetSubnetsSubnet { return v.Subnets }).(GetSubnetsSubnetArrayOutput)
}

func (o GetSubnetsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetSubnetsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetsResultOutput{})
}
