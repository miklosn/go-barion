package barion

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"golang.org/x/net/context"
)

func TestPaymentRequestError(t *testing.T) {
	c := resty.New()
	// c.SetDebug(true)
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()

	client := NewClient("http://localhost", c)
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
		t.Fatalf("Should not be err")
	}
	if res != nil {
		t.Fatalf("There shouldn't be a response")
	}
	e, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Should be an ErrorResponse answer")
	}
	if e.status != "501" {
		t.Fatalf("should give back api status")
	}
}

func TestPaymentRequestSuccess(t *testing.T) {
	c := resty.New()
	// c.SetDebug(true)
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()

	client := NewClient("http://localhost", c)
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
	if res.PaymentId != "31dfdaff269e4aa0b7a12e1c0cc2f933" {
		t.Fatalf("Paymentid should be given back")
	}
	if res.GatewayUrl != "https://secure.barion.com/Pay?Id=31dfdaff269e4aa0b7a12e1c0cc2f933" {
		t.Fatalf("GatewayUrl should be given back")
	}
}
