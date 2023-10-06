package domain

import (
	"net/http"

	"github.com/eduardo-paes/clean-go/core/dto"
)

// It's an entity of table product database column
type Product struct {
	ID			int32 `json:"id"`
	Name 		string `json:"name"`
	Price 		float32 `json:"price"`
	Description string `json:"description"`
}

// It's a contract of http adapter layer
type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// It's a contract of business rule layer
type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

// It's a contract of database connection adapter layer
type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}