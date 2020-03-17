package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(60);not null" json:"password" validate:"required,min=8"`
	Name     string `gorm:"size:20;not null" json:"name"`
	Avatar   string `gorm:"size:255;not null" json:"avatar"`
	Point    int64  `gorm:"size:255;default:0;not null" json:"point"`
}
