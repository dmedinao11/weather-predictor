package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "db_mysql:3306",
		DBName: "weather_database",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}
