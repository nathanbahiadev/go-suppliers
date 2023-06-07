package repository

import "github.com/nathanbahiadev/go-suppliers/src/domain/entity"

type SupplierRepositoryInterface interface {
	Create(DB interface{}, s *entity.Supplier) error
	Update(DB interface{}, s *entity.Supplier) error
	Find(DB interface{}, id string) (*entity.Supplier, error)
	FindAll(DB interface{}, id string) ([]*entity.Supplier, error)
}
