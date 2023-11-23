package main

import (
	"fmt"
	"log/slog"
	"net"

	pb "github.com/niewolinsky/tw_employee_data_service/grpc/employee"
	"google.golang.org/grpc"
)

func (app *application) serveGRPC(port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	employeeServiceServer := pb.NewEmployeeServiceServer(app.data_access)

	pb.RegisterEmployeeServiceServer(grpcServer, employeeServiceServer)

	slog.Info(fmt.Sprintf("Starting gRPC server on port %s", port))
	return grpcServer.Serve(listen)
}
