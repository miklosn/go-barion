package barion

import (
	"time"

	"github.com/shopspring/decimal"
)

type paymentStateRequest struct {
	POSKey    string `json:"POSKey,omitempty"`
	PaymentID string `json:"PaymentId"`
}

//go:generate enumer -type=CardType -json
type CardType int

const (
	Unknown CardType = iota
	MasterCard
	Visa
	AmericanExpress
	Electron
	Maestro
)

type BankCardType struct {
	MaskedPan      string
	BankCardType   CardType
	ValidThruYear  string
	ValidThruMonth string
}

type FundingInformation struct {
	AuthorizationCode string
	BankCard          BankCardType
}

type PaymentState struct {
	PaymentID             string `json:"PaymentId"`
	PaymentRequestID      string `json:"PaymentRequestId`
	POSId                 string
	POSName               string
	POSOwnerEmail         string
	Status                PaymentStatus
	PaymentType           PaymentType
	AllowedFundingSources *[]FundingSource `json:"FundingSources,omitempty"`
	FundingSource         *FundingSource
	FundingInformation    *FundingInformation
	GuestCheckout         bool
	CreatedAt             *time.Time
	StartedAt             *time.Time
	CompletedAt           *time.Time
	ValidUntil            *time.Time
	ReservedUntil         *time.Time
	DelayedCaptureUntil   *time.Time
	Total                 decimal.Decimal
	Currency              *Currency
	SuggestedLocale       *Locale
	FraudRiskScore        decimal.Decimal
	CallbackURL           string `json:"CallbackUrl"`
	RedirectURL           string `json:"RedirectUrl"`
}
