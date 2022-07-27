package handlers

import (
	"net/http"
	"strconv"

	"github.com/dijotmathews/microservices-go/data"
	"github.com/gorilla/mux"
)

// GetProducts ...
// swagger:route GET /products products listProducts
// Returns a list of products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetProduct ...
// swagger:route GET /product/{id} product
// Returns a single product
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, perr := strconv.Atoi(vars["id"])
	if perr != nil {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)

	}
	prod := data.GetProduct(id)
	err := prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)

	}

}
