package api

import (
	"context"

	"task-management-be/internal/pkg/hash"

	"task-management-be/internal/generated/openapi/server"
	"task-management-be/internal/generated/sql"
)

func (a *API) CreateAccount(ctx context.Context, request server.CreateAccountRequestObject) (server.CreateAccountResponseObject, error) {
	newAcc := *request.Body

	password, err := hash.HashPassword(newAcc.Password)
	if err != nil {
		return server.CreateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid password")}, nil
	}

	param := sql.InsertNewAccountParams{
		Username: newAcc.Username,
		Password: password,
		Role:     sql.RoleEnum(newAcc.Role),
	}

	err = a.DBClient.InsertNewAccount(ctx, param)
	if err != nil {
		return server.CreateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Username already exists")}, nil
	}

	return server.CreateAccount200JSONResponse(true), nil
}

func (a *API) UpdateAccount(ctx context.Context, request server.UpdateAccountRequestObject) (server.UpdateAccountResponseObject, error) {
	reqBody := *request.Body

	curAccount, err := a.DBClient.GetUserByUserName(ctx, reqBody.Username)
	if err != nil {
		return server.UpdateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Username doesn't exists")}, nil
	}

	password := curAccount.Password
	if reqBody.NewPassword != nil {
		password, err = hash.HashPassword(*reqBody.NewPassword)
		if err != nil {
			return server.UpdateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid password")}, nil
		}
	}

	params := sql.UpdateAccountByAdminParams{
		Username:    reqBody.Username,
		Newusername: PointerValueWithDefault(reqBody.NewUsername, reqBody.Username),
		Newpassword: password,
		Newactive:   PointerValueWithDefault(reqBody.Active, curAccount.Active),
		Newrole:     sql.RoleEnum(PointerValueWithDefault(reqBody.NewRole, server.Role(curAccount.Role))),
	}

	err = a.DBClient.UpdateAccountByAdmin(ctx, params)
	if err != nil {
		return server.UpdateAccount400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Username already exists")}, nil
	}

	return server.UpdateAccount200JSONResponse(true), nil
}
