package api

import (
	"context"

	"task-management-be/internal/generated/openapi/server"
	"task-management-be/internal/generated/sql"
	"task-management-be/internal/pkg/middleware"
)

func (a *API) UpdateTaskStatus(ctx context.Context, request server.UpdateTaskStatusRequestObject) (server.UpdateTaskStatusResponseObject, error) {
	reqBody := *request.Body

	username, err := middleware.GetUserName(ctx)
	if err != nil {
		return server.UpdateTaskStatus400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Unauthorized")}, nil
	}
	convertedUsername, err := StringToPgtypeText(username)
	if err != nil {
		return server.UpdateTaskStatus400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid username")}, nil
	}
	isValidTask, err := a.DBClient.CheckValidTask(ctx, sql.CheckValidTaskParams{
		ID:       reqBody.TaskID,
		Username: convertedUsername,
	})
	if !isValidTask {
		return server.UpdateTaskStatus400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid task")}, nil
	}

	param := sql.UpdateTaskStatusParams{
		ID:     reqBody.TaskID,
		Status: sql.Status(reqBody.Status),
	}
	err = a.DBClient.UpdateTaskStatus(ctx, param)
	if err != nil {
		return server.UpdateTaskStatus500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
	}

	return server.UpdateTaskStatus200JSONResponse(true), nil
}

func (a *API) GetTasksByEmployee(ctx context.Context, request server.GetTasksByEmployeeRequestObject) (server.GetTasksByEmployeeResponseObject, error) {
	username, err := middleware.GetUserName(ctx)
	if err != nil {
		return server.GetTasksByEmployee400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Unauthorized")}, nil
	}

	convertedUsername, err := StringToPgtypeText(username)
	if err != nil {
		return server.GetTasksByEmployee400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid username")}, nil
	}

	tasks, err := a.DBClient.GetTasksByAssignee(ctx, convertedUsername)
	if err != nil {
		return server.GetTasksByEmployee500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
	}

	res := []server.Task{}
	for _, task := range tasks {
		var assignee *string
		if task.Assignee.Valid {
			assignee = &task.Assignee.String
		}

		res = append(res, server.Task{
			Id:           task.ID,
			Title:        task.Title,
			Description:  task.Description.String,
			Status:       server.Status(task.Status),
			Assignee:     assignee,
			CreationDate: task.CreatedAt.Time,
		})
	}

	return server.GetTasksByEmployee200JSONResponse(res), nil
}
