package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Server config
	RESTPort        string `envconfig:"REST_PORT" default:"8000"`
	GracefulTimeout int    `envconfig:"GRACEFUL_TIMEOUT" default:"5"`
	WriteTimeout    int    `envconfig:"WRITE_TIMEOUT" default:"15"`
	ReadTimeout     int    `envconfig:"READ_TIMEOUT" default:"15"`
	HostAdress      string `envconfig:"HOST_ADRESS" DEFAULT:"127.0.0.1"`

	// MySQL Database config
	MySQLPort         string `envconfig:"MYSQL_PORT" default:"3306"`
	MySQLHost         string `envconfig:"MYSQL_HOST" default:"127.0.0.1"`
	MySQLUser         string `envconfig:"MYSQL_USER" default:"user"`
	MySQLPassword     string `envconfig:"MYSQL_PASSWORD" default:"password"`
	MySQLDatabaseName string `envconfig:"MYSQL_DATABASE_NAME" default:"bareksa_news"`
	MySQLDSNFormat    string `envconfig:"MYSQL_DSN_FORMAT" default:"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
