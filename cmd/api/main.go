package main

import (
	config "github.com/niewolinsky/tw_employee_data_service/config"
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
	postgres_client, app_port := config.InitConfig()
	defer postgres_client.Close()

	app := &application{
		data_access: data.New(postgres_client),
		validator:   validator.New(),
	}

	err := app.serve(app_port)
	if err != nil {
		slog.Error("failed starting server", err)
	}
	slog.Info("stopped server")
}
