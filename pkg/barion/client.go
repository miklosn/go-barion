package barion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type logger func(...interface{})

type client struct {
	poskey  string
	baseurl string
	r       *resty.Client
	logger  logger
}

func NewClient(baseurl string, poskey string, r *resty.Client) *client {
	r.SetHeader("User-Agent", "github.com/miklosn/go-barion")
	return &client{
		r:       r,
		baseurl: baseurl,
		poskey:  poskey,
		logger:  log.Print,
	}
}

func (c *client) SetLogger(logger logger) {
	c.logger = logger
}

func (c *client) PaymentRequest(ctx context.Context, request *PaymentRequest) (*PaymentRequestResponse, error) {
	url := c.baseurl + "/Payment/Start"
	request.POSKey = c.poskey
	request.CallbackURL = "https://root-2c3992af.localhost.run/callback"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req := c.r.R()
	req.SetContext(ctx)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(body)
	req.SetResult(&PaymentRequestResponse{})
	req.SetError(&ErrorResponse{})

	res, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, err)
	}

	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*ErrorResponse)
			x.status = res.Status()
		}
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, res.Error().(*ErrorResponse))
	}
	return res.Result().(*PaymentRequestResponse), nil
}

func (c *client) GetPaymentState(ctx context.Context, PaymentId string) (*PaymentState, error) {
	url := c.baseurl + "/Payment/GetPaymentState"
	req := c.r.R()
	req.SetContext(ctx)
	req.SetHeader("Content-Type", "application/json")
	req.SetQueryParam("POSKey", c.poskey)
	req.SetQueryParam("PaymentId", PaymentId)
	req.SetResult(PaymentState{})
	req.SetError(ErrorResponse{})
	res, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("1Error getting payment state for payment id %s: %v", PaymentId, err)
	}
	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*ErrorResponse)
			x.status = res.Status()
			return nil, fmt.Errorf("2Error getting payment state for payment id %s: %v", PaymentId, res.Error())
		}
		return nil, fmt.Errorf("3Error getting payment state for payment id %s: %v", PaymentId, &ErrorResponse{
			status: res.Status(),
		})
	}
	return res.Result().(*PaymentState), nil
}
