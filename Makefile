.PHONY: tools openapi go

GO_RACE_DETECTOR ?= false

tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.2
	go install ariga.io/atlas/cmd/atlas@v0.13.1
	go install github.com/sqlc-dev/sqlc/cmd/sqlc

openapi: openapi/ internal/generated/openapi
	rm -f internal/generated/openapi/models/* internal/generated/openapi/client/* internal/generated/openapi/server/*
	oapi-codegen --config openapi/models.yaml openapi/openapi.yaml
	oapi-codegen --config openapi/client.yaml openapi/openapi.yaml
	oapi-codegen --config openapi/server.yaml openapi/openapi.yaml

go: openapi
	go build -race="${GO_RACE_DETECTOR}" -ldflags="-s -w" -o dist/ ./...
