package api

import (
	"context"
)

type API struct{}

func NewAPI(_ context.Context) *API {

	return &API{}
}
