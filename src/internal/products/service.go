package products

import (
	"github.com/caiowWillian/first-crud-golang/src/pkg/databases/mongo"
	"github.com/google/uuid"
)

type Service interface {
	CreateProduct(product Product) (string, error)
	GetAllProducts() ([]Product, error)
}

type service struct {
	repo mongo.Repository
}

func NewService(repo mongo.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(product Product) (string, error) {
	product.Id = uuid.NewString()
	err := s.repo.Insert(mongo.MongoOperation{"teste", "teste", product})

	if err != nil {
		return "", err
	}
	return product.Id, nil
}

func (s *service) GetAllProducts() ([]Product, error) {
	var product []Product
	err := s.repo.GetAll(mongo.MongoOperation{"teste", "teste", nil}, &product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
