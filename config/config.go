package config

import (
	"context"
	"database/sql"
	"flag"
	"os"

	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type configuration struct {
	version           string
	rest_api_portport string
	grpc_api_port     string
	env               string
	db                struct {
		dsn string
	}
}

func initializeMySQLClient(cfg configuration) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Check if the database is accessible
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitConfig() (*sql.DB, string, string) {
	config := configuration{}

	err := godotenv.Load()
	if err != nil {
		slog.Error("failed loading environment variables", err)
	}
	slog.Info("environment variables loaded")

	//?APP
	flag.StringVar(&config.rest_api_portport, "rest_api_port", os.Getenv("REST_API_PORT"), "application server port")
	flag.StringVar(&config.grpc_api_port, "grpc_api_port", os.Getenv("GRPC_API_PORT"), "application server port")
	flag.StringVar(&config.version, "version", os.Getenv("APP_VERSION"), "application version")
	flag.StringVar(&config.env, "env", os.Getenv("APP_ENVIRONMENT"), "application environment")

	//?MYSQL
	flag.StringVar(&config.db.dsn, "db-dsn", os.Getenv("MYSQL_DSN"), "mysql dsn")

	flag.Parse()
	slog.Info("command line variables loaded")

	mysqlClient, err := initializeMySQLClient(config)
	if err != nil {
		slog.Error("failed initializing MySQL client", err)
	}
	slog.Info("MySQL client initialized")

	return mysqlClient, config.rest_api_portport, config.grpc_api_port
}
