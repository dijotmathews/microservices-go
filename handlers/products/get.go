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
// swagger:route GET /{id} products users returnProduct
//
// Returns product filtered by ID.
//
// This will show available product by default.
//
//
//     Consumes:
//     - application/json
//
//
//     Produces:
//     - application/json
//
//
//     Schemes: http
//
//
//     Parameters:
//       + name: limit
//         in: query
//         description: maximum numnber of results to return
//         required: false
//         type: integer
//         format: int32
//
//
//     Responses:
//       default: genericError
//       200: someResponse
//       422: validationError
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
