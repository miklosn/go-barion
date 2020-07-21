package barion

//go:generate enumer -type=PaymentStatus -json
type PaymentStatus int

const (
	Prepared PaymentStatus = iota
	Started
	InProgress
	Waiting
	Reserved
	Authorized
	Canceled
	Succeeded
	Failed
	PartiallySucceeded
	Expired
)
