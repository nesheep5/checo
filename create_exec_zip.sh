#!/bin/bash
rm -rf pkg dist
source compile.sh
zip checo_exec.zip -r ./pkg
mkdir dist
mv checo_exec.zip dist
rm -rf pkg