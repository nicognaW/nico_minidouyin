#!/bin/bash

# 判断 gen 文件夹是否存在
if [ -d "gen" ]; then
  echo "gen文件夹已存在"
else
  echo "gen文件夹不存在，创建"
  # 新建 gen 文件夹
  mkdir gen
fi

# 判断web文件夹是否存在
if [ -d "service/web" ]; then
  echo "基础服务已创建，执行更新逻辑"
  mv service/web/go.mod.bak service/web/go.mod
  cd service/web && hz update \
    --model_dir gen \
    --proto_path ../../idl \
    --idl ../../idl/douyin/"$1".proto && cd ../../
else
  echo "基础服务未创建，执行创建逻辑"
  # 新建web文件夹
  mkdir service/web

  mv go.mod go.mod.bak

  hz new \
    --module nico_minidouyin/service/web \
    --service web \
    --model_dir gen \
    --out_dir service/web \
    --proto_path idl \
    --idl idl/douyin/"$1".proto

  mv go.mod.bak go.mod
  rm service/web/.gitignore
fi

mv service/web/go.mod service/web/go.mod.bak
cp -R service/web/gen/ gen/
rm -rf service/web/gen
mkdir service/web/gen
cp -R gen/api service/web/gen/
rm -rf gen/api

for file in service/web/biz/handler/"$1"/*_service.go; do
  echo 修正 "$file"
  # 使用sed命令替换字符串
  sed 's/nico_minidouyin\/service\/web\/gen\/douyin\//nico_minidouyin\/gen\/douyin\//g' $file >$file"_new"
  # 删除原有文件
  rm $file
  # 重命名替换后的文件
  mv $file"_new" $file
done

cd gen &&
  protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path ../idl/ ../idl/douyin/"$1".proto &&
  mv douyin/"$1"_grpc.pb.go douyin/"$1" &&
  cd ../
