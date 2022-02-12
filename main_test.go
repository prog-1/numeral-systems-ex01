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
		{"", 1, false},
		{"", 16, false},
		{"1", 0, false},
		{"123", 0, false},
		{"0123459", -1, false},
		{"32178BC1", -195, false},
		{"123", 10, true},
		{"123", 3, false},
		{"123", 4, true},
		{"101010110011100001", 2, true},
		{"101010110011100001", 3, true},
		{"101010110011100001", 1, false},
		{"101020110011100001", 2, false},
		{"01A234C981EF", 16, true},
		{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 36, true},
		{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 37, false},
		{"123", 123, false},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := isValidNumber(tc.num, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
