-- name: InsertNewTask :exec
INSERT INTO task (id, title, description)
VALUES (@id, @title, @description);

-- name: AssignTask :exec
UPDATE task
SET assignee = @username
WHERE id = @id;
