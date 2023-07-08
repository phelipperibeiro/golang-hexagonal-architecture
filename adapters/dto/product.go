package adapters_dto

import "github.com/phelipperibeiro/golang-hexagonal-architecture/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (productDto *Product) Bind(product *application.Product) (*application.Product, error) {
	if productDto.ID != "" {
		product.ID = productDto.ID
	}
	product.Name = productDto.Name
	product.Price = productDto.Price
	product.Status = productDto.Status
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
