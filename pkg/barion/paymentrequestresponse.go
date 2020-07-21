package barion

type PaymentRequestResponse struct {
	Errors           *[]barionError
	PaymentID        string
	PaymentRequestID string
	Status           PaymentStatus
	QRUrl            string
	GatewayURL       string
	RedirectURL      string
	CallbackURL      string
}
