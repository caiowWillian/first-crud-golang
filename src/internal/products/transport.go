package products

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, s Service, router *mux.Router) http.Handler {
	router.Use(commonMiddleware)

	router.Methods(http.MethodPost).Path("/product").Handler(httptransport.NewServer(
		makeCreateProduct(s),
		decodeProductReq,
		encodeResponse,
	))

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
