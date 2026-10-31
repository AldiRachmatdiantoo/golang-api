package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_golang_api")
	if err != nil {
		panic("Tidak dapat tersambung ke database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Gagal menjalankan koneksi! " + err.Error())
	}
	return db
}