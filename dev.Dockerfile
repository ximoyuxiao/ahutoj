# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

WORKDIR /build

COPY  ./web ./web

COPY ./go.* ./

RUN  go env -w GO111MODULE=on && go env -w GOPROXY=goproxy.cn,direct && go mod tidy && \
    go build -o ./gateway web/service/gateway/gateway.go && \
    go build -o ./oj web/service/ahutoj/ahutoj.go && \
    go build -o ./originproblem web/service/originproblem/originproblem.go && \
    go build -o ./originJudge web/service/originJudge/originJudge.go &&\
    go build -o ./forum web/service/forum/forum.go

FROM alpine:3.16 as gateway

WORKDIR /app

COPY  --from=build /build/gateway ./gateway

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
RUN touch ahutoj.log &&\
    chmod +x ./gateway

EXPOSE 4433

ENTRYPOINT ["/app/gateway"]

FROM alpine:3.16 as problem

WORKDIR /app

COPY  --from=build /build/originproblem ./originproblem

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
RUN touch ahutoj.log &&\
    chmod +x ./originproblem

ENTRYPOINT ["/app/originproblem"]

FROM alpine:3.16 as origin

WORKDIR /app

COPY  --from=build /build/originJudge ./originJudge

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
RUN touch ahutoj.log &&\
    chmod +x ./originJudge

ENTRYPOINT ["/app/originJudge"]

FROM alpine:3.16 as oj

WORKDIR /app

COPY  --from=build /build/oj ./oj

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && \
    touch ahutoj.log &&\
    chmod +x ./oj &&\
    apk add --no-cache curl

HEALTHCHECK --interval=10s CMD curl --fail http://gateway:4433/api/ping/ || exit 1
#配置探针用于注册路由

EXPOSE 4212

ENTRYPOINT ["/app/oj"]

FROM alpine:3.16 as forum

WORKDIR /app

COPY  --from=build /build/forum ./forum

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && \
    touch ahutoj.log &&\
    chmod +x ./forum &&\
    apk add --no-cache curl

EXPOSE  4498

HEALTHCHECK --interval=10s CMD curl --fail http://gateway:4433/api/solution/ping || exit 1

ENTRYPOINT ["/app/forum"]

