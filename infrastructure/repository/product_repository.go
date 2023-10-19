package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wendryosales/desafio-01-gRPC/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (repository ProductRepositoryDb) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	error := repository.Db.Find(&products).Error

	if error != nil {
		return nil, error
	}

	return products, nil
}

func (r ProductRepositoryDb) Create(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}
