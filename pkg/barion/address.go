package barion

type BillingAddress struct {
	Country string `json:"Country,omitempty"`
	City    string `json:"City,omitempty"`
	Region  string `json:"Region,omitempty"`
	Zip     string `json:"Zip,omitempty"`
	Street  string `json:"Street,omitempty"`
	Street2 string `json:"Street2,omitempty"`
	Street3 string `json:"Street3,omitempty"`
}

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
