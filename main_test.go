package main

import "testing"

func TestIsValidNumber(t *testing.T) {
	for _, tc := range []struct {
		name          string
		number        string
		base          int
		isValidNumber bool
	}{
		{"Empty", "", 1, true},
		// If number is an empty string, then the number does not exist(it is emptiness), and the emptiness belongs to any base
		{"Binary number", "010111001", 2, true},
		{"Non-binary number", "1102001", 2, false},
		{"Ternary number", "0110201", 3, true},
		{"Non-ternary number", "012001A", 3, false},
		{"Decimal number", "142980", 10, true},
		{"Non-decimal number", "19542B", 10, false},
		{"Hexadecimal number", "14673ACE", 16, true},
		{"Non-hexadecimal number", "11427MN", 16, false},
		{"Vigesimal number", "F853C21GE", 20, true},
		{"Non-vigesimal number", "QSL5847332DI", 20, false},
		{"Hexatrigesimal number", "MNO3749XYZ2235478VU", 36, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidNumber(tc.number, tc.base)
			if got != tc.isValidNumber {
				t.Errorf("ERR: TestIsValidNumber(%s): got: %v, want: %v", tc.name, got, tc.isValidNumber)
			}
		})
	}
}
