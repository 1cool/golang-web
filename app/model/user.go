package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	NickName string `gorm:"size:20" json:"nick_name"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Point    int64  `gorm:"size:255;default:0" json:"point"`
}
