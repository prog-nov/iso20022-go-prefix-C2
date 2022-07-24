package iso20022

// Cryptographic Key to exchange.
type CryptographicKey2 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType2Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType5 `xml:"KeyVal"`
}

func (c *CryptographicKey2) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey2) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey2) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}

func (c *CryptographicKey2) SetType(value string) {
	c.Type = (*CryptographicKeyType2Code)(&value)
}

func (c *CryptographicKey2) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey2) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey2) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey2) AddKeyValue() *ContentInformationType5 {
	c.KeyValue = new(ContentInformationType5)
	return c.KeyValue
}
