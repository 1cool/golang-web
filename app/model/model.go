package model

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"golang-web/config"
)

func SetUp() {
	var buffer bytes.Buffer
	buffer.WriteString(config.AppConfig.DbUsername)
	buffer.WriteString(":")
	buffer.WriteString(config.AppConfig.DbPassword)
	buffer.WriteString("@")
	buffer.WriteString(config.AppConfig.DbHost)
	buffer.WriteString("/")
	buffer.WriteString(config.AppConfig.DbDatabase)
	buffer.WriteString("?charset=")
	buffer.WriteString(config.AppConfig.DbCharset)
	buffer.WriteString("&parseTime=True&loc=Local")
	dbConn := buffer.String()

	db, _ := gorm.Open(config.AppConfig.DbConnection, dbConn)
	defer db.Close()
}
