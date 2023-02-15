#!/bin/bash

sed -i 's+github.com/pulumi/registry+github.com/Jason-Zhang9309/registry+g' docs/config/_default/config.yml
sed -i 's+https://github.com/pulumi/registry+https://github.com/Jason-Zhang9309/registry+g' docs/tools/resourcedocsgen/cmd/docs/docs.go
sed -i '26,28d' docs/scripts/build-site.sh
sed -i 's+resourcedocsgen docs registry --commitSha "${REGISTRY_COMMIT}" --logtostderr+resourcedocsgen docs registry --logtostderr+g' docs/scripts/build-site.sh
