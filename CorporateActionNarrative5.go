package iso20022

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative5 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation3 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation3 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation1 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`
}

func (c *CorporateActionNarrative5) AddAdditionalText() *UpdatedAdditionalInformation3 {
	c.AdditionalText = new(UpdatedAdditionalInformation3)
	return c.AdditionalText
}

func (c *CorporateActionNarrative5) AddNarrativeVersion() *UpdatedAdditionalInformation3 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation3)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative5) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative5) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative5) AddSecurityRestriction() *UpdatedAdditionalInformation1 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation1)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative5) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative5) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}
