run-client:lint
	go run cmd/client/main.go
run-server:lint
	go run cmd/server/main.go
proto:
	protoc --go-grpc_out=proto  proto/ping.proto
build-client:lint
	go build -o bin/client cmd/client/main.go
build-server:lint
	go build -o bin/server cmd/server/main.go
lint:
	golangci-lint run
lint-fix:
	golangci-lint run --fix

markdownlint:
	markdownlint-cli2 --config ./docs/decisions/.markdownlint.yml docs/decisions/*.md

plantuml:
	find ./docs/C4/ -name "*.puml" -type f -exec plantuml {} +
