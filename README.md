[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bdronneau_memoriesbox&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bdronneau_memoriesbox) [![CI](https://github.com/bdronneau/memoriesbox/actions/workflows/ci.yaml/badge.svg)](https://github.com/bdronneau/memoriesbox/actions/workflows/ci.yaml) ![Docker Pulls](https://img.shields.io/docker/pulls/bdronneau/memoriesbox) [![docker Build](https://github.com/bdronneau/memoriesbox/actions/workflows/docker.yaml/badge.svg)](https://github.com/bdronneau/memoriesbox/actions/workflows/docker.yaml)

# Memories box

This app store your memories in postgres backend.

## Configuration

This application used [ff](https://github.com/peterbourgon/ff/) in order to get configuration (prefix `MEMORIESBOX`).

Check [env_example](./.env_example) for main variables.

## Development

### Dependencies

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [SQLBoiler](github.com/volatiletech/sqlboiler) is used for generate models based on PostgreSQL.
- [Migrate](https://github.com/golang-migrate/migrate) is used to follow update on schema database.

### Application dependencies

You can use docker with docker-compose to handle deps.

```shell script
docker-compose -f docker-compose-dev.yml up -d
```

In order to migrate down you can run:

```shell script
docker-compose -f docker-compose-dev.yml up --scale migrate_up=0 --scale migrate_down=1 -d
```

### Usage

Run `make help` to have all targets with description

Copy `.env_example` to `.env` with your custom values.

## Credits

- Usage of [Memory loss icons created by Good Ware - Flaticon](https://www.flaticon.com/free-icons/memory-loss)
