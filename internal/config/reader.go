package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}
	viper.AutomaticEnv()

	return &Config{
		TZ: loadString("TZ"),
		Database: Database{
			MySql: MySql{
				Host:     loadString("DATABASE_MYSQL_HOST"),
				Port:     loadInt("DATABASE_MYSQL_PORT"),
				Username: loadString("DATABASE_MYSQL_USERNAME"),
				Password: loadString("DATABASE_MYSQL_PASSWORD"),
				DbName:   loadString("DATABASE_MYSQL_DBNAME"),
			},
		},
	}, nil
}
