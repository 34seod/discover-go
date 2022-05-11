package customTesting

import "testing"

type aaa interface {
	int | float32 | float64
}

func TestFib(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{6, 8},

	for _, c := range cases {
		got := seq.Fib(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
