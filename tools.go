//go:build tools

//go:generate make openapi

package tools

import (
	_ "ariga.io/atlas/cmd/atlas"
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
