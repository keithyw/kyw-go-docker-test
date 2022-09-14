package conf

import (
	"os"
)

type Config struct {
	MysqlUser string
	MysqlPass string
	MysqlHost string
	MysqlDBName string
	Port string
	GrpcHost string
}

func NewConfig() (*Config, error) {
	return &Config{
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
		os.Getenv("PORT"),
		os.Getenv("GRPC_HOST"),
	}, nil
}