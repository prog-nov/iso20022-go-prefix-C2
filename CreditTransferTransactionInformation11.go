package iso20022

// Set of elements used to provide information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation11 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SettlementTimeIndication *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty"`

	// Provides information on the requested settlement time(s) of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount16 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount16 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount16 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount16 `xml:"IntrmyAgt3Acct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation11) AddPaymentIdentification() *PaymentIdentification3 {
	c.PaymentIdentification = new(PaymentIdentification3)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation11) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	c.PaymentTypeInformation = new(PaymentTypeInformation21)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation11) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation11) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation11) SetSettlementPriority(value string) {
	c.SettlementPriority = (*Priority3Code)(&value)
}

func (c *CreditTransferTransactionInformation11) AddSettlementTimeIndication() *SettlementDateTimeIndication1 {
	c.SettlementTimeIndication = new(SettlementDateTimeIndication1)
	return c.SettlementTimeIndication
}

func (c *CreditTransferTransactionInformation11) AddSettlementTimeRequest() *SettlementTimeRequest2 {
	c.SettlementTimeRequest = new(SettlementTimeRequest2)
	return c.SettlementTimeRequest
}

func (c *CreditTransferTransactionInformation11) SetAcceptanceDateTime(value string) {
	c.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (c *CreditTransferTransactionInformation11) SetPoolingAdjustmentDate(value string) {
	c.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation11) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation11) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CreditTransferTransactionInformation11) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransactionInformation11) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	c.ChargesInformation = append(c.ChargesInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation11) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransactionInformation11) AddPreviousInstructingAgentAccount() *CashAccount16 {
	c.PreviousInstructingAgentAccount = new(CashAccount16)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransactionInformation11) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.InstructingAgent
}

func (c *CreditTransferTransactionInformation11) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.InstructedAgent
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent1Account() *CashAccount16 {
	c.IntermediaryAgent1Account = new(CashAccount16)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent2Account() *CashAccount16 {
	c.IntermediaryAgent2Account = new(CashAccount16)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation11) AddIntermediaryAgent3Account() *CashAccount16 {
	c.IntermediaryAgent3Account = new(CashAccount16)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation11) AddUltimateDebtor() *PartyIdentification32 {
	c.UltimateDebtor = new(PartyIdentification32)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation11) AddInitiatingParty() *PartyIdentification32 {
	c.InitiatingParty = new(PartyIdentification32)
	return c.InitiatingParty
}

func (c *CreditTransferTransactionInformation11) AddDebtor() *PartyIdentification32 {
	c.Debtor = new(PartyIdentification32)
	return c.Debtor
}

func (c *CreditTransferTransactionInformation11) AddDebtorAccount() *CashAccount16 {
	c.DebtorAccount = new(CashAccount16)
	return c.DebtorAccount
}

func (c *CreditTransferTransactionInformation11) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.DebtorAgent
}

func (c *CreditTransferTransactionInformation11) AddDebtorAgentAccount() *CashAccount16 {
	c.DebtorAgentAccount = new(CashAccount16)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransactionInformation11) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation11) AddCreditorAgentAccount() *CashAccount16 {
	c.CreditorAgentAccount = new(CashAccount16)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation11) AddCreditor() *PartyIdentification32 {
	c.Creditor = new(PartyIdentification32)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation11) AddCreditorAccount() *CashAccount16 {
	c.CreditorAccount = new(CashAccount16)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation11) AddUltimateCreditor() *PartyIdentification32 {
	c.UltimateCreditor = new(PartyIdentification32)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation11) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation11) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstructionForNextAgent = append(c.InstructionForNextAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation11) AddPurpose() *Purpose2Choice {
	c.Purpose = new(Purpose2Choice)
	return c.Purpose
}

func (c *CreditTransferTransactionInformation11) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation11) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation11) AddRemittanceInformation() *RemittanceInformation5 {
	c.RemittanceInformation = new(RemittanceInformation5)
	return c.RemittanceInformation
}
