package customTesting

import "testing"

func AssertIntEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("assertion failed: %d != %d", a, b)
	}
}
