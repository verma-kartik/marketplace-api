package models

import (
	"context"
)

type Product struct {
	Name         string  `json:"name"`
	SerialNumber int32   `json:"serialNumber,primary_key"`
	Quantity     int32   `json:"quantity"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
}

// ProductRepository this interface defines all the methods our
// services needs in order to operate
type ProductRepository interface {
	GetProductBySerialN(context.Context, int32) (Product, error)
	CreateProduct(context.Context, *Product) error
}
