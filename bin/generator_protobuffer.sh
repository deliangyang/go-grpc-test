#!/usr/bin/env bash

BASE="`dirname $0`/.."
SRC="$BASE/proto"

rm -rf ${BASE}/php/*
rm -rf ${BASE}/go/*

mkdir -p ${BASE}/php
mkdir -p ${BASE}/go

for FILE in `find ${SRC} -type f -name "*.proto"`
do
    echo "generate ${FILE} ..."
    protoc --proto_path=${SRC} \
        --go_out=plugins=grpc:${BASE}/go \
        --php_out=${BASE}/php \
        --grpc_out=${BASE}/php \
        --plugin=protoc-gen-grpc=`which grpc_php_plugin` \
        ${FILE}
done
