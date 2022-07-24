package iso20022

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral17 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral8 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral2 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`
}

func (c *Collateral17) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral17) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) AddSecuritiesCollateral() *SecuritiesCollateral8 {
	newValue := new(SecuritiesCollateral8)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral17) AddCashCollateral() *CashCollateral2 {
	newValue := new(CashCollateral2)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral17) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}
