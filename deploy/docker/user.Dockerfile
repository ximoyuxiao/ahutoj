# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update \
#    && apk --no-cache add gcompat gcc g++ make musl

WORKDIR /build

COPY --link ./go.* .

RUN go env -w GOPROXY=goproxy.cn && go mod tidyCOPY --link . .

RUN go build -o ./useranalytics web/service/useranalytics/useranalytics.go

FROM alpine:3.16 as image

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/useranalytics /usr/bin/useranalytics

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/useranalytics
#ENTRYPOINT ["tail", "-f", "/dev/null"]

ENTRYPOINT ["/usr/bin/useranalytics"]

