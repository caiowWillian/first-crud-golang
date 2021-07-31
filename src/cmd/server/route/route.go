package route

import (
	"context"

	"github.com/caiowWillian/first-crud-golang/src/internal/products"
	"github.com/gorilla/mux"
)

func MakeRoutes(ctx context.Context, router *mux.Router) {
	products.NewHTTPServer(ctx, products.NewService(), router)
}
