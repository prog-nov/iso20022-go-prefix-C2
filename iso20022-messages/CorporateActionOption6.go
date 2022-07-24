package iso20022

// Provides information about the corporate action option.
type CorporateActionOption6 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption4Choice `xml:"OptnTp"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (c *CorporateActionOption6) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption6) AddOptionType() *CorporateActionOption4Choice {
	c.OptionType = new(CorporateActionOption4Choice)
	return c.OptionType
}

func (c *CorporateActionOption6) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	c.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return c.InstructedOrQuantityToReceive
}
