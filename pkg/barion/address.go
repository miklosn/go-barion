package barion

// BillingAddress defines the billing address associated with the payment, if applicable.
type BillingAddress struct {
	Country string `json:"Country"`           // The payer's country code in ISO-3166-1 format. E.g. HU or DE.
	City    string `json:"City,omitempty"`    // The complete name of the city of the recipient address.
	Region  string `json:"Region,omitempty"`  // The country subdivision code of the recipient address in ISO-3166-2 format
	Zip     string `json:"Zip,omitempty"`     // The zip code of the recipient address.
	Street  string `json:"Street,omitempty"`  // The shipping street address with house number and other details.
	Street2 string `json:"Street2,omitempty"` // Street address, continued.
	Street3 string `json:"Street3,omitempty"` // Street address, continued.
}

// ShippingAddress defines the shipping address associated with the payment, if applicable.
type ShippingAddress struct {
	Country  string `json:"Country,omitempty"`
	City     string `json:"City,omitempty"`
	Region   string `json:"Region,omitempty"`
	Zip      string `json:"Zip,omitempty"`
	Street   string `json:"Street,omitempty"`
	Street2  string `json:"Street2,omitempty"`
	Street3  string `json:"Street3,omitempty"`
	FullName string `json:"FullName,omitempty"`
}
