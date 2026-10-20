package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	DBDriver := "mysql"
	DBUser := "root"
	DBPass := ""
	DBName := "learn-go"
	db, err := sql.Open(DBDriver, DBUser+":"+DBPass+"@/"+DBName)
	return db, err
}
