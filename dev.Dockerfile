FROM golang:alpine

RUN --mount=type=cache,target=/var/cache/apk \
  sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories \
    && apk update \
    && apk --no-cache add gcompat gcc g++ zsh git curl make

# 安装 poj 编译依赖
RUN  apk --no-cache add hiredis-dev rabbitmq-c-dev  mysql-dev nlohmann-json\
  && go env -w GOPROXY=goproxy.cn

WORKDIR /root