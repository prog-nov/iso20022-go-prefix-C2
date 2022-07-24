package iso20022

// Provides information about the corporate action option.
type CorporateActionOption135 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat23Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat18Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus4Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType12Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate55 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate84 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice65 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption56 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption61 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption54 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative36 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption135) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption135) AddOptionType() *CorporateActionOption23Choice {
	c.OptionType = new(CorporateActionOption23Choice)
	return c.OptionType
}

func (c *CorporateActionOption135) AddFractionDisposition() *FractionDispositionType31Choice {
	c.FractionDisposition = new(FractionDispositionType31Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption135) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddOptionFeatures() *OptionFeaturesFormat23Choice {
	newValue := new(OptionFeaturesFormat23Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat18Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat18Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption135) AddOptionAvailabilityStatus() *OptionAvailabilityStatus4Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus4Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption135) AddCertificationBreakdownType() *BeneficiaryCertificationType12Choice {
	newValue := new(BeneficiaryCertificationType12Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption135) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption135) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption135) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption135) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption135) AddDateDetails() *CorporateActionDate55 {
	c.DateDetails = new(CorporateActionDate55)
	return c.DateDetails
}

func (c *CorporateActionOption135) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption135) AddRateAndAmountDetails() *CorporateActionRate84 {
	c.RateAndAmountDetails = new(CorporateActionRate84)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption135) AddPriceDetails() *CorporateActionPrice65 {
	c.PriceDetails = new(CorporateActionPrice65)
	return c.PriceDetails
}

func (c *CorporateActionOption135) AddSecuritiesQuantity() *SecuritiesOption56 {
	c.SecuritiesQuantity = new(SecuritiesOption56)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption135) AddSecuritiesMovementDetails() *SecuritiesOption61 {
	newValue := new(SecuritiesOption61)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddCashMovementDetails() *CashOption54 {
	newValue := new(CashOption54)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddAdditionalInformation() *CorporateActionNarrative36 {
	c.AdditionalInformation = new(CorporateActionNarrative36)
	return c.AdditionalInformation
}
