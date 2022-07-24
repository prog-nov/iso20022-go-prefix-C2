package iso20022

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative28 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation1 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation1 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation1 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation1 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation1 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation1 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative28) AddAdditionalText() *UpdatedAdditionalInformation1 {
	c.AdditionalText = new(UpdatedAdditionalInformation1)
	return c.AdditionalText
}

func (c *CorporateActionNarrative28) AddNarrativeVersion() *UpdatedAdditionalInformation1 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation1)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative28) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative28) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative28) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative28) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

func (c *CorporateActionNarrative28) AddPartyContactNarrative() *UpdatedAdditionalInformation1 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation1)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative28) AddRegistrationDetails() *UpdatedAdditionalInformation1 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation1)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative28) AddBasketOrIndexInformation() *UpdatedAdditionalInformation1 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation1)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative28) AddCertificationBreakdown() *UpdatedAdditionalInformation1 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation1)
	return c.CertificationBreakdown
}
