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