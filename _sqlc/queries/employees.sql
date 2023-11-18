-- EMPLOYEES
-- name: GetEmployee :one
SELECT * FROM employees 
WHERE employee_id = ?;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY employee_id
LIMIT ? OFFSET ?;

-- name: CreateEmployee :execresult
INSERT INTO employees (
    first_name, last_name, email, ip_address
) VALUES (
    ?, ?, ?, ?
);

-- name: UpdateEmployee :execresult
UPDATE employees
SET 
    first_name = ?,
    last_name = ?,
    email = ?,
    ip_address = ?
WHERE employee_id = ?;

-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE employee_id = ?;
