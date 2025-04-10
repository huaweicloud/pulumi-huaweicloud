#!/bin/bash

sed -i 's+github.com/pulumi/registry+github.com/chnsz/registry+g' docs/config/_default/config.yml
sed -i 's+https://github.com/pulumi/registry+https://github.com/chnsz/registry+g' docs/tools/resourcedocsgen/cmd/docs/docs.go
sed -i '26,28d' docs/scripts/build-site.sh
sed -i 's+go build -o "${GOPATH}/bin/resourcedocsgen" .+go build .+g' docs/scripts/build-site.sh
sed -i 's+resourcedocsgen docs registry --commitSha "${REGISTRY_COMMIT}" --logtostderr+./resourcedocsgen docs registry --logtostderr+g' docs/scripts/build-site.sh
sed -i '3iexport HUGO_BASEURL="https://huaweicloud.github.io/pulumi-huaweicloud/"' docs/scripts/build-site.sh
