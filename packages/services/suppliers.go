package services

import (
	"bytes"
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/logging"
	"fmt"
	"go.uber.org/zap"
	"sync"

	"github.com/xuri/excelize/v2"
)

type SupplierService interface {
	ImportSuppliersFromFile(ctx context.Context, file []byte) error
	GetSuppliers(ctx context.Context, page, pageSize int) ([]domain.Supplier, error)
	FindSupplierById(ctx context.Context, id int) (*domain.Supplier, error)
}

type supplierService struct {
	repo domain.SupplierRepository
}

func NewSupplierService(repo domain.SupplierRepository) SupplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) ImportSuppliersFromFile(ctx context.Context, file []byte) error {
	f, err := excelize.OpenReader(bytes.NewReader(file))
	if err != nil {
		return fmt.Errorf("failed to open excel file: %w", err)
	}

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return fmt.Errorf("no sheet found in file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		logging.Log.Error("failed to fetch rows", zap.Error(err))
		return err
	}

	const numWorkers = 7
	const batchSize = 1000

	var wg sync.WaitGroup
	rowChan := make(chan []string, numWorkers)
	supplierChan := make(chan domain.Supplier, batchSize)
	errChan := make(chan error, 1)

	worker := func() {
		defer wg.Done()
		for row := range rowChan {
			for len(row) < 55 {
				row = append(row, "")
			}

			supplier := domain.Supplier{
				PartnerId:                     row[0],
				PartnerName:                   row[1],
				CustomerId:                    row[2],
				CustomerName:                  row[3],
				CustomerDomainName:            row[4],
				CustomerCountry:               row[5],
				MpnId:                         row[6],
				Tier2MpnId:                    row[7],
				InvoiceNumber:                 row[8],
				ProductId:                     row[9],
				SkuId:                         row[10],
				AvailabilityId:                row[11],
				SkuName:                       row[12],
				ProductName:                   row[13],
				PublisherName:                 row[14],
				PublisherId:                   row[15],
				SubscriptionDescription:       row[16],
				SubscriptionId:                row[17],
				ChargeStartDate:               row[18],
				ChargeEndDate:                 row[19],
				UsageDate:                     row[20],
				MeterType:                     row[21],
				MeterCategory:                 row[22],
				MeterId:                       row[23],
				MeterSubCategory:              row[24],
				MeterName:                     row[25],
				MeterRegion:                   row[26],
				Unit:                          row[27],
				ResourceLocation:              row[28],
				ConsumedService:               row[29],
				ResourceGroup:                 row[30],
				ResourceURI:                   row[31],
				ChargeType:                    row[32],
				UnitPrice:                     row[33],
				Quantity:                      row[34],
				UnitType:                      row[35],
				BillingPreTaxTotal:            row[36],
				BillingCurrency:               row[37],
				PricingPreTaxTotal:            row[38],
				PricingCurrency:               row[39],
				ServiceInfo1:                  row[40],
				ServiceInfo2:                  row[41],
				Tags:                          row[42],
				AdditionalInfo:                row[43],
				EffectiveUnitPrice:            row[44],
				PCToBCExchangeRate:            row[45],
				PCToBCExchangeRateDate:        row[46],
				EntitlementId:                 row[47],
				EntitlementDescription:        row[48],
				PartnerEarnedCreditPercentage: row[49],
				CreditPercentage:              row[50],
				CreditType:                    row[51],
				BenefitOrderId:                row[52],
				BenefitId:                     row[53],
				BenefitType:                   row[54],
			}

			select {
			case supplierChan <- supplier:
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	go func() {
		for _, row := range rows[1:] {
			rowChan <- row
		}
		close(rowChan)
		wg.Wait()
		close(supplierChan)
	}()

	if err := s.repo.SaveSuppliers(ctx, supplierChan, batchSize); err != nil {
		logging.Log.Error("failed to save suppliers", zap.Error(err))
		return err
	}

	select {
	case err := <-errChan:
		return err
	default:
	}

	logging.Log.Info("successfully imported suppliers")
	return nil
}

func (s *supplierService) GetSuppliers(ctx context.Context, page, pageSize int) ([]domain.Supplier, error) {
	return s.repo.FindAllSuppliers(ctx, page, pageSize)
}

func (s *supplierService) FindSupplierById(ctx context.Context, id int) (*domain.Supplier, error) {
	return s.repo.FindSupplierById(ctx, id)
}
