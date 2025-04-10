// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdbforopengauss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb opengauss instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforOpenGauss"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDBforOpenGauss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDBforOpenGauss.GetOpengaussInstance(ctx, &gaussdbforopengauss.GetOpengaussInstanceArgs{
//				Name: pulumi.StringRef("gaussdb-instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOpengaussInstance(ctx *pulumi.Context, args *LookupOpengaussInstanceArgs, opts ...pulumi.InvokeOption) (*LookupOpengaussInstanceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOpengaussInstanceResult
	err := ctx.Invoke("huaweicloud:GaussDBforOpenGauss/getOpengaussInstance:getOpengaussInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOpengaussInstance.
type LookupOpengaussInstanceArgs struct {
	// Specifies the name of the instance.
	Name *string `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region *string `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getOpengaussInstance.
type LookupOpengaussInstanceResult struct {
	// Indicates the availability zone where the node resides.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates the advanced backup policy. Structure is documented below.
	BackupStrategies []GetOpengaussInstanceBackupStrategy `pulumi:"backupStrategies"`
	// Indicates the count of coordinator node.
	CoordinatorNum int `pulumi:"coordinatorNum"`
	// Indicates the database information. Structure is documented below.
	Datastores []GetOpengaussInstanceDatastore `pulumi:"datastores"`
	// Indicates the default username.
	DbUserName string `pulumi:"dbUserName"`
	// Indicates the enterprise project id.
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// Indicates the instance specifications.
	Flavor string `pulumi:"flavor"`
	// Indicates the instance ha information. Structure is documented below.
	Has []GetOpengaussInstanceHa `pulumi:"has"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates the maintenance window.
	MaintenanceWindow string `pulumi:"maintenanceWindow"`
	// Indicates the node name.
	Name string `pulumi:"name"`
	// Indicates the instance nodes information. Structure is documented below.
	Nodes []GetOpengaussInstanceNode `pulumi:"nodes"`
	// Indicates the database port.
	Port int `pulumi:"port"`
	// Indicates the list of private IP address of the nodes.
	PrivateIps []string `pulumi:"privateIps"`
	// Indicates the public IP address of the DB instance.
	PublicIps []string `pulumi:"publicIps"`
	Region    string   `pulumi:"region"`
	// Indicates the replica num.
	ReplicaNum int `pulumi:"replicaNum"`
	// Indicates the security group ID.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Indicates the sharding num.
	ShardingNum int `pulumi:"shardingNum"`
	// Indicates the node status.
	Status   string `pulumi:"status"`
	SubnetId string `pulumi:"subnetId"`
	// Indicates the switch strategy.
	SwitchStrategy string `pulumi:"switchStrategy"`
	// Indicates the default username.
	TimeZone string `pulumi:"timeZone"`
	// Indicates the volume type. Value options: **ULTRAHIGH**, **ESSD**.
	Type string `pulumi:"type"`
	// Indicates the volume information. Structure is documented below.
	Volumes []GetOpengaussInstanceVolume `pulumi:"volumes"`
	VpcId   string                       `pulumi:"vpcId"`
}

func LookupOpengaussInstanceOutput(ctx *pulumi.Context, args LookupOpengaussInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupOpengaussInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpengaussInstanceResult, error) {
			args := v.(LookupOpengaussInstanceArgs)
			r, err := LookupOpengaussInstance(ctx, &args, opts...)
			var s LookupOpengaussInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpengaussInstanceResultOutput)
}

// A collection of arguments for invoking getOpengaussInstance.
type LookupOpengaussInstanceOutputArgs struct {
	// Specifies the name of the instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the instance. If omitted, the provider-level region will
	// be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Specifies the VPC ID.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupOpengaussInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpengaussInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getOpengaussInstance.
type LookupOpengaussInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupOpengaussInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpengaussInstanceResult)(nil)).Elem()
}

func (o LookupOpengaussInstanceResultOutput) ToLookupOpengaussInstanceResultOutput() LookupOpengaussInstanceResultOutput {
	return o
}

func (o LookupOpengaussInstanceResultOutput) ToLookupOpengaussInstanceResultOutputWithContext(ctx context.Context) LookupOpengaussInstanceResultOutput {
	return o
}

// Indicates the availability zone where the node resides.
func (o LookupOpengaussInstanceResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Indicates the advanced backup policy. Structure is documented below.
func (o LookupOpengaussInstanceResultOutput) BackupStrategies() GetOpengaussInstanceBackupStrategyArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []GetOpengaussInstanceBackupStrategy { return v.BackupStrategies }).(GetOpengaussInstanceBackupStrategyArrayOutput)
}

// Indicates the count of coordinator node.
func (o LookupOpengaussInstanceResultOutput) CoordinatorNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) int { return v.CoordinatorNum }).(pulumi.IntOutput)
}

// Indicates the database information. Structure is documented below.
func (o LookupOpengaussInstanceResultOutput) Datastores() GetOpengaussInstanceDatastoreArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []GetOpengaussInstanceDatastore { return v.Datastores }).(GetOpengaussInstanceDatastoreArrayOutput)
}

// Indicates the default username.
func (o LookupOpengaussInstanceResultOutput) DbUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.DbUserName }).(pulumi.StringOutput)
}

// Indicates the enterprise project id.
func (o LookupOpengaussInstanceResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Indicates the instance specifications.
func (o LookupOpengaussInstanceResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// Indicates the instance ha information. Structure is documented below.
func (o LookupOpengaussInstanceResultOutput) Has() GetOpengaussInstanceHaArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []GetOpengaussInstanceHa { return v.Has }).(GetOpengaussInstanceHaArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOpengaussInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates the maintenance window.
func (o LookupOpengaussInstanceResultOutput) MaintenanceWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.MaintenanceWindow }).(pulumi.StringOutput)
}

// Indicates the node name.
func (o LookupOpengaussInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the instance nodes information. Structure is documented below.
func (o LookupOpengaussInstanceResultOutput) Nodes() GetOpengaussInstanceNodeArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []GetOpengaussInstanceNode { return v.Nodes }).(GetOpengaussInstanceNodeArrayOutput)
}

// Indicates the database port.
func (o LookupOpengaussInstanceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) int { return v.Port }).(pulumi.IntOutput)
}

// Indicates the list of private IP address of the nodes.
func (o LookupOpengaussInstanceResultOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []string { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

// Indicates the public IP address of the DB instance.
func (o LookupOpengaussInstanceResultOutput) PublicIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []string { return v.PublicIps }).(pulumi.StringArrayOutput)
}

func (o LookupOpengaussInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the replica num.
func (o LookupOpengaussInstanceResultOutput) ReplicaNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) int { return v.ReplicaNum }).(pulumi.IntOutput)
}

// Indicates the security group ID.
func (o LookupOpengaussInstanceResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Indicates the sharding num.
func (o LookupOpengaussInstanceResultOutput) ShardingNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) int { return v.ShardingNum }).(pulumi.IntOutput)
}

// Indicates the node status.
func (o LookupOpengaussInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupOpengaussInstanceResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Indicates the switch strategy.
func (o LookupOpengaussInstanceResultOutput) SwitchStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.SwitchStrategy }).(pulumi.StringOutput)
}

// Indicates the default username.
func (o LookupOpengaussInstanceResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// Indicates the volume type. Value options: **ULTRAHIGH**, **ESSD**.
func (o LookupOpengaussInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Indicates the volume information. Structure is documented below.
func (o LookupOpengaussInstanceResultOutput) Volumes() GetOpengaussInstanceVolumeArrayOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) []GetOpengaussInstanceVolume { return v.Volumes }).(GetOpengaussInstanceVolumeArrayOutput)
}

func (o LookupOpengaussInstanceResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpengaussInstanceResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpengaussInstanceResultOutput{})
}
