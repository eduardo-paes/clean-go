package productusecase

import (
	"github.com/eduardo-paes/clean-go/core/domain"
	"github.com/eduardo-paes/clean-go/core/dto"
)

func (usecase UseCase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.Repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}