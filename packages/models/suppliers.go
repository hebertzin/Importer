package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID                            uint           `gorm:"primaryKey"`
	PartnerId                     string         `gorm:"size:255"`
	PartnerName                   string         `gorm:"size:255"`
	CustomerId                    string         `gorm:"size:255"`
	CustomerName                  string         `gorm:"size:255"`
	CustomerDomainName            string         `gorm:"size:255"`
	CustomerCountry               string         `gorm:"size:255"`
	MpnId                         string         `gorm:"size:255"`
	Tier2MpnId                    string         `gorm:"size:255"`
	InvoiceNumber                 string         `gorm:"size:255"`
	ProductId                     string         `gorm:"size:255"`
	SkuId                         string         `gorm:"size:255"`
	AvailabilityId                string         `gorm:"size:255"`
	SkuName                       string         `gorm:"size:255"`
	ProductName                   string         `gorm:"size:255"`
	PublisherName                 string         `gorm:"size:255"`
	PublisherId                   string         `gorm:"size:255"`
	SubscriptionDescription       string         `gorm:"size:255"`
	SubscriptionId                string         `gorm:"size:255"`
	ChargeStartDate               string         `gorm:"size:255"`
	ChargeEndDate                 string         `gorm:"size:255"`
	UsageDate                     string         `gorm:"size:255"`
	MeterType                     string         `gorm:"size:255"`
	MeterCategory                 string         `gorm:"size:255"`
	MeterId                       string         `gorm:"size:255"`
	MeterSubCategory              string         `gorm:"size:255"`
	MeterName                     string         `gorm:"size:255"`
	MeterRegion                   string         `gorm:"size:255"`
	Unit                          string         `gorm:"size:255"`
	ResourceLocation              string         `gorm:"size:255"`
	ConsumedService               string         `gorm:"size:255"`
	ResourceGroup                 string         `gorm:"size:255"`
	ResourceURI                   string         `gorm:"size:255"`
	ChargeType                    string         `gorm:"size:255"`
	UnitPrice                     string         `gorm:"size:255"`
	Quantity                      string         `gorm:"size:255"`
	UnitType                      string         `gorm:"size:255"`
	BillingPreTaxTotal            string         `gorm:"size:255"`
	BillingCurrency               string         `gorm:"size:255"`
	PricingPreTaxTotal            string         `gorm:"size:255"`
	PricingCurrency               string         `gorm:"size:255"`
	ServiceInfo1                  string         `gorm:"size:255"`
	ServiceInfo2                  string         `gorm:"size:255"`
	Tags                          string         `gorm:"size:255"`
	AdditionalInfo                string         `gorm:"size:255"`
	EffectiveUnitPrice            string         `gorm:"size:255"`
	PCToBCExchangeRate            string         `gorm:"size:255"`
	PCToBCExchangeRateDate        string         `gorm:"size:255"`
	EntitlementId                 string         `gorm:"size:255"`
	EntitlementDescription        string         `gorm:"size:255"`
	PartnerEarnedCreditPercentage string         `gorm:"size:255"`
	CreditPercentage              string         `gorm:"size:255"`
	CreditType                    string         `gorm:"size:255"`
	BenefitOrderId                string         `gorm:"size:255"`
	BenefitId                     string         `gorm:"size:255"`
	BenefitType                   string         `gorm:"size:255"`
	CreatedAt                     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt                     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt                     gorm.DeletedAt `gorm:"index"`
}
