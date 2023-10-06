package productservice

import "github.com/eduardo-paes/clean-go/core/domain"

type Service struct {
	UseCase domain.ProductUseCase
}

func New(usecase domain.ProductUseCase) domain.ProductService {
	return &Service {
		UseCase: usecase,
	}
}