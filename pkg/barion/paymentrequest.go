package barion

type PaymentRequest struct {
	POSKey               string           `json:"POSKey,omitempty"`
	PaymentType          PaymentType      `json:"PaymentType"`
	ReservationPeriod    TimeSpan         `json:"ReservationPeriod,omitempty"`
	DelayedCapturePeriod TimeSpan         `json:"DelayedCapturePeriod,omitempty"`
	PaymentWindow        TimeSpan         `json:"PaymentWindow,omitempty"`
	GuestCheckout        bool             `json:"GuestCheckout"`
	InitiateRecurrence   bool             `json:"InitiateRecurrence"`
	RecurrenceID         string           `json:"RecurrenceId,omitempty"`
	FundingSources       *[]FundingSource `json:"FundingSources,omitempty"`
	PaymentRequestID     string           `json:"PaymentRequestId,omitempty"`
	PayerHint            string           `json:"PayerHint,omitempty"`
	CardHolderNameHint   string           `json:"CardHolderNameHint,omitempty"`
	RecurrenceType       RecurrenceType   `json:"RecurrenceType,omitempty"`
	RedirectURL          string           `json:"RedirectUrl,omitempty"`
	CallbackURL          string           `json:"CallbackUrl,omitempty"`
	Transactions         *[]PaymentTransaction
	OrderNumber          string           `json:"OrderNumber,omitempty"`
	ShippingAddress      *ShippingAddress `json:"ShippingAddress,omitempty"`
	Locale               Locale           `json:"Locale"`
	Currency             Currency         `json:"Currency"`
	PayerPhoneNumber     string           `json:"PayerPhoneNumber,omitempty"`
	PayerWorkPhoneNumber string           `json:"PayerWorkPhoneNumber,omitempty"`
	PayerHomeNumber      string           `json:"PayerHomeNumber,omitempty"`
	BillingAddress       *BillingAddress  `json:"BillingAddress,omitempty"`

	// Won't support:
	// - PayerAccount
	// Questionable:
	// - PurchaseInformation
	// - ChallengePreference
}
