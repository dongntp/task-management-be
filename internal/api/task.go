package api

import (
	"context"
	"github.com/google/uuid"

	"task-management-be/internal/generated/openapi/server"
	"task-management-be/internal/generated/sql"
)

func (a *API) CreateTask(ctx context.Context, request server.CreateTaskRequestObject) (server.CreateTaskResponseObject, error) {
	reqBody := *request.Body

	taskID := generateUUID()

	description, err := StringToPgtypeText(reqBody.Description)
	if err != nil {
		return server.CreateTask400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid description")}, nil
	}

	param := sql.InsertNewTaskParams{
		ID:          taskID,
		Title:       reqBody.Title,
		Description: description,
	}
	err = a.DBClient.InsertNewTask(ctx, param)
	if err != nil {
		return server.CreateTask500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
	}

	return server.CreateTask200JSONResponse(taskID), nil
}

func (a *API) AssignTask(ctx context.Context, request server.AssignTaskRequestObject) (server.AssignTaskResponseObject, error) {
	reqBody := *request.Body

	username, err := StringToPgtypeText(reqBody.Username)
	if err != nil {
		return server.AssignTask400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid username")}, nil
	}

	param := sql.AssignTaskParams{
		ID:       reqBody.TaskID,
		Username: username,
	}
	err = a.DBClient.AssignTask(ctx, param)
	if err != nil {
		return server.AssignTask500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
	}

	return server.AssignTask200JSONResponse(true), nil
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}
