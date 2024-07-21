package domain

import (
	"context"
	"enube-challenge/packages/models"
)

type Supplier interface {
	SaveSuppliers(ctx context.Context, suppliersChan <-chan models.Supplier, batchSize int) error
	FindAllSuppliers(ctx context.Context, page, pageSize int) ([]models.Supplier, error)
	FindSupplierById(ctx context.Context, id int) (*models.Supplier, error)
}
