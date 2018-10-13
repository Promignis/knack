#!/bin/bash

app_name="zeno"
base_dir="dist/$app_name.app/Contents/MacOS"
contents_dir="dist/$app_name.app/Contents"
osx_defaults_dir="osx-default/"
mkdir -p ${base_dir}
go build -o "${base_dir}/$app_name"
cp -r js-runtime "${base_dir}/."
cp -r styles "${base_dir}/."
cp -r js "${base_dir}/."
cp -r views "${base_dir}/."
cp manifest.json "${base_dir}/."
cp "${osx_defaults_dir}Info.plist" ${contents_dir}
cp -r "${osx_defaults_dir}Resources" ${contents_dir}
