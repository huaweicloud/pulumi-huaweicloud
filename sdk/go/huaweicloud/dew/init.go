// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dew

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
	case "huaweicloud:Dew/key:Key":
		r = &Key{}
	case "huaweicloud:Dew/keypair:Keypair":
		r = &Keypair{}
	case "huaweicloud:Dew/secret:Secret":
		r = &Secret{}
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
		"Dew/key",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Dew/keypair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Dew/secret",
		&module{version},
	)
}