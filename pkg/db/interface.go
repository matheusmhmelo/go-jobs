package db

import (
	"database/sql"
)

type Database interface {
	Connect(Config)
	Execute(string, ...interface{}) (sql.Result, error)
	Get(string, ...interface{}) (map[string]interface{}, error)
	Fetch(string, ...interface{}) ([]map[string]interface{}, error)
}


type Config interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabase() string
	GetSSL() string
}