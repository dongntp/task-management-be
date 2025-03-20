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

func (a *API) GetTasksByEmployer(ctx context.Context, request server.GetTasksByEmployerRequestObject) (server.GetTasksByEmployerResponseObject, error) {
	params := request.Params

	var (
		tasks []sql.Task
		err   error
	)

	switch {
	case params.Assignee != nil:
		assignee, e := StringToPgtypeText(*params.Assignee)
		if e != nil {
			return server.GetTasksByEmployer400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("Invalid assignee")}, nil
		}
		tasks, err = a.DBClient.GetTasksByAssignee(ctx, assignee)
	case params.Status != nil:
		tasks, err = a.DBClient.GetTasksByStatus(ctx, sql.Status(*params.Status))
	case params.OrderBy != nil:
		if params.OrderDirection == nil {
			return server.GetTasksByEmployer400JSONResponse{ClientErrorJSONResponse: server.ClientErrorJSONResponse("`orderDirection` must be specified with `orderBy`")}, nil
		}
		if *params.OrderBy == server.OrderByCreationDate {
			if *params.OrderDirection == server.ASC {
				tasks, err = a.DBClient.GetSortedTasksByTimeAsc(ctx)
			} else {
				tasks, err = a.DBClient.GetSortedTasksByTimeDesc(ctx)
			}
		} else {
			if *params.OrderDirection == server.ASC {
				tasks, err = a.DBClient.GetSortedTasksByStatusAsc(ctx)
			} else {
				tasks, err = a.DBClient.GetSortedTasksByStatusDesc(ctx)
			}
		}
	default:
		tasks, err = a.DBClient.GetAllTasks(ctx)
	}

	if err != nil {
		return server.GetTasksByEmployer500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
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

	return server.GetTasksByEmployer200JSONResponse(res), nil
}

func (a *API) GetEmployeeSummary(ctx context.Context, request server.GetEmployeeSummaryRequestObject) (server.GetEmployeeSummaryResponseObject, error) {
	employees, err := a.DBClient.GetEmployeeSummary(ctx)
	if err != nil {
		return server.GetEmployeeSummary500JSONResponse{ServerErrorJSONResponse: server.ServerErrorJSONResponse("Internal server error")}, nil
	}

	res := []server.EmployeeSummary{}
	for _, e := range employees {
		res = append(res, server.EmployeeSummary{
			Employee:       e.Username,
			TotalTasks:     e.TotalTasks,
			TotalCompleted: e.TotalCompleted,
		})
	}

	return server.GetEmployeeSummary200JSONResponse(res), nil
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}
