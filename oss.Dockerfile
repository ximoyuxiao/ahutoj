
FROM alpine:3.16 as oj

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories && apk update
#  && apk --no-cache add hiredis rabbitmq-c  mysql-dev nlohmann-json

WORKDIR /app

COPY --link --from=build /build/oss /usr/bin/oss

COPY --link ./config/config.yaml.bak /app/config.yaml

RUN chmod +x /usr/bin/oss

EXPOSE 4466

ENTRYPOINT ["/usr/bin/oss"]