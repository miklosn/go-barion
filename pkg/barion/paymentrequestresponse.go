package barion

type ErrorResponse struct {
	status string
	Errors []Error
}

func (e *ErrorResponse) Error() string {
	return e.status
}

type Error struct {
	ErrorCode   string
	Title       string
	Description string
	EndPoint    string
	AuthData    string
	HappenedAt  string
}

type PaymentRequestResponse struct {
	Errors           []Error
	PaymentId        string
	PaymentRequestId string
	Status           PaymentStatus
	QRUrl            string
	GatewayUrl       string
	RedirectUrl      string
	CallbackUrl      string
}
