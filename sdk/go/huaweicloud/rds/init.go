// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "huaweicloud:Rds/account:Account":
		r = &Account{}
	case "huaweicloud:Rds/database:Database":
		r = &Database{}
	case "huaweicloud:Rds/database_privilege:Database_privilege":
		r = &Database_privilege{}
	case "huaweicloud:Rds/instance:Instance":
		r = &Instance{}
	case "huaweicloud:Rds/parametergroup:Parametergroup":
		r = &Parametergroup{}
	case "huaweicloud:Rds/readReplicaInstance:ReadReplicaInstance":
		r = &ReadReplicaInstance{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := huaweicloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/database_privilege",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/parametergroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Rds/readReplicaInstance",
		&module{version},
	)
}
