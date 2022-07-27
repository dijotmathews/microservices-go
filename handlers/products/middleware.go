package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dijotmathews/microservices-go/data"
)

// MiddlewareProductValidation validates the product in request and calles the next if fine.
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validating the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(rw, fmt.Sprintf("Error validating product %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		//calling the next handler or the final handler.
		next.ServeHTTP(rw, r)
	})
}
