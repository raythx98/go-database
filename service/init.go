package service

import (
	"database/sql"
	"github.com/raythx98/go-database/service/receiver"
	database "github.com/raythx98/go-database/sqlc/output"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/go_url_shortener")
	if err != nil {
		return nil, err
	}

	receiver.DbInstance.Queries = database.New(db)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
