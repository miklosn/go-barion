package main

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	. "github.com/miklosn/go-barion/pkg/barion"
	"github.com/shopspring/decimal"
	"log"
)

const (
	//BaseURL = "https://apia.test.barion.com/v2"
	BaseURL = "https://www.google.com"
	//BaseURL = "http://localhost:8000"
)

func main() {
	m := PaymentRequest{
		POSKey:           "c25fb8a5685f459482bb3ce47c732bd4-",
		PaymentType:      Immediate,
		PaymentRequestId: "EXMPLSHOP-PM-001",
		FundingSources:   []FundingSource{All},
		Currency:         HUF,
		Locale:           HU,
		GuestCheckout:    true,
		RedirectUrl:      "https://example.com/test",
		Transactions: &[]PaymentTransaction{
			{
				POSTransactionId: "EXMPLSHOP-PM-001/TR001",
				Payee:            "miklos.niedermayer@cray.one",
				Total:            decimal.NewFromInt(37),
				Comment:          "A brief description of the transaction",
				Items: []Item{
					{
						Name:        "iPhone 7 smart case",
						Description: "Durable elegant phone case / matte black",
						Quantity:    decimal.NewFromInt(1),
						Unit:        "piece",
						UnitPrice:   decimal.NewFromInt(37),
						ItemTotal:   decimal.NewFromInt(37),
						SKU:         "EXMPLSHOP/SKU/PHC-01",
					},
				},
			},
		},
	}

	barion := NewClient(BaseURL)

	response, err := barion.PaymentRequest(context.TODO(), &m)
	if err != nil {
		log.Fatal(spew.Sdump(err))
	}
	spew.Dump(response)
}
