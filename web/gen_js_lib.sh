#!/bin/bash
# leehom Chen clh021@gmail.com
cd "$( dirname "${BASH_SOURCE[0]}" )"
currentPath=`pwd`
ProjectDir=$(dirname "${currentPath}")
webDir="${ProjectDir}/web"
serverDir="${ProjectDir}/server"
PROTOC_GEN_TS_PATH=`npx which protoc-gen-ts`
OUT_DIR="${webDir}/src/lib"
mkdir -p ${OUT_DIR}
protoc --plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} --js_out=import_style=commonjs,binary:${OUT_DIR} --ts_out=${OUT_DIR} -I=${serverDir} ${serverDir}/shell.proto
