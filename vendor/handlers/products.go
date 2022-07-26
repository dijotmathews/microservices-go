package handlers

import (
	"context"
	"data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is ...
type Products struct {
	l *log.Logger
}

// NewProducts is
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// UpdateProducts ...
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, perr := strconv.Atoi(vars["id"])
	if perr != nil {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)

	}

	p.l.Println("handle PUT product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	fperr := data.UpdateProduct(id, &prod)

	if fperr == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if fperr != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return

	}

}

// AddProduct ...
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

// GetProducts ...
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// KeyProduct ...
type KeyProduct struct{}

//MiddlewareProductValidation ...
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
