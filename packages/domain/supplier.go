package domain

import (
	"context"
)

type Supplier struct {
	ID                            int
	PartnerId                     string
	PartnerName                   string
	CustomerId                    string
	CustomerName                  string
	CustomerDomainName            string
	CustomerCountry               string
	MpnId                         string
	Tier2MpnId                    string
	InvoiceNumber                 string
	ProductId                     string
	SkuId                         string
	AvailabilityId                string
	SkuName                       string
	ProductName                   string
	PublisherName                 string
	PublisherId                   string
	SubscriptionDescription       string
	SubscriptionId                string
	ChargeStartDate               string
	ChargeEndDate                 string
	UsageDate                     string
	MeterType                     string
	MeterCategory                 string
	MeterId                       string
	MeterSubCategory              string
	MeterName                     string
	MeterRegion                   string
	Unit                          string
	ResourceLocation              string
	ConsumedService               string
	ResourceGroup                 string
	ResourceURI                   string
	ChargeType                    string
	UnitPrice                     string
	Quantity                      string
	UnitType                      string
	BillingPreTaxTotal            string
	BillingCurrency               string
	PricingPreTaxTotal            string
	PricingCurrency               string
	ServiceInfo1                  string
	ServiceInfo2                  string
	Tags                          string
	AdditionalInfo                string
	EffectiveUnitPrice            string
	PCToBCExchangeRate            string
	PCToBCExchangeRateDate        string
	EntitlementId                 string
	EntitlementDescription        string
	PartnerEarnedCreditPercentage string
	CreditPercentage              string
	CreditType                    string
	BenefitOrderId                string
	BenefitId                     string
	BenefitType                   string
}

type SupplierRepository interface {
	SaveSuppliers(ctx context.Context, suppliersChan <-chan Supplier, batchSize int) error
	FindAllSuppliers(ctx context.Context, page, pageSize int) ([]Supplier, error)
	FindSupplierById(ctx context.Context, id int) (*Supplier, error)
}
