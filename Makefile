.PHONY: tools openapi go

GO_RACE_DETECTOR ?= false

tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.2
	go install github.com/sqlc-dev/sqlc/cmd/sqlc

tools-atlas:
	curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.28.1 sh -s -- --community --yes --no-install --output ./dist/tools/atlas
	chmod +x ./dist/tools/atlas

gen-query:
	rm -f internal/generated/sql/*.go
	sqlc generate

validate-schema:
	db/new-migration.sh DO_NOT_COMMIT_RUN_MANUALLY

reset-schema:
	rm -f db/migrations/*
	db/new-migration.sh init

openapi: openapi/ internal/generated/openapi
	rm -f internal/generated/openapi/models/* internal/generated/openapi/client/* internal/generated/openapi/server/*
	oapi-codegen --config openapi/models.yaml openapi/openapi.yaml
	oapi-codegen --config openapi/client.yaml openapi/openapi.yaml
	oapi-codegen --config openapi/server.yaml openapi/openapi.yaml

go: openapi gen-query
	go build -race="${GO_RACE_DETECTOR}" -ldflags="-s -w" -o dist/ ./...
