#!/bin/bash

base_dir="dist/zeno.app/Contents/MacOS"
rm -rf ${base_dir} && ./build-app.sh
