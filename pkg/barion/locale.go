package barion

import "fmt"

type Locale int

const (
	CZ Locale = iota
	DE
	US
	ES
	FR
	HU
	SK
	SI
)

func (l Locale) String() string {
	return [...]string{"cs-CZ", "de-DE", "en-US", "es-ES", "fr-FR", "hu-HU", "sk-SK", "sl-SI"}[l]
}

func (l Locale) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *Locale) UnmarshalText(text []byte) error {
	switch string(text) {
	case "cs-CZ":
		*l = CZ
	case "de-DE":
		*l = DE
	case "en-US":
		*l = US
	case "es-ES":
		*l = ES
	case "fr-FR":
		*l = FR
	case "hu-HU":
		*l = HU
	case "sk-SK":
		*l = SK
	case "sl-SI":
		*l = SI
	default:
		return fmt.Errorf("Invalid locale value %s", text)
	}
	return nil
}
