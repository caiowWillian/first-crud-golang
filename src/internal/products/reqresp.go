package products

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	Product struct {
		Id        string  `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Inventory int     `json:"inventory"`
	}

	createProductResponse struct {
		Id         string `json:"id"`
		statusCode int
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
