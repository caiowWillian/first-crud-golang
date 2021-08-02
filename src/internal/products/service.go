package products

import (
	"github.com/caiowWillian/first-crud-golang/src/pkg/databases/mongo"
	"github.com/google/uuid"
)

type Service interface {
	CreateProduct(product Product) (string, error)
	GetAllProducts() ([]Product, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateProduct(product Product) (string, error) {
	product.Id = uuid.NewString()
	err := mongo.Repo().Insert(mongo.MongoOperation{"teste", "teste", product})

	if err != nil {
		return "", err
	}
	return product.Id, nil
}

func (s *service) GetAllProducts() ([]Product, error) {
	var product []Product
	err := mongo.Repo().GetAll(mongo.MongoOperation{"teste", "teste", nil}, &product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
