package main

import "testing"

func TestIsValidNumber(t *testing.T) {
	for _, tc := range []struct {
		testnum int
		num     string
		base    int
		want    bool
	}{
		{testnum: 1, num: "", base: 0, want: true},
		{testnum: 2, num: "123", base: 0, want: false},
		{testnum: 3, num: "", base: 10, want: true},
		{testnum: 4, num: "123", base: 10, want: true},
		{testnum: 5, num: "123", base: 3, want: false},
		{testnum: 6, num: "ABC", base: 16, want: true},
		{testnum: 7, num: "AZ", base: 36, want: true},
		{testnum: 8, num: "1A", base: 1, want: false},
		{testnum: 9, num: "12xd", base: 36, want: false},
	} {
		if got := IsValidNumber(tc.num, tc.base); got != tc.want {
			t.Errorf("Tets number %v fail: got = %v, want = %v", tc.testnum, got, tc.want)
		}
	}
}
