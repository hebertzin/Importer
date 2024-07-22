package repository

import (
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/logging"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) domain.SupplierRepository {
	return &repository{db: db}
}

func (r *repository) SaveSuppliers(ctx context.Context, suppliersChan <-chan domain.Supplier, batchSize int) error {
	var suppliers []domain.Supplier
	for supplier := range suppliersChan {
		suppliers = append(suppliers, supplier)
		if len(suppliers) >= batchSize {
			if err := r.uploadSuppliers(ctx, suppliers); err != nil {
				return err
			}
			suppliers = nil
		}
	}
	if len(suppliers) > 0 {
		if err := r.uploadSuppliers(ctx, suppliers); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) uploadSuppliers(ctx context.Context, suppliers []domain.Supplier) error {
	if len(suppliers) == 0 {
		log.Println("No suppliers to upload")
		return nil
	}

	log.Printf("Uploading %d suppliers", len(suppliers))

	if err := r.db.WithContext(ctx).Create(&suppliers).Error; err != nil {
		log.Printf("Error uploading suppliers: %v", err)
		return err
	}

	logging.Log.Info("Suppliers uploaded successfully")
	return nil
}

func (r *repository) FindAllSuppliers(ctx context.Context, page, pageSize int) ([]domain.Supplier, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var suppliers []domain.Supplier
	result := r.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&suppliers)

	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

func (r *repository) FindSupplierById(ctx context.Context, id int) (*domain.Supplier, error) {
	var supplier domain.Supplier
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, result.Error
	}
	return &supplier, nil
}
