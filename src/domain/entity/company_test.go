package entity_test

import (
	"testing"

	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidateCompanyWhenIDIsEmptyShouldRaiseAnError(t *testing.T) {
	Company := entity.Company{}
	err := Company.Validate()
	assert.Equal(t, err.Error(), "company id cannot be blank")
}

func TestValidateCompanyWhenUserIDIsEmptyShouldRaiseAnError(t *testing.T) {
	Company := entity.Company{ID: "1"}
	err := Company.Validate()
	assert.Equal(t, err.Error(), "company user id cannot be blank")
}

func TestValidateCompanyWhenCNPJIsIEmptyShouldRaiseAnError(t *testing.T) {
	Company := entity.Company{
		ID:   "1",
		User: entity.User{ID: "1"},
	}
	err := Company.Validate()
	assert.Equal(t, err.Error(), "cnpj is invalid")
}

func TestValidateCompanyWhenCNPJIsInvalidShouldRaiseAnError(t *testing.T) {
	company := entity.Company{
		ID:   "1",
		User: entity.User{ID: "1"},
		CNPJ: entity.CNPJ{Value: "00.000.000/0001-999"},
	}

	err := company.Validate()
	assert.Equal(t, err.Error(), "cnpj is invalid")
}

func TestValidateCompanyWhenNameIsEmptyShouldRaiseAnError(t *testing.T) {
	company := entity.Company{
		ID:   "1",
		User: entity.User{ID: "1"},
		CNPJ: entity.CNPJ{Value: "00.000.000/0001-00"},
	}

	err := company.Validate()
	assert.Equal(t, err.Error(), "company name cannot be blank")
}

func TestCreateCompany(t *testing.T) {
	company := entity.Company{
		ID:   "1",
		User: entity.User{ID: "1"},
		CNPJ: entity.CNPJ{Value: "00.000.000/0001-00"},
		Name: "Company",
	}
	err := company.Create()

	assert.Nil(t, err)
	assert.NotEqual(t, company.ID, "")
}
