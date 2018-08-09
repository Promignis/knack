#!/bin/bash

base_dir="dist/zeno.app/Contents/MacOS"
contents_dir="dist/zeno.app/Contents"
osx_defaults_dir="osx-default/"
mkdir -p ${base_dir}
go build -o "${base_dir}/zeno"
cp -r runtime "${base_dir}/."
cp -r js "${base_dir}/."
cp -r views "${base_dir}/."
cp "${osx_defaults_dir}Info.plist" ${contents_dir}
cp -r "${osx_defaults_dir}Resources" ${contents_dir}
