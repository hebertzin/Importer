package domain

import (
	"context"
	"enube-challenge/packages/models"
)

type Supplier interface {
	Upload(ctx context.Context, suppliers *models.Supplier) error
}
