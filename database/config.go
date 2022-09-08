package database

import "fmt"

type Config struct {
	ServeName 	string
	User 		string
	Password 	string
	DB		string
}

var getConnectionString = func(config Config) string{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}