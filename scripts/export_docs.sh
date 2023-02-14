#!/bin/bash

mkdir dist
cp -r docs/public/css/ dist/
cp -r docs/public/fonts/ dist/
cp -r docs/public/images/ dist/
cp -r docs/public/js/ dist/
cp -r docs/public/logos/ dist/
mkdir -p dist/registry/packages
cp -r docs/public/registry/packages/huaweicloud/ dist/registry/packages
mkdir -p dist/registry/packages/navs
cp -r docs/public/registry/packages/navs/huaweicloud.json dist/registry/packages/navs
