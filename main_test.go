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
		{"3", 2, false},
		{"12", -1, false},
		{"123", 10, true},
		{"123", 3, false},
		{"153264221", 7, true},
		{"356", 10, true},
		{"ABCD", 14, true},
		{"KQUXCD", 9, false},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := isValidNumber(tc.num, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
