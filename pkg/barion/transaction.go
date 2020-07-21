package barion

import (
	"github.com/shopspring/decimal"
)

type Item struct {
	Name        string          `json:"Name,omitempty"`
	Description string          `json:"Description,omitempty"`
	ImageURL    string          `json:"ImageUrl,omitempty"`
	Quantity    decimal.Decimal `json:"Quantity,omitempty"`
	Unit        string          `json:"Unit,omitempty"`
	UnitPrice   decimal.Decimal `json:"UnitPrice,omitempty"`
	ItemTotal   decimal.Decimal `json:"ItemTotal,omitempty"`
	SKU         string          `json:"SKU,omitempty"`
}

type PayeeTransaction struct {
	POSTransactionID string
	Payee            string
	Total            decimal.Decimal
	Comment          string
}

type PaymentTransaction struct {
	POSTransactionID  string
	Payee             string
	Total             decimal.Decimal
	Comment           string
	PayeeTransactions *[]PayeeTransaction `json:"PayeeTransactions,omitempty"`
	Items             *[]Item
}
