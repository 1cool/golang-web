package model

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	"golang-web/config"
)

var Db *gorm.DB

func SetUp() {
	var AppConfig config.App
	_ = envconfig.Process("myapp", &AppConfig)
	var buffer bytes.Buffer
	buffer.WriteString(AppConfig.DbUsername)
	buffer.WriteString(":")
	buffer.WriteString(AppConfig.DbPassword)
	buffer.WriteString("@tcp(")
	buffer.WriteString(AppConfig.DbHost)
	buffer.WriteString(":")
	buffer.WriteString(AppConfig.DbPort)
	buffer.WriteString(")/")
	buffer.WriteString(AppConfig.DbDatabase)
	buffer.WriteString("?charset=")
	buffer.WriteString(AppConfig.DbCharset)
	buffer.WriteString("&parseTime=True&loc=Local")
	dbConn := buffer.String()
	fmt.Println(AppConfig.DbConnection, dbConn)
	Db, _ = gorm.Open(AppConfig.DbConnection, dbConn)
	//defer db.Close()
	//Db.AutoMigrate(&user.User{})
}
