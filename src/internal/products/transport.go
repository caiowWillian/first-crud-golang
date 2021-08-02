package products

import (
	"context"
	"net/http"

	"github.com/caiowWillian/first-crud-golang/src/pkg/encodedError"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, s Service, router *mux.Router) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodedError.EncodeError),
	}

	router.Use(commonMiddleware)

	router.Methods(http.MethodPost).Path("/product").Handler(httptransport.NewServer(
		makeCreateProduct(s),
		decodeProductReq,
		encodeResponse,
		opts...,
	))

	router.Methods(http.MethodGet).Path("/product").Handler(httptransport.NewServer(
		makeGetAllProducts(s),
		decodeGetProductReq,
		httptransport.EncodeJSONResponse,
		opts...,
	))

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
