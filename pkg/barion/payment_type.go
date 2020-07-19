package barion

//go:generate enumer -type=PaymentType -json
type PaymentType int

const (
	Immediate PaymentType = iota
	Reservation
	DelayedCapture
)
