package barion

//go:generate enumer -type=FundingSource -json
type FundingSource int

const (
	All FundingSource = iota
	Balance
	BankCard
	BankTransfer
)
