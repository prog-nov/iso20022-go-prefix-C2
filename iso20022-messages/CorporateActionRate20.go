package iso20022

// Specifies rates.
type CorporateActionRate20 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat2Choice `xml:"GrssDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate20) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate20) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate20) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate20) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate20) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate20) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}
