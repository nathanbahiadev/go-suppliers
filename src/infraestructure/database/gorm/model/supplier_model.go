package model

import "gorm.io/gorm"

type SupplierModel struct {
	gorm.Model

	ID        string
	Name      string
	CompanyID string
	Company   CompanyModel `gorm:"foreignKey:CompanyID"`
}
