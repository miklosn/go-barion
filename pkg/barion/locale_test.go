package barion

import (
	"testing"
)

func TestUnmarshalInvalid(t *testing.T) {
	var locale Locale
	value := "xxx"
	err := locale.UnmarshalText([]byte(value))
	if err == nil {
		t.Fatal()
	}
}

func TestUnmarshalValidTable(t *testing.T) {
	var locale Locale

	tables := []struct {
		input  string
		output Locale
	}{
		{"hu-HU", HU},
	}

	for _, table := range tables {
		value := table.input
		locale.UnmarshalText([]byte(value))
		if locale != table.output {
			t.Fatal()
		}
	}
}
