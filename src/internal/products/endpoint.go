package products

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateProduct(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(Product)
		//s.CreateProduct(req)
		return createProductResponse{"oneid", http.StatusCreated}, nil
	}
}
