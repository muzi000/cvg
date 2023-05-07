package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "rootwolf", "127.0.0.1", 3306, "vuln")
	fmt.Println(connStr)
	var err error
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	return DB.Ping()
}
