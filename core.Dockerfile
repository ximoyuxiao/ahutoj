# syntax=docker/dockerfile:1.4

FROM alpine:3.16 as build

RUN --mount=type=cache,target=/var/cache/apk \
  sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update \
    && apk add gcompat gcc g++ make musl \
    && apk add hiredis-dev rabbitmq-c-dev  mysql-dev  \
    && apk add --no-cache nlohmann-json

WORKDIR /build

COPY  --link ./core .

#COPY --link ./go/* .

RUN make judged

FROM alpine:3.16 as image

RUN --mount=type=cache,target=/var/cache/apk \
    apk update \
    && apk add hiredis rabbitmq-c  mysql-dev \
    && apk add --no-cache nlohmann-json

WORKDIR /app

COPY --from=build /build/judged /app/judged
COPY ./core/config.conf /app/config.conf
#COPY --link ./config/config.conf /app/config.conf

RUN chmod +x /app/judged
# ENTRYPOINT ["tail", "-f", "/dev/null"]

ENTRYPOINT ["/app/judged"]

