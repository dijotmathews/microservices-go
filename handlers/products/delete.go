package handlers

import (
	"net/http"
	"strconv"

	"github.com/dijotmathews/microservices-go/data"
	"github.com/gorilla/mux"
)

// DeleteProduct ...
func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, perr := strconv.Atoi(vars["id"])
	if perr != nil {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)

	}

	p.l.Println("handle DELETE product")
	// prod := r.Context().Value(KeyProduct{}).(data.Product)

	fperr := data.DeleteProduct(id)

	if fperr == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		// data.ToJSON(&GenericError{Message: fperr.Error()}, rw)
		return
	}

	if fperr != nil {
		p.l.Println("[ERROR] deleting record", fperr)

		rw.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Message: fperr.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}
