package handlers

import (
	"data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return

		}
		p.l.Println("got id ", id)
		p.updateProducts(id, rw, r)
		return
	}
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle PUT product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to un marshal json", http.StatusBadRequest)
	}
	fperr := data.UpdateProduct(id, prod)

	if fperr == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if fperr != nil {
		http.Error(rw, "Invalid URL", http.StatusInternalServerError)
		return

	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to un marshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
