# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

WORKDIR /build

COPY  ./web ./web

COPY ./go.* ./

RUN  go env -w GO111MODULE=on && go env -w GOPROXY=goproxy.cn,direct && go mod tidy && \
    go build -o ./gateway web/service/gateway/gateway.go && \
    go build -o ./oj web/service/ahutoj/ahutoj.go && \
    go build -o ./persistence web/service/persistence/persistence.go && \
    go build -o ./originproblem web/service/originproblem/originproblem.go && \
    go build -o ./originJudge web/service/originJudge/originJudge.go && \
    go build -o ./oss web/service/oss/oss.go

FROM alpine:3.16 as gateway

WORKDIR /app

COPY  --from=build /build/gateway ./gateway

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
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

FROM alpine:3.16 as persistence

WORKDIR /app

COPY  --from=build /build/persistence ./persistence

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
RUN   touch ahutoj.log &&\
    chmod +x ./persistence

ENTRYPOINT ["/app/persistence"]

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

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
RUN touch ahutoj.log &&\
    chmod +x ./oj

EXPOSE 4212

ENTRYPOINT ["/app/oj"]

FROM alpine:3.16 as user

WORKDIR /app

COPY  --from=build /build/useranalytics ./useranalytics

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x ./useranalytics

ENTRYPOINT ["/app/useranalytics"]

FROM alpine:3.16 as oss

WORKDIR /app

COPY  --from=build /build/oss ./oss

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    hmod +x ./oss

EXPOSE 4466

ENTRYPOINT ["/app/oss"]