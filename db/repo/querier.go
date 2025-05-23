// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"context"
)

type Querier interface {
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	GetCustomersByNameLike(ctx context.Context, name string) ([]Customer, error)
	GetOrder(ctx context.Context, id string) (Order, error)
	GetTotalRevenueForProduct(ctx context.Context, productID *string) (int64, error)
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) ([]Customer, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) ([]Product, error)
}

var _ Querier = (*Queries)(nil)
