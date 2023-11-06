# syntax=docker/dockerfile:1.4

FROM golang:alpine as build

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update \
#    && apk --no-cache add gcompat gcc g++ make musl

WORKDIR /build

COPY --link ./go.* .

RUN go env -w GOPROXY=goproxy.cn && go mod tidy

COPY --link . .

RUN go build -o ./gateway web/service/gateway/gateway.go

RUN go build -o ./oj web/service/ahutoj/ahutoj.go

RUN go build -o ./persistence web/service/persistence/persistence.go

RUN go build -o ./originproblem web/service/originproblem/originproblem.go

RUN go build -o ./originJudge web/service/originJudge/originJudge.go

RUN go build -o ./oss web/service/oss/oss.go

FROM alpine:3.16 as gateway

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/gateway /usr/bin/gateway

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/gateway

EXPOSE 4433

ENTRYPOINT ["/usr/bin/gateway"]

FROM alpine:3.16 as problem

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/originproblem /usr/bin/originproblem

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/originproblem

ENTRYPOINT ["/usr/bin/originproblem"]

FROM alpine:3.16 as persistence

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/persistence /usr/bin/persistence

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/persistence

ENTRYPOINT ["/usr/bin/persistence"]

FROM alpine:3.16 as origin

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/originJudge /usr/bin/originJudge

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/originJudge

ENTRYPOINT ["/usr/bin/originJudge"]

FROM alpine:3.16 as oj

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/oj /usr/bin/oj

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/oj

EXPOSE 4212

ENTRYPOINT ["/usr/bin/oj"]

FROM alpine:3.16 as user

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/useranalytics /usr/bin/useranalytics

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/useranalytics
#ENTRYPOINT ["tail", "-f", "/dev/null"]

ENTRYPOINT ["/usr/bin/useranalytics"]

FROM alpine:3.16 as oss

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/oss /usr/bin/oss

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/oss

EXPOSE 4466

ENTRYPOINT ["/usr/bin/oss"]