package products

import (
	"context"
	"errors"
	"net/http"

	"github.com/caiowWillian/first-crud-golang/src/pkg/encodedError"
	"github.com/go-kit/kit/endpoint"
)

func makeCreateProduct(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Product)
		id, err := s.CreateProduct(req)

		if err != nil {
			return nil, errors.New("Internal Server Error")
		}

		return createProductPostResponse{id, http.StatusCreated}, nil
	}
}

func makeGetAllProducts(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetAllProducts()

		if err != nil {
			return nil, encodedError.InternalServerError
		}

		if products == nil {
			return []Product{}, nil
		}
		return products, nil
	}
}
