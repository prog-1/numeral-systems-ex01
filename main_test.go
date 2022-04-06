package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestIsValidNumber(t *testing.T) {
	for _, tc := range []struct {
		num  string
		base int
		want bool
	}{
		{num: "123", base: 10, want: true},
		{num: "123", base: 1, want: false},
		{num: "123AB", base: 16, want: true},
		{num: "123ZZZ", base: 3, want: false},
		{num: "123456789", base: 10, want: true},
		{num: "123AB", base: 10, want: false},
		{num: "", base: 10, want: true},
		{num: "000", base: 0, want: false},
		{num: "000", base: 1, want: false}, // i found only theories about base 0 or 1, but noone use them in any practical case, so i think false here will be better
	} {
		if got := isValidNumber(tc.num, tc.base); got != tc.want {
			t.Errorf("isValidNumber(%v, %v) = (%v), want = (%v)", tc.num, tc.base, got, tc.want)
		}
	}
}
