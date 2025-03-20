-- name: GetAllEmployers :many
SELECT username FROM account WHERE role = 'Employer';

-- name: GetUserByUserName :one
SELECT role, password, active FROM account WHERE username = @username;

-- name: InsertNewAccount :exec
INSERT INTO account (username, password, role)
VALUES (@username, @password, @role);

-- name: UpdateAccountByAdmin :exec
UPDATE account
SET
  username = @newUsername,
  password = @newPassword,
  active = @newActive,
  role = @newRole
WHERE username = @username;

-- name: GetEmployeeSummary :many
SELECT account.username, coalesce(t.total_tasks, 0), coalesce(t.total_completed, 0)
FROM account
LEFT JOIN (
  SELECT assignee, COUNT(*) AS total_tasks, COUNT(
    CASE WHEN status = 'Completed' THEN 1 ELSE NULL END
  ) as total_completed
  FROM task
  GROUP BY assignee
) AS t ON t.assignee = account.username
WHERE role = 'Employee';
