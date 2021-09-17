package util

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user   = "fox"
	pass   = "Transformers"
	ip     = "127.0.0.1"
	port   = "3306"
	dbName = "challenge"
)

// CONEXION DB
func InitDB() *sql.DB {

	path := strings.Join([]string{user, ":", pass, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	// maximo conexiones
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(5)

	// Verificacion de connect
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
