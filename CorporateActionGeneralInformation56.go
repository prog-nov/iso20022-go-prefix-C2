package iso20022

// General information about the corporate action event.
type CorporateActionGeneralInformation56 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType13Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation56) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) AddEventType() *CorporateActionEventType13Choice {
	c.EventType = new(CorporateActionEventType13Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation56) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation56) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}
