package products

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/caiowWillian/first-crud-golang/src/pkg/encodedError"
)

type (
	Product struct {
		Id        string  `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Inventory int     `json:"inventory"`
	}

	createProductPostResponse struct {
		Id         string `json:"id"`
		StatusCode int    `json:"-"`
	}

	createProductGet struct {
		StatusCode int `json:"-"`
	}

	errResponse struct {
		StatusCode int
		err        string
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.(createProductPostResponse).StatusCode)
	return json.NewEncoder(w).Encode(response)
}

func decodeProductReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req Product
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, encodedError.BadRequest
	}
	return req, nil
}

func decodeGetProductReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
