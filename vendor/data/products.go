package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product ...
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products ...
type Products []*Product

// Validate ...
func (p *Product) Validate() error {
	return nil
}

// FromJSON ...
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON ...
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}

// UpdateProduct ...
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

// ErrProductNotFound ...
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// AddProduct ...
func AddProduct(p *Product) Products {
	p.ID = getNextID()
	productList = append(productList, p)
	return productList
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// GetProducts ...
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "latte",
		Description: "frothy milky coffee",
		Price:       2.33,
		SKU:         "333sss",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "espresso",
		Description: "black coffee",
		Price:       2.33,
		SKU:         "333sss",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
