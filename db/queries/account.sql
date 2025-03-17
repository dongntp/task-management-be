-- name: GetAllEmployers :many
SELECT username FROM account WHERE role = 'Employer';
