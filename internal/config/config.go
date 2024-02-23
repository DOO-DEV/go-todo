package config

type Config struct {
	TZ          string
	LogLevel    string
	AppDebug    bool
	Database    Database
	HttpApi     HttpApi
	HealthToken string
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
