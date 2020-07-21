package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-resty/resty/v2"
	. "github.com/miklosn/go-barion/pkg/barion"
	"github.com/namsral/flag"
)

func main() {
	var (
		baseurl string
		poskey  string
	)
	flag.StringVar(&baseurl, "baseurl", "https://api.test.barion.com/v2", "base url")
	flag.StringVar(&poskey, "poskey", "", "pos key")
	flag.Parse()

	c := resty.New()
	c.SetDebug(true)
	barion := NewClient(baseurl, poskey, c)
	barion.SetLogger(log.Print)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Post("/callback", barion.CallbackHandler(func(state *PaymentState) {
		log.Printf("PaymentState callback result: %v", state)
	}))
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(":3000", r)
		log.Println(err)
		wg.Done()
	}()
	/*
		 := PaymentRequest{
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
		}*/

	/*
		response, err := barion.PaymentRequest(context.TODO(), &m)
		if err != nil {
			log.Fatal(spew.Sdump(err))
		}
		spew.Dump(response)
		log.Println(response.PaymentId)

		paymentState, err := barion.GetPaymentState(context.TODO(), response.PaymentId)
		if err != nil {
			log.Fatal(spew.Sdump(err))
		}
		spew.Dump(paymentState)
	*/

	wg.Wait()
}
