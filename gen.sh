#!/bin/bash

# feed kitex_gen generation
echo "generate feed kitex_gen"
kitex --module "nico_minidouyin" -I idl/ idl/douyin/rpc/feed.proto

# feed service generation
echo "generate feed service"
cd service/feed && kitex --module "nico_minidouyin" --service feed --use nico_minidouyin/kitex_gen/ -I ../../idl/ ../../idl/douyin/rpc/feed.proto
