# Huaweicloud Pulumi Provider

<!-- markdownlint-disable-next-line MD033 -->
<a href="https://www.huaweicloud.com/"><img width="450px" height="102px" src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg"></a>

The Huaweicloud Pulumi Provider lets you manage [Huaweicloud](https://www.huaweicloud.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @huaweicloudos/huaweicloud
```

or `yarn`:

```bash
yarn add @huaweicloudos/huaweicloud
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_huaweicloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/huaweicloud/pulumi-huaweicloud/sdk/go/...
```

## Configuration

The following configuration points are available for the `huaweicloud` provider:

- `huaweicloud:region` - (Optional) This is the Huawei Cloud region. It must be provided when using `static credentials`
  authentication, but it can also be sourced from the `HW_REGION_NAME` environment variables.

- `huaweicloud:access_key` - (Optional) The access key of the HuaweiCloud to use. If omitted, the `HW_ACCESS_KEY` environment
  variable is used.

- `huaweicloud:secret_key` - (Optional) The secret key of the HuaweiCloud to use. If omitted, the `HW_SECRET_KEY` environment
  variable is used.

- `huaweicloud:shared_config_file` - (Optional) The path to the shared config file. If omitted, the `HW_SHARED_CONFIG_FILE`
  environment variable is used.

- `huaweicloud:profile` - (Optional) The profile name as set in the shared config file. If omitted, the `HW_PROFILE` environment
  variable is used. Defaults to the `current` profile in the shared config file.

- `huaweicloud:assume_role` - (Optional) Configuration block for an assumed role. See below. Only one assume_role
  block may be in the configuration.

- `huaweicloud:project_name` - (Optional) The Name of the project to login with. If omitted, the `HW_PROJECT_NAME` environment
  variable or `region` is used.

- `huaweicloud:domain_name` - (Optional) The [Account name](https://support.huaweicloud.com/en-us/usermanual-iam/iam_01_0552.html)
  of IAM to scope to. If omitted, the `HW_DOMAIN_NAME` environment variable is used.

- `huaweicloud:security_token` - (Optional) The security token to authenticate with a
  [temporary security credential](https://support.huaweicloud.com/intl/en-us/iam_faq/iam_01_0620.html). If omitted,
  the `HW_SECURITY_TOKEN` environment variable is used.

- `huaweicloud:cloud` - (Optional) The endpoint of the cloud provider. If omitted, the
  `HW_CLOUD` environment variable is used. Defaults to `myhuaweicloud.com`.

- `huaweicloud:auth_url` - (Optional, Required before 1.14.0) The Identity authentication URL. If omitted, the
  `HW_AUTH_URL` environment variable is used. Defaults to `https://iam.{{region}}.{{cloud}}:443/v3`.

- `huaweicloud:insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the
  `HW_INSECURE` environment variable is used.

- `huaweicloud:max_retries` - (Optional) This is the maximum number of times an API call is retried, in the case where
  requests are being throttled or experiencing transient failures. The delay between the subsequent API calls increases
  exponentially. The default value is `5`. If omitted, the `HW_MAX_RETRIES` environment variable is used.

- `huaweicloud:enterprise_project_id` - (Optional) Default Enterprise Project ID for supported resources.
  If omitted, the `HW_ENTERPRISE_PROJECT_ID` environment variable is used.

- `huaweicloud:endpoints` - (Optional) Configuration block in key/value pairs for customizing service endpoints. The following
  endpoints support to be customized: autoscaling, ecs, ims, vpc, nat, evs, obs, sfs, cce, rds, dds, iam. An example
  provider configuration:

  ```hcl
  provider "huaweicloud" {
    ...
    endpoints = {
      ecs = "https://ecs-customizing-endpoint.com"
    }
  }
  ```

The `assume_role` block supports:

- `agency_name` - (Required) The name of the agency for assume role.
  If omitted, the `HW_ASSUME_ROLE_AGENCY_NAME` environment variable is used.
- `domain_name` - (Required) The name of the agency domain for assume role.
  If omitted, the `HW_ASSUME_ROLE_DOMAIN_NAME` environment variable is used.

## Reference

For detailed reference documentation, please visit [the Huaweicloud Pulumi Provider](https://huaweicloud.github.io/pulumi-huaweicloud/).
