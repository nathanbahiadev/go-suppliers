package repository

import "github.com/nathanbahiadev/go-suppliers/src/domain/entity"

type CompanyRepositoryInterface interface {
	Update(DB interface{}, c *entity.Company) error
}
