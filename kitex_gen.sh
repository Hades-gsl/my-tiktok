#!/bin/bash

start_dir=$(pwd) 

mkdir -p kitex_gen
mkdir -p service

for file in idl/*.thrift; do
    filename=$(basename "$file" .thrift)
    kitex -module "tiktok" "tiktok/"$file"
    mkdir -p service/"$filename"
    cd service/"$filename"
    kitex -module "tiktok" -service "$filename" -use tiktok/kitex_gen/ tiktok/"$file"
    cd "$start_dir"
done

go mod tidy