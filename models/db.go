package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/woodylan/go-websocket/config"
	"log"
)

var Db *sql.DB
var Db2 *sql.DB

func init() {
	log.Println("config: DB")
	var err error
	driver := ViperConfig.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&&parseTime=true", ViperConfig.Db.User, ViperConfig.Db.Password,
		ViperConfig.Db.Address, ViperConfig.Db.Database)
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}
