package employee

import (
	"context"
	"database/sql"

	data "github.com/niewolinsky/tw_employee_data_service/data"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ EmployeeServiceServer = (*employeeServiceServer)(nil)

type employeeServiceServer struct {
	UnimplementedEmployeeServiceServer
	dataAccess *data.Queries
}

func NewEmployeeServiceServer(dataAccess *data.Queries) EmployeeServiceServer {
	return &employeeServiceServer{
		dataAccess: dataAccess,
	}
}

func (s *employeeServiceServer) GetEmployee(ctx context.Context, req *GetEmployeeRequest) (*EmployeeResponse, error) {
	employee, err := s.dataAccess.GetEmployee(ctx, req.GetEmployeeId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "employee not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get employee: %v", err)
	}

	return &EmployeeResponse{
		EmployeeId: employee.EmployeeID,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		Email:      employee.Email,
		IpAddress:  employee.IpAddress,
	}, nil
}

func (s *employeeServiceServer) ListEmployees(ctx context.Context, req *ListEmployeesRequest) (*ListEmployeesResponse, error) {
	employees, err := s.dataAccess.ListEmployees(ctx, data.ListEmployeesParams{Limit: req.GetLimit(), Offset: req.GetOffset()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list employees: %v", err)
	}

	var employeeResponses []*EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, &EmployeeResponse{
			EmployeeId: employee.EmployeeID,
			FirstName:  employee.FirstName,
			LastName:   employee.LastName,
			Email:      employee.Email,
			IpAddress:  employee.IpAddress,
		})
	}

	return &ListEmployeesResponse{
		Employees: employeeResponses,
	}, nil
}
