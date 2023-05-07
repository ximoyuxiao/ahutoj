# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update \
#    && apk --no-cache add gcompat gcc g++ make musl

WORKDIR /build

COPY --link ./go.* .

RUN go env -w GOPROXY=goproxy.cn && go mod tidy

COPY --link . .

RUN go build -o ./oj web/service/ahutoj/ahutoj.go

FROM alpine:3.16 as image

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/oj /usr/bin/oj

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/oj
#ENTRYPOINT ["tail", "-f", "/dev/null"]

EXPOSE 4212

ENTRYPOINT ["/usr/bin/oj"]

