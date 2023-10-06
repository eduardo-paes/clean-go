package productusecase

import (
	"github.com/eduardo-paes/clean-go/core/domain"
	"github.com/eduardo-paes/clean-go/core/dto"
)

func (usecase UseCase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.Repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}