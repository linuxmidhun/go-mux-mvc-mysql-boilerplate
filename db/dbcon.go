package db

import (
	"database/sql"
)

// Conn .
func Conn() (con *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "[Password]"
	dbName := "sample"
	con, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return con
}
