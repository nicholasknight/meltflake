package meltflake

import (
	"testing"
	"time"
)

func TestMelt(t *testing.T) {
	expectedTime := time.Date(2017, 11, 19, 20, 6, 51, 707000000, time.UTC).String()
	sf := Melt(381898139189116930, Discord)
	gotTime := sf.Time.UTC().String()
	if gotTime != expectedTime {
		t.Errorf("Expected %s, got %s", expectedTime, gotTime)
	}

	if sf.WorkerID != 0 {
		t.Errorf("Expected worker ID 0 in %s", sf)
	}

	if sf.ProcessID != 0 {
		t.Errorf("Expected process ID 0 in %s", sf)
	}

	if sf.Increment != 2 {
		t.Errorf("Expected increment 2 in %s", sf)
	}
}

func TestStringer(t *testing.T) {
	expected := "2017-11-19 20:06:51.707 Z, w 0, p 0, i 2"
	sf := Melt(381898139189116930, Discord)
	got := sf.String()
	if expected != got {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
