package config

type Config struct {
	TZ       string
	Database Database
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
