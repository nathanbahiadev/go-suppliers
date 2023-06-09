package model

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	ID       string
	Email    string `gorm:"unique"`
	Password string
	Name     string
}
