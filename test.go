package main

import (
	"fmt"
	"testing"
)

func TestValidNumber(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  string
		base int
		want bool
	}{
		{"empty num", "", 4, true},
		{"simple false", "1234", 4, false},
		{"simple true", "5612896", 10, true},
		{"simple true with letters", "1354ADSFA", 10, true},
		{"simple false with letters", "120123ADGD", 4, false},
		{"wrong base", "89623496", 4, false},
		{"only letters,true", "DFSDGFHJH", 10, true},
		{"only letters,false", "ARFGSRHYK", 4, false},
	} {
		if got := ValidNumber(tc.num, tc.base); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.num, tc.base, got, tc.want)

		}
	}
	if ok {
		fmt.Println("PASS")
	}
}
