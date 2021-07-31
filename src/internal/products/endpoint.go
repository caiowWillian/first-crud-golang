package products

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateProduct(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(Product)
		//s.CreateProduct(req)

		return CreateProductResponse{"oneid"}, nil
	}
}
