package barion

//go:generate enumer -type=Currency -json
type Currency int

const (
	HUF Currency = iota
	EUR
	USD
	CZK
)
