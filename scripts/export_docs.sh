#!/bin/bash

mkdir dist
cp -r docs/public/css/ dist/
cp -r docs/public/fonts/ dist/
cp -r docs/public/icons/ dist/
cp -r docs/public/images/ dist/
cp -r docs/public/js/ dist/
cp -r docs/public/logos/ dist/
mkdir -p dist/registry/packages
cp -r docs/public/registry/packages/huaweicloud/ dist/registry/packages
mkdir -p dist/registry/packages/navs
cp -r docs/public/registry/packages/navs/huaweicloud.json dist/registry/packages/navs
cp -r pulumi-huaweicloud/assets/index.html dist/index.html

sed -i 's+/fonts/+/pulumi-huaweicloud/fonts/+g' dist/css/bundle.*
sed -i 's+/registry/+/pulumi-huaweicloud/registry/+g' dist/js/bundle.*
sed -i 's+/logos/+/pulumi-huaweicloud/logos/+g' dist/js/bundle.*
sed -i 's+/icons/+/pulumi-huaweicloud/icons/+g' dist/js/bundle.*
sed -i 's+/images/+/pulumi-huaweicloud/images/+g' dist/css/bundle.*

sed -i 's+<a href=https://github.com/huaweicloud/pulumi-huaweicloud/tree/main/examples>pulumi-huaweicloud repository</a>+<a target=_blank href=https://github.com/huaweicloud/pulumi-huaweicloud/tree/main/examples>pulumi-huaweicloud repository</a>+g' dist/registry/packages/huaweicloud/index.html
sed -i 's+<a href=https://github.com/huaweicloud/pulumi-huaweicloud/blob/main/README.md>HuaweiCloud provider</a>+<a target=_blank href=https://github.com/huaweicloud/pulumi-huaweicloud/blob/main/README.md>HuaweiCloud provider</a>+g' dist/registry/packages/huaweicloud/installation-configuration/index.html

find dist/registry/packages/huaweicloud/api-docs/ -type f -name "*.html" -exec sed -i 's+base-directory=/registry/+base-directory=/pulumi-huaweicloud/registry/+g' {} +
find dist/registry/packages/huaweicloud/api-docs/ -type f -name "*.html" -exec sed -i 's+<a href=https://github.com/huaweicloud/pulumi-huaweicloud>+<a target=_blank href=https://github.com/huaweicloud/pulumi-huaweicloud>+g' {} +
find dist/registry/packages/huaweicloud/api-docs/ -type f -name "*.html" -exec sed -i 's+<a href=https://github.com/huaweicloud/terraform-provider-huaweicloud>+<a target=_blank href=https://github.com/huaweicloud/terraform-provider-huaweicloud>+g' {} +

echo >> dist/css/bundle.*
echo '.footer, .top-nav-bar, .hide-on-pinned, .dot-background-container, .header-container, .breadcrumb {display: none !important;}' >> dist/css/bundle.*
