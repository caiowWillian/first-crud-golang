package products

import (
	"context"
	"errors"
	"net/http"

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
