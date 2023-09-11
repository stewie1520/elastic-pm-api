# define PATH for db migration
export MIGRATION_DIR_PATH := db/migrations

.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: build
build: clean
	go build -o bin/elasticpmapi cmd/main/main.go

.PHONY: debug
debug:
	dlv debug --headless --only-same-user --listen :2345 --api-version 2 ./cmd/main/main.go

.PHONY: watch
watch:
	air

.PHONY: migrate
migrate:
	CGO_ENABLED=0 go run -tags migrate github.com/stewie1520/elasticpmapi/cmd/migrate

.PHONY: gen-migration
gen-migration:
	migrate create -ext sql -dir ./${MIGRATION_DIR_PATH} -seq $(name)

.PHONY: gen-sqlc
gen-sqlc:
	sqlc generate