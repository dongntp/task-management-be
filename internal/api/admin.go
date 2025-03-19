package api

import (
	"context"

	"task-management-be/internal/generated/openapi/server"
	"task-management-be/internal/generated/sql"
)

func (a *API) CreateAccount(ctx context.Context, request server.CreateAccountRequestObject) (server.CreateAccountResponseObject, error) {
	newAcc := *request.Body

	param := sql.InsertNewAccountParams{
		Username: newAcc.Username,
		Password: newAcc.Password,
		Role:     sql.RoleEnum(newAcc.Role),
	}

	err := a.DBClient.InsertNewAccount(ctx, param)
	if err != nil {
		return server.CreateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Username already exists")}, nil
	}

	return server.CreateAccount200JSONResponse(true), nil
}
