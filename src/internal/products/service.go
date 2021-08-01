package products

import (
	"fmt"

	"github.com/caiowWillian/first-crud-golang/src/pkg/databases/mongo"
)

type Service interface {
	CreateProduct(product Product) bool
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateProduct(product Product) bool {
	err := mongo.Repo().Insert(mongo.MongoOperation{"teste", "teste", product})
	fmt.Println(err)
	return true
}
