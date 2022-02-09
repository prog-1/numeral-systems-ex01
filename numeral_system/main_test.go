package main

import (
	"testing"
)

func TestMaxAwardsperLetreture(t *testing.T) {
	for _, tc := range []struct {
		name   string
		num    string
		b      int
		answer bool
	}{
		{"from README", "1234", 10, true},
		{"from README", "123", 3, false},
		{"mixed", "9U", 36, true},
		{"only letters", "AF87", 16, true},
		{"random", "1015", 2, false},
		{"full dictionary", "0123456789abcdefghijklmnopqrstuvwxyz", 36, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidNumber(tc.num, tc.b); got != tc.answer {
				t.Errorf("isValidNumber(%v, %v) = %v, want = %v", tc.num, tc.b, got, tc.answer)
			}
		})
	}
}
