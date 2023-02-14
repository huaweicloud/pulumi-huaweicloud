---
title: HuaweiCloud
meta_desc: Provides an overview of the HuaweiCloud Provider for Pulumi.
layout: overview
---

The HuaweiCloud provider for Pulumi can be used to provision any of the monitoring resources available in
[HuaweiCloud](https://www.huaweicloud.com/). The HuaweiCloud provider must be configured with the `region`,
`access_key` and `secret_key` in order to deploy HuaweiCloud resources.

## Example

{{< chooser language "javascript,typescript,go,python" >}}

{{% choosable language javascript %}}

```javascript
const huaweicloud = require("@pulumi/huaweicloud");

const vpc = new huaweicloud.VPC("my-vpc", {
    name: "test",
    cidr: "192.168.0.0/16",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as huaweicloud from "@pulumi/huaweicloud";

const vpc = new huaweicloud.VPC("my-vpc", {
    name: "test",
    cidr: "192.168.0.0/16",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_huaweicloud as huaweicloud

vpc = huaweicloud.VPC("my-vpc",
    name = "test",
    cidr = "192.168.0.0/16"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
    huaweicloud "github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := huaweicloud.NewVPC(ctx, "my-vpc", &huaweicloud.VPCArgs{
            Name: pulumi.StringPtr("test"),
            Cidr: pulumi.String("192.168.0.0/16"),
        })
        if err != nil {
            return err
        }
        return nil
    })
}
```

{{% /choosable %}}
{{< /chooser >}}

> You could find more complete and detailed examples in the [pulumi-huaweicloud repository](https://github.com/huaweicloud/pulumi-huaweicloud/tree/main/examples)
