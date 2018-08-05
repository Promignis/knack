#!/bin/bash

base_dir="dist/zeno.app/Contents/MacOS"
contents_dir="dist/zeno.app/Contents"
mkdir -p ${base_dir}
go build -o "${base_dir}/zeno"
cp -r runtime "${base_dir}/."
cp -r js "${base_dir}/."
cp -r views "${base_dir}/."
cp Info.plist "${contents_dir}"
cp -r Resources ${contents_dir}
