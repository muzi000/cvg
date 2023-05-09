package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muzi000/vuln-go/config"
	"github.com/muzi000/vuln-go/logging"
)

var DB *sql.DB

func InitDB() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DatabaseSetting.UserName, config.DatabaseSetting.Password, config.DatabaseSetting.Host, config.DatabaseSetting.Port, config.DatabaseSetting.DBName)
	//fmt.Println(connStr)
	var err error
	DB, err = sql.Open(config.DatabaseSetting.DBType, connStr)
	if err != nil {
		logging.LogPrint("err", err.Error())
		return err
	}
	return DB.Ping()
}
