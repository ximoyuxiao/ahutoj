# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

WORKDIR /build

COPY --link ./ ./

RUN  go env -w GO111MODULE=on && go env -w GOPROXY=goproxy.cn,direct && go mod tidy && \
    go build -o ./gateway web/service/gateway/gateway.go && \
    go build -o ./oj web/service/ahutoj/ahutoj.go && \
    go build -o ./persistence web/service/persistence/persistence.go && \
    go build -o ./originproblem web/service/originproblem/originproblem.go && \
    go build -o ./originJudge web/service/originJudge/originJudge.go && \
    go build -o ./oss web/service/oss/oss.go

FROM alpine:3.16 as gateway

WORKDIR /app

COPY --link --from=build /build/gateway /usr/bin/gateway

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x /usr/bin/gateway

EXPOSE 4433

ENTRYPOINT ["/usr/bin/gateway"]

FROM alpine:3.16 as problem

WORKDIR /app

COPY --link --from=build /build/originproblem /usr/bin/originproblem

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x /usr/bin/originproblem

ENTRYPOINT ["/usr/bin/originproblem"]

FROM alpine:3.16 as persistence

WORKDIR /app

COPY --link --from=build /build/persistence /usr/bin/persistence

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x /usr/bin/persistence

ENTRYPOINT ["/usr/bin/persistence"]

FROM alpine:3.16 as origin

WORKDIR /app

COPY --link --from=build /build/originJudge /usr/bin/originJudge

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x /usr/bin/originJudge

ENTRYPOINT ["/usr/bin/originJudge"]

FROM alpine:3.16 as oj

WORKDIR /app

COPY --link --from=build /build/oj /usr/bin/oj

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    touch ahutoj.log &&\
    chmod +x /usr/bin/oj

EXPOSE 4212

ENTRYPOINT ["/usr/bin/oj"]

FROM alpine:3.16 as user

WORKDIR /app

COPY --link --from=build /build/useranalytics /usr/bin/useranalytics

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    chmod +x /usr/bin/useranalytics

ENTRYPOINT ["/usr/bin/useranalytics"]

FROM alpine:3.16 as oss

WORKDIR /app

COPY --link --from=build /build/oss /usr/bin/oss

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update && \
    hmod +x /usr/bin/oss

EXPOSE 4466

ENTRYPOINT ["/usr/bin/oss"]