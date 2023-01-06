package timeparse

import (
	"testing"
)

func TestParseTimeString(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("err at valid input '%v'", err)
		}
		if !data.ok && err == nil {
			t.Errorf("no error at invalid input '%v'", data.time)
		}
	}
}
