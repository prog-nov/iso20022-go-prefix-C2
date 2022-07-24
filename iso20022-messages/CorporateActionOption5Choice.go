package iso20022

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption5Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption6Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption5Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption6Code)(&value)
}

func (c *CorporateActionOption5Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
