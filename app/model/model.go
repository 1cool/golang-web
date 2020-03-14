package model

import "github.com/jinzhu/gorm"

func SetUp() {
	db, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
}
