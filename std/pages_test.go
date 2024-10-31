package std

import "testing"

func TestParseRange(t *testing.T) {
	start, end := ParseRange(1, 10, 1)
	if start != 0 || end != 1 {
		t.Error(start, end)
	}

	start, end = ParseRange(1, 1, 10)
	if start != 0 || end != 0 {
		t.Error(start, end)
	}
}
