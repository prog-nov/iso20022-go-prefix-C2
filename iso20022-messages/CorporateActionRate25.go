package iso20022

// Specifies rates related to a corporate action option.
type CorporateActionRate25 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat5Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat5Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat6Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat6Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate25) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate25) AddGrossDividendRate() *GrossDividendRateFormat5Choice {
	newValue := new(GrossDividendRateFormat5Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddNetDividendRate() *NetDividendRateFormat5Choice {
	newValue := new(NetDividendRateFormat5Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate25) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddMaximumAllowedOversubscriptionRate() *RateFormat6Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat6Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate25) AddProrationRate() *RateFormat6Choice {
	c.ProrationRate = new(RateFormat6Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate25) AddWithholdingTaxRate() *RateFormat6Choice {
	c.WithholdingTaxRate = new(RateFormat6Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate25) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate25) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}
