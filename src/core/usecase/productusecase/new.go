package productusecase

import "github.com/eduardo-paes/clean-go/core/domain"

type UseCase struct {
    Repository domain.ProductRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.ProductRepository) domain.ProductUseCase {
    return &UseCase{
        Repository: repository,
    }
}