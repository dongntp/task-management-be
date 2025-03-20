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
