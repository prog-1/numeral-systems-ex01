package main

import (
	"testing"
)

func TestIsValidNumber(t *testing.T) {
	for _, tc := range []struct {
		num  string
		base int
		want bool
	}{
		{"", 0, false},
		{"1", 2, true},
		{"77", 6, false},
		{"34545", 39, false},
		{"32178BC1", -195, false},
		{"1001101010100110", 2, true},
		{"221101010", 10, true},
		{"72882299020", 4, false},
		{"AAAA3526832323256", 13, true},
		{"WXYZ", 7, false},
		{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 36, true},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := isValidNumber(tc.num, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
