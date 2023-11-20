package main

import (
	"github.com/niewolinsky/tw_employee_data_service/config"
	data "github.com/niewolinsky/tw_employee_data_service/data"

	"log/slog"
	"sync"

	"github.com/go-playground/validator/v10"
)

type application struct {
	data_access *data.Queries
	validator   *validator.Validate
	wait_group  sync.WaitGroup
}

func main() {
	mysqlClient, restApiPort, grpcApiPort := config.InitConfig()
	defer mysqlClient.Close()

	app := &application{
		data_access: data.New(mysqlClient),
		validator:   validator.New(),
	}

	// ? start HTTP server in a goroutine to serve both HTTP and GRPC
	go func() {
		err := app.serveREST(restApiPort)
		if err != nil {
			slog.Error("failed starting HTTP server", err)
		}
	}()

	err := app.serveGRPC(grpcApiPort)
	if err != nil {
		slog.Error("failed starting gRPC server", err)
	}

	slog.Info("stopped server")
}
