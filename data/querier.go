// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package data

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (sql.Result, error)
	DeleteEmployee(ctx context.Context, employeeID int32) error
	// EMPLOYEES
	GetEmployee(ctx context.Context, employeeID int32) (Employees, error)
	ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employees, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)