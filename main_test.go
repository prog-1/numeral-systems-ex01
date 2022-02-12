package main

import (
	"testing"
)

func TestMaxAwardsperLetreture(t *testing.T) {
	for _, tc := range []struct {
		name     string
		number   string
		base     int
		isNumber bool
	}{
		{"Binary number", "1010110010101", 2, true},
		{"Non-binary number", "1010110210101", 2, false},
		{"Decimal number", "0123456789", 10, true},
		{"Non-decimal number", "0123456789ABCDEF", 10, false},
		{"HEX number", "0123456789ABCDEF", 16, true},
		{"Non-HEX number", "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 16, false},
		{"base 36 number", "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 36, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidNumber(tc.number, tc.base); got != tc.isNumber {
				t.Errorf("isValidNumber(%v, %v) = %v, want = %v", tc.number, tc.base, got, tc.isNumber)
			}
		})
	}
}
