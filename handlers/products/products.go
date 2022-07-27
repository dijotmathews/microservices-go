// Package classification of Product API
//
// Documentation for Product API
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// 	- application/json
//
// Produces:
//  - application/json
// swagger:meta

package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dijotmathews/microservices-go/data"

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

// KeyProduct ...
type KeyProduct struct{}
