ifneq ("$(wildcard .env)","")
	include .env
	export
endif

PACKAGES ?= ./...

## help: Display list of commands
.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sort

## build: build all components for prod
.PHONY: build
build: go-build

## lint: golang
.PHONY: lint
lint: go-lint

## go-dev: build app golang
.PHONY: go-dev
go-dev:
	go run cmd/webapp/main.go

## go-migrate-up: run cli debug
.PHONY: go-migrate-up
go-migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ./db/migrations up

## go-lint: lint golang
.PHONY: go-lint
go-lint:
	golangci-lint run $(PACKAGES)

## go-build: build app golang
.PHONY: go-build
go-build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o bin/waldo cmd/webapp/main.go

## go-mocks: clean and generate mocks
.PHONY: go-mocks
go-mocks:
	find . -name "mocks" -type d -exec rm -r "{}" \+
	go generate -run mockgen $(PACKAGES)

## go-tests: run test on golang
.PHONY: go-tests
go-tests:
	go test -cover $(PACKAGES)

## go-tests-coverage: run test on golang with cover output
.PHONY: go-tests-coverage
go-tests-coverage:
	go test -coverprofile=cover.out $(PACKAGES) && go tool cover -html=cover.out

## go-tests-dbless: run test without DB
.PHONY: v
go-tests-dbless:
	go test $$(go list $(PACKAGES) | grep -v memoriesbox/pkg/db/models)

## deps-run: start sidecars
.PHONY: deps-run
deps-run:
	docker-compose -f docker-compose-dev.yml up -d

## deps-logs: get logs sidecars
.PHONY: deps-logs
deps-logs:
	docker-compose -f docker-compose-dev.yml logs -f

## db-migrate-up: launch migrate up
.PHONY: db-migrate-up
db-migrate-up:
	migrate -database "${POSTGRESQL_URL}" -path ./db/migrations up

## db-migrate-down: revert last migrate
.PHONY: db-migrate-down
db-migrate-down:
	migrate -database "${POSTGRESQL_URL}"" -path ./db/migrations down

## db-sqlboiler: update db models
.PHONY: db-sqlboiler
db-sqlboiler:
	sqlboiler psql

## db-migrate-create: add a new migration files (make db-migrate-create name=foo)
.PHONY: db-migrate-create
db-migrate-create:
	migrate create -ext sql -dir db/migrations -digits=10 -seq $(name)

## docker-build: build container
.PHONY: docker-build
docker-build:
	docker buildx build --platform linux/amd64,linux/arm,linux/arm64 .

## nerdctl-build: build container
.PHONY: docker-build
nerdctl-build:
	nerdctl build --platform=amd64,arm64 .
