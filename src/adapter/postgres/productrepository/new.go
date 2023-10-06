package productrepository

import (
	"github.com/eduardo-paes/clean-go/adapter/postgres"
	"github.com/eduardo-paes/clean-go/core/domain"
)

type Repository struct {
    db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
    return &Repository{
        db: db,
    }
}