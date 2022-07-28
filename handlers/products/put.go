package handlers

import (
	"net/http"
	"strconv"

	"github.com/dijotmathews/microservices-go/data"
	"github.com/gorilla/mux"
)

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
