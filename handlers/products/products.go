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
)

// Products is ...
type Products struct {
	l *log.Logger
}

// NewProducts is
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// KeyProduct ...
type KeyProduct struct{}
