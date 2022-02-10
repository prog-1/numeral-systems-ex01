package main

import "testing"

func TestIsValidNumber(t *testing.T) {
	tests := []struct {
		name string
		num  string
		base int
		want bool
	}{
		{
			name: "example from exercise",
			num:  "123",
			base: 10,
			want: true,
		},
		{
			name: "example from exercise 2",
			num:  "123",
			base: 3,
			want: false,
		},
		{
			name: "not numbers",
			num:  ";:/",
			base: 10,
			want: false,
		},

		{
			name: "base biggers than 10",
			num:  "ABDC",
			base: 16,
			want: true,
		},
	}
	for _, a := range tests {
		t.Run(a.num, func(t *testing.T) {
			if got := IsValidNumber(a.num, a.base); got != a.want {
				t.Errorf("%s isValidNumber() = %v, want %v", a.name, got, a.want)
			}
		})
	}
}
