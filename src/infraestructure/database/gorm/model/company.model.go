package model

import "gorm.io/gorm"

type CompanyModel struct {
	gorm.Model

	ID        string
	CNPJ      string
	Name      string
	UserID    string
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
