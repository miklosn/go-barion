package barion

import (
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	var timespan TimeSpan = TimeSpan(time.Minute * 30)
	value, err := timespan.MarshalText()
	if err != nil {
		t.Fatal()
	}
	if string(value) != "0.00:30:00" {
		t.Fatalf("Received %s", string(value))
	}
}

func TestMarshal2d(t *testing.T) {
	var timespan TimeSpan = TimeSpan(time.Hour * 4)
	value, err := timespan.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(value) != "0.04:00:00" {
		t.Fatalf("Received %s", string(value))
	}
}

func TestMarshal7d(t *testing.T) {
	var timespan TimeSpan = TimeSpan(time.Hour * 169)
	value, err := timespan.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(value) != "7.01:00:00" {
		t.Fatalf("Received %s", string(value))
	}
}

func TestUnMarshal1Second(t *testing.T) {
	var input = "0.00:00:1"
	var timespan TimeSpan
	err := timespan.UnmarshalText([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if timespan != 1000000000 {
		t.Fatal()
	}
}

func TestUnMarshal1Day1Sec(t *testing.T) {
	var input = "1.00:00:1"
	var timespan TimeSpan
	err := timespan.UnmarshalText([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if timespan != 86401000000000 {
		t.Fatal()
	}
}
