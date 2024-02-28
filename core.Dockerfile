# syntax=docker/dockerfile:1.4

FROM alpine:3.16 as build

# 如果有问题把上下的这句都给去掉，--mount
#RUN --mount=type=cache,target=/var/cache/apk
RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories &&  \
    apk update && \
    apk add gcompat gcc g++ make musl && \
    apk add hiredis-dev rabbitmq-c-dev  mysql-dev && \
    apk add --no-cache nlohmann-json

WORKDIR /build

COPY   ./core .

RUN make judged

FROM alpine:3.16 as image

WORKDIR /app

#RUN --mount=type=cache,target=/var/cache/apk
RUN   sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories &&  \
    apk update && \
    apk add hiredis rabbitmq-c  mysql-dev && \
    apk add --no-cache nlohmann-json \

#RUN apk add gdb ncurses-libs python3 expat
#GDB调试所需环境

COPY --from=build /build/judged /app/judged

COPY  ./core/config.conf /app/config.conf

RUN chmod +x /app/judged

ENTRYPOINT  /app/judged && tail -f /dev/null
# 避免容器退出

