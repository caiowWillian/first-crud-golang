package products

import "github.com/caiowWillian/first-crud-golang/src/pkg/databases/mongo"

type Service interface {
	CreateProduct(product Product) bool
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateProduct(product Product) bool {
	mongo.Repo().Insert(mongo.MongoOperation{"teste", "teste", product})
	return true
}
