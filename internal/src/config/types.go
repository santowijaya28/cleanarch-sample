package config

type Config struct {
	DatabaseConfig DatabaseType
}

type DatabaseType struct {
	Host       string
	Port       string
	DbName     string
	DbUserName string
	DbPwd      string
}
