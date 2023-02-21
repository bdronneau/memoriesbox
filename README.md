[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_memoriesbox&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bdronneau_memoriesbox) [![Golang Build](https://github.com/bdronneau/memoriesbox/actions/workflows/build.yaml/badge.svg)](https://github.com/bdronneau/memoriesbox/actions/workflows/build.yaml) ![Docker Pulls](https://img.shields.io/docker/pulls/bdronneau/memoriesbox) [![docker Build](https://github.com/bdronneau/memoriesbox/actions/workflows/docker.yaml/badge.svg)](https://github.com/bdronneau/memoriesbox/actions/workflows/docker.yaml)

# Memories box

Based on Bonjour poney mojo, this project is memories oriented.

## Configuration

This application used [ff](https://github.com/peterbourgon/ff/) in order to get configuration (prefix `MEMORIESBOX`).

Check [env_example](./.env_example) for main variables.

## Development

### Dependencies

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [SQLBoiler](github.com/volatiletech/sqlboiler) is used for generate models based on PostgreSQL.
- [Migrate](https://github.com/golang-migrate/migrate) is used to follow update on schema database.

### Application dependencies

Dependencies are handle in docker.

```shell script
docker-compose -f docker-compose-dev.yml up -d
```

In order to migrate down you can run:

```shell script
docker-compose -f docker-compose-dev.yml up --scale migrate_up=0 --scale migrate_down=1 -d
```

### SQLBoiler

Run [SQLBoiler](github.com/volatiletech/sqlboiler) to update generated models

```shell script
PGPASSWORD=passwordToChange sqlboiler psql
```

### Run

```shell script
go run cmd/webapp/main.go
```

### Lint

```shell script
golangci-lint run ./...
```

### Build

```shell script
CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o bin/memoriesbox cmd/webapp/main.go
```

### Docker build

```shell sceript
docker buildx build --platform linux/amd64,linux/arm,linux/arm64 .
```

or

```shell script
nerdctl build --platform=amd64,arm64 .
```
