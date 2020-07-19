package barion

type RecurrenceType int

const (
	MerchantInitiatedPayment RecurrenceType = 0
	OneClickPayment          RecurrenceType = 10
)
