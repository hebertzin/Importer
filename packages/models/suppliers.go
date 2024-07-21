package models

type Supplier struct {
	ID                            uint   `gorm:"primaryKey"`
	PartnerId                     string `gorm:"type:text"`
	PartnerName                   string `gorm:"type:text"`
	CustomerId                    string `gorm:"type:text"`
	CustomerName                  string `gorm:"type:text"`
	CustomerDomainName            string `gorm:"type:text"`
	CustomerCountry               string `gorm:"type:text"`
	MpnId                         string `gorm:"type:text"`
	Tier2MpnId                    string `gorm:"type:text"`
	InvoiceNumber                 string `gorm:"type:text"`
	ProductId                     string `gorm:"type:text"`
	SkuId                         string `gorm:"type:text"`
	AvailabilityId                string `gorm:"type:text"`
	SkuName                       string `gorm:"type:text"`
	ProductName                   string `gorm:"type:text"`
	PublisherName                 string `gorm:"type:text"`
	PublisherId                   string `gorm:"type:text"`
	SubscriptionDescription       string `gorm:"type:text"`
	SubscriptionId                string `gorm:"type:text"`
	ChargeStartDate               string `gorm:"type:text"`
	ChargeEndDate                 string `gorm:"type:text"`
	UsageDate                     string `gorm:"type:text"`
	MeterType                     string `gorm:"type:text"`
	MeterCategory                 string `gorm:"type:text"`
	MeterId                       string `gorm:"type:text"`
	MeterSubCategory              string `gorm:"type:text"`
	MeterName                     string `gorm:"type:text"`
	MeterRegion                   string `gorm:"type:text"`
	Unit                          string `gorm:"type:text"`
	ResourceLocation              string `gorm:"type:text"`
	ConsumedService               string `gorm:"type:text"`
	ResourceGroup                 string `gorm:"type:text"`
	ResourceURI                   string `gorm:"type:text"`
	ChargeType                    string `gorm:"type:text"`
	UnitPrice                     string `gorm:"type:text"`
	Quantity                      string `gorm:"type:text"`
	UnitType                      string `gorm:"type:text"`
	BillingPreTaxTotal            string `gorm:"type:text"`
	BillingCurrency               string `gorm:"type:text"`
	PricingPreTaxTotal            string `gorm:"type:text"`
	PricingCurrency               string `gorm:"type:text"`
	ServiceInfo1                  string `gorm:"type:text"`
	ServiceInfo2                  string `gorm:"type:text"`
	Tags                          string `gorm:"type:text"`
	AdditionalInfo                string `gorm:"type:text"`
	EffectiveUnitPrice            string `gorm:"type:text"`
	PCToBCExchangeRate            string `gorm:"type:text"`
	PCToBCExchangeRateDate        string `gorm:"type:text"`
	EntitlementId                 string `gorm:"type:text"`
	EntitlementDescription        string `gorm:"type:text"`
	PartnerEarnedCreditPercentage string `gorm:"type:text"`
	CreditPercentage              string `gorm:"type:text"`
	CreditType                    string `gorm:"type:text"`
	BenefitOrderId                string `gorm:"type:text"`
	BenefitId                     string `gorm:"type:text"`
	BenefitType                   string `gorm:"type:text"`
}
