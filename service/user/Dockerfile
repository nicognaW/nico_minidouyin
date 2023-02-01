# 编译镜像
FROM golang:1.19.5-bullseye AS build
ENV TZ=Asia/Shanghai

# 配置 git
RUN apt update
RUN apt install git

# 获取文件
RUN mkdir -p /source
WORKDIR /source
COPY . .

# 编译
RUN go build -o user /source/service/feed

# 运行环境
FROM gcr.io/distroless/base-debian11
ENV TZ=Asia/Shanghai

# RUN mkdir -p /data/apps/nico-minidouyin-user
WORKDIR /data/apps/nico-minidouyin-user

# 收集数据
COPY --from=build /source .
ENTRYPOINT ["./user"]
