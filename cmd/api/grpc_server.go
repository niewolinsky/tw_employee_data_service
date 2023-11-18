package main

import (
	"log"
	"net"

	employeeService "github.com/niewolinsky/tw_employee_data_service/internal/grpc/employee"
	"google.golang.org/grpc"
)

func startGrpcServer() {
	lis, err := net.Listen("tcp", ":50051") // Replace with your desired port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	employeeServiceServer := employeeService.NewEmployeeServiceServer()

	employeeService.RegisterEmployeeServiceServer(grpcServer, employeeServiceServer)

	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
