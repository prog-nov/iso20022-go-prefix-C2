package iso20022

// Provides information about the corporate action option.
type CorporateActionOption3 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat2Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat2Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType1Choice `xml:"CertfctnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate8 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod5 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate5 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice6 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption1 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption6 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption4 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative5 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption3) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption3) AddOptionType() *CorporateActionOption2Choice {
	c.OptionType = new(CorporateActionOption2Choice)
	return c.OptionType
}

func (c *CorporateActionOption3) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption3) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddOptionFeatures() *OptionFeaturesFormat2Choice {
	newValue := new(OptionFeaturesFormat2Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat2Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat2Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption3) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption3) AddCertificationType() *BeneficiaryCertificationType1Choice {
	newValue := new(BeneficiaryCertificationType1Choice)
	c.CertificationType = append(c.CertificationType, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption3) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption3) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption3) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption3) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) AddSecurityIdentification() *SecurityIdentification11 {
	c.SecurityIdentification = new(SecurityIdentification11)
	return c.SecurityIdentification
}

func (c *CorporateActionOption3) AddDateDetails() *CorporateActionDate8 {
	c.DateDetails = new(CorporateActionDate8)
	return c.DateDetails
}

func (c *CorporateActionOption3) AddPeriodDetails() *CorporateActionPeriod5 {
	c.PeriodDetails = new(CorporateActionPeriod5)
	return c.PeriodDetails
}

func (c *CorporateActionOption3) AddRateAndAmountDetails() *CorporateActionRate5 {
	c.RateAndAmountDetails = new(CorporateActionRate5)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption3) AddPriceDetails() *CorporateActionPrice6 {
	c.PriceDetails = new(CorporateActionPrice6)
	return c.PriceDetails
}

func (c *CorporateActionOption3) AddSecuritiesQuantity() *SecuritiesOption1 {
	c.SecuritiesQuantity = new(SecuritiesOption1)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption3) AddSecuritiesMovementDetails() *SecuritiesOption6 {
	newValue := new(SecuritiesOption6)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddCashMovementDetails() *CashOption4 {
	newValue := new(CashOption4)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddAdditionalInformation() *CorporateActionNarrative5 {
	c.AdditionalInformation = new(CorporateActionNarrative5)
	return c.AdditionalInformation
}
