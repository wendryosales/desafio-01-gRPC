package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface {
	FindAll() ([]*Product, error)
	Create(product *Product) (*Product, error)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	Id          int32   `json:"id"          gorm:"primary_key"       valid:"int,optional"`
	Name        string  `json:"name"        gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"-"`
	Price       float32 `json:"price"       gorm:"type:float"        valid:"notnull"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if p.Price <= 0 {
		return errors.New("The price must be greater than zero")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
