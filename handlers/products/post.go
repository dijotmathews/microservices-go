package handlers

import (
	"net/http"

	"github.com/dijotmathews/microservices-go/data"
)

// AddProduct ...
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
