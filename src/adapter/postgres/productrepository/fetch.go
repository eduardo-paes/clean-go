package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/eduardo-paes/clean-go/core/domain"
	"github.com/eduardo-paes/clean-go/core/dto"
)

func (repository Repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
    ctx := context.Background()
    products := []domain.Product{}
    total := int32(0)

    pagin := paginate.Instance(domain.Product{})

    query, queryCount := pagin.Query("SELECT * FROM product").
        Sort(pagination.Sort).
        Desc(pagination.Descending).
        Page(pagination.Page).
        RowsPerPage(pagination.ItemsPerPage).
        SearchBy(pagination.Search, "name", "description").
        Select()

    {
        rows, err := repository.db.Query(
            ctx,
            *query,
        )

        if err != nil {
            return nil, err
        }

        for rows.Next() {
            product := domain.Product{}

            rows.Scan(
                &product.ID,
                &product.Name,
                &product.Price,
                &product.Description,
            )

            products = append(products, product)
        }
    }

    {
        err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

        if err != nil {
            return nil, err
        }
    }

    return &domain.Pagination[[]domain.Product]{
        Items: products,
        Total: total,
    }, nil
}