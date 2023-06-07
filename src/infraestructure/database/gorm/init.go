package database

import (
	"fmt"

	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func StartDatabase() {
	fmt.Println("************** Iniciando banco de dados **************")

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("************** Banco de dados conectado **************")
	fmt.Println("**************** Realizando migrações ****************")

	DB.AutoMigrate(&model.UserModel{})
	DB.AutoMigrate(&model.CompanyModel{})
	DB.AutoMigrate(&model.SupplierModel{})

	fmt.Println("*************** Migrações finalizadas ****************")
	fmt.Println("************** Banco de dados pronto! ****************")
}
