#!/bin/bash

base_dir="dist/zeno.app/Contents/MacOS"
mkdir -p ${base_dir}
go build -o "${base_dir}/zeno"
cp -r runtime "${base_dir}/."
cp -r js "${base_dir}/."
cp -r views "${base_dir}/."
