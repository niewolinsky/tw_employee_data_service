// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: employees.sql

package data

import (
	"context"
	"database/sql"
)

const createEmployee = `-- name: CreateEmployee :execresult
INSERT INTO employees (
    first_name, last_name, email, ip_address
) VALUES (
    ?, ?, ?, ?
)
`

type CreateEmployeeParams struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Email     string `json:"email" validate:"required,email"`
	IpAddress string `json:"ip_address" validate:"required,ipv4"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEmployee,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.IpAddress,
	)
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE employee_id = ?
`

func (q *Queries) DeleteEmployee(ctx context.Context, employeeID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, employeeID)
	return err
}

const getEmployee = `-- name: GetEmployee :one
SELECT employee_id, first_name, last_name, email, ip_address FROM employees 
WHERE employee_id = ?
`

// EMPLOYEES
func (q *Queries) GetEmployee(ctx context.Context, employeeID int32) (Employees, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, employeeID)
	var i Employees
	err := row.Scan(
		&i.EmployeeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IpAddress,
	)
	return i, err
}

const listEmployees = `-- name: ListEmployees :many
SELECT employee_id, first_name, last_name, email, ip_address FROM employees
ORDER BY employee_id
LIMIT ? OFFSET ?
`

type ListEmployeesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employees, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employees
	for rows.Next() {
		var i Employees
		if err := rows.Scan(
			&i.EmployeeID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.IpAddress,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmployee = `-- name: UpdateEmployee :execresult
UPDATE employees
SET 
    first_name = ?,
    last_name = ?,
    email = ?,
    ip_address = ?
WHERE employee_id = ?
`

type UpdateEmployeeParams struct {
	FirstName  string `json:"first_name" validate:"required,alpha"`
	LastName   string `json:"last_name" validate:"required,alpha"`
	Email      string `json:"email" validate:"required,email"`
	IpAddress  string `json:"ip_address" validate:"required,ipv4"`
	EmployeeID int32  `json:"employee_id"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateEmployee,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.IpAddress,
		arg.EmployeeID,
	)
}
