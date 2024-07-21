package services

import (
	"bytes"
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type SupplierService interface {
	ImportSuppliersFromFile(ctx context.Context, file []byte) error
}

type supplierService struct {
	repo domain.Supplier
}

func NewSupplierService(repo domain.Supplier) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) ImportSuppliersFromFile(ctx context.Context, file []byte) error {
	f, err := excelize.OpenReader(bytes.NewReader(file))
	if err != nil {
		return err
	}

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return fmt.Errorf("no sheet found in file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	for _, row := range rows[1:] {
		if len(row) < 55 {
			return fmt.Errorf("row has insufficient columns: %v", row)
		}

		supplier := models.Supplier{
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

		if err := s.repo.Upload(ctx, &supplier); err != nil {
			return err
		}
	}

	return nil
}
