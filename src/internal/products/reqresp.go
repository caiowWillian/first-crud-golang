package products

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	Product struct {
		id        string  `json:"id"`
		name      string  `json:"name"`
		price     float64 `json:"price"`
		inventory int     `json:"inventory"`
	}

	CreateProductResponse struct {
		id string `json:"id"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeProductReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req Product
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
