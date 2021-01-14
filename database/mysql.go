package database

import (
	"database/sql"
	"log"
	. "main/config"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", Conf.Database.Mysql.Url)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("数据库连接成功...")
}
