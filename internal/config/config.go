package config

import "time"

type Config struct {
	TZ          string
	LogLevel    string
	AppDebug    bool
	Database    Database
	HttpApi     HttpApi
	HealthToken string
	UserToken   UserToken
}

type Database struct {
	MySql MySql
}

type MySql struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type HttpApi struct {
	Host string
	Port int
}

type UserToken struct {
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
	PrivateKey         string
	PublicKey          string
	PrivateKeyFilePath string
	PublicKeyFilePath  string
}
