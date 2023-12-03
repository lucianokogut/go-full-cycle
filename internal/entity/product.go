package entity

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NeProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

// Contrato de acesso aos dados via Repository com mapeamento da entidade utilizada
type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

// Definição de uma Interface para acesso ao Repository
