package barion

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"golang.org/x/net/context"
)

func TestPaymentRequestError(t *testing.T) {
	c := resty.New()
	//c.SetDebug(true)
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()

	client := NewClient("http://localhost", "XXX", c)
	response := httpmock.NewStringResponse(501, `
			{
			  "Errors": [
			    {
			      "ErrorCode": "NotExistingPaymentId",
			      "Title": "The given payment id is invalid",
			      "Description": "The given payment id(510bbd38-08c9-40b9-9e21-c6f6cb9eb4f1) is invalid!",
			      "EndPoint": "https://api.barion.com/v2/Payment/GetPaymentState?POSKey=5b65f8f7a5f4405bb19495383dcf170a&PaymentId=510bbd3808c940b99e21c6f6cb9eb4f1",
			      "AuthData": "shop@example.com",
			      "HappenedAt": "2017-02-06T09:48:13.7694921Z"
			    }
			  ]
			}
		`)
	response.Header.Set("content-type", "application/json")
	httpmock.RegisterResponder("POST", "http://localhost/Payment/Start", httpmock.ResponderFromResponse(response))
	res, err := client.PaymentRequest(context.Background(), &PaymentRequest{})
	if err == nil {
		t.Fatalf("Should be err")
	}
	if res != nil {
		t.Fatalf("There shouldn't be a response")
	}
}

func TestPaymentRequestSuccess(t *testing.T) {
	c := resty.New()
	// c.SetDebug(true)
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()

	client := NewClient("http://localhost", "xxx", c)
	response := httpmock.NewStringResponse(200, `
		{
		  "PaymentId": "31dfdaff269e4aa0b7a12e1c0cc2f933",
		  "PaymentRequestId": "EXMPLSHOP-PM-001",
		  "Status": "Prepared",
		  "QRUrl": "https://api.barion.com/qr/generate?paymentId=31dfdaff-269e-4aa0-b7a1-2e1c0cc2f933&size=Large",
		  "Transactions": [
		    {
		      "POSTransactionId": "EXMPLSHOP-PM-001/TR001",
		      "TransactionId": "a4ead1d965bd48b3934847504c67ec8c",
		      "Status": "Prepared",
		      "Currency": "EUR",
		      "TransactionTime": "2017-02-03T13:07:10.215",
		      "RelatedId": null
		    }
		  ],
		  "RecurrenceResult": "None",
		  "GatewayUrl": "https://secure.barion.com/Pay?Id=31dfdaff269e4aa0b7a12e1c0cc2f933",
		  "RedirectUrl": "https://webshop.example.com/Redirect?paymentId=31dfdaff269e4aa0b7a12e1c0cc2f933",
		  "CallbackUrls": "https://webshop.example.com/blank.html?paymentId=31dfdaff269e4aa0b7a12e1c0cc2f933",
		  "Errors": []
		}
		`)
	response.Header.Set("content-type", "application/json")
	httpmock.RegisterResponder("POST", "http://localhost/Payment/Start", httpmock.ResponderFromResponse(response))
	res, err := client.PaymentRequest(context.Background(), &PaymentRequest{})
	if err != nil {
		t.Fatalf("Should not be err")
	}
	if res.PaymentID != "31dfdaff269e4aa0b7a12e1c0cc2f933" {
		t.Fatalf("Paymentid should be given back")
	}
	if res.GatewayURL != "https://secure.barion.com/Pay?Id=31dfdaff269e4aa0b7a12e1c0cc2f933" {
		t.Fatalf("GatewayUrl should be given back")
	}
}

func TestGetPaymentState(t *testing.T) {
	c := resty.New()
	// c.SetDebug(true)
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()
	client := NewClient("http://localhost", "xxx", c)
	response := httpmock.NewStringResponse(200, `
		{"PaymentId":"8a879c1fa4c9ea118bbd001dd8b71cc4","PaymentRequestId":"EXMPLSHOP-PM-001","OrderNumber":null,"POSId":"84717d42b27647c5bbb2bf3f6d7756f3","POSName":"Teszt 91","POSOwnerEmail":"miklos.niedermayer@cray.one","Status":"Prepared","PaymentType":"Immediate","FundingSource":null,"FundingInformation":null,"AllowedFundingSources":["All"],"GuestCheckout":true,"CreatedAt":"2020-07-19T09:42:14.021Z","ValidUntil":"2020-07-19T10:12:14.021Z","CompletedAt":null,"ReservedUntil":null,"DelayedCaptureUntil":null,"Transactions":[{"TransactionId":"8b879c1fa4c9ea118bbd001dd8b71cc4","POSTransactionId":"EXMPLSHOP-PM-001/TR001","TransactionTime":"2020-07-19T09:42:14.021Z","Total":37.00,"Currency":"HUF","Payer":null,"Payee":{"Name":{"LoginName":"miklos.niedermayer@cray.one","FirstName":"Miklós","LastName":"Niedermayer","OrganizationName":null},"Email":"miklos.niedermayer@cray.one"},"Comment":"A brief description of the transaction","Status":"Prepared","TransactionType":"Unspecified","Items":[{"Name":"iPhone 7 smart case","Description":"Durable elegant phone case / matte black","Quantity":1.00,"Unit":"piece","UnitPrice":37.00,"ItemTotal":37.00,"SKU":"EXMPLSHOP/SKU/PHC-01"}],"RelatedId":null,"POSId":"84717d42b27647c5bbb2bf3f6d7756f3","PaymentId":"8a879c1fa4c9ea118bbd001dd8b71cc4"}],"Total":37.00,"SuggestedLocale":"hu-HU","FraudRiskScore":null,"RedirectUrl":"https://example.com/test?paymentId=8a879c1fa4c9ea118bbd001dd8b71cc4","CallbackUrl":null,"Currency":"HUF","Errors":[]}
	`)
	response.Header.Set("content-type", "application/json")
	httpmock.RegisterResponder("GET", "http://localhost/Payment/GetPaymentState", httpmock.ResponderFromResponse(response))
	res, err := client.GetPaymentState(context.Background(), "8a879c1fa4c9ea118bbd001dd8b71cc4")
	if err != nil {
		t.Fatalf("Should not be err")
	}
	if res.PaymentID != "8a879c1fa4c9ea118bbd001dd8b71cc4" {
		t.Fatalf("Paymentid should be given back")
	}
}
