FROM golang:1.20 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o bin/memoriesbox cmd/webapp/main.go

FROM alpine as fetcher

WORKDIR /app

RUN apk --update add tzdata ca-certificates zip \
 && cd /usr/share/zoneinfo/ \
 && zip -q -r -0 /app/zoneinfo.zip .

FROM scratch

ENV APP_NAME memoriesbox
ENV ZONEINFO /zoneinfo.zip

EXPOSE 1080
ENTRYPOINT [ "/memoriesbox" ]

COPY --from=fetcher /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=fetcher /app/zoneinfo.zip /
COPY --from=builder /app/bin/${APP_NAME} /
