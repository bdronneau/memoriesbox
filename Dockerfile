FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o bin/memoriesbox cmd/webapp/main.go

FROM migrate/migrate:v4.17.1 AS migrate

FROM alpine AS fetcher
ENV DOCKERIZE_VERSION v0.6.1

WORKDIR /app

RUN apk --update add tzdata ca-certificates zip \
 && cd /usr/share/zoneinfo/ \
 && zip -q -r -0 /app/zoneinfo.zip .

FROM scratch

USER 2000:2000

ENV APP_NAME memoriesbox
ENV ZONEINFO /zoneinfo.zip

EXPOSE 1080
ENTRYPOINT [ "/memoriesbox" ]

COPY --from=fetcher /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=migrate /usr/local/bin/migrate /usr/local/bin/
COPY --from=fetcher --chown=2000:2000 --chmod=0600 /app/zoneinfo.zip /
COPY --from=builder --chown=2000:2000 /app/bin/${APP_NAME} /
COPY --from=builder --chown=2000:2000 /app/db /opt
