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
		TZ:       loadString("TZ"),
		LogLevel: loadString("LOG_LEVEL"),
		AppDebug: loadBool("APP_DEBUG"),
		Database: Database{
			MySql: MySql{
				Host:     loadString("DATABASE_MYSQL_HOST"),
				Port:     loadInt("DATABASE_MYSQL_PORT"),
				Username: loadString("DATABASE_MYSQL_USERNAME"),
				Password: loadString("DATABASE_MYSQL_PASSWORD"),
				DbName:   loadString("DATABASE_MYSQL_DBNAME"),
			},
		},
		HttpApi: HttpApi{
			Host: loadString("HTTP_API_HOST"),
			Port: loadInt("HTTP_API_PORT"),
		},
		HealthToken: loadString("HEALTH_TOKEN"),
		UserToken: UserToken{
			AccessTokenTTL:     loadDuration("ACCESS_TOKEN_TTL"),
			RefreshTokenTTL:    loadDuration("REFRESH_TOKEN_TTL"),
			PrivateKey:         loadString("TOKEN_PRIVATE_KEY"),
			PublicKey:          loadString("TOKEN_PUBLIC_KEY"),
			PrivateKeyFilePath: loadString("TOKEN_PRIVATE_KEY_FILE_PATH"),
			PublicKeyFilePath:  loadString("TOKEN_PUBLIC_KEY_KEY_FILE_PATH"),
		},
	}, nil
}
