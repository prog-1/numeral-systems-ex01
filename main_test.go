package main

import "testing"

func TestIsValidNumber(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  string
		base int
		want bool
	}{
		{name: "valid number with the base 10", num: "123", base: 10, want: true},
		{name: "invalid number witj invalid base", num: "123", base: 1, want: false},
		{name: "valid number with letters and base", num: "123AB", base: 16, want: true},
		{name: "inval num, val base", num: "123ZZZ", base: 3, want: false},
		{name: "valid number with the base 10", num: "123456789", base: 10, want: true},
		{name: "valid num, not right base", num: "123AB", base: 10, want: false},
		{name: "empty string and valid num", num: "", base: 10, want: true},
		{name: "valid number with the invalid base", num: "000", base: 0, want: false},
		{name: "valid number with the invalid base", num: "000", base: 1, want: false}, // i found only theories about base 0 or 1, but noone use them in any practical case, so i think false here will be better
	} {
		if got := isValidNumber(tc.num, tc.base); got != tc.want {
			t.Errorf("isValidNumber(%v, %v) = (%v), want = (%v)", tc.num, tc.base, got, tc.want)
		}
	}
}
