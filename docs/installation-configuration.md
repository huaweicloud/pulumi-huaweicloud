---
title: HuaweiCloud Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi HuaweiCloud Provider.
layout: installation
---

## Installation

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

### .NET

> *TBA*

## Configuring Credentials

The Pulumi HuaweiCloud Provider needs to be configured with HuaweiCloud credentials before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

+ 1. Set the environment variables HW_ACCESS_KEY and HW_SECRET_KEY:

```bash
export HW_ACCESS_KEY=XXXXXXXXXXXXXX
export HW_SECRET_KEY=YYYYYYYYYYYYYY
```

+ 2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

```bash
pulumi config set huaweicloud:access_key XXXXXXXXXXXXXX --secret
pulumi config set huaweicloud:secret_key YYYYYYYYYYYYYY --secret
```

Remember to pass --secret when setting huaweicloud:access_key and huaweicloud:secret_key so that they are properly encrypted.
The complete list of configuration parameters is in the
[HuaweiCloud provider](https://github.com/huaweicloud/pulumi-huaweicloud/blob/main/README.md).
