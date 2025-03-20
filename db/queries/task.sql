-- name: InsertNewTask :exec
INSERT INTO task (id, title, description)
VALUES (@id, @title, @description);

-- name: AssignTask :exec
UPDATE task
SET assignee = @username
WHERE id = @id;

-- name: UpdateTaskStatus :exec
UPDATE task
SET status = @status
WHERE id = @id;

-- name: GetAllTasks :many
SELECT * FROM task;

-- name: GetTasksByAssignee :many
SELECT * FROM task WHERE assignee = @assignee;

-- name: GetTasksByStatus :many
SELECT * FROM task WHERE status = @status;

-- name: GetSortedTasksById :many
SELECT * FROM task ORDER BY id @orderDirection;

-- name: GetSortedTasksByStatus :many
SELECT * FROM task ORDER BY status @orderDirection;
