package web

import (
	"encoding/json"
	"net/http"

	"br.com.lucianokogut/go-full-cycle-esquenta/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
