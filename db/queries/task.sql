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

-- name: GetSortedTasksByTimeAsc :many
SELECT * FROM task ORDER BY created_at ASC;

-- name: GetSortedTasksByTimeDesc :many
SELECT * FROM task ORDER BY created_at DESC;

-- name: GetSortedTasksByStatusAsc :many
SELECT * FROM task ORDER BY status Asc;

-- name: GetSortedTasksByStatusDesc :many
SELECT * FROM task ORDER BY status DESC;
