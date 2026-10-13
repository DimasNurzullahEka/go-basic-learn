package config

import "database/sql"

func DBConnection() (*sql.DB, error) {
	DBDriver := "mysql"
	DBUser := "root"
	DBPass := ""
	DBName := "learn_go"
	db, err := sql.Open(DBDriver, DBUser+":"+DBPass+"@/"+DBName)
	return db, err
}
