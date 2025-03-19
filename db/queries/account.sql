-- name: GetAllEmployers :many
SELECT username FROM account WHERE role = 'Employer';

-- name: GetRoleByUser :one
SELECT role FROM account WHERE username = @username AND password = @password;

-- name: InsertNewAccount :exec
INSERT INTO account (username, password, role)
VALUES (@username, @password, @role);
