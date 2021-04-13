package repository

import (
	_ "github.com/lib/pq"
	"github.com/matheusmhmelo/go-jobs/internal/config"
	"github.com/matheusmhmelo/go-jobs/pkg/db"
	"os"
)

func Start(){
	postgres := config.Postgres{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Ssl:	  os.Getenv("DB_SSL"),
	}

	Db = &db.PostgreSQL{}
	Db.Connect(postgres)
}

var Db db.Database