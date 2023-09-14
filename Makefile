# define PATH for db migration
export MIGRATION_DIR_PATH := db/migrations

LOCAL_BIN:=$(CURDIR)/bin-deps
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: clean
clean: ### Clean the build directory
	rm -rf ./bin/*

.PHONY: build
build: clean ### Build the binary file
	go build -o bin/elasticpmapi cmd/main/main.go

.PHONY: debug
debug: ### Debug the main app, you need to attach the client to port 2345
	dlv debug --headless --only-same-user --listen :2345 --api-version 2 ./cmd/main/main.go -- -debug=true

.PHONY: watch
watch: ### Run make build to build the binary file and run it, will restart on file change
	air

.PHONY: migrate
migrate: ### Run the migration
	CGO_ENABLED=0 go run -tags migrate github.com/stewie1520/elasticpmapi/cmd/migrate

.PHONY: gen-migration
gen-migration: ### Generate a new migration file
	migrate create -ext sql -dir ./${MIGRATION_DIR_PATH} -seq $(name)

.PHONY: gen-sqlc
gen-sqlc: ### Generate sqlc
	sqlc generate

.PHONY: linter-golangci
linter-golangci: ### check by golangci linter
	golangci-lint run

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/go-delve/delve/cmd/dlv@latest
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v1.54.2
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(LOCAL_BIN)