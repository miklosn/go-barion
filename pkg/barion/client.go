package barion

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type client struct {
	baseurl string
	r       *resty.Client
}

func NewClient(baseurl string, r *resty.Client) *client {
	return &client{
		r:       r,
		baseurl: baseurl,
	}
}

func (c *client) PaymentRequest(ctx context.Context, request *PaymentRequest) (*PaymentRequestResponse, error) {
	url := c.baseurl + "/Payment/Start"

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
		return nil, err
	}

	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*ErrorResponse)
			x.status = res.Status()
		}
		return nil, res.Error().(*ErrorResponse)
	}
	return res.Result().(*PaymentRequestResponse), nil
}
