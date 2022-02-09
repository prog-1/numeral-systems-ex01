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
		{"simple", "1", 36, true},
		{"from README", "123", 10, true},
		{"from README", "123", 3, false},
		{"mixed", "9u", 36, true},
		{"only letters", "la", 36, true},
		{"random", "8439", 3, false},
		{"full dictionary", "0123456789abcdefghijklmnopqrstuvwxyz", 36, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidNumber(tc.num, tc.b); got != tc.answer {
				t.Errorf("isValidNumber(%v, %v) = %v, want = %v", tc.num, tc.b, got, tc.answer)
			}
		})
	}
}
