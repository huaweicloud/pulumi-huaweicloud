// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

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
	case "huaweicloud:Elb/certificate:Certificate":
		r = &Certificate{}
	case "huaweicloud:Elb/l7policy:L7policy":
		r = &L7policy{}
	case "huaweicloud:Elb/l7rule:L7rule":
		r = &L7rule{}
	case "huaweicloud:Elb/listener:Listener":
		r = &Listener{}
	case "huaweicloud:Elb/loadbalancer:Loadbalancer":
		r = &Loadbalancer{}
	case "huaweicloud:Elb/member:Member":
		r = &Member{}
	case "huaweicloud:Elb/monitor:Monitor":
		r = &Monitor{}
	case "huaweicloud:Elb/pool:Pool":
		r = &Pool{}
	case "huaweicloud:Elb/whitelist:Whitelist":
		r = &Whitelist{}
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
		"Elb/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/l7policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/l7rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/loadbalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/member",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/monitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/pool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Elb/whitelist",
		&module{version},
	)
}
