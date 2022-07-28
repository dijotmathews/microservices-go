package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product ...
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products ...
type Products []*Product

// Validate ...
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
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

// ToJSON ...
func (p *Product) ToJSON(w io.Writer) error {
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

// DeleteProduct ...
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])
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

// GetProduct ...
func GetProduct(id int) Product {
	for i := 0; i < len(productList); i++ {
		if id == productList[i].ID {
			return *productList[i]
		}
	}
	return Product{}
}

// GetProductByID ...
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// findIndexByProductID finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
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
		ID:          2,
		Name:        "espresso",
		Description: "black coffee",
		Price:       2.33,
		SKU:         "333sss",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
