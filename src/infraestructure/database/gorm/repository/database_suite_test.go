package repository_test

import (
	"fmt"
	"testing"

	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/model"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseSuiteTest struct {
	suite.Suite
	DB *gorm.DB
}

func (suite *DatabaseSuiteTest) SetupTest() {
	DB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	DB.AutoMigrate(
		&model.UserModel{},
		&model.CompanyModel{},
		&model.SupplierModel{},
	)

	suite.DB = DB
}

func (suite *DatabaseSuiteTest) TearDownTest() {
	suite.DB.Unscoped().Delete(&model.UserModel{}, nil)
	suite.DB.Unscoped().Delete(&model.CompanyModel{}, nil)
	suite.DB.Unscoped().Delete(&model.SupplierModel{}, nil)
}

func TestExecuteDatabaseSuiteTest(t *testing.T) {
	suite.Run(t, new(DatabaseSuiteTest))
}
