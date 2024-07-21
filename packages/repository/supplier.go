package repository

import (
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/models"
	"gorm.io/gorm"
)

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) domain.Supplier {
	return &supplierRepository{
		db: db,
	}
}

func (r *supplierRepository) Upload(ctx context.Context, supplier *models.Supplier) error {
	return r.db.Create(&supplier).Error
}
