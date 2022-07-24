package iso20022

// General information about the corporate action event.
type CorporateActionGeneralInformation85 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType33Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes66 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation85) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation85) AddEventType() *CorporateActionEventType33Choice {
	c.EventType = new(CorporateActionEventType33Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation85) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation85) AddUnderlyingSecurity() *FinancialInstrumentAttributes66 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes66)
	return c.UnderlyingSecurity
}
